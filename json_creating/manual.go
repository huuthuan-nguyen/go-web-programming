package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	post := Post{
		Id: 1,
		Content: "Hello World!",
		Author: Author{
			Id: 2,
			Name: "Thuan Nguyen",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "Have a good day!",
				Author: "Betty",
			},
			Comment{
				Id: 4,
				Content: "Have a great time!",
				Author: "Billy",
			},
		},
	}

	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}