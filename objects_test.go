package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipleResourceValue(t *testing.T) {
	var val Value

	resources := []*ResourceValue{
		NewResourceValue(0, Integer(100)).(*ResourceValue),
		NewResourceValue(1, Integer(200)).(*ResourceValue),
		NewResourceValue(2, Integer(300)).(*ResourceValue),
		NewResourceValue(3, Integer(400)).(*ResourceValue),
	}
	val = NewMultipleResourceValue(0, resources).(*MultipleResourceValue)

	assert.NotNil(t, val)
	assert.Equal(t, VALUETYPE_MULTIRESOURCE, val.GetType())
	assert.Equal(t, VALUETYPE_RESOURCE, val.GetContainedType())
	assert.Equal(t, "100,200,300,400,", val.GetStringValue())
}
