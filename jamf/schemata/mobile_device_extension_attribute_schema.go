package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceExtensionAttribute resource defined in the Terraform configuration
func MobileDeviceExtensionAttributeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"date_type": {
			Type:     schema.TypeString,
			Default:  "String",
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"input_type": {
			Type: schema.TypeList, //GoType: MobileDeviceExtensionAttributeInputType
			Elem: &schema.Resource{
				Schema: MobileDeviceExtensionAttributeInputTypeSchema(),
			},
			Optional: true,
		},

		"inventory_display": {
			Type:     schema.TypeString,
			Default:  "General",
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying MobileDeviceExtensionAttribute resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceExtensionAttributeResourceData(d *schema.ResourceData, m *models.MobileDeviceExtensionAttribute) {
	d.Set("date_type", m.DateType)
	d.Set("description", m.Description)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("input_type", SetMobileDeviceExtensionAttributeInputTypeSubResourceData([]*models.MobileDeviceExtensionAttributeInputType{m.InputType}))
	d.Set("inventory_display", m.InventoryDisplay)
	d.Set("name", m.Name)
}

// Iterate throught and update the MobileDeviceExtensionAttribute resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceExtensionAttributeSubResourceData(m []*models.MobileDeviceExtensionAttribute) (d []*map[string]interface{}) {
	for _, mobileDeviceExtensionAttribute := range m {
		if mobileDeviceExtensionAttribute != nil {
			properties := make(map[string]interface{})
			properties["date_type"] = mobileDeviceExtensionAttribute.DateType
			properties["description"] = mobileDeviceExtensionAttribute.Description
			properties["id"] = mobileDeviceExtensionAttribute.ID
			properties["input_type"] = SetMobileDeviceExtensionAttributeInputTypeSubResourceData([]*models.MobileDeviceExtensionAttributeInputType{mobileDeviceExtensionAttribute.InputType})
			properties["inventory_display"] = mobileDeviceExtensionAttribute.InventoryDisplay
			properties["name"] = mobileDeviceExtensionAttribute.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceExtensionAttribute resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceExtensionAttributeModel(d *schema.ResourceData) *models.MobileDeviceExtensionAttribute {
	dateType := d.Get("date_type").(string)
	description := d.Get("description").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	var inputType *models.MobileDeviceExtensionAttributeInputType = nil //hit complex
	input_typeInterface, input_typeIsSet := d.GetOk("input_type")
	if input_typeIsSet {
		input_typeMap := input_typeInterface.([]interface{})[0].(map[string]interface{})
		inputType = MobileDeviceExtensionAttributeInputTypeModel(input_typeMap)
	}
	inventoryDisplay := d.Get("inventory_display").(string)
	name := d.Get("name").(string)

	return &models.MobileDeviceExtensionAttribute{
		DateType:         &dateType,
		Description:      description,
		ID:               int32(id),
		InputType:        inputType,
		InventoryDisplay: &inventoryDisplay,
		Name:             &name,
	}
}

// Retrieve property field names for updating the MobileDeviceExtensionAttribute resource
func GetMobileDeviceExtensionAttributePropertyFields() (t []string) {
	return []string{
		"date_type",
		"description",
		"id",
		"input_type",
		"inventory_display",
		"name",
	}
}
