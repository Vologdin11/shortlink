package repository

import (
	"testing"

	conf "shortlink/pkg/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLink(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		config := conf.NewConfig("localhost", "5432", "postgres", "qweasd", "shortlink", 5, 4)
		postgres := Postgres{}
		err := postgres.InitPostgres(config)
		require.NoError(t, err)

		shortlilnk, err := postgres.GetLink("http://localhost:8000/gl")
		require.NoError(t, err)
		assert.Equal(t, "https://google.com", shortlilnk)
	})
}
