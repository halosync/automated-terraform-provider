package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerExtensionAttribute resource defined in the Terraform configuration
func ComputerExtensionAttributeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"input_type": {
			Type: schema.TypeList, //GoType: ComputerExtensionAttributeInputType
			Elem: &schema.Resource{
				Schema: ComputerExtensionAttributeInputTypeSchema(),
			},
			Optional: true,
		},

		"inventory_display": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"recon_display": {
			Type:     schema.TypeString,
			Default:  "Extension Attributes",
			Optional: true,
		},
	}
}

// Update the underlying ComputerExtensionAttribute resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerExtensionAttributeResourceData(d *schema.ResourceData, m *models.ComputerExtensionAttribute) {
	d.Set("data_type", m.DataType)
	d.Set("description", m.Description)
	d.Set("enabled", m.Enabled)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("input_type", SetComputerExtensionAttributeInputTypeSubResourceData([]*models.ComputerExtensionAttributeInputType{m.InputType}))
	d.Set("inventory_display", m.InventoryDisplay)
	d.Set("name", m.Name)
	d.Set("recon_display", m.ReconDisplay)
}

// Iterate throught and update the ComputerExtensionAttribute resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerExtensionAttributeSubResourceData(m []*models.ComputerExtensionAttribute) (d []*map[string]interface{}) {
	for _, computerExtensionAttribute := range m {
		if computerExtensionAttribute != nil {
			properties := make(map[string]interface{})
			properties["data_type"] = computerExtensionAttribute.DataType
			properties["description"] = computerExtensionAttribute.Description
			properties["enabled"] = computerExtensionAttribute.Enabled
			properties["id"] = computerExtensionAttribute.ID
			properties["input_type"] = SetComputerExtensionAttributeInputTypeSubResourceData([]*models.ComputerExtensionAttributeInputType{computerExtensionAttribute.InputType})
			properties["inventory_display"] = computerExtensionAttribute.InventoryDisplay
			properties["name"] = computerExtensionAttribute.Name
			properties["recon_display"] = computerExtensionAttribute.ReconDisplay
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerExtensionAttribute resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerExtensionAttributeModel(d *schema.ResourceData) *models.ComputerExtensionAttribute {
	dataType := d.Get("data_type").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	var inputType *models.ComputerExtensionAttributeInputType = nil //hit complex
	input_typeInterface, input_typeIsSet := d.GetOk("input_type")
	if input_typeIsSet {
		input_typeMap := input_typeInterface.([]interface{})[0].(map[string]interface{})
		inputType = ComputerExtensionAttributeInputTypeModel(input_typeMap)
	}
	inventoryDisplay := d.Get("inventory_display").(string)
	name := d.Get("name").(string)
	reconDisplay := d.Get("recon_display").(string)

	return &models.ComputerExtensionAttribute{
		DataType:         dataType,
		Description:      description,
		Enabled:          enabled,
		ID:               int32(id),
		InputType:        inputType,
		InventoryDisplay: inventoryDisplay,
		Name:             &name,
		ReconDisplay:     &reconDisplay,
	}
}

// Retrieve property field names for updating the ComputerExtensionAttribute resource
func GetComputerExtensionAttributePropertyFields() (t []string) {
	return []string{
		"data_type",
		"description",
		"enabled",
		"id",
		"input_type",
		"inventory_display",
		"name",
		"recon_display",
	}
}
