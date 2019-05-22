package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gophersises/02_url_shortener/handlers"
	"gophersises/02_url_shortener/storage"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	LogLevel string `default:"DEBUG"`

	Port          int           `required:"true"`
	GracefullWait time.Duration `default:"15s"`

	UseDB bool `default:"true"`

	RedisAddr      string        `default:"localhost:6379"`
	RedisPassword  string        `default:""`
	RedisDB        int           `default:"0"`
	RedisRWTimeout time.Duration `default:"5s"`
}

func main() {
	var c config
	err := envconfig.Process("urlshort", &c)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("CONFIG: %+v", c)

	err = storage.Init(c.UseDB, c.RedisAddr, c.RedisPassword, c.RedisDB, c.RedisRWTimeout)
	if err != nil {
		log.Fatal(err)
	}

	router := newRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf(`:%d`, c.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), c.GracefullWait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Printf("grasefully shutted down")
	os.Exit(0)
}

func newRouter() *mux.Router {
	r := mux.NewRouter() //.StrictSlash(true)

	r.HandleFunc("/new", handlers.NewHandler).Methods("GET", "POST")
	r.HandleFunc("/{hash}", handlers.IndexHandler).Methods("GET")

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	r.Use(handlers.PanicMiddleware)
	r.Use(handlers.RequestLoggingMiddleware)

	return r
}
