package main

import (
	"context"
	pb "github.com/keronscribe/learn-go/grpc-unary/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

type routeGuideServer struct {
	features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature {
			{
				Name: "東京鐵塔",
				Location: &pb.Point{
					Latitude:  353931000,
					Longitude: 139444400,
				},
			},
			{
				Name: "淺草寺",
				Location: &pb.Point{
					Latitude:  357147651,
					Longitude: 139794466,
				},
			},
			{
				Name: "晴空塔",
				Location: &pb.Point{
					Latitude:  357100670,
					Longitude: 139808511,
				},
			},
		},
	}
}

func (s *routeGuideServer) GetFeature(cxt context.Context, point *pb.Point) (*pb.Feature, error){
	for _,feature := range s.features{
		if proto.Equal(feature.Location, point){
			return feature, nil
		}
	}
	return nil, nil
}

func main() {
	// 生成一個listener
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener a the address")
	}

	// server
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(lis))
}