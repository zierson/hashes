package strix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMD5(t *testing.T) {
	assert.Equal(t, "cc03e747a6afbbcbf8be7668acfebee5", ToMD5("test123"))
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", ToMD5(""))
}

func TestRandomMD5(t *testing.T) {
	assert.Regexp(t, "^[a-z0-9]{32}$", RandomMD5())
}
