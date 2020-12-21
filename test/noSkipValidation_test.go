package test

import (
	"encoding/json"
	"fmt"
	"testing"

	noSkipValidation "github.com/daverlo/generate/test/noSkipValidation_gen"

	"github.com/stretchr/testify/assert"
)

func TestNoSkipValidation(t *testing.T) {
	// this just tests the name generation works correctly
	s := noSkipValidation.Simple{
		Name: "name",
	}

	v, err := json.Marshal(&s)
	assert.Error(t, err, "json: error calling MarshalJSON for type *noSkipValidation.Simple: address is a required field")
	fmt.Println(v)
}
