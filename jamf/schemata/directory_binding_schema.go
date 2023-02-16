package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DirectoryBinding resource defined in the Terraform configuration
func DirectoryBindingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computer_ou": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"domain": {
			Type:     schema.TypeString,
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

		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying DirectoryBinding resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDirectoryBindingResourceData(d *schema.ResourceData, m *models.DirectoryBinding) {
	d.Set("computer_ou", m.ComputerOu)
	d.Set("domain", m.Domain)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("password", m.Password)
	d.Set("priority", m.Priority)
	d.Set("type", m.Type)
	d.Set("username", m.Username)
}

// Iterate throught and update the DirectoryBinding resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDirectoryBindingSubResourceData(m []*models.DirectoryBinding) (d []*map[string]interface{}) {
	for _, directoryBinding := range m {
		if directoryBinding != nil {
			properties := make(map[string]interface{})
			properties["computer_ou"] = directoryBinding.ComputerOu
			properties["domain"] = directoryBinding.Domain
			properties["id"] = directoryBinding.ID
			properties["name"] = directoryBinding.Name
			properties["password"] = directoryBinding.Password
			properties["priority"] = directoryBinding.Priority
			properties["type"] = directoryBinding.Type
			properties["username"] = directoryBinding.Username
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DirectoryBinding resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DirectoryBindingModel(d *schema.ResourceData) *models.DirectoryBinding {
	computerOu := d.Get("computer_ou").(string)
	domain := d.Get("domain").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	password := d.Get("password").(string)
	priority := int64(d.Get("priority").(int))
	typeVar := d.Get("type").(string)
	username := d.Get("username").(string)

	return &models.DirectoryBinding{
		ComputerOu: computerOu,
		Domain:     domain,
		ID:         int32(id),
		Name:       &name,
		Password:   password,
		Priority:   priority,
		Type:       typeVar,
		Username:   username,
	}
}

// Retrieve property field names for updating the DirectoryBinding resource
func GetDirectoryBindingPropertyFields() (t []string) {
	return []string{
		"computer_ou",
		"domain",
		"id",
		"name",
		"password",
		"priority",
		"type",
		"username",
	}
}
