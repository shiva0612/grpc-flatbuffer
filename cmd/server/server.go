package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	sm "shiv/api/models"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ApierServer struct {
	sm.UnimplementedUserAPIServer
}

func newServer() *ApierServer {
	return &ApierServer{}
}

func (s *ApierServer) Session_charging(ctx context.Context, req *sm.UserRequest) (*flatbuffers.Builder, error) {
	obj := req.UnPack()
	byt, _ := json.MarshalIndent(obj, " ", "  ")
	fmt.Println("---request json received--- ")
	fmt.Println(string(byt))
	fmt.Println("--------------------------------")

	response := getresponse()
	return response, nil
}

func getresponse() *flatbuffers.Builder {
	var err error
	respBytes, err := ioutil.ReadFile("op.json")
	UserResponseT := &sm.UserResponseT{}
	json.Unmarshal(respBytes, UserResponseT)
	if err != nil {
		panic(err.Error())
	}
	b := flatbuffers.NewBuilder(0)
	b.Finish(UserResponseT.Pack(b))
	return b
}
