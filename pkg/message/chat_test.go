package message

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	validChatMessage = `
{
    "destinationPath": "/irc/freenode/#tenyks",
    "originPath": "/irc/freenode/mr-user-person",
    "direct": false,
    "content": "this is a test message",
    "timestamp": "2020-08-21T03:23:30-07:00"
}`
	validNoOriginChatMessage = `
{
    "destinationPath": "/irc/freenode/#tenyks",
    "direct": false,
    "content": "this is a test message",
    "timestamp": "2020-08-21T03:23:30-07:00"
}`
	invalidChatMessage = `
{
    "invalidKey": "foobar",
    "direct": false,
    "content": "this is a test message"
}`
)

func TestChatMessageValidation(t *testing.T) {
	msg := ChatMessage{}
	buf := bytes.NewBufferString(validChatMessage)

	require.NoError(t, msg.Validator().Validate(buf.Bytes()))
	require.NoError(t, msg.Decode(buf))

	expectedTime, err := time.Parse(time.RFC3339, "2020-08-21T03:23:30-07:00")
	require.NoError(t, err)

	require.Equal(t, "/irc/freenode/#tenyks", msg.DestinationPath)
	require.Equal(t, "/irc/freenode/mr-user-person", msg.OriginPath)
	require.Equal(t, false, msg.Direct)
	require.Equal(t, "this is a test message", msg.Content)
	require.Equal(t, expectedTime, msg.Timestamp)

	{
		msg := ChatMessage{}
		buf := bytes.NewBufferString(validNoOriginChatMessage)

		require.NoError(t, msg.Validator().Validate(buf.Bytes()))
		require.NoError(t, msg.Decode(buf))

		expectedTime, err := time.Parse(time.RFC3339, "2020-08-21T03:23:30-07:00")
		require.NoError(t, err)

		require.Equal(t, "/irc/freenode/#tenyks", msg.DestinationPath)
		require.Equal(t, false, msg.Direct)
		require.Equal(t, "this is a test message", msg.Content)
		require.Equal(t, expectedTime, msg.Timestamp)
	}

	{
		msg := ChatMessage{}
		buf := bytes.NewBufferString(invalidChatMessage)

		require.Error(t, msg.Validator().Validate(buf.Bytes()))
	}
}
