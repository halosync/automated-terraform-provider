package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceApplication resource defined in the Terraform configuration
func MobileDeviceApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_configuration": {
			Type: schema.TypeList, //GoType: MobileDeviceApplicationAppConfiguration
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationAppConfigurationSchema(),
			},
			Optional: true,
		},

		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceApplicationGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: MobileDeviceApplicationScope
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationScopeSchema(),
			},
			Optional: true,
		},

		"self_service": {
			Type: schema.TypeList, //GoType: MobileDeviceApplicationSelfService
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationSelfServiceSchema(),
			},
			Optional: true,
		},

		"vpp": {
			Type: schema.TypeList, //GoType: MobileDeviceApplicationVpp
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationVppSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceApplication resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceApplicationResourceData(d *schema.ResourceData, m *models.MobileDeviceApplication) {
	d.Set("app_configuration", SetMobileDeviceApplicationAppConfigurationSubResourceData([]*models.MobileDeviceApplicationAppConfiguration{m.AppConfiguration}))
	d.Set("general", SetMobileDeviceApplicationGeneralSubResourceData([]*models.MobileDeviceApplicationGeneral{m.General}))
	d.Set("scope", SetMobileDeviceApplicationScopeSubResourceData([]*models.MobileDeviceApplicationScope{m.Scope}))
	d.Set("self_service", SetMobileDeviceApplicationSelfServiceSubResourceData([]*models.MobileDeviceApplicationSelfService{m.SelfService}))
	d.Set("vpp", SetMobileDeviceApplicationVppSubResourceData([]*models.MobileDeviceApplicationVpp{m.Vpp}))
}

// Iterate throught and update the MobileDeviceApplication resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceApplicationSubResourceData(m []*models.MobileDeviceApplication) (d []*map[string]interface{}) {
	for _, mobileDeviceApplication := range m {
		if mobileDeviceApplication != nil {
			properties := make(map[string]interface{})
			properties["app_configuration"] = SetMobileDeviceApplicationAppConfigurationSubResourceData([]*models.MobileDeviceApplicationAppConfiguration{mobileDeviceApplication.AppConfiguration})
			properties["general"] = SetMobileDeviceApplicationGeneralSubResourceData([]*models.MobileDeviceApplicationGeneral{mobileDeviceApplication.General})
			properties["scope"] = SetMobileDeviceApplicationScopeSubResourceData([]*models.MobileDeviceApplicationScope{mobileDeviceApplication.Scope})
			properties["self_service"] = SetMobileDeviceApplicationSelfServiceSubResourceData([]*models.MobileDeviceApplicationSelfService{mobileDeviceApplication.SelfService})
			properties["vpp"] = SetMobileDeviceApplicationVppSubResourceData([]*models.MobileDeviceApplicationVpp{mobileDeviceApplication.Vpp})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceApplication resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceApplicationModel(d *schema.ResourceData) *models.MobileDeviceApplication {
	var appConfiguration *models.MobileDeviceApplicationAppConfiguration = nil //hit complex
	app_configurationInterface, app_configurationIsSet := d.GetOk("app_configuration")
	if app_configurationIsSet {
		app_configurationMap := app_configurationInterface.([]interface{})[0].(map[string]interface{})
		appConfiguration = MobileDeviceApplicationAppConfigurationModel(app_configurationMap)
	}
	var general *models.MobileDeviceApplicationGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceApplicationGeneralModel(generalMap)
	}
	var scope *models.MobileDeviceApplicationScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = MobileDeviceApplicationScopeModel(scopeMap)
	}
	var selfService *models.MobileDeviceApplicationSelfService = nil //hit complex
	self_serviceInterface, self_serviceIsSet := d.GetOk("self_service")
	if self_serviceIsSet {
		self_serviceMap := self_serviceInterface.([]interface{})[0].(map[string]interface{})
		selfService = MobileDeviceApplicationSelfServiceModel(self_serviceMap)
	}
	var vpp *models.MobileDeviceApplicationVpp = nil //hit complex
	vppInterface, vppIsSet := d.GetOk("vpp")
	if vppIsSet {
		vppMap := vppInterface.([]interface{})[0].(map[string]interface{})
		vpp = MobileDeviceApplicationVppModel(vppMap)
	}

	return &models.MobileDeviceApplication{
		AppConfiguration: appConfiguration,
		General:          general,
		Scope:            scope,
		SelfService:      selfService,
		Vpp:              vpp,
	}
}

// Retrieve property field names for updating the MobileDeviceApplication resource
func GetMobileDeviceApplicationPropertyFields() (t []string) {
	return []string{
		"app_configuration",
		"general",
		"scope",
		"self_service",
		"vpp",
	}
}
