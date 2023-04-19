// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sshpublickey "github.com/golingon/terraproviders/azurerm/3.52.0/sshpublickey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSshPublicKey creates a new instance of [SshPublicKey].
func NewSshPublicKey(name string, args SshPublicKeyArgs) *SshPublicKey {
	return &SshPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SshPublicKey)(nil)

// SshPublicKey represents the Terraform resource azurerm_ssh_public_key.
type SshPublicKey struct {
	Name      string
	Args      SshPublicKeyArgs
	state     *sshPublicKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SshPublicKey].
func (spk *SshPublicKey) Type() string {
	return "azurerm_ssh_public_key"
}

// LocalName returns the local name for [SshPublicKey].
func (spk *SshPublicKey) LocalName() string {
	return spk.Name
}

// Configuration returns the configuration (args) for [SshPublicKey].
func (spk *SshPublicKey) Configuration() interface{} {
	return spk.Args
}

// DependOn is used for other resources to depend on [SshPublicKey].
func (spk *SshPublicKey) DependOn() terra.Reference {
	return terra.ReferenceResource(spk)
}

// Dependencies returns the list of resources [SshPublicKey] depends_on.
func (spk *SshPublicKey) Dependencies() terra.Dependencies {
	return spk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SshPublicKey].
func (spk *SshPublicKey) LifecycleManagement() *terra.Lifecycle {
	return spk.Lifecycle
}

// Attributes returns the attributes for [SshPublicKey].
func (spk *SshPublicKey) Attributes() sshPublicKeyAttributes {
	return sshPublicKeyAttributes{ref: terra.ReferenceResource(spk)}
}

// ImportState imports the given attribute values into [SshPublicKey]'s state.
func (spk *SshPublicKey) ImportState(av io.Reader) error {
	spk.state = &sshPublicKeyState{}
	if err := json.NewDecoder(av).Decode(spk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spk.Type(), spk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SshPublicKey] has state.
func (spk *SshPublicKey) State() (*sshPublicKeyState, bool) {
	return spk.state, spk.state != nil
}

// StateMust returns the state for [SshPublicKey]. Panics if the state is nil.
func (spk *SshPublicKey) StateMust() *sshPublicKeyState {
	if spk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spk.Type(), spk.LocalName()))
	}
	return spk.state
}

// SshPublicKeyArgs contains the configurations for azurerm_ssh_public_key.
type SshPublicKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicKey: string, required
	PublicKey terra.StringValue `hcl:"public_key,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *sshpublickey.Timeouts `hcl:"timeouts,block"`
}
type sshPublicKeyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("name"))
}

// PublicKey returns a reference to field public_key of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("public_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_ssh_public_key.
func (spk sshPublicKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](spk.ref.Append("tags"))
}

func (spk sshPublicKeyAttributes) Timeouts() sshpublickey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sshpublickey.TimeoutsAttributes](spk.ref.Append("timeouts"))
}

type sshPublicKeyState struct {
	Id                string                      `json:"id"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	PublicKey         string                      `json:"public_key"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Timeouts          *sshpublickey.TimeoutsState `json:"timeouts"`
}