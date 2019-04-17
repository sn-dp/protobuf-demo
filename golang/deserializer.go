package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sn-dp/protobuf-demo/golang/proto"
	"io/ioutil"
)

func main(){
	input, e := ioutil.ReadFile("/Users/ksandeep/golang_ws/src/github.com/sn-dp/protobuf-demo/tran1")
	if e !=  nil {
		fmt.Print("error reading file")
	}
	tran1 := &com_thoughtworks_zpay.CompletedTransaction{}
	proto.Unmarshal(input, tran1)

	fmt.Print(tran1)

}
