package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerExtensionAttributes resource defined in the Terraform configuration
func ComputerExtensionAttributesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying ComputerExtensionAttributes resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerExtensionAttributesResourceData(d *schema.ResourceData, m *models.ComputerExtensionAttributes) {
}

// Iterate throught and update the ComputerExtensionAttributes resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerExtensionAttributesSubResourceData(m []*models.ComputerExtensionAttributes) (d []*map[string]interface{}) {
	for _, computerExtensionAttributes := range m {
		if computerExtensionAttributes != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerExtensionAttributes resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerExtensionAttributesModel(d *schema.ResourceData) *models.ComputerExtensionAttributes {

	return &models.ComputerExtensionAttributes{}
}

// Retrieve property field names for updating the ComputerExtensionAttributes resource
func GetComputerExtensionAttributesPropertyFields() (t []string) {
	return []string{}
}
