package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	sm "shiv/api/models"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

func main() {

	conn := getConnection()
	defer conn.Close()
	client := sm.NewUserAPIClient(conn)

	builder := getrequest()

	response, err := client.Session_charging(context.TODO(), builder, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Println(err.Error())
	}

	sessionResultT := response.UnPack()
	byt, _ := json.MarshalIndent(sessionResultT, " ", "  ")
	fmt.Println(string(byt))

}

func getConnection() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	return conn
}

func getrequest() *flatbuffers.Builder {
	var err error
	reqBytes, err := ioutil.ReadFile("ip.json")
	if err != nil {
		panic(err.Error())
	}

	UserRequestT := &sm.UserRequestT{}
	json.Unmarshal(reqBytes, UserRequestT)

	b := flatbuffers.NewBuilder(0)
	b.Finish(UserRequestT.Pack(b))

	return b

}
