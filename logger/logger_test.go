package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerNew(t *testing.T) {
	assert.NotNil(t, New())
}
