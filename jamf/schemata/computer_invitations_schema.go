package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerInvitations resource defined in the Terraform configuration
func ComputerInvitationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying ComputerInvitations resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerInvitationsResourceData(d *schema.ResourceData, m *models.ComputerInvitations) {
}

// Iterate throught and update the ComputerInvitations resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerInvitationsSubResourceData(m []*models.ComputerInvitations) (d []*map[string]interface{}) {
	for _, computerInvitations := range m {
		if computerInvitations != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerInvitations resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerInvitationsModel(d *schema.ResourceData) *models.ComputerInvitations {

	return &models.ComputerInvitations{}
}

// Retrieve property field names for updating the ComputerInvitations resource
func GetComputerInvitationsPropertyFields() (t []string) {
	return []string{}
}
