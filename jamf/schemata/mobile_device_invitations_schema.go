package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceInvitations resource defined in the Terraform configuration
func MobileDeviceInvitationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceInvitations resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceInvitationsResourceData(d *schema.ResourceData, m *models.MobileDeviceInvitations) {
}

// Iterate throught and update the MobileDeviceInvitations resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceInvitationsSubResourceData(m []*models.MobileDeviceInvitations) (d []*map[string]interface{}) {
	for _, mobileDeviceInvitations := range m {
		if mobileDeviceInvitations != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceInvitations resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceInvitationsModel(d *schema.ResourceData) *models.MobileDeviceInvitations {

	return &models.MobileDeviceInvitations{}
}

// Retrieve property field names for updating the MobileDeviceInvitations resource
func GetMobileDeviceInvitationsPropertyFields() (t []string) {
	return []string{}
}
