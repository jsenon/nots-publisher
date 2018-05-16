// Package web demoserver.
//
// the purpose of this package is to provide Web HTML Interface
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
package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nats-io/nats"
	"go.uber.org/zap"
)

// Message struct defin body sent by client
type Message struct {
	Nats    string `json:"nats"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// Publish func to display all server on table view
func Publish(res http.ResponseWriter, req *http.Request) {

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("Failed to create zap logger",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return
	}
	defer logger.Sync() // nolint: errcheck
	// var rs Server

	// Retrieve Body

	var msg Message

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("Failed to read body",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return
	}

	err = json.Unmarshal(body, &msg) // nolint: gas
	if err != nil {
		logger.Error("Failed to unmarshal",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return
	}

	url := msg.Nats
	subject := msg.Subject
	message := msg.Message

	// Connect to server; defer close
	natsConnection, err := nats.Connect(url)
	if err != nil {
		logger.Error("Failed to connect",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return
	}
	defer natsConnection.Close()
	logger.Info("Connected",
		zap.String("target", url),
		zap.String("ServerID", natsConnection.ConnectedServerId()),
		zap.String("ConnectedServer", natsConnection.ConnectedUrl()),
	)

	// Publish message on subject
	err = natsConnection.Publish(subject, []byte(message))
	if err != nil {
		logger.Error("Failed to publish",
			zap.String("status", "ERROR"),
			zap.Int("statusCode", 500),
			zap.Duration("backoff", time.Second),
			zap.Error(err),
		)
		return
	}
	logger.Info("Publish",
		zap.String("topic", subject),
	)
}
