#######################################
# VPC module variables
#######################################

variable "cluster_id" {
  type = string
}

variable "dedicated_hosts_master" {
  type = list(map(string))
  default = []
}

variable "dedicated_hosts_worker" {
  type = list(map(string))
  default = []
}

variable "public_endpoints" {
  type = bool
}

variable "region" {
  type = string
}

variable "resource_group_id" {
  type = string
}

variable "tags" {
  type = list(string)
}

variable "zone_list" {
  type = list(string)
}