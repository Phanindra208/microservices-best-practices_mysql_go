package db

import (
	"time"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// CreateOptsFromConfig creates opts from config
func createOptsFromConfig() Opts {
	return Opts{
		RecentDialogueThreshold: 5 * 60.0,
		NewDialogueThreshold:    3 * 60 * 60.0,
	}
}

// SearchOpts for SearchHistory function.
type SearchOpts struct {
	UserID             string
	Query              string
	Intent             string
	Limit              int
	Offset             int
	IgnorePreformatted bool
	Start              time.Time
	End                time.Time
}

// Manager ... stores/retrieves dialogue records.
type Manager interface {
	Register(teacher_ID string, register apimodels.Register) error
	Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error)
	SuspendStudent(teacher_ID string, Pagestate apimodels.SuspendStudents) error
	Retrievefornotifications(teacher_ID apimodels.RetrieveForNotifications) (apimodels.Recipients, error)
}
