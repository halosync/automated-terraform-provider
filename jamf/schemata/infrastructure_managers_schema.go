package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the InfrastructureManagers resource defined in the Terraform configuration
func InfrastructureManagersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying InfrastructureManagers resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetInfrastructureManagersResourceData(d *schema.ResourceData, m *models.InfrastructureManagers) {
}

// Iterate throught and update the InfrastructureManagers resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetInfrastructureManagersSubResourceData(m []*models.InfrastructureManagers) (d []*map[string]interface{}) {
	for _, infrastructureManagers := range m {
		if infrastructureManagers != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate InfrastructureManagers resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func InfrastructureManagersModel(d *schema.ResourceData) *models.InfrastructureManagers {

	return &models.InfrastructureManagers{}
}

// Retrieve property field names for updating the InfrastructureManagers resource
func GetInfrastructureManagersPropertyFields() (t []string) {
	return []string{}
}
