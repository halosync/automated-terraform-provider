package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the HealthcareListener resource defined in the Terraform configuration
func HealthcareListenerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying HealthcareListener resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetHealthcareListenerResourceData(d *schema.ResourceData, m *models.HealthcareListener) {
}

// Iterate throught and update the HealthcareListener resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetHealthcareListenerSubResourceData(m []*models.HealthcareListener) (d []*map[string]interface{}) {
	for _, healthcareListener := range m {
		if healthcareListener != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate HealthcareListener resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func HealthcareListenerModel(d *schema.ResourceData) *models.HealthcareListener {

	return &models.HealthcareListener{}
}

// Retrieve property field names for updating the HealthcareListener resource
func GetHealthcareListenerPropertyFields() (t []string) {
	return []string{}
}
