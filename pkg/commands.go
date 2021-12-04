package ocfollow

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m Model) getPodCmd() tea.Cmd {
	pod, err := m.clientset.CoreV1().Pods(m.namespace).Get(context.TODO(), m.podName, metav1.GetOptions{})
	if err != nil {
		return func() tea.Msg {
			return err
		}
	}
	return func() tea.Msg {
		return pod
	}
}
