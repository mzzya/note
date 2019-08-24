package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"blog/example/grpc/pb"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gs := pb.NewGoodsSvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = gs.Create(ctx, &pb.Goods{
		ID:         1,
		Name:       "aaa",
		Bn:         "bbb",
		Price:      10,
		Pic:        "http://img.alicdn.com/imgextra/i1/735276822/TB2.bqVjfxNTKJjy0FjXXX6yVXa_!!735276822-2-beehive-scenes.png",
		Content:    "河马男朋友抱枕靠枕床头靠垫大靠背睡觉枕头大号床上长条枕可爱女",
		CreateTime: time.Now().Unix(),
		UpdateTime: 0,
	})
	if err != nil {
		log.Printf("create_error:%s\n", err)
	}
	g, err := gs.Get(ctx, &pb.IDRequest{ID: 1})
	if err != nil {
		log.Printf("create_error:%s\n", err)
	}
	g.Data.UpdateTime = time.Now().Unix()
	_, err = gs.Update(ctx, g.Data)
	if err != nil {
		log.Printf("create_error:%s\n", err)
	}
	g, err = gs.Get(ctx, &pb.IDRequest{ID: 1})
	if err != nil {
		log.Printf("create_error:%s\n", err)
	}
	log.Printf("%#v\n", g.Data)
	list, err := gs.List(ctx, &pb.PageRequest{PageIndex: 1, PageSize: 2})
	if err != nil {
		log.Printf("create_error:%s\n", err)
	}
	log.Printf("List:%#v\n", list)
}
