package main

import (
	"log"
	"os"

	"guanyu.dev/arithmetic/internal/adapters/app/api"
	"guanyu.dev/arithmetic/internal/adapters/core/arithmetic"
	rpc "guanyu.dev/arithmetic/internal/adapters/framework/left/grpc"
	"guanyu.dev/arithmetic/internal/adapters/framework/right/db"
	"guanyu.dev/arithmetic/internal/ports"
)

func main() {
	var err error

	var dbAdapter ports.DbPort
	var arithAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbDriverName := os.Getenv("DB_DRIVER_NAME")
	dbSourceName := os.Getenv("DB_SOURCE_NAME")

	dbAdapter, err = db.NewAdapter(dbDriverName, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	arithAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, arithAdapter)
	grpcAdapter = rpc.NewAdapter(appAdapter)

	grpcAdapter.Run()
}
