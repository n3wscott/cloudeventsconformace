package canonical

import (
	"encoding/json"
	"fmt"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical/http"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical/v01"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical/v02"
	"log"
	"strings"
)

type PrintVersion interface {
	Print(b *strings.Builder)
}

func Print(r http.Request) {
	var b strings.Builder

	v := Categorize(r)

	if v != nil {
		v.Print(&b)
	}

	fmt.Println(b.String())
}

type Versioner struct {
	CloudEventsVersion string `json:"cloudEventsVersion, omitempty"` // version 0.1
	SpecVersion        string `json:"specversion, omitempty"`        // version 0.2
}

// We will categorize based on newest version first for matching.
func Categorize(r http.Request) PrintVersion {

	log.Print("Headers: %+v", r.Header)
	// check for v0.2
	contentType := r.Header.Get("content-type")

	if strings.HasPrefix(contentType, "application/cloudevents+json") {
		// then this should be a Structured Json HTTP cloudevent.
		// it might have headers but they are suppose to be redundant to the body.
		version := &Versioner{}
		if err := json.Unmarshal(r.Body, version); err != nil {
			log.Printf("failed to find version: %s", err)
		}

		if version.SpecVersion == "0.2" {
			log.Println("Found HTTP Structured Json v0.2 cloud event.")
			return v02.Parse(r)
		}

		if version.CloudEventsVersion == "0.1" {
			log.Println("Found HTTP Structured Json v0.1 cloud event.")
			return v01.Parse(r)
		}
	} else if strings.HasPrefix(contentType, "application/json") {
		// this might be Binary Json cloudevent
		version := r.Header.Get("Ce-specversion")
		if version == "0.2" {
			log.Println("Found HTTP Binary Json v0.2 cloud event.")
			return v02.Parse(r)
		}

		version = r.Header.Get("Ce-cloudEventsVersion")
		if version == "0.1" {
			log.Println("Found HTTP Binary Json v0.1 cloud event.")
			return v01.Parse(r)
		}
	}

	return nil
}
