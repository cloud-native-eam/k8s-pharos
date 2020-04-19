package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPods(cs *kubernetes.Clientset) int {

	pods, err := cs.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
	panic(err.Error())
	}
	numItems := len(pods.Items)
	return numItems
}
