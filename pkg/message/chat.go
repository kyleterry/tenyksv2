package message

import (
	"encoding/json"
	"io"
	"time"

	"github.com/xeipuuv/gojsonschema"
)

var chatMessageSchema = `
{
	"$schema": "https://json-schema.org/draft/2019-09/schema#",
    "$id": "https://tenyks.io/messages.chat.schema.json",
    "title": "Chat message",
	"description": "Tenyks chat message schema",
	"type": "object",
	"required": [
        "destinationPath",
		"content",
		"timestamp"
	],
	"properties": {
		"destinationPath": {
			"type": "string",
            "description": "the path to the target that got the message or is intended to recieve the message"
		},
		"originPath": {
			"type": "string",
            "description": "the path to the target that sent the message"
		},
		"direct": {
			"type": "boolean",
            "description": "whether the message was direct or intended to be direct"
		},
        "mention": {
            "type": "boolean",
            "description": "whether the message contains the connection nick in a channel message"
        },
		"content": {
			"type": "string",
            "description": "the content of the message"
		},
		"timestamp": {
			"type": "string",
			"format": "date-time",
            "description": "when the message was created"
		}
	},
	"additionalProperties": false
}`

// ChatMessage is a message derived from a human (or robot)
// usually from a channel or a direct message mechanism
type ChatMessage struct {
	DestinationPath string    `json:"destinationPath"`
	OriginPath      string    `json:"originPath"`
	Direct          bool      `json:"direct"`
	Mention         bool      `json:"mention"`
	Content         string    `json:"content"`
	Timestamp       time.Time `json:"timestamp"`
}

func (cm *ChatMessage) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(cm)
}

func (cm *ChatMessage) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(cm)
}

func (cm *ChatMessage) Validator() Validator {
	return &JSONSchemaValidator{
		SchemaLoader: gojsonschema.NewStringLoader(chatMessageSchema),
	}
}
