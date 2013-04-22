package main

import "fmt"


type data struct {
	message string
	channel string
}

func main(){
	m := make(map[string]map[int]data)
	m["test"] = map[int]data {
		0:data{"hello", "test"}
	}
	fmt.Println(m["test"][0])
}