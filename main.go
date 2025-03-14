package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Mrhb787/hospital-ward-manager/common"
	"github.com/Mrhb787/hospital-ward-manager/configs"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

func main() {

	// service begin
	serviceStart := time.Now()

	// get app config
	appConfig := configs.New()

	// initalize services
	healthService := health.NewService()

	// database service
	dbService := database.NewService(appConfig.Host)

	// http handler
	handler := transport.GetHttpHandler(common.BASE, transport.HttpHandlerRequest{
		HealthService: healthService,
		DbService:     dbService,
	})
	h := handler.NewHandler()
	addr := fmt.Sprintf(":%s", appConfig.ServiceConfig.ServicePort)
	httpAddr := flag.String("http.addr", addr, "HTTP listen address")
	flag.Parse()

	// start server
	common.Go(func() {
		func() {
			err := http.ListenAndServe(*httpAddr, h)
			if err != nil {
				fmt.Println("Failed to start server!", err)
			}
		}()
	})
	fmt.Println("Service startup ended!", fmt.Sprintf("Startup time: %d", time.Since(serviceStart).Milliseconds()))
}
