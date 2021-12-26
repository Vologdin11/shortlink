package inmemory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLink(t *testing.T) {
	t.Run("link in db", func(t *testing.T) {
		db := NewInmemory()
		err := db.AddLink("ozon", "asd")
		require.NoError(t, err)

		link, err := db.GetLink("asd")

		assert.NoError(t, err)
		assert.Equal(t, "ozon", link)
	})

	t.Run("no link", func(t *testing.T) {
		db := NewInmemory()

		link, err := db.GetLink("asd")

		assert.Error(t, err)
		assert.Equal(t, "", link)
	})
}

func TestGetShortLink(t *testing.T) {
	t.Run("shortlink in db", func(t *testing.T) {
		db := NewInmemory()
		err := db.AddLink("ozon", "asd")
		require.NoError(t, err)

		link, err := db.GetShortLink("ozon")

		assert.NoError(t, err)
		assert.Equal(t, "asd", link)
	})

	t.Run("no shortlink", func(t *testing.T) {
		db := NewInmemory()

		link, err := db.GetShortLink("ozon")

		assert.Error(t, err)
		assert.Equal(t, "", link)
	})
}
