package main

import (
	"fmt"
	"os"
	"encoding/xml"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id: "1",
		Content: "Hello World!",
		Author: Author{
			Id: "2",
			Name: "Thuan Nguyen",
		},
	}

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	xmlFile.Write([]byte(xml.Header))

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file", err)
		return
	}
}