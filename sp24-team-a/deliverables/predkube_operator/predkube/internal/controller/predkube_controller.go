package controller

import (
	"context"
	"strconv"
	// "fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "k8s.io/api/apps/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	predkubev1 "test.com/predkube/api/v1"
)

// ScalerReconciler reconciles a Scaling object
type ScalerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.test.com,resources=scalings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.test.com,resources=scalings/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.test.com,resources=scalings/finalizers,verbs=update
//+kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;update;patch

// Reconcile is part of the main Kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var scaling predkubev1.Scaling
    if err := r.Get(ctx, req.NamespacedName, &scaling); err != nil {
        logger.Error(err, "Unable to fetch Scaling")
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // Find the targeted Deployment in the specified namespace
    var deployment appsv1.Deployment
    if err := r.Get(ctx, client.ObjectKey{Name: scaling.Spec.TargetDeployment, Namespace: scaling.Spec.TargetNamespace}, &deployment); err != nil {
        if errors.IsNotFound(err) {
            logger.Error(err, "Deployment not found", "Namespace", scaling.Spec.TargetNamespace)
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

	// // Find the targeted Deployment in the specified namespace
    // var deployment appsv1.Deployment
    // if err := r.Get(ctx, client.ObjectKey{Name: scaling.Spec.TargetPodName, Namespace: "default"}, &deployment); err != nil {
    //     if errors.IsNotFound(err) {
    //         logger.Error(err, "Deployment not found", "Namespace", "default")
    //         return ctrl.Result{}, nil
    //     }
    //     return ctrl.Result{}, err
    // }

	// Read min and max replicas from the deployment annotations
    minReplicaCount, err := strconv.Atoi(deployment.Annotations["scaler.kubernetes.io/min-replicas"])
    if err != nil {
        logger.Error(err, "Failed to read min replicas from annotations")
        minReplicaCount = 1 // Default value if not set or error
    }
    maxReplicaCount, err := strconv.Atoi(deployment.Annotations["scaler.kubernetes.io/max-replicas"])
    if err != nil {
        logger.Error(err, "Failed to read max replicas from annotations")
        maxReplicaCount = 10 // Default value if not set or error
    }

	// Adjust the replica count within the allowed minimum and maximum
	desiredReplicas := scaling.Spec.TargetReplicaCount
	if desiredReplicas < minReplicaCount {
		desiredReplicas = minReplicaCount
	} else if desiredReplicas > maxReplicaCount {
		desiredReplicas = maxReplicaCount
	}

	// Update the Deployment if necessary
	desiredReplicas_int32 := int32(desiredReplicas)
	if *deployment.Spec.Replicas != desiredReplicas_int32 {
		deployment.Spec.Replicas = &desiredReplicas_int32
		if err := r.Update(ctx, &deployment); err != nil {
			logger.Error(err, "Failed to update Deployment replica count")
			return ctrl.Result{}, err
		}
		logger.Info("Updated Deployment replica count", "Deployment", deployment.Name, "NewReplicas", desiredReplicas_int32)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&predkubev1.Scaling{}).
		Complete(r)
}
