// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Witness Witness cluster configuration of the MCV.
// swagger:model witness
type Witness struct {

	// customized hardware
	CustomizedHardware *CustomizedHardware `json:"customized_hardware,omitempty"`

	// Number of hardware servers.
	// Required: true
	Quantity *int64 `json:"quantity"`

	// List of shared storage configurations.
	SharedStorages []*SharedStorageConfig `json:"shared_storages"`

	// The template ID that maps to a preconfigured hardware specification to be ordered. To see supported templates, use the `GET /v2/templates` API. To customize your hardware configuration, use `customized_hardware`. This will be ignored if `customized_hardware` is also specified.
	TemplateID string `json:"template_id,omitempty"`

	// Data center location. Pick one data center by using `GET /v2/mzr_locations` API.
	// Required: true
	WitnessLocation *string `json:"witness_location"`
}

// Validate validates this witness
func (m *Witness) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomizedHardware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedStorages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWitnessLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Witness) validateCustomizedHardware(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomizedHardware) { // not required
		return nil
	}

	if m.CustomizedHardware != nil {
		if err := m.CustomizedHardware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customized_hardware")
			}
			return err
		}
	}

	return nil
}

func (m *Witness) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateSharedStorages(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedStorages) { // not required
		return nil
	}

	for i := 0; i < len(m.SharedStorages); i++ {
		if swag.IsZero(m.SharedStorages[i]) { // not required
			continue
		}

		if m.SharedStorages[i] != nil {
			if err := m.SharedStorages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shared_storages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Witness) validateWitnessLocation(formats strfmt.Registry) error {

	if err := validate.Required("witness_location", "body", m.WitnessLocation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Witness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Witness) UnmarshalBinary(b []byte) error {
	var res Witness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}