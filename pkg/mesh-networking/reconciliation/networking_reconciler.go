package reconciliation

import (
	"context"
	"github.com/solo-io/smh/pkg/mesh-networking/translation/istio"
	"github.com/solo-io/smh/pkg/mesh-networking/translation/reporter"
	"github.com/solo-io/smh/pkg/mesh-networking/validation"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/snapshot/input"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type networkingReconciler struct {
	ctx          context.Context
	builder      input.Builder
	validator    validation.Validator
	reporter     reporter.Reporter
	translator   istio.Translator
	masterClient client.Client
}

func Start(
	ctx context.Context,
	builder input.Builder,
	translator istio.Translator,
	masterClient client.Client,
	mgr manager.Manager,
) error {
	d := &networkingReconciler{
		ctx:          ctx,
		builder:      builder,
		translator:   translator,
		masterClient: masterClient,
	}

	return input.RegisterSingleClusterReconciler(ctx, mgr, d.reconcile)
}

// reconcile global state
func (d *networkingReconciler) reconcile() error {
	inputSnap, err := d.builder.BuildSnapshot(d.ctx, "mesh-networking")
	if err != nil {
		// failed to read from cache; should never happen
		return err
	}

	translatorSnapshot := d.validator.Validate(d.ctx, inputSnap)

	istioSnap, err := d.translator.Translate(translatorSnapshot, d.reporter)
	if err != nil {
		// internal translator errors should never happen
		return err
	}

	if err := istioSnap.Apply(d.ctx, d.masterClient); err != nil {
		return err
	}

	return inputSnap.SyncStatuses(d.ctx)
}