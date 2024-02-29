# Build Definition to create the VM Template
build {

    name = "debian-12"
    sources = [
        "source.file.preseed",
        "proxmox-iso.debian-12"
    ]

    # Provisioning post-installation commands
    provisioner "shell" {
        inline = [
            "sudo apt -y update",
            "sudo apt -y dist-upgrade",
            "sudo apt -y install curl wget git nano htop screen unzip zip lsof",
            "sudo apt -y autoremove --purge",
            "sudo apt -y clean",
            "sudo apt -y autoclean",
            "sudo sync"
        ]
    }
}