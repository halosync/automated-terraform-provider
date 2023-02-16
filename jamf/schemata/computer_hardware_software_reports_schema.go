package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerHardwareSoftwareReports resource defined in the Terraform configuration
func ComputerHardwareSoftwareReportsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"font_report": {
			Type: schema.TypeList, //GoType: ComputerHardwareSoftwareReportsFontReport
			Elem: &schema.Resource{
				Schema: ComputerHardwareSoftwareReportsFontReportSchema(),
			},
			Optional: true,
		},

		"hardware_report": {
			Type: schema.TypeList, //GoType: ComputerHardwareSoftwareReportsHardwareReport
			Elem: &schema.Resource{
				Schema: ComputerHardwareSoftwareReportsHardwareReportSchema(),
			},
			Optional: true,
		},

		"plugin_report": {
			Type: schema.TypeList, //GoType: ComputerHardwareSoftwareReportsPluginReport
			Elem: &schema.Resource{
				Schema: ComputerHardwareSoftwareReportsPluginReportSchema(),
			},
			Optional: true,
		},

		"software_report": {
			Type: schema.TypeList, //GoType: ComputerHardwareSoftwareReportsSoftwareReport
			Elem: &schema.Resource{
				Schema: ComputerHardwareSoftwareReportsSoftwareReportSchema(),
			},
			Optional: true,
		},
	}
}

// Update the underlying ComputerHardwareSoftwareReports resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerHardwareSoftwareReportsResourceData(d *schema.ResourceData, m *models.ComputerHardwareSoftwareReports) {
	d.Set("font_report", SetComputerHardwareSoftwareReportsFontReportSubResourceData([]*models.ComputerHardwareSoftwareReportsFontReport{m.FontReport}))
	d.Set("hardware_report", SetComputerHardwareSoftwareReportsHardwareReportSubResourceData([]*models.ComputerHardwareSoftwareReportsHardwareReport{m.HardwareReport}))
	d.Set("plugin_report", SetComputerHardwareSoftwareReportsPluginReportSubResourceData([]*models.ComputerHardwareSoftwareReportsPluginReport{m.PluginReport}))
	d.Set("software_report", SetComputerHardwareSoftwareReportsSoftwareReportSubResourceData([]*models.ComputerHardwareSoftwareReportsSoftwareReport{m.SoftwareReport}))
}

// Iterate throught and update the ComputerHardwareSoftwareReports resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerHardwareSoftwareReportsSubResourceData(m []*models.ComputerHardwareSoftwareReports) (d []*map[string]interface{}) {
	for _, computerHardwareSoftwareReports := range m {
		if computerHardwareSoftwareReports != nil {
			properties := make(map[string]interface{})
			properties["font_report"] = SetComputerHardwareSoftwareReportsFontReportSubResourceData([]*models.ComputerHardwareSoftwareReportsFontReport{computerHardwareSoftwareReports.FontReport})
			properties["hardware_report"] = SetComputerHardwareSoftwareReportsHardwareReportSubResourceData([]*models.ComputerHardwareSoftwareReportsHardwareReport{computerHardwareSoftwareReports.HardwareReport})
			properties["plugin_report"] = SetComputerHardwareSoftwareReportsPluginReportSubResourceData([]*models.ComputerHardwareSoftwareReportsPluginReport{computerHardwareSoftwareReports.PluginReport})
			properties["software_report"] = SetComputerHardwareSoftwareReportsSoftwareReportSubResourceData([]*models.ComputerHardwareSoftwareReportsSoftwareReport{computerHardwareSoftwareReports.SoftwareReport})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerHardwareSoftwareReports resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerHardwareSoftwareReportsModel(d *schema.ResourceData) *models.ComputerHardwareSoftwareReports {
	var fontReport *models.ComputerHardwareSoftwareReportsFontReport = nil //hit complex
	font_reportInterface, font_reportIsSet := d.GetOk("font_report")
	if font_reportIsSet {
		font_reportMap := font_reportInterface.([]interface{})[0].(map[string]interface{})
		fontReport = ComputerHardwareSoftwareReportsFontReportModel(font_reportMap)
	}
	var hardwareReport *models.ComputerHardwareSoftwareReportsHardwareReport = nil //hit complex
	hardware_reportInterface, hardware_reportIsSet := d.GetOk("hardware_report")
	if hardware_reportIsSet {
		hardware_reportMap := hardware_reportInterface.([]interface{})[0].(map[string]interface{})
		hardwareReport = ComputerHardwareSoftwareReportsHardwareReportModel(hardware_reportMap)
	}
	var pluginReport *models.ComputerHardwareSoftwareReportsPluginReport = nil //hit complex
	plugin_reportInterface, plugin_reportIsSet := d.GetOk("plugin_report")
	if plugin_reportIsSet {
		plugin_reportMap := plugin_reportInterface.([]interface{})[0].(map[string]interface{})
		pluginReport = ComputerHardwareSoftwareReportsPluginReportModel(plugin_reportMap)
	}
	var softwareReport *models.ComputerHardwareSoftwareReportsSoftwareReport = nil //hit complex
	software_reportInterface, software_reportIsSet := d.GetOk("software_report")
	if software_reportIsSet {
		software_reportMap := software_reportInterface.([]interface{})[0].(map[string]interface{})
		softwareReport = ComputerHardwareSoftwareReportsSoftwareReportModel(software_reportMap)
	}

	return &models.ComputerHardwareSoftwareReports{
		FontReport:     fontReport,
		HardwareReport: hardwareReport,
		PluginReport:   pluginReport,
		SoftwareReport: softwareReport,
	}
}

// Retrieve property field names for updating the ComputerHardwareSoftwareReports resource
func GetComputerHardwareSoftwareReportsPropertyFields() (t []string) {
	return []string{
		"font_report",
		"hardware_report",
		"plugin_report",
		"software_report",
	}
}
