package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Group resource defined in the Terraform configuration
func GroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_level": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"members": {
			Type: schema.TypeList, //GoType: []*GroupMembersItems0
			Elem: &schema.Resource{
				Schema: GroupMembersItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
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
			Type: schema.TypeList, //GoType: GroupPrivileges
			Elem: &schema.Resource{
				Schema: GroupPrivilegesSchema(),
			},
			Optional: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: Site
			Elem: &schema.Resource{
				Schema: SiteSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Group resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGroupResourceData(d *schema.ResourceData, m *models.Group) {
	d.Set("access_level", m.AccessLevel)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("members", SetGroupMembersItems0SubResourceData(m.Members))
	d.Set("name", m.Name)
	d.Set("privilege_set", m.PrivilegeSet)
	d.Set("privileges", SetGroupPrivilegesSubResourceData([]*models.GroupPrivileges{m.Privileges}))
	d.Set("site", SetSiteSubResourceData([]*models.Site{m.Site}))
}

// Iterate throught and update the Group resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGroupSubResourceData(m []*models.Group) (d []*map[string]interface{}) {
	for _, group := range m {
		if group != nil {
			properties := make(map[string]interface{})
			properties["access_level"] = group.AccessLevel
			properties["id"] = group.ID
			properties["members"] = SetGroupMembersItems0SubResourceData(group.Members)
			properties["name"] = group.Name
			properties["privilege_set"] = group.PrivilegeSet
			properties["privileges"] = SetGroupPrivilegesSubResourceData([]*models.GroupPrivileges{group.Privileges})
			properties["site"] = SetSiteSubResourceData([]*models.Site{group.Site})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Group resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GroupModel(d *schema.ResourceData) *models.Group {
	accessLevel := d.Get("access_level").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	members := d.Get("members").([]*models.GroupMembersItems0)
	name := d.Get("name").(string)
	privilegeSet := d.Get("privilege_set").(string)
	var privileges *models.GroupPrivileges = nil //hit complex
	privilegesInterface, privilegesIsSet := d.GetOk("privileges")
	if privilegesIsSet {
		privilegesMap := privilegesInterface.([]interface{})[0].(map[string]interface{})
		privileges = GroupPrivilegesModel(privilegesMap)
	}
	var site *models.Site = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteModel(siteMap)
	}

	return &models.Group{
		AccessLevel:  accessLevel,
		ID:           int32(id),
		Members:      members,
		Name:         &name,
		PrivilegeSet: privilegeSet,
		Privileges:   privileges,
		Site:         site,
	}
}

// Retrieve property field names for updating the Group resource
func GetGroupPropertyFields() (t []string) {
	return []string{
		"access_level",
		"id",
		"members",
		"name",
		"privilege_set",
		"privileges",
		"site",
	}
}
