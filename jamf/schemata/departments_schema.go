package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Departments resource defined in the Terraform configuration
func DepartmentsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Departments resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDepartmentsResourceData(d *schema.ResourceData, m *models.Departments) {
}

// Iterate throught and update the Departments resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDepartmentsSubResourceData(m []*models.Departments) (d []*map[string]interface{}) {
	for _, departments := range m {
		if departments != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Departments resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DepartmentsModel(d *schema.ResourceData) *models.Departments {

	return &models.Departments{}
}

// Retrieve property field names for updating the Departments resource
func GetDepartmentsPropertyFields() (t []string) {
	return []string{}
}
