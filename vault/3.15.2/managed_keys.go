// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	managedkeys "github.com/golingon/terraproviders/vault/3.15.2/managedkeys"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagedKeys creates a new instance of [ManagedKeys].
func NewManagedKeys(name string, args ManagedKeysArgs) *ManagedKeys {
	return &ManagedKeys{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagedKeys)(nil)

// ManagedKeys represents the Terraform resource vault_managed_keys.
type ManagedKeys struct {
	Name      string
	Args      ManagedKeysArgs
	state     *managedKeysState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagedKeys].
func (mk *ManagedKeys) Type() string {
	return "vault_managed_keys"
}

// LocalName returns the local name for [ManagedKeys].
func (mk *ManagedKeys) LocalName() string {
	return mk.Name
}

// Configuration returns the configuration (args) for [ManagedKeys].
func (mk *ManagedKeys) Configuration() interface{} {
	return mk.Args
}

// DependOn is used for other resources to depend on [ManagedKeys].
func (mk *ManagedKeys) DependOn() terra.Reference {
	return terra.ReferenceResource(mk)
}

// Dependencies returns the list of resources [ManagedKeys] depends_on.
func (mk *ManagedKeys) Dependencies() terra.Dependencies {
	return mk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagedKeys].
func (mk *ManagedKeys) LifecycleManagement() *terra.Lifecycle {
	return mk.Lifecycle
}

// Attributes returns the attributes for [ManagedKeys].
func (mk *ManagedKeys) Attributes() managedKeysAttributes {
	return managedKeysAttributes{ref: terra.ReferenceResource(mk)}
}

// ImportState imports the given attribute values into [ManagedKeys]'s state.
func (mk *ManagedKeys) ImportState(av io.Reader) error {
	mk.state = &managedKeysState{}
	if err := json.NewDecoder(av).Decode(mk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mk.Type(), mk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagedKeys] has state.
func (mk *ManagedKeys) State() (*managedKeysState, bool) {
	return mk.state, mk.state != nil
}

// StateMust returns the state for [ManagedKeys]. Panics if the state is nil.
func (mk *ManagedKeys) StateMust() *managedKeysState {
	if mk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mk.Type(), mk.LocalName()))
	}
	return mk.state
}

// ManagedKeysArgs contains the configurations for vault_managed_keys.
type ManagedKeysArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Aws: min=0
	Aws []managedkeys.Aws `hcl:"aws,block" validate:"min=0"`
	// Azure: min=0
	Azure []managedkeys.Azure `hcl:"azure,block" validate:"min=0"`
	// Pkcs: min=0
	Pkcs []managedkeys.Pkcs `hcl:"pkcs,block" validate:"min=0"`
}
type managedKeysAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of vault_managed_keys.
func (mk managedKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mk.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_managed_keys.
func (mk managedKeysAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(mk.ref.Append("namespace"))
}

func (mk managedKeysAttributes) Aws() terra.SetValue[managedkeys.AwsAttributes] {
	return terra.ReferenceAsSet[managedkeys.AwsAttributes](mk.ref.Append("aws"))
}

func (mk managedKeysAttributes) Azure() terra.SetValue[managedkeys.AzureAttributes] {
	return terra.ReferenceAsSet[managedkeys.AzureAttributes](mk.ref.Append("azure"))
}

func (mk managedKeysAttributes) Pkcs() terra.SetValue[managedkeys.PkcsAttributes] {
	return terra.ReferenceAsSet[managedkeys.PkcsAttributes](mk.ref.Append("pkcs"))
}

type managedKeysState struct {
	Id        string                   `json:"id"`
	Namespace string                   `json:"namespace"`
	Aws       []managedkeys.AwsState   `json:"aws"`
	Azure     []managedkeys.AzureState `json:"azure"`
	Pkcs      []managedkeys.PkcsState  `json:"pkcs"`
}
