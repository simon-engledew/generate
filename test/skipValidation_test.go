package test

import (
	"encoding/json"
	"testing"

	skipValidation "github.com/simon-engledew/generate/test/skipValidation_gen"
	"gotest.tools/assert"
)

func TestSkipValidation(t *testing.T) {
	// this just tests the name generation works correctly
	s := skipValidation.Simple{
		Address: &skipValidation.Address{
			Street: "street",
		},
	}

	_, err := json.Marshal(&s)
	assert.Equal(t, err, nil)
}
