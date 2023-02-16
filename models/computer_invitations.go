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

// ComputerInvitations computer invitations
//
// swagger:model computer_invitations
type ComputerInvitations []*ComputerInvitationsItems0

// Validate validates this computer invitations
func (m ComputerInvitations) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this computer invitations based on the context it is used
func (m ComputerInvitations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// ComputerInvitationsItems0 computer invitations items0
//
// swagger:model ComputerInvitationsItems0
type ComputerInvitationsItems0 struct {

	// computer invitation
	ComputerInvitation *ComputerInvitationsItems0ComputerInvitation `json:"computer_invitation,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this computer invitations items0
func (m *ComputerInvitationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputerInvitation(formats); err != nil {
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

func (m *ComputerInvitationsItems0) validateComputerInvitation(formats strfmt.Registry) error {
	if swag.IsZero(m.ComputerInvitation) { // not required
		return nil
	}

	if m.ComputerInvitation != nil {
		if err := m.ComputerInvitation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer_invitation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer_invitation")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerInvitationsItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this computer invitations items0 based on the context it is used
func (m *ComputerInvitationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputerInvitation(ctx, formats); err != nil {
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

func (m *ComputerInvitationsItems0) contextValidateComputerInvitation(ctx context.Context, formats strfmt.Registry) error {

	if m.ComputerInvitation != nil {
		if err := m.ComputerInvitation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer_invitation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer_invitation")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerInvitationsItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerInvitationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerInvitationsItems0) UnmarshalBinary(b []byte) error {
	var res ComputerInvitationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerInvitationsItems0ComputerInvitation computer invitations items0 computer invitation
//
// swagger:model ComputerInvitationsItems0ComputerInvitation
type ComputerInvitationsItems0ComputerInvitation struct {

	// expiration date
	// Example: 2012-05-07 11:13:35
	ExpirationDate string `json:"expiration_date,omitempty"`

	// expiration date epoch
	// Example: 1336407215609
	ExpirationDateEpoch int64 `json:"expiration_date_epoch,omitempty"`

	// expiration date utc
	// Example: 2012-05-07T11:13:35.609-0500
	ExpirationDateUtc string `json:"expiration_date_utc,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// invitation
	// Example: 7.249663721735319e+37
	Invitation int64 `json:"invitation,omitempty"`

	// invitation type
	// Example: USER_INITATIED_EMAIL
	InvitationType string `json:"invitation_type,omitempty"`
}

// Validate validates this computer invitations items0 computer invitation
func (m *ComputerInvitationsItems0ComputerInvitation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer invitations items0 computer invitation based on context it is used
func (m *ComputerInvitationsItems0ComputerInvitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerInvitationsItems0ComputerInvitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerInvitationsItems0ComputerInvitation) UnmarshalBinary(b []byte) error {
	var res ComputerInvitationsItems0ComputerInvitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
