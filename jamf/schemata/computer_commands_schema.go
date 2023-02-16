package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerCommands resource defined in the Terraform configuration
func ComputerCommandsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying ComputerCommands resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerCommandsResourceData(d *schema.ResourceData, m *models.ComputerCommands) {
}

// Iterate throught and update the ComputerCommands resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerCommandsSubResourceData(m []*models.ComputerCommands) (d []*map[string]interface{}) {
	for _, computerCommands := range m {
		if computerCommands != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerCommands resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerCommandsModel(d *schema.ResourceData) *models.ComputerCommands {

	return &models.ComputerCommands{}
}

// Retrieve property field names for updating the ComputerCommands resource
func GetComputerCommandsPropertyFields() (t []string) {
	return []string{}
}
