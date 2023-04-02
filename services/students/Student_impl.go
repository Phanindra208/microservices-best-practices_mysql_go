package studentService

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// GetUserFeed
func (ls *studentService) Register(teacher_ID string, register apimodels.Register) error {

	return ls.db.Register(teacher_ID, register)

}

// GetUserFrendsFeed
func (ls *studentService) Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error) {

	return ls.db.Getcommonstudents(teacher_ID)

}
