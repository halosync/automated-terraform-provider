package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LicensedSoftwareDefintion resource defined in the Terraform configuration
func LicensedSoftwareDefintionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"compare_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"version": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying LicensedSoftwareDefintion resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLicensedSoftwareDefintionResourceData(d *schema.ResourceData, m *models.LicensedSoftwareDefintion) {
	d.Set("compare_type", m.CompareType)
	d.Set("name", m.Name)
	d.Set("version", m.Version)
}

// Iterate throught and update the LicensedSoftwareDefintion resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLicensedSoftwareDefintionSubResourceData(m []*models.LicensedSoftwareDefintion) (d []*map[string]interface{}) {
	for _, licensedSoftwareDefintion := range m {
		if licensedSoftwareDefintion != nil {
			properties := make(map[string]interface{})
			properties["compare_type"] = licensedSoftwareDefintion.CompareType
			properties["name"] = licensedSoftwareDefintion.Name
			properties["version"] = licensedSoftwareDefintion.Version
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LicensedSoftwareDefintion resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LicensedSoftwareDefintionModel(d *schema.ResourceData) *models.LicensedSoftwareDefintion {
	compareType := d.Get("compare_type").(string)
	name := d.Get("name").(string)
	version := d.Get("version").(string)

	return &models.LicensedSoftwareDefintion{
		CompareType: compareType,
		Name:        name,
		Version:     version,
	}
}

// Retrieve property field names for updating the LicensedSoftwareDefintion resource
func GetLicensedSoftwareDefintionPropertyFields() (t []string) {
	return []string{
		"compare_type",
		"name",
		"version",
	}
}
