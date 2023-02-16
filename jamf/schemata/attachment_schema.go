package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Attachment resource defined in the Terraform configuration
func AttachmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"filename": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"uri": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying Attachment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAttachmentResourceData(d *schema.ResourceData, m *models.Attachment) {
	d.Set("filename", m.Filename)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("uri", m.URI)
}

// Iterate throught and update the Attachment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAttachmentSubResourceData(m []*models.Attachment) (d []*map[string]interface{}) {
	for _, attachment := range m {
		if attachment != nil {
			properties := make(map[string]interface{})
			properties["filename"] = attachment.Filename
			properties["id"] = attachment.ID
			properties["uri"] = attachment.URI
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Attachment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AttachmentModel(d *schema.ResourceData) *models.Attachment {
	filename := d.Get("filename").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	uri := d.Get("uri").(string)

	return &models.Attachment{
		Filename: filename,
		ID:       int32(id),
		URI:      uri,
	}
}

// Retrieve property field names for updating the Attachment resource
func GetAttachmentPropertyFields() (t []string) {
	return []string{
		"filename",
		"id",
		"uri",
	}
}
