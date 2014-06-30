package main

import (
   "fmt"
   zmq "github.com/pebbe/zmq4"
   vampire "./vampire"
)

func main() {
	server, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	server.Connect("ipc:///tmp/helloworld.ipc")

	vamp := vampire.NewLord("ipc:///tmp/helloworld.ipc")
        if vamp == nil  {
		panic("Cannot Construct")
	}
	if err := vampire.GetErr(vamp); err != nil {
		panic(err)
	}

	_, err = server.Send("hello",0)
        if err != nil  {
		panic(err)
	}
        received,err := vampire.GetStakeBlock(vamp)
        if err != nil  {
		panic(err)
	}
	fmt.Println(received + " World")

	vamp2 := vampire.NewMinion("ipc:///tmp/helloworld.ipc")
	server2, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		panic(err)
	}
	server2.Bind("ipc:///tmp/helloworld.ipc")
	_, err = server2.Send("hello2",0)
        if err != nil  {
		panic(err)
	}
        received,err = vampire.GetStakeBlock(vamp2)
        if err != nil  {
		panic(err)
	}
	fmt.Println(received + " World")
 
}
   
