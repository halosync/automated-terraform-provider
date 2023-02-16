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

// ComputerHardwareSoftwareReports computer hardware software reports
//
// swagger:model computer_hardware_software_reports
type ComputerHardwareSoftwareReports struct {

	// font report
	FontReport *ComputerHardwareSoftwareReportsFontReport `json:"font_report,omitempty"`

	// hardware report
	HardwareReport *ComputerHardwareSoftwareReportsHardwareReport `json:"hardware_report,omitempty"`

	// plugin report
	PluginReport *ComputerHardwareSoftwareReportsPluginReport `json:"plugin_report,omitempty"`

	// software report
	SoftwareReport *ComputerHardwareSoftwareReportsSoftwareReport `json:"software_report,omitempty"`
}

// Validate validates this computer hardware software reports
func (m *ComputerHardwareSoftwareReports) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFontReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardwareReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginReport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareReport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerHardwareSoftwareReports) validateFontReport(formats strfmt.Registry) error {
	if swag.IsZero(m.FontReport) { // not required
		return nil
	}

	if m.FontReport != nil {
		if err := m.FontReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("font_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("font_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) validateHardwareReport(formats strfmt.Registry) error {
	if swag.IsZero(m.HardwareReport) { // not required
		return nil
	}

	if m.HardwareReport != nil {
		if err := m.HardwareReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) validatePluginReport(formats strfmt.Registry) error {
	if swag.IsZero(m.PluginReport) { // not required
		return nil
	}

	if m.PluginReport != nil {
		if err := m.PluginReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugin_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plugin_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) validateSoftwareReport(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareReport) { // not required
		return nil
	}

	if m.SoftwareReport != nil {
		if err := m.SoftwareReport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_report")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computer hardware software reports based on the context it is used
func (m *ComputerHardwareSoftwareReports) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFontReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHardwareReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePluginReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareReport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerHardwareSoftwareReports) contextValidateFontReport(ctx context.Context, formats strfmt.Registry) error {

	if m.FontReport != nil {
		if err := m.FontReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("font_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("font_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) contextValidateHardwareReport(ctx context.Context, formats strfmt.Registry) error {

	if m.HardwareReport != nil {
		if err := m.HardwareReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) contextValidatePluginReport(ctx context.Context, formats strfmt.Registry) error {

	if m.PluginReport != nil {
		if err := m.PluginReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugin_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plugin_report")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerHardwareSoftwareReports) contextValidateSoftwareReport(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareReport != nil {
		if err := m.SoftwareReport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_report")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_report")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReports) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReports) UnmarshalBinary(b []byte) error {
	var res ComputerHardwareSoftwareReports
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerHardwareSoftwareReportsFontReport computer hardware software reports font report
//
// swagger:model ComputerHardwareSoftwareReportsFontReport
type ComputerHardwareSoftwareReportsFontReport struct {

	// date time
	// Example: 2017-07-07 18:37:04
	DateTime string `json:"date_time,omitempty"`

	// date time epoch
	// Example: 1499470624555
	DateTimeEpoch string `json:"date_time_epoch,omitempty"`

	// date time utc
	// Example: 2017-07-07T18:37:04.555-0500
	DateTimeUtc string `json:"date_time_utc,omitempty"`

	// name
	// Example: Al Nile.ttc
	Name string `json:"name,omitempty"`

	// path
	// Example: /Library/Fonts/Al Nile.ttc
	Path string `json:"path,omitempty"`

	// type
	// Enum: [Added Deleted]
	Type string `json:"type,omitempty"`

	// version
	// Example: n/a
	Version string `json:"version,omitempty"`
}

// Validate validates this computer hardware software reports font report
func (m *ComputerHardwareSoftwareReportsFontReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerHardwareSoftwareReportsFontReportTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Added","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerHardwareSoftwareReportsFontReportTypeTypePropEnum = append(computerHardwareSoftwareReportsFontReportTypeTypePropEnum, v)
	}
}

const (

	// ComputerHardwareSoftwareReportsFontReportTypeAdded captures enum value "Added"
	ComputerHardwareSoftwareReportsFontReportTypeAdded string = "Added"

	// ComputerHardwareSoftwareReportsFontReportTypeDeleted captures enum value "Deleted"
	ComputerHardwareSoftwareReportsFontReportTypeDeleted string = "Deleted"
)

// prop value enum
func (m *ComputerHardwareSoftwareReportsFontReport) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerHardwareSoftwareReportsFontReportTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerHardwareSoftwareReportsFontReport) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("font_report"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer hardware software reports font report based on context it is used
func (m *ComputerHardwareSoftwareReportsFontReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsFontReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsFontReport) UnmarshalBinary(b []byte) error {
	var res ComputerHardwareSoftwareReportsFontReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerHardwareSoftwareReportsHardwareReport computer hardware software reports hardware report
//
// swagger:model ComputerHardwareSoftwareReportsHardwareReport
type ComputerHardwareSoftwareReportsHardwareReport struct {

	// n i c speed
	// Example: n/a
	NICSpeed string `json:"NIC_speed,omitempty"`

	// boot partition used percent
	// Example: 19
	BootPartitionUsedPercent int64 `json:"boot_partition_used_percent,omitempty"`

	// core count
	// Example: 2
	CoreCount int64 `json:"core_count,omitempty"`

	// date time
	// Example: 2017-07-07 18:37:04
	DateTime string `json:"date_time,omitempty"`

	// date time epoch
	// Example: 1499470624555
	DateTimeEpoch string `json:"date_time_epoch,omitempty"`

	// date time utc
	// Example: 2017-07-07T18:37:04.555-0500
	DateTimeUtc string `json:"date_time_utc,omitempty"`

	// make
	// Example: Apple
	Make string `json:"make,omitempty"`

	// model identifier
	// Example: MacBookPro11,1
	ModelIdentifier string `json:"model_identifier,omitempty"`

	// open ram slots
	// Example: 0
	OpenRAMSlots int64 `json:"open_ram_slots,omitempty"`

	// operating system
	// Example: Mac OS X 10.12.4
	OperatingSystem string `json:"operating_system,omitempty"`

	// optical drive
	OpticalDrive string `json:"optical_drive,omitempty"`

	// processor count
	// Example: 1
	ProcessorCount int64 `json:"processor_count,omitempty"`

	// processor speed mhz
	// Example: 2600
	ProcessorSpeedMhz int64 `json:"processor_speed_mhz,omitempty"`

	// serial number
	// Example: C02Q7KHTGFWF
	SerialNumber string `json:"serial_number,omitempty"`

	// service pack
	ServicePack string `json:"service_pack,omitempty"`

	// total harddrive size
	// Example: 500.28 GB
	TotalHarddriveSize string `json:"total_harddrive_size,omitempty"`

	// total ram mb
	// Example: 16384
	TotalRAMMb int64 `json:"total_ram_mb,omitempty"`
}

// Validate validates this computer hardware software reports hardware report
func (m *ComputerHardwareSoftwareReportsHardwareReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer hardware software reports hardware report based on context it is used
func (m *ComputerHardwareSoftwareReportsHardwareReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsHardwareReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsHardwareReport) UnmarshalBinary(b []byte) error {
	var res ComputerHardwareSoftwareReportsHardwareReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerHardwareSoftwareReportsPluginReport computer hardware software reports plugin report
//
// swagger:model ComputerHardwareSoftwareReportsPluginReport
type ComputerHardwareSoftwareReportsPluginReport struct {

	// date time
	// Example: 2017-07-07 18:37:04
	DateTime string `json:"date_time,omitempty"`

	// date time epoch
	// Example: 1499470624555
	DateTimeEpoch string `json:"date_time_epoch,omitempty"`

	// date time utc
	// Example: 2017-07-07T18:37:04.555-0500
	DateTimeUtc string `json:"date_time_utc,omitempty"`

	// name
	// Example: Quartz Composer.webplugin
	Name string `json:"name,omitempty"`

	// path
	// Example: /Library/Internet Plug-Ins/Quartz Composer.webplugin
	Path string `json:"path,omitempty"`

	// type
	// Enum: [Added Removed]
	Type string `json:"type,omitempty"`

	// version
	// Example: 1.4
	Version string `json:"version,omitempty"`
}

// Validate validates this computer hardware software reports plugin report
func (m *ComputerHardwareSoftwareReportsPluginReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerHardwareSoftwareReportsPluginReportTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Added","Removed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerHardwareSoftwareReportsPluginReportTypeTypePropEnum = append(computerHardwareSoftwareReportsPluginReportTypeTypePropEnum, v)
	}
}

const (

	// ComputerHardwareSoftwareReportsPluginReportTypeAdded captures enum value "Added"
	ComputerHardwareSoftwareReportsPluginReportTypeAdded string = "Added"

	// ComputerHardwareSoftwareReportsPluginReportTypeRemoved captures enum value "Removed"
	ComputerHardwareSoftwareReportsPluginReportTypeRemoved string = "Removed"
)

// prop value enum
func (m *ComputerHardwareSoftwareReportsPluginReport) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerHardwareSoftwareReportsPluginReportTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerHardwareSoftwareReportsPluginReport) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("plugin_report"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer hardware software reports plugin report based on context it is used
func (m *ComputerHardwareSoftwareReportsPluginReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsPluginReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsPluginReport) UnmarshalBinary(b []byte) error {
	var res ComputerHardwareSoftwareReportsPluginReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerHardwareSoftwareReportsSoftwareReport computer hardware software reports software report
//
// swagger:model ComputerHardwareSoftwareReportsSoftwareReport
type ComputerHardwareSoftwareReportsSoftwareReport struct {

	// date time
	// Example: 2017-07-07 18:37:04
	DateTime string `json:"date_time,omitempty"`

	// date time epoch
	// Example: 1499470624555
	DateTimeEpoch string `json:"date_time_epoch,omitempty"`

	// date time utc
	// Example: 2017-07-07T18:37:04.555-0500
	DateTimeUtc string `json:"date_time_utc,omitempty"`

	// name
	// Example: Parallels Desktop.app
	Name string `json:"name,omitempty"`

	// path
	// Example: /Applications/Parallels Desktop.app
	Path string `json:"path,omitempty"`

	// type
	// Enum: [Added Deleted]
	Type string `json:"type,omitempty"`

	// version
	// Example: 9
	Version string `json:"version,omitempty"`
}

// Validate validates this computer hardware software reports software report
func (m *ComputerHardwareSoftwareReportsSoftwareReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var computerHardwareSoftwareReportsSoftwareReportTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Added","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		computerHardwareSoftwareReportsSoftwareReportTypeTypePropEnum = append(computerHardwareSoftwareReportsSoftwareReportTypeTypePropEnum, v)
	}
}

const (

	// ComputerHardwareSoftwareReportsSoftwareReportTypeAdded captures enum value "Added"
	ComputerHardwareSoftwareReportsSoftwareReportTypeAdded string = "Added"

	// ComputerHardwareSoftwareReportsSoftwareReportTypeDeleted captures enum value "Deleted"
	ComputerHardwareSoftwareReportsSoftwareReportTypeDeleted string = "Deleted"
)

// prop value enum
func (m *ComputerHardwareSoftwareReportsSoftwareReport) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, computerHardwareSoftwareReportsSoftwareReportTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComputerHardwareSoftwareReportsSoftwareReport) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("software_report"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this computer hardware software reports software report based on context it is used
func (m *ComputerHardwareSoftwareReportsSoftwareReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsSoftwareReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerHardwareSoftwareReportsSoftwareReport) UnmarshalBinary(b []byte) error {
	var res ComputerHardwareSoftwareReportsSoftwareReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
