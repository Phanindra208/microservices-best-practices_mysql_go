package teacherService

import (
	"github.com/iAmPlus/microservice/models/apimodels"
)

// GetUserFeed
// GetUserFrendsFeed
func (ls *teacherService) SuspendStudent(teacher_ID string, suspend_students apimodels.SuspendStudents) error {

	return ls.db.SuspendStudent(teacher_ID, suspend_students)

}

// GetUserFrendsFeed
func (ls *teacherService) Retrievefornotifications(retrieve_for_notifications apimodels.RetrieveForNotifications) (apimodels.Recipients, error) {

	return ls.db.Retrievefornotifications(retrieve_for_notifications)

}
