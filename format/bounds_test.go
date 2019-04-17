package format

import (
	"testing"

	"github.com/trivago/gollum/core"
	"github.com/trivago/tgo/ttesting"
)

func TestFormatterTrim(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.Trim")
	config.Override("LeftSeparator", "|")
	config.Override("RightSeparator", "|")

	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*Trim)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("|foo bar foobar|"), nil, core.InvalidStreamID)

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	expect.Equal("foo bar foobar", msg.String())
}

func TestFormatterTrimWithSpaces(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.Trim")
	config.Override("LeftSeparator", " ")
	config.Override("RightSeparator", "  ")

	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*Trim)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte(" foo bar foobar  "), nil, core.InvalidStreamID)

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	expect.Equal("foo bar foobar", msg.String())
}

func TestFormatterTrimTarget(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.Trim")
	config.Override("LeftSeparator", "|")
	config.Override("RightSeparator", "|")
	config.Override("Target", "foo")

	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*Trim)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("|foo bar foobar|"), nil, core.InvalidStreamID)
	msg.GetMetadata().Set("foo", []byte("|foo bar foobar|second foo bar|"))

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	foo, err := msg.GetMetadata().Bytes("foo")
	expect.NoError(err)
	expect.Equal("|foo bar foobar|", msg.String())
	expect.Equal("foo bar foobar|second foo bar", string(foo))
}