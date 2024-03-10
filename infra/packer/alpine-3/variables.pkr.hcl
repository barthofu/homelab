variable "proxmox_api_url" {
    type = string
    description = "The Proxmox API url to connect to"
}

variable "proxmox_api_token_id" {
    type = string
    description = "The Proxmox API token ID"
}

variable "proxmox_api_token_secret" {
    type = string
    sensitive = true
    description = "The Proxmox API token secret"
}

variable "http_server_host" {
    type = string
    description = "The host to bind the HTTP server to"
}

variable "http_server_port" {
    type = number
    description = "The port to bind the HTTP server to"
}

variable "root_password" {
  type        = string
  description = "root password to use during the setup process."
  sensitive   = true
}

variable "username" {
    type = string
    description = "The username to use for SSH connections"
}

variable "password" {
    type = string
    sensitive = true
    description = "The password to use for SSH connections"
}

##### Optional Variables #####

variable "dns_servers" {
  type        = list(string)
  description = "Sets the DNS servers during the setup-alpine install."
  default     = []

  validation {
    condition     = var.dns_servers != null
    error_message = "The DNS server list must not be null. An empty list is allowed."
  }
}