// Code generated by go-swagger; DO NOT EDIT.

package student

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/iAmPlus/microservice/models/apimodels"
)

// NewCreateRegisterParams creates a new CreateRegisterParams object
// no default values defined in spec.
func NewCreateRegisterParams() CreateRegisterParams {

	return CreateRegisterParams{}
}

// CreateRegisterParams contains all the bound params for the create register operation
// typically these are obtained from a http.Request
//
// swagger:parameters Create_Register
type CreateRegisterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Payload *apimodels.Register
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateRegisterParams() beforehand.
func (o *CreateRegisterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body apimodels.Register
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("payload", "body", ""))
			} else {
				res = append(res, errors.NewParseError("payload", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Payload = &body
			}
		}
	} else {
		res = append(res, errors.Required("payload", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
