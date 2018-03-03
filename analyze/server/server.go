package main

import (
	"fmt"
	pb "stock/analyze/proto"
	"stock/etc"
	"stock/rpc"

	"golang.org/x/net/context"
)

var (
	port = etc.Int("service", "analyze")
)

type server struct {
	*db
}

func newServer() *server {
	db := NewDb()
	return &server{
		db: db,
	}
}

func (s *server) CodeNameList(ctx context.Context, req *pb.CodeNameListReq) (*pb.CodeNameListResp, error) {
	codeName, err := s.db.codeNameList(req.Sybmol)
	if err != nil {
		fmt.Println("code name error ==>", err)
		return nil, err
	}
	return &pb.CodeNameListResp{Codenames: codeName}, nil
}

func (s *server) GrowingCompare(ctx context.Context, req *pb.GrowingCompareReq) (*pb.GrowingCompareResp, error) {
	growths, err := s.db.GrowingCompare(req.Code)
	if err != nil {
		fmt.Println("growing compare error ==>", err)
		return nil, err
	}
	return &pb.GrowingCompareResp{Growths: growths}, nil
}

func (s *server) ProfitCompare(ctx context.Context, req *pb.ProfitCompareReq) (*pb.ProfitCompareResp, error) {
	profits, err := s.db.ProfitCompare(req.Code)
	if err != nil {
		fmt.Println("profits compare error ==>", err)
		return nil, err
	}
	return &pb.ProfitCompareResp{Profits: profits}, nil
}

func (s *server) Close() error {
	return s.db.Close()
}

func main() {
	s := rpc.NewServer()
	pb.RegisterAnalyzeServiceServer(s.Server, newServer())
	s.Serve(int(port))
}
