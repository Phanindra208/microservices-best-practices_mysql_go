// Code generated by go-swagger; DO NOT EDIT.

package apimodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Student Student structure
//
// swagger:model student
type Student struct {

	// student id
	StudentID string `json:"student_id,omitempty"`

	// student email
	StudentEmail string `json:"student_email,omitempty"`

	// student status
	StudentStatus string `json:"status,omitempty"`
}

// Validate validates this student
func (m *Student) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Student) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Student) UnmarshalBinary(b []byte) error {
	var res Student
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
