package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LdapServer resource defined in the Terraform configuration
func LdapServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection": {
			Type: schema.TypeList, //GoType: LdapServerConnection
			Elem: &schema.Resource{
				Schema: LdapServerConnectionSchema(),
			},
			Optional: true,
		},

		"mappings_for_users": {
			Type: schema.TypeList, //GoType: LdapServerMappingsForUsers
			Elem: &schema.Resource{
				Schema: LdapServerMappingsForUsersSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying LdapServer resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLdapServerResourceData(d *schema.ResourceData, m *models.LdapServer) {
	d.Set("connection", SetLdapServerConnectionSubResourceData([]*models.LdapServerConnection{m.Connection}))
	d.Set("mappings_for_users", SetLdapServerMappingsForUsersSubResourceData([]*models.LdapServerMappingsForUsers{m.MappingsForUsers}))
}

// Iterate throught and update the LdapServer resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLdapServerSubResourceData(m []*models.LdapServer) (d []*map[string]interface{}) {
	for _, ldapServer := range m {
		if ldapServer != nil {
			properties := make(map[string]interface{})
			properties["connection"] = SetLdapServerConnectionSubResourceData([]*models.LdapServerConnection{ldapServer.Connection})
			properties["mappings_for_users"] = SetLdapServerMappingsForUsersSubResourceData([]*models.LdapServerMappingsForUsers{ldapServer.MappingsForUsers})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LdapServer resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LdapServerModel(d *schema.ResourceData) *models.LdapServer {
	var connection *models.LdapServerConnection = nil //hit complex
	connectionInterface, connectionIsSet := d.GetOk("connection")
	if connectionIsSet {
		connectionMap := connectionInterface.([]interface{})[0].(map[string]interface{})
		connection = LdapServerConnectionModel(connectionMap)
	}
	var mappingsForUsers *models.LdapServerMappingsForUsers = nil //hit complex
	mappings_for_usersInterface, mappings_for_usersIsSet := d.GetOk("mappings_for_users")
	if mappings_for_usersIsSet {
		mappings_for_usersMap := mappings_for_usersInterface.([]interface{})[0].(map[string]interface{})
		mappingsForUsers = LdapServerMappingsForUsersModel(mappings_for_usersMap)
	}

	return &models.LdapServer{
		Connection:       connection,
		MappingsForUsers: mappingsForUsers,
	}
}

// Retrieve property field names for updating the LdapServer resource
func GetLdapServerPropertyFields() (t []string) {
	return []string{
		"connection",
		"mappings_for_users",
	}
}
