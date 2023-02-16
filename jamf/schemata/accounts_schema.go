package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Accounts resource defined in the Terraform configuration
func AccountsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"groups": {
			Type: schema.TypeList, //GoType: AccountsGroups
			Elem: &schema.Resource{
				Schema: AccountsGroupsSchema(),
			},
			Optional: true,
		},

		"users": {
			Type: schema.TypeList, //GoType: AccountsUsers
			Elem: &schema.Resource{
				Schema: AccountsUsersSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Accounts resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAccountsResourceData(d *schema.ResourceData, m *models.Accounts) {
	d.Set("groups", SetAccountsGroupsSubResourceData([]*models.AccountsGroups{m.Groups}))
	d.Set("users", SetAccountsUsersSubResourceData([]*models.AccountsUsers{m.Users}))
}

// Iterate throught and update the Accounts resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAccountsSubResourceData(m []*models.Accounts) (d []*map[string]interface{}) {
	for _, accounts := range m {
		if accounts != nil {
			properties := make(map[string]interface{})
			properties["groups"] = SetAccountsGroupsSubResourceData([]*models.AccountsGroups{accounts.Groups})
			properties["users"] = SetAccountsUsersSubResourceData([]*models.AccountsUsers{accounts.Users})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Accounts resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AccountsModel(d *schema.ResourceData) *models.Accounts {
	var groups *models.AccountsGroups = nil //hit complex
	groupsInterface, groupsIsSet := d.GetOk("groups")
	if groupsIsSet {
		groupsMap := groupsInterface.([]interface{})[0].(map[string]interface{})
		groups = AccountsGroupsModel(groupsMap)
	}
	var users *models.AccountsUsers = nil //hit complex
	usersInterface, usersIsSet := d.GetOk("users")
	if usersIsSet {
		usersMap := usersInterface.([]interface{})[0].(map[string]interface{})
		users = AccountsUsersModel(usersMap)
	}

	return &models.Accounts{
		Groups: groups,
		Users:  users,
	}
}

// Retrieve property field names for updating the Accounts resource
func GetAccountsPropertyFields() (t []string) {
	return []string{
		"groups",
		"users",
	}
}
