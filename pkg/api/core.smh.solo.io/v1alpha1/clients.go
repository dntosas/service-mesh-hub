// Code generated by skv2. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// clienset for the core.smh.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the core.smh.solo.io/v1alpha1/v1alpha1 APIs
	Settings() SettingsClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (*clientSet, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) *clientSet {
	return &clientSet{client: client}
}

// clienset for the core.smh.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) Settings() SettingsClient {
	return NewSettingsClient(c.client)
}

// Reader knows how to read and list Settingss.
type SettingsReader interface {
	// Get retrieves a Settings for the given object key
	GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error)

	// List retrieves list of Settingss for a given namespace and list options.
	ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error)
}

// Writer knows how to create, delete, and update Settingss.
type SettingsWriter interface {
	// Create saves the Settings object.
	CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error

	// Delete deletes the Settings object.
	DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Settings object.
	UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

	// If the Settings object exists, update its spec. Otherwise, create the Settings object.
	UpsertSettingsSpec(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

	// Patch patches the given Settings object.
	PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Settings objects matching the given options.
	DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a Settings object.
type SettingsStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Settings object.
	UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error

	// Patch patches the given Settings object's subresource.
	PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Settingss.
type SettingsClient interface {
	SettingsReader
	SettingsWriter
	SettingsStatusWriter
}

type settingsClient struct {
	client client.Client
}

func NewSettingsClient(client client.Client) *settingsClient {
	return &settingsClient{client: client}
}

func (c *settingsClient) GetSettings(ctx context.Context, key client.ObjectKey) (*Settings, error) {
	obj := &Settings{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *settingsClient) ListSettings(ctx context.Context, opts ...client.ListOption) (*SettingsList, error) {
	list := &SettingsList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *settingsClient) CreateSettings(ctx context.Context, obj *Settings, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *settingsClient) DeleteSettings(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Settings{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *settingsClient) UpdateSettings(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *settingsClient) UpsertSettingsSpec(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
	existing, err := c.GetSettings(ctx, client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateSettings(ctx, obj)
		}
		return err
	}
	existing.Spec = obj.Spec
	return c.client.Update(ctx, existing, opts...)
}

func (c *settingsClient) PatchSettings(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *settingsClient) DeleteAllOfSettings(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Settings{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *settingsClient) UpdateSettingsStatus(ctx context.Context, obj *Settings, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *settingsClient) PatchSettingsStatus(ctx context.Context, obj *Settings, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}