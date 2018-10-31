package main // import "set-filezilla-config"

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	// err := doc.ReadFromFile("./tomcat-conf.xml")
	err := doc.ReadFromFile(os.Getenv("TOMCAT_PATH") + "/conf/server.xml")
	if err != nil {
		panic(err)
	}

	root := doc.SelectElement("Server")
	fmt.Println("ROOT element:", root.Tag)

	// Modify
	for _, service := range root.SelectElements("Service") {
		fmt.Println("CHILD element:", service.Tag)

		for _, engine := range service.SelectElements("Engine") {
			fmt.Println("CHILD element:", engine.Tag)

			for _, host := range engine.SelectElements("Host") {
				fmt.Println("CHILD element:", host.Tag)

				for _, context := range host.SelectElements("Context") {
					fmt.Println("CHILD element:", context.Tag)
					context.RemoveAttr("docBase")
					context.CreateAttr("docBase", os.Getenv("current_workspace"))
				}
			}
		}
	}

	// Output to console.
	// doc.WriteTo(os.Stdout)

	// Output to file.
	doc.Indent(2)
	f, err := os.Create(os.Getenv("TOMCAT_PATH") + "/conf/server.xml")
	doc.WriteTo(f)
}
