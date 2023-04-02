package teacherService

import "github.com/iAmPlus/microservice/models/apimodels"

type Database interface {
	SuspendStudent(teacher_ID string, Pagestate apimodels.SuspendStudents) error
	Retrievefornotifications(teacher_ID apimodels.RetrieveForNotifications) (apimodels.Recipients, error)
}

type teacherService struct {
	db Database
}

// New creates new Feed Service.
func New(m Database) *teacherService {
	return &teacherService{m}
}
