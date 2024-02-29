# Bartho's Homelab

This repository holds the code and configuration for my homelab.

## Why a homelab?

My fascination with IT has driven me to establish my own home lab. While I've been utilizing various VPS services, I yearned for greater control over my infrastructure, along with the opportunity to delve into domains I had previously overlooked (*gazing at you, Network shit*).

## Philosophy and design goals

My homelab should be a place where I can experiment with new technologies and learn new things while being capable of hosting important services and production things.

The big plan is to have a fully automated and reproducible infrastructure. I aim to swiftly dismantle and rebuild the entire setup within minutes, with the capability to deploy new services in a matter of seconds. The crucial aspect for me is achieving reproducibility alongside flexibility.

## Features

- [ ] Automated and reproducible infrastructure deployment (Pulumi, Ansible, Packer, etc)
- [ ] GitOps practices
- [ ] Data security
  - [ ] On-site backups
    - [ ] Proxmox snapshots (if used)
    - [ ] Docker volumes (if needed)
    - [ ] Personal and important files
  - [ ] Off-site encrypted backups of my important files (Blackblaze B2, iDrive, Wasabi, etc)
- [ ] SSO everywhere on my services (Authelia)
- [ ] Full security over my network, software and hardware
- [ ] Monitoring and alert system
- [ ] Modular overall architecture (easy to remove or add features/components)

### Bonus
- [ ] Infrastructure testing

## Hardware

*[insert photo]*

- 1 × HP `EliteDesk 800 G4 Mini`:
    - CPU: `Intel Core i5-8500 @ 3.40GHz`
    - RAM: `32GB`
    - NVMe: `512GB`
    - SSD: `1TB`
- Yottamaster `5-bay USB3.0`:
    - 2 × `8TB` HDD
    - 1 × `2TB` HDD

The Yottamaster DAS enclosure is connected to the HP EliteDesk via USB3.0. The HP EliteDesk is running Proxmox and the Yottamaster is used for backups and file storage.

## Services

Here is a list of services I want to run on my homelab:

## Tooling

2. [Packer](https://www.packer.io/) to create the base VMs templates in Proxmox
3. [Pulumi](https://www.pulumi.com/) to provision the infrastructure:
   1. Proxmox VMs and CTs
   2. Cloudflare
   3. OCI
   4. Tailscale
   5. Amazon S3 Glacier / Blackblaze B2 / iDrive / Wasabi
4. [Ansible](https://www.ansible.com/) to configure the VMs and CTs post-provisioning
5. [k3s](https://k3s.io/) cluster for running all my services (deployed with Ansible?)
6. [Renovate](https://www.whitesourcesoftware.com/free-developer-tools/renovate/) to automate versions upgrades inside this repository
7. [Flux](https://fluxcd.io/) to automate the deployment of the whole infrastructure using subsquent tools