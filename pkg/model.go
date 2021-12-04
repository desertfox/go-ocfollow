package ocfollow

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/desertfox/go-ocfollow/pkg/list"
)

type Model struct {
	namespace string
	pod       string

	list list.Model
}

func NewModel(n, p string, l list.Model) Model {
	return Model{n, p, l}
}

func (m *Model) setPod(p string) {
	m.pod = p
}

func (m *Model) clearPod() {
	m.setPod("")
}

func (m Model) Init() tea.Cmd {
	//get namespace from kubeconfig
	return nil
}
