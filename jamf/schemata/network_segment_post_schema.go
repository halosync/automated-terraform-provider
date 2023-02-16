package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NetworkSegmentPost resource defined in the Terraform configuration
func NetworkSegmentPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ending_address": {
			Type:     schema.TypeString,
			Required: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"override_buildings": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"override_departments": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"starting_address": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying NetworkSegmentPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkSegmentPostResourceData(d *schema.ResourceData, m *models.NetworkSegmentPost) {
	d.Set("ending_address", m.EndingAddress)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("override_buildings", m.OverrideBuildings)
	d.Set("override_departments", m.OverrideDepartments)
	d.Set("starting_address", m.StartingAddress)
}

// Iterate throught and update the NetworkSegmentPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkSegmentPostSubResourceData(m []*models.NetworkSegmentPost) (d []*map[string]interface{}) {
	for _, networkSegmentPost := range m {
		if networkSegmentPost != nil {
			properties := make(map[string]interface{})
			properties["ending_address"] = networkSegmentPost.EndingAddress
			properties["id"] = networkSegmentPost.ID
			properties["name"] = networkSegmentPost.Name
			properties["override_buildings"] = networkSegmentPost.OverrideBuildings
			properties["override_departments"] = networkSegmentPost.OverrideDepartments
			properties["starting_address"] = networkSegmentPost.StartingAddress
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NetworkSegmentPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkSegmentPostModel(d *schema.ResourceData) *models.NetworkSegmentPost {
	endingAddress := d.Get("ending_address").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	overrideBuildings := d.Get("override_buildings").(bool)
	overrideDepartments := d.Get("override_departments").(bool)
	startingAddress := d.Get("starting_address").(string)

	return &models.NetworkSegmentPost{
		EndingAddress:       &endingAddress,
		ID:                  int32(id),
		Name:                &name,
		OverrideBuildings:   &overrideBuildings,
		OverrideDepartments: &overrideDepartments,
		StartingAddress:     &startingAddress,
	}
}

// Retrieve property field names for updating the NetworkSegmentPost resource
func GetNetworkSegmentPostPropertyFields() (t []string) {
	return []string{
		"ending_address",
		"id",
		"name",
		"override_buildings",
		"override_departments",
		"starting_address",
	}
}
