// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	policysetdefinition "github.com/golingon/terraproviders/azurerm/3.49.0/policysetdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPolicySetDefinition creates a new instance of [PolicySetDefinition].
func NewPolicySetDefinition(name string, args PolicySetDefinitionArgs) *PolicySetDefinition {
	return &PolicySetDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PolicySetDefinition)(nil)

// PolicySetDefinition represents the Terraform resource azurerm_policy_set_definition.
type PolicySetDefinition struct {
	Name      string
	Args      PolicySetDefinitionArgs
	state     *policySetDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PolicySetDefinition].
func (psd *PolicySetDefinition) Type() string {
	return "azurerm_policy_set_definition"
}

// LocalName returns the local name for [PolicySetDefinition].
func (psd *PolicySetDefinition) LocalName() string {
	return psd.Name
}

// Configuration returns the configuration (args) for [PolicySetDefinition].
func (psd *PolicySetDefinition) Configuration() interface{} {
	return psd.Args
}

// DependOn is used for other resources to depend on [PolicySetDefinition].
func (psd *PolicySetDefinition) DependOn() terra.Reference {
	return terra.ReferenceResource(psd)
}

// Dependencies returns the list of resources [PolicySetDefinition] depends_on.
func (psd *PolicySetDefinition) Dependencies() terra.Dependencies {
	return psd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PolicySetDefinition].
func (psd *PolicySetDefinition) LifecycleManagement() *terra.Lifecycle {
	return psd.Lifecycle
}

// Attributes returns the attributes for [PolicySetDefinition].
func (psd *PolicySetDefinition) Attributes() policySetDefinitionAttributes {
	return policySetDefinitionAttributes{ref: terra.ReferenceResource(psd)}
}

// ImportState imports the given attribute values into [PolicySetDefinition]'s state.
func (psd *PolicySetDefinition) ImportState(av io.Reader) error {
	psd.state = &policySetDefinitionState{}
	if err := json.NewDecoder(av).Decode(psd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psd.Type(), psd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PolicySetDefinition] has state.
func (psd *PolicySetDefinition) State() (*policySetDefinitionState, bool) {
	return psd.state, psd.state != nil
}

// StateMust returns the state for [PolicySetDefinition]. Panics if the state is nil.
func (psd *PolicySetDefinition) StateMust() *policySetDefinitionState {
	if psd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psd.Type(), psd.LocalName()))
	}
	return psd.state
}

// PolicySetDefinitionArgs contains the configurations for azurerm_policy_set_definition.
type PolicySetDefinitionArgs struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// PolicyType: string, required
	PolicyType terra.StringValue `hcl:"policy_type,attr" validate:"required"`
	// PolicyDefinitionGroup: min=0
	PolicyDefinitionGroup []policysetdefinition.PolicyDefinitionGroup `hcl:"policy_definition_group,block" validate:"min=0"`
	// PolicyDefinitionReference: min=1
	PolicyDefinitionReference []policysetdefinition.PolicyDefinitionReference `hcl:"policy_definition_reference,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *policysetdefinition.Timeouts `hcl:"timeouts,block"`
}
type policySetDefinitionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("management_group_id"))
}

// Metadata returns a reference to field metadata of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("parameters"))
}

// PolicyType returns a reference to field policy_type of azurerm_policy_set_definition.
func (psd policySetDefinitionAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("policy_type"))
}

func (psd policySetDefinitionAttributes) PolicyDefinitionGroup() terra.SetValue[policysetdefinition.PolicyDefinitionGroupAttributes] {
	return terra.ReferenceAsSet[policysetdefinition.PolicyDefinitionGroupAttributes](psd.ref.Append("policy_definition_group"))
}

func (psd policySetDefinitionAttributes) PolicyDefinitionReference() terra.ListValue[policysetdefinition.PolicyDefinitionReferenceAttributes] {
	return terra.ReferenceAsList[policysetdefinition.PolicyDefinitionReferenceAttributes](psd.ref.Append("policy_definition_reference"))
}

func (psd policySetDefinitionAttributes) Timeouts() policysetdefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[policysetdefinition.TimeoutsAttributes](psd.ref.Append("timeouts"))
}

type policySetDefinitionState struct {
	Description               string                                               `json:"description"`
	DisplayName               string                                               `json:"display_name"`
	Id                        string                                               `json:"id"`
	ManagementGroupId         string                                               `json:"management_group_id"`
	Metadata                  string                                               `json:"metadata"`
	Name                      string                                               `json:"name"`
	Parameters                string                                               `json:"parameters"`
	PolicyType                string                                               `json:"policy_type"`
	PolicyDefinitionGroup     []policysetdefinition.PolicyDefinitionGroupState     `json:"policy_definition_group"`
	PolicyDefinitionReference []policysetdefinition.PolicyDefinitionReferenceState `json:"policy_definition_reference"`
	Timeouts                  *policysetdefinition.TimeoutsState                   `json:"timeouts"`
}
