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

// LdapServers ldap servers
//
// swagger:model ldap_servers
type LdapServers []*LdapServersItems0

// Validate validates this ldap servers
func (m LdapServers) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this ldap servers based on the context it is used
func (m LdapServers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// LdapServersItems0 ldap servers items0
//
// swagger:model LdapServersItems0
type LdapServersItems0 struct {

	// ldap server
	LdapServer *IDName `json:"ldap_server,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this ldap servers items0
func (m *LdapServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLdapServer(formats); err != nil {
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

func (m *LdapServersItems0) validateLdapServer(formats strfmt.Registry) error {
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

func (m *LdapServersItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this ldap servers items0 based on the context it is used
func (m *LdapServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLdapServer(ctx, formats); err != nil {
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

func (m *LdapServersItems0) contextValidateLdapServer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LdapServersItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LdapServersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServersItems0) UnmarshalBinary(b []byte) error {
	var res LdapServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
