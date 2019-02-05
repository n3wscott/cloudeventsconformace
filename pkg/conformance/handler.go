package conformance

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/knative/pkg/cloudevents"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical"
	"io"
	"log"
	"net/http"
)

var _ http.Handler = (*handler)(nil)

func NewHandler() http.Handler {
	return &handler{}
}

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var rawData io.Reader
	eventContext, err := cloudevents.FromRequest(&rawData, r)
	if err != nil {
		log.Printf("Failed to handle request: %s %s", err, spew.Sdump(r))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid request`))
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, struct{}{}, eventContext)

	//	data := rawData

	log.Println(eventContext)
	canonical.Print(eventContext, r)

	w.WriteHeader(http.StatusNoContent)
}
