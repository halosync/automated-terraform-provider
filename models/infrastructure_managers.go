// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InfrastructureManagers infrastructure managers
//
// swagger:model infrastructure_managers
type InfrastructureManagers []*InfrastructureManagersItems0

// Validate validates this infrastructure managers
func (m InfrastructureManagers) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this infrastructure managers based on the context it is used
func (m InfrastructureManagers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// InfrastructureManagersItems0 infrastructure managers items0
//
// swagger:model InfrastructureManagersItems0
type InfrastructureManagersItems0 struct {

	// infrastructure manager
	InfrastructureManager *InfrastructureManager `json:"infrastructure_manager,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this infrastructure managers items0
func (m *InfrastructureManagersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfrastructureManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfrastructureManagersItems0) validateInfrastructureManager(formats strfmt.Registry) error {
	if swag.IsZero(m.InfrastructureManager) { // not required
		return nil
	}

	if m.InfrastructureManager != nil {
		if err := m.InfrastructureManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infrastructure_manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infrastructure_manager")
			}
			return err
		}
	}

	return nil
}

func (m *InfrastructureManagersItems0) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := m.Size.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("size")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("size")
		}
		return err
	}

	return nil
}

// ContextValidate validate this infrastructure managers items0 based on the context it is used
func (m *InfrastructureManagersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfrastructureManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfrastructureManagersItems0) contextValidateInfrastructureManager(ctx context.Context, formats strfmt.Registry) error {

	if m.InfrastructureManager != nil {
		if err := m.InfrastructureManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infrastructure_manager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infrastructure_manager")
			}
			return err
		}
	}

	return nil
}

func (m *InfrastructureManagersItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Size.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("size")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("size")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfrastructureManagersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfrastructureManagersItems0) UnmarshalBinary(b []byte) error {
	var res InfrastructureManagersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
