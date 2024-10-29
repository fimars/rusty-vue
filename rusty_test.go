package rusty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRusty(t *testing.T) {
	assert.Equal(t, Rusty(), "<button>Count is: 0</button>")
}
