package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceHistoryApp resource defined in the Terraform configuration
func MobileDeviceHistoryAppSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bundle_size": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"dynamic_size": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"management_status": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"version": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceHistoryApp resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceHistoryAppResourceData(d *schema.ResourceData, m *models.MobileDeviceHistoryApp) {
	d.Set("bundle_size", m.BundleSize)
	d.Set("dynamic_size", m.DynamicSize)
	d.Set("management_status", m.ManagementStatus)
	d.Set("name", m.Name)
	d.Set("version", m.Version)
}

// Iterate throught and update the MobileDeviceHistoryApp resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceHistoryAppSubResourceData(m []*models.MobileDeviceHistoryApp) (d []*map[string]interface{}) {
	for _, mobileDeviceHistoryApp := range m {
		if mobileDeviceHistoryApp != nil {
			properties := make(map[string]interface{})
			properties["bundle_size"] = mobileDeviceHistoryApp.BundleSize
			properties["dynamic_size"] = mobileDeviceHistoryApp.DynamicSize
			properties["management_status"] = mobileDeviceHistoryApp.ManagementStatus
			properties["name"] = mobileDeviceHistoryApp.Name
			properties["version"] = mobileDeviceHistoryApp.Version
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceHistoryApp resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceHistoryAppModel(d *schema.ResourceData) *models.MobileDeviceHistoryApp {
	bundleSize := d.Get("bundle_size").(string)
	dynamicSize := d.Get("dynamic_size").(string)
	managementStatus := d.Get("management_status").(string)
	name := d.Get("name").(string)
	version := d.Get("version").(string)

	return &models.MobileDeviceHistoryApp{
		BundleSize:       bundleSize,
		DynamicSize:      dynamicSize,
		ManagementStatus: managementStatus,
		Name:             name,
		Version:          version,
	}
}

// Retrieve property field names for updating the MobileDeviceHistoryApp resource
func GetMobileDeviceHistoryAppPropertyFields() (t []string) {
	return []string{
		"bundle_size",
		"dynamic_size",
		"management_status",
		"name",
		"version",
	}
}
