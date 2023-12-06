package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCubeConundrum(t *testing.T) {
	assert.Equal(t, 2237, cubeConundrum("input.txt", 12, 13, 14))
}
