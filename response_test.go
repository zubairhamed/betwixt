package betwixt

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/canopus"
	"testing"
)

func TestResponse(t *testing.T) {
	assert.Equal(t, CoapCodeCreated, Created().GetResponseCode())
	assert.Equal(t, Empty(), Created().GetResponseValue())

	assert.Equal(t, CoapCodeDeleted, Deleted().GetResponseCode())
	assert.Equal(t, Empty(), Deleted().GetResponseValue())

	assert.Equal(t, CoapCodeChanged, Changed().GetResponseCode())
	assert.Equal(t, Empty(), Changed().GetResponseValue())

	assert.Equal(t, CoapCodeContent, Content(String("this is a string")).GetResponseCode())
	assert.Equal(t, "this is a string", Content(String("this is a string")).GetResponseValue().GetValue())

	assert.Equal(t, CoapCodeBadRequest, BadRequest().GetResponseCode())
	assert.Equal(t, Empty(), BadRequest().GetResponseValue())

	assert.Equal(t, CoapCodeUnauthorized, Unauthorized().GetResponseCode())
	assert.Equal(t, Empty(), Unauthorized().GetResponseValue())

	assert.Equal(t, CoapCodeNotFound, NotFound().GetResponseCode())
	assert.Equal(t, Empty(), NotFound().GetResponseValue())

	assert.Equal(t, CoapCodeMethodNotAllowed, MethodNotAllowed().GetResponseCode())
	assert.Equal(t, Empty(), MethodNotAllowed().GetResponseValue())

	assert.Equal(t, CoapCodeConflict, Conflict().GetResponseCode())
	assert.Equal(t, Empty(), Conflict().GetResponseValue())
}
