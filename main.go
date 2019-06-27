package main

import (
	"./server"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Trainer sdsd
type Trainer struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Item   string `json:"item"`
	Qty    int    `json:"qty"`
	Status string `json:"status"`
	Size   struct {
		H   int    `json:"h"`
		W   int    `json:"w"`
		Uom string `json:"uom"`
	} `json:"size"`
	Tags []string `json:"tags"`
}

// Logger ddd
//func Logger() http.Handler {
//    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//        fmt.Println(time.Now(), r.Method, r.URL)
//        router.ServeHTTP(w, r) // dispatch the request
//    })
//}

func main() {
	// This is the domain the server should accept connections for.

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

collection := client.Database("test").Collection("inventory")
// create a value into which the result can be decoded
var result Trainer

err = collection.Find(context.TODO(), filter).Decode(&result)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found a single document: %+v\n", result)
	handler := server.NewRouter()
	// http.ListenAndServe( ":443", handler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		srv.ListenAndServe()

	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

}
