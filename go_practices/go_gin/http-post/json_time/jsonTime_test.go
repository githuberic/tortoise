package json_time

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONTime_UnmarshalJSON(t *testing.T) {
	var timeUnmarshalJSON JSONTime
	err := timeUnmarshalJSON.UnmarshalJSON([]byte(`"1970-01-26"`))
	assert.Nil(t, err)
}

func TestJSONTime_MarshalJSON(t *testing.T) {
	var timeUnmarshalJSON JSONTime
	marshalJSON, err := timeUnmarshalJSON.MarshalJSON()

	assert.Nil(t, err)
	assert.NotNil(t, marshalJSON)
}

func TestJSONTime_String(t *testing.T) {
	var timeUnmarshalJSON JSONTime
	s := timeUnmarshalJSON.String()

	assert.NotNil(t, s)
}
