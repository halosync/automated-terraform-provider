package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AllowedFileExtension resource defined in the Terraform configuration
func AllowedFileExtensionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"extension": {
			Type:     schema.TypeString,
			Required: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Update the underlying AllowedFileExtension resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAllowedFileExtensionResourceData(d *schema.ResourceData, m *models.AllowedFileExtension) {
	d.Set("extension", m.Extension)
	d.Set("id", strconv.Itoa(int(m.ID)))
}

// Iterate throught and update the AllowedFileExtension resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAllowedFileExtensionSubResourceData(m []*models.AllowedFileExtension) (d []*map[string]interface{}) {
	for _, allowedFileExtension := range m {
		if allowedFileExtension != nil {
			properties := make(map[string]interface{})
			properties["extension"] = allowedFileExtension.Extension
			properties["id"] = allowedFileExtension.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AllowedFileExtension resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AllowedFileExtensionModel(d *schema.ResourceData) *models.AllowedFileExtension {
	extension := d.Get("extension").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))

	return &models.AllowedFileExtension{
		Extension: &extension,
		ID:        int32(id),
	}
}

// Retrieve property field names for updating the AllowedFileExtension resource
func GetAllowedFileExtensionPropertyFields() (t []string) {
	return []string{
		"extension",
		"id",
	}
}
