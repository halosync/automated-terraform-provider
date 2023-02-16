package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DistributionPointPost resource defined in the Terraform configuration
func DistributionPointPostSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection_type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"context": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"enable_load_balancing": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"failover_point": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_downloads_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"http_password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_url": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"http_username": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"ip_address": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"is_master": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"no_authentication_required": {
			Type:     schema.TypeBool,
			Default:  true,
			Optional: true,
		},

		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"port": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"protocol": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"read_only_password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"read_only_username": {
			Type:     schema.TypeString,
			Required: true,
		},

		"read_write_password": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"read_write_username": {
			Type:     schema.TypeString,
			Required: true,
		},

		"share_name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"share_port": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"ssh_username": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"username_password_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"workgroup_or_domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Update the underlying DistributionPointPost resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDistributionPointPostResourceData(d *schema.ResourceData, m *models.DistributionPointPost) {
	d.Set("connection_type", m.ConnectionType)
	d.Set("context", m.Context)
	d.Set("enable_load_balancing", m.EnableLoadBalancing)
	d.Set("failover_point", m.FailoverPoint)
	d.Set("http_downloads_enabled", m.HTTPDownloadsEnabled)
	d.Set("http_password", m.HTTPPassword)
	d.Set("http_url", m.HTTPURL)
	d.Set("http_username", m.HTTPUsername)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("ip_address", m.IPAddress)
	d.Set("is_master", m.IsMaster)
	d.Set("name", m.Name)
	d.Set("no_authentication_required", m.NoAuthenticationRequired)
	d.Set("password", m.Password)
	d.Set("port", m.Port)
	d.Set("protocol", m.Protocol)
	d.Set("read_only_password", m.ReadOnlyPassword)
	d.Set("read_only_username", m.ReadOnlyUsername)
	d.Set("read_write_password", m.ReadWritePassword)
	d.Set("read_write_username", m.ReadWriteUsername)
	d.Set("share_name", m.ShareName)
	d.Set("share_port", m.SharePort)
	d.Set("ssh_username", m.SSHUsername)
	d.Set("username_password_required", m.UsernamePasswordRequired)
	d.Set("workgroup_or_domain", m.WorkgroupOrDomain)
}

// Iterate throught and update the DistributionPointPost resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDistributionPointPostSubResourceData(m []*models.DistributionPointPost) (d []*map[string]interface{}) {
	for _, distributionPointPost := range m {
		if distributionPointPost != nil {
			properties := make(map[string]interface{})
			properties["connection_type"] = distributionPointPost.ConnectionType
			properties["context"] = distributionPointPost.Context
			properties["enable_load_balancing"] = distributionPointPost.EnableLoadBalancing
			properties["failover_point"] = distributionPointPost.FailoverPoint
			properties["http_downloads_enabled"] = distributionPointPost.HTTPDownloadsEnabled
			properties["http_password"] = distributionPointPost.HTTPPassword
			properties["http_url"] = distributionPointPost.HTTPURL
			properties["http_username"] = distributionPointPost.HTTPUsername
			properties["id"] = distributionPointPost.ID
			properties["ip_address"] = distributionPointPost.IPAddress
			properties["is_master"] = distributionPointPost.IsMaster
			properties["name"] = distributionPointPost.Name
			properties["no_authentication_required"] = distributionPointPost.NoAuthenticationRequired
			properties["password"] = distributionPointPost.Password
			properties["port"] = distributionPointPost.Port
			properties["protocol"] = distributionPointPost.Protocol
			properties["read_only_password"] = distributionPointPost.ReadOnlyPassword
			properties["read_only_username"] = distributionPointPost.ReadOnlyUsername
			properties["read_write_password"] = distributionPointPost.ReadWritePassword
			properties["read_write_username"] = distributionPointPost.ReadWriteUsername
			properties["share_name"] = distributionPointPost.ShareName
			properties["share_port"] = distributionPointPost.SharePort
			properties["ssh_username"] = distributionPointPost.SSHUsername
			properties["username_password_required"] = distributionPointPost.UsernamePasswordRequired
			properties["workgroup_or_domain"] = distributionPointPost.WorkgroupOrDomain
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DistributionPointPost resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DistributionPointPostModel(d *schema.ResourceData) *models.DistributionPointPost {
	connectionType := d.Get("connection_type").(string)
	context := d.Get("context").(string)
	enableLoadBalancing := d.Get("enable_load_balancing").(bool)
	failoverPoint := d.Get("failover_point").(string)
	hTTPDownloadsEnabled := d.Get("http_downloads_enabled").(bool)
	hTTPPassword := d.Get("http_password").(string)
	hTTPURL := d.Get("http_url").(string)
	hTTPUsername := d.Get("http_username").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	iPAddress := d.Get("ip_address").(string)
	isMaster := d.Get("is_master").(bool)
	name := d.Get("name").(string)
	noAuthenticationRequired := d.Get("no_authentication_required").(bool)
	password := d.Get("password").(string)
	port := int64(d.Get("port").(int))
	protocol := d.Get("protocol").(string)
	readOnlyPassword := d.Get("read_only_password").(string)
	readOnlyUsername := d.Get("read_only_username").(string)
	readWritePassword := d.Get("read_write_password").(string)
	readWriteUsername := d.Get("read_write_username").(string)
	shareName := d.Get("share_name").(string)
	sharePort := int64(d.Get("share_port").(int))
	sSHUsername := d.Get("ssh_username").(string)
	usernamePasswordRequired := d.Get("username_password_required").(bool)
	workgroupOrDomain := d.Get("workgroup_or_domain").(string)

	return &models.DistributionPointPost{
		ConnectionType:           connectionType,
		Context:                  context,
		EnableLoadBalancing:      &enableLoadBalancing,
		FailoverPoint:            failoverPoint,
		HTTPDownloadsEnabled:     hTTPDownloadsEnabled,
		HTTPPassword:             hTTPPassword,
		HTTPURL:                  hTTPURL,
		HTTPUsername:             hTTPUsername,
		ID:                       int32(id),
		IPAddress:                iPAddress,
		IsMaster:                 isMaster,
		Name:                     &name,
		NoAuthenticationRequired: &noAuthenticationRequired,
		Password:                 password,
		Port:                     port,
		Protocol:                 protocol,
		ReadOnlyPassword:         readOnlyPassword,
		ReadOnlyUsername:         &readOnlyUsername,
		ReadWritePassword:        readWritePassword,
		ReadWriteUsername:        &readWriteUsername,
		ShareName:                &shareName,
		SharePort:                sharePort,
		SSHUsername:              sSHUsername,
		UsernamePasswordRequired: &usernamePasswordRequired,
		WorkgroupOrDomain:        workgroupOrDomain,
	}
}

// Retrieve property field names for updating the DistributionPointPost resource
func GetDistributionPointPostPropertyFields() (t []string) {
	return []string{
		"connection_type",
		"context",
		"enable_load_balancing",
		"failover_point",
		"http_downloads_enabled",
		"http_password",
		"http_url",
		"http_username",
		"id",
		"ip_address",
		"is_master",
		"name",
		"no_authentication_required",
		"password",
		"port",
		"protocol",
		"read_only_password",
		"read_only_username",
		"read_write_password",
		"read_write_username",
		"share_name",
		"share_port",
		"ssh_username",
		"username_password_required",
		"workgroup_or_domain",
	}
}
