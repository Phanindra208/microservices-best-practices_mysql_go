package teacherhandlers

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// PaymentService Process Payment.
type TeacherService interface {
	SuspendStudent(teacher_ID string, Pagestate apimodels.SuspendStudents) error
	Retrievefornotifications(notificationpayload apimodels.RetrieveForNotifications) (apimodels.Recipients, error)
}

var teacherService TeacherService

// Init initializes package.
func Init(ls TeacherService) {
	teacherService = ls
}
