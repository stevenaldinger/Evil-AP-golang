# Create wpa_supplicant config

1. `wpa_passphrase NETGEAR12345 passwordhere > wpa.conf`

# Run wpa_supplicant config

1. `wpa_supplicant -D nl80211 -i wlp1s0 -c wpa.conf`
