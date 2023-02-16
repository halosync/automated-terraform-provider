package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerHistory resource defined in the Terraform configuration
func ComputerHistorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"audits": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryAuditsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryAuditsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"casper_imaging_logs": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryCasperImagingLogsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryCasperImagingLogsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"casper_remote_logs": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryCasperRemoteLogsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryCasperRemoteLogsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"commands": {
			Type: schema.TypeList, //GoType: ComputerHistoryCommands
			Elem: &schema.Resource{
				Schema: ComputerHistoryCommandsSchema(),
			},
			Optional: true,
		},

		"computer_usage_logs": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryComputerUsageLogsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryComputerUsageLogsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: ComputerHistoryGeneral
			Elem: &schema.Resource{
				Schema: ComputerHistoryGeneralSchema(),
			},
			Optional: true,
		},

		"mac_app_store_applications": {
			Type: schema.TypeList, //GoType: ComputerHistoryMacAppStoreApplications
			Elem: &schema.Resource{
				Schema: ComputerHistoryMacAppStoreApplicationsSchema(),
			},
			Optional: true,
		},

		"policy_logs": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryPolicyLogsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryPolicyLogsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"screen_sharing_logs": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryScreenSharingLogsItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryScreenSharingLogsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"user_location": {
			Type: schema.TypeList, //GoType: []*ComputerHistoryUserLocationItems0
			Elem: &schema.Resource{
				Schema: ComputerHistoryUserLocationItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying ComputerHistory resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerHistoryResourceData(d *schema.ResourceData, m *models.ComputerHistory) {
	d.Set("audits", SetComputerHistoryAuditsItems0SubResourceData(m.Audits))
	d.Set("casper_imaging_logs", SetComputerHistoryCasperImagingLogsItems0SubResourceData(m.CasperImagingLogs))
	d.Set("casper_remote_logs", SetComputerHistoryCasperRemoteLogsItems0SubResourceData(m.CasperRemoteLogs))
	d.Set("commands", SetComputerHistoryCommandsSubResourceData([]*models.ComputerHistoryCommands{m.Commands}))
	d.Set("computer_usage_logs", SetComputerHistoryComputerUsageLogsItems0SubResourceData(m.ComputerUsageLogs))
	d.Set("general", SetComputerHistoryGeneralSubResourceData([]*models.ComputerHistoryGeneral{m.General}))
	d.Set("mac_app_store_applications", SetComputerHistoryMacAppStoreApplicationsSubResourceData([]*models.ComputerHistoryMacAppStoreApplications{m.MacAppStoreApplications}))
	d.Set("policy_logs", SetComputerHistoryPolicyLogsItems0SubResourceData(m.PolicyLogs))
	d.Set("screen_sharing_logs", SetComputerHistoryScreenSharingLogsItems0SubResourceData(m.ScreenSharingLogs))
	d.Set("user_location", SetComputerHistoryUserLocationItems0SubResourceData(m.UserLocation))
}

// Iterate throught and update the ComputerHistory resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerHistorySubResourceData(m []*models.ComputerHistory) (d []*map[string]interface{}) {
	for _, computerHistory := range m {
		if computerHistory != nil {
			properties := make(map[string]interface{})
			properties["audits"] = SetComputerHistoryAuditsItems0SubResourceData(computerHistory.Audits)
			properties["casper_imaging_logs"] = SetComputerHistoryCasperImagingLogsItems0SubResourceData(computerHistory.CasperImagingLogs)
			properties["casper_remote_logs"] = SetComputerHistoryCasperRemoteLogsItems0SubResourceData(computerHistory.CasperRemoteLogs)
			properties["commands"] = SetComputerHistoryCommandsSubResourceData([]*models.ComputerHistoryCommands{computerHistory.Commands})
			properties["computer_usage_logs"] = SetComputerHistoryComputerUsageLogsItems0SubResourceData(computerHistory.ComputerUsageLogs)
			properties["general"] = SetComputerHistoryGeneralSubResourceData([]*models.ComputerHistoryGeneral{computerHistory.General})
			properties["mac_app_store_applications"] = SetComputerHistoryMacAppStoreApplicationsSubResourceData([]*models.ComputerHistoryMacAppStoreApplications{computerHistory.MacAppStoreApplications})
			properties["policy_logs"] = SetComputerHistoryPolicyLogsItems0SubResourceData(computerHistory.PolicyLogs)
			properties["screen_sharing_logs"] = SetComputerHistoryScreenSharingLogsItems0SubResourceData(computerHistory.ScreenSharingLogs)
			properties["user_location"] = SetComputerHistoryUserLocationItems0SubResourceData(computerHistory.UserLocation)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerHistory resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerHistoryModel(d *schema.ResourceData) *models.ComputerHistory {
	audits := d.Get("audits").([]*models.ComputerHistoryAuditsItems0)
	casperImagingLogs := d.Get("casper_imaging_logs").([]*models.ComputerHistoryCasperImagingLogsItems0)
	casperRemoteLogs := d.Get("casper_remote_logs").([]*models.ComputerHistoryCasperRemoteLogsItems0)
	var commands *models.ComputerHistoryCommands = nil //hit complex
	commandsInterface, commandsIsSet := d.GetOk("commands")
	if commandsIsSet {
		commandsMap := commandsInterface.([]interface{})[0].(map[string]interface{})
		commands = ComputerHistoryCommandsModel(commandsMap)
	}
	computerUsageLogs := d.Get("computer_usage_logs").([]*models.ComputerHistoryComputerUsageLogsItems0)
	var general *models.ComputerHistoryGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerHistoryGeneralModel(generalMap)
	}
	var macAppStoreApplications *models.ComputerHistoryMacAppStoreApplications = nil //hit complex
	mac_app_store_applicationsInterface, mac_app_store_applicationsIsSet := d.GetOk("mac_app_store_applications")
	if mac_app_store_applicationsIsSet {
		mac_app_store_applicationsMap := mac_app_store_applicationsInterface.([]interface{})[0].(map[string]interface{})
		macAppStoreApplications = ComputerHistoryMacAppStoreApplicationsModel(mac_app_store_applicationsMap)
	}
	policyLogs := d.Get("policy_logs").([]*models.ComputerHistoryPolicyLogsItems0)
	screenSharingLogs := d.Get("screen_sharing_logs").([]*models.ComputerHistoryScreenSharingLogsItems0)
	userLocation := d.Get("user_location").([]*models.ComputerHistoryUserLocationItems0)

	return &models.ComputerHistory{
		Audits:                  audits,
		CasperImagingLogs:       casperImagingLogs,
		CasperRemoteLogs:        casperRemoteLogs,
		Commands:                commands,
		ComputerUsageLogs:       computerUsageLogs,
		General:                 general,
		MacAppStoreApplications: macAppStoreApplications,
		PolicyLogs:              policyLogs,
		ScreenSharingLogs:       screenSharingLogs,
		UserLocation:            userLocation,
	}
}

// Retrieve property field names for updating the ComputerHistory resource
func GetComputerHistoryPropertyFields() (t []string) {
	return []string{
		"audits",
		"casper_imaging_logs",
		"casper_remote_logs",
		"commands",
		"computer_usage_logs",
		"general",
		"mac_app_store_applications",
		"policy_logs",
		"screen_sharing_logs",
		"user_location",
	}
}
