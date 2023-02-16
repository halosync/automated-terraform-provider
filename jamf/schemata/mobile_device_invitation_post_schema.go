package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceInvitationPost resource defined in the Terraform configuration
func MobileDeviceInvitationPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_multiple_uses": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"date_sent": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"date_sent_epoch": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"date_sent_utc": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"enrolled_into_site": {
			Type: schema.TypeList, //GoType: MobileDeviceInvitationPostEnrolledIntoSite
			Elem: &schema.Resource{
				Schema: MobileDeviceInvitationPostEnrolledIntoSiteSchema(),
			},
			Optional: true,
		},

		"expiration_date": {
			Type:     schema.TypeString,
			Default:  "Unlimited",
			Optional: true,
		},

		"expiration_date_epoch": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"expiration_date_utc": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"invitation": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"invitation_type": {
			Type:     schema.TypeString,
			Default:  "USER_INITIATED_URL",
			Optional: true,
		},

		"keep_existing_site_membership": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"last_action": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"message": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"reply_to": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"require_login": {
			Type:     schema.TypeBool,
			Default:  true,
			Optional: true,
		},

		"sent_from": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"sent_to": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: SiteObject
			Elem: &schema.Resource{
				Schema: SiteObjectSchema(),
			},
			Optional: true,
		},

		"subject": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"target_ios": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying MobileDeviceInvitationPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceInvitationPostResourceData(d *schema.ResourceData, m *models.MobileDeviceInvitationPost) {
	d.Set("allow_multiple_uses", m.AllowMultipleUses)
	d.Set("date_sent", m.DateSent)
	d.Set("date_sent_epoch", m.DateSentEpoch)
	d.Set("date_sent_utc", m.DateSentUtc)
	d.Set("enrolled_into_site", SetMobileDeviceInvitationPostEnrolledIntoSiteSubResourceData([]*models.MobileDeviceInvitationPostEnrolledIntoSite{m.EnrolledIntoSite}))
	d.Set("expiration_date", m.ExpirationDate)
	d.Set("expiration_date_epoch", m.ExpirationDateEpoch)
	d.Set("expiration_date_utc", m.ExpirationDateUtc)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("invitation", m.Invitation)
	d.Set("invitation_type", m.InvitationType)
	d.Set("keep_existing_site_membership", m.KeepExistingSiteMembership)
	d.Set("last_action", m.LastAction)
	d.Set("message", m.Message)
	d.Set("reply_to", m.ReplyTo)
	d.Set("require_login", m.RequireLogin)
	d.Set("sent_from", m.SentFrom)
	d.Set("sent_to", m.SentTo)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("subject", m.Subject)
	d.Set("target_ios", m.TargetIos)
	d.Set("username", m.Username)
}

// Iterate throught and update the MobileDeviceInvitationPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceInvitationPostSubResourceData(m []*models.MobileDeviceInvitationPost) (d []*map[string]interface{}) {
	for _, mobileDeviceInvitationPost := range m {
		if mobileDeviceInvitationPost != nil {
			properties := make(map[string]interface{})
			properties["allow_multiple_uses"] = mobileDeviceInvitationPost.AllowMultipleUses
			properties["date_sent"] = mobileDeviceInvitationPost.DateSent
			properties["date_sent_epoch"] = mobileDeviceInvitationPost.DateSentEpoch
			properties["date_sent_utc"] = mobileDeviceInvitationPost.DateSentUtc
			properties["enrolled_into_site"] = SetMobileDeviceInvitationPostEnrolledIntoSiteSubResourceData([]*models.MobileDeviceInvitationPostEnrolledIntoSite{mobileDeviceInvitationPost.EnrolledIntoSite})
			properties["expiration_date"] = mobileDeviceInvitationPost.ExpirationDate
			properties["expiration_date_epoch"] = mobileDeviceInvitationPost.ExpirationDateEpoch
			properties["expiration_date_utc"] = mobileDeviceInvitationPost.ExpirationDateUtc
			properties["id"] = mobileDeviceInvitationPost.ID
			properties["invitation"] = mobileDeviceInvitationPost.Invitation
			properties["invitation_type"] = mobileDeviceInvitationPost.InvitationType
			properties["keep_existing_site_membership"] = mobileDeviceInvitationPost.KeepExistingSiteMembership
			properties["last_action"] = mobileDeviceInvitationPost.LastAction
			properties["message"] = mobileDeviceInvitationPost.Message
			properties["reply_to"] = mobileDeviceInvitationPost.ReplyTo
			properties["require_login"] = mobileDeviceInvitationPost.RequireLogin
			properties["sent_from"] = mobileDeviceInvitationPost.SentFrom
			properties["sent_to"] = mobileDeviceInvitationPost.SentTo
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{mobileDeviceInvitationPost.Site})
			properties["subject"] = mobileDeviceInvitationPost.Subject
			properties["target_ios"] = mobileDeviceInvitationPost.TargetIos
			properties["username"] = mobileDeviceInvitationPost.Username
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceInvitationPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceInvitationPostModel(d *schema.ResourceData) *models.MobileDeviceInvitationPost {
	allowMultipleUses := d.Get("allow_multiple_uses").(bool)
	var enrolledIntoSite *models.MobileDeviceInvitationPostEnrolledIntoSite = nil //hit complex
	enrolled_into_siteInterface, enrolled_into_siteIsSet := d.GetOk("enrolled_into_site")
	if enrolled_into_siteIsSet {
		enrolled_into_siteMap := enrolled_into_siteInterface.([]interface{})[0].(map[string]interface{})
		enrolledIntoSite = MobileDeviceInvitationPostEnrolledIntoSiteModel(enrolled_into_siteMap)
	}
	expirationDate := d.Get("expiration_date").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	invitationType := d.Get("invitation_type").(string)
	keepExistingSiteMembership := d.Get("keep_existing_site_membership").(bool)
	message := d.Get("message").(string)
	requireLogin := d.Get("require_login").(bool)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	targetIos := d.Get("target_ios").(string)
	username := d.Get("username").(string)

	return &models.MobileDeviceInvitationPost{
		AllowMultipleUses:          &allowMultipleUses,
		EnrolledIntoSite:           enrolledIntoSite,
		ExpirationDate:             &expirationDate,
		ID:                         int32(id),
		InvitationType:             &invitationType,
		KeepExistingSiteMembership: &keepExistingSiteMembership,
		Message:                    message,
		RequireLogin:               &requireLogin,
		Site:                       site,
		TargetIos:                  targetIos,
		Username:                   username,
	}
}

// Retrieve property field names for updating the MobileDeviceInvitationPost resource
func GetMobileDeviceInvitationPostPropertyFields() (t []string) {
	return []string{
		"allow_multiple_uses",
		"enrolled_into_site",
		"expiration_date",
		"id",
		"invitation_type",
		"keep_existing_site_membership",
		"message",
		"require_login",
		"site",
		"target_ios",
		"username",
	}
}
