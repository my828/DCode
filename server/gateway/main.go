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

	context := handlers.NewHandlerContext(signingKey, redisStore, socketStore)
	websocket := handlers.NewWebSocket(context)

	router := mux.NewRouter()
	router.HandleFunc("/dcode", HeartBeatHandler)
	router.HandleFunc("/ws/{pageID}", websocket.WebSocketConnectionHandler)
	router.HandleFunc("/dcode/v1/new", context.NewSessionHandler)
	router.HandleFunc("/dcode/v1/{pageID}/extend", context.SessionExtensionHandler)
	router.HandleFunc("/dcode/v1/{pageID}", context.GetPageHandler)

	// adds CORS middleware around handlers
	cors := handlers.NewCORSHandler(router)

	messagesChannel := rabbitStore.Consume()
	log.Println(messagesChannel)
	if messagesChannel != nil {
		go socketStore.Notify(messagesChannel)
	}

	log.Printf("Server is listening on port: %s\n", gatewayAddress)
	log.Fatal(http.ListenAndServe(gatewayAddress, cors))
	// log.Fatal(http.ListenAndServeTLS(gatewayAddress, tlsCertificate, tlsKey, cors))
}
