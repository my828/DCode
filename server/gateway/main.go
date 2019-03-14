package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"DCode/server/gateway/handlers"

	"DCode/server/gateway/sessions"

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
	// tlsCertificate := os.Getenv("TLSCERT")
	// tlsKey := os.Getenv("TLSKEY")
	redisAddress := os.Getenv("REDISADDRESS")
	mqAddress := os.Getenv("RABBITMQADDRESS")
	mqName := os.Getenv("RABBITMQNAME")

	// connect to Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	redisStore := sessions.NewRedisStore(redisDB, time.Hour*48)

	socketStore := handlers.NewSocketStore()
	rabbitStore := handlers.NewRabbitStore(mqAddress, mqName)

	context := handlers.NewHandlerContext(signingKey, redisStore, socketStore, rabbitStore)
	websocket := handlers.NewWebSocket(socketStore, context)

	router := mux.NewRouter()

	router.HandleFunc("/dcode", HeartBeatHandler)
	router.HandleFunc("/dcode/v1/new", context.NewSessionHandler)
	router.HandleFunc("/dcode/v1/{pageID}/extend", context.SessionExtensionHandler)
	// for websocket connections
	router.HandleFunc("/ws/{pageID}", websocket.WebSocketConnectionHandler)
	// @TODO:
	router.Handle("/dcode/v1/{pageID}", nil)

	// adds CORS middleware around handlers
	cors := handlers.NewCORSHandler(router)

	log.Printf("Server is listening on port: %s\n", gatewayAddress)
	log.Fatal(http.ListenAndServe(gatewayAddress, cors))
	// log.Fatal(http.ListenAndServeTLS(gatewayAddress, tlsCertificate, tlsKey, cors))
}

func createRabbitChannel(mqAddress string, mqName string) {

}
