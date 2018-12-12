package gorm_helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConnectionString(t *testing.T) {
	str, err := BuildConnectionString("postgres", nil)
	assert.Nil(t, err)
	assert.NotNil(t, str)
	assert.NotEmpty(t, str)
}

func TestBuildConnectionString2(t *testing.T) {
	str, err := BuildConnectionString("foobar", nil)
	assert.Error(t, err)
	assert.Empty(t, str)
}
