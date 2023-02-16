package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the HealthcareListenerRule resource defined in the Terraform configuration
func HealthcareListenerRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"adt_message": {
			Type:     schema.TypeString,
			Required: true,
		},

		"adt_message_field": {
			Type:     schema.TypeString,
			Required: true,
		},

		"device_inventory_field": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"hcl_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"mdm_command": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"mdm_command_additional_data": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"notification_emails": {
			Type: schema.TypeList, //GoType: HealthcareListenerRuleNotificationEmails
			Elem: &schema.Resource{
				Schema: HealthcareListenerRuleNotificationEmailsSchema(),
			},
			Optional: true,
		},

		"notification_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"notification_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"notify_unsupported_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"operating_system": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying HealthcareListenerRule resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetHealthcareListenerRuleResourceData(d *schema.ResourceData, m *models.HealthcareListenerRule) {
	d.Set("adt_message", m.AdtMessage)
	d.Set("adt_message_field", m.AdtMessageField)
	d.Set("device_inventory_field", m.DeviceInventoryField)
	d.Set("hcl_id", m.HclID)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("mdm_command", m.MdmCommand)
	d.Set("mdm_command_additional_data", m.MdmCommandAdditionalData)
	d.Set("name", m.Name)
	d.Set("notification_emails", SetHealthcareListenerRuleNotificationEmailsSubResourceData([]*models.HealthcareListenerRuleNotificationEmails{m.NotificationEmails}))
	d.Set("notification_enabled", m.NotificationEnabled)
	d.Set("notification_threshold", m.NotificationThreshold)
	d.Set("notify_unsupported_enabled", m.NotifyUnsupportedEnabled)
	d.Set("operating_system", m.OperatingSystem)
}

// Iterate throught and update the HealthcareListenerRule resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetHealthcareListenerRuleSubResourceData(m []*models.HealthcareListenerRule) (d []*map[string]interface{}) {
	for _, healthcareListenerRule := range m {
		if healthcareListenerRule != nil {
			properties := make(map[string]interface{})
			properties["adt_message"] = healthcareListenerRule.AdtMessage
			properties["adt_message_field"] = healthcareListenerRule.AdtMessageField
			properties["device_inventory_field"] = healthcareListenerRule.DeviceInventoryField
			properties["hcl_id"] = healthcareListenerRule.HclID
			properties["id"] = healthcareListenerRule.ID
			properties["mdm_command"] = healthcareListenerRule.MdmCommand
			properties["mdm_command_additional_data"] = healthcareListenerRule.MdmCommandAdditionalData
			properties["name"] = healthcareListenerRule.Name
			properties["notification_emails"] = SetHealthcareListenerRuleNotificationEmailsSubResourceData([]*models.HealthcareListenerRuleNotificationEmails{healthcareListenerRule.NotificationEmails})
			properties["notification_enabled"] = healthcareListenerRule.NotificationEnabled
			properties["notification_threshold"] = healthcareListenerRule.NotificationThreshold
			properties["notify_unsupported_enabled"] = healthcareListenerRule.NotifyUnsupportedEnabled
			properties["operating_system"] = healthcareListenerRule.OperatingSystem
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate HealthcareListenerRule resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func HealthcareListenerRuleModel(d *schema.ResourceData) *models.HealthcareListenerRule {
	adtMessage := d.Get("adt_message").(string)
	adtMessageField := d.Get("adt_message_field").(string)
	deviceInventoryField := int64(d.Get("device_inventory_field").(int))
	hclID := int64(d.Get("hcl_id").(int))
	id, _ := strconv.Atoi(d.Get("id").(string))
	mdmCommand := d.Get("mdm_command").(string)
	mdmCommandAdditionalData := d.Get("mdm_command_additional_data").(string)
	name := d.Get("name").(string)
	var notificationEmails *models.HealthcareListenerRuleNotificationEmails = nil //hit complex
	notification_emailsInterface, notification_emailsIsSet := d.GetOk("notification_emails")
	if notification_emailsIsSet {
		notification_emailsMap := notification_emailsInterface.([]interface{})[0].(map[string]interface{})
		notificationEmails = HealthcareListenerRuleNotificationEmailsModel(notification_emailsMap)
	}
	notificationEnabled := d.Get("notification_enabled").(bool)
	notificationThreshold := int64(d.Get("notification_threshold").(int))
	notifyUnsupportedEnabled := d.Get("notify_unsupported_enabled").(bool)
	operatingSystem := d.Get("operating_system").(string)

	return &models.HealthcareListenerRule{
		AdtMessage:               &adtMessage,
		AdtMessageField:          &adtMessageField,
		DeviceInventoryField:     &deviceInventoryField,
		HclID:                    hclID,
		ID:                       int32(id),
		MdmCommand:               mdmCommand,
		MdmCommandAdditionalData: mdmCommandAdditionalData,
		Name:                     &name,
		NotificationEmails:       notificationEmails,
		NotificationEnabled:      &notificationEnabled,
		NotificationThreshold:    &notificationThreshold,
		NotifyUnsupportedEnabled: &notifyUnsupportedEnabled,
		OperatingSystem:          operatingSystem,
	}
}

// Retrieve property field names for updating the HealthcareListenerRule resource
func GetHealthcareListenerRulePropertyFields() (t []string) {
	return []string{
		"adt_message",
		"adt_message_field",
		"device_inventory_field",
		"hcl_id",
		"id",
		"mdm_command",
		"mdm_command_additional_data",
		"name",
		"notification_emails",
		"notification_enabled",
		"notification_threshold",
		"notify_unsupported_enabled",
		"operating_system",
	}
}
