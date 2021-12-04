package ocfollow

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.podName == "" {
		return m.list.View()
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.showPodDescribe(),
		m.showPodLogs(),
	)
}

func (m Model) showPodDescribe() string {
	return m.describeStr
}

func (m Model) showPodLogs() string {
	return m.podLogs
}
