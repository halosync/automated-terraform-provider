package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedComputerSearches resource defined in the Terraform configuration
func AdvancedComputerSearchesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying AdvancedComputerSearches resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedComputerSearchesResourceData(d *schema.ResourceData, m *models.AdvancedComputerSearches) {
}

// Iterate throught and update the AdvancedComputerSearches resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedComputerSearchesSubResourceData(m []*models.AdvancedComputerSearches) (d []*map[string]interface{}) {
	for _, advancedComputerSearches := range m {
		if advancedComputerSearches != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedComputerSearches resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedComputerSearchesModel(d *schema.ResourceData) *models.AdvancedComputerSearches {

	return &models.AdvancedComputerSearches{}
}

// Retrieve property field names for updating the AdvancedComputerSearches resource
func GetAdvancedComputerSearchesPropertyFields() (t []string) {
	return []string{}
}
