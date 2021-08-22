package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/keronscribe/learn-go/grpc/route" // 透過 `proto` 生成的 Client Stub
	"google.golang.org/grpc" // grpc 包
)

// Unary
// 傳一個 request，拿一個 response
func getFeat (client pb.RouteGuideClient) {
	feature,err := client.GetFeature(context.Background(), &pb.Point{
		Latitude: 353931000,
		Longitude: 139444400,
	})

	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

// Client-side streaming
// 傳一個 request，拿多個 response
// 好處是後端可以找到一個回傳一個，不需要全部找完一起回傳
func getListFeatures (client pb.RouteGuideClient){
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  355000000,
			Longitude: 138000000,

		},
		Hi: &pb.Point{
			Latitude:  360000000,
			Longitude: 140000000,

		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		feature, err := serverStream.Recv()
		if err == io.EOF {
			// stream 結束
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

func main ()  {
	// Dial 會是一個撥請求，第一個參數會說他要播向哪裡，之後是一些 option
	// grpc.WithInsecure() -> 因為現在 server 端沒有提供證書，所以需要使用 Insecure 來跳過驗證
	// grpc.WithBlock() -> 如果沒有成功就不讓他往下走的一個選項
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("Client cannot dail grpc server")
	}

	defer conn.Close() // 最後的時候要關掉 conn

	// 拿到 Stub(這裡被重新命名為 pb) 中生成的 NewRouteGuideClient
	client := pb.NewRouteGuideClient(conn)
	// 像是調用本地函數一樣的去調用這個 GetFeature
	// 其實中間有完成 server 的互動

	//getFeat(client)
	getListFeatures(client)
}

