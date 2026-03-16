package main

import (
	"net/http"

	"github.com/akrylysov/algnhsa"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gkatanacio/go-lambdalith-template/internal/sample"
)

var router *http.ServeMux

func init() {
	router = http.NewServeMux()

	sampleService := sample.NewService(sample.ConfigFromEnv())
	sampleController := sample.NewController(sampleService)

	router.HandleFunc("GET /hello", sampleController.Hello)
	router.HandleFunc("POST /echo", sampleController.Echo)
}

func main() {
	handler := algnhsa.New(router, nil)

	lambda.Start(handler)
}
