package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceGroup resource defined in the Terraform configuration
func MobileDeviceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"criteria": {
			Type: schema.TypeList, //GoType: []*MobileDeviceGroupCriteriaItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceGroupCriteriaItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"is_smart": {
			Type:     schema.TypeBool,
			Required: true,
		},

		"mobile_devices": {
			Type: schema.TypeList, //GoType: []*MobileDeviceGroupMobileDevicesItems0
			Elem: &schema.Resource{
				Schema: MobileDeviceGroupMobileDevicesItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: SiteObject
			Elem: &schema.Resource{
				Schema: SiteObjectSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceGroup resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceGroupResourceData(d *schema.ResourceData, m *models.MobileDeviceGroup) {
	d.Set("criteria", SetMobileDeviceGroupCriteriaItems0SubResourceData(m.Criteria))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("is_smart", m.IsSmart)
	d.Set("mobile_devices", SetMobileDeviceGroupMobileDevicesItems0SubResourceData(m.MobileDevices))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
}

// Iterate throught and update the MobileDeviceGroup resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceGroupSubResourceData(m []*models.MobileDeviceGroup) (d []*map[string]interface{}) {
	for _, mobileDeviceGroup := range m {
		if mobileDeviceGroup != nil {
			properties := make(map[string]interface{})
			properties["criteria"] = SetMobileDeviceGroupCriteriaItems0SubResourceData(mobileDeviceGroup.Criteria)
			properties["id"] = mobileDeviceGroup.ID
			properties["is_smart"] = mobileDeviceGroup.IsSmart
			properties["mobile_devices"] = SetMobileDeviceGroupMobileDevicesItems0SubResourceData(mobileDeviceGroup.MobileDevices)
			properties["name"] = mobileDeviceGroup.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{mobileDeviceGroup.Site})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceGroup resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceGroupModel(d *schema.ResourceData) *models.MobileDeviceGroup {
	criteria := d.Get("criteria").([]*models.MobileDeviceGroupCriteriaItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
	isSmart := d.Get("is_smart").(bool)
	mobileDevices := d.Get("mobile_devices").([]*models.MobileDeviceGroupMobileDevicesItems0)
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}

	return &models.MobileDeviceGroup{
		Criteria:      criteria,
		ID:            int32(id),
		IsSmart:       &isSmart,
		MobileDevices: mobileDevices,
		Name:          &name,
		Site:          site,
	}
}

// Retrieve property field names for updating the MobileDeviceGroup resource
func GetMobileDeviceGroupPropertyFields() (t []string) {
	return []string{
		"criteria",
		"id",
		"is_smart",
		"mobile_devices",
		"name",
		"site",
	}
}
