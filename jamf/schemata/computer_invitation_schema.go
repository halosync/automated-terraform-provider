package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the ComputerInvitation resource defined in the Terraform configuration
func ComputerInvitationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"create_account_if_does_not_exist": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"enrolled_into_site": {
			Type: schema.TypeList, //GoType: ComputerInvitationEnrolledIntoSite
			Elem: &schema.Resource{
				Schema: ComputerInvitationEnrolledIntoSiteSchema(),
			},
			Optional: true,
		},

		"expiration_date": {
			Type:     schema.TypeString,
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

		"hide_account": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"invitation": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		"invitation_status": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"invitation_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"invited_user_uuid": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"keep_existing_site_membership": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"lock_down_ssh": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"multiple_users_allowed": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"site": {
			Type: schema.TypeList, //GoType: SiteObject
			Elem: &schema.Resource{
				Schema: SiteObjectSchema(),
			},
			Optional: true,
		},

		"ssh_password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"ssh_username": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"times_used": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

// Update the underlying ComputerInvitation resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetComputerInvitationResourceData(d *schema.ResourceData, m *models.ComputerInvitation) {
	d.Set("create_account_if_does_not_exist", m.CreateAccountIfDoesNotExist)
	d.Set("enrolled_into_site", SetComputerInvitationEnrolledIntoSiteSubResourceData([]*models.ComputerInvitationEnrolledIntoSite{m.EnrolledIntoSite}))
	d.Set("expiration_date", m.ExpirationDate)
	d.Set("expiration_date_epoch", m.ExpirationDateEpoch)
	d.Set("expiration_date_utc", m.ExpirationDateUtc)
	d.Set("hide_account", m.HideAccount)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("invitation", m.Invitation)
	d.Set("invitation_status", m.InvitationStatus)
	d.Set("invitation_type", m.InvitationType)
	d.Set("invited_user_uuid", m.InvitedUserUUID)
	d.Set("keep_existing_site_membership", m.KeepExistingSiteMembership)
	d.Set("lock_down_ssh", m.LockDownSSH)
	d.Set("multiple_users_allowed", m.MultipleUsersAllowed)
	d.Set("site", SetSiteObjectSubResourceData([]*models.SiteObject{m.Site}))
	d.Set("ssh_password", m.SSHPassword)
	d.Set("ssh_username", m.SSHUsername)
	d.Set("times_used", m.TimesUsed)
}

// Iterate throught and update the ComputerInvitation resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetComputerInvitationSubResourceData(m []*models.ComputerInvitation) (d []*map[string]interface{}) {
	for _, computerInvitation := range m {
		if computerInvitation != nil {
			properties := make(map[string]interface{})
			properties["create_account_if_does_not_exist"] = computerInvitation.CreateAccountIfDoesNotExist
			properties["enrolled_into_site"] = SetComputerInvitationEnrolledIntoSiteSubResourceData([]*models.ComputerInvitationEnrolledIntoSite{computerInvitation.EnrolledIntoSite})
			properties["expiration_date"] = computerInvitation.ExpirationDate
			properties["expiration_date_epoch"] = computerInvitation.ExpirationDateEpoch
			properties["expiration_date_utc"] = computerInvitation.ExpirationDateUtc
			properties["hide_account"] = computerInvitation.HideAccount
			properties["id"] = computerInvitation.ID
			properties["invitation"] = computerInvitation.Invitation
			properties["invitation_status"] = computerInvitation.InvitationStatus
			properties["invitation_type"] = computerInvitation.InvitationType
			properties["invited_user_uuid"] = computerInvitation.InvitedUserUUID
			properties["keep_existing_site_membership"] = computerInvitation.KeepExistingSiteMembership
			properties["lock_down_ssh"] = computerInvitation.LockDownSSH
			properties["multiple_users_allowed"] = computerInvitation.MultipleUsersAllowed
			properties["site"] = SetSiteObjectSubResourceData([]*models.SiteObject{computerInvitation.Site})
			properties["ssh_password"] = computerInvitation.SSHPassword
			properties["ssh_username"] = computerInvitation.SSHUsername
			properties["times_used"] = computerInvitation.TimesUsed
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ComputerInvitation resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ComputerInvitationModel(d *schema.ResourceData) *models.ComputerInvitation {
	createAccountIfDoesNotExist := d.Get("create_account_if_does_not_exist").(bool)
	var enrolledIntoSite *models.ComputerInvitationEnrolledIntoSite = nil //hit complex
	enrolled_into_siteInterface, enrolled_into_siteIsSet := d.GetOk("enrolled_into_site")
	if enrolled_into_siteIsSet {
		enrolled_into_siteMap := enrolled_into_siteInterface.([]interface{})[0].(map[string]interface{})
		enrolledIntoSite = ComputerInvitationEnrolledIntoSiteModel(enrolled_into_siteMap)
	}
	expirationDate := d.Get("expiration_date").(string)
	hideAccount := d.Get("hide_account").(bool)
	id, _ := strconv.Atoi(d.Get("id").(string))
	invitationType := d.Get("invitation_type").(string)
	keepExistingSiteMembership := d.Get("keep_existing_site_membership").(bool)
	lockDownSSH := d.Get("lock_down_ssh").(bool)
	multipleUsersAllowed := d.Get("multiple_users_allowed").(bool)
	var site *models.SiteObject = nil //hit complex
	siteInterface, siteIsSet := d.GetOk("site")
	if siteIsSet {
		siteMap := siteInterface.([]interface{})[0].(map[string]interface{})
		site = SiteObjectModel(siteMap)
	}
	sSHPassword := d.Get("ssh_password").(string)
	sSHUsername := d.Get("ssh_username").(string)

	return &models.ComputerInvitation{
		CreateAccountIfDoesNotExist: &createAccountIfDoesNotExist,
		EnrolledIntoSite:            enrolledIntoSite,
		ExpirationDate:              expirationDate,
		HideAccount:                 &hideAccount,
		ID:                          int32(id),
		InvitationType:              invitationType,
		KeepExistingSiteMembership:  &keepExistingSiteMembership,
		LockDownSSH:                 &lockDownSSH,
		MultipleUsersAllowed:        &multipleUsersAllowed,
		Site:                        site,
		SSHPassword:                 sSHPassword,
		SSHUsername:                 sSHUsername,
	}
}

// Retrieve property field names for updating the ComputerInvitation resource
func GetComputerInvitationPropertyFields() (t []string) {
	return []string{
		"create_account_if_does_not_exist",
		"enrolled_into_site",
		"expiration_date",
		"hide_account",
		"id",
		"invitation_type",
		"keep_existing_site_membership",
		"lock_down_ssh",
		"multiple_users_allowed",
		"site",
		"ssh_password",
		"ssh_username",
	}
}
