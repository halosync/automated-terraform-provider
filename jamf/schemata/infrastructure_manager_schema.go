package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the InfrastructureManager resource defined in the Terraform configuration
func InfrastructureManagerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"last_check_in": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"last_reported_ip": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"operating_system": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"operating_system_version": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"recurring_check_in_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying InfrastructureManager resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetInfrastructureManagerResourceData(d *schema.ResourceData, m *models.InfrastructureManager) {
	d.Set("hostname", m.Hostname)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("last_check_in", m.LastCheckIn)
	d.Set("last_reported_ip", m.LastReportedIP)
	d.Set("name", m.Name)
	d.Set("operating_system", m.OperatingSystem)
	d.Set("operating_system_version", m.OperatingSystemVersion)
	d.Set("recurring_check_in_frequency", m.RecurringCheckInFrequency)
	d.Set("uuid", m.UUID)
}

// Iterate throught and update the InfrastructureManager resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetInfrastructureManagerSubResourceData(m []*models.InfrastructureManager) (d []*map[string]interface{}) {
	for _, infrastructureManager := range m {
		if infrastructureManager != nil {
			properties := make(map[string]interface{})
			properties["hostname"] = infrastructureManager.Hostname
			properties["id"] = infrastructureManager.ID
			properties["last_check_in"] = infrastructureManager.LastCheckIn
			properties["last_reported_ip"] = infrastructureManager.LastReportedIP
			properties["name"] = infrastructureManager.Name
			properties["operating_system"] = infrastructureManager.OperatingSystem
			properties["operating_system_version"] = infrastructureManager.OperatingSystemVersion
			properties["recurring_check_in_frequency"] = infrastructureManager.RecurringCheckInFrequency
			properties["uuid"] = infrastructureManager.UUID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate InfrastructureManager resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func InfrastructureManagerModel(d *schema.ResourceData) *models.InfrastructureManager {
	hostname := d.Get("hostname").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	lastCheckIn := d.Get("last_check_in").(string)
	lastReportedIP := d.Get("last_reported_ip").(string)
	name := d.Get("name").(string)
	operatingSystem := d.Get("operating_system").(string)
	operatingSystemVersion := d.Get("operating_system_version").(string)
	recurringCheckInFrequency := int64(d.Get("recurring_check_in_frequency").(int))
	uuid := d.Get("uuid").(string)

	return &models.InfrastructureManager{
		Hostname:                  hostname,
		ID:                        int32(id),
		LastCheckIn:               lastCheckIn,
		LastReportedIP:            lastReportedIP,
		Name:                      name,
		OperatingSystem:           operatingSystem,
		OperatingSystemVersion:    operatingSystemVersion,
		RecurringCheckInFrequency: recurringCheckInFrequency,
		UUID:                      uuid,
	}
}

// Retrieve property field names for updating the InfrastructureManager resource
func GetInfrastructureManagerPropertyFields() (t []string) {
	return []string{
		"hostname",
		"id",
		"last_check_in",
		"last_reported_ip",
		"name",
		"operating_system",
		"operating_system_version",
		"recurring_check_in_frequency",
		"uuid",
	}
}
