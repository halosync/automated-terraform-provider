package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DockItems resource defined in the Terraform configuration
func DockItemsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying DockItems resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDockItemsResourceData(d *schema.ResourceData, m *models.DockItems) {
}

// Iterate throught and update the DockItems resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDockItemsSubResourceData(m []*models.DockItems) (d []*map[string]interface{}) {
	for _, dockItems := range m {
		if dockItems != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DockItems resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DockItemsModel(d *schema.ResourceData) *models.DockItems {

	return &models.DockItems{}
}

// Retrieve property field names for updating the DockItems resource
func GetDockItemsPropertyFields() (t []string) {
	return []string{}
}
