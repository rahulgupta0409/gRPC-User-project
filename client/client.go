package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	pb "github.com/totality-assignment/user-grpc-service/proto"
)

var (
	addr = flag.String("addr", "0.0.0.0:3001", "connect server address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connection server error: %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewUserServiceClient(conn)
	stream1, err := grpcClient.GetUserById(context.Background())
	stream2, err := grpcClient.GetUserListByIds(context.Background())
	if err != nil {
		log.Fatal("receive stream error: %v", err)
	}
	var i int32
	for i = 1; i < 10; i++ {
		stream1.Send(&pb.UserRequest{UserId: i})
		resp, err := stream1.Recv()
		if err != nil {
			log.Fatalf("resp error: %v", err)
		}

		log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
	}

	stream2.Send(&pb.UserList{
		UserRequestList: []*pb.UserRequest{{UserId: 1}, {UserId: 2}, {UserId: 3}, {UserId: 200}},
	})
	resp, err := stream2.Recv()
	if err != nil {
		log.Fatalf("resp error: %v", err)
	}

	log.Printf("[RECEIVED RESPONSE]: %v\n", resp)
}
