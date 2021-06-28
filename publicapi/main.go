package main

import (
	"github.com/m3o/services/pkg/tracing"
	"github.com/m3o/services/publicapi/handler"
	pb "github.com/m3o/services/publicapi/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("publicapi1"),
		service.Version("latest"),
	)

	store.DefaultStore.Init(store.Table("publicapi"))

	// Register handler
	pb.RegisterPublicapiHandler(srv.Server(), handler.NewPublicAPIHandler(srv))
	pb.RegisterExploreHandler(srv.Server(), handler.NewExploreAPIHandler(srv))
	traceCloser := tracing.SetupOpentracing("publicapi")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
