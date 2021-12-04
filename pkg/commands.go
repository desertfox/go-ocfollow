package ocfollow

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m Model) getPodCmd() tea.Cmd {
	pod, _ := m.clientset.CoreV1().Pods(m.namespace).Get(context.TODO(), m.podName, metav1.GetOptions{})
	return func() tea.Msg {
		return pod
	}
}
