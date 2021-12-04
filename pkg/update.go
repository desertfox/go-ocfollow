package ocfollow

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	v1 "k8s.io/api/core/v1"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case *v1.Pod:
		m.setPod(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, KeyMap.Select):
			m.setPodName(m.list.GetCursor())
			return m, m.getPodCmd()
		case key.Matches(msg, KeyMap.Back):
			m.clearPod()
		case key.Matches(msg, KeyMap.Quit):
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
