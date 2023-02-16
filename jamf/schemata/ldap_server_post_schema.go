package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LdapServerPost resource defined in the Terraform configuration
func LdapServerPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection": {
			Type: schema.TypeList, //GoType: LdapServerPostConnection
			Elem: &schema.Resource{
				Schema: LdapServerPostConnectionSchema(),
			},
			Optional: true,
		},

		"mappings_for_users": {
			Type: schema.TypeList, //GoType: LdapServerPostMappingsForUsers
			Elem: &schema.Resource{
				Schema: LdapServerPostMappingsForUsersSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying LdapServerPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLdapServerPostResourceData(d *schema.ResourceData, m *models.LdapServerPost) {
	d.Set("connection", SetLdapServerPostConnectionSubResourceData([]*models.LdapServerPostConnection{m.Connection}))
	d.Set("mappings_for_users", SetLdapServerPostMappingsForUsersSubResourceData([]*models.LdapServerPostMappingsForUsers{m.MappingsForUsers}))
}

// Iterate throught and update the LdapServerPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLdapServerPostSubResourceData(m []*models.LdapServerPost) (d []*map[string]interface{}) {
	for _, ldapServerPost := range m {
		if ldapServerPost != nil {
			properties := make(map[string]interface{})
			properties["connection"] = SetLdapServerPostConnectionSubResourceData([]*models.LdapServerPostConnection{ldapServerPost.Connection})
			properties["mappings_for_users"] = SetLdapServerPostMappingsForUsersSubResourceData([]*models.LdapServerPostMappingsForUsers{ldapServerPost.MappingsForUsers})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LdapServerPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LdapServerPostModel(d *schema.ResourceData) *models.LdapServerPost {
	var connection *models.LdapServerPostConnection = nil //hit complex
	connectionInterface, connectionIsSet := d.GetOk("connection")
	if connectionIsSet {
		connectionMap := connectionInterface.([]interface{})[0].(map[string]interface{})
		connection = LdapServerPostConnectionModel(connectionMap)
	}
	var mappingsForUsers *models.LdapServerPostMappingsForUsers = nil //hit complex
	mappings_for_usersInterface, mappings_for_usersIsSet := d.GetOk("mappings_for_users")
	if mappings_for_usersIsSet {
		mappings_for_usersMap := mappings_for_usersInterface.([]interface{})[0].(map[string]interface{})
		mappingsForUsers = LdapServerPostMappingsForUsersModel(mappings_for_usersMap)
	}

	return &models.LdapServerPost{
		Connection:       connection,
		MappingsForUsers: mappingsForUsers,
	}
}

// Retrieve property field names for updating the LdapServerPost resource
func GetLdapServerPostPropertyFields() (t []string) {
	return []string{
		"connection",
		"mappings_for_users",
	}
}
