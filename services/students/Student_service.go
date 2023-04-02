package studentService

import "github.com/iAmPlus/microservice/models/apimodels"

type Database interface {
	Register(teacher_ID string, register apimodels.Register) error
	Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error)
}

type studentService struct {
	db Database
}

// New creates new Feed Service.
func New(m Database) *studentService {
	return &studentService{m}
}
