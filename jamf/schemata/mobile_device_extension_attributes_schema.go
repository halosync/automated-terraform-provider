package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceExtensionAttributes resource defined in the Terraform configuration
func MobileDeviceExtensionAttributesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying MobileDeviceExtensionAttributes resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceExtensionAttributesResourceData(d *schema.ResourceData, m *models.MobileDeviceExtensionAttributes) {
}

// Iterate throught and update the MobileDeviceExtensionAttributes resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceExtensionAttributesSubResourceData(m []*models.MobileDeviceExtensionAttributes) (d []*map[string]interface{}) {
	for _, mobileDeviceExtensionAttributes := range m {
		if mobileDeviceExtensionAttributes != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceExtensionAttributes resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceExtensionAttributesModel(d *schema.ResourceData) *models.MobileDeviceExtensionAttributes {

	return &models.MobileDeviceExtensionAttributes{}
}

// Retrieve property field names for updating the MobileDeviceExtensionAttributes resource
func GetMobileDeviceExtensionAttributesPropertyFields() (t []string) {
	return []string{}
}
