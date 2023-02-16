package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NetworkSegments resource defined in the Terraform configuration
func NetworkSegmentsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying NetworkSegments resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkSegmentsResourceData(d *schema.ResourceData, m *models.NetworkSegments) {
}

// Iterate throught and update the NetworkSegments resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkSegmentsSubResourceData(m []*models.NetworkSegments) (d []*map[string]interface{}) {
	for _, networkSegments := range m {
		if networkSegments != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NetworkSegments resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkSegmentsModel(d *schema.ResourceData) *models.NetworkSegments {

	return &models.NetworkSegments{}
}

// Retrieve property field names for updating the NetworkSegments resource
func GetNetworkSegmentsPropertyFields() (t []string) {
	return []string{}
}
