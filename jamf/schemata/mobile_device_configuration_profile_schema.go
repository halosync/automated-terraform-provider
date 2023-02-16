package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceConfigurationProfile resource defined in the Terraform configuration
func MobileDeviceConfigurationProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceConfigurationProfileGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceConfigurationProfileGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: MobileDeviceConfigurationProfileScope
			Elem: &schema.Resource{
				Schema: MobileDeviceConfigurationProfileScopeSchema(),
			},
			Optional: true,
		},

		"self_service": {
			Type: schema.TypeList, //GoType: MobileDeviceConfigurationProfileSelfService
			Elem: &schema.Resource{
				Schema: MobileDeviceConfigurationProfileSelfServiceSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceConfigurationProfile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceConfigurationProfileResourceData(d *schema.ResourceData, m *models.MobileDeviceConfigurationProfile) {
	d.Set("general", SetMobileDeviceConfigurationProfileGeneralSubResourceData([]*models.MobileDeviceConfigurationProfileGeneral{m.General}))
	d.Set("scope", SetMobileDeviceConfigurationProfileScopeSubResourceData([]*models.MobileDeviceConfigurationProfileScope{m.Scope}))
	d.Set("self_service", SetMobileDeviceConfigurationProfileSelfServiceSubResourceData([]*models.MobileDeviceConfigurationProfileSelfService{m.SelfService}))
}

// Iterate throught and update the MobileDeviceConfigurationProfile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceConfigurationProfileSubResourceData(m []*models.MobileDeviceConfigurationProfile) (d []*map[string]interface{}) {
	for _, mobileDeviceConfigurationProfile := range m {
		if mobileDeviceConfigurationProfile != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMobileDeviceConfigurationProfileGeneralSubResourceData([]*models.MobileDeviceConfigurationProfileGeneral{mobileDeviceConfigurationProfile.General})
			properties["scope"] = SetMobileDeviceConfigurationProfileScopeSubResourceData([]*models.MobileDeviceConfigurationProfileScope{mobileDeviceConfigurationProfile.Scope})
			properties["self_service"] = SetMobileDeviceConfigurationProfileSelfServiceSubResourceData([]*models.MobileDeviceConfigurationProfileSelfService{mobileDeviceConfigurationProfile.SelfService})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceConfigurationProfile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceConfigurationProfileModel(d *schema.ResourceData) *models.MobileDeviceConfigurationProfile {
	var general *models.MobileDeviceConfigurationProfileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceConfigurationProfileGeneralModel(generalMap)
	}
	var scope *models.MobileDeviceConfigurationProfileScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = MobileDeviceConfigurationProfileScopeModel(scopeMap)
	}
	var selfService *models.MobileDeviceConfigurationProfileSelfService = nil //hit complex
	self_serviceInterface, self_serviceIsSet := d.GetOk("self_service")
	if self_serviceIsSet {
		self_serviceMap := self_serviceInterface.([]interface{})[0].(map[string]interface{})
		selfService = MobileDeviceConfigurationProfileSelfServiceModel(self_serviceMap)
	}

	return &models.MobileDeviceConfigurationProfile{
		General:     general,
		Scope:       scope,
		SelfService: selfService,
	}
}

// Retrieve property field names for updating the MobileDeviceConfigurationProfile resource
func GetMobileDeviceConfigurationProfilePropertyFields() (t []string) {
	return []string{
		"general",
		"scope",
		"self_service",
	}
}
