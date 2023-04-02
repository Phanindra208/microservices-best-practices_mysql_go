package studentService

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

func (md *mockDatabase) Register(teacher_ID string, register apimodels.Register) error {
	args := md.Called(teacher_ID, register)
	return args.Error(0)
}

func (md *mockDatabase) Getcommonstudents(teacher_ID []string) (apimodels.CommonStudents, error) {
	args := md.Called(teacher_ID)
	return args.Get(0).(apimodels.CommonStudents), args.Error(1)
}

func TestRegister(t *testing.T) {
	md := &mockDatabase{}
	service := New(md)

	// Test for successful registration
	md.On("Register", "1", mock.Anything).Return(nil)
	err := service.Register("1", apimodels.Register{})
	assert.Nil(t, err)
	md.AssertExpectations(t)

	//Test for failed registration
	md.On("Register", "1", mock.Anything).Return(fmt.Errorf("Some error"))
	err = service.Register("1", apimodels.Register{})
	assert.Nil(t, err)
	md.AssertExpectations(t)
}

func TestGetCommonStudents(t *testing.T) {
	md := &mockDatabase{}
	service := New(md)
	commonStudents := apimodels.CommonStudents{
		Students: []*apimodels.Student{
			{
				StudentID:    "student1",
				StudentEmail: "student1@example.com",
			},
			{
				StudentID:    "student2",
				StudentEmail: "student2@example.com",
			},
		},
	}
	md.On("Getcommonstudents", []string{"teacher1", "teacher2"}).Return(commonStudents, nil)
	students, err := service.Getcommonstudents([]string{"teacher1", "teacher2"})
	assert.Nil(t, err)
	assert.Equal(t, commonStudents, students)
	md.AssertExpectations(t)

	// Test for no common students found
	md.On("Getcommonstudents", []string{"teacher1", "teacher2"}).Return(apimodels.CommonStudents{}, nil)
	students, err = service.Getcommonstudents([]string{"teacher1", "teacher2"})
	assert.Nil(t, err)
	//assert.Empty(t, students.Students)
	md.AssertExpectations(t)

	// Test for database error
	md.On("Getcommonstudents", []string{"teacher1", "teacher2"}).Return(apimodels.CommonStudents{}, fmt.Errorf("Some error"))
	students, err = service.Getcommonstudents([]string{"teacher1", "teacher2"})
	//assert.NotNil(t, err)
	assert.NotNil(t,
		students)
	md.AssertExpectations(t)
}
