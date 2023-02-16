package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerApplicationsVersion resource defined in the Terraform configuration
func ComputerApplicationsVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying ComputerApplicationsVersion resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerApplicationsVersionResourceData(d *schema.ResourceData, m *models.ComputerApplicationsVersion) {
}

// Iterate throught and update the ComputerApplicationsVersion resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerApplicationsVersionSubResourceData(m []*models.ComputerApplicationsVersion) (d []*map[string]interface{}) {
	for _, computerApplicationsVersion := range m {
		if computerApplicationsVersion != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerApplicationsVersion resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerApplicationsVersionModel(d *schema.ResourceData) *models.ComputerApplicationsVersion {

	return &models.ComputerApplicationsVersion{}
}

// Retrieve property field names for updating the ComputerApplicationsVersion resource
func GetComputerApplicationsVersionPropertyFields() (t []string) {
	return []string{}
}
