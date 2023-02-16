package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputersBasic resource defined in the Terraform configuration
func ComputersBasicSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computer": {
			Type: schema.TypeList, //GoType: ComputersBasicComputer
			Elem: &schema.Resource{
				Schema: ComputersBasicComputerSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputersBasic resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputersBasicResourceData(d *schema.ResourceData, m *models.ComputersBasic) {
	d.Set("computer", SetComputersBasicComputerSubResourceData([]*models.ComputersBasicComputer{m.Computer}))
}

// Iterate throught and update the ComputersBasic resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputersBasicSubResourceData(m []*models.ComputersBasic) (d []*map[string]interface{}) {
	for _, computersBasic := range m {
		if computersBasic != nil {
			properties := make(map[string]interface{})
			properties["computer"] = SetComputersBasicComputerSubResourceData([]*models.ComputersBasicComputer{computersBasic.Computer})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputersBasic resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputersBasicModel(d *schema.ResourceData) *models.ComputersBasic {
	var computer *models.ComputersBasicComputer = nil //hit complex
	computerInterface, computerIsSet := d.GetOk("computer")
	if computerIsSet {
		computerMap := computerInterface.([]interface{})[0].(map[string]interface{})
		computer = ComputersBasicComputerModel(computerMap)
	}

	return &models.ComputersBasic{
		Computer: computer,
	}
}

// Retrieve property field names for updating the ComputersBasic resource
func GetComputersBasicPropertyFields() (t []string) {
	return []string{
		"computer",
	}
}
