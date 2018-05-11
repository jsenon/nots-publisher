//go:generate swagger generate spec
// Package main demoserver.
//
// the purpose of this application is to provide an CMDB application
// that will store information in mongodb backend
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package main

import (
	"net/http"
	"time"

	"github.com/jsenon/nats-publisher/web"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Failed to create zap logger",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
	defer logger.Sync() // nolint: errcheck

	// Web Part
	http.HandleFunc("/publish", web.Publish)

	// // API Part
	// http.HandleFunc("/healthz", api.Health)
	// http.HandleFunc("/.well-known", api.Wellknown)
	// http.HandleFunc("/play", a.Play)
	// http.HandleFunc("/ping", api.Pong)

	// If no Jaeger variable set we don't propagate header
	err = http.ListenAndServe(":9010", nil)
	if err != nil {
		logger.Error("Failed to start web server",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
	}
}
