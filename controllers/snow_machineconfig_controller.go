package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	anywherev1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/aws"
	"github.com/aws/eks-anywhere/pkg/providers/snow"
)

type ClientBuilder interface {
	Build(ctx context.Context) (aws.Clients, error)
}

// SnowMachineConfigReconciler reconciles a SnowMachineConfig object
type SnowMachineConfigReconciler struct {
	client        client.Client
	log           logr.Logger
	clientBuilder ClientBuilder
}

func NewSnowMachineConfigReconciler(client client.Client, log logr.Logger, clientBuilder ClientBuilder) *SnowMachineConfigReconciler {
	return &SnowMachineConfigReconciler{
		client:        client,
		log:           log,
		clientBuilder: clientBuilder,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *SnowMachineConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&anywherev1.SnowMachineConfig{}).
		Complete(r)
}

// TODO: add here kubebuilder permissions as needed
func (r *SnowMachineConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	log := r.log.WithValues("snowMachineConfig", req.NamespacedName)

	// Fetch the SnowMachineConfig object
	snowMachineConfig := &anywherev1.SnowMachineConfig{}
	log.Info("Reconciling snowmachineconfig")
	if err := r.client.Get(ctx, req.NamespacedName, snowMachineConfig); err != nil {
		return ctrl.Result{}, err
	}

	// Initialize the patch helper
	patchHelper, err := patch.NewHelper(snowMachineConfig, r.client)
	if err != nil {
		return ctrl.Result{}, err
	}

	defer func() {
		// Always attempt to patch the object and status after each reconciliation.
		patchOpts := []patch.Option{}

		if err := patchHelper.Patch(ctx, snowMachineConfig, patchOpts...); err != nil {
			reterr = kerrors.NewAggregate([]error{reterr, fmt.Errorf("patching snowmachineconfig: %v", err)})
		}
	}()

	// There's no need to go any further if the SnowMachineConfig is marked for deletion.
	if !snowMachineConfig.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, reterr
	}

	result, err := r.reconcile(ctx, snowMachineConfig)
	if err != nil {
		reterr = kerrors.NewAggregate([]error{reterr, fmt.Errorf("reconciling snowmachineconfig: %v", err)})
	}
	return result, reterr
}

func (r *SnowMachineConfigReconciler) reconcile(ctx context.Context, snowMachineConfig *anywherev1.SnowMachineConfig) (_ ctrl.Result, reterr error) {
	var allErrs []error
	deviceClientMap, err := r.clientBuilder.Build(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	// Setting the aws client map on every reconcile based on the secrets at that point of time
	validator := snow.NewValidator(deviceClientMap)
	if err := validator.ValidateMachineDeviceIPs(ctx, snowMachineConfig); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validator.ValidateEC2ImageExistsOnDevice(ctx, snowMachineConfig); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validator.ValidateEC2SshKeyNameExists(ctx, snowMachineConfig); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) > 0 {
		snowMachineConfig.Status.SpecValid = false
		aggregate := kerrors.NewAggregate(allErrs)
		failureMessage := aggregate.Error()
		snowMachineConfig.Status.FailureMessage = &failureMessage
		return ctrl.Result{}, aggregate
	}
	snowMachineConfig.Status.SpecValid = true
	snowMachineConfig.Status.FailureMessage = nil
	return ctrl.Result{}, nil
}
