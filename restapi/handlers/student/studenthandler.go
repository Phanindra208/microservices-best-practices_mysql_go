package studenthandler

import (
	"github.com/iAmPlus/microservice/restapi/operations/student"
	"github.com/iAmPlus/microservice/restapi/responder"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserFrendsFeed
func Register(params student.CreateRegisterParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	err := studentService.Register(params.Payload.TeacherID, *params.Payload)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.NoContent().Created(201)

}

// GetUserFrendsFeed
func Getcommonstudents(params student.GetCommonStudentsParams) middleware.Responder {

	resp := responder.New(params.HTTPRequest)
	result, err := studentService.Getcommonstudents(*&params.TeacherID)
	if err != nil {
		return resp.Status(500).Error(500, err.Error())
	}
	return resp.OK(result)

}
