package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/totality-assignment/user-grpc-service/proto"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 3001, "listening port")

type bidirectionalStreamServer struct {
	pb.UnimplementedUserServiceServer
}

type User struct {
	UserId    int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

var users = map[int32]pb.UserResponse{
	1:  {UserId: 1, FirstName: "Steve", City: "LA", Phone: "123-456-7890", Height: 5.8, Married: true},
	2:  {UserId: 2, FirstName: "Alice", City: "New York", Phone: "555-123-4567", Height: 5.6, Married: false},
	3:  {UserId: 3, FirstName: "John", City: "Chicago", Phone: "777-999-8888", Height: 6.0, Married: true},
	4:  {UserId: 4, FirstName: "Eva", City: "San Francisco", Phone: "123-789-4560", Height: 5.7, Married: false},
	5:  {UserId: 5, FirstName: "Michael", City: "Houston", Phone: "999-555-1234", Height: 6.2, Married: false},
	6:  {UserId: 6, FirstName: "Linda", City: "Seattle", Phone: "444-444-4444", Height: 5.9, Married: true},
	7:  {UserId: 7, FirstName: "David", City: "Boston", Phone: "111-222-3333", Height: 5.8, Married: true},
	8:  {UserId: 8, FirstName: "Olivia", City: "Los Angeles", Phone: "987-654-3210", Height: 5.6, Married: false},
	9:  {UserId: 9, FirstName: "Daniel", City: "San Diego", Phone: "333-666-9999", Height: 6.1, Married: true},
	10: {UserId: 10, FirstName: "Sophia", City: "Miami", Phone: "222-777-5555", Height: 5.4, Married: true},
	11: {UserId: 11, FirstName: "William", City: "Dallas", Phone: "444-555-6666", Height: 5.9, Married: false},
	12: {UserId: 12, FirstName: "Mia", City: "Phoenix", Phone: "555-777-8888", Height: 5.7, Married: false},
	13: {UserId: 13, FirstName: "Emily", City: "San Antonio", Phone: "999-888-7777", Height: 6.0, Married: true},
	14: {UserId: 14, FirstName: "James", City: "Atlanta", Phone: "123-789-4561", Height: 5.6, Married: false},
	15: {UserId: 15, FirstName: "Lucas", City: "Denver", Phone: "555-555-5555", Height: 6.3, Married: true},
	16: {UserId: 16, FirstName: "Ava", City: "Minneapolis", Phone: "444-333-2222", Height: 5.5, Married: true},
	17: {UserId: 17, FirstName: "Benjamin", City: "Portland", Phone: "777-999-5555", Height: 5.8, Married: false},
	18: {UserId: 18, FirstName: "Sophie", City: "Charlotte", Phone: "111-111-1111", Height: 5.8, Married: false},
	19: {UserId: 19, FirstName: "Jack", City: "San Francisco", Phone: "123-456-7891", Height: 6.1, Married: true},
	20: {UserId: 20, FirstName: "Isabella", City: "Tampa", Phone: "555-111-3333", Height: 5.7, Married: true},
}

func (s *bidirectionalStreamServer) GetUserById(stream pb.UserService_GetUserByIdServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		u, ok := users[req.UserId]

		if !ok {
			u = pb.UserResponse{UserId: 100, FirstName: "Not Found", City: "Not Found", Phone: "Not Found", Height: 10000.99, Married: false}
		}
		err = stream.Send(&u)
		if err != nil {
			return err
		}
		log.Printf("[RECEVIED REQUEST] : %v\n", req)
	}
}

func (s *bidirectionalStreamServer) GetUserListByIds(stream pb.UserService_GetUserListByIdsServer) error {

	var u pb.UserResponseList
	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	for i := 0; i < len(req.GetUserRequestList()); i++ {
		u1, ok := users[req.GetUserRequestList()[i].UserId]
		if !ok {
			u1 = pb.UserResponse{UserId: req.GetUserRequestList()[i].UserId, FirstName: "Not Found", City: "Not Found", Phone: "Not Found", Height: 10000.99, Married: false}
		}
		u.UserResponseList = append(u.UserResponseList, &u1)

	}

	err = stream.Send(&u)
	if err != nil {
		return err
	}
	log.Printf("[RECEVIED REQUEST] : %v\n", req)
	return nil

}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("listen error: %v", err)
	} else {
		log.Printf("server listen: ", addr)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &bidirectionalStreamServer{})
	grpcServer.Serve(listener)

}
