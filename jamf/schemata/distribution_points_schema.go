package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DistributionPoints resource defined in the Terraform configuration
func DistributionPointsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying DistributionPoints resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDistributionPointsResourceData(d *schema.ResourceData, m *models.DistributionPoints) {
}

// Iterate throught and update the DistributionPoints resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDistributionPointsSubResourceData(m []*models.DistributionPoints) (d []*map[string]interface{}) {
	for _, distributionPoints := range m {
		if distributionPoints != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DistributionPoints resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DistributionPointsModel(d *schema.ResourceData) *models.DistributionPoints {

	return &models.DistributionPoints{}
}

// Retrieve property field names for updating the DistributionPoints resource
func GetDistributionPointsPropertyFields() (t []string) {
	return []string{}
}
