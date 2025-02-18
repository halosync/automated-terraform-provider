// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComputersBasic computers
//
// swagger:model computers_basic
type ComputersBasic struct {

	// computer
	Computer *ComputersBasicComputer `json:"computer,omitempty"`
}

// Validate validates this computers basic
func (m *ComputersBasic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputersBasic) validateComputer(formats strfmt.Registry) error {
	if swag.IsZero(m.Computer) { // not required
		return nil
	}

	if m.Computer != nil {
		if err := m.Computer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computers basic based on the context it is used
func (m *ComputersBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputersBasic) contextValidateComputer(ctx context.Context, formats strfmt.Registry) error {

	if m.Computer != nil {
		if err := m.Computer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputersBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputersBasic) UnmarshalBinary(b []byte) error {
	var res ComputersBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputersBasicComputer computers basic computer
//
// swagger:model ComputersBasicComputer
type ComputersBasicComputer struct {

	// building
	// Example: East
	Building string `json:"building,omitempty"`

	// department
	// Example: Accounting
	Department string `json:"department,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// mac address
	// Example: E0:AC:CB:97:36:G4
	MacAddress string `json:"mac_address,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// model
	// Example: 13-inch MacBook Pro (Mid 2016)
	Model string `json:"model,omitempty"`

	// Name of the computer
	// Example: Lauras MacBook Pro
	Name string `json:"name,omitempty"`

	// report date epoch
	// Example: 1499470624555
	ReportDateEpoch string `json:"report_date_epoch,omitempty"`

	// report date utc
	// Example: 2017-07-07T18:37:04.555-0500
	ReportDateUtc string `json:"report_date_utc,omitempty"`

	// serial number
	// Example: C02Q7KHTGFWF
	SerialNumber string `json:"serial_number,omitempty"`

	// udid
	// Example: 55900BDC-347C-58B1-D249-F32244B11D30
	Udid string `json:"udid,omitempty"`

	// username
	// Example: Laura
	Username string `json:"username,omitempty"`
}

// Validate validates this computers basic computer
func (m *ComputersBasicComputer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computers basic computer based on context it is used
func (m *ComputersBasicComputer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputersBasicComputer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputersBasicComputer) UnmarshalBinary(b []byte) error {
	var res ComputersBasicComputer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
