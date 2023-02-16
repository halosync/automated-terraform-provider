package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceCommands resource defined in the Terraform configuration
func MobileDeviceCommandsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceCommands resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceCommandsResourceData(d *schema.ResourceData, m *models.MobileDeviceCommands) {
}

// Iterate throught and update the MobileDeviceCommands resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceCommandsSubResourceData(m []*models.MobileDeviceCommands) (d []*map[string]interface{}) {
	for _, mobileDeviceCommands := range m {
		if mobileDeviceCommands != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceCommands resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceCommandsModel(d *schema.ResourceData) *models.MobileDeviceCommands {

	return &models.MobileDeviceCommands{}
}

// Retrieve property field names for updating the MobileDeviceCommands resource
func GetMobileDeviceCommandsPropertyFields() (t []string) {
	return []string{}
}
