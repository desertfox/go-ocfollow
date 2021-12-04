package ocfollow

import (
	"fmt"

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
	return fmt.Sprintf("%v", m.pod)
}

func (m Model) showPodLogs() string {
	return "logs"
}
