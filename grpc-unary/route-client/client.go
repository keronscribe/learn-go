package main

import (
	"context"
	"fmt"
	pb "github.com/keronscribe/learn-go/grpc-unary/route" // 透過 `proto` 生成的 Client Stub
	"google.golang.org/grpc"                              // grpc 包
	"log"
)

func getFeat (client pb.RouteGuideClient){
	feature,err := client.GetFeature(context.Background(),&pb.Point{
		Latitude: 353931000,
		Longitude: 139444400,
	})
	if err != nil{
		log.Fatalln()
	}
	fmt.Println(feature)
}

func main () {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil{
		log.Fatalln("Client cannot dail grpc server")
	}

	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	getFeat(client)
}