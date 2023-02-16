package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ManagedPreferenceProfilePost resource defined in the Terraform configuration
func ManagedPreferenceProfilePostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: ManagedPreferenceProfilePostGeneral
			Elem: &schema.Resource{
				Schema: ManagedPreferenceProfilePostGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ManagedPreferenceProfilePost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetManagedPreferenceProfilePostResourceData(d *schema.ResourceData, m *models.ManagedPreferenceProfilePost) {
	d.Set("general", SetManagedPreferenceProfilePostGeneralSubResourceData([]*models.ManagedPreferenceProfilePostGeneral{m.General}))
}

// Iterate throught and update the ManagedPreferenceProfilePost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetManagedPreferenceProfilePostSubResourceData(m []*models.ManagedPreferenceProfilePost) (d []*map[string]interface{}) {
	for _, managedPreferenceProfilePost := range m {
		if managedPreferenceProfilePost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetManagedPreferenceProfilePostGeneralSubResourceData([]*models.ManagedPreferenceProfilePostGeneral{managedPreferenceProfilePost.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ManagedPreferenceProfilePost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ManagedPreferenceProfilePostModel(d *schema.ResourceData) *models.ManagedPreferenceProfilePost {
	var general *models.ManagedPreferenceProfilePostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ManagedPreferenceProfilePostGeneralModel(generalMap)
	}

	return &models.ManagedPreferenceProfilePost{
		General: general,
	}
}

// Retrieve property field names for updating the ManagedPreferenceProfilePost resource
func GetManagedPreferenceProfilePostPropertyFields() (t []string) {
	return []string{
		"general",
	}
}
