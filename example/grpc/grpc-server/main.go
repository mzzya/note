package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"

	"blog/example/grpc/dbmodel"
	"blog/example/grpc/pb"

	"github.com/pkg/errors"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123123@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	db.DB().SetConnMaxLifetime(time.Hour)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	err = db.DB().Ping()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	g := &dbmodel.Goods{}
	if !db.HasTable(g) {
		err = db.CreateTable(g).Error
		if err != nil {
			panic(err)
			os.Exit(1)
		}
	}
}

func Convert(in interface{}, out interface{}) error {
	bts, err := json.Marshal(in)
	if err != nil {
		return errors.WithMessage(err, "json_marshal")
	}
	err = json.Unmarshal(bts, out)
	if err != nil {
		return errors.WithMessage(err, "json_unmarshal")
	}
	return nil
}

// GoodsServer 商品服务
type GoodsServer struct{}

// Create 创建商品
func (s *GoodsServer) Create(ctx context.Context, in *pb.Goods) (*pb.BaseResponse, error) {
	g := &dbmodel.Goods{}
	err := Convert(in, g)
	if err != nil {
		return nil, errors.WithMessage(err, "convert")
	}
	err = db.Create(g).Error
	if err != nil {
		return nil, errors.WithMessage(err, "create")
	}
	return &pb.BaseResponse{Status: true}, nil
}

// Update 更新商品
func (s *GoodsServer) Update(ctx context.Context, in *pb.Goods) (*pb.BaseResponse, error) {
	g := &dbmodel.Goods{}
	err := Convert(in, g)
	if err != nil {
		return nil, errors.WithMessage(err, "convert")
	}
	err = db.Model(g).Updates(g).Error
	if err != nil {
		return nil, errors.WithMessage(err, "update")
	}
	return &pb.BaseResponse{Status: true}, nil
}

// Delete 删除商品
func (s *GoodsServer) Delete(ctx context.Context, in *pb.Goods) (*pb.BaseResponse, error) {
	g := &dbmodel.Goods{}
	err := Convert(in, g)
	if err != nil {
		return nil, errors.WithMessage(err, "convert")
	}
	err = db.Delete(g).Error
	if err != nil {
		return nil, errors.WithMessage(err, "delete")
	}
	return &pb.BaseResponse{Status: true}, nil
}

// Get 根据ID获取商品
func (s *GoodsServer) Get(ctx context.Context, in *pb.IDRequest) (*pb.GoodsResponse, error) {
	g := &dbmodel.Goods{}
	err := db.Where("id = ?", in.ID).First(g).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, errors.WithMessage(err, "db_first")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	result := &pb.GoodsResponse{}
	err = Convert(g, &result.Data)
	if err != nil {
		return nil, errors.WithMessage(err, "convert")
	}
	result.BaseResponse = &pb.BaseResponse{Status: true}
	return result, nil
}

// List 分页获取商品
func (s *GoodsServer) List(ctx context.Context, in *pb.PageRequest) (*pb.GoodsListResponse, error) {
	goods := make([]dbmodel.Goods, 0, in.PageSize)
	err := db.Limit(in.PageSize).Offset((in.PageIndex - 1) * in.PageSize).Find(&goods).Error
	if err != nil {
		return nil, errors.WithMessage(err, "db_find")
	}
	result := &pb.GoodsListResponse{
		Data: make([]*pb.Goods, 0, len(goods)),
	}
	err = Convert(goods, result.Data)
	if err != nil {
		return nil, errors.WithMessage(err, "convert")
	}
	result.BaseResponse = &pb.BaseResponse{Status: true}
	return result, nil
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoodsSvcServer(s, &GoodsServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
