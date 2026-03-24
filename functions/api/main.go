package main

import (
	"log"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gkatanacio/go-lambdalith-template/internal/sample"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()

	sampleService := sample.NewService(sample.ConfigFromEnv())
	sampleController := sample.NewController(sampleService)

	mux.HandleFunc("GET /hello", sampleController.Hello)
	mux.HandleFunc("POST /echo", sampleController.Echo)
}

func main() {
	csrfProtectedMux := http.NewCrossOriginProtection().Handler(mux)

	if os.Getenv("APP_ENV") == "local" {
		server := &http.Server{
			Addr:    ":8080",
			Handler: csrfProtectedMux,
		}
		log.Printf("listening on port %s", server.Addr)
		log.Fatal(server.ListenAndServe())
	}

	handler := algnhsa.New(csrfProtectedMux, nil)

	lambda.Start(handler)
}
