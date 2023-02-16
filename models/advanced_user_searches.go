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

// AdvancedUserSearches advanced user searches
//
// swagger:model advanced_user_searches
type AdvancedUserSearches []*AdvancedUserSearchesItems0

// Validate validates this advanced user searches
func (m AdvancedUserSearches) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this advanced user searches based on the context it is used
func (m AdvancedUserSearches) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// AdvancedUserSearchesItems0 advanced user searches items0
//
// swagger:model AdvancedUserSearchesItems0
type AdvancedUserSearchesItems0 struct {

	// advanced user search
	AdvancedUserSearch *IDName `json:"advanced_user_search,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this advanced user searches items0
func (m *AdvancedUserSearchesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedUserSearch(formats); err != nil {
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

func (m *AdvancedUserSearchesItems0) validateAdvancedUserSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.AdvancedUserSearch) { // not required
		return nil
	}

	if m.AdvancedUserSearch != nil {
		if err := m.AdvancedUserSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced_user_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced_user_search")
			}
			return err
		}
	}

	return nil
}

func (m *AdvancedUserSearchesItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this advanced user searches items0 based on the context it is used
func (m *AdvancedUserSearchesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvancedUserSearch(ctx, formats); err != nil {
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

func (m *AdvancedUserSearchesItems0) contextValidateAdvancedUserSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.AdvancedUserSearch != nil {
		if err := m.AdvancedUserSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced_user_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("advanced_user_search")
			}
			return err
		}
	}

	return nil
}

func (m *AdvancedUserSearchesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AdvancedUserSearchesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvancedUserSearchesItems0) UnmarshalBinary(b []byte) error {
	var res AdvancedUserSearchesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
