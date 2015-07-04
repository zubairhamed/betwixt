package betwixt

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/go-commons/typeval"
	"testing"
)

func TestMultipleResourceValue(t *testing.T) {
	var val typeval.Value

	resources := []*ResourceValue{
		NewResourceValue(0, typeval.Integer(100)).(*ResourceValue),
		NewResourceValue(1, typeval.Integer(200)).(*ResourceValue),
		NewResourceValue(2, typeval.Integer(300)).(*ResourceValue),
		NewResourceValue(3, typeval.Integer(400)).(*ResourceValue),
	}
	val = NewMultipleResourceValue(0, resources).(*MultipleResourceValue)

	assert.NotNil(t, val)
	assert.Equal(t, typeval.VALUETYPE_MULTIRESOURCE, val.GetType())
	assert.Equal(t, typeval.VALUETYPE_RESOURCE, val.GetContainedType())
	assert.Equal(t, "100,200,300,400,", val.GetStringValue())
}
