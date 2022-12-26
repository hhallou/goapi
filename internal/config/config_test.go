//go:build unit
// +build unit

package config

import (
	"testing"

	"os"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Nil(t, nil, nil)
}

var portConfigTests = []struct {
	key   string
	value string
	out   string
	desc  string
}{
	{"APP_PORT", "", "8901", "Default port configuration expect"},
	{"APP_PORT", "421", "421", "Customized configuration expect"},
}

func TestGet_Port(t *testing.T) {
	os.Setenv("ENVIRONMENT", "")
	defer os.Unsetenv("ENVIRONMENT")

	for _, test := range portConfigTests {
		// Arrange
		os.Setenv(test.key, test.value)
		defer os.Unsetenv(test.key)

		//Act
		InitConfig()

		//Assert
		assert.Equal(t, test.out, ConfigValues.AppPort, fmt.Sprintf("error with %s on %s", os.Getenv("ENVIRONMENT"), test.desc))
	}
}
