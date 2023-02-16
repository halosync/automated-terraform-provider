package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDevicePost resource defined in the Terraform configuration
func MobileDevicePostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDevicePostGeneral
			Elem: &schema.Resource{
				Schema: MobileDevicePostGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDevicePost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDevicePostResourceData(d *schema.ResourceData, m *models.MobileDevicePost) {
	d.Set("general", SetMobileDevicePostGeneralSubResourceData([]*models.MobileDevicePostGeneral{m.General}))
}

// Iterate throught and update the MobileDevicePost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDevicePostSubResourceData(m []*models.MobileDevicePost) (d []*map[string]interface{}) {
	for _, mobileDevicePost := range m {
		if mobileDevicePost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDevicePostGeneralSubResourceData([]*models.MobileDevicePostGeneral{mobileDevicePost.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDevicePost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDevicePostModel(d *schema.ResourceData) *models.MobileDevicePost {
	var general *models.MobileDevicePostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDevicePostGeneralModel(generalMap)
	}

	return &models.MobileDevicePost{
		General: general,
	}
}

// Retrieve property field names for updating the MobileDevicePost resource
func GetMobileDevicePostPropertyFields() (t []string) {
	return []string{
		"general",
	}
}
