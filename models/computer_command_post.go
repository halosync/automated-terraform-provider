// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComputerCommandPost computer command post
//
// swagger:model computer_command_post
type ComputerCommandPost struct {

	// computers
	Computers *ComputerCommandPostComputers `json:"computers,omitempty"`

	// general
	General *ComputerCommandPostGeneral `json:"general,omitempty"`
}

// Validate validates this computer command post
func (m *ComputerCommandPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneral(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerCommandPost) validateComputers(formats strfmt.Registry) error {
	if swag.IsZero(m.Computers) { // not required
		return nil
	}

	if m.Computers != nil {
		if err := m.Computers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computers")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerCommandPost) validateGeneral(formats strfmt.Registry) error {
	if swag.IsZero(m.General) { // not required
		return nil
	}

	if m.General != nil {
		if err := m.General.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computer command post based on the context it is used
func (m *ComputerCommandPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeneral(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerCommandPost) contextValidateComputers(ctx context.Context, formats strfmt.Registry) error {

	if m.Computers != nil {
		if err := m.Computers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computers")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerCommandPost) contextValidateGeneral(ctx context.Context, formats strfmt.Registry) error {

	if m.General != nil {
		if err := m.General.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerCommandPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerCommandPost) UnmarshalBinary(b []byte) error {
	var res ComputerCommandPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerCommandPostComputers computer command post computers
//
// swagger:model ComputerCommandPostComputers
type ComputerCommandPostComputers struct {

	// computer
	Computer *ComputerCommandPostComputersComputer `json:"computer,omitempty"`
}

// Validate validates this computer command post computers
func (m *ComputerCommandPostComputers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerCommandPostComputers) validateComputer(formats strfmt.Registry) error {
	if swag.IsZero(m.Computer) { // not required
		return nil
	}

	if m.Computer != nil {
		if err := m.Computer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computers" + "." + "computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computers" + "." + "computer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computer command post computers based on the context it is used
func (m *ComputerCommandPostComputers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerCommandPostComputers) contextValidateComputer(ctx context.Context, formats strfmt.Registry) error {

	if m.Computer != nil {
		if err := m.Computer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computers" + "." + "computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computers" + "." + "computer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerCommandPostComputers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerCommandPostComputers) UnmarshalBinary(b []byte) error {
	var res ComputerCommandPostComputers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerCommandPostComputersComputer computer command post computers computer
//
// swagger:model ComputerCommandPostComputersComputer
type ComputerCommandPostComputersComputer struct {

	// id
	// Example: 1
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this computer command post computers computer
func (m *ComputerCommandPostComputersComputer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerCommandPostComputersComputer) validateID(formats strfmt.Registry) error {

	if err := validate.Required("computers"+"."+"computer"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer command post computers computer based on context it is used
func (m *ComputerCommandPostComputersComputer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerCommandPostComputersComputer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerCommandPostComputersComputer) UnmarshalBinary(b []byte) error {
	var res ComputerCommandPostComputersComputer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerCommandPostGeneral computer command post general
//
// swagger:model ComputerCommandPostGeneral
type ComputerCommandPostGeneral struct {

	// Command type. (UnlockUserAccount and DeleteUser require a DEP enrolled device, SettingsEnableBluetooth and SettingsDisableBluetooth require macOS 10.13.4 or later, EnableRemoteDesktop and DisableRemoteDesktop require macOS 10.14.4 or later)
	// Required: true
	// Enum: [DeviceLock EraseDevice UnmanageDevice BlankPush UnlockUserAccount DeleteUser SettingsEnableBluetooth SettingsDisableBluetooth DisableRemoteDesktop EnableRemoteDesktop]
	Command *string `json:"command"`

	// Passcode to apply for DeviceLock and EraseDevice commands
	// Example: 123456
	Passcode string `json:"passcode,omitempty"`

	// Username for UnlockUserAccount and DeleteUser commands
	// Example: administrator
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this computer command post general
func (m *ComputerCommandPostGeneral) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerCommandPostGeneralTypeCommandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DeviceLock","EraseDevice","UnmanageDevice","BlankPush","UnlockUserAccount","DeleteUser","SettingsEnableBluetooth","SettingsDisableBluetooth","DisableRemoteDesktop","EnableRemoteDesktop"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerCommandPostGeneralTypeCommandPropEnum = append(computerCommandPostGeneralTypeCommandPropEnum, v)
	}
}

const (

	// ComputerCommandPostGeneralCommandDeviceLock captures enum value "DeviceLock"
	ComputerCommandPostGeneralCommandDeviceLock string = "DeviceLock"

	// ComputerCommandPostGeneralCommandEraseDevice captures enum value "EraseDevice"
	ComputerCommandPostGeneralCommandEraseDevice string = "EraseDevice"

	// ComputerCommandPostGeneralCommandUnmanageDevice captures enum value "UnmanageDevice"
	ComputerCommandPostGeneralCommandUnmanageDevice string = "UnmanageDevice"

	// ComputerCommandPostGeneralCommandBlankPush captures enum value "BlankPush"
	ComputerCommandPostGeneralCommandBlankPush string = "BlankPush"

	// ComputerCommandPostGeneralCommandUnlockUserAccount captures enum value "UnlockUserAccount"
	ComputerCommandPostGeneralCommandUnlockUserAccount string = "UnlockUserAccount"

	// ComputerCommandPostGeneralCommandDeleteUser captures enum value "DeleteUser"
	ComputerCommandPostGeneralCommandDeleteUser string = "DeleteUser"

	// ComputerCommandPostGeneralCommandSettingsEnableBluetooth captures enum value "SettingsEnableBluetooth"
	ComputerCommandPostGeneralCommandSettingsEnableBluetooth string = "SettingsEnableBluetooth"

	// ComputerCommandPostGeneralCommandSettingsDisableBluetooth captures enum value "SettingsDisableBluetooth"
	ComputerCommandPostGeneralCommandSettingsDisableBluetooth string = "SettingsDisableBluetooth"

	// ComputerCommandPostGeneralCommandDisableRemoteDesktop captures enum value "DisableRemoteDesktop"
	ComputerCommandPostGeneralCommandDisableRemoteDesktop string = "DisableRemoteDesktop"

	// ComputerCommandPostGeneralCommandEnableRemoteDesktop captures enum value "EnableRemoteDesktop"
	ComputerCommandPostGeneralCommandEnableRemoteDesktop string = "EnableRemoteDesktop"
)

// prop value enum
func (m *ComputerCommandPostGeneral) validateCommandEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerCommandPostGeneralTypeCommandPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerCommandPostGeneral) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("general"+"."+"command", "body", m.Command); err != nil {
		return err
	}

	// value enum
	if err := m.validateCommandEnum("general"+"."+"command", "body", *m.Command); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer command post general based on context it is used
func (m *ComputerCommandPostGeneral) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerCommandPostGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerCommandPostGeneral) UnmarshalBinary(b []byte) error {
	var res ComputerCommandPostGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
