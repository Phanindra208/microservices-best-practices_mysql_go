package studenthandler

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// PaymentService Process Payment.
type StudentService interface {
	Register(teacher_ID string, register apimodels.Register) error
	Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error)
}

var studentService StudentService

// Init initializes package.
func Init(ls StudentService) {
	studentService = ls
}
