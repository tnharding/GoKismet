package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	for {
		conn, err := net.Dial("tcp", "192.168.82.1:2501")
		if err != nil { //error connecting to kismet server
			fmt.Println("Error connecting to kismet server.")
			time.Sleep(time.Millisecond * 60000) //We will try to connect once a minute
			continue
		}
		//We have connected to the server.
		//Do something
		fmt.Println("We have connected to the kismet server.")
		conn.Close()                         //close the connection
		time.Sleep(time.Millisecond * 60000) //We will try to connect once a minute
	}
}
