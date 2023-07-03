// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmipSecretScope creates a new instance of [KmipSecretScope].
func NewKmipSecretScope(name string, args KmipSecretScopeArgs) *KmipSecretScope {
	return &KmipSecretScope{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmipSecretScope)(nil)

// KmipSecretScope represents the Terraform resource vault_kmip_secret_scope.
type KmipSecretScope struct {
	Name      string
	Args      KmipSecretScopeArgs
	state     *kmipSecretScopeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmipSecretScope].
func (kss *KmipSecretScope) Type() string {
	return "vault_kmip_secret_scope"
}

// LocalName returns the local name for [KmipSecretScope].
func (kss *KmipSecretScope) LocalName() string {
	return kss.Name
}

// Configuration returns the configuration (args) for [KmipSecretScope].
func (kss *KmipSecretScope) Configuration() interface{} {
	return kss.Args
}

// DependOn is used for other resources to depend on [KmipSecretScope].
func (kss *KmipSecretScope) DependOn() terra.Reference {
	return terra.ReferenceResource(kss)
}

// Dependencies returns the list of resources [KmipSecretScope] depends_on.
func (kss *KmipSecretScope) Dependencies() terra.Dependencies {
	return kss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmipSecretScope].
func (kss *KmipSecretScope) LifecycleManagement() *terra.Lifecycle {
	return kss.Lifecycle
}

// Attributes returns the attributes for [KmipSecretScope].
func (kss *KmipSecretScope) Attributes() kmipSecretScopeAttributes {
	return kmipSecretScopeAttributes{ref: terra.ReferenceResource(kss)}
}

// ImportState imports the given attribute values into [KmipSecretScope]'s state.
func (kss *KmipSecretScope) ImportState(av io.Reader) error {
	kss.state = &kmipSecretScopeState{}
	if err := json.NewDecoder(av).Decode(kss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kss.Type(), kss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmipSecretScope] has state.
func (kss *KmipSecretScope) State() (*kmipSecretScopeState, bool) {
	return kss.state, kss.state != nil
}

// StateMust returns the state for [KmipSecretScope]. Panics if the state is nil.
func (kss *KmipSecretScope) StateMust() *kmipSecretScopeState {
	if kss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kss.Type(), kss.LocalName()))
	}
	return kss.state
}

// KmipSecretScopeArgs contains the configurations for vault_kmip_secret_scope.
type KmipSecretScopeArgs struct {
	// Force: bool, optional
	Force terra.BoolValue `hcl:"force,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
}
type kmipSecretScopeAttributes struct {
	ref terra.Reference
}

// Force returns a reference to field force of vault_kmip_secret_scope.
func (kss kmipSecretScopeAttributes) Force() terra.BoolValue {
	return terra.ReferenceAsBool(kss.ref.Append("force"))
}

// Id returns a reference to field id of vault_kmip_secret_scope.
func (kss kmipSecretScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kss.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_kmip_secret_scope.
func (kss kmipSecretScopeAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(kss.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kmip_secret_scope.
func (kss kmipSecretScopeAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(kss.ref.Append("path"))
}

// Scope returns a reference to field scope of vault_kmip_secret_scope.
func (kss kmipSecretScopeAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(kss.ref.Append("scope"))
}

type kmipSecretScopeState struct {
	Force     bool   `json:"force"`
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
	Path      string `json:"path"`
	Scope     string `json:"scope"`
}
