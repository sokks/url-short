package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	"github.com/sokks/url-short/handlers"
	"github.com/sokks/url-short/storage"
)

type config struct {
	LogLevel   string `default:"DEBUG"`
	LogFile    string `default:""`
	LogFullReq bool   `default:"false"`

	Port          int           `required:"true"`
	GracefullWait time.Duration `default:"15s"`

	UseDB bool `default:"true"`

	RedisAddr      string        `default:""`
	RedisPassword  string        `default:""`
	RedisDB        int           `default:"0"`
	RedisRWTimeout time.Duration `default:"5s"`

	BaseURL        string `default:""`
	MaxRehashTries int    `default:"5"`
}

func main() {
	var c config
	err := envconfig.Process("urlshort", &c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CONFIG: %+v", c)

	if c.LogFile != "" {
		f, err := os.OpenFile(c.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	fmt.Println("started at ", time.Now())

	err = storage.Init(c.UseDB, c.RedisAddr, c.RedisPassword, c.RedisDB, c.RedisRWTimeout)
	if err != nil {
		log.Fatal(err)
	}
	handlers.Init(c.BaseURL, c.MaxRehashTries)

	router := newRouter(c.LogFullReq)
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
		log.Println("started at ", time.Now())
	}()

	// graceful shutdown
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), c.GracefullWait)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to shutdown properly, err: %s", err)
		os.Exit(1)
	}
	log.Printf("grasefully shutted down")
	os.Exit(0)
}

func newRouter(useRequestLoggingMiddlware bool) *mux.Router {
	r := mux.NewRouter() //.StrictSlash(true)

	r.HandleFunc("/new", handlers.NewHandler).Methods("POST")
	r.HandleFunc("/{hash}", handlers.IndexHandler).Methods("GET")

	r.Use(handlers.PanicMiddleware)
	if useRequestLoggingMiddlware {
		r.Use(handlers.RequestLoggingMiddleware)
	}

	return r
}
