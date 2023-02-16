package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the LicensedSoftwareAll resource defined in the Terraform configuration
func LicensedSoftwareAllSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying LicensedSoftwareAll resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLicensedSoftwareAllResourceData(d *schema.ResourceData, m *models.LicensedSoftwareAll) {
}

// Iterate throught and update the LicensedSoftwareAll resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLicensedSoftwareAllSubResourceData(m []*models.LicensedSoftwareAll) (d []*map[string]interface{}) {
	for _, licensedSoftwareAll := range m {
		if licensedSoftwareAll != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate LicensedSoftwareAll resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LicensedSoftwareAllModel(d *schema.ResourceData) *models.LicensedSoftwareAll {

	return &models.LicensedSoftwareAll{}
}

// Retrieve property field names for updating the LicensedSoftwareAll resource
func GetLicensedSoftwareAllPropertyFields() (t []string) {
	return []string{}
}
