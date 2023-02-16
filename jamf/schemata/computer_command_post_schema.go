package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerCommandPost resource defined in the Terraform configuration
func ComputerCommandPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computers": {
			Type: schema.TypeList, //GoType: ComputerCommandPostComputers
			Elem: &schema.Resource{
				Schema: ComputerCommandPostComputersSchema(),
			},
			Optional: true,
		},

		"general": {
			Type: schema.TypeList, //GoType: ComputerCommandPostGeneral
			Elem: &schema.Resource{
				Schema: ComputerCommandPostGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputerCommandPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerCommandPostResourceData(d *schema.ResourceData, m *models.ComputerCommandPost) {
	d.Set("computers", SetComputerCommandPostComputersSubResourceData([]*models.ComputerCommandPostComputers{m.Computers}))
	d.Set("general", SetComputerCommandPostGeneralSubResourceData([]*models.ComputerCommandPostGeneral{m.General}))
}

// Iterate throught and update the ComputerCommandPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerCommandPostSubResourceData(m []*models.ComputerCommandPost) (d []*map[string]interface{}) {
	for _, computerCommandPost := range m {
		if computerCommandPost != nil {
			properties := make(map[string]interface{})
			properties["computers"] = SetComputerCommandPostComputersSubResourceData([]*models.ComputerCommandPostComputers{computerCommandPost.Computers})
			properties["general"] = SetComputerCommandPostGeneralSubResourceData([]*models.ComputerCommandPostGeneral{computerCommandPost.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerCommandPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerCommandPostModel(d *schema.ResourceData) *models.ComputerCommandPost {
	var computers *models.ComputerCommandPostComputers = nil //hit complex
	computersInterface, computersIsSet := d.GetOk("computers")
	if computersIsSet {
		computersMap := computersInterface.([]interface{})[0].(map[string]interface{})
		computers = ComputerCommandPostComputersModel(computersMap)
	}
	var general *models.ComputerCommandPostGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerCommandPostGeneralModel(generalMap)
	}

	return &models.ComputerCommandPost{
		Computers: computers,
		General:   general,
	}
}

// Retrieve property field names for updating the ComputerCommandPost resource
func GetComputerCommandPostPropertyFields() (t []string) {
	return []string{
		"computers",
		"general",
	}
}
