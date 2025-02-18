// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
//
// swagger:model account
type Account struct {

	// access level
	// Enum: [Full Access Site Access Group Access]
	AccessLevel string `json:"access_level,omitempty"`

	// directory user
	DirectoryUser bool `json:"directory_user,omitempty"`

	// email
	// Example: john.smith@company.com
	Email string `json:"email,omitempty"`

	// email address
	// Example: john.smith@company.com
	EmailAddress string `json:"email_address,omitempty"`

	// enabled
	// Enum: [Enabled Disabled]
	Enabled string `json:"enabled,omitempty"`

	// force password change
	ForcePasswordChange bool `json:"force_password_change,omitempty"`

	// full name
	// Example: John Smith
	FullName string `json:"full_name,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// ldap server
	LdapServer *AccountLdapServer `json:"ldap_server,omitempty"`

	// Name of the account
	// Example: John Smith
	// Required: true
	Name *string `json:"name"`

	// privilege set
	// Enum: [Administrator Auditor Enrollment Only Custom]
	PrivilegeSet string `json:"privilege_set,omitempty"`

	// privileges
	Privileges *AccountPrivileges `json:"privileges,omitempty"`

	// site
	Site *SiteObject `json:"site,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdapServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilegeSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileges(formats); err != nil {
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

var accountTypeAccessLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Full Access","Site Access","Group Access"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeAccessLevelPropEnum = append(accountTypeAccessLevelPropEnum, v)
	}
}

const (

	// AccountAccessLevelFullAccess captures enum value "Full Access"
	AccountAccessLevelFullAccess string = "Full Access"

	// AccountAccessLevelSiteAccess captures enum value "Site Access"
	AccountAccessLevelSiteAccess string = "Site Access"

	// AccountAccessLevelGroupAccess captures enum value "Group Access"
	AccountAccessLevelGroupAccess string = "Group Access"
)

// prop value enum
func (m *Account) validateAccessLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountTypeAccessLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateAccessLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessLevelEnum("access_level", "body", m.AccessLevel); err != nil {
		return err
	}

	return nil
}

var accountTypeEnabledPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enabled","Disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeEnabledPropEnum = append(accountTypeEnabledPropEnum, v)
	}
}

const (

	// AccountEnabledEnabled captures enum value "Enabled"
	AccountEnabledEnabled string = "Enabled"

	// AccountEnabledDisabled captures enum value "Disabled"
	AccountEnabledDisabled string = "Disabled"
)

// prop value enum
func (m *Account) validateEnabledEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountTypeEnabledPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateEnabled(formats strfmt.Registry) error {
	if swag.IsZero(m.Enabled) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnabledEnum("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateLdapServer(formats strfmt.Registry) error {
	if swag.IsZero(m.LdapServer) { // not required
		return nil
	}

	if m.LdapServer != nil {
		if err := m.LdapServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldap_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldap_server")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var accountTypePrivilegeSetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Administrator","Auditor","Enrollment Only","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypePrivilegeSetPropEnum = append(accountTypePrivilegeSetPropEnum, v)
	}
}

const (

	// AccountPrivilegeSetAdministrator captures enum value "Administrator"
	AccountPrivilegeSetAdministrator string = "Administrator"

	// AccountPrivilegeSetAuditor captures enum value "Auditor"
	AccountPrivilegeSetAuditor string = "Auditor"

	// AccountPrivilegeSetEnrollmentOnly captures enum value "Enrollment Only"
	AccountPrivilegeSetEnrollmentOnly string = "Enrollment Only"

	// AccountPrivilegeSetCustom captures enum value "Custom"
	AccountPrivilegeSetCustom string = "Custom"
)

// prop value enum
func (m *Account) validatePrivilegeSetEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountTypePrivilegeSetPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Account) validatePrivilegeSet(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivilegeSet) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivilegeSetEnum("privilege_set", "body", m.PrivilegeSet); err != nil {
		return err
	}

	return nil
}

func (m *Account) validatePrivileges(formats strfmt.Registry) error {
	if swag.IsZero(m.Privileges) { // not required
		return nil
	}

	if m.Privileges != nil {
		if err := m.Privileges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privileges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privileges")
			}
			return err
		}
	}

	return nil
}

func (m *Account) validateSite(formats strfmt.Registry) error {
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

// ContextValidate validate this account based on the context it is used
func (m *Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLdapServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivileges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) contextValidateLdapServer(ctx context.Context, formats strfmt.Registry) error {

	if m.LdapServer != nil {
		if err := m.LdapServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldap_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ldap_server")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidatePrivileges(ctx context.Context, formats strfmt.Registry) error {

	if m.Privileges != nil {
		if err := m.Privileges.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privileges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privileges")
			}
			return err
		}
	}

	return nil
}

func (m *Account) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountLdapServer account ldap server
//
// swagger:model AccountLdapServer
type AccountLdapServer struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Directory Server Name
	Name string `json:"name,omitempty"`
}

// Validate validates this account ldap server
func (m *AccountLdapServer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account ldap server based on context it is used
func (m *AccountLdapServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountLdapServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountLdapServer) UnmarshalBinary(b []byte) error {
	var res AccountLdapServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivileges account privileges
//
// swagger:model AccountPrivileges
type AccountPrivileges struct {

	// casper admin
	CasperAdmin []*AccountPrivilegesCasperAdminItems0 `json:"casper_admin"`

	// casper imaging
	CasperImaging []*AccountPrivilegesCasperImagingItems0 `json:"casper_imaging"`

	// casper remote
	CasperRemote []*AccountPrivilegesCasperRemoteItems0 `json:"casper_remote"`

	// jss actions
	JssActions []*AccountPrivilegesJssActionsItems0 `json:"jss_actions"`

	// jss objects
	JssObjects []*AccountPrivilegesJssObjectsItems0 `json:"jss_objects"`

	// jss settings
	JssSettings []*AccountPrivilegesJssSettingsItems0 `json:"jss_settings"`

	// recon
	Recon []*AccountPrivilegesReconItems0 `json:"recon"`
}

// Validate validates this account privileges
func (m *AccountPrivileges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCasperAdmin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCasperImaging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCasperRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJssActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJssObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJssSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPrivileges) validateCasperAdmin(formats strfmt.Registry) error {
	if swag.IsZero(m.CasperAdmin) { // not required
		return nil
	}

	for i := 0; i < len(m.CasperAdmin); i++ {
		if swag.IsZero(m.CasperAdmin[i]) { // not required
			continue
		}

		if m.CasperAdmin[i] != nil {
			if err := m.CasperAdmin[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_admin" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_admin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateCasperImaging(formats strfmt.Registry) error {
	if swag.IsZero(m.CasperImaging) { // not required
		return nil
	}

	for i := 0; i < len(m.CasperImaging); i++ {
		if swag.IsZero(m.CasperImaging[i]) { // not required
			continue
		}

		if m.CasperImaging[i] != nil {
			if err := m.CasperImaging[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_imaging" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_imaging" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateCasperRemote(formats strfmt.Registry) error {
	if swag.IsZero(m.CasperRemote) { // not required
		return nil
	}

	for i := 0; i < len(m.CasperRemote); i++ {
		if swag.IsZero(m.CasperRemote[i]) { // not required
			continue
		}

		if m.CasperRemote[i] != nil {
			if err := m.CasperRemote[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_remote" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_remote" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateJssActions(formats strfmt.Registry) error {
	if swag.IsZero(m.JssActions) { // not required
		return nil
	}

	for i := 0; i < len(m.JssActions); i++ {
		if swag.IsZero(m.JssActions[i]) { // not required
			continue
		}

		if m.JssActions[i] != nil {
			if err := m.JssActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateJssObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.JssObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.JssObjects); i++ {
		if swag.IsZero(m.JssObjects[i]) { // not required
			continue
		}

		if m.JssObjects[i] != nil {
			if err := m.JssObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateJssSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.JssSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.JssSettings); i++ {
		if swag.IsZero(m.JssSettings[i]) { // not required
			continue
		}

		if m.JssSettings[i] != nil {
			if err := m.JssSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) validateRecon(formats strfmt.Registry) error {
	if swag.IsZero(m.Recon) { // not required
		return nil
	}

	for i := 0; i < len(m.Recon); i++ {
		if swag.IsZero(m.Recon[i]) { // not required
			continue
		}

		if m.Recon[i] != nil {
			if err := m.Recon[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "recon" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "recon" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account privileges based on the context it is used
func (m *AccountPrivileges) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCasperAdmin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCasperImaging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCasperRemote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJssActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJssObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJssSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountPrivileges) contextValidateCasperAdmin(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CasperAdmin); i++ {

		if m.CasperAdmin[i] != nil {
			if err := m.CasperAdmin[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_admin" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_admin" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateCasperImaging(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CasperImaging); i++ {

		if m.CasperImaging[i] != nil {
			if err := m.CasperImaging[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_imaging" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_imaging" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateCasperRemote(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CasperRemote); i++ {

		if m.CasperRemote[i] != nil {
			if err := m.CasperRemote[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "casper_remote" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "casper_remote" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateJssActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JssActions); i++ {

		if m.JssActions[i] != nil {
			if err := m.JssActions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateJssObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JssObjects); i++ {

		if m.JssObjects[i] != nil {
			if err := m.JssObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateJssSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JssSettings); i++ {

		if m.JssSettings[i] != nil {
			if err := m.JssSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "jss_settings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "jss_settings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountPrivileges) contextValidateRecon(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recon); i++ {

		if m.Recon[i] != nil {
			if err := m.Recon[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + "recon" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("privileges" + "." + "recon" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivileges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivileges) UnmarshalBinary(b []byte) error {
	var res AccountPrivileges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesCasperAdminItems0 account privileges casper admin items0
//
// swagger:model AccountPrivilegesCasperAdminItems0
type AccountPrivilegesCasperAdminItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges casper admin items0
func (m *AccountPrivilegesCasperAdminItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges casper admin items0 based on context it is used
func (m *AccountPrivilegesCasperAdminItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesCasperAdminItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesCasperAdminItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesCasperAdminItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesCasperImagingItems0 account privileges casper imaging items0
//
// swagger:model AccountPrivilegesCasperImagingItems0
type AccountPrivilegesCasperImagingItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges casper imaging items0
func (m *AccountPrivilegesCasperImagingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges casper imaging items0 based on context it is used
func (m *AccountPrivilegesCasperImagingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesCasperImagingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesCasperImagingItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesCasperImagingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesCasperRemoteItems0 account privileges casper remote items0
//
// swagger:model AccountPrivilegesCasperRemoteItems0
type AccountPrivilegesCasperRemoteItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges casper remote items0
func (m *AccountPrivilegesCasperRemoteItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges casper remote items0 based on context it is used
func (m *AccountPrivilegesCasperRemoteItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesCasperRemoteItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesCasperRemoteItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesCasperRemoteItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesJssActionsItems0 account privileges jss actions items0
//
// swagger:model AccountPrivilegesJssActionsItems0
type AccountPrivilegesJssActionsItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges jss actions items0
func (m *AccountPrivilegesJssActionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges jss actions items0 based on context it is used
func (m *AccountPrivilegesJssActionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesJssActionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesJssActionsItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesJssActionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesJssObjectsItems0 account privileges jss objects items0
//
// swagger:model AccountPrivilegesJssObjectsItems0
type AccountPrivilegesJssObjectsItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges jss objects items0
func (m *AccountPrivilegesJssObjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges jss objects items0 based on context it is used
func (m *AccountPrivilegesJssObjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesJssObjectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesJssObjectsItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesJssObjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesJssSettingsItems0 account privileges jss settings items0
//
// swagger:model AccountPrivilegesJssSettingsItems0
type AccountPrivilegesJssSettingsItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges jss settings items0
func (m *AccountPrivilegesJssSettingsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges jss settings items0 based on context it is used
func (m *AccountPrivilegesJssSettingsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesJssSettingsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesJssSettingsItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesJssSettingsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountPrivilegesReconItems0 account privileges recon items0
//
// swagger:model AccountPrivilegesReconItems0
type AccountPrivilegesReconItems0 struct {

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this account privileges recon items0
func (m *AccountPrivilegesReconItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account privileges recon items0 based on context it is used
func (m *AccountPrivilegesReconItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountPrivilegesReconItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountPrivilegesReconItems0) UnmarshalBinary(b []byte) error {
	var res AccountPrivilegesReconItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
