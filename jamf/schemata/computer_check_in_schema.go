package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerCheckIn resource defined in the Terraform configuration
func ComputerCheckInSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apply_computer_level_managed_preferences": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"apply_user_level_managed_preferences": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"check_for_policies_at_login_logout": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"check_for_policies_at_startup": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"check_in_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"create_login_logout_hooks": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"create_startup_script": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"display_status_to_user": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"ensure_ssh_is_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"hide_restore_partition": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"log_startup_event": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"log_username": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"perform_login_actions_in_background": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

// Update the underlying ComputerCheckIn resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerCheckInResourceData(d *schema.ResourceData, m *models.ComputerCheckIn) {
	d.Set("apply_computer_level_managed_preferences", m.ApplyComputerLevelManagedPreferences)
	d.Set("apply_user_level_managed_preferences", m.ApplyUserLevelManagedPreferences)
	d.Set("check_for_policies_at_login_logout", m.CheckForPoliciesAtLoginLogout)
	d.Set("check_for_policies_at_startup", m.CheckForPoliciesAtStartup)
	d.Set("check_in_frequency", m.CheckInFrequency)
	d.Set("create_login_logout_hooks", m.CreateLoginLogoutHooks)
	d.Set("create_startup_script", m.CreateStartupScript)
	d.Set("display_status_to_user", m.DisplayStatusToUser)
	d.Set("ensure_ssh_is_enabled", m.EnsureSSHIsEnabled)
	d.Set("hide_restore_partition", m.HideRestorePartition)
	d.Set("log_startup_event", m.LogStartupEvent)
	d.Set("log_username", m.LogUsername)
	d.Set("perform_login_actions_in_background", m.PerformLoginActionsInBackground)
}

// Iterate throught and update the ComputerCheckIn resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerCheckInSubResourceData(m []*models.ComputerCheckIn) (d []*map[string]interface{}) {
	for _, computerCheckIn := range m {
		if computerCheckIn != nil {
			properties := make(map[string]interface{})
			properties["apply_computer_level_managed_preferences"] = computerCheckIn.ApplyComputerLevelManagedPreferences
			properties["apply_user_level_managed_preferences"] = computerCheckIn.ApplyUserLevelManagedPreferences
			properties["check_for_policies_at_login_logout"] = computerCheckIn.CheckForPoliciesAtLoginLogout
			properties["check_for_policies_at_startup"] = computerCheckIn.CheckForPoliciesAtStartup
			properties["check_in_frequency"] = computerCheckIn.CheckInFrequency
			properties["create_login_logout_hooks"] = computerCheckIn.CreateLoginLogoutHooks
			properties["create_startup_script"] = computerCheckIn.CreateStartupScript
			properties["display_status_to_user"] = computerCheckIn.DisplayStatusToUser
			properties["ensure_ssh_is_enabled"] = computerCheckIn.EnsureSSHIsEnabled
			properties["hide_restore_partition"] = computerCheckIn.HideRestorePartition
			properties["log_startup_event"] = computerCheckIn.LogStartupEvent
			properties["log_username"] = computerCheckIn.LogUsername
			properties["perform_login_actions_in_background"] = computerCheckIn.PerformLoginActionsInBackground
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerCheckIn resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerCheckInModel(d *schema.ResourceData) *models.ComputerCheckIn {
	applyComputerLevelManagedPreferences := d.Get("apply_computer_level_managed_preferences").(bool)
	applyUserLevelManagedPreferences := d.Get("apply_user_level_managed_preferences").(bool)
	checkForPoliciesAtLoginLogout := d.Get("check_for_policies_at_login_logout").(bool)
	checkForPoliciesAtStartup := d.Get("check_for_policies_at_startup").(bool)
	checkInFrequency := int64(d.Get("check_in_frequency").(int))
	createLoginLogoutHooks := d.Get("create_login_logout_hooks").(bool)
	createStartupScript := d.Get("create_startup_script").(bool)
	displayStatusToUser := d.Get("display_status_to_user").(bool)
	ensureSSHIsEnabled := d.Get("ensure_ssh_is_enabled").(bool)
	hideRestorePartition := d.Get("hide_restore_partition").(bool)
	logStartupEvent := d.Get("log_startup_event").(bool)
	logUsername := d.Get("log_username").(bool)
	performLoginActionsInBackground := d.Get("perform_login_actions_in_background").(bool)

	return &models.ComputerCheckIn{
		ApplyComputerLevelManagedPreferences: applyComputerLevelManagedPreferences,
		ApplyUserLevelManagedPreferences:     applyUserLevelManagedPreferences,
		CheckForPoliciesAtLoginLogout:        checkForPoliciesAtLoginLogout,
		CheckForPoliciesAtStartup:            checkForPoliciesAtStartup,
		CheckInFrequency:                     checkInFrequency,
		CreateLoginLogoutHooks:               createLoginLogoutHooks,
		CreateStartupScript:                  createStartupScript,
		DisplayStatusToUser:                  displayStatusToUser,
		EnsureSSHIsEnabled:                   ensureSSHIsEnabled,
		HideRestorePartition:                 hideRestorePartition,
		LogStartupEvent:                      logStartupEvent,
		LogUsername:                          logUsername,
		PerformLoginActionsInBackground:      performLoginActionsInBackground,
	}
}

// Retrieve property field names for updating the ComputerCheckIn resource
func GetComputerCheckInPropertyFields() (t []string) {
	return []string{
		"apply_computer_level_managed_preferences",
		"apply_user_level_managed_preferences",
		"check_for_policies_at_login_logout",
		"check_for_policies_at_startup",
		"check_in_frequency",
		"create_login_logout_hooks",
		"create_startup_script",
		"display_status_to_user",
		"ensure_ssh_is_enabled",
		"hide_restore_partition",
		"log_startup_event",
		"log_username",
		"perform_login_actions_in_background",
	}
}
