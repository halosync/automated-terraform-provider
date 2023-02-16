package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Location resource defined in the Terraform configuration
func LocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"building": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"department": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"email_address": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"phone": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"phone_number": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"position": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"real_name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"realname": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"room": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying Location resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLocationResourceData(d *schema.ResourceData, m *models.Location) {
	d.Set("building", m.Building)
	d.Set("department", m.Department)
	d.Set("email_address", m.EmailAddress)
	d.Set("phone", m.Phone)
	d.Set("phone_number", m.PhoneNumber)
	d.Set("position", m.Position)
	d.Set("real_name", m.RealName)
	d.Set("realname", m.Realname)
	d.Set("room", m.Room)
	d.Set("username", m.Username)
}

// Iterate throught and update the Location resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLocationSubResourceData(m []*models.Location) (d []*map[string]interface{}) {
	for _, location := range m {
		if location != nil {
			properties := make(map[string]interface{})
			properties["building"] = location.Building
			properties["department"] = location.Department
			properties["email_address"] = location.EmailAddress
			properties["phone"] = location.Phone
			properties["phone_number"] = location.PhoneNumber
			properties["position"] = location.Position
			properties["real_name"] = location.RealName
			properties["realname"] = location.Realname
			properties["room"] = location.Room
			properties["username"] = location.Username
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Location resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LocationModel(d *schema.ResourceData) *models.Location {
	building := d.Get("building").(string)
	department := d.Get("department").(string)
	emailAddress := d.Get("email_address").(string)
	phone := d.Get("phone").(string)
	phoneNumber := d.Get("phone_number").(string)
	position := d.Get("position").(string)
	realName := d.Get("real_name").(string)
	realname := d.Get("realname").(string)
	room := d.Get("room").(string)
	username := d.Get("username").(string)

	return &models.Location{
		Building:     building,
		Department:   department,
		EmailAddress: emailAddress,
		Phone:        phone,
		PhoneNumber:  phoneNumber,
		Position:     position,
		RealName:     realName,
		Realname:     realname,
		Room:         room,
		Username:     username,
	}
}

// Retrieve property field names for updating the Location resource
func GetLocationPropertyFields() (t []string) {
	return []string{
		"building",
		"department",
		"email_address",
		"phone",
		"phone_number",
		"position",
		"real_name",
		"realname",
		"room",
		"username",
	}
}
