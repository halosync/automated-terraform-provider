package schemata

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Schema mapping representing the MobileDeviceEnrollmentProfile resource defined in the Terraform configuration
func MobileDeviceEnrollmentProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attachments": {
			Type: schema.TypeList, //GoType: []*MobileDeviceEnrollmentProfileAttachmentsItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceEnrollmentProfileAttachmentsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"general": {
			Type: schema.TypeList, //GoType: MobileDeviceEnrollmentProfileGeneral
			Elem: &schema.Resource{
				Schema: MobileDeviceEnrollmentProfileGeneralSchema(),
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

		"purchasing": {
			Type: schema.TypeList, //GoType: Purchasing
			Elem: &schema.Resource{
				Schema: PurchasingSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceEnrollmentProfile resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceEnrollmentProfileResourceData(d *schema.ResourceData, m *models.MobileDeviceEnrollmentProfile) {
	d.Set("attachments", SetMobileDeviceEnrollmentProfileAttachmentsItems0SubResourceData(m.Attachments))
	d.Set("general", SetMobileDeviceEnrollmentProfileGeneralSubResourceData([]*models.MobileDeviceEnrollmentProfileGeneral{m.General}))
	d.Set("location", SetLocationSubResourceData([]*models.Location{m.Location}))
	d.Set("purchasing", SetPurchasingSubResourceData([]*models.Purchasing{m.Purchasing}))
}

// Iterate throught and update the MobileDeviceEnrollmentProfile resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceEnrollmentProfileSubResourceData(m []*models.MobileDeviceEnrollmentProfile) (d []*map[string]interface{}) {
	for _, mobileDeviceEnrollmentProfile := range m {
		if mobileDeviceEnrollmentProfile != nil {
			properties := make(map[string]interface{})
			properties["attachments"] = SetMobileDeviceEnrollmentProfileAttachmentsItems0SubResourceData(mobileDeviceEnrollmentProfile.Attachments)
			properties["general"] = SetMobileDeviceEnrollmentProfileGeneralSubResourceData([]*models.MobileDeviceEnrollmentProfileGeneral{mobileDeviceEnrollmentProfile.General})
			properties["location"] = SetLocationSubResourceData([]*models.Location{mobileDeviceEnrollmentProfile.Location})
			properties["purchasing"] = SetPurchasingSubResourceData([]*models.Purchasing{mobileDeviceEnrollmentProfile.Purchasing})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceEnrollmentProfile resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceEnrollmentProfileModel(d *schema.ResourceData) *models.MobileDeviceEnrollmentProfile {
	attachments := d.Get("attachments").([]*models.MobileDeviceEnrollmentProfileAttachmentsItems0)
	var general *models.MobileDeviceEnrollmentProfileGeneral = nil //hit complex
	generalInterface, generalIsSet := d.GetOk("general")
	if generalIsSet {
		generalMap := generalInterface.([]interface{})[0].(map[string]interface{})
		general = MobileDeviceEnrollmentProfileGeneralModel(generalMap)
	}
	var location *models.Location = nil //hit complex
	locationInterface, locationIsSet := d.GetOk("location")
	if locationIsSet {
		locationMap := locationInterface.([]interface{})[0].(map[string]interface{})
		location = LocationModel(locationMap)
	}
	var purchasing *models.Purchasing = nil //hit complex
	purchasingInterface, purchasingIsSet := d.GetOk("purchasing")
	if purchasingIsSet {
		purchasingMap := purchasingInterface.([]interface{})[0].(map[string]interface{})
		purchasing = PurchasingModel(purchasingMap)
	}

	return &models.MobileDeviceEnrollmentProfile{
		Attachments: attachments,
		General:     general,
		Location:    location,
		Purchasing:  purchasing,
	}
}

// Retrieve property field names for updating the MobileDeviceEnrollmentProfile resource
func GetMobileDeviceEnrollmentProfilePropertyFields() (t []string) {
	return []string{
		"attachments",
		"general",
		"location",
		"purchasing",
	}
}
