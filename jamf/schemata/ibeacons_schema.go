package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Ibeacons resource defined in the Terraform configuration
func IbeaconsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying Ibeacons resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIbeaconsResourceData(d *schema.ResourceData, m *models.Ibeacons) {
}

// Iterate throught and update the Ibeacons resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIbeaconsSubResourceData(m []*models.Ibeacons) (d []*map[string]interface{}) {
	for _, ibeacons := range m {
		if ibeacons != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Ibeacons resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IbeaconsModel(d *schema.ResourceData) *models.Ibeacons {

	return &models.Ibeacons{}
}

// Retrieve property field names for updating the Ibeacons resource
func GetIbeaconsPropertyFields() (t []string) {
	return []string{}
}
