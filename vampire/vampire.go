package vampire

/* 
The Receiver of a ZMQ PULL socket on a single endpoint
*/

import (
	zmq "github.com/pebbe/zmq4"
)

type Vampire struct {
	ownsSocket	bool
	endPoint	string
	theSocket	*zmq.Socket
	err		error
}


func NewLord(where string) *Vampire {
	return newVamp(where,true)
}

func NewMinion(where string) *Vampire {
	return newVamp(where,false)
}

func newVamp(where string, owns bool) *Vampire {
	vamp := new(Vampire)
	vamp.ownsSocket = false
	vamp.endPoint = where
	vamp.theSocket,vamp.err = zmq.NewSocket(zmq.PULL)
	if vamp.err != nil {	
		return nil
	}
	if owns {
		vamp.theSocket.Bind(where)
	} else {
		vamp.theSocket.Connect(where)
	}
	
	return vamp
}	

func GetErr(vamp *Vampire) error {
	return vamp.err
}

func GetStakeBlock(vamp *Vampire) (string, error) {
	if vamp.err != nil {
		return "",GetErr(vamp)
	}	
	received := ""
	received, vamp.err = vamp.theSocket.Recv(0)
	if vamp.err != nil {
		return "",GetErr(vamp)
	}
	return received, nil
}
