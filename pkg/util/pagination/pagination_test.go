package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOffset(t *testing.T) {
	t.Run("success determine", func(t *testing.T) {
		result := GetOffset(3, 5)
		expected := 10

		assert.Equal(t, expected, result)
	})

	t.Run("success with minus input", func(t *testing.T) {
		result := GetOffset(-1, 5)
		expected := 0

		assert.Equal(t, expected, result)
	})
}
