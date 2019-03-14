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
	mqAddress := os.Getenv("RABBITADDRESS")
	mqName := os.Getenv("RABBITNAME")

	// connect to Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})
	redisStore := sessions.NewRedisStore(redisDB, time.Hour*48)

	rabbitStore := handlers.NewRabbitStore(mqAddress, mqName)
	socketStore := handlers.NewSocketStore(rabbitStore, redisStore)
	messagesChannel := rabbitStore.Consume()
	go socketStore.Notify(messagesChannel)

	context := handlers.NewHandlerContext(signingKey, redisStore, socketStore)
	websocket := handlers.NewWebSocket(context)

	router := mux.NewRouter()
	router.HandleFunc("/dcode", HeartBeatHandler)
	router.HandleFunc("/ws/{pageID}", websocket.WebSocketConnectionHandler)
	router.HandleFunc("/dcode/v1/new", context.NewSessionHandler)
	router.HandleFunc("/dcode/v1/{pageID}/extend", context.SessionExtensionHandler)
	router.HandleFunc("/dcode/v1/{pageID}", context.GetPageHandler)

	// go ss.write

	// adds CORS middleware around handlers
	cors := handlers.NewCORSHandler(router)

	log.Printf("Server is listening on port: %s\n", gatewayAddress)
	log.Fatal(http.ListenAndServe(gatewayAddress, cors))
	// log.Fatal(http.ListenAndServeTLS(gatewayAddress, tlsCertificate, tlsKey, cors))
}

func createRabbitChannel(mqAddress string, mqName string) {

}
