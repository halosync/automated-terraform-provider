package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Logflush resource defined in the Terraform configuration
func LogflushSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"computers": {
			Type: schema.TypeList, //GoType: LogflushComputers
			Elem: &schema.Resource{
				Schema: LogflushComputersSchema(),
			},
			Optional: true,
		},

		"interval": {
			Type:     schema.TypeString,
			Required: true,
		},

		"log": {
			Type:     schema.TypeString,
			Required: true,
		},

		"log_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

// Update the underlying Logflush resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLogflushResourceData(d *schema.ResourceData, m *models.Logflush) {
	d.Set("computers", SetLogflushComputersSubResourceData([]*models.LogflushComputers{m.Computers}))
	d.Set("interval", m.Interval)
	d.Set("log", m.Log)
	d.Set("log_id", m.LogID)
}

// Iterate throught and update the Logflush resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetLogflushSubResourceData(m []*models.Logflush) (d []*map[string]interface{}) {
	for _, logflush := range m {
		if logflush != nil {
			properties := make(map[string]interface{})
			properties["computers"] = SetLogflushComputersSubResourceData([]*models.LogflushComputers{logflush.Computers})
			properties["interval"] = logflush.Interval
			properties["log"] = logflush.Log
			properties["log_id"] = logflush.LogID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Logflush resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LogflushModel(d *schema.ResourceData) *models.Logflush {
	var computers *models.LogflushComputers = nil //hit complex
	computersInterface, computersIsSet := d.GetOk("computers")
	if computersIsSet {
		computersMap := computersInterface.([]interface{})[0].(map[string]interface{})
		computers = LogflushComputersModel(computersMap)
	}
	interval := d.Get("interval").(string)
	log := d.Get("log").(string)
	logID := int64(d.Get("log_id").(int))

	return &models.Logflush{
		Computers: computers,
		Interval:  &interval,
		Log:       &log,
		LogID:     &logID,
	}
}

// Retrieve property field names for updating the Logflush resource
func GetLogflushPropertyFields() (t []string) {
	return []string{
		"computers",
		"interval",
		"log",
		"log_id",
	}
}
