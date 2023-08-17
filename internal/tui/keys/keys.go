package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Left   key.Binding
	Right  key.Binding
	Submit key.Binding
	Help   key.Binding
	Quit   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Right},
		{k.Submit, k.Quit},
		{k.Help},
	}
}

var Keys = KeyMap{
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
