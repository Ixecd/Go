package main

import "fmt"
import "encoding/json"

type Movie struct {
	Title string 	`json:"Title"`
	Year int		`json:"Year"`
	Price int		`json:"Price"`
	Actors []string	`json:"Actors"`
}

func main() {

	movie := Movie{"x", 2024, 50, []string{"actor1", "actor2"}}

	// 编码过程,将结构体转换为json
	// Marshal -> 元帅、组织者、整理、调动.
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程 jsonStr -> struct
	// jsonStr = {"Title":"x","Year":2024,"Price":50,"Actors":["actor1","actor2"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)

	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	fmt.Printf("%v\n", myMovie)
}