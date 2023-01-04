package main

import (
	"context"
	"ercd-test/cmd/back/handler"
	"ercd-test/internal/conf"
	"ercd-test/internal/logger"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := flag.String("port", "8090", "Http server Default port is 8090")
	flag.Parse()

	addr := fmt.Sprintf("0.0.0.0:%s", *port)

	conf.InitYaml()

	// router
	router := gin.Default()

	// handler
	h, err := handler.NewHandler(&handler.HConfig{
		R: router,
	})
	if err != nil {
		logger.Logrus.Fatalf("Failed in handler: %v", err)
	}

	// register router
	h.Register()

	// new http server instance
	server := http.Server{
		Addr:           addr,
		Handler:        h.R,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	// initialized server
	go func() {
		logger.Logrus.Println("Http server listening on port: ", *port)
		if err := server.ListenAndServe(); err != nil {
			logger.Logrus.Fatalf("Failed to initialized server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	// graceful shutdown
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Logrus.Fatalf("failed to shutdown http server: %v", err)
	}
}
