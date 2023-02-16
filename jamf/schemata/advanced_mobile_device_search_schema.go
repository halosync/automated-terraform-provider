package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedMobileDeviceSearch resource defined in the Terraform configuration
func AdvancedMobileDeviceSearchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"criteria": {
			Type: schema.TypeList, //GoType: []*AdvancedMobileDeviceSearchCriteriaItems0
			Elem: &schema.Resource{
				Schema: AdvancedMobileDeviceSearchCriteriaItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"display_fields": {
			Type: schema.TypeList, //GoType: []*AdvancedMobileDeviceSearchDisplayFieldsItems0
			Elem: &schema.Resource{
				Schema: AdvancedMobileDeviceSearchDisplayFieldsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"mobile_devices": {
			Type: schema.TypeList, //GoType: []*AdvancedMobileDeviceSearchMobileDevicesItems0
			Elem: &schema.Resource{
				Schema: AdvancedMobileDeviceSearchMobileDevicesItems0Schema(),
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

		"sort_1": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"sort_2": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"sort_3": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"view_as": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying AdvancedMobileDeviceSearch resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedMobileDeviceSearchResourceData(d *schema.ResourceData, m *models.AdvancedMobileDeviceSearch) {
	d.Set("criteria", SetAdvancedMobileDeviceSearchCriteriaItems0SubResourceData(m.Criteria))
	d.Set("display_fields", SetAdvancedMobileDeviceSearchDisplayFieldsItems0SubResourceData(m.DisplayFields))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("mobile_devices", SetAdvancedMobileDeviceSearchMobileDevicesItems0SubResourceData(m.MobileDevices))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("sort_1", m.Sort1)
	d.Set("sort_2", m.Sort2)
	d.Set("sort_3", m.Sort3)
	d.Set("view_as", m.ViewAs)
}

// Iterate throught and update the AdvancedMobileDeviceSearch resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedMobileDeviceSearchSubResourceData(m []*models.AdvancedMobileDeviceSearch) (d []*map[string]interface{}) {
	for _, advancedMobileDeviceSearch := range m {
		if advancedMobileDeviceSearch != nil {
			properties := make(map[string]interface{})
			properties["criteria"] = SetAdvancedMobileDeviceSearchCriteriaItems0SubResourceData(advancedMobileDeviceSearch.Criteria)
			properties["display_fields"] = SetAdvancedMobileDeviceSearchDisplayFieldsItems0SubResourceData(advancedMobileDeviceSearch.DisplayFields)
			properties["id"] = advancedMobileDeviceSearch.ID
			properties["mobile_devices"] = SetAdvancedMobileDeviceSearchMobileDevicesItems0SubResourceData(advancedMobileDeviceSearch.MobileDevices)
			properties["name"] = advancedMobileDeviceSearch.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{advancedMobileDeviceSearch.Site})
			properties["sort_1"] = advancedMobileDeviceSearch.Sort1
			properties["sort_2"] = advancedMobileDeviceSearch.Sort2
			properties["sort_3"] = advancedMobileDeviceSearch.Sort3
			properties["view_as"] = advancedMobileDeviceSearch.ViewAs
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedMobileDeviceSearch resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedMobileDeviceSearchModel(d *schema.ResourceData) *models.AdvancedMobileDeviceSearch {
	criteria := d.Get("criteria").([]*models.AdvancedMobileDeviceSearchCriteriaItems0)
	displayFields := d.Get("display_fields").([]*models.AdvancedMobileDeviceSearchDisplayFieldsItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
	mobileDevices := d.Get("mobile_devices").([]*models.AdvancedMobileDeviceSearchMobileDevicesItems0)
	name := d.Get("name").(string)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	sort1 := d.Get("sort_1").(string)
	sort2 := d.Get("sort_2").(string)
	sort3 := d.Get("sort_3").(string)
	viewAs := d.Get("view_as").(string)

	return &models.AdvancedMobileDeviceSearch{
		Criteria:      criteria,
		DisplayFields: displayFields,
		ID:            int32(id),
		MobileDevices: mobileDevices,
		Name:          &name,
		Site:          site,
		Sort1:         sort1,
		Sort2:         sort2,
		Sort3:         sort3,
		ViewAs:        viewAs,
	}
}

// Retrieve property field names for updating the AdvancedMobileDeviceSearch resource
func GetAdvancedMobileDeviceSearchPropertyFields() (t []string) {
	return []string{
		"criteria",
		"display_fields",
		"id",
		"mobile_devices",
		"name",
		"site",
		"sort_1",
		"sort_2",
		"sort_3",
		"view_as",
	}
}
