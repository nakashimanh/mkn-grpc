package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"database/sql"


	"github.com/nakashimanh/mkn-grpc/grpc/driver"
	"github.com/nakashimanh/mkn-grpc/grpc/models"
	"github.com/nakashimanh/mkn-grpc/mikanpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}
var db *sql.DB

func (*server) Mikan(ctx context.Context, req *mikanpb.MikanRequest) (*mikanpb.MikanResponse, error) {
	fmt.Printf("Mikan function was invoked with %v\n", req)
	name := req.GetMikan().GetName()
	kind := req.GetMikan().GetKind()
	quality := req.GetMikan().GetQuality()
	result := "Response Mikan= " + name + " kind= " + kind + " quality= " + strconv.FormatInt(quality, 10)
	res := &mikanpb.MikanResponse{
		Result: result,
	}
	return res, nil
}

func (*server) RegisterMikan(ctx context.Context, req *mikanpb.RegisterMikanRequest) (*mikanpb.RegisterMikanResponse, error) {
	fmt.Printf("RegisterMikan function was invoked with %v\n", req)
	var mikan models.Mikan
	name := req.GetMikan().GetName()
	kind := req.GetMikan().GetKind()
	quality := req.GetMikan().GetQuality()
	result := "Response New Mikan = " + string(name) + " kind= " + kind + " quality= " + strconv.FormatInt(quality, 10)
	db = driver.ConnectDB()
	stmt := "insert into mikans (name, kind,quality) values($1, $2, $3) RETURNING id;"
	err := db.QueryRow(stmt, name, kind, quality).Scan(&mikan.ID)

	res := &mikanpb.RegisterMikanResponse{
		Result: result,
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

func main() {
	fmt.Println("Starting Mikan Service")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mikanpb.RegisterMikanServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
