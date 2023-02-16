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

// Classes classes
//
// swagger:model classes
type Classes []*ClassesItems0

// Validate validates this classes
func (m Classes) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this classes based on the context it is used
func (m Classes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// ClassesItems0 classes items0
//
// swagger:model ClassesItems0
type ClassesItems0 struct {

	// class
	Class *ClassesItems0Class `json:"class,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this classes items0
func (m *ClassesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClass(formats); err != nil {
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

func (m *ClassesItems0) validateClass(formats strfmt.Registry) error {
	if swag.IsZero(m.Class) { // not required
		return nil
	}

	if m.Class != nil {
		if err := m.Class.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("class")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("class")
			}
			return err
		}
	}

	return nil
}

func (m *ClassesItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this classes items0 based on the context it is used
func (m *ClassesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClass(ctx, formats); err != nil {
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

func (m *ClassesItems0) contextValidateClass(ctx context.Context, formats strfmt.Registry) error {

	if m.Class != nil {
		if err := m.Class.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("class")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("class")
			}
			return err
		}
	}

	return nil
}

func (m *ClassesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ClassesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassesItems0) UnmarshalBinary(b []byte) error {
	var res ClassesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassesItems0Class classes items0 class
//
// swagger:model ClassesItems0Class
type ClassesItems0Class struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the class
	// Example: Biology 101
	Name string `json:"name,omitempty"`
}

// Validate validates this classes items0 class
func (m *ClassesItems0Class) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this classes items0 class based on context it is used
func (m *ClassesItems0Class) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassesItems0Class) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassesItems0Class) UnmarshalBinary(b []byte) error {
	var res ClassesItems0Class
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
