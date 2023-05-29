package informer

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
)

type DynamicInformer struct {
	informer cache.SharedIndexInformer
	client   dynamic.DynamicClient
}

func (d *DynamicInformer) Get(namespace, name string) error {
	return nil
}

func NewInformer() *DynamicInformer {
	return nil
}
