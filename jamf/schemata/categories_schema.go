package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Categories resource defined in the Terraform configuration
func CategoriesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Categories resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCategoriesResourceData(d *schema.ResourceData, m *models.Categories) {
}

// Iterate throught and update the Categories resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCategoriesSubResourceData(m []*models.Categories) (d []*map[string]interface{}) {
	for _, categories := range m {
		if categories != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Categories resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CategoriesModel(d *schema.ResourceData) *models.Categories {

	return &models.Categories{}
}

// Retrieve property field names for updating the Categories resource
func GetCategoriesPropertyFields() (t []string) {
	return []string{}
}
