# Resource Definiation for the VM Template
source "proxmox-iso" "debian-12" {
 
    # Proxmox Connection Settings
    proxmox_url = "${var.proxmox_api_url}"
    username = "${var.proxmox_api_token_id}"
    token = "${var.proxmox_api_token_secret}"
    insecure_skip_tls_verify = true
    
    # VM General Settings
    node = "homelab"
    vm_id = 1000
    vm_name = "debian-12-base"
    template_description = "Debian 12 base image"

    # VM OS Settings
    # (Option 1) Local ISO File
    iso_file = "local:iso/debian-12.5.0-amd64-netinst.iso"
    iso_checksum = "33c08e56c83d13007e4a5511b9bf2c4926c4aa12fd5dd56d493c0653aecbab380988c5bf1671dbaea75c582827797d98c4a611f7fb2b131fbde2c677d5258ec9"
    # - or -
    # (Option 2) Download ISO
    # iso_url = "https://releases.ubuntu.com/22.04/ubuntu-22.04-live-server-amd64.iso"
    # iso_checksum = "84aeaf7823c8c61baa0ae862d0a06b03409394800000b3235854a6b38eb4856f"
    iso_storage_pool = "local"
    unmount_iso = true

    os         = "l26"
    qemu_agent = true
    cores      = "1"
    sockets    = "1"
    memory     = "1024"

    scsi_controller = "virtio-scsi-pci"

    disks {
        disk_size = "10G"
        storage_pool = "local-lvm"
        type = "virtio"
    }

    network_adapters {
        model = "virtio"
        bridge = "vmbr0"
        firewall = "false"
    }

    cloud_init = true
    cloud_init_storage_pool = "local-lvm"

    boot = null
    boot_command = [
        "<esc><wait>",
        "auto console-keymaps-at/keymap=fr console-setup/ask_detect=false debconf/frontend=noninteractive fb=false preseed/url=${local.http_url}/preseed.cfg",
        "<wait><enter>"
    ]
    boot_wait = "6s"

    http_directory    = "./http" 
    http_bind_address = null
    http_interface    = null
    http_port_min     = var.http_server_port
    http_port_max     = var.http_server_port 

    ssh_username = var.username
    ssh_password = var.password
    ssh_private_key_file = null
    ssh_timeout = "15m"
}