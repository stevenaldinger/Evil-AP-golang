package ip

import (
	"fmt"
	"os"
	"os/exec"
	"net"
)

func checkIfIPAvailable(ip string) bool {
	var (
		cmdOut          []byte
		err             error
	)

	// root@deviant:~# if ping -c1 -w3 10.0.0.1 >/dev/null 2>&1; then echo "IP taken"; else echo "IP not taken"; fi
	// IP not taken
	// root@deviant:~# if ping -c1 -w3 192.168.0.23 >/dev/null 2>&1; then echo "IP taken"; else echo "IP not taken"; fi
	// IP taken

	cmdName := "sh"
	cmdArgs := []string{"-c", "if ping -c1 -w3 " + ip + " >/dev/null 2>&1; then echo -n \"IP taken\"; else echo -n \"IP not taken\"; fi"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error checking if IP address is available: ", err)
		return false
	}

	output := string(cmdOut)

	if output == "IP not taken" {
		return true
	} else {
		return false
	}
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// https://gist.github.com/kotakanbe/d3059af990252ba89a82
func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func FindAvailableIPInCIDR(cidr string) string {
	hostArray, _ := hosts(cidr)

	for _, host := range hostArray {
		if checkIfIPAvailable(host) {
			return host
		}
	}

	return ""
}

// https://www.lifewire.com/what-is-a-private-ip-address-2625970
func AssignIPAddress(interfaceName, ip, netmask string) error {
	// Which IP Addresses Are Private?
	//
	// The Internet Assigned Numbers Authority (IANA) reserves the following IP address blocks for use as private IP addresses:
	//
	// 	 10.0.0.0 to 10.255.255.255
	// 	 172.16.0.0 to 172.31.255.255
	// 	 192.168.0.0 to 192.168.255.255
	//

	// https://www.tecmint.com/ifconfig-command-examples/
	// ifconfig eth0 172.16.25.125 netmask 255.255.255.0
	var (
		err             error
	)

	cmdName := "ifconfig"
	cmdArgs := []string{interfaceName, ip, "netmask", netmask}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error assigning IP to interface: ", err)
		return err
	}

	return nil
}
