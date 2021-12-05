package list

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	v1 "k8s.io/api/core/v1"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *v1.PodList:
		m.setPodList(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, KeyMap.Up):
			m.cursorUp()
		case key.Matches(msg, KeyMap.Down):
			m.cursorDown()
		}
	}

	return m, nil
}
