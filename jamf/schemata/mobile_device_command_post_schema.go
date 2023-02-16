package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceCommandPost resource defined in the Terraform configuration
func MobileDeviceCommandPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceCommandPostGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceCommandPostGeneralSchema(),
			},
			Optional: true,
		},

		"mobile_devices": {
			Type: schema.TypeList, //GoType: MobileDeviceCommandPostMobileDevices
			Elem: &schema.Resource{
				Schema: MobileDeviceCommandPostMobileDevicesSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceCommandPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceCommandPostResourceData(d *schema.ResourceData, m *models.MobileDeviceCommandPost) {
	d.Set("general", SetMobileDeviceCommandPostGeneralSubResourceData([]*models.MobileDeviceCommandPostGeneral{m.General}))
	d.Set("mobile_devices", SetMobileDeviceCommandPostMobileDevicesSubResourceData([]*models.MobileDeviceCommandPostMobileDevices{m.MobileDevices}))
}

// Iterate throught and update the MobileDeviceCommandPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceCommandPostSubResourceData(m []*models.MobileDeviceCommandPost) (d []*map[string]interface{}) {
	for _, mobileDeviceCommandPost := range m {
		if mobileDeviceCommandPost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDeviceCommandPostGeneralSubResourceData([]*models.MobileDeviceCommandPostGeneral{mobileDeviceCommandPost.General})
			properties["mobile_devices"] = SetMobileDeviceCommandPostMobileDevicesSubResourceData([]*models.MobileDeviceCommandPostMobileDevices{mobileDeviceCommandPost.MobileDevices})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceCommandPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceCommandPostModel(d *schema.ResourceData) *models.MobileDeviceCommandPost {
	var general *models.MobileDeviceCommandPostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceCommandPostGeneralModel(generalMap)
	}
	var mobileDevices *models.MobileDeviceCommandPostMobileDevices = nil //hit complex
	mobile_devicesInterface, mobile_devicesIsSet := d.GetOk("mobile_devices")
	if mobile_devicesIsSet {
		mobile_devicesMap := mobile_devicesInterface.([]interface{})[0].(map[string]interface{})
		mobileDevices = MobileDeviceCommandPostMobileDevicesModel(mobile_devicesMap)
	}

	return &models.MobileDeviceCommandPost{
		General:       general,
		MobileDevices: mobileDevices,
	}
}

// Retrieve property field names for updating the MobileDeviceCommandPost resource
func GetMobileDeviceCommandPostPropertyFields() (t []string) {
	return []string{
		"general",
		"mobile_devices",
	}
}
