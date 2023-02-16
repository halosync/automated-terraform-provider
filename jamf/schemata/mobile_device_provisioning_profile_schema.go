package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceProvisioningProfile resource defined in the Terraform configuration
func MobileDeviceProvisioningProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceProvisioningProfileGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceProvisioningProfileGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceProvisioningProfile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceProvisioningProfileResourceData(d *schema.ResourceData, m *models.MobileDeviceProvisioningProfile) {
	d.Set("general", SetMobileDeviceProvisioningProfileGeneralSubResourceData([]*models.MobileDeviceProvisioningProfileGeneral{m.General}))
}

// Iterate throught and update the MobileDeviceProvisioningProfile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceProvisioningProfileSubResourceData(m []*models.MobileDeviceProvisioningProfile) (d []*map[string]interface{}) {
	for _, mobileDeviceProvisioningProfile := range m {
		if mobileDeviceProvisioningProfile != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDeviceProvisioningProfileGeneralSubResourceData([]*models.MobileDeviceProvisioningProfileGeneral{mobileDeviceProvisioningProfile.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceProvisioningProfile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceProvisioningProfileModel(d *schema.ResourceData) *models.MobileDeviceProvisioningProfile {
	var general *models.MobileDeviceProvisioningProfileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceProvisioningProfileGeneralModel(generalMap)
	}

	return &models.MobileDeviceProvisioningProfile{
		General: general,
	}
}

// Retrieve property field names for updating the MobileDeviceProvisioningProfile resource
func GetMobileDeviceProvisioningProfilePropertyFields() (t []string) {
	return []string{
		"general",
	}
}
