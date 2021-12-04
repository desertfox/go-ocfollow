package list

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up   key.Binding
	Down key.Binding
}

var KeyMap = keyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
}
