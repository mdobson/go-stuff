package main

import "fmt"

func print(str string){
	for i := 0; i < 10; i++ {
        fmt.Println(str, ":", i)
    }
}

func main(){
	print("non concurrent")

	go print("concurrent")
	go print("concurrent")

	var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}