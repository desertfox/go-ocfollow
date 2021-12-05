package ocfollow

import (
	"k8s.io/client-go/kubernetes"
)

func newTestModel(n, p string, clientset kubernetes.Interface) Model {
	return NewModel(n, clientset)
}
