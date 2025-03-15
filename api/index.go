package api

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/Mrhb787/hospital-ward-manager/configs"
	"github.com/Mrhb787/hospital-ward-manager/service/http/database"
	"github.com/Mrhb787/hospital-ward-manager/service/http/health"
	"github.com/Mrhb787/hospital-ward-manager/transport"
)

//go:embed main.html
var mainHTML []byte

var dbService database.Service

func init() {
	// app config
	appConfig := configs.New()

	// database service
	dbService := database.NewService(appConfig.Host, nil)

	// connect client
	client, err := dbService.Client()
	if err != nil || client == nil {
		log.Fatal("failed to connect database", err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		// Serve index.html
		serveMainHTML(w)
		return
	}

	// initalize services
	healthService := health.NewService()

	// http handler
	h := transport.NewHandler(transport.HttpHandlerRequest{
		HealthService: healthService,
		DbService:     dbService,
	})

	// serve http
	h.ServeHTTP(w, r)

}

func serveMainHTML(w http.ResponseWriter) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the embedded HTML to the response
	w.Write(mainHTML)
}
