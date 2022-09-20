package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewStringPointer(t *testing.T) {
	s := NewStringPointer("foo")
	sExpected := "foo"

	assert.Equal(t, &sExpected, s)
}

func Test_NewStringSlice(t *testing.T) {
	sSlice := NewStringSlice()

	assert.Equal(t, make([]string, 0), sSlice)
}
