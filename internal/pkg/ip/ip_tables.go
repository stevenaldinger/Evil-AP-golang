package ip

import (
	"fmt"
	"os"
	"os/exec"
)

// iptables -t nat -A POSTROUTING --out-interface wlan0 -j MASQUERADE
// iptables -A FORWARD --in-interface wlan1 -j ACCEPT
func AddIPTablesRules(apInterface, internetInterface string) error {
	var (
		err             error
	)

	cmdName := "iptables"
	cmdArgs := []string{"-t", "nat", "-A", "POSTROUTING", "--out-interface", internetInterface, "-j", "MASQUERADE"}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error adding iptables nat rule: ", err)
		return err
	}

	cmdName = "iptables"
	cmdArgs = []string{"-A", "FORWARD", "--in-interface", apInterface, "-j", "ACCEPT"}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error adding iptables forward rule: ", err)
		return err
	}

	return nil
}

// iptables -t nat -D POSTROUTING --out-interface wlan0 -j MASQUERADE
// iptables -D FORWARD --in-interface wlan1 -j ACCEPT
func DeleteIPTablesRules(apInterface, internetInterface string) error {
	var (
		err             error
	)

	cmdName := "iptables"
	cmdArgs := []string{"-t", "nat", "-D", "POSTROUTING", "--out-interface", internetInterface, "-j", "MASQUERADE"}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error deleting iptables nat rule: ", err)
		return err
	}

	cmdName = "iptables"
	cmdArgs = []string{"-D", "FORWARD", "--in-interface", apInterface, "-j", "ACCEPT"}

	if _, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error deleting iptables forward rule: ", err)
		return err
	}

	return nil
}
