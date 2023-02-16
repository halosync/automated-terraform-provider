package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceCommand resource defined in the Terraform configuration
func MobileDeviceCommandSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceCommandGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceCommandGeneralSchema(),
			},
			Optional: true,
		},

		"mobile_devices": {
			Type: schema.TypeList, //GoType: MobileDeviceCommandMobileDevices
			Elem: &schema.Resource{
				Schema: MobileDeviceCommandMobileDevicesSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceCommand resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceCommandResourceData(d *schema.ResourceData, m *models.MobileDeviceCommand) {
	d.Set("general", SetMobileDeviceCommandGeneralSubResourceData([]*models.MobileDeviceCommandGeneral{m.General}))
	d.Set("mobile_devices", SetMobileDeviceCommandMobileDevicesSubResourceData([]*models.MobileDeviceCommandMobileDevices{m.MobileDevices}))
}

// Iterate throught and update the MobileDeviceCommand resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceCommandSubResourceData(m []*models.MobileDeviceCommand) (d []*map[string]interface{}) {
	for _, mobileDeviceCommand := range m {
		if mobileDeviceCommand != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDeviceCommandGeneralSubResourceData([]*models.MobileDeviceCommandGeneral{mobileDeviceCommand.General})
			properties["mobile_devices"] = SetMobileDeviceCommandMobileDevicesSubResourceData([]*models.MobileDeviceCommandMobileDevices{mobileDeviceCommand.MobileDevices})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceCommand resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceCommandModel(d *schema.ResourceData) *models.MobileDeviceCommand {
	var general *models.MobileDeviceCommandGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceCommandGeneralModel(generalMap)
	}
	var mobileDevices *models.MobileDeviceCommandMobileDevices = nil //hit complex
	mobile_devicesInterface, mobile_devicesIsSet := d.GetOk("mobile_devices")
	if mobile_devicesIsSet {
		mobile_devicesMap := mobile_devicesInterface.([]interface{})[0].(map[string]interface{})
		mobileDevices = MobileDeviceCommandMobileDevicesModel(mobile_devicesMap)
	}

	return &models.MobileDeviceCommand{
		General:       general,
		MobileDevices: mobileDevices,
	}
}

// Retrieve property field names for updating the MobileDeviceCommand resource
func GetMobileDeviceCommandPropertyFields() (t []string) {
	return []string{
		"general",
		"mobile_devices",
	}
}
