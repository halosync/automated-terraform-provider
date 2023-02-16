package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerPost resource defined in the Terraform configuration
func ComputerPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"general": {
			Type: schema.TypeList, //GoType: ComputerPostGeneral
			Elem: &schema.Resource{
				Schema: ComputerPostGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputerPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerPostResourceData(d *schema.ResourceData, m *models.ComputerPost) {
	d.Set("general", SetComputerPostGeneralSubResourceData([]*models.ComputerPostGeneral{m.General}))
}

// Iterate throught and update the ComputerPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerPostSubResourceData(m []*models.ComputerPost) (d []*map[string]interface{}) {
	for _, computerPost := range m {
		if computerPost != nil {
			properties := make(map[string]interface{})
			properties["general"] = SetComputerPostGeneralSubResourceData([]*models.ComputerPostGeneral{computerPost.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerPostModel(d *schema.ResourceData) *models.ComputerPost {
	var general *models.ComputerPostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerPostGeneralModel(generalMap)
	}

	return &models.ComputerPost{
		General: general,
	}
}

// Retrieve property field names for updating the ComputerPost resource
func GetComputerPostPropertyFields() (t []string) {
	return []string{
		"general",
	}
}
