package message

import (
	"encoding/json"
	"io"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

var controlMessageSchema = `
{
	"$schema": "https://json-schema.org/draft/2019-09/schema#",
    "$id": "https://tenyks.io/messages.control.schema.json",
    "title": "Control message",
	"description": "Tenyks control message schema",
	"type": "object",
	"required": [
		"oid",
		"kind",
		"timestamp",
	],
	"properties": {
		"oid": {
			"type": "string",
			"format": "uuid"
		},
		"kind": {
			"type": "string"
		},
		"content": {
			"type": "string"
		},
		"timestamp": {
			"type": "string",
			"format": "date-time"
		}
	},
	"additionalProperties": false
}`

// ControlMessage is used to syncronize and communicate with services
// on an operational level. It's usually used to ping a service (and vice versa)
// to see if it still responds and can recieve messages
type ControlMessage struct {
	OID       string    `json:"oid"`
	Kind      string    `json:"kind"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

func (cm *ControlMessage) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(cm)
}

func (cm *ControlMessage) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(cm)
}

func (cm *ControlMessage) Validator() Validator {
	return &JSONSchemaValidator{
		SchemaLoader: gojsonschema.NewStringLoader(controlMessageSchema),
	}
}
