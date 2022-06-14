/*
Copyright 2022.

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

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	distributedsecretsv1alpha1 "github.com/DeeAjayi/distributed-secrets/api/v1alpha1"
)

// DistributedSecretsReconciler reconciles a DistributedSecrets object
type DistributedSecretsReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=distributed-secrets.distributed-secrets.com,resources=distributedsecrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=distributed-secrets.distributed-secrets.com,resources=distributedsecrets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=distributed-secrets.distributed-secrets.com,resources=distributedsecrets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DistributedSecrets object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *DistributedSecretsReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Initialize controller
	distributedSecrets := &distributedsecretsv1alpha1.DistributedSecrets{}
	err := r.Get(ctx, req.NamespacedName, distributedSecrets)

	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Distributed Secret not found")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get Distributed Secret")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DistributedSecretsReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&distributedsecretsv1alpha1.DistributedSecrets{}).
		Complete(r)
}
