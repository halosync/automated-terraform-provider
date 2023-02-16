package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Commandflush resource defined in the Terraform configuration
func CommandflushSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mobile_devices": {
			Type: schema.TypeList, //GoType: CommandflushMobileDevices
			Elem: &schema.Resource{
				Schema: CommandflushMobileDevicesSchema(),
			},
			Optional: true,
		},

		"status": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying Commandflush resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCommandflushResourceData(d *schema.ResourceData, m *models.Commandflush) {
	d.Set("mobile_devices", SetCommandflushMobileDevicesSubResourceData([]*models.CommandflushMobileDevices{m.MobileDevices}))
	d.Set("status", m.Status)
}

// Iterate throught and update the Commandflush resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCommandflushSubResourceData(m []*models.Commandflush) (d []*map[string]interface{}) {
	for _, commandflush := range m {
		if commandflush != nil {
			properties := make(map[string]interface{})
			properties["mobile_devices"] = SetCommandflushMobileDevicesSubResourceData([]*models.CommandflushMobileDevices{commandflush.MobileDevices})
			properties["status"] = commandflush.Status
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Commandflush resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CommandflushModel(d *schema.ResourceData) *models.Commandflush {
	var mobileDevices *models.CommandflushMobileDevices = nil //hit complex
	mobile_devicesInterface, mobile_devicesIsSet := d.GetOk("mobile_devices")
	if mobile_devicesIsSet {
		mobile_devicesMap := mobile_devicesInterface.([]interface{})[0].(map[string]interface{})
		mobileDevices = CommandflushMobileDevicesModel(mobile_devicesMap)
	}
	status := d.Get("status").(string)

	return &models.Commandflush{
		MobileDevices: mobileDevices,
		Status:        &status,
	}
}

// Retrieve property field names for updating the Commandflush resource
func GetCommandflushPropertyFields() (t []string) {
	return []string{
		"mobile_devices",
		"status",
	}
}
