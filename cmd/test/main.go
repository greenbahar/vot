package main

import (
	"context"
	"fmt"
	"github.com/greenbahar/vot/x/vot/types"
	"log"

	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

func main() {
	//flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	queryClient := types.NewQueryClient(conn)
	msgClient := types.NewMsgClient(conn)

	/*
		params := &types.QueryGetNextVoteRequest{}
		res, resErr := queryClient.NextVote(context.Background(), params)
		if resErr != nil {
			fmt.Println(resErr)
		}
		fmt.Println(res)
	*/

	params2 := &types.QueryAllStoredVoteRequest{}
	res2, res2Err := queryClient.StoredVoteAll(context.Background(), params2)
	if res2Err != nil {
		fmt.Println(res2Err)
	}
	fmt.Println(res2)

	params3 := &types.MsgCreateVote{
		Question: "are you robot?",
		Options:  "yes//no",
		Creator:  "cosmos142dr726hgx3rf2u6cj53efse3h63c22laf6mf4",
		Days:     1,
	}
	result, resultErr := msgClient.CreateVote(context.Background(), params3)
	if resultErr != nil {
		fmt.Println(resultErr)
	}
	fmt.Println(result)
}
