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
	Target string `envconfig:"TARGET" default:"http://localhost:8080x" required:"true"`
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
	ce *cloudevents.Client
}

func (e *Loopback) handler(ctx context.Context, msg json.RawMessage) {
	metadata := cloudevents.FromContext(ctx)
	_ = metadata

	log.Printf("Message: %s", string(msg))

	if err := e.ce.Send(msg); err != nil {
		log.Printf("failed to send cloudevent: %s\n", err)
	}
}

func _main(args []string, env envConfig) int {
	e := &Loopback{
		ce: cloudevents.NewClient(env.Target, cloudevents.Builder{
			EventTypeVersion: "v1alpha1",
			EventType:        events.ResponseEventType,
			Source:           "n3wscott.loopback",
		}),
	}

	log.Printf("listening on port %s\n", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), cloudevents.Handler(e.handler)))
	return 0
}
