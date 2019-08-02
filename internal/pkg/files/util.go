package files

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// WriteStringToFile takes a file path and a content string and writes the
// content to a file with the given path
func WriteStringToFile(filePath, str string) {
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()

	f.WriteString(str)

	// flush
	f.Sync()
}
