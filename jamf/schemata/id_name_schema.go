package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the IDName resource defined in the Terraform configuration
func IDNameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying IDName resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIDNameResourceData(d *schema.ResourceData, m *models.IDName) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
}

// Iterate throught and update the IDName resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIDNameSubResourceData(m []*models.IDName) (d []*map[string]interface{}) {
	for _, idName := range m {
		if idName != nil {
			properties := make(map[string]interface{})
			properties["id"] = idName.ID
			properties["name"] = idName.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate IDName resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IDNameModel(d *schema.ResourceData) *models.IDName {
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)

	return &models.IDName{
		ID:   int32(id),
		Name: name,
	}
}

// Retrieve property field names for updating the IDName resource
func GetIDNamePropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
