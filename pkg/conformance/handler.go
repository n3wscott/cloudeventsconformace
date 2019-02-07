package conformance

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical"
	canonicalhttp "github.com/n3wscott/cloudeventsconformace/pkg/canonical/http"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to handle request: %s %s", err, spew.Sdump(r))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Invalid request`))
		return
	}
	req := canonicalhttp.Request{
		Header: r.Header,
		Body:   body,
	}

	canonical.Print(req)

	w.WriteHeader(http.StatusNoContent)
}
