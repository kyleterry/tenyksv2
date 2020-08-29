package irc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNamesReply(t *testing.T) {
	raw := ":user!~nick@server/mask 353 nick = #tenyks :nick1 nick2 nick3 nick4"

	msg, err := NewRawMessageDecoder().Decode(raw)
	require.NoError(t, err)

	nr := NamesReply{m: msg}

	require.NoError(t, nr.Validate())

	require.Equal(t, []string{"nick1", "nick2", "nick3", "nick4"}, nr.Names())
	require.Equal(t, "#tenyks", nr.Channel())
}
