package ocfollow

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/desertfox/go-ocfollow/pkg/list"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type Model struct {
	namespace string
	podName   string
	pod       *v1.Pod

	list        list.Model
	describeStr string

	clientset kubernetes.Interface
}

func NewModel(n, p string, l list.Model, c kubernetes.Interface) Model {
	return Model{n, p, &v1.Pod{}, l, "", c}
}

func (m *Model) setPod(p *v1.Pod) {
	m.pod = p
}

func (m *Model) setPodName(p string) {
	m.podName = p
}

func (m *Model) clearPod() {
	m.setPodName("")
}

func (m Model) Init() tea.Cmd {
	return nil
}
