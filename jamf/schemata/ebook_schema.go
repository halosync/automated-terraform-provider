package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Ebook resource defined in the Terraform configuration
func EbookSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: EbookGeneral
			Elem: &schema.Resource{
				Schema: EbookGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: EbookScope
			Elem: &schema.Resource{
				Schema: EbookScopeSchema(),
			},
			Optional: true,
		},

		"self_service": {
			Type: schema.TypeList, //GoType: EbookSelfService
			Elem: &schema.Resource{
				Schema: EbookSelfServiceSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Ebook resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEbookResourceData(d *schema.ResourceData, m *models.Ebook) {
	d.Set("general", SetEbookGeneralSubResourceData([]*models.EbookGeneral{m.General}))
	d.Set("scope", SetEbookScopeSubResourceData([]*models.EbookScope{m.Scope}))
	d.Set("self_service", SetEbookSelfServiceSubResourceData([]*models.EbookSelfService{m.SelfService}))
}

// Iterate throught and update the Ebook resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEbookSubResourceData(m []*models.Ebook) (d []*map[string]interface{}) {
	for _, ebook := range m {
		if ebook != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetEbookGeneralSubResourceData([]*models.EbookGeneral{ebook.General})
			properties["scope"] = SetEbookScopeSubResourceData([]*models.EbookScope{ebook.Scope})
			properties["self_service"] = SetEbookSelfServiceSubResourceData([]*models.EbookSelfService{ebook.SelfService})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Ebook resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EbookModel(d *schema.ResourceData) *models.Ebook {
	var general *models.EbookGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = EbookGeneralModel(generalMap)
	}
	var scope *models.EbookScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = EbookScopeModel(scopeMap)
	}
	var selfService *models.EbookSelfService = nil //hit complex
	self_serviceInterface, self_serviceIsSet := d.GetOk("self_service")
	if self_serviceIsSet {
		self_serviceMap := self_serviceInterface.([]interface{})[0].(map[string]interface{})
		selfService = EbookSelfServiceModel(self_serviceMap)
	}

	return &models.Ebook{
		General:     general,
		Scope:       scope,
		SelfService: selfService,
	}
}

// Retrieve property field names for updating the Ebook resource
func GetEbookPropertyFields() (t []string) {
	return []string{
		"general",
		"scope",
		"self_service",
	}
}
