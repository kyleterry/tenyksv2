package message

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// JSONSchemaValidator will validate json messages against
// the configured Schema
type JSONSchemaValidator struct {
	SchemaLoader gojsonschema.JSONLoader
}

func (j JSONSchemaValidator) Validate(b []byte) error {
	schema, err := gojsonschema.NewSchema(j.SchemaLoader)
	if err != nil {
		return err
	}

	result, err := schema.Validate(gojsonschema.NewBytesLoader(b))
	if err != nil {
		return err
	}

	if !result.Valid() {
		return fmt.Errorf("schema validation failed: %s", result.Errors())
	}

	return nil
}
