package main

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"blog/example/grpc/pb"
)

func TestGoodsServer(t *testing.T) {
	gs := GoodsServer{}
	Convey("GoodsServer", t, func() {
		//result, err := gs.Create(context.Background(), &pb.Goods{
		//	ID:         1,
		//	Name:       "aaa",
		//	Bn:         "bbb",
		//	Price:      10,
		//	Pic:        "http://img.alicdn.com/imgextra/i1/735276822/TB2.bqVjfxNTKJjy0FjXXX6yVXa_!!735276822-2-beehive-scenes.png",
		//	Content:    "河马男朋友抱枕靠枕床头靠垫大靠背睡觉枕头大号床上长条枕可爱女",
		//	CreateTime: time.Now().Unix(),
		//	UpdateTime: 0,
		//})
		//So(err, ShouldEqual, nil)
		//t.Logf("%#v\n", result)
		//g, err := gs.Get(context.Background(), &pb.IDRequest{ID: 1})
		//So(err, ShouldEqual, nil)
		//t.Logf("%#v", g)
		//g.Data.UpdateTime = time.Now().Unix()
		//result, err = gs.Update(context.Background(), g.Data)
		//So(err, ShouldEqual, nil)
		//t.Logf("%#v\n", result)
		//result, err = gs.Delete(context.Background(), &pb.Goods{ID: 1})
		//So(err, ShouldEqual, nil)
		//t.Logf("%#v\n", result)
		g, err := gs.List(context.Background(), &pb.PageRequest{PageIndex: 1, PageSize: 10})
		So(err, ShouldEqual, nil)
		t.Logf("%#v", g)
	})

}
