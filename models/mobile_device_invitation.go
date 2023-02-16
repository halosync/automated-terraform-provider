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

// MobileDeviceInvitation mobile device invitation
//
// swagger:model mobile_device_invitation
type MobileDeviceInvitation struct {

	// date sent
	// Example: 2012-05-07 11:13:35
	// Read Only: true
	DateSent string `json:"date_sent,omitempty"`

	// date sent epoch
	// Example: 1336407215609
	// Read Only: true
	DateSentEpoch int64 `json:"date_sent_epoch,omitempty"`

	// date sent utc
	// Example: 2012-05-07T11:13:35.609-0500
	// Read Only: true
	DateSentUtc string `json:"date_sent_utc,omitempty"`

	// enrolled into site
	EnrolledIntoSite *MobileDeviceInvitationEnrolledIntoSite `json:"enrolled_into_site,omitempty"`

	// Use 'Unlimited' to specify no expiration
	// Example: 2012-05-07 11:13:35
	ExpirationDate *string `json:"expiration_date,omitempty"`

	// expiration date epoch
	// Example: 1336407215609
	// Read Only: true
	ExpirationDateEpoch int64 `json:"expiration_date_epoch,omitempty"`

	// expiration date utc
	// Example: 2012-05-07T11:13:35.609-0500
	// Read Only: true
	ExpirationDateUtc string `json:"expiration_date_utc,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// invitation
	// Example: 7.249663721735319e+37
	// Read Only: true
	Invitation int64 `json:"invitation,omitempty"`

	// invitation type
	// Example: USER_INITATIED_EMAIL
	InvitationType *string `json:"invitation_type,omitempty"`

	// keep existing site membership
	KeepExistingSiteMembership *bool `json:"keep_existing_site_membership,omitempty"`

	// last action
	// Example: NONE
	// Read Only: true
	LastAction string `json:"last_action,omitempty"`

	// login required
	LoginRequired *bool `json:"login_required,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// multiple uses allowed
	MultipleUsesAllowed *bool `json:"multiple_uses_allowed,omitempty"`

	// reply to
	// Example: it@company.com
	// Read Only: true
	ReplyTo string `json:"reply_to,omitempty"`

	// sent from
	// Example: Jamf Pro
	// Read Only: true
	SentFrom string `json:"sent_from,omitempty"`

	// sent to
	// Example: user@company.com
	// Read Only: true
	SentTo string `json:"sent_to,omitempty"`

	// site
	Site *SiteObject `json:"site,omitempty"`

	// subject
	// Example: Enroll your device
	// Read Only: true
	Subject string `json:"subject,omitempty"`

	// target ios
	// Example: iOS 4
	TargetIos string `json:"target_ios,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this mobile device invitation
func (m *MobileDeviceInvitation) Validate(formats strfmt.Registry) error {
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

func (m *MobileDeviceInvitation) validateEnrolledIntoSite(formats strfmt.Registry) error {
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

func (m *MobileDeviceInvitation) validateSite(formats strfmt.Registry) error {
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

// ContextValidate validate this mobile device invitation based on the context it is used
func (m *MobileDeviceInvitation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateSent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateSentEpoch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateSentUtc(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateLastAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSentFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSentTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceInvitation) contextValidateDateSent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "date_sent", "body", string(m.DateSent)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateDateSentEpoch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "date_sent_epoch", "body", int64(m.DateSentEpoch)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateDateSentUtc(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "date_sent_utc", "body", string(m.DateSentUtc)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateEnrolledIntoSite(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MobileDeviceInvitation) contextValidateExpirationDateEpoch(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiration_date_epoch", "body", int64(m.ExpirationDateEpoch)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateExpirationDateUtc(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "expiration_date_utc", "body", string(m.ExpirationDateUtc)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateInvitation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "invitation", "body", int64(m.Invitation)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateLastAction(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_action", "body", string(m.LastAction)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateReplyTo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reply_to", "body", string(m.ReplyTo)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateSentFrom(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sent_from", "body", string(m.SentFrom)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateSentTo(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sent_to", "body", string(m.SentTo)); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceInvitation) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MobileDeviceInvitation) contextValidateSubject(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "subject", "body", string(m.Subject)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceInvitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceInvitation) UnmarshalBinary(b []byte) error {
	var res MobileDeviceInvitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MobileDeviceInvitationEnrolledIntoSite mobile device invitation enrolled into site
//
// swagger:model MobileDeviceInvitationEnrolledIntoSite
type MobileDeviceInvitationEnrolledIntoSite struct {

	// id
	ID *int64 `json:"id,omitempty"`

	// name
	Name *string `json:"name,omitempty"`
}

// Validate validates this mobile device invitation enrolled into site
func (m *MobileDeviceInvitationEnrolledIntoSite) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mobile device invitation enrolled into site based on context it is used
func (m *MobileDeviceInvitationEnrolledIntoSite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceInvitationEnrolledIntoSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceInvitationEnrolledIntoSite) UnmarshalBinary(b []byte) error {
	var res MobileDeviceInvitationEnrolledIntoSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
