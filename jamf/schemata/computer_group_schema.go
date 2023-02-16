package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerGroup resource defined in the Terraform configuration
func ComputerGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computers": {
			Type: schema.TypeList, //GoType: []*ComputerGroupComputersItems0
			Elem: &schema.Resource{
				Schema: ComputerGroupComputersItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"criteria": {
			Type: schema.TypeList, //GoType: []*ComputerGroupCriteriaItems0
			Elem: &schema.Resource{
				Schema: ComputerGroupCriteriaItems0Schema(),
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
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
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

// Update the underlying ComputerGroup resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerGroupResourceData(d *schema.ResourceData, m *models.ComputerGroup) {
	d.Set("computers", SetComputerGroupComputersItems0SubResourceData(m.Computers))
	d.Set("criteria", SetComputerGroupCriteriaItems0SubResourceData(m.Criteria))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("is_smart", m.IsSmart)
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
}

// Iterate throught and update the ComputerGroup resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerGroupSubResourceData(m []*models.ComputerGroup) (d []*map[string]interface{}) {
	for _, computerGroup := range m {
		if computerGroup != nil {
			properties := make(map[string]interface{})
			properties["computers"] = SetComputerGroupComputersItems0SubResourceData(computerGroup.Computers)
			properties["criteria"] = SetComputerGroupCriteriaItems0SubResourceData(computerGroup.Criteria)
			properties["id"] = computerGroup.ID
			properties["is_smart"] = computerGroup.IsSmart
			properties["name"] = computerGroup.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{computerGroup.Site})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerGroup resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerGroupModel(d *schema.ResourceData) *models.ComputerGroup {
	computers := d.Get("computers").([]*models.ComputerGroupComputersItems0)
	criteria := d.Get("criteria").([]*models.ComputerGroupCriteriaItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
	isSmart := d.Get("is_smart").(bool)
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}

	return &models.ComputerGroup{
		Computers: computers,
		Criteria:  criteria,
		ID:        int32(id),
		IsSmart:   isSmart,
		Name:      name,
		Site:      site,
	}
}

// Retrieve property field names for updating the ComputerGroup resource
func GetComputerGroupPropertyFields() (t []string) {
	return []string{
		"computers",
		"criteria",
		"id",
		"is_smart",
		"name",
		"site",
	}
}
