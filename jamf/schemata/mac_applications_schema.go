package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MacApplications resource defined in the Terraform configuration
func MacApplicationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MacApplications resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMacApplicationsResourceData(d *schema.ResourceData, m *models.MacApplications) {
}

// Iterate throught and update the MacApplications resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMacApplicationsSubResourceData(m []*models.MacApplications) (d []*map[string]interface{}) {
	for _, macApplications := range m {
		if macApplications != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MacApplications resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MacApplicationsModel(d *schema.ResourceData) *models.MacApplications {

	return &models.MacApplications{}
}

// Retrieve property field names for updating the MacApplications resource
func GetMacApplicationsPropertyFields() (t []string) {
	return []string{}
}
