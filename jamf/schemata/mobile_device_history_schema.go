package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceHistory resource defined in the Terraform configuration
func MobileDeviceHistorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applications": {
			Type: schema.TypeList, //GoType: MobileDeviceHistoryApplications
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryApplicationsSchema(),
			},
			Optional: true,
		},

		"audits": {
			Type: schema.TypeList, //GoType: []*MobileDeviceHistoryAuditsItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryAuditsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"ebooks": {
			Type: schema.TypeList, //GoType: MobileDeviceHistoryEbooks
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryEbooksSchema(),
			},
			Optional: true,
		},

		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceHistoryGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryGeneralSchema(),
			},
			Optional: true,
		},

		"management_commands": {
			Type: schema.TypeList, //GoType: MobileDeviceHistoryManagementCommands
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryManagementCommandsSchema(),
			},
			Optional: true,
		},

		"user_location": {
			Type: schema.TypeList, //GoType: []*MobileDeviceHistoryUserLocationItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceHistoryUserLocationItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying MobileDeviceHistory resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceHistoryResourceData(d *schema.ResourceData, m *models.MobileDeviceHistory) {
	d.Set("applications", SetMobileDeviceHistoryApplicationsSubResourceData([]*models.MobileDeviceHistoryApplications{m.Applications}))
	d.Set("audits", SetMobileDeviceHistoryAuditsItems0SubResourceData(m.Audits))
	d.Set("ebooks", SetMobileDeviceHistoryEbooksSubResourceData([]*models.MobileDeviceHistoryEbooks{m.Ebooks}))
	d.Set("general", SetMobileDeviceHistoryGeneralSubResourceData([]*models.MobileDeviceHistoryGeneral{m.General}))
	d.Set("management_commands", SetMobileDeviceHistoryManagementCommandsSubResourceData([]*models.MobileDeviceHistoryManagementCommands{m.ManagementCommands}))
	d.Set("user_location", SetMobileDeviceHistoryUserLocationItems0SubResourceData(m.UserLocation))
}

// Iterate throught and update the MobileDeviceHistory resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceHistorySubResourceData(m []*models.MobileDeviceHistory) (d []*map[string]interface{}) {
	for _, mobileDeviceHistory := range m {
		if mobileDeviceHistory != nil {
			properties := make(map[string]interface{})
			properties["applications"] = SetMobileDeviceHistoryApplicationsSubResourceData([]*models.MobileDeviceHistoryApplications{mobileDeviceHistory.Applications})
			properties["audits"] = SetMobileDeviceHistoryAuditsItems0SubResourceData(mobileDeviceHistory.Audits)
			properties["ebooks"] = SetMobileDeviceHistoryEbooksSubResourceData([]*models.MobileDeviceHistoryEbooks{mobileDeviceHistory.Ebooks})
			properties["general"] = SetMobileDeviceHistoryGeneralSubResourceData([]*models.MobileDeviceHistoryGeneral{mobileDeviceHistory.General})
			properties["management_commands"] = SetMobileDeviceHistoryManagementCommandsSubResourceData([]*models.MobileDeviceHistoryManagementCommands{mobileDeviceHistory.ManagementCommands})
			properties["user_location"] = SetMobileDeviceHistoryUserLocationItems0SubResourceData(mobileDeviceHistory.UserLocation)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceHistory resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceHistoryModel(d *schema.ResourceData) *models.MobileDeviceHistory {
	var applications *models.MobileDeviceHistoryApplications = nil //hit complex
	applicationsInterface, applicationsIsSet := d.GetOk("applications")
	if applicationsIsSet {
		applicationsMap := applicationsInterface.([]interface{})[0].(map[string]interface{})
		applications = MobileDeviceHistoryApplicationsModel(applicationsMap)
	}
	audits := d.Get("audits").([]*models.MobileDeviceHistoryAuditsItems0)
	var ebooks *models.MobileDeviceHistoryEbooks = nil //hit complex
	ebooksInterface, ebooksIsSet := d.GetOk("ebooks")
	if ebooksIsSet {
		ebooksMap := ebooksInterface.([]interface{})[0].(map[string]interface{})
		ebooks = MobileDeviceHistoryEbooksModel(ebooksMap)
	}
	var general *models.MobileDeviceHistoryGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceHistoryGeneralModel(generalMap)
	}
	var managementCommands *models.MobileDeviceHistoryManagementCommands = nil //hit complex
	management_commandsInterface, management_commandsIsSet := d.GetOk("management_commands")
	if management_commandsIsSet {
		management_commandsMap := management_commandsInterface.([]interface{})[0].(map[string]interface{})
		managementCommands = MobileDeviceHistoryManagementCommandsModel(management_commandsMap)
	}
	userLocation := d.Get("user_location").([]*models.MobileDeviceHistoryUserLocationItems0)

	return &models.MobileDeviceHistory{
		Applications:       applications,
		Audits:             audits,
		Ebooks:             ebooks,
		General:            general,
		ManagementCommands: managementCommands,
		UserLocation:       userLocation,
	}
}

// Retrieve property field names for updating the MobileDeviceHistory resource
func GetMobileDeviceHistoryPropertyFields() (t []string) {
	return []string{
		"applications",
		"audits",
		"ebooks",
		"general",
		"management_commands",
		"user_location",
	}
}
