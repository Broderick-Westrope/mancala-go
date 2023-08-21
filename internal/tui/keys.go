package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Left   key.Binding
	Right  key.Binding
	Submit key.Binding
	Help   key.Binding
	Quit   key.Binding
}

// ShortHelp returns the key bindings that should be displayed by default.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns the key bindings that should be displayed after the user has toggled help.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Right},
		{k.Submit, k.Quit},
		{k.Help},
	}
}

var keys = keyMap{
	Left: key.NewBinding(
		key.WithKeys("left", "h", "a"),
		key.WithHelp("←/h/a", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l", "d"),
		key.WithHelp("→/l/d", "move right"),
	),
	Submit: key.NewBinding(
		key.WithKeys(" ", "enter"),
		key.WithHelp("enter/space", "submit selection"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q/escape/ctrl+c", "quit"),
	),
}
