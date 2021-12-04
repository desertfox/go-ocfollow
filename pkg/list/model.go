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
	m.cursor++
}

func (m *Model) cursorDown() {
	m.cursor--
}

func (m Model) GetCursor() string {
	return m.pods[m.cursor]
}
