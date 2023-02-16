package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the JSONWebTokenConfiguration resource defined in the Terraform configuration
func JSONWebTokenConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"token_expiry": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

// Update the underlying JSONWebTokenConfiguration resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetJSONWebTokenConfigurationResourceData(d *schema.ResourceData, m *models.JSONWebTokenConfiguration) {
	d.Set("disabled", m.Disabled)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("token_expiry", m.TokenExpiry)
}

// Iterate throught and update the JSONWebTokenConfiguration resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetJSONWebTokenConfigurationSubResourceData(m []*models.JSONWebTokenConfiguration) (d []*map[string]interface{}) {
	for _, jsonWebTokenConfiguration := range m {
		if jsonWebTokenConfiguration != nil {
			properties := make(map[string]interface{})
			properties["disabled"] = jsonWebTokenConfiguration.Disabled
			properties["id"] = jsonWebTokenConfiguration.ID
			properties["name"] = jsonWebTokenConfiguration.Name
			properties["token_expiry"] = jsonWebTokenConfiguration.TokenExpiry
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate JSONWebTokenConfiguration resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func JSONWebTokenConfigurationModel(d *schema.ResourceData) *models.JSONWebTokenConfiguration {
	disabled := d.Get("disabled").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	tokenExpiry := int64(d.Get("token_expiry").(int))

	return &models.JSONWebTokenConfiguration{
		Disabled:    &disabled,
		ID:          int32(id),
		Name:        &name,
		TokenExpiry: &tokenExpiry,
	}
}

// Retrieve property field names for updating the JSONWebTokenConfiguration resource
func GetJSONWebTokenConfigurationPropertyFields() (t []string) {
	return []string{
		"disabled",
		"id",
		"name",
		"token_expiry",
	}
}
