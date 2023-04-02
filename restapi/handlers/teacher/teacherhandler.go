package teacherhandlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/iAmPlus/microservice/restapi/operations/teacher"
	"github.com/iAmPlus/microservice/restapi/responder"
)

func SuspendStudent(params teacher.SuspendStudentParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	err := teacherService.SuspendStudent(params.Payload.StudentID, *params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.NoContent().Created(201)

}

// GetUserFrendsFeed
func Retrievefornotifications(params teacher.RetrieveForNotificationsParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := teacherService.Retrievefornotifications(*params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}
