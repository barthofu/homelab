#!/usr/bin/env bash

# Copyright (c) 2021-2024 tteck
# Author: tteck (tteckster)
# License: MIT
# https://github.com/tteck/Proxmox/raw/main/LICENSE

function header_info {
clear
cat <<"EOF"
  ______      _ __                __
 /_  __/___ _(_) /_____________ _/ /__
  / / / __ `/ / / ___/ ___/ __ `/ / _ \
 / / / /_/ / / (__  ) /__/ /_/ / /  __/
/_/  \__,_/_/_/____/\___/\__,_/_/\___/

EOF
}
header_info
set -e
while true; do
  read -p "This will add Tailscale to an existing LXC Container ONLY. Proceed(y/n)?" yn
  case $yn in
  [Yy]*) break ;;
  [Nn]*) exit ;;
  *) echo "Please answer yes or no." ;;
  esac
done
header_info
echo "Loading..."
function msg() {
  local TEXT="$1"
  echo -e "$TEXT"
}

NODE=$(hostname)
MSG_MAX_LENGTH=0
while read -r line; do
  TAG=$(echo "$line" | awk '{print $1}')
  ITEM=$(echo "$line" | awk '{print substr($0,36)}')
  OFFSET=2
  if [[ $((${#ITEM} + $OFFSET)) -gt ${MSG_MAX_LENGTH:-} ]]; then
    MSG_MAX_LENGTH=$((${#ITEM} + $OFFSET))
  fi
  CTID_MENU+=("$TAG" "$ITEM " "OFF")
done < <(pct list | awk 'NR>1')

while [ -z "${CTID:+x}" ]; do
  CTID=$(whiptail --backtitle "Proxmox VE Helper Scripts" --title "Containers on $NODE" --radiolist \
    "\nSelect a container to add Tailscale to:\n" \
    16 $(($MSG_MAX_LENGTH + 23)) 6 \
    "${CTID_MENU[@]}" 3>&1 1>&2 2>&3) || exit
done

CTID_CONFIG_PATH=/etc/pve/lxc/${CTID}.conf
cat <<EOF >>$CTID_CONFIG_PATH
lxc.cgroup2.devices.allow: c 10:200 rwm
lxc.mount.entry: /dev/net/tun dev/net/tun none bind,create=file
EOF
header_info
msg "Installing Tailscale..."
lxc-attach -n $CTID -- apk add curl
lxc-attach -n $CTID -- curl -fsSL "https://tailscale.com/install.sh" | sh -s -- || exit
lxc-attach -n $CTID -- sh -c "echo 'net.ipv4.ip_forward = 1' | tee -a /etc/sysctl.d/99-tailscale.conf"
lxc-attach -n $CTID -- sh -c "echo 'net.ipv6.conf.all.forwarding = 1' | tee -a /etc/sysctl.d/99-tailscale.conf"
lxc-attach -n $CTID -- sh -c "echo 'sysctl -p /etc/sysctl.d/99-tailscale.conf' | tee -a ~/.ashrc"
lxc-attach -n $CTID -- sh -c "echo 'export ENV=$HOME/.ashrc' | tee -a ~/.profile"
msg "Installed Tailscale"
sleep 2
msg "Restarting LXC..."
lxc-stop -n $CTID
lxc-start -n $CTID
msg "\e[1;32m ✔ Completed Successfully!\e[0m"
msg "\e[1;31m Run tailscale up in the LXC console\e[0m"