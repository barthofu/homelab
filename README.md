# Bartho's Homelab

This repository holds the code and configuration for my homelab.

## Features

- Full GitOps (terraform, ansible, etc)
    - TF proxmox provider?
    - Renovate or smthg else to automate versions upgrades (with aproval!)
    - CI and CD
- On-site backups
    - Proxmox snapshots (if used)
    - Docker volumes (if needed)
    - Personal and important files
- Off-site encrypted backups of my important files (Blackblaze B2, iDrive, Wasabi, etc)
- SSO everywhere on my services (Authelia)
- Full security over my network, software and hardware
- Monitoring and alert system
- Traefik
- Modular overall architecture (easy to remove or add features/components)
- Private container registry

Bonus (if not too complicated):
- Infrastructure testing

## Services

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

- [ ]  

### Media

### 

## The plan

The plan is to have a fully automated and secure homelab. I want to be able to deploy new services with a single command and have them fully secured and monitored. I also want to have a full backup system in place, both on-site and off-site.

### Deployment

- I will mainly use [pulumi](https://www.pulumi.com/) to deploy my entire infrastrucure, such as:
    - Proxmox (and its VMs)
    - Cloudflare
    - OCI

I should then verify if pulumi can be used to deploy to deploy the kubernetes cluster on my provisionned Proxmox. If not, I will use certainly use [ansible](https://www.ansible.com/).