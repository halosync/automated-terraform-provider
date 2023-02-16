package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerInventoryCollection resource defined in the Terraform configuration
func ComputerInventoryCollectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_services": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"applications": {
			Type: schema.TypeList, //GoType: []*ComputerInventoryCollectionApplicationsItems0
			Elem: &schema.Resource{
				Schema: ComputerInventoryCollectionApplicationsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"available_software_updates": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"computer_location_information": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"fonts": {
			Type: schema.TypeList, //GoType: []*ComputerInventoryCollectionFontsItems0
			Elem: &schema.Resource{
				Schema: ComputerInventoryCollectionFontsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"hidden_accounts": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"home_directory_sizes": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"inclue_applications": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"inclue_fonts": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"inclue_plugins": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"local_user_accounts": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"mobile_device_app_purchasing_info": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"package_receipts": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"plugins": {
			Type: schema.TypeList, //GoType: []*ComputerInventoryCollectionPluginsItems0
			Elem: &schema.Resource{
				Schema: ComputerInventoryCollectionPluginsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"printers": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

// Update the underlying ComputerInventoryCollection resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerInventoryCollectionResourceData(d *schema.ResourceData, m *models.ComputerInventoryCollection) {
	d.Set("active_services", m.ActiveServices)
	d.Set("applications", SetComputerInventoryCollectionApplicationsItems0SubResourceData(m.Applications))
	d.Set("available_software_updates", m.AvailableSoftwareUpdates)
	d.Set("computer_location_information", m.ComputerLocationInformation)
	d.Set("fonts", SetComputerInventoryCollectionFontsItems0SubResourceData(m.Fonts))
	d.Set("hidden_accounts", m.HiddenAccounts)
	d.Set("home_directory_sizes", m.HomeDirectorySizes)
	d.Set("inclue_applications", m.InclueApplications)
	d.Set("inclue_fonts", m.InclueFonts)
	d.Set("inclue_plugins", m.IncluePlugins)
	d.Set("local_user_accounts", m.LocalUserAccounts)
	d.Set("mobile_device_app_purchasing_info", m.MobileDeviceAppPurchasingInfo)
	d.Set("package_receipts", m.PackageReceipts)
	d.Set("plugins", SetComputerInventoryCollectionPluginsItems0SubResourceData(m.Plugins))
	d.Set("printers", m.Printers)
}

// Iterate throught and update the ComputerInventoryCollection resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerInventoryCollectionSubResourceData(m []*models.ComputerInventoryCollection) (d []*map[string]interface{}) {
	for _, computerInventoryCollection := range m {
		if computerInventoryCollection != nil {
			properties := make(map[string]interface{})
			properties["active_services"] = computerInventoryCollection.ActiveServices
			properties["applications"] = SetComputerInventoryCollectionApplicationsItems0SubResourceData(computerInventoryCollection.Applications)
			properties["available_software_updates"] = computerInventoryCollection.AvailableSoftwareUpdates
			properties["computer_location_information"] = computerInventoryCollection.ComputerLocationInformation
			properties["fonts"] = SetComputerInventoryCollectionFontsItems0SubResourceData(computerInventoryCollection.Fonts)
			properties["hidden_accounts"] = computerInventoryCollection.HiddenAccounts
			properties["home_directory_sizes"] = computerInventoryCollection.HomeDirectorySizes
			properties["inclue_applications"] = computerInventoryCollection.InclueApplications
			properties["inclue_fonts"] = computerInventoryCollection.InclueFonts
			properties["inclue_plugins"] = computerInventoryCollection.IncluePlugins
			properties["local_user_accounts"] = computerInventoryCollection.LocalUserAccounts
			properties["mobile_device_app_purchasing_info"] = computerInventoryCollection.MobileDeviceAppPurchasingInfo
			properties["package_receipts"] = computerInventoryCollection.PackageReceipts
			properties["plugins"] = SetComputerInventoryCollectionPluginsItems0SubResourceData(computerInventoryCollection.Plugins)
			properties["printers"] = computerInventoryCollection.Printers
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerInventoryCollection resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerInventoryCollectionModel(d *schema.ResourceData) *models.ComputerInventoryCollection {
	activeServices := d.Get("active_services").(bool)
	applications := d.Get("applications").([]*models.ComputerInventoryCollectionApplicationsItems0)
	availableSoftwareUpdates := d.Get("available_software_updates").(bool)
	computerLocationInformation := d.Get("computer_location_information").(bool)
	fonts := d.Get("fonts").([]*models.ComputerInventoryCollectionFontsItems0)
	hiddenAccounts := d.Get("hidden_accounts").(bool)
	homeDirectorySizes := d.Get("home_directory_sizes").(bool)
	inclueApplications := d.Get("inclue_applications").(bool)
	inclueFonts := d.Get("inclue_fonts").(bool)
	incluePlugins := d.Get("inclue_plugins").(bool)
	localUserAccounts := d.Get("local_user_accounts").(bool)
	mobileDeviceAppPurchasingInfo := d.Get("mobile_device_app_purchasing_info").(bool)
	packageReceipts := d.Get("package_receipts").(bool)
	plugins := d.Get("plugins").([]*models.ComputerInventoryCollectionPluginsItems0)
	printers := d.Get("printers").(bool)

	return &models.ComputerInventoryCollection{
		ActiveServices:                activeServices,
		Applications:                  applications,
		AvailableSoftwareUpdates:      availableSoftwareUpdates,
		ComputerLocationInformation:   computerLocationInformation,
		Fonts:                         fonts,
		HiddenAccounts:                hiddenAccounts,
		HomeDirectorySizes:            homeDirectorySizes,
		InclueApplications:            inclueApplications,
		InclueFonts:                   inclueFonts,
		IncluePlugins:                 incluePlugins,
		LocalUserAccounts:             localUserAccounts,
		MobileDeviceAppPurchasingInfo: mobileDeviceAppPurchasingInfo,
		PackageReceipts:               packageReceipts,
		Plugins:                       plugins,
		Printers:                      printers,
	}
}

// Retrieve property field names for updating the ComputerInventoryCollection resource
func GetComputerInventoryCollectionPropertyFields() (t []string) {
	return []string{
		"active_services",
		"applications",
		"available_software_updates",
		"computer_location_information",
		"fonts",
		"hidden_accounts",
		"home_directory_sizes",
		"inclue_applications",
		"inclue_fonts",
		"inclue_plugins",
		"local_user_accounts",
		"mobile_device_app_purchasing_info",
		"package_receipts",
		"plugins",
		"printers",
	}
}
