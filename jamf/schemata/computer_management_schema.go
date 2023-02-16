package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerManagement resource defined in the Terraform configuration
func ComputerManagementSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ebooks": {
			Type: schema.TypeList, //GoType: []*ComputerManagementEbooksItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementEbooksItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: ComputerManagementGeneral
			Elem: &schema.Resource{
				Schema: ComputerManagementGeneralSchema(),
			},
			Optional: true,
		},

		"mac_app_store_apps": {
			Type: schema.TypeList, //GoType: []*ComputerManagementMacAppStoreAppsItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementMacAppStoreAppsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"managed_preference_profiles": {
			Type: schema.TypeList, //GoType: []*ComputerManagementManagedPreferenceProfilesItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementManagedPreferenceProfilesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"os_x_configuration_profiles": {
			Type: schema.TypeList, //GoType: []*ComputerManagementOsxConfigurationProfilesItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementOsxConfigurationProfilesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"patch_policies": {
			Type: schema.TypeList, //GoType: []*ComputerManagementPatchPoliciesItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementPatchPoliciesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"patch_reporting_software_titles": {
			Type: schema.TypeList, //GoType: []*ComputerManagementPatchReportingSoftwareTitlesItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementPatchReportingSoftwareTitlesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"policies": {
			Type: schema.TypeList, //GoType: []*ComputerManagementPoliciesItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementPoliciesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"restricted_software": {
			Type: schema.TypeList, //GoType: []*ComputerManagementRestrictedSoftwareItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementRestrictedSoftwareItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"smart_groups": {
			Type: schema.TypeList, //GoType: []*ComputerManagementSmartGroupsItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementSmartGroupsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"static_groups": {
			Type: schema.TypeList, //GoType: []*ComputerManagementStaticGroupsItems0
			Elem: &schema.Resource{
				Schema: ComputerManagementStaticGroupsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Update the underlying ComputerManagement resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerManagementResourceData(d *schema.ResourceData, m *models.ComputerManagement) {
	d.Set("ebooks", SetComputerManagementEbooksItems0SubResourceData(m.Ebooks))
	d.Set("general", SetComputerManagementGeneralSubResourceData([]*models.ComputerManagementGeneral{m.General}))
	d.Set("mac_app_store_apps", SetComputerManagementMacAppStoreAppsItems0SubResourceData(m.MacAppStoreApps))
	d.Set("managed_preference_profiles", SetComputerManagementManagedPreferenceProfilesItems0SubResourceData(m.ManagedPreferenceProfiles))
	d.Set("os_x_configuration_profiles", SetComputerManagementOsxConfigurationProfilesItems0SubResourceData(m.OsxConfigurationProfiles))
	d.Set("patch_policies", SetComputerManagementPatchPoliciesItems0SubResourceData(m.PatchPolicies))
	d.Set("patch_reporting_software_titles", SetComputerManagementPatchReportingSoftwareTitlesItems0SubResourceData(m.PatchReportingSoftwareTitles))
	d.Set("policies", SetComputerManagementPoliciesItems0SubResourceData(m.Policies))
	d.Set("restricted_software", SetComputerManagementRestrictedSoftwareItems0SubResourceData(m.RestrictedSoftware))
	d.Set("smart_groups", SetComputerManagementSmartGroupsItems0SubResourceData(m.SmartGroups))
	d.Set("static_groups", SetComputerManagementStaticGroupsItems0SubResourceData(m.StaticGroups))
}

// Iterate throught and update the ComputerManagement resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerManagementSubResourceData(m []*models.ComputerManagement) (d []*map[string]interface{}) {
	for _, computerManagement := range m {
		if computerManagement != nil {
			properties := make(map[string]interface{})
			properties["ebooks"] = SetComputerManagementEbooksItems0SubResourceData(computerManagement.Ebooks)
			properties["general"] = SetComputerManagementGeneralSubResourceData([]*models.ComputerManagementGeneral{computerManagement.General})
			properties["mac_app_store_apps"] = SetComputerManagementMacAppStoreAppsItems0SubResourceData(computerManagement.MacAppStoreApps)
			properties["managed_preference_profiles"] = SetComputerManagementManagedPreferenceProfilesItems0SubResourceData(computerManagement.ManagedPreferenceProfiles)
			properties["os_x_configuration_profiles"] = SetComputerManagementOsxConfigurationProfilesItems0SubResourceData(computerManagement.OsxConfigurationProfiles)
			properties["patch_policies"] = SetComputerManagementPatchPoliciesItems0SubResourceData(computerManagement.PatchPolicies)
			properties["patch_reporting_software_titles"] = SetComputerManagementPatchReportingSoftwareTitlesItems0SubResourceData(computerManagement.PatchReportingSoftwareTitles)
			properties["policies"] = SetComputerManagementPoliciesItems0SubResourceData(computerManagement.Policies)
			properties["restricted_software"] = SetComputerManagementRestrictedSoftwareItems0SubResourceData(computerManagement.RestrictedSoftware)
			properties["smart_groups"] = SetComputerManagementSmartGroupsItems0SubResourceData(computerManagement.SmartGroups)
			properties["static_groups"] = SetComputerManagementStaticGroupsItems0SubResourceData(computerManagement.StaticGroups)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerManagement resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerManagementModel(d *schema.ResourceData) *models.ComputerManagement {
	ebooks := d.Get("ebooks").([]*models.ComputerManagementEbooksItems0)
	var general *models.ComputerManagementGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerManagementGeneralModel(generalMap)
	}
	macAppStoreApps := d.Get("mac_app_store_apps").([]*models.ComputerManagementMacAppStoreAppsItems0)
	managedPreferenceProfiles := d.Get("managed_preference_profiles").([]*models.ComputerManagementManagedPreferenceProfilesItems0)
	osxConfigurationProfiles := d.Get("os_x_configuration_profiles").([]*models.ComputerManagementOsxConfigurationProfilesItems0)
	patchPolicies := d.Get("patch_policies").([]*models.ComputerManagementPatchPoliciesItems0)
	patchReportingSoftwareTitles := d.Get("patch_reporting_software_titles").([]*models.ComputerManagementPatchReportingSoftwareTitlesItems0)
	policies := d.Get("policies").([]*models.ComputerManagementPoliciesItems0)
	restrictedSoftware := d.Get("restricted_software").([]*models.ComputerManagementRestrictedSoftwareItems0)
	smartGroups := d.Get("smart_groups").([]*models.ComputerManagementSmartGroupsItems0)
	staticGroups := d.Get("static_groups").([]*models.ComputerManagementStaticGroupsItems0)

	return &models.ComputerManagement{
		Ebooks:                       ebooks,
		General:                      general,
		MacAppStoreApps:              macAppStoreApps,
		ManagedPreferenceProfiles:    managedPreferenceProfiles,
		OsxConfigurationProfiles:     osxConfigurationProfiles,
		PatchPolicies:                patchPolicies,
		PatchReportingSoftwareTitles: patchReportingSoftwareTitles,
		Policies:                     policies,
		RestrictedSoftware:           restrictedSoftware,
		SmartGroups:                  smartGroups,
		StaticGroups:                 staticGroups,
	}
}

// Retrieve property field names for updating the ComputerManagement resource
func GetComputerManagementPropertyFields() (t []string) {
	return []string{
		"ebooks",
		"general",
		"mac_app_store_apps",
		"managed_preference_profiles",
		"os_x_configuration_profiles",
		"patch_policies",
		"patch_reporting_software_titles",
		"policies",
		"restricted_software",
		"smart_groups",
		"static_groups",
	}
}
