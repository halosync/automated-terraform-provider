package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LdapServers resource defined in the Terraform configuration
func LdapServersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying LdapServers resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLdapServersResourceData(d *schema.ResourceData, m *models.LdapServers) {
}

// Iterate throught and update the LdapServers resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLdapServersSubResourceData(m []*models.LdapServers) (d []*map[string]interface{}) {
	for _, ldapServers := range m {
		if ldapServers != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LdapServers resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LdapServersModel(d *schema.ResourceData) *models.LdapServers {

	return &models.LdapServers{}
}

// Retrieve property field names for updating the LdapServers resource
func GetLdapServersPropertyFields() (t []string) {
	return []string{}
}
