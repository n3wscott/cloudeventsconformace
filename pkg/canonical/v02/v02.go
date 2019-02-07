package v02

import (
	"fmt"
	"github.com/n3wscott/cloudeventsconformace/pkg/canonical/http"
	"strings"
)

type Integer int32
type String string
type URI string
type URIRef string
type Timestamp string
type Map map[String]Any
type Any []byte // Map, String, Binary, or Integer
type Binary []byte

// Integer - A 32-bit whole number.
// String - Sequence of printable Unicode characters.
// Binary - Sequence of bytes.
// Map - String-indexed dictionary of Any-typed values.
// Any - Either a String, or a Binary, or a Map, or an Integer.
// URI-reference - String expression conforming to URI-reference as defined in RFC 3986 §4.1.
// Timestamp - String expression as defined in RFC 3339.

func (s String) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s URI) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s URIRef) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s Timestamp) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s Any) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s Map) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return fmt.Sprintf("%v", s)
}

type Message struct {
	event_type  String    // (called type)
	specversion String    //
	source      URIRef    //
	id          String    //
	time        Timestamp // +optional
	schemaurl   URI       // +optional
	contenttype String    // +optional
	data        Any       // +optional
	// Everything else is an extension.
}

func Parse(r http.Request) Message {
	// TODO
	return Message{}
}

func (m Message) Print(b *strings.Builder) {
	b.WriteString(fmt.Sprintf("type: %v\n", m.event_type))
	b.WriteString(fmt.Sprintf("specversion: %v\n", m.specversion))
	b.WriteString(fmt.Sprintf("source: %v\n", m.source))
	b.WriteString(fmt.Sprintf("id: %v\n", m.id))
	b.WriteString(fmt.Sprintf("time: %v\n", m.time))
	b.WriteString(fmt.Sprintf("schemaurl: %v\n", m.schemaurl))
	b.WriteString(fmt.Sprintf("contenttype: %v\n", m.contenttype))
	b.WriteString(fmt.Sprintf("data: %v\n", m.data))
}

/*

type
Type: String
Description: Type of occurrence which has happened. Often this attribute is used for routing, observability, policy enforcement, etc. The format of this is producer defined and might include information such as the version of the eventtype - see Versioning of Attributes in the Primer for more information.
Constraints:
REQUIRED
MUST be a non-empty string
SHOULD be prefixed with a reverse-DNS name. The prefixed domain dictates the organization which defines the semantics of this event type.
Examples
com.github.pull.create
com.example.object.delete.v2

specversion
Type: String
Description: The version of the CloudEvents specification which the event uses. This enables the interpretation of the context. Compliant event producers MUST use a value of 0.2 when referring to this version of the specification.
Constraints:
REQUIRED
MUST be a non-empty string

source
Type: URI-reference
Description: This describes the event producer. Often this will include information such as the type of the event source, the organization publishing the event, the process that produced the event, and some unique identifiers. The exact syntax and semantics behind the data encoded in the URI is event producer defined.
Constraints:
REQUIRED
Examples
https://github.com/cloudevents/spec/pull/123
/cloudevents/spec/pull/123
urn:event:from:myapi/resourse/123
mailto:cncf-wg-serverless@lists.cncf.io

id
Type: String
Description: ID of the event. The semantics of this string are explicitly undefined to ease the implementation of producers. Enables deduplication.
Examples:
A database commit ID
Constraints:
REQUIRED
MUST be a non-empty string
MUST be unique within the scope of the producer

time
Type: Timestamp
Description: Timestamp of when the event happened.
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 3339

schemaurl
Type: URI
Description: A link to the schema that the data attribute adheres to. Incompatible changes to the schema SHOULD be reflected by a different URL. See Versioning of Attributes in the Primer for more information.
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 3986

contenttype
Type: String per RFC 2046
Description: Content type of the data attribute value. This attribute enables the data attribute to carry any type of content, whereby format and encoding might differ from that of the chosen event format. For example, an event rendered using the JSON envelope format might carry an XML payload in its data attribute, and the consumer is informed by this attribute being set to "application/xml". The rules for how the data attribute content is rendered for different contenttype values are defined in the event format specifications; for example, the JSON event format defines the relationship in section 3.1.
When this attribute is omitted, the "data" attribute simply follows the event format's encoding rules. For the JSON event format, the "data" attribute value can therefore be a JSON object, array, or value.
For the binary mode of some of the CloudEvents transport bindings, where the "data" content is immediately mapped into the payload of the transport frame, this field is directly mapped to the respective transport or application protocol's content-type metadata property. Normative rules for the binary mode and the content-type metadata mapping can be found in the respective transport mapping specifications.
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 2046
For Media Type examples see IANA Media Types

Data Attribute
As defined by the term Data, CloudEvents MAY include domain-specific information about the occurrence. When present, this information will be encapsulated within the data attribute.

data
Type: Any
Description: The event payload. The payload depends on the type and the schemaurl. It is encoded into a media format which is specified by the contenttype attribute (e.g. application/json).
Constraints:
OPTIONAL

*/
