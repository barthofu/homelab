source "proxmox-iso" "k3os" {

  # Proxmox Connection
  proxmox_url              = "${var.proxmox_api_url}"
  username                 = "${var.proxmox_api_token_id}"
  token                    = "${var.proxmox_api_token_secret}"
  insecure_skip_tls_verify = true

  # VM
  node                 = "homelab"
  vm_id                = 1001
  vm_name              = "k3os-${var.k3os_version}-base"
  template_description = "k3os ${var.k3os_version}"

  # ISO
  iso_file         = "local:iso/k3os-amd64.iso"
  iso_checksum     = null
  iso_storage_pool = "local"
  unmount_iso      = true

  # Hardware
  os              = "l26"
  cores           = 2
  memory          = "2048"
  scsi_controller = "virtio-scsi-pci"

  # Storage
  disks {
    disk_size    = "8G"
    storage_pool = "local-lvm"
    type         = "virtio"
  }

  # Networking
  network_adapters {
    model    = "virtio"
    bridge   = "vmbr0"
    firewall = "false"
  }

  # Cloud-Init
  cloud_init  = false
 
  # Boot
  # boot_command = concat(
  #   local.boot_command_pre, 
  #   local.boot_command_args, 
  #   local.boot_command_args_proxmox, 
  #   local.boot_command_post
  # )
  boot_command = [
    "rancher",
    "<enter>",
    "sudo k3os install",
    "<enter>",
    "1",
    "<enter>",
    "Y",
    "<enter>",
    "http://{{ .HTTPIP }}:{{ .HTTPPort }}/cloud.yml",
    "<enter>",
    "Y",
    "<enter>"
  ],
  boot_wait    = "90s"
  boot         = "order=virtio0;ide2"

  # HTTP
  http_directory    = "./http"
  http_bind_address = null
  http_interface    = null
  http_port_min     = var.http_server_port
  http_port_max     = var.http_server_port

  # Misc
  qemu_agent = true
}

  # template_name        = "k3os-${var.k3os_version}"
  # template_description = <<EOF
  #   k3os ${var.k3os_version}
  #   generated on ${timestamp()}"
  #   git ref: ${var.git_ref}
  #   git sha: ${var.git_sha}
  # EOF

 