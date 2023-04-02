package teacherService

import (
	"fmt"
	"testing"

	"github.com/iAmPlus/microservice/models/apimodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockDatabase is a mock implementation of the Database interface
type mockDatabase struct {
	mock.Mock
}

func (md *mockDatabase) SuspendStudent(teacher_ID string, suspend_students apimodels.SuspendStudents) error {
	args := md.Called(teacher_ID, suspend_students)
	return args.Error(0)
}

func (md *mockDatabase) Retrievefornotifications(retrieve_for_notifications apimodels.RetrieveForNotifications) (apimodels.Recipients, error) {
	args := md.Called(retrieve_for_notifications)
	return args.Get(0).(apimodels.Recipients), args.Error(1)
}

func TestSuspendStudent(t *testing.T) {
	md := &mockDatabase{}
	service := New(md)

	// Test for successful suspension
	md.On("SuspendStudent", "1", mock.Anything).Return(nil)
	err := service.SuspendStudent("1", apimodels.SuspendStudents{})
	assert.Nil(t, err)
	md.AssertExpectations(t)

	// Test for failed suspension
	md.On("SuspendStudent", "1", mock.Anything).Return(fmt.Errorf("Some error"))
	err = service.SuspendStudent("1", apimodels.SuspendStudents{})
	assert.Nil(t, err)
	md.AssertExpectations(t)
}

func TestRetrieveForNotifications(t *testing.T) {
	md := &mockDatabase{}
	service := New(md)

	// Test for successful retrieval
	md.On("Retrievefornotifications", mock.Anything).Return(apimodels.Recipients{}, nil)
	students, err := service.Retrievefornotifications(apimodels.RetrieveForNotifications{})
	assert.Nil(t, err)
	assert.NotNil(t, students)
	md.AssertExpectations(t)

	// Test for failed retrieval
	md.On("Retrievefornotifications", mock.Anything).Return(apimodels.Recipients{}, fmt.Errorf("Some error"))
	students, err = service.Retrievefornotifications(apimodels.RetrieveForNotifications{})
	assert.Nil(t, err)
	assert.NotNil(t, students)
	md.AssertExpectations(t)
}
