package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceGroups resource defined in the Terraform configuration
func MobileDeviceGroupsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceGroups resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceGroupsResourceData(d *schema.ResourceData, m *models.MobileDeviceGroups) {
}

// Iterate throught and update the MobileDeviceGroups resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceGroupsSubResourceData(m []*models.MobileDeviceGroups) (d []*map[string]interface{}) {
	for _, mobileDeviceGroups := range m {
		if mobileDeviceGroups != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceGroups resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceGroupsModel(d *schema.ResourceData) *models.MobileDeviceGroups {

	return &models.MobileDeviceGroups{}
}

// Retrieve property field names for updating the MobileDeviceGroups resource
func GetMobileDeviceGroupsPropertyFields() (t []string) {
	return []string{}
}
