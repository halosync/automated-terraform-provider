package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ManagedPreferenceProfile resource defined in the Terraform configuration
func ManagedPreferenceProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: ManagedPreferenceProfileGeneral
			Elem: &schema.Resource{
				Schema: ManagedPreferenceProfileGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ManagedPreferenceProfile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetManagedPreferenceProfileResourceData(d *schema.ResourceData, m *models.ManagedPreferenceProfile) {
	d.Set("general", SetManagedPreferenceProfileGeneralSubResourceData([]*models.ManagedPreferenceProfileGeneral{m.General}))
}

// Iterate throught and update the ManagedPreferenceProfile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetManagedPreferenceProfileSubResourceData(m []*models.ManagedPreferenceProfile) (d []*map[string]interface{}) {
	for _, managedPreferenceProfile := range m {
		if managedPreferenceProfile != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetManagedPreferenceProfileGeneralSubResourceData([]*models.ManagedPreferenceProfileGeneral{managedPreferenceProfile.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ManagedPreferenceProfile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ManagedPreferenceProfileModel(d *schema.ResourceData) *models.ManagedPreferenceProfile {
	var general *models.ManagedPreferenceProfileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ManagedPreferenceProfileGeneralModel(generalMap)
	}

	return &models.ManagedPreferenceProfile{
		General: general,
	}
}

// Retrieve property field names for updating the ManagedPreferenceProfile resource
func GetManagedPreferenceProfilePropertyFields() (t []string) {
	return []string{
		"general",
	}
}
