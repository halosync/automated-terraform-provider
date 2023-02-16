package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ActivationCode resource defined in the Terraform configuration
func ActivationCodeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"code": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"organization_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying ActivationCode resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetActivationCodeResourceData(d *schema.ResourceData, m *models.ActivationCode) {
	d.Set("code", m.Code)
	d.Set("organization_name", m.OrganizationName)
}

// Iterate throught and update the ActivationCode resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetActivationCodeSubResourceData(m []*models.ActivationCode) (d []*map[string]interface{}) {
	for _, activationCode := range m {
		if activationCode != nil {
			properties := make(map[string]interface{})
			properties["code"] = activationCode.Code
			properties["organization_name"] = activationCode.OrganizationName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ActivationCode resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ActivationCodeModel(d *schema.ResourceData) *models.ActivationCode {
	code := d.Get("code").(string)
	organizationName := d.Get("organization_name").(string)

	return &models.ActivationCode{
		Code:             code,
		OrganizationName: organizationName,
	}
}

// Retrieve property field names for updating the ActivationCode resource
func GetActivationCodePropertyFields() (t []string) {
	return []string{
		"code",
		"organization_name",
	}
}
