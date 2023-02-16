package schemata

import (
)

// Schema mapping representing the Package resource defined in the Terraform configuration
func PackageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_uninstalled": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"category": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"filename": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"fill_existing_users": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"fill_user_template": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"id": {
			Type: schema.TypeString,
			Computed: true,
		},
		
		"info": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"install_if_reported_available": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"name": {
			Type: schema.TypeString,
			Required: true,
		},
		
		"notes": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"os_requirements": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"priority": {
			Type: schema.TypeInt,
			Optional: true,
		},
		
		"reboot_required": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"reinstall_option": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"required_processor": {
			Type: schema.TypeString,
			Default: "None",
			Optional: true,
		},
		
		"send_notification": {
			Type: schema.TypeBool,
			Optional: true,
		},
		
		"switch_with_package": {
			Type: schema.TypeString,
			Optional: true,
		},
		
		"triggering_files": {
			Type: schema.TypeString,
			Optional: true,
		},
		
	}
}



// Update the underlying Package resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPackageResourceData(d *schema.ResourceData, m *models.Package) {
	d.Set("allow_uninstalled", m.AllowUninstalled)
	d.Set("category", m.Category)
	d.Set("filename", m.Filename)
	d.Set("fill_existing_users", m.FillExistingUsers)
	d.Set("fill_user_template", m.FillUserTemplate)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("info", m.Info)
	d.Set("install_if_reported_available", m.InstallIfReportedAvailable)
	d.Set("name", m.Name)
	d.Set("notes", m.Notes)
	d.Set("os_requirements", m.OsRequirements)
	d.Set("priority", m.Priority)
	d.Set("reboot_required", m.RebootRequired)
	d.Set("reinstall_option", m.ReinstallOption)
	d.Set("required_processor", m.RequiredProcessor)
	d.Set("send_notification", m.SendNotification)
	d.Set("switch_with_package", m.SwitchWithPackage)
	d.Set("triggering_files", m.TriggeringFiles)
}

// Iterate throught and update the Package resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPackageSubResourceData(m []*models.Package) (d []*map[string]interface{}) {
	for _, package := range m {
		if package != nil {
			properties := make(map[string]interface{})
			properties["allow_uninstalled"] = package.AllowUninstalled
			properties["category"] = package.Category
			properties["filename"] = package.Filename
			properties["fill_existing_users"] = package.FillExistingUsers
			properties["fill_user_template"] = package.FillUserTemplate
			properties["id"] = package.ID
			properties["info"] = package.Info
			properties["install_if_reported_available"] = package.InstallIfReportedAvailable
			properties["name"] = package.Name
			properties["notes"] = package.Notes
			properties["os_requirements"] = package.OsRequirements
			properties["priority"] = package.Priority
			properties["reboot_required"] = package.RebootRequired
			properties["reinstall_option"] = package.ReinstallOption
			properties["required_processor"] = package.RequiredProcessor
			properties["send_notification"] = package.SendNotification
			properties["switch_with_package"] = package.SwitchWithPackage
			properties["triggering_files"] = package.TriggeringFiles
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Package resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PackageModel(d *schema.ResourceData) *models.Package {
	allowUninstalled := d.Get("allow_uninstalled").(bool)
	category := d.Get("category").(string)
	filename := d.Get("filename").(string)
	fillExistingUsers := d.Get("fill_existing_users").(bool)
	fillUserTemplate := d.Get("fill_user_template").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	info := d.Get("info").(string)
	installIfReportedAvailable := d.Get("install_if_reported_available").(bool)
	name := d.Get("name").(string)
	notes := d.Get("notes").(string)
	osRequirements := d.Get("os_requirements").(string)
	priority := int64(d.Get("priority").(int))
	rebootRequired := d.Get("reboot_required").(bool)
	reinstallOption := d.Get("reinstall_option").(string)
	requiredProcessor := d.Get("required_processor").(string)
	sendNotification := d.Get("send_notification").(bool)
	switchWithPackage := d.Get("switch_with_package").(string)
	triggeringFiles := d.Get("triggering_files").(string)
	
	return &models.Package {
		AllowUninstalled: allowUninstalled,
		Category: category,
		Filename: &filename,
		FillExistingUsers: fillExistingUsers,
		FillUserTemplate: fillUserTemplate,
		ID: int32(id),
		Info: info,
		InstallIfReportedAvailable: installIfReportedAvailable,
		Name: &name,
		Notes: notes,
		OsRequirements: osRequirements,
		Priority: priority,
		RebootRequired: rebootRequired,
		ReinstallOption: reinstallOption,
		RequiredProcessor: &requiredProcessor,
		SendNotification: sendNotification,
		SwitchWithPackage: switchWithPackage,
		TriggeringFiles: triggeringFiles,
	}
}

// Retrieve property field names for updating the Package resource 
func GetPackagePropertyFields() (t []string) {
	return []string{
		"allow_uninstalled",
		"category",
		"filename",
		"fill_existing_users",
		"fill_user_template",
		"id",
		"info",
		"install_if_reported_available",
		"name",
		"notes",
		"os_requirements",
		"priority",
		"reboot_required",
		"reinstall_option",
		"required_processor",
		"send_notification",
		"switch_with_package",
		"triggering_files",
	}
}