# Homelab / Ansible

This directory contains the Ansible playbooks and roles used to manage the homelab.

## Directory Structure

- `roles/`: Contains the roles used by the playbooks.
- `playbooks/`: Contains the playbooks used to manage the homelab.
- `inventory/`: Contains the inventory files used by the playbooks.
- `group_vars/`: Contains the group variables used by the playbooks.
- `host_vars/`: Contains the host variables used by the playbooks.

## Usage

### SSH Configuration

1.
> eval \`ssh-agent -s\`
1. `ssh-add ../../secrets/homelab-ansible`