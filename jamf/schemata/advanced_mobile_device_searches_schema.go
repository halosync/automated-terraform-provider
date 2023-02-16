package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedMobileDeviceSearches resource defined in the Terraform configuration
func AdvancedMobileDeviceSearchesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying AdvancedMobileDeviceSearches resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedMobileDeviceSearchesResourceData(d *schema.ResourceData, m *models.AdvancedMobileDeviceSearches) {
}

// Iterate throught and update the AdvancedMobileDeviceSearches resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedMobileDeviceSearchesSubResourceData(m []*models.AdvancedMobileDeviceSearches) (d []*map[string]interface{}) {
	for _, advancedMobileDeviceSearches := range m {
		if advancedMobileDeviceSearches != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedMobileDeviceSearches resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedMobileDeviceSearchesModel(d *schema.ResourceData) *models.AdvancedMobileDeviceSearches {

	return &models.AdvancedMobileDeviceSearches{}
}

// Retrieve property field names for updating the AdvancedMobileDeviceSearches resource
func GetAdvancedMobileDeviceSearchesPropertyFields() (t []string) {
	return []string{}
}
