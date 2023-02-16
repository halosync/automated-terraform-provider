package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the GsxConnection resource defined in the Terraform configuration
func GsxConnectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_number": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"region": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"uri": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying GsxConnection resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGsxConnectionResourceData(d *schema.ResourceData, m *models.GsxConnection) {
	d.Set("account_number", m.AccountNumber)
	d.Set("enabled", m.Enabled)
	d.Set("region", m.Region)
	d.Set("uri", m.URI)
	d.Set("username", m.Username)
}

// Iterate throught and update the GsxConnection resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGsxConnectionSubResourceData(m []*models.GsxConnection) (d []*map[string]interface{}) {
	for _, gsxConnection := range m {
		if gsxConnection != nil {
			properties := make(map[string]interface{})
			properties["account_number"] = gsxConnection.AccountNumber
			properties["enabled"] = gsxConnection.Enabled
			properties["region"] = gsxConnection.Region
			properties["uri"] = gsxConnection.URI
			properties["username"] = gsxConnection.Username
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate GsxConnection resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GsxConnectionModel(d *schema.ResourceData) *models.GsxConnection {
	accountNumber := int64(d.Get("account_number").(int))
	enabled := d.Get("enabled").(bool)
	region := d.Get("region").(string)
	uri := d.Get("uri").(string)
	username := d.Get("username").(string)

	return &models.GsxConnection{
		AccountNumber: accountNumber,
		Enabled:       enabled,
		Region:        region,
		URI:           uri,
		Username:      username,
	}
}

// Retrieve property field names for updating the GsxConnection resource
func GetGsxConnectionPropertyFields() (t []string) {
	return []string{
		"account_number",
		"enabled",
		"region",
		"uri",
		"username",
	}
}
