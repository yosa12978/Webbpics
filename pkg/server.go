package pkg

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	r := NewRouter()
	server := http.Server{
		Addr:    ":8089",
		Handler: r,
	}
	go server.ListenAndServe()
	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	stat := <-out
	log.Println("Server shut down with signal", stat.String())
}
