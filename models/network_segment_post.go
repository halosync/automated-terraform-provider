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

// NetworkSegmentPost network segment post
//
// swagger:model network_segment_post
type NetworkSegmentPost struct {

	// ending address
	// Example: 10.10.1.1
	// Required: true
	EndingAddress *string `json:"ending_address"`

	// id
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
}

// Validate validates this network segment post
func (m *NetworkSegmentPost) Validate(formats strfmt.Registry) error {
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

func (m *NetworkSegmentPost) validateEndingAddress(formats strfmt.Registry) error {

	if err := validate.Required("ending_address", "body", m.EndingAddress); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSegmentPost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSegmentPost) validateStartingAddress(formats strfmt.Registry) error {

	if err := validate.Required("starting_address", "body", m.StartingAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network segment post based on context it is used
func (m *NetworkSegmentPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSegmentPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSegmentPost) UnmarshalBinary(b []byte) error {
	var res NetworkSegmentPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
