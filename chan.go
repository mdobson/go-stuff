package main

import "fmt"

func main() {

	//Method of communication between them
	messages := make(chan string)

	//Separate process
	//We Queue up a simple ping message into the channel.
	go func(){ 
		fmt.Println("Queue message into channel in goroutine")
		messages <- "ping"
	}()


	//Separate process
	go func(){
		//We use a loop to keep the process alive
		//The select statement to waits for the channel to send a message to us.
		for {
			select {
				case msg:= <-messages:
					fmt.Println("Dequeue message from channel:",msg)
			}
		}
	}()

	
	fmt.Println("End of main execution wait for input...")
	var input string
    fmt.Scanln(&input)	
}