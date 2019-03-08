package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	gatewayAddress := os.Getenv("GATEWAYADDRESS")
	// tlsCertificate := os.Getenv("TLSCERT")
	// tlsKey := os.Getenv("TLSKEY")

	// redisAddress := os.Getenv("REDISADDRESS")

	router := mux.NewRouter()

	// @TODO: Replace with actual handlers
	router.Handle("/v1/pages", nil)
	router.Handle("v1/pages/{pageID}", nil)
	router.Handle("v1/pages/{pageID}/canvas", nil)
	router.Handle("v1/pages/{pageID}/editor", nil)

	// adds CORS middleware around handlers
	// cors := handlers.NewCORSHandler(router)

	log.Printf("Server is listening on port: %s\n", gatewayAddress)
	// log.Fatalf(http.ListenAndServeTLS(gatewayAddress, tlsCertificate, tlsKey, cors))
}
