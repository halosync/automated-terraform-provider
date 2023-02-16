package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the NetworkSegment resource defined in the Terraform configuration
func NetworkSegmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"building": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"department": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"distribution_point": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"distribution_server": {
			Type:     schema.TypeString,
			Optional: true,
		},

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

		"swu_server": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"url": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying NetworkSegment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkSegmentResourceData(d *schema.ResourceData, m *models.NetworkSegment) {
	d.Set("building", m.Building)
	d.Set("department", m.Department)
	d.Set("distribution_point", m.DistributionPoint)
	d.Set("distribution_server", m.DistributionServer)
	d.Set("ending_address", m.EndingAddress)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("override_buildings", m.OverrideBuildings)
	d.Set("override_departments", m.OverrideDepartments)
	d.Set("starting_address", m.StartingAddress)
	d.Set("swu_server", m.SwuServer)
	d.Set("url", m.URL)
}

// Iterate throught and update the NetworkSegment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkSegmentSubResourceData(m []*models.NetworkSegment) (d []*map[string]interface{}) {
	for _, networkSegment := range m {
		if networkSegment != nil {
			properties := make(map[string]interface{})
			properties["building"] = networkSegment.Building
			properties["department"] = networkSegment.Department
			properties["distribution_point"] = networkSegment.DistributionPoint
			properties["distribution_server"] = networkSegment.DistributionServer
			properties["ending_address"] = networkSegment.EndingAddress
			properties["id"] = networkSegment.ID
			properties["name"] = networkSegment.Name
			properties["override_buildings"] = networkSegment.OverrideBuildings
			properties["override_departments"] = networkSegment.OverrideDepartments
			properties["starting_address"] = networkSegment.StartingAddress
			properties["swu_server"] = networkSegment.SwuServer
			properties["url"] = networkSegment.URL
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NetworkSegment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkSegmentModel(d *schema.ResourceData) *models.NetworkSegment {
	building := d.Get("building").(string)
	department := d.Get("department").(string)
	distributionPoint := d.Get("distribution_point").(string)
	distributionServer := d.Get("distribution_server").(string)
	endingAddress := d.Get("ending_address").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	name := d.Get("name").(string)
	overrideBuildings := d.Get("override_buildings").(bool)
	overrideDepartments := d.Get("override_departments").(bool)
	startingAddress := d.Get("starting_address").(string)
	swuServer := d.Get("swu_server").(string)
	url := d.Get("url").(string)

	return &models.NetworkSegment{
		Building:            building,
		Department:          department,
		DistributionPoint:   distributionPoint,
		DistributionServer:  distributionServer,
		EndingAddress:       &endingAddress,
		ID:                  int32(id),
		Name:                &name,
		OverrideBuildings:   &overrideBuildings,
		OverrideDepartments: &overrideDepartments,
		StartingAddress:     &startingAddress,
		SwuServer:           swuServer,
		URL:                 url,
	}
}

// Retrieve property field names for updating the NetworkSegment resource
func GetNetworkSegmentPropertyFields() (t []string) {
	return []string{
		"building",
		"department",
		"distribution_point",
		"distribution_server",
		"ending_address",
		"id",
		"name",
		"override_buildings",
		"override_departments",
		"starting_address",
		"swu_server",
		"url",
	}
}
