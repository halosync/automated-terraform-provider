package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerCommand resource defined in the Terraform configuration
func ComputerCommandSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computers": {
			Type: schema.TypeList, //GoType: ComputerCommandComputers
			Elem: &schema.Resource{
				Schema: ComputerCommandComputersSchema(),
			},
			Optional: true,
		},

		"general": {
			Type: schema.TypeList, //GoType: ComputerCommandGeneral
			Elem: &schema.Resource{
				Schema: ComputerCommandGeneralSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputerCommand resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerCommandResourceData(d *schema.ResourceData, m *models.ComputerCommand) {
	d.Set("computers", SetComputerCommandComputersSubResourceData([]*models.ComputerCommandComputers{m.Computers}))
	d.Set("general", SetComputerCommandGeneralSubResourceData([]*models.ComputerCommandGeneral{m.General}))
}

// Iterate throught and update the ComputerCommand resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerCommandSubResourceData(m []*models.ComputerCommand) (d []*map[string]interface{}) {
	for _, computerCommand := range m {
		if computerCommand != nil {
			properties := make(map[string]interface{})
			properties["computers"] = SetComputerCommandComputersSubResourceData([]*models.ComputerCommandComputers{computerCommand.Computers})
			properties["general"] = SetComputerCommandGeneralSubResourceData([]*models.ComputerCommandGeneral{computerCommand.General})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerCommand resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerCommandModel(d *schema.ResourceData) *models.ComputerCommand {
	var computers *models.ComputerCommandComputers = nil //hit complex
	computersInterface, computersIsSet := d.GetOk("computers")
	if computersIsSet {
		computersMap := computersInterface.([]interface{})[0].(map[string]interface{})
		computers = ComputerCommandComputersModel(computersMap)
	}
	var general *models.ComputerCommandGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerCommandGeneralModel(generalMap)
	}

	return &models.ComputerCommand{
		Computers: computers,
		General:   general,
	}
}

// Retrieve property field names for updating the ComputerCommand resource
func GetComputerCommandPropertyFields() (t []string) {
	return []string{
		"computers",
		"general",
	}
}
