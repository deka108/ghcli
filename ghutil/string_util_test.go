package ghutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSnakeCase(t *testing.T) {
	require.Equal(t, "abc", ToSnakeCase("abc"))
	require.Equal(t, "abc_def", ToSnakeCase("abcDef"))
}
