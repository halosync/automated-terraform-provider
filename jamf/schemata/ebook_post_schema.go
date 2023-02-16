package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the EbookPost resource defined in the Terraform configuration
func EbookPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: EbookPostGeneral
			Elem: &schema.Resource{
				Schema: EbookPostGeneralSchema(),
			},
			Optional: true,
		},

		"scope": {
			Type: schema.TypeList, //GoType: EbookPostScope
			Elem: &schema.Resource{
				Schema: EbookPostScopeSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying EbookPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEbookPostResourceData(d *schema.ResourceData, m *models.EbookPost) {
	d.Set("general", SetEbookPostGeneralSubResourceData([]*models.EbookPostGeneral{m.General}))
	d.Set("scope", SetEbookPostScopeSubResourceData([]*models.EbookPostScope{m.Scope}))
}

// Iterate throught and update the EbookPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEbookPostSubResourceData(m []*models.EbookPost) (d []*map[string]interface{}) {
	for _, ebookPost := range m {
		if ebookPost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetEbookPostGeneralSubResourceData([]*models.EbookPostGeneral{ebookPost.General})
			properties["scope"] = SetEbookPostScopeSubResourceData([]*models.EbookPostScope{ebookPost.Scope})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate EbookPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EbookPostModel(d *schema.ResourceData) *models.EbookPost {
	var general *models.EbookPostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = EbookPostGeneralModel(generalMap)
	}
	var scope *models.EbookPostScope = nil //hit complex
	scopeInterface, scopeIsSet := d.GetOk("scope")
	if scopeIsSet {
		scopeMap := scopeInterface.([]interface{})[0].(map[string]interface{})
		scope = EbookPostScopeModel(scopeMap)
	}

	return &models.EbookPost{
		General: general,
		Scope:   scope,
	}
}

// Retrieve property field names for updating the EbookPost resource
func GetEbookPostPropertyFields() (t []string) {
	return []string{
		"general",
		"scope",
	}
}
