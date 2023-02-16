package schemata

import (
	"strconv"
	"tf-provider-example/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Schema mapping representing the DistributionPoint resource defined in the Terraform configuration
func DistributionPointSchema() map[string]*schema.Schema {
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

		"failover_point_url": {
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

		"local_path": {
			Type:     schema.TypeString,
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

// Update the underlying DistributionPoint resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDistributionPointResourceData(d *schema.ResourceData, m *models.DistributionPoint) {
	d.Set("connection_type", m.ConnectionType)
	d.Set("context", m.Context)
	d.Set("enable_load_balancing", m.EnableLoadBalancing)
	d.Set("failover_point", m.FailoverPoint)
	d.Set("failover_point_url", m.FailoverPointURL)
	d.Set("http_downloads_enabled", m.HTTPDownloadsEnabled)
	d.Set("http_password", m.HTTPPassword)
	d.Set("http_url", m.HTTPURL)
	d.Set("http_username", m.HTTPUsername)
	d.Set("id", strconv.Itoa(int(m.ID)))
	d.Set("ip_address", m.IPAddress)
	d.Set("is_master", m.IsMaster)
	d.Set("local_path", m.LocalPath)
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

// Iterate throught and update the DistributionPoint resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDistributionPointSubResourceData(m []*models.DistributionPoint) (d []*map[string]interface{}) {
	for _, distributionPoint := range m {
		if distributionPoint != nil {
			properties := make(map[string]interface{})
			properties["connection_type"] = distributionPoint.ConnectionType
			properties["context"] = distributionPoint.Context
			properties["enable_load_balancing"] = distributionPoint.EnableLoadBalancing
			properties["failover_point"] = distributionPoint.FailoverPoint
			properties["failover_point_url"] = distributionPoint.FailoverPointURL
			properties["http_downloads_enabled"] = distributionPoint.HTTPDownloadsEnabled
			properties["http_password"] = distributionPoint.HTTPPassword
			properties["http_url"] = distributionPoint.HTTPURL
			properties["http_username"] = distributionPoint.HTTPUsername
			properties["id"] = distributionPoint.ID
			properties["ip_address"] = distributionPoint.IPAddress
			properties["is_master"] = distributionPoint.IsMaster
			properties["local_path"] = distributionPoint.LocalPath
			properties["name"] = distributionPoint.Name
			properties["no_authentication_required"] = distributionPoint.NoAuthenticationRequired
			properties["password"] = distributionPoint.Password
			properties["port"] = distributionPoint.Port
			properties["protocol"] = distributionPoint.Protocol
			properties["read_only_password"] = distributionPoint.ReadOnlyPassword
			properties["read_only_username"] = distributionPoint.ReadOnlyUsername
			properties["read_write_password"] = distributionPoint.ReadWritePassword
			properties["read_write_username"] = distributionPoint.ReadWriteUsername
			properties["share_name"] = distributionPoint.ShareName
			properties["share_port"] = distributionPoint.SharePort
			properties["ssh_username"] = distributionPoint.SSHUsername
			properties["username_password_required"] = distributionPoint.UsernamePasswordRequired
			properties["workgroup_or_domain"] = distributionPoint.WorkgroupOrDomain
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DistributionPoint resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DistributionPointModel(d *schema.ResourceData) *models.DistributionPoint {
	connectionType := d.Get("connection_type").(string)
	context := d.Get("context").(string)
	enableLoadBalancing := d.Get("enable_load_balancing").(bool)
	failoverPoint := d.Get("failover_point").(string)
	failoverPointURL := d.Get("failover_point_url").(string)
	hTTPDownloadsEnabled := d.Get("http_downloads_enabled").(bool)
	hTTPPassword := d.Get("http_password").(string)
	hTTPURL := d.Get("http_url").(string)
	hTTPUsername := d.Get("http_username").(string)
	id, _ := strconv.Atoi(d.Get("id").(string))
	iPAddress := d.Get("ip_address").(string)
	isMaster := d.Get("is_master").(bool)
	localPath := d.Get("local_path").(string)
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

	return &models.DistributionPoint{
		ConnectionType:           connectionType,
		Context:                  context,
		EnableLoadBalancing:      &enableLoadBalancing,
		FailoverPoint:            failoverPoint,
		FailoverPointURL:         failoverPointURL,
		HTTPDownloadsEnabled:     &hTTPDownloadsEnabled,
		HTTPPassword:             hTTPPassword,
		HTTPURL:                  hTTPURL,
		HTTPUsername:             hTTPUsername,
		ID:                       int32(id),
		IPAddress:                iPAddress,
		IsMaster:                 &isMaster,
		LocalPath:                localPath,
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

// Retrieve property field names for updating the DistributionPoint resource
func GetDistributionPointPropertyFields() (t []string) {
	return []string{
		"connection_type",
		"context",
		"enable_load_balancing",
		"failover_point",
		"failover_point_url",
		"http_downloads_enabled",
		"http_password",
		"http_url",
		"http_username",
		"id",
		"ip_address",
		"is_master",
		"local_path",
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
