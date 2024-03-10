# Resource definiation for the VM Template
source "proxmox-iso" "debian-12" {

  # Proxmox Connection
  proxmox_url              = "${var.proxmox_api_url}"
  username                 = "${var.proxmox_api_token_id}"
  token                    = "${var.proxmox_api_token_secret}"
  insecure_skip_tls_verify = true

  # VM
  node                 = "homelab"
  vm_id                = 1000
  vm_name              = "debian-12-base"
  template_description = "Debian 12 base image"

  # ISO
  iso_file         = "local:iso/debian-12.5.0-amd64-netinst.iso"
  iso_checksum     = "33c08e56c83d13007e4a5511b9bf2c4926c4aa12fd5dd56d493c0653aecbab380988c5bf1671dbaea75c582827797d98c4a611f7fb2b131fbde2c677d5258ec9"
  iso_storage_pool = "local"
  unmount_iso      = true

  # Hardware
  os              = "l26"
  qemu_agent      = true
  cores           = "1"
  sockets         = "1"
  memory          = "1024"
  scsi_controller = "virtio-scsi-pci"

  # Storage
  disks {
    disk_size    = "8G"
    storage_pool = "local-lvm"
    type         = "virtio"
  }

  # Network
  network_adapters {
    model    = "virtio"
    bridge   = "vmbr0"
    firewall = "false"
  }

  # Cloud-Init
  cloud_init              = true
  cloud_init_storage_pool = "local-lvm"

  # Boot
  boot = null
  boot_command = [
    "<esc><wait>",
    "auto console-keymaps-at/keymap=fr console-setup/ask_detect=false debconf/frontend=noninteractive fb=false preseed/url=${local.http_url}/preseed.cfg",
    "<wait><enter>"
  ]
  boot_wait = "6s"

  # HTTP
  http_directory    = "./http"
  http_bind_address = null
  http_interface    = null
  http_port_min     = var.http_server_port
  http_port_max     = var.http_server_port

  # SSH
  ssh_username         = var.username
  ssh_password         = var.password
  ssh_private_key_file = null
  ssh_timeout          = "15m"
}