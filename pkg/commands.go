package ocfollow

import (
	"bytes"
	"context"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubectl/pkg/describe"
)

type podDescribeMsg string

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

func (m Model) getPodDescribeCmd() tea.Cmd {
	podDescribeClient := describe.PodDescriber{Interface: m.clientset}

	description, err := podDescribeClient.Describe(m.namespace, m.podName, describe.DescriberSettings{})

	if err != nil {
		return func() tea.Msg {
			return err
		}
	}
	return func() tea.Msg {
		return podDescribeMsg(description)
	}
}

func (m Model) getPodLogsCmd() tea.Cmd {
	req := m.clientset.CoreV1().Pods(m.namespace).GetLogs(m.podName, &v1.PodLogOptions{})

	podLogs, _ := req.Stream(context.TODO())
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, podLogs)

	return func() tea.Msg {
		return buf.String()
	}
}
