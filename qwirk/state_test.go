package qwirk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStates(t *testing.T) {
	stateTests := []struct {
		name     string
		value    State
		expected string
	}{
		{
			"Ready",
			Ready,
			"Ready",
		},
		{
			"InProgress",
			InProgress,
			"InProgress",
		},
		{
			"Done",
			Done,
			"Done",
		},
	}
	for _, tt := range stateTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.String())
		})
	}
}
