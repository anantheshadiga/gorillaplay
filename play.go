package gorillaplay

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anantheshadiga/gorillaplay/hang"
	"github.com/anantheshadiga/gorillaplay/jump"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logo = `
                  ."'".
              .-./ _=_ \.-.
             {  (,(oYo),) }}
             {{ |   "   |} }
             { { \(---)/  }}
             {{  }'-=-'{ } }
             { { }._:_.{  }}
             {{  } -:- { } }
             {_{ }'==='{  _}
            ((((\)     (/))))
`

type Gorilla struct {
	server *http.Server
}

func (g *Gorilla) Start() (err error) {
	fmt.Println(logo)

	var zapConfig = zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	zapConfig.OutputPaths = []string{"stdout"}
	logger, logError := zapConfig.Build()
	if logError != nil {
		return errors.New("Error while initializing logger")
	}

	r := mux.NewRouter()
	r.HandleFunc("/gorilla", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(logo)) })
	r.Handle("/jump", jump.NewHandler(logger))
	r.Handle("/hang", hang.NewHandler(logger))
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:4242",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  100 * time.Second,
		IdleTimeout:  100 * time.Second,
	}

	g.server = srv

	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}(srv)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	err = g.Close()
	if err != nil {
		log.Fatal("Error Closing Server: ", err.Error())
	}

	os.Exit(9)
	return nil
}

func (g *Gorilla) Close() error {
	// Context shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(), 300*time.Second,
	)

	defer cancel()

	return g.server.Shutdown(ctx)
}

func New() *Gorilla {
	return &Gorilla{}
}
