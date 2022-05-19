package rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"guanyu.dev/arithmetic/internal/adapters/framework/left/grpc/pb"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, err.Error())
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
