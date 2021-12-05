package list

import (
	tea "github.com/charmbracelet/bubbletea"
	v1 "k8s.io/api/core/v1"
)

type Model struct {
	podNameList []string
	cursor      int
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
	if m.cursor > len(m.podNameList)-1 {
		m.resetCursor()
	}
}

func (m *Model) resetCursor() {
	m.cursor = 0
}

func (m Model) GetCursor() string {
	return m.podNameList[m.cursor]
}

func (m *Model) setPodList(pl *v1.PodList) {
	var podNameList []string
	for _, pod := range pl.Items {
		podNameList = append(podNameList, pod.Name)
	}

	m.podNameList = podNameList
}
