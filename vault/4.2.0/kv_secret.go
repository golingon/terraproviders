// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewKvSecret creates a new instance of [KvSecret].
func NewKvSecret(name string, args KvSecretArgs) *KvSecret {
	return &KvSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KvSecret)(nil)

// KvSecret represents the Terraform resource vault_kv_secret.
type KvSecret struct {
	Name      string
	Args      KvSecretArgs
	state     *kvSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KvSecret].
func (ks *KvSecret) Type() string {
	return "vault_kv_secret"
}

// LocalName returns the local name for [KvSecret].
func (ks *KvSecret) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [KvSecret].
func (ks *KvSecret) Configuration() interface{} {
	return ks.Args
}

// DependOn is used for other resources to depend on [KvSecret].
func (ks *KvSecret) DependOn() terra.Reference {
	return terra.ReferenceResource(ks)
}

// Dependencies returns the list of resources [KvSecret] depends_on.
func (ks *KvSecret) Dependencies() terra.Dependencies {
	return ks.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KvSecret].
func (ks *KvSecret) LifecycleManagement() *terra.Lifecycle {
	return ks.Lifecycle
}

// Attributes returns the attributes for [KvSecret].
func (ks *KvSecret) Attributes() kvSecretAttributes {
	return kvSecretAttributes{ref: terra.ReferenceResource(ks)}
}

// ImportState imports the given attribute values into [KvSecret]'s state.
func (ks *KvSecret) ImportState(av io.Reader) error {
	ks.state = &kvSecretState{}
	if err := json.NewDecoder(av).Decode(ks.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ks.Type(), ks.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KvSecret] has state.
func (ks *KvSecret) State() (*kvSecretState, bool) {
	return ks.state, ks.state != nil
}

// StateMust returns the state for [KvSecret]. Panics if the state is nil.
func (ks *KvSecret) StateMust() *kvSecretState {
	if ks.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ks.Type(), ks.LocalName()))
	}
	return ks.state
}

// KvSecretArgs contains the configurations for vault_kv_secret.
type KvSecretArgs struct {
	// DataJson: string, required
	DataJson terra.StringValue `hcl:"data_json,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}
type kvSecretAttributes struct {
	ref terra.Reference
}

// Data returns a reference to field data of vault_kv_secret.
func (ks kvSecretAttributes) Data() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("data"))
}

// DataJson returns a reference to field data_json of vault_kv_secret.
func (ks kvSecretAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("data_json"))
}

// Id returns a reference to field id of vault_kv_secret.
func (ks kvSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_kv_secret.
func (ks kvSecretAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_kv_secret.
func (ks kvSecretAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("path"))
}

type kvSecretState struct {
	Data      map[string]string `json:"data"`
	DataJson  string            `json:"data_json"`
	Id        string            `json:"id"`
	Namespace string            `json:"namespace"`
	Path      string            `json:"path"`
}
