# Resource definiation for the VM Template
source "proxmox-iso" "alpine-3" {

  # Proxmox Connection
  proxmox_url              = "${var.proxmox_api_url}"
  username                 = "${var.proxmox_api_token_id}"
  token                    = "${var.proxmox_api_token_secret}"
  insecure_skip_tls_verify = true

  # VM Template
  node                 = "homelab"
  vm_id                = 1001
  template_name        = "alpine-3-base"
  template_description = "Alpine 3.19 base image"

  # ISO
  iso_file         = "local:iso/alpine-standard-3.19.1-x86_64.iso"
  iso_storage_pool = "local"
  iso_checksum     = "63e62f5a52cfe73a6cb137ecbb111b7d48356862a1dfe50d8fdd977d727da192"
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
  cloud_init = true

  # Boot
  boot = null
  boot_command = [
    "root<enter><wait>",
    "ifconfig eth0 up && udhcpc -i eth0<enter><wait5>", # Start networking with DHCP
    "wget ${local.http_url}/answers<enter><wait>",      #Replace CR if file was generated on Windows machine
    "sed -i 's/\\r$//g' $PWD/answers<enter><wait>",
    "USERANSERFILE=1 setup-alpine -f $PWD/answers<enter><wait10>", # Run alpine installer
    "${var.root_password}<enter><wait>",
    "${var.root_password}<enter><wait>",
    "no<enter><wait10>",
    "y<enter><wait20>",
    "reboot<enter>",
    "<wait30>",
    "root<enter><wait>",
    "${var.root_password}<enter><wait>",
    "wget ${local.http_url}/setup.sh<enter><wait>",
    "chmod +x $PWD/setup.sh<enter><wait>",
    "sed -i 's/\\r$//g' $PWD/setup.sh<enter><wait>",
    "$PWD/setup.sh<enter><wait>",
    "wget ${local.http_url}/provision.sh<enter><wait>",
    "chmod +x $PWD/provision.sh<enter><wait>",
    "sed -i 's/\\r$//g' $PWD/provision.sh<enter><wait>",
    "$PWD/provision.sh<enter><wait>",
    "rm -f $PWD/setup.sh<enter><wait>",
    "rm -f $PWD/provision.sh<enter><wait>",
    "setup-keymap fr fr<enter><wait>",
  ]
  boot_wait = "10s"

  # HTTP
  http_directory    = "./http"
  http_bind_address = null
  http_interface    = null
  http_port_min     = var.http_server_port
  http_port_max     = var.http_server_port

  # SSH
  ssh_handshake_attempts    = 100
  ssh_username              = "${var.username}"
  ssh_password              = "${var.password}"
  ssh_timeout               = "15m"
  ssh_clear_authorized_keys = true
}
