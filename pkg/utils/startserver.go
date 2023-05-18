package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// Starting server without graceful shutdown
func StartServer(a *fiber.App, url string, l *log.Logger) {
	if err := a.Listen(url); err != nil {
		l.Println("Error starting server: ", err)
	}
}

// Starting server with graceful shutdown
func StartServerWithGracefulShutdown(a *fiber.App, url string, l *log.Logger) { // TODO fix this

	l.Printf("Starting server on %s",  url)

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt)

		<-sigint

		l.Println("Received interrupt signal, gracefully shutting down server")

		err := a.Shutdown()

		if err != nil {
			l.Println("Error shutting down server: ", err)
		}

		close(idleConnsClosed)
	}()

	err := a.Listen(url)

	if err != nil {
		l.Println("Error starting server: ", err)
	}

	<-idleConnsClosed
}
