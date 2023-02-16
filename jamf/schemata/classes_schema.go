package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Classes resource defined in the Terraform configuration
func ClassesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Classes resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClassesResourceData(d *schema.ResourceData, m *models.Classes) {
}

// Iterate throught and update the Classes resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClassesSubResourceData(m []*models.Classes) (d []*map[string]interface{}) {
	for _, classes := range m {
		if classes != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Classes resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClassesModel(d *schema.ResourceData) *models.Classes {

	return &models.Classes{}
}

// Retrieve property field names for updating the Classes resource
func GetClassesPropertyFields() (t []string) {
	return []string{}
}
