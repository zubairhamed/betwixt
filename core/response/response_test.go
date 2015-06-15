package response

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/canopus"
	"testing"
	"github.com/zubairhamed/go-commons/typeval"
)

func TestResponse(t *testing.T) {
	assert.Equal(t, COAPCODE_201_CREATED, Created().GetResponseCode())
	assert.Equal(t, typeval.Empty(), Created().GetResponseValue())

	assert.Equal(t, COAPCODE_202_DELETED, Deleted().GetResponseCode())
	assert.Equal(t, typeval.Empty(), Deleted().GetResponseValue())

	assert.Equal(t, COAPCODE_204_CHANGED, Changed().GetResponseCode())
	assert.Equal(t, typeval.Empty(), Changed().GetResponseValue())

	assert.Equal(t, COAPCODE_205_CONTENT, Content(typeval.String("this is a string")).GetResponseCode())
	assert.Equal(t, "this is a string", Content(typeval.String("this is a string")).GetResponseValue().GetValue())

	assert.Equal(t, COAPCODE_400_BAD_REQUEST, BadRequest().GetResponseCode())
	assert.Equal(t, typeval.Empty(), BadRequest().GetResponseValue())

	assert.Equal(t, COAPCODE_401_UNAUTHORIZED, Unauthorized().GetResponseCode())
	assert.Equal(t, typeval.Empty(), Unauthorized().GetResponseValue())

	assert.Equal(t, COAPCODE_404_NOT_FOUND, NotFound().GetResponseCode())
	assert.Equal(t, typeval.Empty(), NotFound().GetResponseValue())

	assert.Equal(t, COAPCODE_405_METHOD_NOT_ALLOWED, MethodNotAllowed().GetResponseCode())
	assert.Equal(t, typeval.Empty(), MethodNotAllowed().GetResponseValue())

	assert.Equal(t, COAPCODE_409_CONFLICT, Conflict().GetResponseCode())
	assert.Equal(t, typeval.Empty(), Conflict().GetResponseValue())
}
