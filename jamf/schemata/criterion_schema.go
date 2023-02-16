package schemata

import (
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the Criterion resource defined in the Terraform configuration
func CriterionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"and_or": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"closing_paren": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"opening_paren": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"search_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying Criterion resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCriterionResourceData(d *schema.ResourceData, m *models.Criterion) {
	d.Set("and_or", m.AndOr)
	d.Set("closing_paren", m.ClosingParen)
	d.Set("name", m.Name)
	d.Set("opening_paren", m.OpeningParen)
	d.Set("priority", m.Priority)
	d.Set("search_type", m.SearchType)
	d.Set("value", m.Value)
}

// Iterate throught and update the Criterion resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCriterionSubResourceData(m []*models.Criterion) (d []*map[string]interface{}) {
	for _, criterion := range m {
		if criterion != nil {
			properties := make(map[string]interface{})
			properties["and_or"] = criterion.AndOr
			properties["closing_paren"] = criterion.ClosingParen
			properties["name"] = criterion.Name
			properties["opening_paren"] = criterion.OpeningParen
			properties["priority"] = criterion.Priority
			properties["search_type"] = criterion.SearchType
			properties["value"] = criterion.Value
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Criterion resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CriterionModel(d *schema.ResourceData) *models.Criterion {
	andOr := d.Get("and_or").(string)
	closingParen := d.Get("closing_paren").(bool)
	name := d.Get("name").(string)
	openingParen := d.Get("opening_paren").(bool)
	priority := int64(d.Get("priority").(int))
	searchType := d.Get("search_type").(string)
	value := d.Get("value").(string)

	return &models.Criterion{
		AndOr:        andOr,
		ClosingParen: &closingParen,
		Name:         name,
		OpeningParen: &openingParen,
		Priority:     priority,
		SearchType:   searchType,
		Value:        value,
	}
}

// Retrieve property field names for updating the Criterion resource
func GetCriterionPropertyFields() (t []string) {
	return []string{
		"and_or",
		"closing_paren",
		"name",
		"opening_paren",
		"priority",
		"search_type",
		"value",
	}
}
