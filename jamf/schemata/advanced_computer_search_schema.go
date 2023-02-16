package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the AdvancedComputerSearch resource defined in the Terraform configuration
func AdvancedComputerSearchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computers": {
			Type: schema.TypeList, //GoType: []*AdvancedComputerSearchComputersItems0
			Elem: &schema.Resource{
				Schema: AdvancedComputerSearchComputersItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"criteria": {
			Type: schema.TypeList, //GoType: []*AdvancedComputerSearchCriteriaItems0
			Elem: &schema.Resource{
				Schema: AdvancedComputerSearchCriteriaItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"display_fields": {
			Type: schema.TypeList, //GoType: []*AdvancedComputerSearchDisplayFieldsItems0
			Elem: &schema.Resource{
				Schema: AdvancedComputerSearchDisplayFieldsItems0Schema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
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

// Update the underlying AdvancedComputerSearch resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAdvancedComputerSearchResourceData(d *schema.ResourceData, m *models.AdvancedComputerSearch) {
	d.Set("computers", SetAdvancedComputerSearchComputersItems0SubResourceData(m.Computers))
	d.Set("criteria", SetAdvancedComputerSearchCriteriaItems0SubResourceData(m.Criteria))
	d.Set("display_fields", SetAdvancedComputerSearchDisplayFieldsItems0SubResourceData(m.DisplayFields))
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("name", m.Name)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("sort_1", m.Sort1)
	d.Set("sort_2", m.Sort2)
	d.Set("sort_3", m.Sort3)
	d.Set("view_as", m.ViewAs)
}

// Iterate throught and update the AdvancedComputerSearch resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAdvancedComputerSearchSubResourceData(m []*models.AdvancedComputerSearch) (d []*map[string]interface{}) {
	for _, advancedComputerSearch := range m {
		if advancedComputerSearch != nil {
			properties := make(map[string]interface{})
			properties["computers"] = SetAdvancedComputerSearchComputersItems0SubResourceData(advancedComputerSearch.Computers)
			properties["criteria"] = SetAdvancedComputerSearchCriteriaItems0SubResourceData(advancedComputerSearch.Criteria)
			properties["display_fields"] = SetAdvancedComputerSearchDisplayFieldsItems0SubResourceData(advancedComputerSearch.DisplayFields)
			properties["id"] = advancedComputerSearch.ID
			properties["name"] = advancedComputerSearch.Name
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{advancedComputerSearch.Site})
			properties["sort_1"] = advancedComputerSearch.Sort1
			properties["sort_2"] = advancedComputerSearch.Sort2
			properties["sort_3"] = advancedComputerSearch.Sort3
			properties["view_as"] = advancedComputerSearch.ViewAs
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AdvancedComputerSearch resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AdvancedComputerSearchModel(d *schema.ResourceData) *models.AdvancedComputerSearch {
	computers := d.Get("computers").([]*models.AdvancedComputerSearchComputersItems0)
	criteria := d.Get("criteria").([]*models.AdvancedComputerSearchCriteriaItems0)
	displayFields := d.Get("display_fields").([]*models.AdvancedComputerSearchDisplayFieldsItems0)
	id, _ := strconv.Atoi(d.Get("id").(string))
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

	return &models.AdvancedComputerSearch{
		Computers:     computers,
		Criteria:      criteria,
		DisplayFields: displayFields,
		ID:            int32(id),
		Name:          &name,
		Site:          site,
		Sort1:         sort1,
		Sort2:         sort2,
		Sort3:         sort3,
		ViewAs:        viewAs,
	}
}

// Retrieve property field names for updating the AdvancedComputerSearch resource
func GetAdvancedComputerSearchPropertyFields() (t []string) {
	return []string{
		"computers",
		"criteria",
		"display_fields",
		"id",
		"name",
		"site",
		"sort_1",
		"sort_2",
		"sort_3",
		"view_as",
	}
}
