package ocfollow

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPodCmd(t *testing.T) {
	clientset := fake.NewSimpleClientset(
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "test",
				Namespace:   "test",
				Annotations: map[string]string{},
			},
		},
	)

	m := newTestModel("test", "test", clientset)

	cmd := m.getPodCmd()

	got := cmd()
	_, want := got.(*v1.Pod)

	if !want {
		t.Errorf("Incorrect type returned type: %T, value: %v", got, got)
	}

}
