package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/botless/slack/pkg/events"
	"github.com/kelseyhightower/envconfig"
	"github.com/knative/pkg/cloudevents"
	"log"
	"net/http"
	"os"
)

type envConfig struct {
	// Port is server port to be listened.
	Port string `envconfig:"PORT" default:"8081"`

	// Target where to send received event back to
	Target string `envconfig:"TARGET" default:"http://localhost:8080" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	os.Exit(_main(os.Args[1:], env))
}

type Loopback struct {
	ce0 *cloudevents.Client
	ce1 *cloudevents.Client
	i   int
}

func (e *Loopback) handler(ctx context.Context, msg json.RawMessage) {
	metadata := cloudevents.FromContext(ctx)
	_ = metadata

	log.Printf("Message: %s", string(msg))

	if e.i%2 == 0 {
		if err := e.ce0.Send(msg); err != nil {
			log.Printf("failed to send cloudevent: %s\n", err)
		}
	} else {
		if err := e.ce1.Send(msg); err != nil {
			log.Printf("failed to send cloudevent: %s\n", err)
		}
	}
	e.i++
}

func _main(args []string, env envConfig) int {
	e := &Loopback{
		ce0: cloudevents.NewClient(env.Target, cloudevents.Builder{
			EventTypeVersion: "v1alpha1",
			EventType:        events.ResponseEventType,
			Source:           "n3wscott.loopback",
			Encoding:         cloudevents.StructuredV01,
		}),
		ce1: cloudevents.NewClient(env.Target, cloudevents.Builder{
			EventTypeVersion: "v1alpha1",
			EventType:        events.ResponseEventType,
			Source:           "n3wscott.loopback",
			Encoding:         cloudevents.BinaryV01,
		}),
	}

	log.Printf("listening on port %s\n", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), cloudevents.Handler(e.handler)))
	return 0
}
