package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Class resource defined in the Terraform configuration
func ClassSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apple_tvs": {
			Type: schema.TypeList, //GoType: []*ClassAppleTvsItems0
			Elem: &schema.Resource{
				Schema: ClassAppleTvsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"meeting_times": {
			Type: schema.TypeList, //GoType: ClassMeetingTimes
			Elem: &schema.Resource{
				Schema: ClassMeetingTimesSchema(),
			},
			Optional: true,
		},

		"mobile_device_group": {
			Type: schema.TypeList, //GoType: IDName
			Elem: &schema.Resource{
				Schema: IDNameSchema(),
			},
			Optional: true,
		},

		"mobile_device_group_id": {
			Type: schema.TypeList, //GoType: []*ClassMobileDeviceGroupIDItems0
			Elem: &schema.Resource{
				Schema: ClassMobileDeviceGroupIDItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"mobile_devices": {
			Type: schema.TypeList, //GoType: []*ClassMobileDevicesItems0
			Elem: &schema.Resource{
				Schema: ClassMobileDevicesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: SiteObject
			Elem: &schema.Resource{
				Schema: SiteObjectSchema(),
			},
			Optional: true,
		},

		"source": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"student_group_ids": {
			Type: schema.TypeList, //GoType: []*ClassStudentGroupIdsItems0
			Elem: &schema.Resource{
				Schema: ClassStudentGroupIdsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"students": {
			Type: schema.TypeList, //GoType: []*ClassStudentsItems0
			Elem: &schema.Resource{
				Schema: ClassStudentsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"teacher_group_ids": {
			Type: schema.TypeList, //GoType: []*ClassTeacherGroupIdsItems0
			Elem: &schema.Resource{
				Schema: ClassTeacherGroupIdsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"teacher_ids": {
			Type: schema.TypeList, //GoType: []*ClassTeacherIdsItems0
			Elem: &schema.Resource{
				Schema: ClassTeacherIdsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"teachers": {
			Type: schema.TypeList, //GoType: []*ClassTeachersItems0
			Elem: &schema.Resource{
				Schema: ClassTeachersItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying Class resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClassResourceData(d *schema.ResourceData, m *models.Class) {
	d.Set("apple_tvs", SetClassAppleTvsItems0SubResourceData(m.AppleTvs))
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("meeting_times", SetClassMeetingTimesSubResourceData([]*models.ClassMeetingTimes{m.MeetingTimes}))
	d.Set("mobile_device_group", SetIDNameSubResourceData([]*models.IDName{m.MobileDeviceGroup}))
	d.Set("mobile_device_group_id", SetClassMobileDeviceGroupIDItems0SubResourceData(m.MobileDeviceGroupID))
	d.Set("mobile_devices", SetClassMobileDevicesItems0SubResourceData(m.MobileDevices))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("source", m.Source)
	d.Set("student_group_ids", SetClassStudentGroupIdsItems0SubResourceData(m.StudentGroupIds))
	d.Set("students", SetClassStudentsItems0SubResourceData(m.Students))
	d.Set("teacher_group_ids", SetClassTeacherGroupIdsItems0SubResourceData(m.TeacherGroupIds))
	d.Set("teacher_ids", SetClassTeacherIdsItems0SubResourceData(m.TeacherIds))
	d.Set("teachers", SetClassTeachersItems0SubResourceData(m.Teachers))
}

// Iterate throught and update the Class resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClassSubResourceData(m []*models.Class) (d []*map[string]interface{}) {
	for _, class := range m {
		if class != nil {
			properties := make(map[string]interface{})
			properties["apple_tvs"] = SetClassAppleTvsItems0SubResourceData(class.AppleTvs)
			properties["description"] = class.Description
			properties["id"] = class.ID
			properties["meeting_times"] = SetClassMeetingTimesSubResourceData([]*models.ClassMeetingTimes{class.MeetingTimes})
			properties["mobile_device_group"] = SetIDNameSubResourceData([]*models.IDName{class.MobileDeviceGroup})
			properties["mobile_device_group_id"] = SetClassMobileDeviceGroupIDItems0SubResourceData(class.MobileDeviceGroupID)
			properties["mobile_devices"] = SetClassMobileDevicesItems0SubResourceData(class.MobileDevices)
			properties["name"] = class.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{class.Site})
			properties["source"] = class.Source
			properties["student_group_ids"] = SetClassStudentGroupIdsItems0SubResourceData(class.StudentGroupIds)
			properties["students"] = SetClassStudentsItems0SubResourceData(class.Students)
			properties["teacher_group_ids"] = SetClassTeacherGroupIdsItems0SubResourceData(class.TeacherGroupIds)
			properties["teacher_ids"] = SetClassTeacherIdsItems0SubResourceData(class.TeacherIds)
			properties["teachers"] = SetClassTeachersItems0SubResourceData(class.Teachers)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Class resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClassModel(d *schema.ResourceData) *models.Class {
	appleTvs := d.Get("apple_tvs").([]*models.ClassAppleTvsItems0)
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	var meetingTimes *models.ClassMeetingTimes = nil //hit complex
	meeting_timesInterface, meeting_timesIsSet := d.GetOk("meeting_times")
	if meeting_timesIsSet {
		meeting_timesMap := meeting_timesInterface.([]interface{})[0].(map[string]interface{})
		meetingTimes = ClassMeetingTimesModel(meeting_timesMap)
	}
	var mobileDeviceGroup *models.IDName = nil //hit complex
	mobile_device_groupInterface, mobile_device_groupIsSet := d.GetOk("mobile_device_group")
	if mobile_device_groupIsSet {
		mobile_device_groupMap := mobile_device_groupInterface.([]interface{})[0].(map[string]interface{})
		mobileDeviceGroup = IDNameModel(mobile_device_groupMap)
	}
	mobileDeviceGroupID := d.Get("mobile_device_group_id").([]*models.ClassMobileDeviceGroupIDItems0)
	mobileDevices := d.Get("mobile_devices").([]*models.ClassMobileDevicesItems0)
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	source := d.Get("source").(string)
	studentGroupIds := d.Get("student_group_ids").([]*models.ClassStudentGroupIdsItems0)
	students := d.Get("students").([]*models.ClassStudentsItems0)
	teacherGroupIds := d.Get("teacher_group_ids").([]*models.ClassTeacherGroupIdsItems0)
	teacherIds := d.Get("teacher_ids").([]*models.ClassTeacherIdsItems0)
	teachers := d.Get("teachers").([]*models.ClassTeachersItems0)

	return &models.Class{
		AppleTvs:            appleTvs,
		Description:         description,
		ID:                  int32(id),
		MeetingTimes:        meetingTimes,
		MobileDeviceGroup:   mobileDeviceGroup,
		MobileDeviceGroupID: mobileDeviceGroupID,
		MobileDevices:       mobileDevices,
		Name:                &name,
		Site:                site,
		Source:              source,
		StudentGroupIds:     studentGroupIds,
		Students:            students,
		TeacherGroupIds:     teacherGroupIds,
		TeacherIds:          teacherIds,
		Teachers:            teachers,
	}
}

// Retrieve property field names for updating the Class resource
func GetClassPropertyFields() (t []string) {
	return []string{
		"apple_tvs",
		"description",
		"id",
		"meeting_times",
		"mobile_device_group",
		"mobile_device_group_id",
		"mobile_devices",
		"name",
		"site",
		"source",
		"student_group_ids",
		"students",
		"teacher_group_ids",
		"teacher_ids",
		"teachers",
	}
}
