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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	kleffv1 "kleff.io/api/v1"
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
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
// kleff.io/internal/controller/tenant_controller.go

func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	tenant := &kleffv1.Tenant{}
	if err := r.Get(ctx, req.NamespacedName, tenant); err != nil {
		// If the Tenant resource is not found, it might have been deleted.
		// We don't need to reconcile anymore, so we can ignore this error.
		// For any other error, we return it to retry the reconciliation.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciliating Tenant", "userId", tenant.Spec.UserId, "plan", tenant.Spec.Plan)

	namespaceReady := false

	if err := r.addNamespaceIfNotExists(ctx, tenant); err != nil {
		log.Error(err, "Failed to add Namespace for tenant")
		addCondition(&tenant.Status, "NamespaceNotReady", metav1.ConditionFalse, "NamespaceNotReady", "Failed to add namespace for tenant")
	} else {
		namespaceReady = true
	}

	if namespaceReady {
		addCondition(&tenant.Status, "TenantReady", metav1.ConditionTrue, "AllSubresourcesReady", "All subresources are ready")
	}

	log.Info("Reconciliating Complete!")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("tenant-controller")
	return ctrl.NewControllerManagedBy(mgr).
		For(&kleffv1.Tenant{}).
		Named("tenant").
		Complete(r)
}

func (r *TenantReconciler) addNamespaceIfNotExists(ctx context.Context, tenant *kleffv1.Tenant) error {
	log := log.FromContext(ctx)

	namespace := &corev1.Namespace{}
	namespaceName := tenant.Spec.UserId
	print(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	fmt.Println(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)
	log.Info(tenant.Spec.UserId)

	err := r.Get(ctx, client.ObjectKey{Name: namespaceName}, namespace)

	// check if namespace exists
	if err == nil {
		// namespace exists, exit
		return nil
	}

	// tenant namespace doesnt exist, create it
	desiredNamespace := generateDesiredNamespace(namespaceName)
	// here

	if err := r.Create(ctx, desiredNamespace); err != nil {
		return err
	}

	r.recorder.Event(tenant, corev1.EventTypeNormal, "NamespaceReady", "Namespace Created Successfully")
	log.Info("Namespace created", "namespace", namespaceName)
	return nil
}

func generateDesiredNamespace(namespaceName string) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}
}

func addCondition(status *kleffv1.TenantStatus, condType string, statusType metav1.ConditionStatus, reason, message string) {
	for i, existingCondition := range status.Conditions {
		if existingCondition.Type == condType {
			status.Conditions[i].Status = statusType
			status.Conditions[i].Reason = reason
			status.Conditions[i].Message = message
			status.Conditions[i].LastTransitionTime = metav1.Now()
			return
		}
	}

	condition := metav1.Condition{
		Type:               condType,
		Status:             statusType,
		Reason:             reason,
		Message:            message,
		LastTransitionTime: metav1.Now(),
	}
	status.Conditions = append(status.Conditions, condition)
}
