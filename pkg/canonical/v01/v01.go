package v01

import (
	"fmt"
	"strings"
)

func Print(b *strings.Builder, envelope interface{}, data interface{}) {
	PrintMessage(b, Message{})
}

type String string
type URI string
type Timestamp string
type Map map[String]Object
type Object String // Map, String, or Binary
type Binary []byte

// String - Sequence of printable Unicode characters.
// Binary - Sequence of bytes.
// Map - String-indexed dictionary of Object-typed values
// Object - Either a String, or a Binary, or a Map
// URI - String expression conforming to URI-reference as defined in RFC 3986 §4.1.
// Timestamp - String expression as defined in RFC 3339

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

func (s Timestamp) String() string {
	if len(s) == 0 {
		return "nil"
	}
	return string(s)
}

func (s Object) String() string {
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
	eventType          String    //
	eventTypeVersion   String    // +optional
	cloudEventsVersion String    //
	source             URI       //
	eventID            String    //
	eventTime          Timestamp // +optional
	schemaURL          URI       // +optional
	contentType        String    // +optional
	extensions         Map       // +optional
	data               Object    // +optional
}

func PrintMessage(b *strings.Builder, m Message) {
	b.WriteString(fmt.Sprintf("eventType: %v\n", m.eventType))
	b.WriteString(fmt.Sprintf("eventTypeVersion: %v\n", m.eventTypeVersion))
	b.WriteString(fmt.Sprintf("cloudEventsVersion: %v\n", m.cloudEventsVersion))
	b.WriteString(fmt.Sprintf("source: %v\n", m.source))
	b.WriteString(fmt.Sprintf("eventID: %v\n", m.eventID))
	b.WriteString(fmt.Sprintf("eventTime: %v\n", m.eventTime))
	b.WriteString(fmt.Sprintf("schemaURL: %v\n", m.schemaURL))
	b.WriteString(fmt.Sprintf("contentType: %v\n", m.contentType))
	b.WriteString(fmt.Sprintf("extensions: %v\n", m.extensions))
	b.WriteString(fmt.Sprintf("data: %v\n", m.data))
}

/*
eventType
Type: String
Description: Type of occurrence which has happened. Often this property is used for routing, observability, policy enforcement, etc.
Constraints:
REQUIRED
MUST be a non-empty string
SHOULD be prefixed with a reverse-DNS name. The prefixed domain dictates the organization which defines the semantics of this event type.
Examples
com.github.pull.create

eventTypeVersion
Type: String
Description: The version of the eventType. This enables the interpretation of data by eventual consumers, requires the consumer to be knowledgeable about the producer.
Constraints:
OPTIONAL
If present, MUST be a non-empty string

cloudEventsVersion
Type: String
Description: The version of the CloudEvents specification which the event uses. This enables the interpretation of the context.
Constraints:
REQUIRED
MUST be a non-empty string

source
Type: URI
Description: This describes the event producer. Often this will include information such as the type of the event source, the organization publishing the event, and some unique idenfitiers. The exact syntax and semantics behind the data encoded in the URI is event producer defined.
Constraints:
REQUIRED

eventID
Type: String
Description: ID of the event. The semantics of this string are explicitly undefined to ease the implementation of producers. Enables deduplication.
Examples:
A database commit ID
Constraints:
REQUIRED
MUST be a non-empty string
MUST be unique within the scope of the producer

eventTime
Type: Timestamp
Description: Timestamp of when the event happened.
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 3339

schemaURL
Type: URI
Description: A link to the schema that the data attribute adheres to.
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 3986

contentType
Type: String per RFC 2046
Description: Describe the data encoding format
Constraints:
OPTIONAL
If present, MUST adhere to the format specified in RFC 2046
For Media Type examples see IANA Media Types

extensions
Type: Map
Description: This is for additional metadata and this does not have a mandated structure. This enables a place for custom fields a producer or middleware might want to include and provides a place to test metadata before adding them to the CloudEvents specification. See the Extensions document for a list of possible properties.
Constraints:
OPTIONAL
If present, MUST contain at least one entry
Examples:
authorization data

data
Type: Object
Description: The event payload. The payload depends on the eventType, schemaURL and eventTypeVersion, the payload is encoded into a media format which is specified by the contentType attribute (e.g. application/json).
Constraints:
OPTIONAL
*/
