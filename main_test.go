package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/osohq/go-oso"
)

func TestPolicy(t *testing.T) {
	oso, err := oso.NewOso()
	require.NoError(t, err)

	require.NoError(t, oso.LoadFiles([]string{"policy.polar"}))

	assert.NoError(t, oso.Authorize("actor", "action", "resource"))
}
