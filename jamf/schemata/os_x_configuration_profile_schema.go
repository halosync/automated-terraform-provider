package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the OsxConfigurationProfile resource defined in the Terraform configuration
func OsxConfigurationProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: OsxConfigurationProfileGeneral
			Elem: &schema.Resource{
				Schema: OsxConfigurationProfileGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: OsxConfigurationProfileScope
			Elem: &schema.Resource{
				Schema: OsxConfigurationProfileScopeSchema(),
			},
			Optional: true,
		},

		"self_service": {
			Type: schema.TypeList, //GoType: OsxConfigurationProfileSelfService
			Elem: &schema.Resource{
				Schema: OsxConfigurationProfileSelfServiceSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying OsxConfigurationProfile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetOsxConfigurationProfileResourceData(d *schema.ResourceData, m *models.OsxConfigurationProfile) {
	d.Set("general", SetOsxConfigurationProfileGeneralSubResourceData([]*models.OsxConfigurationProfileGeneral{m.General}))
	d.Set("scope", SetOsxConfigurationProfileScopeSubResourceData([]*models.OsxConfigurationProfileScope{m.Scope}))
	d.Set("self_service", SetOsxConfigurationProfileSelfServiceSubResourceData([]*models.OsxConfigurationProfileSelfService{m.SelfService}))
}

// Iterate throught and update the OsxConfigurationProfile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetOsxConfigurationProfileSubResourceData(m []*models.OsxConfigurationProfile) (d []*map[string]interface{}) {
	for _, osxConfigurationProfile := range m {
		if osxConfigurationProfile != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetOsxConfigurationProfileGeneralSubResourceData([]*models.OsxConfigurationProfileGeneral{osxConfigurationProfile.General})
			properties["scope"] = SetOsxConfigurationProfileScopeSubResourceData([]*models.OsxConfigurationProfileScope{osxConfigurationProfile.Scope})
			properties["self_service"] = SetOsxConfigurationProfileSelfServiceSubResourceData([]*models.OsxConfigurationProfileSelfService{osxConfigurationProfile.SelfService})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate OsxConfigurationProfile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func OsxConfigurationProfileModel(d *schema.ResourceData) *models.OsxConfigurationProfile {
	var general *models.OsxConfigurationProfileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = OsxConfigurationProfileGeneralModel(generalMap)
	}
	var scope *models.OsxConfigurationProfileScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = OsxConfigurationProfileScopeModel(scopeMap)
	}
	var selfService *models.OsxConfigurationProfileSelfService = nil //hit complex
	self_serviceInterface, self_serviceIsSet := d.GetOk("self_service")
	if self_serviceIsSet {
		self_serviceMap := self_serviceInterface.([]interface{})[0].(map[string]interface{})
		selfService = OsxConfigurationProfileSelfServiceModel(self_serviceMap)
	}

	return &models.OsxConfigurationProfile{
		General:     general,
		Scope:       scope,
		SelfService: selfService,
	}
}

// Retrieve property field names for updating the OsxConfigurationProfile resource
func GetOsxConfigurationProfilePropertyFields() (t []string) {
	return []string{
		"general",
		"scope",
		"self_service",
	}
}
