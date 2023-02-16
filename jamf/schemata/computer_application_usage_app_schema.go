package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerApplicationUsageApp resource defined in the Terraform configuration
func ComputerApplicationUsageAppSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"foreground": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"open": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"version": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying ComputerApplicationUsageApp resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerApplicationUsageAppResourceData(d *schema.ResourceData, m *models.ComputerApplicationUsageApp) {
	d.Set("foreground", m.Foreground)
	d.Set("name", m.Name)
	d.Set("open", m.Open)
	d.Set("version", m.Version)
}

// Iterate throught and update the ComputerApplicationUsageApp resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerApplicationUsageAppSubResourceData(m []*models.ComputerApplicationUsageApp) (d []*map[string]interface{}) {
	for _, computerApplicationUsageApp := range m {
		if computerApplicationUsageApp != nil {
			properties := make(map[string]interface{})
			properties["foreground"] = computerApplicationUsageApp.Foreground
			properties["name"] = computerApplicationUsageApp.Name
			properties["open"] = computerApplicationUsageApp.Open
			properties["version"] = computerApplicationUsageApp.Version
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerApplicationUsageApp resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerApplicationUsageAppModel(d *schema.ResourceData) *models.ComputerApplicationUsageApp {
	foreground := int64(d.Get("foreground").(int))
	name := d.Get("name").(string)
	open := int64(d.Get("open").(int))
	version := d.Get("version").(string)

	return &models.ComputerApplicationUsageApp{
		Foreground: foreground,
		Name:       name,
		Open:       open,
		Version:    version,
	}
}

// Retrieve property field names for updating the ComputerApplicationUsageApp resource
func GetComputerApplicationUsageAppPropertyFields() (t []string) {
	return []string{
		"foreground",
		"name",
		"open",
		"version",
	}
}
