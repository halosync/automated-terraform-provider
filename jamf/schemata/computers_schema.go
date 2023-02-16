package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Computers resource defined in the Terraform configuration
func ComputersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Computers resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputersResourceData(d *schema.ResourceData, m *models.Computers) {
}

// Iterate throught and update the Computers resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputersSubResourceData(m []*models.Computers) (d []*map[string]interface{}) {
	for _, computers := range m {
		if computers != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Computers resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputersModel(d *schema.ResourceData) *models.Computers {

	return &models.Computers{}
}

// Retrieve property field names for updating the Computers resource
func GetComputersPropertyFields() (t []string) {
	return []string{}
}
