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
    type = string
    sensitive = true
    description = "The root password to use for SSH connections"
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

variable "ssh_public_key" {
    type = string
    description = "The SSH key to use for SSH connections"
}