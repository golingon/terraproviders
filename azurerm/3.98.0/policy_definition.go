// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	policydefinition "github.com/golingon/terraproviders/azurerm/3.98.0/policydefinition"
	"io"
)

// NewPolicyDefinition creates a new instance of [PolicyDefinition].
func NewPolicyDefinition(name string, args PolicyDefinitionArgs) *PolicyDefinition {
	return &PolicyDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PolicyDefinition)(nil)

// PolicyDefinition represents the Terraform resource azurerm_policy_definition.
type PolicyDefinition struct {
	Name      string
	Args      PolicyDefinitionArgs
	state     *policyDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PolicyDefinition].
func (pd *PolicyDefinition) Type() string {
	return "azurerm_policy_definition"
}

// LocalName returns the local name for [PolicyDefinition].
func (pd *PolicyDefinition) LocalName() string {
	return pd.Name
}

// Configuration returns the configuration (args) for [PolicyDefinition].
func (pd *PolicyDefinition) Configuration() interface{} {
	return pd.Args
}

// DependOn is used for other resources to depend on [PolicyDefinition].
func (pd *PolicyDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(pd)
}

// Dependencies returns the list of resources [PolicyDefinition] depends_on.
func (pd *PolicyDefinition) Dependencies() terra.Dependencies {
	return pd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PolicyDefinition].
func (pd *PolicyDefinition) LifecycleManagement() *terra.Lifecycle {
	return pd.Lifecycle
}

// Attributes returns the attributes for [PolicyDefinition].
func (pd *PolicyDefinition) Attributes() policyDefinitionAttributes {
	return policyDefinitionAttributes{ref: terra.ReferenceResource(pd)}
}

// ImportState imports the given attribute values into [PolicyDefinition]'s state.
func (pd *PolicyDefinition) ImportState(av io.Reader) error {
	pd.state = &policyDefinitionState{}
	if err := json.NewDecoder(av).Decode(pd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pd.Type(), pd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PolicyDefinition] has state.
func (pd *PolicyDefinition) State() (*policyDefinitionState, bool) {
	return pd.state, pd.state != nil
}

// StateMust returns the state for [PolicyDefinition]. Panics if the state is nil.
func (pd *PolicyDefinition) StateMust() *policyDefinitionState {
	if pd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pd.Type(), pd.LocalName()))
	}
	return pd.state
}

// PolicyDefinitionArgs contains the configurations for azurerm_policy_definition.
type PolicyDefinitionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, optional
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// PolicyRule: string, optional
	PolicyRule terra.StringValue `hcl:"policy_rule,attr"`
	// PolicyType: string, required
	PolicyType terra.StringValue `hcl:"policy_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *policydefinition.Timeouts `hcl:"timeouts,block"`
}
type policyDefinitionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_definition.
func (pd policyDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_policy_definition.
func (pd policyDefinitionAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("management_group_id"))
}

// Metadata returns a reference to field metadata of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("metadata"))
}

// Mode returns a reference to field mode of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_policy_definition.
func (pd policyDefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("parameters"))
}

// PolicyRule returns a reference to field policy_rule of azurerm_policy_definition.
func (pd policyDefinitionAttributes) PolicyRule() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("policy_rule"))
}

// PolicyType returns a reference to field policy_type of azurerm_policy_definition.
func (pd policyDefinitionAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("policy_type"))
}

// RoleDefinitionIds returns a reference to field role_definition_ids of azurerm_policy_definition.
func (pd policyDefinitionAttributes) RoleDefinitionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pd.ref.Append("role_definition_ids"))
}

func (pd policyDefinitionAttributes) Timeouts() policydefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[policydefinition.TimeoutsAttributes](pd.ref.Append("timeouts"))
}

type policyDefinitionState struct {
	Description       string                          `json:"description"`
	DisplayName       string                          `json:"display_name"`
	Id                string                          `json:"id"`
	ManagementGroupId string                          `json:"management_group_id"`
	Metadata          string                          `json:"metadata"`
	Mode              string                          `json:"mode"`
	Name              string                          `json:"name"`
	Parameters        string                          `json:"parameters"`
	PolicyRule        string                          `json:"policy_rule"`
	PolicyType        string                          `json:"policy_type"`
	RoleDefinitionIds []string                        `json:"role_definition_ids"`
	Timeouts          *policydefinition.TimeoutsState `json:"timeouts"`
}
