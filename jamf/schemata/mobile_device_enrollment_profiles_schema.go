package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceEnrollmentProfiles resource defined in the Terraform configuration
func MobileDeviceEnrollmentProfilesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceEnrollmentProfiles resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceEnrollmentProfilesResourceData(d *schema.ResourceData, m *models.MobileDeviceEnrollmentProfiles) {
}

// Iterate throught and update the MobileDeviceEnrollmentProfiles resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceEnrollmentProfilesSubResourceData(m []*models.MobileDeviceEnrollmentProfiles) (d []*map[string]interface{}) {
	for _, mobileDeviceEnrollmentProfiles := range m {
		if mobileDeviceEnrollmentProfiles != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceEnrollmentProfiles resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceEnrollmentProfilesModel(d *schema.ResourceData) *models.MobileDeviceEnrollmentProfiles {

	return &models.MobileDeviceEnrollmentProfiles{}
}

// Retrieve property field names for updating the MobileDeviceEnrollmentProfiles resource
func GetMobileDeviceEnrollmentProfilesPropertyFields() (t []string) {
	return []string{}
}
