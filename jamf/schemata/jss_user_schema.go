package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the JssUser resource defined in the Terraform configuration
func JssUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"institution": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"license_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"privileges": {
			Type: schema.TypeList, //GoType: []*JssUserPrivilegesItems0
			Elem: &schema.Resource{
				Schema: JssUserPrivilegesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"product": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"version": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying JssUser resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetJssUserResourceData(d *schema.ResourceData, m *models.JssUser) {
	d.Set("institution", m.Institution)
	d.Set("license_type", m.LicenseType)
	d.Set("privileges", SetJssUserPrivilegesItems0SubResourceData(m.Privileges))
	d.Set("product", m.Product)
	d.Set("version", m.Version)
}

// Iterate throught and update the JssUser resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetJssUserSubResourceData(m []*models.JssUser) (d []*map[string]interface{}) {
	for _, jssUser := range m {
		if jssUser != nil {
			properties := make(map[string]interface{})
			properties["institution"] = jssUser.Institution
			properties["license_type"] = jssUser.LicenseType
			properties["privileges"] = SetJssUserPrivilegesItems0SubResourceData(jssUser.Privileges)
			properties["product"] = jssUser.Product
			properties["version"] = jssUser.Version
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate JssUser resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func JssUserModel(d *schema.ResourceData) *models.JssUser {
	institution := d.Get("institution").(string)
	licenseType := d.Get("license_type").(string)
	privileges := d.Get("privileges").([]*models.JssUserPrivilegesItems0)
	product := d.Get("product").(string)
	version := d.Get("version").(string)

	return &models.JssUser{
		Institution: institution,
		LicenseType: licenseType,
		Privileges:  privileges,
		Product:     product,
		Version:     version,
	}
}

// Retrieve property field names for updating the JssUser resource
func GetJssUserPropertyFields() (t []string) {
	return []string{
		"institution",
		"license_type",
		"privileges",
		"product",
		"version",
	}
}
