package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the MobileDeviceInvitation resource defined in the Terraform configuration
func MobileDeviceInvitationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Type: schema.TypeList, //GoType: MobileDeviceInvitationEnrolledIntoSite
			Elem: &schema.Resource{
				Schema: MobileDeviceInvitationEnrolledIntoSiteSchema(),
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

		"login_required": {
			Type:     schema.TypeBool,
			Default:  true,
			Optional: true,
		},

		"message": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"multiple_uses_allowed": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"reply_to": {
			Type:     schema.TypeString,
			Computed: true,
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

// Update the underlying MobileDeviceInvitation resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMobileDeviceInvitationResourceData(d *schema.ResourceData, m *models.MobileDeviceInvitation) {
	d.Set("date_sent", m.DateSent)
	d.Set("date_sent_epoch", m.DateSentEpoch)
	d.Set("date_sent_utc", m.DateSentUtc)
	d.Set("enrolled_into_site", SetMobileDeviceInvitationEnrolledIntoSiteSubResourceData([]*models.MobileDeviceInvitationEnrolledIntoSite{m.EnrolledIntoSite}))
	d.Set("expiration_date", m.ExpirationDate)
	d.Set("expiration_date_epoch", m.ExpirationDateEpoch)
	d.Set("expiration_date_utc", m.ExpirationDateUtc)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("invitation", m.Invitation)
	d.Set("invitation_type", m.InvitationType)
	d.Set("keep_existing_site_membership", m.KeepExistingSiteMembership)
	d.Set("last_action", m.LastAction)
	d.Set("login_required", m.LoginRequired)
	d.Set("message", m.Message)
	d.Set("multiple_uses_allowed", m.MultipleUsesAllowed)
	d.Set("reply_to", m.ReplyTo)
	d.Set("sent_from", m.SentFrom)
	d.Set("sent_to", m.SentTo)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("subject", m.Subject)
	d.Set("target_ios", m.TargetIos)
	d.Set("username", m.Username)
}

// Iterate throught and update the MobileDeviceInvitation resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMobileDeviceInvitationSubResourceData(m []*models.MobileDeviceInvitation) (d []*map[string]interface{}) {
	for _, mobileDeviceInvitation := range m {
		if mobileDeviceInvitation != nil {
			properties := make(map[string]interface{})
			properties["date_sent"] = mobileDeviceInvitation.DateSent
			properties["date_sent_epoch"] = mobileDeviceInvitation.DateSentEpoch
			properties["date_sent_utc"] = mobileDeviceInvitation.DateSentUtc
			properties["enrolled_into_site"] = SetMobileDeviceInvitationEnrolledIntoSiteSubResourceData([]*models.MobileDeviceInvitationEnrolledIntoSite{mobileDeviceInvitation.EnrolledIntoSite})
			properties["expiration_date"] = mobileDeviceInvitation.ExpirationDate
			properties["expiration_date_epoch"] = mobileDeviceInvitation.ExpirationDateEpoch
			properties["expiration_date_utc"] = mobileDeviceInvitation.ExpirationDateUtc
			properties["id"] = mobileDeviceInvitation.ID
			properties["invitation"] = mobileDeviceInvitation.Invitation
			properties["invitation_type"] = mobileDeviceInvitation.InvitationType
			properties["keep_existing_site_membership"] = mobileDeviceInvitation.KeepExistingSiteMembership
			properties["last_action"] = mobileDeviceInvitation.LastAction
			properties["login_required"] = mobileDeviceInvitation.LoginRequired
			properties["message"] = mobileDeviceInvitation.Message
			properties["multiple_uses_allowed"] = mobileDeviceInvitation.MultipleUsesAllowed
			properties["reply_to"] = mobileDeviceInvitation.ReplyTo
			properties["sent_from"] = mobileDeviceInvitation.SentFrom
			properties["sent_to"] = mobileDeviceInvitation.SentTo
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{mobileDeviceInvitation.Site})
			properties["subject"] = mobileDeviceInvitation.Subject
			properties["target_ios"] = mobileDeviceInvitation.TargetIos
			properties["username"] = mobileDeviceInvitation.Username
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MobileDeviceInvitation resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MobileDeviceInvitationModel(d *schema.ResourceData) *models.MobileDeviceInvitation {
	var enrolledIntoSite *models.MobileDeviceInvitationEnrolledIntoSite = nil //hit complex
	enrolled_into_siteInterface, enrolled_into_siteIsSet := d.GetOk("enrolled_into_site")
	if enrolled_into_siteIsSet {
		enrolled_into_siteMap := enrolled_into_siteInterface.([]interface{})[0].(map[string]interface{})
		enrolledIntoSite = MobileDeviceInvitationEnrolledIntoSiteModel(enrolled_into_siteMap)
	}
	expirationDate := d.Get("expiration_date").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	invitationType := d.Get("invitation_type").(string)
	keepExistingSiteMembership := d.Get("keep_existing_site_membership").(bool)
	loginRequired := d.Get("login_required").(bool)
	message := d.Get("message").(string)
	multipleUsesAllowed := d.Get("multiple_uses_allowed").(bool)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	targetIos := d.Get("target_ios").(string)
	username := d.Get("username").(string)

	return &models.MobileDeviceInvitation{
		EnrolledIntoSite:           enrolledIntoSite,
		ExpirationDate:             &expirationDate,
		ID:                         int32(id),
		InvitationType:             &invitationType,
		KeepExistingSiteMembership: &keepExistingSiteMembership,
		LoginRequired:              &loginRequired,
		Message:                    message,
		MultipleUsesAllowed:        &multipleUsesAllowed,
		Site:                       site,
		TargetIos:                  targetIos,
		Username:                   username,
	}
}

// Retrieve property field names for updating the MobileDeviceInvitation resource
func GetMobileDeviceInvitationPropertyFields() (t []string) {
	return []string{
		"enrolled_into_site",
		"expiration_date",
		"id",
		"invitation_type",
		"keep_existing_site_membership",
		"login_required",
		"message",
		"multiple_uses_allowed",
		"site",
		"target_ios",
		"username",
	}
}
