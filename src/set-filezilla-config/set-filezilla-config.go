package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	// err := doc.ReadFromFile("./config/filezilla.xml")
	// err := doc.ReadFromFile("./filezilla.xml")
	err := doc.ReadFromFile(os.Getenv("FILEZILLA_ROOT") + "/config/filezilla.xml")
	if err != nil {
		panic(err)
	}

	root := doc.SelectElement("FileZilla3")
	fmt.Println("ROOT element:", root.Tag)

	// Modify
	for _, settings := range root.SelectElements("Settings") {
		fmt.Println("CHILD element:", settings.Tag)

		for _, setting := range settings.SelectElements("Setting") {
			name := setting.SelectAttrValue("name", "unknown")
			if strings.Contains(name, "Default editor") {
				fmt.Printf("  Default editor: %s (%s)\n", setting.Text(), name)

				// Get parent directory
				// wd, err := os.Getwd()
				wd := os.Getenv("FILEZILLA_ROOT")
				if err != nil {
					panic(err)
				}
				parent := filepath.Dir(wd)

				setting.SetText("2" + parent + "\\notepad2\\notepad2.exe")
			} else if strings.Contains(name, "Last local directory") {
				fmt.Printf("  Last local directory: %s (%s)\n", setting.Text(), name)

				// Get parent directory
				wd, err := os.Getwd()
				if err != nil {
					panic(err)
				}
				// currentWorkspace := filepath.Dir(wd)

				setting.SetText(wd)
			}
		}

	}

	doc.Indent(2)
	// doc.WriteTo(os.Stdout)
	// f, err := os.Create("./config/filezilla.xml")
	// f, err := os.Create("./filezilla.xml")
	f, err := os.Create(os.Getenv("FILEZILLA_ROOT") + "/config/filezilla.xml")
	doc.WriteTo(f)
}
