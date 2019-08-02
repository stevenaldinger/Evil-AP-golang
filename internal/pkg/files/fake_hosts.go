package files

import (
	"fmt"
	"text/template"
	"bytes"
	"os"
)

type FakeHostsVariables struct {
	// ideally an http server
	IP string
}

const fakeHosts = `
{{.IP}} apple.com
{{.IP}} facebook.com
`

func WriteFakeHostsFile(filePath string, vars *FakeHostsVariables) {
	var (
		err error
	)

	t := template.New("fake hosts config template")

	t, err = t.Parse(fakeHosts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing fake hosts config:", err)
		panic(err)
	}

	var tpl bytes.Buffer
	if tErr := t.Execute(&tpl, *vars); tErr != nil {
		fmt.Fprintln(os.Stderr, "Error executing fake hosts config template:", tErr)
    panic(tErr)
	}

	result := tpl.String()

	WriteStringToFile(filePath, result)
}
