package commands

import (
  "fmt"
	"os"
	"os/exec"
)

func RunDNSMasq() *exec.Cmd {
	cmdName := "dnsmasq"
	cmdArgs := []string{"-C", "conf/dnsmasq.conf", "-H", "conf/fake_hosts.conf", "-d"}

  cmd := exec.Command(cmdName, cmdArgs...)

  if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running dnsmasq: ", err)
  }

  return cmd
}

func RunHostAPD() *exec.Cmd {
	cmdName := "hostapd"
	cmdArgs := []string{"conf/hostapd.conf"}

  cmd := exec.Command(cmdName, cmdArgs...)

  if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running hostapd: ", err)
  }

  return cmd
}

func CancelCommand(cmd *exec.Cmd, name string) {
  if err := cmd.Process.Kill(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error trying to kill process: " + name, err)
  }
}
