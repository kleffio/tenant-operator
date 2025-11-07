/*
Copyright 2025.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kleffv1 "kleff.io/api/v1"
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=kleff.kleff.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kleff.kleff.io,resources=tenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kleff.kleff.io,resources=tenants/finalizers,verbs=update

// ========================== TENANT PERMISSIONS =============================
// +kubebuilder:rbac:groups=kleff.kleff.io,resources=tenants/events,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups=kleff.kleff.io,resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kleff.kleff.io,resources=services,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Tenant object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/reconcile
func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	tenant := &kleffv1.Tenant{}
	if err := r.Get(ctx, req.NamespacedName, tenant); err != nil {
		// log.Error(err, "Failed to get tenant")
		log.Error(err, "failed to get tenant")
	}

	log.Info("Reconciliating Tenant", "userId", tenant.Spec.UserId, "plan", tenant.Spec.Plan)
	log.Info("Reconciliating Complete!")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kleffv1.Tenant{}).
		Named("tenant").
		Complete(r)
}
