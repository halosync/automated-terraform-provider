package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the CategoryObject resource defined in the Terraform configuration
func CategoryObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Default:  "Unknown",
			Required: true,
		},
	}
}

// Update the underlying CategoryObject resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCategoryObjectResourceData(d *schema.ResourceData, m *models.CategoryObject) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
}

// Iterate throught and update the CategoryObject resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCategoryObjectSubResourceData(m []*models.CategoryObject) (d []*map[string]interface{}) {
	for _, categoryObject := range m {
		if categoryObject != nil {
			properties := make(map[string]interface{})
			properties["id"] = categoryObject.ID
			properties["name"] = categoryObject.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CategoryObject resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CategoryObjectModel(d *schema.ResourceData) *models.CategoryObject {
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)

	return &models.CategoryObject{
		ID:   int32(id),
		Name: &name,
	}
}

// Retrieve property field names for updating the CategoryObject resource
func GetCategoryObjectPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
