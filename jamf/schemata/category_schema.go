package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Category resource defined in the Terraform configuration
func CategorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

// Update the underlying Category resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCategoryResourceData(d *schema.ResourceData, m *models.Category) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("priority", m.Priority)
}

// Iterate throught and update the Category resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCategorySubResourceData(m []*models.Category) (d []*map[string]interface{}) {
	for _, category := range m {
		if category != nil {
			properties := make(map[string]interface{})
			properties["id"] = category.ID
			properties["name"] = category.Name
			properties["priority"] = category.Priority
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Category resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CategoryModel(d *schema.ResourceData) *models.Category {
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	priority := int64(d.Get("priority").(int))

	return &models.Category{
		ID:       int32(id),
		Name:     &name,
		Priority: &priority,
	}
}

// Retrieve property field names for updating the Category resource
func GetCategoryPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"priority",
	}
}
