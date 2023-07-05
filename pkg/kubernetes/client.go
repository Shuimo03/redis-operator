package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sync"
)

var (
	once      sync.Once
	clientSet *kubernetes.Clientset
)

func newKubeClient() *kubernetes.Clientset {
	once.Do(func() {
		config, err := parserConfig()
		if err != nil {
			panic("failed to parser Config File")
		}
		clientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic("failed to Init client")
		}
	})
	return clientSet
}

func parserConfig() (*rest.Config, error) {
	return &rest.Config{}, nil
}
