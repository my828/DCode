package main

import (
	"DCode/server/gateway/handlers"
	"DCode/server/gateway/sessions"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

// Director is a starter http.Request that will be used in CustomDirector
type Director func(r *http.Request)

// CustomDirector modifies the request object before forwarding it to the microservice
func CustomDirector() {
	// @TODO: redirects to microservice
}

// HeartBeatHandler is a handler to check if the dcode server is alive
func HeartBeatHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from DCode!"))
}

func main() {
	signingKey := os.Getenv("SIGNINGKEY")
	gatewayAddress := os.Getenv("GATEWAYADDRESS")
	if len(gatewayAddress) == 0 {
		gatewayAddress = ":4000"
	}
	// tlsCertificate := os.Getenv("TLSCERT")
	// tlsKey := os.Getenv("TLSKEY")
	redisAddress := os.Getenv("REDISADDRESS")

	redisDB := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	redisStore := sessions.NewRedisStore(redisDB, time.Hour*48)

	context := handlers.NewHandlerContext(signingKey, redisStore)
	socketstore := handlers.NewSocketStore() // will this create a new store everytime main runs? not sure if this is the right thing
	websocket := handlers.NewWebSocket(socketstore,context)
	router := mux.NewRouter()
	
	router.HandleFunc("/dcode", HeartBeatHandler)
	router.HandleFunc("/dcode/v1/new", context.NewSessionHandler)
	router.HandleFunc("/dcode/v1/{pageID}/extend", context.SessionExtensionHandler)
	// for websocket connections 
	router.HandleFunc("/ws/{pageID}", websocket.WebSocketConnectionHandler)
	// @TODO: redirect to microservice
	router.Handle("/dcode/v1/{pageID}", nil)
	// router.Handle("/dcode/v1/{pageID}/canvas", nil)
	// router.Handle("/dcode/v1/{pageID}/editor", nil)

	// adds CORS middleware around handlers
	cors := handlers.NewCORSHandler(router)

	log.Printf("Server is listening on port: %s\n", gatewayAddress)
	log.Fatal(http.ListenAndServe(gatewayAddress, cors))
	// log.Fatal(http.ListenAndServeTLS(gatewayAddress, tlsCertificate, tlsKey, cors))
}
  