package files

import (
	"fmt"
	"text/template"
	"bytes"
	"os"
)

type DnsmasqVariables struct {
	Interface string
	ListenAddress string
}

// dhcp-range=10.0.0.10,10.0.0.250,12h:  Client IP address will range from 10.0.0.10 to 10.0.0.250 and default lease time is 12 hours.
// dhcp-option=3,10.0.0.1:  3 is code for Default Gateway followed by IP of D.G i.e. 10.0.0.1
// dhcp-option=6,10.0.0.1:  6 for DNS Server followed by IP address
const dnsmasqConf = `
interface={{.Interface}}
dhcp-range=10.0.0.10,10.0.0.250,255.255.255.0,12h
dhcp-option=3,10.0.0.1
dhcp-option=6,10.0.0.1
server=8.8.8.8
log-queries
log-dhcp
listen-address={{.ListenAddress}}
`

func WriteDNSMasqConfFile(filePath string, vars *DnsmasqVariables) {
	var (
		err error
	)

	t := template.New("dnsmasq config template")

	t, err = t.Parse(dnsmasqConf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing dnsmasq config:", err)
		panic(err)
	}

	var tpl bytes.Buffer
	if tErr := t.Execute(&tpl, *vars); tErr != nil {
		fmt.Fprintln(os.Stderr, "Error executing dnsmasq config template:", tErr)
    panic(tErr)
	}

	result := tpl.String()

	WriteStringToFile(filePath, result)
}
