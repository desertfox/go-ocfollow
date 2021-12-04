package ocfollow

import (
	"github.com/desertfox/go-ocfollow/pkg/list"
	"k8s.io/client-go/kubernetes"
)

func newTestModel(n, p string, clientset kubernetes.Interface) Model {
	return NewModel(n, p, list.NewModel([]string{}, 0), clientset)
}
