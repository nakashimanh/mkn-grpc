package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/nakashimanh/mkn-grpc/grpc/driver"
	"github.com/nakashimanh/mkn-grpc/grpc/jwt"
	"github.com/nakashimanh/mkn-grpc/grpc/models"
	"github.com/nakashimanh/mkn-grpc/proto/mikanpb"
	"github.com/nakashimanh/mkn-grpc/proto/operator_auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

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

func (*server) OperatorAuth(ctx context.Context, req *operator_auth.OperatorAuthRequest) (*operator_auth.OperatorAuthResponse, error) {
	fmt.Printf("OperatorAuth function was invoked with %v\n", req)
	id := req.GetAuth().GetID()
	deviceID := req.GetAuth().GetDeviceID()
	fmt.Println(id)
	newJWT, err := jwt.GenerateToken(id, deviceID)
	res := &operator_auth.OperatorAuthResponse{
		Jwt:    newJWT,
		Result: "Success",
	}
	return res, err
}

// Middleware
func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		// Skip authorize when GetJWT is requested
		if info.FullMethod != "/operator_auth.OperatorAuthService/OperatorAuth" {
			if err := jwt.TokenVerifyMiddleWare(ctx); err != nil {
				log.Printf("auth err = %v\n", err)
				return nil, err
			}
		}

		resp, err := handler(ctx, req)
		// Logging with grpclog (grpclog.LoggerV2)
		grpclog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(start),
			err)
		return resp, err
	}
}

func main() {
	fmt.Println("Starting Mikan Project Service")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//s := grpc.NewServer()
	s := grpc.NewServer(grpc.UnaryInterceptor(
		unaryServerInterceptor()))

	mikanpb.RegisterMikanServiceServer(s, &server{})
	operator_auth.RegisterOperatorAuthServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
