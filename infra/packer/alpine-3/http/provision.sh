#!/bin/sh

apk update
apk add --no-cache sudo python3
apk add --no-cache cloud-init cloud-utils-growpart e2fsprogs-extra

# Add iso9660 as a valid filesystem. Necessary for cloud-init to mount NoData with proxmox cloud-init drive (/dev/sr[0-9]).
echo 'isofs' > /etc/modules-load.d/isofs.conf
chmod -x /etc/modules-load.d/isofs.conf
setup-cloud-init
echo 'datasource_list: [ NoCloud, ConfigDrive, None ]' > /etc/cloud/cloud.cfg.d/99_pve.cfg
chmod 644 /etc/cloud/cloud.cfg.d/99_pve.cfg