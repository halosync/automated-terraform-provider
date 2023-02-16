package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedUserSearch resource defined in the Terraform configuration
func AdvancedUserSearchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"criteria": {
			Type: schema.TypeList, //GoType: []*AdvancedUserSearchCriteriaItems0
			Elem: &schema.Resource{
				Schema: AdvancedUserSearchCriteriaItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"display_fields": {
			Type: schema.TypeList, //GoType: []*AdvancedUserSearchDisplayFieldsItems0
			Elem: &schema.Resource{
				Schema: AdvancedUserSearchDisplayFieldsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
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

		"users": {
			Type: schema.TypeList, //GoType: []*AdvancedUserSearchUsersItems0
			Elem: &schema.Resource{
				Schema: AdvancedUserSearchUsersItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying AdvancedUserSearch resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedUserSearchResourceData(d *schema.ResourceData, m *models.AdvancedUserSearch) {
	d.Set("criteria", SetAdvancedUserSearchCriteriaItems0SubResourceData(m.Criteria))
	d.Set("display_fields", SetAdvancedUserSearchDisplayFieldsItems0SubResourceData(m.DisplayFields))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("users", SetAdvancedUserSearchUsersItems0SubResourceData(m.Users))
}

// Iterate throught and update the AdvancedUserSearch resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedUserSearchSubResourceData(m []*models.AdvancedUserSearch) (d []*map[string]interface{}) {
	for _, advancedUserSearch := range m {
		if advancedUserSearch != nil {
			properties := make(map[string]interface{})
			properties["criteria"] = SetAdvancedUserSearchCriteriaItems0SubResourceData(advancedUserSearch.Criteria)
			properties["display_fields"] = SetAdvancedUserSearchDisplayFieldsItems0SubResourceData(advancedUserSearch.DisplayFields)
			properties["id"] = advancedUserSearch.ID
			properties["name"] = advancedUserSearch.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{advancedUserSearch.Site})
			properties["users"] = SetAdvancedUserSearchUsersItems0SubResourceData(advancedUserSearch.Users)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedUserSearch resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedUserSearchModel(d *schema.ResourceData) *models.AdvancedUserSearch {
	criteria := d.Get("criteria").([]*models.AdvancedUserSearchCriteriaItems0)
	displayFields := d.Get("display_fields").([]*models.AdvancedUserSearchDisplayFieldsItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	users := d.Get("users").([]*models.AdvancedUserSearchUsersItems0)

	return &models.AdvancedUserSearch{
		Criteria:      criteria,
		DisplayFields: displayFields,
		ID:            int32(id),
		Name:          &name,
		Site:          site,
		Users:         users,
	}
}

// Retrieve property field names for updating the AdvancedUserSearch resource
func GetAdvancedUserSearchPropertyFields() (t []string) {
	return []string{
		"criteria",
		"display_fields",
		"id",
		"name",
		"site",
		"users",
	}
}
