package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/iAmPlus/microservice/models/apimodels"
)

var (
	errMissingDBCredentials = fmt.Errorf("Database Credentials Missing")
	maxDuration             = 100 * time.Second
	indexDuration           = 4 * time.Minute
)

func NewManager(DBUrl string, DBname string) (Manager, error) {
	dbOpts := CreateOptsFromConfig(DBUrl, DBname)
	if dbOpts.DatabaseURI == "" || dbOpts.DatabaseName == "" {
		return nil, errMissingDBCredentials
	}
	db, err := sql.Open("mysql", DBUrl+"/"+DBname)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Unable to create MYSQL DB client, URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}

	um := &manager{dbOpts: dbOpts, client: db}
	err = um.initiateTables()
	if err != nil {
		return nil, fmt.Errorf("Unable to create MYSQL DB teacher manager URL %s, opts : %v : %v", dbOpts.DatabaseURI, dbOpts, err)
	}

	return um, nil
}
func CreateOptsFromConfig(DBUrl string, DBname string) DatabaseOpts {
	return DatabaseOpts{
		DatabaseURI:       DBUrl,
		DatabaseName:      DBname,
		Users_coll:        "users",
		Posts_coll:        "posts",
		Related_post_coll: "related",
	}
}

type Opts struct {
	// time in seconds for considering dialogue records for context.
	RecentDialogueThreshold float64
	// time in seconds for considering dialogue continuation
	NewDialogueThreshold float64
}
type manager struct {
	dbOpts DatabaseOpts
	client *sql.DB
}

// DatabaseOpts options for passing database details
type DatabaseOpts struct {
	// Uri of databse to connect to
	DatabaseURI string
	// Name of database to connect
	DatabaseName      string
	Users_coll        string
	Posts_coll        string
	Related_post_coll string
}

func (um *manager) initiateTables() error {
	var err error
	err = um.createTeacherTable()
	if err != nil {
		return err
	}
	err = um.createStudentTable()
	if err != nil {
		return err
	}
	return nil
}

func (u *manager) createStudentTable() error {
	query := `CREATE TABLE IF NOT EXISTS Student(id INT PRIMARY KEY AUTO_INCREMENT,student_id varchar(50), student_email_id text,student_status text,teacher_id varchar(50),FOREIGN KEY (teacher_id) REFERENCES Teacher(teacher_id))`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := u.client.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
func (u *manager) createTeacherTable() error {
	query := `CREATE TABLE IF NOT EXISTS Teacher(teacher_id varchar(50) PRIMARY KEY, teacher_email_id text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := u.client.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func (u *manager) Register(teacher_ID string, register apimodels.Register) error {
	// Prepare the INSERT statement
	stmt, err := u.client.Prepare("INSERT IGNORE INTO Teacher (teacher_id,teacher_email_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	// Insert a row into the table
	result, err := stmt.Exec(register.TeacherID, register.TeacherID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted row with ID", id)
	for _, v := range register.Students {
		stmt, err := u.client.Prepare("INSERT IGNORE INTO Student (student_id, student_email_id, student_status,teacher_id) VALUES (?, ?,?,?)")
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer stmt.Close()

		// Insert a row into the table
		result, err := stmt.Exec(v.StudentID, v.StudentEmail, "Active", register.TeacherID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Get the ID of the inserted row
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Inserted row with ID", id)

	}

	return nil
}
func (u *manager) Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error) {

	interfaceValues := make([]interface{}, len(teacher_ID))
	for i, v := range teacher_ID {
		interfaceValues[i] = v
	}

	// Create the placeholder string
	placeholders := make([]string, len(teacher_ID))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	placeholderString := "(" + strings.Join(placeholders, ",") + ")"

	rows, err := u.client.Query(`
        SELECT Student_ID,Student_Email_id,Student_Status FROM Student WHERE teacher_id IN`+placeholderString, interfaceValues...)
	if err != nil {
		fmt.Println(err)
		return apimodels.CommonStudents{}, err
	}
	defer rows.Close()
	var students []*apimodels.Student
	for rows.Next() {
		// Create a new User struct
		var student apimodels.Student

		// Scan the columns into the fields of the User struct
		if err := rows.Scan(&student.StudentID, &student.StudentEmail, &student.StudentStatus); err != nil {
			fmt.Println(err)
			return apimodels.CommonStudents{}, err
		}

		// Print the User struct
		students = append(students, &student)

	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return apimodels.CommonStudents{}, err
	}
	com := apimodels.CommonStudents{}
	com.Students = students
	return com, nil
}
func (u *manager) SuspendStudent(teacher_ID string, student apimodels.SuspendStudents) error {
	// Execute the UPDATE statement
	result, err := u.client.Exec(`
	 UPDATE Student SET Student_Status = 'Suspend' WHERE Student_ID = ?`, student.StudentID)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get the number of rows affected by the UPDATE statement
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Print the number of rows affected
	fmt.Println(rowsAffected, "rows affected.")
	return nil
}
func (u *manager) Retrievefornotifications(teacher_ID apimodels.RetrieveForNotifications) (apimodels.Recipients, error) {
	rows, err := u.client.Query(`
        SELECT Student_ID,Student_Email_id,Student_Status FROM Student WHERE teacher_id =?`, teacher_ID.TeacherID)
	if err != nil {
		fmt.Println(err)
		return apimodels.Recipients{}, err
	}
	defer rows.Close()
	var students []*apimodels.Student
	for rows.Next() {
		// Create a new User struct
		var student apimodels.Student

		// Scan the columns into the fields of the User struct
		if err := rows.Scan(&student.StudentID, &student.StudentEmail, &student.StudentStatus); err != nil {
			fmt.Println(err)
			return apimodels.Recipients{}, err
		}

		// Print the User struct
		students = append(students, &student)

	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return apimodels.Recipients{}, err
	}
	com := apimodels.Recipients{}
	com.Students = students
	return com, nil
}
