terraform {
  required_version = ">= 0.12.00"
  required_providers {
    logicmonitor = {
      source  = "logicmonitor.com/com/logicmonitor"
      version = "0.2"
    }
  }
}


variable "api_id" {}
variable "api_key" {}
variable "company" {}

provider "logicmonitor" {
  api_id  = var.api_id
  api_key = var.api_key
  company = var.company
}

resource "logicmonitor_device" "my_device" {
  preferred_collector_id = 1
  custom_properties = [
    {
      name  = "addr"
      value = "127.0.0.1"
    },
    {
      name  = "host"
      value = "localhost"
    }
  ]
  description      = "This is a Cisco Router"
  device_type      = 0
  disable_alerting = true
  display_name     = "Cisco Router"
  enable_netflow   = false
  link             = "www.ciscorouter.com"
  name             = "collector.host"
}

data "logicmonitor_device" "my_devices" {
  filter = "displayName~\"Cisco Router\""
  // make sure that the device is created before we try to query for it
  depends_on = [
    logicmonitor_device.my_device
  ]
}

output "devices" {
  description = "devices list"
  value       = data.logicmonitor_device.my_devices
}
