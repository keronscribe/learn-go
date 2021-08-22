package main

import (
	pb "github.com/keronscribe/learn-go/grpc/route"
	"google.golang.org/grpc"
	"log"
)

func main (){
	conn,err:= grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock()){
		if err != nil{
			log.Fatalln("Client cannot dail grpc server")
		}
	}
	defer conn.Close() // 在程式的最後要關掉 conn

	client := pb.NewRouteGuideClient(conn)
	getFeat(client)
}
