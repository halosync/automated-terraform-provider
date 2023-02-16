package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceApplications resource defined in the Terraform configuration
func MobileDeviceApplicationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceApplications resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceApplicationsResourceData(d *schema.ResourceData, m *models.MobileDeviceApplications) {
}

// Iterate throught and update the MobileDeviceApplications resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceApplicationsSubResourceData(m []*models.MobileDeviceApplications) (d []*map[string]interface{}) {
	for _, mobileDeviceApplications := range m {
		if mobileDeviceApplications != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceApplications resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceApplicationsModel(d *schema.ResourceData) *models.MobileDeviceApplications {

	return &models.MobileDeviceApplications{}
}

// Retrieve property field names for updating the MobileDeviceApplications resource
func GetMobileDeviceApplicationsPropertyFields() (t []string) {
	return []string{}
}
