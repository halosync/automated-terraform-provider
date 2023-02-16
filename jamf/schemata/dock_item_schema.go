package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DockItem resource defined in the Terraform configuration
func DockItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"contents": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"path": {
			Type:     schema.TypeString,
			Required: true,
		},

		"type": {
			Type:     schema.TypeString,
			Default:  "App",
			Required: true,
		},
	}
}

// Update the underlying DockItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDockItemResourceData(d *schema.ResourceData, m *models.DockItem) {
	d.Set("contents", m.Contents)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("path", m.Path)
	d.Set("type", m.Type)
}

// Iterate throught and update the DockItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDockItemSubResourceData(m []*models.DockItem) (d []*map[string]interface{}) {
	for _, dockItem := range m {
		if dockItem != nil {
			properties := make(map[string]interface{})
			properties["contents"] = dockItem.Contents
			properties["id"] = dockItem.ID
			properties["name"] = dockItem.Name
			properties["path"] = dockItem.Path
			properties["type"] = dockItem.Type
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DockItem resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DockItemModel(d *schema.ResourceData) *models.DockItem {
	contents := d.Get("contents").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	path := d.Get("path").(string)
	typeVar := d.Get("type").(string)

	return &models.DockItem{
		Contents: contents,
		ID:       int32(id),
		Name:     &name,
		Path:     &path,
		Type:     &typeVar,
	}
}

// Retrieve property field names for updating the DockItem resource
func GetDockItemPropertyFields() (t []string) {
	return []string{
		"contents",
		"id",
		"name",
		"path",
		"type",
	}
}
