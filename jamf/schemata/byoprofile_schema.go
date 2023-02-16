package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Byoprofile resource defined in the Terraform configuration
func ByoprofileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: ByoprofileGeneral
			Elem: &schema.Resource{
				Schema: ByoprofileGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Byoprofile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetByoprofileResourceData(d *schema.ResourceData, m *models.Byoprofile) {
	d.Set("general", SetByoprofileGeneralSubResourceData([]*models.ByoprofileGeneral{m.General}))
}

// Iterate throught and update the Byoprofile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetByoprofileSubResourceData(m []*models.Byoprofile) (d []*map[string]interface{}) {
	for _, byoprofile := range m {
		if byoprofile != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetByoprofileGeneralSubResourceData([]*models.ByoprofileGeneral{byoprofile.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Byoprofile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ByoprofileModel(d *schema.ResourceData) *models.Byoprofile {
	var general *models.ByoprofileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ByoprofileGeneralModel(generalMap)
	}

	return &models.Byoprofile{
		General: general,
	}
}

// Retrieve property field names for updating the Byoprofile resource
func GetByoprofilePropertyFields() (t []string) {
	return []string{
		"general",
	}
}
