#!/bin/bash

turn_off () {
  sudo nmcli radio wifi off && \
  sudo rfkill unblock wlan

  sudo service systemd-resolved stop
}

turn_on () {
  sudo nmcli radio wifi on
  sudo service systemd-resolved start
}

set_ip () {
  ifconfig eth1 10.0.0.1 netmask 255.255.255.0
}

# restartDnsMasq () {
#   if [ -d /run/systemd/system ]; then
#     systemctl reload --no-block dnsmasq >/dev/null 2>&1 || true
#   else
#     invoke-rc.d dnsmasq restart >/dev/null 2>&1 || true
#   fi
# }

# turn_off
turn_on
