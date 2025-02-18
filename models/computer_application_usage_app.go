// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComputerApplicationUsageApp computer application usage app
//
// swagger:model computer_application_usage_app
type ComputerApplicationUsageApp struct {

	// Number of minutes application was in the foreground
	// Example: 45
	Foreground int64 `json:"foreground,omitempty"`

	// name
	// Example: Safari.app
	Name string `json:"name,omitempty"`

	// Number of minutes the application was open
	// Example: 150
	Open int64 `json:"open,omitempty"`

	// version
	// Example: 11
	Version string `json:"version,omitempty"`
}

// Validate validates this computer application usage app
func (m *ComputerApplicationUsageApp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer application usage app based on context it is used
func (m *ComputerApplicationUsageApp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerApplicationUsageApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerApplicationUsageApp) UnmarshalBinary(b []byte) error {
	var res ComputerApplicationUsageApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
