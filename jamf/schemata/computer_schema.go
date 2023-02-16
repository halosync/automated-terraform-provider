package schemata

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Schema mapping representing the Computer resource defined in the Terraform configuration
func ComputerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificates": {
			Type: schema.TypeList, //GoType: []*ComputerCertificatesItems0
			Elem: &schema.Resource{
				Schema: ComputerCertificatesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"configuration_profiles": {
			Type: schema.TypeList, //GoType: []*ComputerConfigurationProfilesItems0
			Elem: &schema.Resource{
				Schema: ComputerConfigurationProfilesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"extension_attributes": {
			Type: schema.TypeList, //GoType: []*ComputerExtensionAttributesItems0
			Elem: &schema.Resource{
				Schema: ComputerExtensionAttributesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: ComputerGeneral
			Elem: &schema.Resource{
				Schema: ComputerGeneralSchema(),
			},
			Optional: true,
		},

		"groups_accounts": {
			Type: schema.TypeList, //GoType: ComputerGroupsAccounts
			Elem: &schema.Resource{
				Schema: ComputerGroupsAccountsSchema(),
			},
			Optional: true,
		},

		"hardware": {
			Type: schema.TypeList, //GoType: ComputerHardware
			Elem: &schema.Resource{
				Schema: ComputerHardwareSchema(),
			},
			Optional: true,
		},

		"location": {
			Type: schema.TypeList, //GoType: Location
			Elem: &schema.Resource{
				Schema: LocationSchema(),
			},
			Optional: true,
		},

		"peripherals": {
			Type: schema.TypeList, //GoType: ComputerPeripherals
			Elem: &schema.Resource{
				Schema: ComputerPeripheralsSchema(),
			},
			Optional: true,
		},

		"purchasing": {
			Type: schema.TypeList, //GoType: Purchasing
			Elem: &schema.Resource{
				Schema: PurchasingSchema(),
			},
			Optional: true,
		},

		"security": {
			Type: schema.TypeList, //GoType: ComputerSecurity
			Elem: &schema.Resource{
				Schema: ComputerSecuritySchema(),
			},
			Optional: true,
		},

		"software": {
			Type: schema.TypeList, //GoType: ComputerSoftware
			Elem: &schema.Resource{
				Schema: ComputerSoftwareSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying Computer resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerResourceData(d *schema.ResourceData, m *models.Computer) {
	d.Set("certificates", SetComputerCertificatesItems0SubResourceData(m.Certificates))
	d.Set("configuration_profiles", SetComputerConfigurationProfilesItems0SubResourceData(m.ConfigurationProfiles))
	d.Set("extension_attributes", SetComputerExtensionAttributesItems0SubResourceData(m.ExtensionAttributes))
	d.Set("general", SetComputerGeneralSubResourceData([]*models.ComputerGeneral{m.General}))
	d.Set("groups_accounts", SetComputerGroupsAccountsSubResourceData([]*models.ComputerGroupsAccounts{m.GroupsAccounts}))
	d.Set("hardware", SetComputerHardwareSubResourceData([]*models.ComputerHardware{m.Hardware}))
	d.Set("location", SetLocationSubResourceData([]*models.Location{m.Location}))
	d.Set("peripherals", SetComputerPeripheralsSubResourceData([]*models.ComputerPeripherals{m.Peripherals}))
	d.Set("purchasing", SetPurchasingSubResourceData([]*models.Purchasing{m.Purchasing}))
	d.Set("security", SetComputerSecuritySubResourceData([]*models.ComputerSecurity{m.Security}))
	d.Set("software", SetComputerSoftwareSubResourceData([]*models.ComputerSoftware{m.Software}))
}

// Iterate throught and update the Computer resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerSubResourceData(m []*models.Computer) (d []*map[string]interface{}) {
	for _, computer := range m {
		if computer != nil {
			properties := make(map[string]interface{})
			properties["certificates"] = SetComputerCertificatesItems0SubResourceData(computer.Certificates)
			properties["configuration_profiles"] = SetComputerConfigurationProfilesItems0SubResourceData(computer.ConfigurationProfiles)
			properties["extension_attributes"] = SetComputerExtensionAttributesItems0SubResourceData(computer.ExtensionAttributes)
			properties["general"] = SetComputerGeneralSubResourceData([]*models.ComputerGeneral{computer.General})
			properties["groups_accounts"] = SetComputerGroupsAccountsSubResourceData([]*models.ComputerGroupsAccounts{computer.GroupsAccounts})
			properties["hardware"] = SetComputerHardwareSubResourceData([]*models.ComputerHardware{computer.Hardware})
			properties["location"] = SetLocationSubResourceData([]*models.Location{computer.Location})
			properties["peripherals"] = SetComputerPeripheralsSubResourceData([]*models.ComputerPeripherals{computer.Peripherals})
			properties["purchasing"] = SetPurchasingSubResourceData([]*models.Purchasing{computer.Purchasing})
			properties["security"] = SetComputerSecuritySubResourceData([]*models.ComputerSecurity{computer.Security})
			properties["software"] = SetComputerSoftwareSubResourceData([]*models.ComputerSoftware{computer.Software})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Computer resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerModel(d *schema.ResourceData) *models.Computer {
	certificates := d.Get("certificates").([]*models.ComputerCertificatesItems0)
	configurationProfiles := d.Get("configuration_profiles").([]*models.ComputerConfigurationProfilesItems0)
	extensionAttributes := d.Get("extension_attributes").([]*models.ComputerExtensionAttributesItems0)
	var general *models.ComputerGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = ComputerGeneralModel(generalMap)
	}
	var groupsAccounts *models.ComputerGroupsAccounts = nil //hit complex
	groups_accountsInterface, groups_accountsIsSet := d.GetOk("groups_accounts")
	if groups_accountsIsSet {
		groups_accountsMap := groups_accountsInterface.([]interface{})[0].(map[string]interface{})
		groupsAccounts = ComputerGroupsAccountsModel(groups_accountsMap)
	}
	var hardware *models.ComputerHardware = nil //hit complex
	hardwareInterface, hardwareIsSet := d.GetOk("hardware")
	if hardwareIsSet {
		hardwareMap := hardwareInterface.([]interface{})[0].(map[string]interface{})
		hardware = ComputerHardwareModel(hardwareMap)
	}
	var location *models.Location = nil //hit complex
	locationInterface, locationIsSet := d.GetOk("location")
	if locationIsSet {
		locationMap := locationInterface.([]interface{})[0].(map[string]interface{})
		location = LocationModel(locationMap)
	}
	var peripherals *models.ComputerPeripherals = nil //hit complex
	peripheralsInterface, peripheralsIsSet := d.GetOk("peripherals")
	if peripheralsIsSet {
		peripheralsMap := peripheralsInterface.([]interface{})[0].(map[string]interface{})
		peripherals = ComputerPeripheralsModel(peripheralsMap)
	}
	var purchasing *models.Purchasing = nil //hit complex
	purchasingInterface, purchasingIsSet := d.GetOk("purchasing")
	if purchasingIsSet {
		purchasingMap := purchasingInterface.([]interface{})[0].(map[string]interface{})
		purchasing = PurchasingModel(purchasingMap)
	}
	var security *models.ComputerSecurity = nil //hit complex
	securityInterface, securityIsSet := d.GetOk("security")
	if securityIsSet {
		securityMap := securityInterface.([]interface{})[0].(map[string]interface{})
		security = ComputerSecurityModel(securityMap)
	}
	var software *models.ComputerSoftware = nil //hit complex
	softwareInterface, softwareIsSet := d.GetOk("software")
	if softwareIsSet {
		softwareMap := softwareInterface.([]interface{})[0].(map[string]interface{})
		software = ComputerSoftwareModel(softwareMap)
	}

	return &models.Computer{
		Certificates:          certificates,
		ConfigurationProfiles: configurationProfiles,
		ExtensionAttributes:   extensionAttributes,
		General:               general,
		GroupsAccounts:        groupsAccounts,
		Hardware:              hardware,
		Location:              location,
		Peripherals:           peripherals,
		Purchasing:            purchasing,
		Security:              security,
		Software:              software,
	}
}

// Retrieve property field names for updating the Computer resource
func GetComputerPropertyFields() (t []string) {
	return []string{
		"certificates",
		"configuration_profiles",
		"extension_attributes",
		"general",
		"groups_accounts",
		"hardware",
		"location",
		"peripherals",
		"purchasing",
		"security",
		"software",
	}
}
