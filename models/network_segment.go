// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkSegment network segment
//
// swagger:model network_segment
type NetworkSegment struct {

	// building
	Building string `json:"building,omitempty"`

	// department
	Department string `json:"department,omitempty"`

	// distribution point
	DistributionPoint string `json:"distribution_point,omitempty"`

	// distribution server
	DistributionServer string `json:"distribution_server,omitempty"`

	// ending address
	// Example: 10.10.1.1
	// Required: true
	EndingAddress *string `json:"ending_address"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the network segment
	// Example: Amsterdam Office
	// Required: true
	Name *string `json:"name"`

	// override buildings
	OverrideBuildings *bool `json:"override_buildings,omitempty"`

	// override departments
	OverrideDepartments *bool `json:"override_departments,omitempty"`

	// starting address
	// Example: 10.1.1.1
	// Required: true
	StartingAddress *string `json:"starting_address"`

	// swu server
	SwuServer string `json:"swu_server,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this network segment
func (m *NetworkSegment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartingAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSegment) validateEndingAddress(formats strfmt.Registry) error {

	if err := validate.Required("ending_address", "body", m.EndingAddress); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSegment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSegment) validateStartingAddress(formats strfmt.Registry) error {

	if err := validate.Required("starting_address", "body", m.StartingAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network segment based on context it is used
func (m *NetworkSegment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSegment) UnmarshalBinary(b []byte) error {
	var res NetworkSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
