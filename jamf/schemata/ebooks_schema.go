package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Ebooks resource defined in the Terraform configuration
func EbooksSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Ebooks resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEbooksResourceData(d *schema.ResourceData, m *models.Ebooks) {
}

// Iterate throught and update the Ebooks resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEbooksSubResourceData(m []*models.Ebooks) (d []*map[string]interface{}) {
	for _, ebooks := range m {
		if ebooks != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Ebooks resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EbooksModel(d *schema.ResourceData) *models.Ebooks {

	return &models.Ebooks{}
}

// Retrieve property field names for updating the Ebooks resource
func GetEbooksPropertyFields() (t []string) {
	return []string{}
}
