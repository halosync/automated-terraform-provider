package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedUserSearches resource defined in the Terraform configuration
func AdvancedUserSearchesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying AdvancedUserSearches resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedUserSearchesResourceData(d *schema.ResourceData, m *models.AdvancedUserSearches) {
}

// Iterate throught and update the AdvancedUserSearches resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedUserSearchesSubResourceData(m []*models.AdvancedUserSearches) (d []*map[string]interface{}) {
	for _, advancedUserSearches := range m {
		if advancedUserSearches != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedUserSearches resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedUserSearchesModel(d *schema.ResourceData) *models.AdvancedUserSearches {

	return &models.AdvancedUserSearches{}
}

// Retrieve property field names for updating the AdvancedUserSearches resource
func GetAdvancedUserSearchesPropertyFields() (t []string) {
	return []string{}
}
