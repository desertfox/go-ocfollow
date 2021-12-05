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
	podList   *v1.PodList

	list        list.Model
	describeStr string
	podLogs     string

	clientset kubernetes.Interface
}

func NewModel(c kubernetes.Interface) Model {
	return Model{"", "", &v1.Pod{}, &v1.PodList{}, list.Model{}, "", "", c}
}

func (m *Model) setPod(p *v1.Pod) {
	m.pod = p
}

func (m *Model) setPodList(pl *v1.PodList) {
	m.podList = pl
}

func (m *Model) setPodName(p string) {
	m.podName = p
}

func (m *Model) clearPod() {
	m.setPodName("")
}

func (m Model) Init() tea.Cmd {
	return m.getPodListCmd()
}
