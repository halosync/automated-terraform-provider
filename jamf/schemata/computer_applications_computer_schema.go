package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerApplicationsComputer resource defined in the Terraform configuration
func ComputerApplicationsComputerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"mac_address": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"serial_number": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"udid": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying ComputerApplicationsComputer resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerApplicationsComputerResourceData(d *schema.ResourceData, m *models.ComputerApplicationsComputer) {
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("mac_address", m.MacAddress)
	d.Set("name", m.Name)
	d.Set("serial_number", m.SerialNumber)
	d.Set("udid", m.Udid)
}

// Iterate throught and update the ComputerApplicationsComputer resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerApplicationsComputerSubResourceData(m []*models.ComputerApplicationsComputer) (d []*map[string]interface{}) {
	for _, computerApplicationsComputer := range m {
		if computerApplicationsComputer != nil {
			properties := make(map[string]interface{})
			properties["id"] = computerApplicationsComputer.ID
			properties["mac_address"] = computerApplicationsComputer.MacAddress
			properties["name"] = computerApplicationsComputer.Name
			properties["serial_number"] = computerApplicationsComputer.SerialNumber
			properties["udid"] = computerApplicationsComputer.Udid
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerApplicationsComputer resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerApplicationsComputerModel(d *schema.ResourceData) *models.ComputerApplicationsComputer {
	id, _ := strconv.Atoi(d.Get("id").(string))
	macAddress := d.Get("mac_address").(string)
	name := d.Get("name").(string)
	serialNumber := d.Get("serial_number").(string)
	udid := d.Get("udid").(string)

	return &models.ComputerApplicationsComputer{
		ID:           int32(id),
		MacAddress:   macAddress,
		Name:         name,
		SerialNumber: serialNumber,
		Udid:         udid,
	}
}

// Retrieve property field names for updating the ComputerApplicationsComputer resource
func GetComputerApplicationsComputerPropertyFields() (t []string) {
	return []string{
		"id",
		"mac_address",
		"name",
		"serial_number",
		"udid",
	}
}
