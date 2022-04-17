package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/handlers"
	"time"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", handlers.NewProducts(l))
	sm.Handle("/goodbye", handlers.NewGoodbye(l))
	server := http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func(){
		err:= server.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Graceful shutdown", sig)
	server.ListenAndServe()
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(tc)
}