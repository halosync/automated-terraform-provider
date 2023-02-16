package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the HealthcareListenerRules resource defined in the Terraform configuration
func HealthcareListenerRulesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying HealthcareListenerRules resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetHealthcareListenerRulesResourceData(d *schema.ResourceData, m *models.HealthcareListenerRules) {
}

// Iterate throught and update the HealthcareListenerRules resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetHealthcareListenerRulesSubResourceData(m []*models.HealthcareListenerRules) (d []*map[string]interface{}) {
	for _, healthcareListenerRules := range m {
		if healthcareListenerRules != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate HealthcareListenerRules resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func HealthcareListenerRulesModel(d *schema.ResourceData) *models.HealthcareListenerRules {

	return &models.HealthcareListenerRules{}
}

// Retrieve property field names for updating the HealthcareListenerRules resource
func GetHealthcareListenerRulesPropertyFields() (t []string) {
	return []string{}
}
