// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIdentityEntityPolicies creates a new instance of [IdentityEntityPolicies].
func NewIdentityEntityPolicies(name string, args IdentityEntityPoliciesArgs) *IdentityEntityPolicies {
	return &IdentityEntityPolicies{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityEntityPolicies)(nil)

// IdentityEntityPolicies represents the Terraform resource vault_identity_entity_policies.
type IdentityEntityPolicies struct {
	Name      string
	Args      IdentityEntityPoliciesArgs
	state     *identityEntityPoliciesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) Type() string {
	return "vault_identity_entity_policies"
}

// LocalName returns the local name for [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) LocalName() string {
	return iep.Name
}

// Configuration returns the configuration (args) for [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) Configuration() interface{} {
	return iep.Args
}

// DependOn is used for other resources to depend on [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) DependOn() terra.Reference {
	return terra.ReferenceResource(iep)
}

// Dependencies returns the list of resources [IdentityEntityPolicies] depends_on.
func (iep *IdentityEntityPolicies) Dependencies() terra.Dependencies {
	return iep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) LifecycleManagement() *terra.Lifecycle {
	return iep.Lifecycle
}

// Attributes returns the attributes for [IdentityEntityPolicies].
func (iep *IdentityEntityPolicies) Attributes() identityEntityPoliciesAttributes {
	return identityEntityPoliciesAttributes{ref: terra.ReferenceResource(iep)}
}

// ImportState imports the given attribute values into [IdentityEntityPolicies]'s state.
func (iep *IdentityEntityPolicies) ImportState(av io.Reader) error {
	iep.state = &identityEntityPoliciesState{}
	if err := json.NewDecoder(av).Decode(iep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iep.Type(), iep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityEntityPolicies] has state.
func (iep *IdentityEntityPolicies) State() (*identityEntityPoliciesState, bool) {
	return iep.state, iep.state != nil
}

// StateMust returns the state for [IdentityEntityPolicies]. Panics if the state is nil.
func (iep *IdentityEntityPolicies) StateMust() *identityEntityPoliciesState {
	if iep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iep.Type(), iep.LocalName()))
	}
	return iep.state
}

// IdentityEntityPoliciesArgs contains the configurations for vault_identity_entity_policies.
type IdentityEntityPoliciesArgs struct {
	// EntityId: string, required
	EntityId terra.StringValue `hcl:"entity_id,attr" validate:"required"`
	// Exclusive: bool, optional
	Exclusive terra.BoolValue `hcl:"exclusive,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policies: set of string, required
	Policies terra.SetValue[terra.StringValue] `hcl:"policies,attr" validate:"required"`
}
type identityEntityPoliciesAttributes struct {
	ref terra.Reference
}

// EntityId returns a reference to field entity_id of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) EntityId() terra.StringValue {
	return terra.ReferenceAsString(iep.ref.Append("entity_id"))
}

// EntityName returns a reference to field entity_name of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) EntityName() terra.StringValue {
	return terra.ReferenceAsString(iep.ref.Append("entity_name"))
}

// Exclusive returns a reference to field exclusive of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) Exclusive() terra.BoolValue {
	return terra.ReferenceAsBool(iep.ref.Append("exclusive"))
}

// Id returns a reference to field id of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iep.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(iep.ref.Append("namespace"))
}

// Policies returns a reference to field policies of vault_identity_entity_policies.
func (iep identityEntityPoliciesAttributes) Policies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iep.ref.Append("policies"))
}

type identityEntityPoliciesState struct {
	EntityId   string   `json:"entity_id"`
	EntityName string   `json:"entity_name"`
	Exclusive  bool     `json:"exclusive"`
	Id         string   `json:"id"`
	Namespace  string   `json:"namespace"`
	Policies   []string `json:"policies"`
}
