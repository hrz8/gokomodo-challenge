package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToInt(t *testing.T) {
	t.Run("success parse", func(t *testing.T) {
		result := ParseStringToInt("12")
		expected := 12

		assert.Equal(t, expected, result)
	})

	t.Run("success parse nil", func(t *testing.T) {
		result := ParseStringToInt("sometext")
		expected := 0

		assert.Equal(t, expected, result)
	})
}
