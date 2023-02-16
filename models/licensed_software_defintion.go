// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LicensedSoftwareDefintion licensed software defintion
//
// swagger:model licensed_software_defintion
type LicensedSoftwareDefintion struct {

	// compare type
	// Enum: [like is]
	CompareType string `json:"compare_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	// Example: 14
	Version string `json:"version,omitempty"`
}

// Validate validates this licensed software defintion
func (m *LicensedSoftwareDefintion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompareType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var licensedSoftwareDefintionTypeCompareTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["like","is"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		licensedSoftwareDefintionTypeCompareTypePropEnum = append(licensedSoftwareDefintionTypeCompareTypePropEnum, v)
	}
}

const (

	// LicensedSoftwareDefintionCompareTypeLike captures enum value "like"
	LicensedSoftwareDefintionCompareTypeLike string = "like"

	// LicensedSoftwareDefintionCompareTypeIs captures enum value "is"
	LicensedSoftwareDefintionCompareTypeIs string = "is"
)

// prop value enum
func (m *LicensedSoftwareDefintion) validateCompareTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, licensedSoftwareDefintionTypeCompareTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LicensedSoftwareDefintion) validateCompareType(formats strfmt.Registry) error {
	if swag.IsZero(m.CompareType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompareTypeEnum("compare_type", "body", m.CompareType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this licensed software defintion based on context it is used
func (m *LicensedSoftwareDefintion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicensedSoftwareDefintion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicensedSoftwareDefintion) UnmarshalBinary(b []byte) error {
	var res LicensedSoftwareDefintion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
