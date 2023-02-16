package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Account resource defined in the Terraform configuration
func AccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"directory_user": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"email": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"email_address": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enabled": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"force_password_change": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"ldap_server": {
			Type: schema.TypeList, //GoType: AccountLdapServer
			Elem: &schema.Resource{
				Schema: AccountLdapServerSchema(),
			},
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"privilege_set": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"privileges": {
			Type: schema.TypeList, //GoType: AccountPrivileges
			Elem: &schema.Resource{
				Schema: AccountPrivilegesSchema(),
			},
			Optional: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: SiteObject
			Elem: &schema.Resource{
				Schema: SiteObjectSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Account resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAccountResourceData(d *schema.ResourceData, m *models.Account) {
	d.Set("access_level", m.AccessLevel)
	d.Set("directory_user", m.DirectoryUser)
	d.Set("email", m.Email)
	d.Set("email_address", m.EmailAddress)
	d.Set("enabled", m.Enabled)
	d.Set("force_password_change", m.ForcePasswordChange)
	d.Set("full_name", m.FullName)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("ldap_server", SetAccountLdapServerSubResourceData([]*models.AccountLdapServer{m.LdapServer}))
	d.Set("name", m.Name)
	d.Set("privilege_set", m.PrivilegeSet)
	d.Set("privileges", SetAccountPrivilegesSubResourceData([]*models.AccountPrivileges{m.Privileges}))
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
}

// Iterate throught and update the Account resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAccountSubResourceData(m []*models.Account) (d []*map[string]interface{}) {
	for _, account := range m {
		if account != nil {
			properties := make(map[string]interface{})
			properties["access_level"] = account.AccessLevel
			properties["directory_user"] = account.DirectoryUser
			properties["email"] = account.Email
			properties["email_address"] = account.EmailAddress
			properties["enabled"] = account.Enabled
			properties["force_password_change"] = account.ForcePasswordChange
			properties["full_name"] = account.FullName
			properties["id"] = account.ID
			properties["ldap_server"] = SetAccountLdapServerSubResourceData([]*models.AccountLdapServer{account.LdapServer})
			properties["name"] = account.Name
			properties["privilege_set"] = account.PrivilegeSet
			properties["privileges"] = SetAccountPrivilegesSubResourceData([]*models.AccountPrivileges{account.Privileges})
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{account.Site})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Account resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AccountModel(d *schema.ResourceData) *models.Account {
	accessLevel := d.Get("access_level").(string)
	directoryUser := d.Get("directory_user").(bool)
	email := d.Get("email").(string)
	emailAddress := d.Get("email_address").(string)
	enabled := d.Get("enabled").(string)
	forcePasswordChange := d.Get("force_password_change").(bool)
	fullName := d.Get("full_name").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	var ldapServer *models.AccountLdapServer = nil //hit complex
	ldap_serverInterface, ldap_serverIsSet := d.GetOk("ldap_server")
	if ldap_serverIsSet {
		ldap_serverMap := ldap_serverInterface.([]interface{})[0].(map[string]interface{})
		ldapServer = AccountLdapServerModel(ldap_serverMap)
	}
	name := d.Get("name").(string)
	privilegeSet := d.Get("privilege_set").(string)
	var privileges *models.AccountPrivileges = nil //hit complex
	privilegesInterface, privilegesIsSet := d.GetOk("privileges")
	if privilegesIsSet {
		privilegesMap := privilegesInterface.([]interface{})[0].(map[string]interface{})
		privileges = AccountPrivilegesModel(privilegesMap)
	}
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}

	return &models.Account{
		AccessLevel:         accessLevel,
		DirectoryUser:       directoryUser,
		Email:               email,
		EmailAddress:        emailAddress,
		Enabled:             enabled,
		ForcePasswordChange: forcePasswordChange,
		FullName:            fullName,
		ID:                  int32(id),
		LdapServer:          ldapServer,
		Name:                &name,
		PrivilegeSet:        privilegeSet,
		Privileges:          privileges,
		Site:                site,
	}
}

// Retrieve property field names for updating the Account resource
func GetAccountPropertyFields() (t []string) {
	return []string{
		"access_level",
		"directory_user",
		"email",
		"email_address",
		"enabled",
		"force_password_change",
		"full_name",
		"id",
		"ldap_server",
		"name",
		"privilege_set",
		"privileges",
		"site",
	}
}
