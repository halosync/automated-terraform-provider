package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerApplications resource defined in the Terraform configuration
func ComputerApplicationsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"unique_computers": {
			Type: schema.TypeList, //GoType: ComputerApplicationsUniqueComputers
			Elem: &schema.Resource{
				Schema: ComputerApplicationsUniqueComputersSchema(),
			},
			Optional: true,
		},

		"versions": {
			Type: schema.TypeList, //GoType: ComputerApplicationsVersions
			Elem: &schema.Resource{
				Schema: ComputerApplicationsVersionsSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputerApplications resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerApplicationsResourceData(d *schema.ResourceData, m *models.ComputerApplications) {
	d.Set("unique_computers", SetComputerApplicationsUniqueComputersSubResourceData([]*models.ComputerApplicationsUniqueComputers{m.UniqueComputers}))
	d.Set("versions", SetComputerApplicationsVersionsSubResourceData([]*models.ComputerApplicationsVersions{m.Versions}))
}

// Iterate throught and update the ComputerApplications resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerApplicationsSubResourceData(m []*models.ComputerApplications) (d []*map[string]interface{}) {
	for _, computerApplications := range m {
		if computerApplications != nil {
			properties := make(map[string]interface{})
			properties["unique_computers"] = SetComputerApplicationsUniqueComputersSubResourceData([]*models.ComputerApplicationsUniqueComputers{computerApplications.UniqueComputers})
			properties["versions"] = SetComputerApplicationsVersionsSubResourceData([]*models.ComputerApplicationsVersions{computerApplications.Versions})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerApplications resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerApplicationsModel(d *schema.ResourceData) *models.ComputerApplications {
	var uniqueComputers *models.ComputerApplicationsUniqueComputers = nil //hit complex
	unique_computersInterface, unique_computersIsSet := d.GetOk("unique_computers")
	if unique_computersIsSet {
		unique_computersMap := unique_computersInterface.([]interface{})[0].(map[string]interface{})
		uniqueComputers = ComputerApplicationsUniqueComputersModel(unique_computersMap)
	}
	var versions *models.ComputerApplicationsVersions = nil //hit complex
	versionsInterface, versionsIsSet := d.GetOk("versions")
	if versionsIsSet {
		versionsMap := versionsInterface.([]interface{})[0].(map[string]interface{})
		versions = ComputerApplicationsVersionsModel(versionsMap)
	}

	return &models.ComputerApplications{
		UniqueComputers: uniqueComputers,
		Versions:        versions,
	}
}

// Retrieve property field names for updating the ComputerApplications resource
func GetComputerApplicationsPropertyFields() (t []string) {
	return []string{
		"unique_computers",
		"versions",
	}
}
