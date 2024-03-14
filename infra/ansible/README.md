# Homelab / Ansible

This directory contains the Ansible playbooks and roles used to manage the homelab.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Prerequisites](#prerequisites)
- [Directory Structure](#directory-structure)
- [Usage](#usage)
  - [Install NAS](#install-nas)
  - [Install K3S cluster](#install-k3s-cluster)
  - [Uninstall K3S cluster](#uninstall-k3s-cluster)
  - [Upgrade K3S cluster](#upgrade-k3s-cluster)
  - [Reboot K3S cluster](#reboot-k3s-cluster)
- [Secrets](#secrets)

## Prerequisites

1. Install [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
2. Put the vault password in the `.vault_pass` file.
    ```bash
    echo "<vault_password>" > .vault_pass
    ```

## Directory Structure

```bash
/infra/ansible
├── ansible.cfg # ansible configuration file
├── .vault_pass # git-ignored file containing the vault password
├── assets/ # resources and assets like scripts, binaries, etc.
├── roles/ # roles used by the playbooks
├── playbooks/ # playbooks used to manage the homelab
└── inventory/
    ├── hosts.yml # hosts definitions
    ├── groups.yml # groups definitions
    ├── group_vars/
    |   ├── <group>/
    |   |   ├── main.yml # main file containing the vars for the group
    |   |   └── secrets.yml # encrypted file containing the secrets vars for the group
    └── host_vars/ # host variables

```

## Usage

### Install NAS

```bash
ansible-playbook -i inventory playbooks/nas/install.yml
```

### Install K3S cluster

```bash
ansible-playbook -i inventory playbooks/k3s/install.yml
```

### Uninstall K3S cluster

```bash
ansible-playbook -i inventory playbooks/k3s/uninstall.yml
```

### Upgrade K3S cluster

```bash
ansible-playbook -i inventory playbooks/k3s/upgrade.yml
```

### Reboot K3S cluster

```bash
ansible-playbook -i inventory playbooks/k3s/reboot.yml
```

## Secrets

Secrets are stored in the `inventory/group_vars/*/secrets.yml` files. These files are encrypted using Ansible Vault. 

In order to play the playbooks that use these secrets, you need to provide the vault password in the git-ignored `.vault_pass` file.

Then to edit a `secrets.yml` file, use the following command:
```bash
ansible-vault edit inventory/group_vars/<group>/secrets.yml
```

> [!NOTE]
> The file will be automatically decrypted using the password issued from the `.vault_pass` file.