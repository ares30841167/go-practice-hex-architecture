package rpc

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"guanyu.dev/arithmetic/internal/adapters/app/api"
	"guanyu.dev/arithmetic/internal/adapters/core/arithmetic"
	"guanyu.dev/arithmetic/internal/adapters/framework/left/grpc/pb"
	"guanyu.dev/arithmetic/internal/adapters/framework/right/db"
	"guanyu.dev/arithmetic/internal/ports"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	var dbAdapter ports.DbPort
	var arithAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort

	dbDriverName := os.Getenv("DB_DRIVER_NAME")
	dbSourceName := os.Getenv("DB_SOURCE_NAME")

	dbAdapter, err = db.NewAdapter(dbDriverName, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}

	arithAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, arithAdapter)
	grpcAdapter := NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer.GetValue(), int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer.GetValue(), int32(0))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer.GetValue(), int32(1))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v", err)
	}

	require.Equal(t, answer.GetValue(), int32(1))
}
