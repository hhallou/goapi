//go:build unit
// +build unit

package logger

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	InitLogger()

	log := getLogger()

	assert.NotNil(t, log)
	assert.Equal(t, log.Formatter, &logrus.JSONFormatter{})
	assert.Equal(t, log.Level, logrus.ErrorLevel)
	assert.Equal(t, log.Out, os.Stderr)
}

func TestWithoutInitLogger(t *testing.T) {
	log := getLogger()

	assert.NotNil(t, log)
}
