package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	jsonData := `[{"First":"James","Last":"Bond","Age":33},{"First":"vaibhav","Last":"Singh","Age":21}]`
	bs := []byte(jsonData)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println(i, v.First, v.Last, v.Age)
	}
}
