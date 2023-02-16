package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Building resource defined in the Terraform configuration
func BuildingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying Building resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBuildingResourceData(d *schema.ResourceData, m *models.Building) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
}

// Iterate throught and update the Building resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBuildingSubResourceData(m []*models.Building) (d []*map[string]interface{}) {
	for _, building := range m {
		if building != nil {
			properties := make(map[string]interface{})
			properties["id"] = building.ID
			properties["name"] = building.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Building resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func BuildingModel(d *schema.ResourceData) *models.Building {
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)

	return &models.Building{
		ID:   int32(id),
		Name: &name,
	}
}

// Retrieve property field names for updating the Building resource
func GetBuildingPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
