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

// ComputerInvitation computer invitation
//
// swagger:model computer_invitation
type ComputerInvitation struct {

	// create account if does not exist
	CreateAccountIfDoesNotExist *bool `json:"create_account_if_does_not_exist,omitempty"`

	// enrolled into site
	EnrolledIntoSite *ComputerInvitationEnrolledIntoSite `json:"enrolled_into_site,omitempty"`

	// Use 'Unlimited' to specify no expiration date
	// Example: 2012-05-07 11:13:35
	ExpirationDate string `json:"expiration_date,omitempty"`

	// expiration date epoch
	// Example: 1336407215609
	// Read Only: true
	ExpirationDateEpoch int64 `json:"expiration_date_epoch,omitempty"`

	// expiration date utc
	// Example: 2012-05-07T11:13:35.609-0500
	// Read Only: true
	ExpirationDateUtc string `json:"expiration_date_utc,omitempty"`

	// hide account
	HideAccount *bool `json:"hide_account,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// invitation
	// Example: 7.249663721735319e+37
	// Read Only: true
	Invitation int64 `json:"invitation,omitempty"`

	// invitation status
	// Example: INVITATION_EXPIRED
	// Read Only: true
	InvitationStatus string `json:"invitation_status,omitempty"`

	// invitation type
	// Example: USER_INITATIED_EMAIL
	InvitationType string `json:"invitation_type,omitempty"`

	// invited user uuid
	// Example: B87E9AA8-C3DE-4034-821E-1B7D51FD4956
	// Read Only: true
	InvitedUserUUID string `json:"invited_user_uuid,omitempty"`

	// keep existing site membership
	KeepExistingSiteMembership *bool `json:"keep_existing_site_membership,omitempty"`

	// lock down ssh
	LockDownSSH *bool `json:"lock_down_ssh,omitempty"`

	// multiple users allowed
	MultipleUsersAllowed *bool `json:"multiple_users_allowed,omitempty"`

	// site
	Site *SiteObject `json:"site,omitempty"`

	// ssh password
	// Example: accountpassword
	SSHPassword string `json:"ssh_password,omitempty"`

	// ssh username
	// Example: jamfadmin
	SSHUsername string `json:"ssh_username,omitempty"`

	// times used
	// Example: 0
	// Read Only: true
	TimesUsed int64 `json:"times_used,omitempty"`
}

// Validate validates this computer invitation
func (m *ComputerInvitation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrolledIntoSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerInvitation) validateEnrolledIntoSite(formats strfmt.Registry) error {
	if swag.IsZero(m.EnrolledIntoSite) { // not required
		return nil
	}

	if m.EnrolledIntoSite != nil {
		if err := m.EnrolledIntoSite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrolled_into_site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enrolled_into_site")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerInvitation) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computer invitation based on the context it is used
func (m *ComputerInvitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnrolledIntoSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpirationDateEpoch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpirationDateUtc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvitation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvitationStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvitedUserUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimesUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerInvitation) contextValidateEnrolledIntoSite(ctx context.Context, formats strfmt.Registry) error {

	if m.EnrolledIntoSite != nil {
		if err := m.EnrolledIntoSite.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrolled_into_site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enrolled_into_site")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerInvitation) contextValidateExpirationDateEpoch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiration_date_epoch", "body", int64(m.ExpirationDateEpoch)); err != nil {
		return err
	}

	return nil
}

func (m *ComputerInvitation) contextValidateExpirationDateUtc(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiration_date_utc", "body", string(m.ExpirationDateUtc)); err != nil {
		return err
	}

	return nil
}

func (m *ComputerInvitation) contextValidateInvitation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "invitation", "body", int64(m.Invitation)); err != nil {
		return err
	}

	return nil
}

func (m *ComputerInvitation) contextValidateInvitationStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "invitation_status", "body", string(m.InvitationStatus)); err != nil {
		return err
	}

	return nil
}

func (m *ComputerInvitation) contextValidateInvitedUserUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "invited_user_uuid", "body", string(m.InvitedUserUUID)); err != nil {
		return err
	}

	return nil
}

func (m *ComputerInvitation) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {
		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerInvitation) contextValidateTimesUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "times_used", "body", int64(m.TimesUsed)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerInvitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerInvitation) UnmarshalBinary(b []byte) error {
	var res ComputerInvitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerInvitationEnrolledIntoSite computer invitation enrolled into site
//
// swagger:model ComputerInvitationEnrolledIntoSite
type ComputerInvitationEnrolledIntoSite struct {

	// id
	ID *int64 `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`
}

// Validate validates this computer invitation enrolled into site
func (m *ComputerInvitationEnrolledIntoSite) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer invitation enrolled into site based on context it is used
func (m *ComputerInvitationEnrolledIntoSite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerInvitationEnrolledIntoSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerInvitationEnrolledIntoSite) UnmarshalBinary(b []byte) error {
	var res ComputerInvitationEnrolledIntoSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
