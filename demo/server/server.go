package route

import (
	"context"
	pb "github.com/keronscribe/learn-go/demo/route" // 透過 `proto` 生成的 Server Stub
	"google.golang.org/grpc"                        // 新增 grpc library
	"google.golang.org/protobuf/proto"
	"log" // 新增
	"net" // 新增
)

type routeGuideServer struct {
	features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
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

func (s *routeGuideServer) GetFeature(cxt context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(listen))
}
