package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the OsxConfigurationProfiles resource defined in the Terraform configuration
func OsxConfigurationProfilesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying OsxConfigurationProfiles resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetOsxConfigurationProfilesResourceData(d *schema.ResourceData, m *models.OsxConfigurationProfiles) {
}

// Iterate throught and update the OsxConfigurationProfiles resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetOsxConfigurationProfilesSubResourceData(m []*models.OsxConfigurationProfiles) (d []*map[string]interface{}) {
	for _, osxConfigurationProfiles := range m {
		if osxConfigurationProfiles != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate OsxConfigurationProfiles resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func OsxConfigurationProfilesModel(d *schema.ResourceData) *models.OsxConfigurationProfiles {

	return &models.OsxConfigurationProfiles{}
}

// Retrieve property field names for updating the OsxConfigurationProfiles resource
func GetOsxConfigurationProfilesPropertyFields() (t []string) {
	return []string{}
}
