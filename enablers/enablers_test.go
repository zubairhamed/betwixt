package enablers

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/response"
	"testing"
)

func TestNullEnabler(t *testing.T) {
	en := NewNullEnabler()

	assert.Equal(t, response.MethodNotAllowed(), en.OnRead(0, 0, nil))
	assert.Equal(t, response.MethodNotAllowed(), en.OnDelete(0, nil))
	assert.Equal(t, response.MethodNotAllowed(), en.OnWrite(0, 0, nil))
	assert.Equal(t, response.MethodNotAllowed(), en.OnExecute(0, 0, nil))
	assert.Equal(t, response.MethodNotAllowed(), en.OnCreate(0, 0, nil))
}
