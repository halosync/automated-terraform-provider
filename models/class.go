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

// Class class
//
// swagger:model class
type Class struct {

	// apple tvs
	AppleTvs []*ClassAppleTvsItems0 `json:"apple_tvs"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// meeting times
	MeetingTimes *ClassMeetingTimes `json:"meeting_times,omitempty"`

	// mobile device group
	MobileDeviceGroup *IDName `json:"mobile_device_group,omitempty"`

	// mobile device group id
	MobileDeviceGroupID []*ClassMobileDeviceGroupIDItems0 `json:"mobile_device_group_id"`

	// mobile devices
	MobileDevices []*ClassMobileDevicesItems0 `json:"mobile_devices"`

	// Name of the class
	// Example: Math 101
	// Required: true
	Name *string `json:"name"`

	// site
	Site *SiteObject `json:"site,omitempty"`

	// source
	// Example: N/A
	Source string `json:"source,omitempty"`

	// student group ids
	StudentGroupIds []*ClassStudentGroupIdsItems0 `json:"student_group_ids"`

	// students
	Students []*ClassStudentsItems0 `json:"students"`

	// teacher group ids
	TeacherGroupIds []*ClassTeacherGroupIdsItems0 `json:"teacher_group_ids"`

	// teacher ids
	TeacherIds []*ClassTeacherIdsItems0 `json:"teacher_ids"`

	// teachers
	Teachers []*ClassTeachersItems0 `json:"teachers"`
}

// Validate validates this class
func (m *Class) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppleTvs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeetingTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileDeviceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileDeviceGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStudentGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStudents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeacherGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeacherIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeachers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Class) validateAppleTvs(formats strfmt.Registry) error {
	if swag.IsZero(m.AppleTvs) { // not required
		return nil
	}

	for i := 0; i < len(m.AppleTvs); i++ {
		if swag.IsZero(m.AppleTvs[i]) { // not required
			continue
		}

		if m.AppleTvs[i] != nil {
			if err := m.AppleTvs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apple_tvs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apple_tvs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateMeetingTimes(formats strfmt.Registry) error {
	if swag.IsZero(m.MeetingTimes) { // not required
		return nil
	}

	if m.MeetingTimes != nil {
		if err := m.MeetingTimes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meeting_times")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meeting_times")
			}
			return err
		}
	}

	return nil
}

func (m *Class) validateMobileDeviceGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDeviceGroup) { // not required
		return nil
	}

	if m.MobileDeviceGroup != nil {
		if err := m.MobileDeviceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device_group")
			}
			return err
		}
	}

	return nil
}

func (m *Class) validateMobileDeviceGroupID(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDeviceGroupID) { // not required
		return nil
	}

	for i := 0; i < len(m.MobileDeviceGroupID); i++ {
		if swag.IsZero(m.MobileDeviceGroupID[i]) { // not required
			continue
		}

		if m.MobileDeviceGroupID[i] != nil {
			if err := m.MobileDeviceGroupID[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_device_group_id" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_device_group_id" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateMobileDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.MobileDevices); i++ {
		if swag.IsZero(m.MobileDevices[i]) { // not required
			continue
		}

		if m.MobileDevices[i] != nil {
			if err := m.MobileDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Class) validateSite(formats strfmt.Registry) error {
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

func (m *Class) validateStudentGroupIds(formats strfmt.Registry) error {
	if swag.IsZero(m.StudentGroupIds) { // not required
		return nil
	}

	for i := 0; i < len(m.StudentGroupIds); i++ {
		if swag.IsZero(m.StudentGroupIds[i]) { // not required
			continue
		}

		if m.StudentGroupIds[i] != nil {
			if err := m.StudentGroupIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("student_group_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("student_group_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateStudents(formats strfmt.Registry) error {
	if swag.IsZero(m.Students) { // not required
		return nil
	}

	for i := 0; i < len(m.Students); i++ {
		if swag.IsZero(m.Students[i]) { // not required
			continue
		}

		if m.Students[i] != nil {
			if err := m.Students[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("students" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("students" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateTeacherGroupIds(formats strfmt.Registry) error {
	if swag.IsZero(m.TeacherGroupIds) { // not required
		return nil
	}

	for i := 0; i < len(m.TeacherGroupIds); i++ {
		if swag.IsZero(m.TeacherGroupIds[i]) { // not required
			continue
		}

		if m.TeacherGroupIds[i] != nil {
			if err := m.TeacherGroupIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teacher_group_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teacher_group_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateTeacherIds(formats strfmt.Registry) error {
	if swag.IsZero(m.TeacherIds) { // not required
		return nil
	}

	for i := 0; i < len(m.TeacherIds); i++ {
		if swag.IsZero(m.TeacherIds[i]) { // not required
			continue
		}

		if m.TeacherIds[i] != nil {
			if err := m.TeacherIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teacher_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teacher_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) validateTeachers(formats strfmt.Registry) error {
	if swag.IsZero(m.Teachers) { // not required
		return nil
	}

	for i := 0; i < len(m.Teachers); i++ {
		if swag.IsZero(m.Teachers[i]) { // not required
			continue
		}

		if m.Teachers[i] != nil {
			if err := m.Teachers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teachers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teachers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this class based on the context it is used
func (m *Class) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppleTvs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeetingTimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMobileDeviceGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMobileDeviceGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMobileDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStudentGroupIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStudents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeacherGroupIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeacherIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeachers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Class) contextValidateAppleTvs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppleTvs); i++ {

		if m.AppleTvs[i] != nil {
			if err := m.AppleTvs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apple_tvs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apple_tvs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateMeetingTimes(ctx context.Context, formats strfmt.Registry) error {

	if m.MeetingTimes != nil {
		if err := m.MeetingTimes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meeting_times")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meeting_times")
			}
			return err
		}
	}

	return nil
}

func (m *Class) contextValidateMobileDeviceGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.MobileDeviceGroup != nil {
		if err := m.MobileDeviceGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device_group")
			}
			return err
		}
	}

	return nil
}

func (m *Class) contextValidateMobileDeviceGroupID(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MobileDeviceGroupID); i++ {

		if m.MobileDeviceGroupID[i] != nil {
			if err := m.MobileDeviceGroupID[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_device_group_id" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_device_group_id" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateMobileDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MobileDevices); i++ {

		if m.MobileDevices[i] != nil {
			if err := m.MobileDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Class) contextValidateStudentGroupIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StudentGroupIds); i++ {

		if m.StudentGroupIds[i] != nil {
			if err := m.StudentGroupIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("student_group_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("student_group_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateStudents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Students); i++ {

		if m.Students[i] != nil {
			if err := m.Students[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("students" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("students" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateTeacherGroupIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TeacherGroupIds); i++ {

		if m.TeacherGroupIds[i] != nil {
			if err := m.TeacherGroupIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teacher_group_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teacher_group_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateTeacherIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TeacherIds); i++ {

		if m.TeacherIds[i] != nil {
			if err := m.TeacherIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teacher_ids" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teacher_ids" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Class) contextValidateTeachers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teachers); i++ {

		if m.Teachers[i] != nil {
			if err := m.Teachers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teachers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teachers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Class) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Class) UnmarshalBinary(b []byte) error {
	var res Class
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassAppleTvsItems0 class apple tvs items0
//
// swagger:model ClassAppleTvsItems0
type ClassAppleTvsItems0 struct {

	// apple tv
	AppleTv *ClassAppleTvsItems0AppleTv `json:"apple_tv,omitempty"`
}

// Validate validates this class apple tvs items0
func (m *ClassAppleTvsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppleTv(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassAppleTvsItems0) validateAppleTv(formats strfmt.Registry) error {
	if swag.IsZero(m.AppleTv) { // not required
		return nil
	}

	if m.AppleTv != nil {
		if err := m.AppleTv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apple_tv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apple_tv")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this class apple tvs items0 based on the context it is used
func (m *ClassAppleTvsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppleTv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassAppleTvsItems0) contextValidateAppleTv(ctx context.Context, formats strfmt.Registry) error {

	if m.AppleTv != nil {
		if err := m.AppleTv.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apple_tv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apple_tv")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClassAppleTvsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassAppleTvsItems0) UnmarshalBinary(b []byte) error {
	var res ClassAppleTvsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassAppleTvsItems0AppleTv class apple tvs items0 apple tv
//
// swagger:model ClassAppleTvsItems0AppleTv
type ClassAppleTvsItems0AppleTv struct {

	// airplay password
	// Example: password
	AirplayPassword string `json:"airplay_password,omitempty"`

	// device id
	// Example: E0:AC:CB:97:36:g5
	DeviceID string `json:"device_id,omitempty"`

	// name
	// Example: Apple TV
	Name string `json:"name,omitempty"`

	// udid
	// Example: 3e8c9775cb3302ed9e645adf632cfa533adc3aa8
	Udid string `json:"udid,omitempty"`

	// wifi mac address
	// Example: E0:AC:CB:97:36:G4
	WifiMacAddress string `json:"wifi_mac_address,omitempty"`
}

// Validate validates this class apple tvs items0 apple tv
func (m *ClassAppleTvsItems0AppleTv) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class apple tvs items0 apple tv based on context it is used
func (m *ClassAppleTvsItems0AppleTv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassAppleTvsItems0AppleTv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassAppleTvsItems0AppleTv) UnmarshalBinary(b []byte) error {
	var res ClassAppleTvsItems0AppleTv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassMeetingTimes class meeting times
//
// swagger:model ClassMeetingTimes
type ClassMeetingTimes struct {

	// meeting time
	MeetingTime *ClassMeetingTimesMeetingTime `json:"meeting_time,omitempty"`
}

// Validate validates this class meeting times
func (m *ClassMeetingTimes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeetingTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassMeetingTimes) validateMeetingTime(formats strfmt.Registry) error {
	if swag.IsZero(m.MeetingTime) { // not required
		return nil
	}

	if m.MeetingTime != nil {
		if err := m.MeetingTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meeting_times" + "." + "meeting_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meeting_times" + "." + "meeting_time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this class meeting times based on the context it is used
func (m *ClassMeetingTimes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeetingTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassMeetingTimes) contextValidateMeetingTime(ctx context.Context, formats strfmt.Registry) error {

	if m.MeetingTime != nil {
		if err := m.MeetingTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meeting_times" + "." + "meeting_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meeting_times" + "." + "meeting_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClassMeetingTimes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassMeetingTimes) UnmarshalBinary(b []byte) error {
	var res ClassMeetingTimes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassMeetingTimesMeetingTime class meeting times meeting time
//
// swagger:model ClassMeetingTimesMeetingTime
type ClassMeetingTimesMeetingTime struct {

	// days
	// Example: M W F
	// Enum: [M T W Th F Sa Su]
	Days string `json:"days,omitempty"`

	// end time
	// Example: 1345
	EndTime int64 `json:"end_time,omitempty"`

	// start time
	// Example: 1300
	StartTime int64 `json:"start_time,omitempty"`
}

// Validate validates this class meeting times meeting time
func (m *ClassMeetingTimesMeetingTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var classMeetingTimesMeetingTimeTypeDaysPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["M","T","W","Th","F","Sa","Su"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		classMeetingTimesMeetingTimeTypeDaysPropEnum = append(classMeetingTimesMeetingTimeTypeDaysPropEnum, v)
	}
}

const (

	// ClassMeetingTimesMeetingTimeDaysM captures enum value "M"
	ClassMeetingTimesMeetingTimeDaysM string = "M"

	// ClassMeetingTimesMeetingTimeDaysT captures enum value "T"
	ClassMeetingTimesMeetingTimeDaysT string = "T"

	// ClassMeetingTimesMeetingTimeDaysW captures enum value "W"
	ClassMeetingTimesMeetingTimeDaysW string = "W"

	// ClassMeetingTimesMeetingTimeDaysTh captures enum value "Th"
	ClassMeetingTimesMeetingTimeDaysTh string = "Th"

	// ClassMeetingTimesMeetingTimeDaysF captures enum value "F"
	ClassMeetingTimesMeetingTimeDaysF string = "F"

	// ClassMeetingTimesMeetingTimeDaysSa captures enum value "Sa"
	ClassMeetingTimesMeetingTimeDaysSa string = "Sa"

	// ClassMeetingTimesMeetingTimeDaysSu captures enum value "Su"
	ClassMeetingTimesMeetingTimeDaysSu string = "Su"
)

// prop value enum
func (m *ClassMeetingTimesMeetingTime) validateDaysEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, classMeetingTimesMeetingTimeTypeDaysPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClassMeetingTimesMeetingTime) validateDays(formats strfmt.Registry) error {
	if swag.IsZero(m.Days) { // not required
		return nil
	}

	// value enum
	if err := m.validateDaysEnum("meeting_times"+"."+"meeting_time"+"."+"days", "body", m.Days); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this class meeting times meeting time based on context it is used
func (m *ClassMeetingTimesMeetingTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassMeetingTimesMeetingTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassMeetingTimesMeetingTime) UnmarshalBinary(b []byte) error {
	var res ClassMeetingTimesMeetingTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassMobileDeviceGroupIDItems0 class mobile device group ID items0
//
// swagger:model ClassMobileDeviceGroupIDItems0
type ClassMobileDeviceGroupIDItems0 struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`
}

// Validate validates this class mobile device group ID items0
func (m *ClassMobileDeviceGroupIDItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class mobile device group ID items0 based on context it is used
func (m *ClassMobileDeviceGroupIDItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassMobileDeviceGroupIDItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassMobileDeviceGroupIDItems0) UnmarshalBinary(b []byte) error {
	var res ClassMobileDeviceGroupIDItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassMobileDevicesItems0 class mobile devices items0
//
// swagger:model ClassMobileDevicesItems0
type ClassMobileDevicesItems0 struct {

	// mobile device
	MobileDevice *ClassMobileDevicesItems0MobileDevice `json:"mobile_device,omitempty"`
}

// Validate validates this class mobile devices items0
func (m *ClassMobileDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMobileDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassMobileDevicesItems0) validateMobileDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDevice) { // not required
		return nil
	}

	if m.MobileDevice != nil {
		if err := m.MobileDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this class mobile devices items0 based on the context it is used
func (m *ClassMobileDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMobileDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClassMobileDevicesItems0) contextValidateMobileDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.MobileDevice != nil {
		if err := m.MobileDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClassMobileDevicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassMobileDevicesItems0) UnmarshalBinary(b []byte) error {
	var res ClassMobileDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassMobileDevicesItems0MobileDevice class mobile devices items0 mobile device
//
// swagger:model ClassMobileDevicesItems0MobileDevice
type ClassMobileDevicesItems0MobileDevice struct {

	// Name of the device
	// Example: Tinas iPad
	Name string `json:"name,omitempty"`

	// udid
	// Example: 270aae10800b6e61a2ee2bbc285eb967050b5984
	Udid string `json:"udid,omitempty"`

	// wifi mac address
	// Example: E0:AC:CB:97:36:G4
	WifiMacAddress string `json:"wifi_mac_address,omitempty"`
}

// Validate validates this class mobile devices items0 mobile device
func (m *ClassMobileDevicesItems0MobileDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class mobile devices items0 mobile device based on context it is used
func (m *ClassMobileDevicesItems0MobileDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassMobileDevicesItems0MobileDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassMobileDevicesItems0MobileDevice) UnmarshalBinary(b []byte) error {
	var res ClassMobileDevicesItems0MobileDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassStudentGroupIdsItems0 class student group ids items0
//
// swagger:model ClassStudentGroupIdsItems0
type ClassStudentGroupIdsItems0 struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this class student group ids items0
func (m *ClassStudentGroupIdsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class student group ids items0 based on context it is used
func (m *ClassStudentGroupIdsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassStudentGroupIdsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassStudentGroupIdsItems0) UnmarshalBinary(b []byte) error {
	var res ClassStudentGroupIdsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassStudentsItems0 class students items0
//
// swagger:model ClassStudentsItems0
type ClassStudentsItems0 struct {

	// Name of the student
	// Example: Joey
	Student string `json:"student,omitempty"`
}

// Validate validates this class students items0
func (m *ClassStudentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class students items0 based on context it is used
func (m *ClassStudentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassStudentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassStudentsItems0) UnmarshalBinary(b []byte) error {
	var res ClassStudentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassTeacherGroupIdsItems0 class teacher group ids items0
//
// swagger:model ClassTeacherGroupIdsItems0
type ClassTeacherGroupIdsItems0 struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this class teacher group ids items0
func (m *ClassTeacherGroupIdsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class teacher group ids items0 based on context it is used
func (m *ClassTeacherGroupIdsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassTeacherGroupIdsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassTeacherGroupIdsItems0) UnmarshalBinary(b []byte) error {
	var res ClassTeacherGroupIdsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassTeacherIdsItems0 class teacher ids items0
//
// swagger:model ClassTeacherIdsItems0
type ClassTeacherIdsItems0 struct {

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this class teacher ids items0
func (m *ClassTeacherIdsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class teacher ids items0 based on context it is used
func (m *ClassTeacherIdsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassTeacherIdsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassTeacherIdsItems0) UnmarshalBinary(b []byte) error {
	var res ClassTeacherIdsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClassTeachersItems0 class teachers items0
//
// swagger:model ClassTeachersItems0
type ClassTeachersItems0 struct {

	// Name of the teacher
	// Example: Mr. Smith
	Teacher string `json:"teacher,omitempty"`
}

// Validate validates this class teachers items0
func (m *ClassTeachersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this class teachers items0 based on context it is used
func (m *ClassTeachersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClassTeachersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClassTeachersItems0) UnmarshalBinary(b []byte) error {
	var res ClassTeachersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
