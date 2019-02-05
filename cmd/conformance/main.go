package main

import (
	"fmt"
	"github.com/botless/slack/pkg/events"
	"github.com/kelseyhightower/envconfig"
	"github.com/n3wscott/cloudeventsconformace/pkg/conformance"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/knative/pkg/cloudevents"
)

type envConfig struct {
	// Port is server port to be listened.
	Port string `envconfig:"PORT" default:"8080"`

	// Target where to send received event back to
	Target string `envconfig:"TARGET" default:"http://localhost:8081" required:"true"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	os.Exit(_main(os.Args[1:], env))
}

type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

type Conformance struct {
	ce      *cloudevents.Client
	handler http.Handler
}

func (c *Conformance) Start() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		data := &Example{
			Message:  "hello, world!",
			Sequence: i,
		}

		if err := c.ce.Send(data, cloudevents.V01EventContext{
			Extensions: map[string]interface{}{
				"example": "example_ext",
			},
		}); err != nil {
			fmt.Printf("failed to send cloudevent: %v\n", err)
		}
	}
	os.Exit(0)
}

func _main(args []string, env envConfig) int {
	c := &Conformance{
		ce: cloudevents.NewClient(env.Target, cloudevents.Builder{
			EventTypeVersion: "v1alpha1",
			EventType:        events.ResponseEventType,
			Source:           "n3wscott.conformance",
		}),
		handler: conformance.NewHandler(),
	}
	go c.Start()

	log.Printf("listening on port %s\n", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), c.handler))
	return 0
}
