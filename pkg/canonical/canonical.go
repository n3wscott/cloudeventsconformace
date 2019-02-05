package canonical

import (
	"fmt"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical/v01"
	"strings"
)

func Print(envelope interface{}, data interface{}) {
	var b strings.Builder

	v01.Print(&b, envelope, data)

	fmt.Println(b.String())
}
