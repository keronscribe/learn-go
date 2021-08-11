package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"net"

	pb "github.com/keronscribe/learn-go/grpc/route"
	"google.golang.org/grpc"
)

// ------ 實現 server stub 的 interface -------
// 定義 RouteGuideServer 的 type
type routeGuideServer struct {
	// 因為在這個interface 中有一個 mustEmbedUnimplementedRouteGuideServer()，實現向上兼容
	// 所以必須要在這裡有這個東西
	features                         []*pb.Feature // DB 要是一個裏面是 feature 的 Slice
	pb.UnimplementedRouteGuideServer               //內嵌進來

}

// 實現 RouteGuideServer 的方法
// Unary
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	// 在 proto 中定義的 Point 或者是 Feature 類型都會出現在生成的 pb stub 裏面，可以直接用他們
	// 這裡用最簡單的 for loop 進行遍歷
	for _, feature := range s.features {
		// 使用 proto 的 Equal 比對傳入的值和資料庫中的值
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

// Client-side Streaming
func (s *routeGuideServer) ListFeature(*pb.Rectangle, pb.RouteGuide_ListFeatureServer) error {
	return nil
}

//Client-side Streaming
func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

// bidirectional Streaming
func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

// ------ 返回我們定義的 RouteGuideServer  -------

func newServer() *routeGuideServer {
	// 回傳他的 address
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
