package analyze

import (
	"log"
	pb "stock/analyze/proto"
	"stock/etc"
	"stock/rpc"

	"golang.org/x/net/context"
)

var host = etc.String("host", "analyze")

var client = pb.NewAnalyzeServiceClient(rpc.NewClient(host))

//
func CodeNameList(ctx context.Context, symbol string) ([]*pb.CodeName, error) {
	resp, err := client.CodeNameList(ctx, &pb.CodeNameListReq{
		Sybmol: symbol,
	})
	if err != nil {
		return nil, err
	}
	return resp.Codenames, nil
}

//
func GrowingCompare(ctx context.Context, code string) ([]*pb.Growth, error) {
	resp, err := client.GrowingCompare(ctx, &pb.GrowingCompareReq{
		Code: code,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp.Growths, nil
}

//
func ProfitCompare(ctx context.Context, code string) ([]*pb.Profit, error) {
	resp, err := client.ProfitCompare(ctx, &pb.ProfitCompareReq{
		Code: code,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp.Profits, nil
}
