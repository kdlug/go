// graceful shutdown
// ctrl + C
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting")
	defer log.Println("ended")

	// creates a channel for receiving system signals
	shutdown := make(chan os.Signal, 1)
	// configures the shutdown channel to receive SIGINT and SIGTERM signals
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	// "<-shutdown" blocks execution, waiting for a signal to be received
	<-shutdown

	log.Println("stopping")
}
