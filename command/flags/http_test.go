package flags

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPFlagsSetToken(t *testing.T) {
	var f HTTPFlags
	require.Empty(t, f.Token())
	require.NoError(t, f.SetToken("foo"))
	require.Equal(t, "foo", f.Token())
}
