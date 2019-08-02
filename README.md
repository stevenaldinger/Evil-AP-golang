# Evil AP in Golang

This is an example of an Evil WiFi Access Point written in golang.

Once a victim connects to the evil access point, DNS overrides on apple.com will make their browser go to a static site that you control.

## Project layout

Check out `cmd/evil-twin/main.go` and `internal/pkg/...` for a look at the code.

```
.
├── cmd
│   └── evil-twin
│       └── main.go
├── conf
│   ├── dnsmasq.conf
│   ├── fake_hosts.conf
│   ├── hostapd.conf
│   └── hostapd.conf.example
├── Dockerfile -> ./build/package/Dockerfile
├── internal
│   └── pkg
│       ├── commands
│       │   └── commands.go
│       ├── files
│       │   ├── dnsmasq_conf.go
│       │   ├── fake_hosts.go
│       │   ├── hostapd_conf.go
│       │   ├── network_manager.go
│       │   └── util.go
│       ├── httpserver
│       │   └── http_server.go
│       └── ip
│           ├── ip_address.go
│           └── ip_tables.go
├── Makefile
├── README.md
├── scripts
│   ├── disable-dns-binding-53.sh
│   └── kill-network-manager.sh
├── static
│   ├── ac
│   │   ├── ac-films
│   │   │   └── 6.0.0
│   │   ├── globalfooter
│   │   │   └── 3
│   │   ├── globalnav
│   │   │   └── 4
│   │   └── localnav
│   │       └── 4
│   ├── facebook.html
│   ├── index.desktop.html
│   ├── index.html
│   └── v
│       └── home
│           └── dz
└── wpa_supp
    ├── README.md
    └── wpa.conf

```
