// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceOrder service order
// swagger:model service_order
type ServiceOrder struct {

	// Service catalog ID.
	// Required: true
	ServiceID *string `json:"service_id"`

	// Input variables for a service installation (if required).
	Variables interface{} `json:"variables,omitempty"`
}

// Validate validates this service order
func (m *ServiceOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOrder) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("service_id", "body", m.ServiceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOrder) UnmarshalBinary(b []byte) error {
	var res ServiceOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}