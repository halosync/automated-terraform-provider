package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Buildings resource defined in the Terraform configuration
func BuildingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Buildings resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBuildingsResourceData(d *schema.ResourceData, m *models.Buildings) {
}

// Iterate throught and update the Buildings resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBuildingsSubResourceData(m []*models.Buildings) (d []*map[string]interface{}) {
	for _, buildings := range m {
		if buildings != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Buildings resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func BuildingsModel(d *schema.ResourceData) *models.Buildings {

	return &models.Buildings{}
}

// Retrieve property field names for updating the Buildings resource
func GetBuildingsPropertyFields() (t []string) {
	return []string{}
}
