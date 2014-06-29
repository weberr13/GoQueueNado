package main

import (
   "fmt"
   zmq "github.com/pebbe/zmq4"
)

func main() {
	client, err := zmq.NewSocket(zmq.PULL)
	if err != nil {
		panic(err)
	}
	server, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}

	server.Bind("ipc:///tmp/helloworld.ipc")
        client.Connect("ipc:///tmp/helloworld.ipc")

	_, err = server.Send("hello",0)
        if err != nil  {
		panic(err)
	}
        received, err := client.Recv(0)
        if err != nil  {
		panic(err)
	}
	fmt.Println(received + " World")
 
}
   
