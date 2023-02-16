package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerGroupPost resource defined in the Terraform configuration
func ComputerGroupPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"criteria": {
			Type: schema.TypeList, //GoType: []*ComputerGroupPostCriteriaItems0
			Elem: &schema.Resource{
				Schema: ComputerGroupPostCriteriaItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"is_smart": {
			Type:     schema.TypeBool,
			Required: true,
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
	}
}

// Update the underlying ComputerGroupPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerGroupPostResourceData(d *schema.ResourceData, m *models.ComputerGroupPost) {
	d.Set("criteria", SetComputerGroupPostCriteriaItems0SubResourceData(m.Criteria))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("is_smart", m.IsSmart)
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
}

// Iterate throught and update the ComputerGroupPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerGroupPostSubResourceData(m []*models.ComputerGroupPost) (d []*map[string]interface{}) {
	for _, computerGroupPost := range m {
		if computerGroupPost != nil {
			properties := make(map[string]interface{})
			properties["criteria"] = SetComputerGroupPostCriteriaItems0SubResourceData(computerGroupPost.Criteria)
			properties["id"] = computerGroupPost.ID
			properties["is_smart"] = computerGroupPost.IsSmart
			properties["name"] = computerGroupPost.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{computerGroupPost.Site})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerGroupPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerGroupPostModel(d *schema.ResourceData) *models.ComputerGroupPost {
	criteria := d.Get("criteria").([]*models.ComputerGroupPostCriteriaItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
	isSmart := d.Get("is_smart").(bool)
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}

	return &models.ComputerGroupPost{
		Criteria: criteria,
		ID:       int32(id),
		IsSmart:  &isSmart,
		Name:     &name,
		Site:     site,
	}
}

// Retrieve property field names for updating the ComputerGroupPost resource
func GetComputerGroupPostPropertyFields() (t []string) {
	return []string{
		"criteria",
		"id",
		"is_smart",
		"name",
		"site",
	}
}
