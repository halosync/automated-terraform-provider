package schemata

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Schema mapping representing the MobileDevice resource defined in the Terraform configuration
func MobileDeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applications": {
			Type: schema.TypeList, //GoType: []*MobileDeviceApplicationsItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceApplicationsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"certificates": {
			Type: schema.TypeList, //GoType: []*MobileDeviceCertificatesItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceCertificatesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"configuration_profiles": {
			Type: schema.TypeList, //GoType: []*MobileDeviceConfigurationProfilesItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceConfigurationProfilesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"extension_attributes": {
			Type: schema.TypeList, //GoType: []*MobileDeviceExtensionAttributesItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceExtensionAttributesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceGeneralSchema(),
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

		"mobile_device_groups": {
			Type: schema.TypeList, //GoType: []*MobileDeviceMobileDeviceGroupsItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceMobileDeviceGroupsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"network": {
			Type: schema.TypeList, //GoType: MobileDeviceNetwork
			Elem: &schema.Resource{
				Schema: MobileDeviceNetworkSchema(),
			},
			Optional: true,
		},

		"provisioning_profiles": {
			Type: schema.TypeList, //GoType: []*MobileDeviceProvisioningProfilesItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceProvisioningProfilesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"purchasing": {
			Type: schema.TypeList, //GoType: Purchasing
			Elem: &schema.Resource{
				Schema: PurchasingSchema(),
			},
			Optional: true,
		},

		"security_object": {
			Type: schema.TypeList, //GoType: MobileDeviceSecurityObject
			Elem: &schema.Resource{
				Schema: MobileDeviceSecurityObjectSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDevice resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceResourceData(d *schema.ResourceData, m *models.MobileDevice) {
	d.Set("applications", SetMobileDeviceApplicationsItems0SubResourceData(m.Applications))
	d.Set("certificates", SetMobileDeviceCertificatesItems0SubResourceData(m.Certificates))
	d.Set("configuration_profiles", SetMobileDeviceConfigurationProfilesItems0SubResourceData(m.ConfigurationProfiles))
	d.Set("extension_attributes", SetMobileDeviceExtensionAttributesItems0SubResourceData(m.ExtensionAttributes))
	d.Set("general", SetMobileDeviceGeneralSubResourceData([]*models.MobileDeviceGeneral{m.General}))
	d.Set("location", SetLocationSubResourceData([]*models.Location{m.Location}))
	d.Set("mobile_device_groups", SetMobileDeviceMobileDeviceGroupsItems0SubResourceData(m.MobileDeviceGroups))
	d.Set("network", SetMobileDeviceNetworkSubResourceData([]*models.MobileDeviceNetwork{m.Network}))
	d.Set("provisioning_profiles", SetMobileDeviceProvisioningProfilesItems0SubResourceData(m.ProvisioningProfiles))
	d.Set("purchasing", SetPurchasingSubResourceData([]*models.Purchasing{m.Purchasing}))
	d.Set("security_object", SetMobileDeviceSecurityObjectSubResourceData([]*models.MobileDeviceSecurityObject{m.SecurityObject}))
}

// Iterate throught and update the MobileDevice resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceSubResourceData(m []*models.MobileDevice) (d []*map[string]interface{}) {
	for _, mobileDevice := range m {
		if mobileDevice != nil {
			properties := make(map[string]interface{})
			properties["applications"] = SetMobileDeviceApplicationsItems0SubResourceData(mobileDevice.Applications)
			properties["certificates"] = SetMobileDeviceCertificatesItems0SubResourceData(mobileDevice.Certificates)
			properties["configuration_profiles"] = SetMobileDeviceConfigurationProfilesItems0SubResourceData(mobileDevice.ConfigurationProfiles)
			properties["extension_attributes"] = SetMobileDeviceExtensionAttributesItems0SubResourceData(mobileDevice.ExtensionAttributes)
			properties["general"] = SetMobileDeviceGeneralSubResourceData([]*models.MobileDeviceGeneral{mobileDevice.General})
			properties["location"] = SetLocationSubResourceData([]*models.Location{mobileDevice.Location})
			properties["mobile_device_groups"] = SetMobileDeviceMobileDeviceGroupsItems0SubResourceData(mobileDevice.MobileDeviceGroups)
			properties["network"] = SetMobileDeviceNetworkSubResourceData([]*models.MobileDeviceNetwork{mobileDevice.Network})
			properties["provisioning_profiles"] = SetMobileDeviceProvisioningProfilesItems0SubResourceData(mobileDevice.ProvisioningProfiles)
			properties["purchasing"] = SetPurchasingSubResourceData([]*models.Purchasing{mobileDevice.Purchasing})
			properties["security_object"] = SetMobileDeviceSecurityObjectSubResourceData([]*models.MobileDeviceSecurityObject{mobileDevice.SecurityObject})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDevice resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceModel(d *schema.ResourceData) *models.MobileDevice {
	applications := d.Get("applications").([]*models.MobileDeviceApplicationsItems0)
	certificates := d.Get("certificates").([]*models.MobileDeviceCertificatesItems0)
	configurationProfiles := d.Get("configuration_profiles").([]*models.MobileDeviceConfigurationProfilesItems0)
	extensionAttributes := d.Get("extension_attributes").([]*models.MobileDeviceExtensionAttributesItems0)
	var general *models.MobileDeviceGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceGeneralModel(generalMap)
	}
	var location *models.Location = nil //hit complex
	locationInterface, locationIsSet := d.GetOk("location")
	if locationIsSet {
		locationMap := locationInterface.([]interface{})[0].(map[string]interface{})
		location = LocationModel(locationMap)
	}
	mobileDeviceGroups := d.Get("mobile_device_groups").([]*models.MobileDeviceMobileDeviceGroupsItems0)
	var network *models.MobileDeviceNetwork = nil //hit complex
	networkInterface, networkIsSet := d.GetOk("network")
	if networkIsSet {
		networkMap := networkInterface.([]interface{})[0].(map[string]interface{})
		network = MobileDeviceNetworkModel(networkMap)
	}
	provisioningProfiles := d.Get("provisioning_profiles").([]*models.MobileDeviceProvisioningProfilesItems0)
	var purchasing *models.Purchasing = nil //hit complex
	purchasingInterface, purchasingIsSet := d.GetOk("purchasing")
	if purchasingIsSet {
		purchasingMap := purchasingInterface.([]interface{})[0].(map[string]interface{})
		purchasing = PurchasingModel(purchasingMap)
	}
	var securityObject *models.MobileDeviceSecurityObject = nil //hit complex
	security_objectInterface, security_objectIsSet := d.GetOk("security_object")
	if security_objectIsSet {
		security_objectMap := security_objectInterface.([]interface{})[0].(map[string]interface{})
		securityObject = MobileDeviceSecurityObjectModel(security_objectMap)
	}

	return &models.MobileDevice{
		Applications:          applications,
		Certificates:          certificates,
		ConfigurationProfiles: configurationProfiles,
		ExtensionAttributes:   extensionAttributes,
		General:               general,
		Location:              location,
		MobileDeviceGroups:    mobileDeviceGroups,
		Network:               network,
		ProvisioningProfiles:  provisioningProfiles,
		Purchasing:            purchasing,
		SecurityObject:        securityObject,
	}
}

// Retrieve property field names for updating the MobileDevice resource
func GetMobileDevicePropertyFields() (t []string) {
	return []string{
		"applications",
		"certificates",
		"configuration_profiles",
		"extension_attributes",
		"general",
		"location",
		"mobile_device_groups",
		"network",
		"provisioning_profiles",
		"purchasing",
		"security_object",
	}
}
