package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"GO_PROJECT/config"
	"GO_PROJECT/db"
	"GO_PROJECT/kafka"
	"GO_PROJECT/logger"
	indexRoute "GO_PROJECT/routes"
	"GO_PROJECT/utils"
)

var exCh = make(chan os.Signal, 3)
var wg sync.WaitGroup

func initialize() {

	fmt.Println("Initializing a Server ...")

	// Load .env file
	godotenv.Load()

	// Handle signals gracefully
	signal.Notify(exCh, syscall.SIGINT)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt) // Listen for SIGINT

	wg.Add(1)
	go func() {
		for {
			defer wg.Done()
			<-done
			// Close resources and terminate gracefully
			fmt.Println("Received Ctrl+C, shutting down...")
			return
		}
	}()

	fmt.Println("Signal Handling Done Sucessfully")
}

func readConfig() {
	config.GetConfigurations()
	fmt.Println("Configuration Read Sucessfully")

}

func makeConnections() {
	db.ConnectDatabases()
	fmt.Println("Connected To Databases Sucessfully")
	kafka.InitialiseKafka()
	fmt.Println("Kafka Connected Successfully")
}

func handleRoutes() {
	//configure http router
	r := mux.NewRouter()
	indexRoute.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println("Routes Handled Sucessfully")
}

func main() {

	// init
	initialize()

	// config read
	readConfig()

	// initialize logger
	logger.Setup(os.Getenv("ENV"))

	// connections
	makeConnections()

	// handle routes
	handleRoutes()

	// start server
	go func() {
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
		fmt.Println("Server is running on port", os.Getenv("PORT"))
	}()

	utils.StartKafkaConsumer()

	// Wait for all goroutines to finish
	wg.Wait()
	// <-exCh
	time.Sleep(1 * time.Second) //waiting for the queued task to complete
	fmt.Println("Server is closing ...")
}
