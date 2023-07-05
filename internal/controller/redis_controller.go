/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	rediscolav1alpha1 "redisoperator/redis/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

// RedisReconciler reconciles a Redis object
type RedisReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=redis.cola.redisoperator,resources=redis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=redis.cola.redisoperator,resources=redis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=redis.cola.redisoperator,resources=redis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Redis object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile

func (r *RedisReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := logf.Log.WithName("controller_redis")
	reqLogger = reqLogger.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)
	reqLogger.Info("Creating Redis")

	innstance := &rediscolav1alpha1.Redis{}
	if err := r.Client.Get(context.TODO(), req.NamespacedName, redisInstance); err != nil {
		reqLogger.Error(err, "failed to get Redis")
		return ctrl.Result{}, nil
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s", redisInstance.Name),
			Namespace: redisInstance.Namespace,
			Labels:    redisInstance.Labels,
		},
		Spec: redisInstance.Spec.Template.Spec,
	}
	//TODO 提供SVC访问地址
	//TODO 通过SC实现持久化

	if err := r.Create(ctx, pod); err != nil {
		reqLogger.Error(err, "failed to create Pod")
		return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
	}
	reqLogger.Info(fmt.Sprintf("the Pod (%s) has created", pod.Name))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RedisReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&rediscolav1alpha1.Redis{}).
		Complete(r)
}
