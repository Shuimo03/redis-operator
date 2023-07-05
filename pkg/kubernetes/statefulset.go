package kubernetes

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func CreateStatefulSet(namespace string, statefulSet *appsv1.StatefulSet) error {
	_, err := newKubeClient().AppsV1().StatefulSets(namespace).Create(context.TODO(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create Redis statefulSet  Error: %v", err)
	}
	log.Println("Create Redis statefulSet successfully")
	return nil
}
