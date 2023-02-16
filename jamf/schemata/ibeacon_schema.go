package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Ibeacon resource defined in the Terraform configuration
func IbeaconSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"major": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"minor": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying Ibeacon resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIbeaconResourceData(d *schema.ResourceData, m *models.Ibeacon) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("major", m.Major)
	d.Set("minor", m.Minor)
	d.Set("name", m.Name)
	d.Set("uuid", m.UUID)
}

// Iterate throught and update the Ibeacon resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIbeaconSubResourceData(m []*models.Ibeacon) (d []*map[string]interface{}) {
	for _, ibeacon := range m {
		if ibeacon != nil {
			properties := make(map[string]interface{})
			properties["id"] = ibeacon.ID
			properties["major"] = ibeacon.Major
			properties["minor"] = ibeacon.Minor
			properties["name"] = ibeacon.Name
			properties["uuid"] = ibeacon.UUID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Ibeacon resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IbeaconModel(d *schema.ResourceData) *models.Ibeacon {
	id, _ := strconv.Atoi(d.Get("id").(string))
	major := d.Get("major").(string)
	minor := d.Get("minor").(string)
	name := d.Get("name").(string)
	uuid := d.Get("uuid").(string)

	return &models.Ibeacon{
		ID:    int32(id),
		Major: major,
		Minor: minor,
		Name:  &name,
		UUID:  &uuid,
	}
}

// Retrieve property field names for updating the Ibeacon resource
func GetIbeaconPropertyFields() (t []string) {
	return []string{
		"id",
		"major",
		"minor",
		"name",
		"uuid",
	}
}
