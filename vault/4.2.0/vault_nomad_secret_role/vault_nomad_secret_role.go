// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_nomad_secret_role

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

// Resource represents the Terraform resource vault_nomad_secret_role.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultNomadSecretRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vnsr *Resource) Type() string {
	return "vault_nomad_secret_role"
}

// LocalName returns the local name for [Resource].
func (vnsr *Resource) LocalName() string {
	return vnsr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vnsr *Resource) Configuration() interface{} {
	return vnsr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vnsr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vnsr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vnsr *Resource) Dependencies() terra.Dependencies {
	return vnsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vnsr *Resource) LifecycleManagement() *terra.Lifecycle {
	return vnsr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vnsr *Resource) Attributes() vaultNomadSecretRoleAttributes {
	return vaultNomadSecretRoleAttributes{ref: terra.ReferenceResource(vnsr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vnsr *Resource) ImportState(state io.Reader) error {
	vnsr.state = &vaultNomadSecretRoleState{}
	if err := json.NewDecoder(state).Decode(vnsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnsr.Type(), vnsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vnsr *Resource) State() (*vaultNomadSecretRoleState, bool) {
	return vnsr.state, vnsr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vnsr *Resource) StateMust() *vaultNomadSecretRoleState {
	if vnsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnsr.Type(), vnsr.LocalName()))
	}
	return vnsr.state
}

// Args contains the configurations for vault_nomad_secret_role.
type Args struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Global: bool, optional
	Global terra.BoolValue `hcl:"global,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policies: list of string, optional
	Policies terra.ListValue[terra.StringValue] `hcl:"policies,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type vaultNomadSecretRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vnsr.ref.Append("backend"))
}

// Global returns a reference to field global of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Global() terra.BoolValue {
	return terra.ReferenceAsBool(vnsr.ref.Append("global"))
}

// Id returns a reference to field id of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnsr.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vnsr.ref.Append("namespace"))
}

// Policies returns a reference to field policies of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Policies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vnsr.ref.Append("policies"))
}

// Role returns a reference to field role of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vnsr.ref.Append("role"))
}

// Type returns a reference to field type of vault_nomad_secret_role.
func (vnsr vaultNomadSecretRoleAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vnsr.ref.Append("type"))
}

type vaultNomadSecretRoleState struct {
	Backend   string   `json:"backend"`
	Global    bool     `json:"global"`
	Id        string   `json:"id"`
	Namespace string   `json:"namespace"`
	Policies  []string `json:"policies"`
	Role      string   `json:"role"`
	Type      string   `json:"type"`
}
