package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSPtr(t *testing.T) {
	s := "Hello world"
	expected := &s
	assert.Equal(t, expected, SPtr(s))
}

func TestF64Ptr(t *testing.T) {
	f := 64.32
	expected := &f
	assert.Equal(t, expected, F64Ptr(f))
}
