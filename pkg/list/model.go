package list

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	pods   []string
	cursor int
}

func NewModel(p []string, c int) Model {
	return Model{p, c}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) cursorUp() {
	m.cursor--
	if m.cursor < 0 {
		m.resetCursor()
	}
}

func (m *Model) cursorDown() {
	m.cursor++
	if m.cursor > len(m.pods)-1 {
		m.resetCursor()
	}
}

func (m *Model) resetCursor() {
	m.cursor = 0
}

func (m Model) GetCursor() string {
	return m.pods[m.cursor]
}
