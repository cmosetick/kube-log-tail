package kubelogtail

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	//"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	//"fmt"
)

/* from client-go examples files:
kubernetes/kubernetes/staging/src/k8s.io/client-go/examples/in-cluster-client-configuration/main.go
 */

 func kubeClient() (*kubernetes.Clientset, error) {
 	// creates the in-cluster config
 	config, err := rest.InClusterConfig()
 	if err != nil {
 		panic(err.Error())
 	}
 	// creates the clientset
 	clientset, err := kubernetes.NewForConfig(config)
 	if err != nil {
 		panic(err.Error())
 	}
	return clientset, nil
}

//we should really do a watch, but this is fine for now
func getPods(clientset *kubernetes.Clientset, namespace string, selector string) (*v1.PodList, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to list pods")
	}
	return pods, nil
}
