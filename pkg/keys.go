package ocfollow

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Select key.Binding
	Back   key.Binding
	Quit   key.Binding
}

var KeyMap = keyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("<enter>", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc"),
		key.WithHelp("q/esc", "quit"),
	),
}
