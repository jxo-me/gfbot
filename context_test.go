package telebot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ IContext = (*nativeContext)(nil)

func TestContext(t *testing.T) {
	t.Run("Get,Set", func(t *testing.T) {
		var c IContext
		c = new(nativeContext)
		c.Set("name", "Jon Snow")
		assert.Equal(t, "Jon Snow", c.Get("name"))
	})
}
