package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceEnrollmentProfilePost resource defined in the Terraform configuration
func MobileDeviceEnrollmentProfilePostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceEnrollmentProfilePostGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceEnrollmentProfilePostGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceEnrollmentProfilePost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceEnrollmentProfilePostResourceData(d *schema.ResourceData, m *models.MobileDeviceEnrollmentProfilePost) {
	d.Set("general", SetMobileDeviceEnrollmentProfilePostGeneralSubResourceData([]*models.MobileDeviceEnrollmentProfilePostGeneral{m.General}))
}

// Iterate throught and update the MobileDeviceEnrollmentProfilePost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceEnrollmentProfilePostSubResourceData(m []*models.MobileDeviceEnrollmentProfilePost) (d []*map[string]interface{}) {
	for _, mobileDeviceEnrollmentProfilePost := range m {
		if mobileDeviceEnrollmentProfilePost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDeviceEnrollmentProfilePostGeneralSubResourceData([]*models.MobileDeviceEnrollmentProfilePostGeneral{mobileDeviceEnrollmentProfilePost.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceEnrollmentProfilePost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceEnrollmentProfilePostModel(d *schema.ResourceData) *models.MobileDeviceEnrollmentProfilePost {
	var general *models.MobileDeviceEnrollmentProfilePostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceEnrollmentProfilePostGeneralModel(generalMap)
	}

	return &models.MobileDeviceEnrollmentProfilePost{
		General: general,
	}
}

// Retrieve property field names for updating the MobileDeviceEnrollmentProfilePost resource
func GetMobileDeviceEnrollmentProfilePostPropertyFields() (t []string) {
	return []string{
		"general",
	}
}
