package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DiskEncryptionConfiguration resource defined in the Terraform configuration
func DiskEncryptionConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file_vault_enabled_users": {
			Type:     schema.TypeString,
			Default:  "Management Account",
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"key_type": {
			Type:     schema.TypeString,
			Default:  "Individual",
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying DiskEncryptionConfiguration resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDiskEncryptionConfigurationResourceData(d *schema.ResourceData, m *models.DiskEncryptionConfiguration) {
	d.Set("file_vault_enabled_users", m.FileVaultEnabledUsers)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("key_type", m.KeyType)
	d.Set("name", m.Name)
}

// Iterate throught and update the DiskEncryptionConfiguration resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDiskEncryptionConfigurationSubResourceData(m []*models.DiskEncryptionConfiguration) (d []*map[string]interface{}) {
	for _, diskEncryptionConfiguration := range m {
		if diskEncryptionConfiguration != nil {
			properties := make(map[string]interface{})
			properties["file_vault_enabled_users"] = diskEncryptionConfiguration.FileVaultEnabledUsers
			properties["id"] = diskEncryptionConfiguration.ID
			properties["key_type"] = diskEncryptionConfiguration.KeyType
			properties["name"] = diskEncryptionConfiguration.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DiskEncryptionConfiguration resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DiskEncryptionConfigurationModel(d *schema.ResourceData) *models.DiskEncryptionConfiguration {
	fileVaultEnabledUsers := d.Get("file_vault_enabled_users").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	keyType := d.Get("key_type").(string)
	name := d.Get("name").(string)

	return &models.DiskEncryptionConfiguration{
		FileVaultEnabledUsers: &fileVaultEnabledUsers,
		ID:                    int32(id),
		KeyType:               &keyType,
		Name:                  &name,
	}
}

// Retrieve property field names for updating the DiskEncryptionConfiguration resource
func GetDiskEncryptionConfigurationPropertyFields() (t []string) {
	return []string{
		"file_vault_enabled_users",
		"id",
		"key_type",
		"name",
	}
}
