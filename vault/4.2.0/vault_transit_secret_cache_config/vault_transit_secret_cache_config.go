// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_transit_secret_cache_config

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource vault_transit_secret_cache_config.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultTransitSecretCacheConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vtscc *Resource) Type() string {
	return "vault_transit_secret_cache_config"
}

// LocalName returns the local name for [Resource].
func (vtscc *Resource) LocalName() string {
	return vtscc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vtscc *Resource) Configuration() interface{} {
	return vtscc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vtscc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vtscc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vtscc *Resource) Dependencies() terra.Dependencies {
	return vtscc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vtscc *Resource) LifecycleManagement() *terra.Lifecycle {
	return vtscc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vtscc *Resource) Attributes() vaultTransitSecretCacheConfigAttributes {
	return vaultTransitSecretCacheConfigAttributes{ref: terra.ReferenceResource(vtscc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vtscc *Resource) ImportState(state io.Reader) error {
	vtscc.state = &vaultTransitSecretCacheConfigState{}
	if err := json.NewDecoder(state).Decode(vtscc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vtscc.Type(), vtscc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vtscc *Resource) State() (*vaultTransitSecretCacheConfigState, bool) {
	return vtscc.state, vtscc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vtscc *Resource) StateMust() *vaultTransitSecretCacheConfigState {
	if vtscc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vtscc.Type(), vtscc.LocalName()))
	}
	return vtscc.state
}

// Args contains the configurations for vault_transit_secret_cache_config.
type Args struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Size: number, required
	Size terra.NumberValue `hcl:"size,attr" validate:"required"`
}

type vaultTransitSecretCacheConfigAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_transit_secret_cache_config.
func (vtscc vaultTransitSecretCacheConfigAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vtscc.ref.Append("backend"))
}

// Id returns a reference to field id of vault_transit_secret_cache_config.
func (vtscc vaultTransitSecretCacheConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vtscc.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_transit_secret_cache_config.
func (vtscc vaultTransitSecretCacheConfigAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vtscc.ref.Append("namespace"))
}

// Size returns a reference to field size of vault_transit_secret_cache_config.
func (vtscc vaultTransitSecretCacheConfigAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(vtscc.ref.Append("size"))
}

type vaultTransitSecretCacheConfigState struct {
	Backend   string  `json:"backend"`
	Id        string  `json:"id"`
	Namespace string  `json:"namespace"`
	Size      float64 `json:"size"`
}
