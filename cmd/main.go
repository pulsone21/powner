package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/pulsone21/powner/internal/api"
)

var (
	protocol string
	url      string
	port     string
	db       string
	server   *http.Server
)

func init() {
	log.Println("Initializing")
	log.Println("Loading ENV variables")
	envFile, _ := godotenv.Read(".env")

	protocol = envFile["PROTOCOL"]
	url = envFile["URL"]
	port = envFile["PORT"]
	db = envFile["DB"]

	log.Printf("Creating Server add:'%v://%v:%v'\n", protocol, url, port)
	var err error
	server, err = api.CreateServer(protocol, url, port, db)
	if err != nil {
		panic(err)
	}

	log.Println("App Initialized")
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//dev := flag.Bool("d", true, "Running in dev mode, True or False")
	//
	//if *dev {
	//	log.Println("Running in dev mode, generating Data")
	//	api.GenerateData()
	//}

	log.Println("Starting Server:")

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
