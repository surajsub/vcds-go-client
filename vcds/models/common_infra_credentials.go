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

// CommonInfraCredentials IBM Cloud infrastructure API credentials
// swagger:model common_infra_credentials
type CommonInfraCredentials struct {

	// IBM Cloud infrastructure account ID
	AccountID string `json:"account_id,omitempty"`

	// IBM Cloud infrastructure API key
	APIKey string `json:"api_key,omitempty"`

	// IBM Cloud infrastructure API user name
	// Required: true
	UserName *string `json:"user_name"`
}

// Validate validates this common infra credentials
func (m *CommonInfraCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonInfraCredentials) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonInfraCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonInfraCredentials) UnmarshalBinary(b []byte) error {
	var res CommonInfraCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
