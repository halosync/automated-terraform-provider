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

// ComputerExtensionAttribute computer extension attribute
//
// swagger:model computer_extension_attribute
type ComputerExtensionAttribute struct {

	// data type
	// Enum: [String Integer Date]
	DataType string `json:"data_type,omitempty"`

	// Description of the extension attribute
	// Example: Number of charge cycles logged on the current battery
	Description string `json:"description,omitempty"`

	// Only applicable to script input type
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// input type
	InputType *ComputerExtensionAttributeInputType `json:"input_type,omitempty"`

	// Category in which to display the extension attribute in Jamf Pro
	// Enum: [General Hardware Operating System User and Location Purchasing Extension Attributes]
	InventoryDisplay string `json:"inventory_display,omitempty"`

	// Extension attribute name
	// Example: Battery Cycle Count
	// Required: true
	Name *string `json:"name"`

	// Pane on which to display the extension attribute in Recon
	// Enum: [Computer User and Location Purchasing Extension Attributes]
	ReconDisplay *string `json:"recon_display,omitempty"`
}

// Validate validates this computer extension attribute
func (m *ComputerExtensionAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInventoryDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconDisplay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerExtensionAttributeTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["String","Integer","Date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerExtensionAttributeTypeDataTypePropEnum = append(computerExtensionAttributeTypeDataTypePropEnum, v)
	}
}

const (

	// ComputerExtensionAttributeDataTypeString captures enum value "String"
	ComputerExtensionAttributeDataTypeString string = "String"

	// ComputerExtensionAttributeDataTypeInteger captures enum value "Integer"
	ComputerExtensionAttributeDataTypeInteger string = "Integer"

	// ComputerExtensionAttributeDataTypeDate captures enum value "Date"
	ComputerExtensionAttributeDataTypeDate string = "Date"
)

// prop value enum
func (m *ComputerExtensionAttribute) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerExtensionAttributeTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerExtensionAttribute) validateDataType(formats strfmt.Registry) error {
	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("data_type", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *ComputerExtensionAttribute) validateInputType(formats strfmt.Registry) error {
	if swag.IsZero(m.InputType) { // not required
		return nil
	}

	if m.InputType != nil {
		if err := m.InputType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input_type")
			}
			return err
		}
	}

	return nil
}

var computerExtensionAttributeTypeInventoryDisplayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["General","Hardware","Operating System","User and Location","Purchasing","Extension Attributes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerExtensionAttributeTypeInventoryDisplayPropEnum = append(computerExtensionAttributeTypeInventoryDisplayPropEnum, v)
	}
}

const (

	// ComputerExtensionAttributeInventoryDisplayGeneral captures enum value "General"
	ComputerExtensionAttributeInventoryDisplayGeneral string = "General"

	// ComputerExtensionAttributeInventoryDisplayHardware captures enum value "Hardware"
	ComputerExtensionAttributeInventoryDisplayHardware string = "Hardware"

	// ComputerExtensionAttributeInventoryDisplayOperatingSystem captures enum value "Operating System"
	ComputerExtensionAttributeInventoryDisplayOperatingSystem string = "Operating System"

	// ComputerExtensionAttributeInventoryDisplayUserAndLocation captures enum value "User and Location"
	ComputerExtensionAttributeInventoryDisplayUserAndLocation string = "User and Location"

	// ComputerExtensionAttributeInventoryDisplayPurchasing captures enum value "Purchasing"
	ComputerExtensionAttributeInventoryDisplayPurchasing string = "Purchasing"

	// ComputerExtensionAttributeInventoryDisplayExtensionAttributes captures enum value "Extension Attributes"
	ComputerExtensionAttributeInventoryDisplayExtensionAttributes string = "Extension Attributes"
)

// prop value enum
func (m *ComputerExtensionAttribute) validateInventoryDisplayEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerExtensionAttributeTypeInventoryDisplayPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerExtensionAttribute) validateInventoryDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.InventoryDisplay) { // not required
		return nil
	}

	// value enum
	if err := m.validateInventoryDisplayEnum("inventory_display", "body", m.InventoryDisplay); err != nil {
		return err
	}

	return nil
}

func (m *ComputerExtensionAttribute) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var computerExtensionAttributeTypeReconDisplayPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Computer","User and Location","Purchasing","Extension Attributes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerExtensionAttributeTypeReconDisplayPropEnum = append(computerExtensionAttributeTypeReconDisplayPropEnum, v)
	}
}

const (

	// ComputerExtensionAttributeReconDisplayComputer captures enum value "Computer"
	ComputerExtensionAttributeReconDisplayComputer string = "Computer"

	// ComputerExtensionAttributeReconDisplayUserAndLocation captures enum value "User and Location"
	ComputerExtensionAttributeReconDisplayUserAndLocation string = "User and Location"

	// ComputerExtensionAttributeReconDisplayPurchasing captures enum value "Purchasing"
	ComputerExtensionAttributeReconDisplayPurchasing string = "Purchasing"

	// ComputerExtensionAttributeReconDisplayExtensionAttributes captures enum value "Extension Attributes"
	ComputerExtensionAttributeReconDisplayExtensionAttributes string = "Extension Attributes"
)

// prop value enum
func (m *ComputerExtensionAttribute) validateReconDisplayEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerExtensionAttributeTypeReconDisplayPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerExtensionAttribute) validateReconDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.ReconDisplay) { // not required
		return nil
	}

	// value enum
	if err := m.validateReconDisplayEnum("recon_display", "body", *m.ReconDisplay); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this computer extension attribute based on the context it is used
func (m *ComputerExtensionAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerExtensionAttribute) contextValidateInputType(ctx context.Context, formats strfmt.Registry) error {

	if m.InputType != nil {
		if err := m.InputType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerExtensionAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerExtensionAttribute) UnmarshalBinary(b []byte) error {
	var res ComputerExtensionAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerExtensionAttributeInputType computer extension attribute input type
//
// swagger:model ComputerExtensionAttributeInputType
type ComputerExtensionAttributeInputType struct {

	// type
	// Enum: [script Text Field LDAP Mapping Pop-up Menu]
	Type *string `json:"type,omitempty"`
}

// Validate validates this computer extension attribute input type
func (m *ComputerExtensionAttributeInputType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerExtensionAttributeInputTypeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["script","Text Field","LDAP Mapping","Pop-up Menu"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerExtensionAttributeInputTypeTypeTypePropEnum = append(computerExtensionAttributeInputTypeTypeTypePropEnum, v)
	}
}

const (

	// ComputerExtensionAttributeInputTypeTypeScript captures enum value "script"
	ComputerExtensionAttributeInputTypeTypeScript string = "script"

	// ComputerExtensionAttributeInputTypeTypeTextField captures enum value "Text Field"
	ComputerExtensionAttributeInputTypeTypeTextField string = "Text Field"

	// ComputerExtensionAttributeInputTypeTypeLDAPMapping captures enum value "LDAP Mapping"
	ComputerExtensionAttributeInputTypeTypeLDAPMapping string = "LDAP Mapping"

	// ComputerExtensionAttributeInputTypeTypePopDashUpMenu captures enum value "Pop-up Menu"
	ComputerExtensionAttributeInputTypeTypePopDashUpMenu string = "Pop-up Menu"
)

// prop value enum
func (m *ComputerExtensionAttributeInputType) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerExtensionAttributeInputTypeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerExtensionAttributeInputType) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("input_type"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer extension attribute input type based on context it is used
func (m *ComputerExtensionAttributeInputType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerExtensionAttributeInputType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerExtensionAttributeInputType) UnmarshalBinary(b []byte) error {
	var res ComputerExtensionAttributeInputType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
