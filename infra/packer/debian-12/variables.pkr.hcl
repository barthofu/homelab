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

variable "username" {
    type = string
    description = "The username to use for SSH connections"
}

variable "password" {
    type = string
    sensitive = true
    description = "The password to use for SSH connections"
}

variable "root_password" {
    type = string
    sensitive = true
    description = "The root password to use for SSH connections"
}

variable "ssh_public_key" {
    type = string
    description = "The SSH key to use for SSH connections"
}

variable "iso_url" {
  type        = string
  description = "URL to an ISO file to upload to Proxmox, and then boot from."
  default     = "https:///debian-cd/current/amd64/iso-cd/debian-12.5.0-amd64-netinst.iso"
}

variable "iso_file" {
  type        = string
  description = "Filename of the ISO file to boot from."
  default     = null
}

variable "iso_checksum" {
  type        = string
  description = "Checksum of the ISO file."
  default     = null
}