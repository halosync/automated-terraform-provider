// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActivationCode activation code
//
// swagger:model activation_code
type ActivationCode struct {

	// code
	// Example: QAPY-BP89-PJEH-87P8-0P89-PQ0A-FT68-00QA
	Code string `json:"code,omitempty"`

	// organization name
	// Example: Jamf Software LLC
	OrganizationName string `json:"organization_name,omitempty"`
}

// Validate validates this activation code
func (m *ActivationCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this activation code based on context it is used
func (m *ActivationCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivationCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivationCode) UnmarshalBinary(b []byte) error {
	var res ActivationCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
