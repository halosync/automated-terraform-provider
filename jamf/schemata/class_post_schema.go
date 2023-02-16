package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ClassPost resource defined in the Terraform configuration
func ClassPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"meeting_times": {
			Type: schema.TypeList, //GoType: ClassPostMeetingTimes
			Elem: &schema.Resource{
				Schema: ClassPostMeetingTimesSchema(),
			},
			Optional: true,
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
	}
}

// Update the underlying ClassPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClassPostResourceData(d *schema.ResourceData, m *models.ClassPost) {
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("meeting_times", SetClassPostMeetingTimesSubResourceData([]*models.ClassPostMeetingTimes{m.MeetingTimes}))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("source", m.Source)
}

// Iterate throught and update the ClassPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClassPostSubResourceData(m []*models.ClassPost) (d []*map[string]interface{}) {
	for _, classPost := range m {
		if classPost != nil {
			properties := make(map[string]interface{})
			properties["description"] = classPost.Description
			properties["id"] = classPost.ID
			properties["meeting_times"] = SetClassPostMeetingTimesSubResourceData([]*models.ClassPostMeetingTimes{classPost.MeetingTimes})
			properties["name"] = classPost.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{classPost.Site})
			properties["source"] = classPost.Source
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ClassPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClassPostModel(d *schema.ResourceData) *models.ClassPost {
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	var meetingTimes *models.ClassPostMeetingTimes = nil //hit complex
	meeting_timesInterface, meeting_timesIsSet := d.GetOk("meeting_times")
	if meeting_timesIsSet {
		meeting_timesMap := meeting_timesInterface.([]interface{})[0].(map[string]interface{})
		meetingTimes = ClassPostMeetingTimesModel(meeting_timesMap)
	}
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	source := d.Get("source").(string)

	return &models.ClassPost{
		Description:  description,
		ID:           int32(id),
		MeetingTimes: meetingTimes,
		Name:         &name,
		Site:         site,
		Source:       source,
	}
}

// Retrieve property field names for updating the ClassPost resource
func GetClassPostPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"meeting_times",
		"name",
		"site",
		"source",
	}
}
