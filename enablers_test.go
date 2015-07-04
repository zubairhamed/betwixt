package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullEnabler(t *testing.T) {
	en := NewNullEnabler()

	assert.Equal(t, MethodNotAllowed(), en.OnRead(0, 0, nil))
	assert.Equal(t, MethodNotAllowed(), en.OnDelete(0, nil))
	assert.Equal(t, MethodNotAllowed(), en.OnWrite(0, 0, nil))
	assert.Equal(t, MethodNotAllowed(), en.OnExecute(0, 0, nil))
	assert.Equal(t, MethodNotAllowed(), en.OnCreate(0, 0, nil))
}
