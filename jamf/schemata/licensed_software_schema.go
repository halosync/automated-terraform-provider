package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LicensedSoftware resource defined in the Terraform configuration
func LicensedSoftwareSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"font_definitions": {
			Type: schema.TypeList, //GoType: []*LicensedSoftwareFontDefinitionsItems0
			Elem: &schema.Resource{
				Schema: LicensedSoftwareFontDefinitionsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: LicensedSoftwareGeneral
			Elem: &schema.Resource{
				Schema: LicensedSoftwareGeneralSchema(),
			},
			Optional: true,
		},

		"licenses": {
			Type: schema.TypeList, //GoType: []*LicensedSoftwareLicensesItems0
			Elem: &schema.Resource{
				Schema: LicensedSoftwareLicensesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"plugin_definitions": {
			Type: schema.TypeList, //GoType: []*LicensedSoftwarePluginDefinitionsItems0
			Elem: &schema.Resource{
				Schema: LicensedSoftwarePluginDefinitionsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"software_definitions": {
			Type: schema.TypeList, //GoType: []*LicensedSoftwareSoftwareDefinitionsItems0
			Elem: &schema.Resource{
				Schema: LicensedSoftwareSoftwareDefinitionsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying LicensedSoftware resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLicensedSoftwareResourceData(d *schema.ResourceData, m *models.LicensedSoftware) {
	d.Set("font_definitions", SetLicensedSoftwareFontDefinitionsItems0SubResourceData(m.FontDefinitions))
	d.Set("general", SetLicensedSoftwareGeneralSubResourceData([]*models.LicensedSoftwareGeneral{m.General}))
	d.Set("licenses", SetLicensedSoftwareLicensesItems0SubResourceData(m.Licenses))
	d.Set("plugin_definitions", SetLicensedSoftwarePluginDefinitionsItems0SubResourceData(m.PluginDefinitions))
	d.Set("software_definitions", SetLicensedSoftwareSoftwareDefinitionsItems0SubResourceData(m.SoftwareDefinitions))
}

// Iterate throught and update the LicensedSoftware resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLicensedSoftwareSubResourceData(m []*models.LicensedSoftware) (d []*map[string]interface{}) {
	for _, licensedSoftware := range m {
		if licensedSoftware != nil {
			properties := make(map[string]interface{})
			properties["font_definitions"] = SetLicensedSoftwareFontDefinitionsItems0SubResourceData(licensedSoftware.FontDefinitions)
			properties["general"] = SetLicensedSoftwareGeneralSubResourceData([]*models.LicensedSoftwareGeneral{licensedSoftware.General})
			properties["licenses"] = SetLicensedSoftwareLicensesItems0SubResourceData(licensedSoftware.Licenses)
			properties["plugin_definitions"] = SetLicensedSoftwarePluginDefinitionsItems0SubResourceData(licensedSoftware.PluginDefinitions)
			properties["software_definitions"] = SetLicensedSoftwareSoftwareDefinitionsItems0SubResourceData(licensedSoftware.SoftwareDefinitions)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LicensedSoftware resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LicensedSoftwareModel(d *schema.ResourceData) *models.LicensedSoftware {
	fontDefinitions := d.Get("font_definitions").([]*models.LicensedSoftwareFontDefinitionsItems0)
	var general *models.LicensedSoftwareGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = LicensedSoftwareGeneralModel(generalMap)
	}
	licenses := d.Get("licenses").([]*models.LicensedSoftwareLicensesItems0)
	pluginDefinitions := d.Get("plugin_definitions").([]*models.LicensedSoftwarePluginDefinitionsItems0)
	softwareDefinitions := d.Get("software_definitions").([]*models.LicensedSoftwareSoftwareDefinitionsItems0)

	return &models.LicensedSoftware{
		FontDefinitions:     fontDefinitions,
		General:             general,
		Licenses:            licenses,
		PluginDefinitions:   pluginDefinitions,
		SoftwareDefinitions: softwareDefinitions,
	}
}

// Retrieve property field names for updating the LicensedSoftware resource
func GetLicensedSoftwarePropertyFields() (t []string) {
	return []string{
		"font_definitions",
		"general",
		"licenses",
		"plugin_definitions",
		"software_definitions",
	}
}
