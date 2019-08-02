package files

import (
	"bufio"
  "fmt"
  "io/ioutil"
	"strings"
	"os"
	"os/exec"
)

func readNetworkManagerConf() string {
  b, err := ioutil.ReadFile("/etc/NetworkManager/NetworkManager.conf")
  if err != nil {
		panic(err)
  }

  str := string(b)

	return str
}

// [main]
// plugins=ifupdown,keyfile
//
// [keyfile]
// unmanaged-devices=mac:00:19:e0:57:86:af
//
// -- or --
//
// [keyfile]
// unmanaged-devices=interface-name:eth*,except:interface-name:eth0;interface-name:wlan*
func checkInterfaceUnmanaged(confString, interfaceName string) (bool, error) {
  scanner := bufio.NewScanner(strings.NewReader(confString))
  for scanner.Scan() {
		// check if "unmanaged-devices" line
		line := scanner.Text()
		if strings.Contains(line, "unmanaged-devices") {
			// if it's an "unmanaged-devices" line check for card name
			if strings.Contains(line, "interface-name:" + interfaceName) {
				return true, nil
			}
		}
  }

  err := scanner.Err()
  return false, err
}

func getInterfaceUnmanagedConfArray(confString, interfaceName string) []string {
	var lines []string
	noUnmanagedDevicesLine := true

  scanner := bufio.NewScanner(strings.NewReader(confString))
  for scanner.Scan() {
		// check if "unmanaged-devices" line
		line := scanner.Text()
		if strings.Contains(line, "unmanaged-devices") {
			noUnmanagedDevicesLine = false
			// add interface to unmanaged-devices line
			lines = append(lines, line + ";interface-name:" + interfaceName)
		} else {
			lines = append(lines, line)
		}
  }

	// if there isn't an "unmanaged-devices" line, add one
	if noUnmanagedDevicesLine {
		for index, line := range lines {
			// if there's already a keyfile section, add line underneath it and return
			if strings.Contains(line, "[keyfile]") {
				linesTemp := append(lines[:index], "unmanaged-devices=interface-name:" + interfaceName)
				lines = append(linesTemp, lines[index:]...)
				return lines
			}
		}

		// add keyfile section
		lines = append(lines, "[keyfile]", "unmanaged-devices=interface-name:" + interfaceName)
		return lines
	}

	return lines
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func restartNetworkManager() {
	var (
		cmdOut          []byte
		err             error
	)

	cmdName := "service"
	cmdArgs := []string{"NetworkManager", "restart"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error restarting NetworkManager: ", err)
		return
	}

	_ = string(cmdOut)
}

func RestoreBackupConf() {
	fileLoc := "/etc/NetworkManager/NetworkManager.conf"
	err := os.Rename(fileLoc + ".bak", fileLoc)
	if err != nil {
		panic(err)
	}

	restartNetworkManager()
}

func EnsureInterfaceUnmanaged(interfaceName string) {
	confString := readNetworkManagerConf()

	interfaceUnmanaged, err := checkInterfaceUnmanaged(confString, interfaceName)

	if err != nil {
		fmt.Println("Error checking if interface is unmanaged!")
	}

	if !interfaceUnmanaged {
		lines := getInterfaceUnmanagedConfArray(confString, interfaceName)
		writeLines(lines, "/etc/NetworkManager/NetworkManager.conf.new")
		// fmt.Println("Configuration written to /etc/NetworkManager/NetworkManager.conf.new!")
		// fmt.Println("Backing up old config and moving the new one to the appropriate place...")

		fileLoc := "/etc/NetworkManager/NetworkManager.conf"
		rnErr := os.Rename(fileLoc, fileLoc + ".bak")
		if rnErr != nil {
			panic(rnErr)
		}

		rnErr = os.Rename(fileLoc + ".new", fileLoc)

		fmt.Println("Restarting network manager...")

		restartNetworkManager()
	} else {
		fmt.Println("Interface already unmanaged, not updating configuration...")
	}
}
