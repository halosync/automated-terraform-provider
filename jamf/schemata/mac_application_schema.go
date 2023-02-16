package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MacApplication resource defined in the Terraform configuration
func MacApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: MacApplicationGeneral
			Elem: &schema.Resource{
				Schema: MacApplicationGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: MacApplicationScope
			Elem: &schema.Resource{
				Schema: MacApplicationScopeSchema(),
			},
			Optional: true,
		},

		"self_service": {
			Type: schema.TypeList, //GoType: MacApplicationSelfService
			Elem: &schema.Resource{
				Schema: MacApplicationSelfServiceSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MacApplication resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMacApplicationResourceData(d *schema.ResourceData, m *models.MacApplication) {
	d.Set("general", SetMacApplicationGeneralSubResourceData([]*models.MacApplicationGeneral{m.General}))
	d.Set("scope", SetMacApplicationScopeSubResourceData([]*models.MacApplicationScope{m.Scope}))
	d.Set("self_service", SetMacApplicationSelfServiceSubResourceData([]*models.MacApplicationSelfService{m.SelfService}))
}

// Iterate throught and update the MacApplication resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMacApplicationSubResourceData(m []*models.MacApplication) (d []*map[string]interface{}) {
	for _, macApplication := range m {
		if macApplication != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetMacApplicationGeneralSubResourceData([]*models.MacApplicationGeneral{macApplication.General})
			properties["scope"] = SetMacApplicationScopeSubResourceData([]*models.MacApplicationScope{macApplication.Scope})
			properties["self_service"] = SetMacApplicationSelfServiceSubResourceData([]*models.MacApplicationSelfService{macApplication.SelfService})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MacApplication resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MacApplicationModel(d *schema.ResourceData) *models.MacApplication {
	var general *models.MacApplicationGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MacApplicationGeneralModel(generalMap)
	}
	var scope *models.MacApplicationScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = MacApplicationScopeModel(scopeMap)
	}
	var selfService *models.MacApplicationSelfService = nil //hit complex
	self_serviceInterface, self_serviceIsSet := d.GetOk("self_service")
	if self_serviceIsSet {
		self_serviceMap := self_serviceInterface.([]interface{})[0].(map[string]interface{})
		selfService = MacApplicationSelfServiceModel(self_serviceMap)
	}

	return &models.MacApplication{
		General:     general,
		Scope:       scope,
		SelfService: selfService,
	}
}

// Retrieve property field names for updating the MacApplication resource
func GetMacApplicationPropertyFields() (t []string) {
	return []string{
		"general",
		"scope",
		"self_service",
	}
}
