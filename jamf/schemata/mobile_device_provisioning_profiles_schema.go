package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceProvisioningProfiles resource defined in the Terraform configuration
func MobileDeviceProvisioningProfilesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceProvisioningProfiles resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceProvisioningProfilesResourceData(d *schema.ResourceData, m *models.MobileDeviceProvisioningProfiles) {
}

// Iterate throught and update the MobileDeviceProvisioningProfiles resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceProvisioningProfilesSubResourceData(m []*models.MobileDeviceProvisioningProfiles) (d []*map[string]interface{}) {
	for _, mobileDeviceProvisioningProfiles := range m {
		if mobileDeviceProvisioningProfiles != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceProvisioningProfiles resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceProvisioningProfilesModel(d *schema.ResourceData) *models.MobileDeviceProvisioningProfiles {

	return &models.MobileDeviceProvisioningProfiles{}
}

// Retrieve property field names for updating the MobileDeviceProvisioningProfiles resource
func GetMobileDeviceProvisioningProfilesPropertyFields() (t []string) {
	return []string{}
}
