package main

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	err := doc.ReadFromFile("computer.xml")
	if err != nil {
		panic(err)
	}

	root := doc.SelectElement("server")
	fmt.Println("ROOT element:", root.Tag)

	// Modify
	for _, service := range root.SelectElements("service") {
		fmt.Println("CHILD element:", service.Tag)
		if con := service.SelectElement("Connector"); con != nil {
			port := con.SelectAttrValue("port", "unknown")
			fmt.Printf("  Connector: %s (%s)\n", con.Text(), port)
			con.RemoveAttr("port")
			con.CreateAttr("port", "19191")
			con.SetText("My Server")
		}
	}

	// Output to console.
	doc.WriteTo(os.Stdout)

	// Output to file.
	// doc.Indent(2)
	// f, err := os.Create("computer_mod.xml")
	// doc.WriteTo(f)
}
