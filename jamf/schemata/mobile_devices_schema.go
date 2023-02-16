package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDevices resource defined in the Terraform configuration
func MobileDevicesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDevices resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDevicesResourceData(d *schema.ResourceData, m *models.MobileDevices) {
}

// Iterate throught and update the MobileDevices resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDevicesSubResourceData(m []*models.MobileDevices) (d []*map[string]interface{}) {
	for _, mobileDevices := range m {
		if mobileDevices != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDevices resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDevicesModel(d *schema.ResourceData) *models.MobileDevices {

	return &models.MobileDevices{}
}

// Retrieve property field names for updating the MobileDevices resource
func GetMobileDevicesPropertyFields() (t []string) {
	return []string{}
}
