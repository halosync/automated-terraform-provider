package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DiskEncryptionConfigurations resource defined in the Terraform configuration
func DiskEncryptionConfigurationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Update the underlying DiskEncryptionConfigurations resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDiskEncryptionConfigurationsResourceData(d *schema.ResourceData, m *models.DiskEncryptionConfigurations) {
}

// Iterate throught and update the DiskEncryptionConfigurations resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDiskEncryptionConfigurationsSubResourceData(m []*models.DiskEncryptionConfigurations) (d []*map[string]interface{}) {
	for _, diskEncryptionConfigurations := range m {
		if diskEncryptionConfigurations != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DiskEncryptionConfigurations resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DiskEncryptionConfigurationsModel(d *schema.ResourceData) *models.DiskEncryptionConfigurations {

	return &models.DiskEncryptionConfigurations{}
}

// Retrieve property field names for updating the DiskEncryptionConfigurations resource
func GetDiskEncryptionConfigurationsPropertyFields() (t []string) {
	return []string{}
}
