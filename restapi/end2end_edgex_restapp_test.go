package restapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPingHandler(t *testing.T) {
	var (
		expected = "pong"
		actual   string
	)

	resp, err := rclient.R().
		SetHeader("Accept", "application/json").
		SetResult(&actual).
		Get("/api/v1/ping")
	require.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, expected, actual)
}
