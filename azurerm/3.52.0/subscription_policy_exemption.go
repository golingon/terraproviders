// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptionpolicyexemption "github.com/golingon/terraproviders/azurerm/3.52.0/subscriptionpolicyexemption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscriptionPolicyExemption creates a new instance of [SubscriptionPolicyExemption].
func NewSubscriptionPolicyExemption(name string, args SubscriptionPolicyExemptionArgs) *SubscriptionPolicyExemption {
	return &SubscriptionPolicyExemption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionPolicyExemption)(nil)

// SubscriptionPolicyExemption represents the Terraform resource azurerm_subscription_policy_exemption.
type SubscriptionPolicyExemption struct {
	Name      string
	Args      SubscriptionPolicyExemptionArgs
	state     *subscriptionPolicyExemptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) Type() string {
	return "azurerm_subscription_policy_exemption"
}

// LocalName returns the local name for [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) LocalName() string {
	return spe.Name
}

// Configuration returns the configuration (args) for [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) Configuration() interface{} {
	return spe.Args
}

// DependOn is used for other resources to depend on [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) DependOn() terra.Reference {
	return terra.ReferenceResource(spe)
}

// Dependencies returns the list of resources [SubscriptionPolicyExemption] depends_on.
func (spe *SubscriptionPolicyExemption) Dependencies() terra.Dependencies {
	return spe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) LifecycleManagement() *terra.Lifecycle {
	return spe.Lifecycle
}

// Attributes returns the attributes for [SubscriptionPolicyExemption].
func (spe *SubscriptionPolicyExemption) Attributes() subscriptionPolicyExemptionAttributes {
	return subscriptionPolicyExemptionAttributes{ref: terra.ReferenceResource(spe)}
}

// ImportState imports the given attribute values into [SubscriptionPolicyExemption]'s state.
func (spe *SubscriptionPolicyExemption) ImportState(av io.Reader) error {
	spe.state = &subscriptionPolicyExemptionState{}
	if err := json.NewDecoder(av).Decode(spe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spe.Type(), spe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubscriptionPolicyExemption] has state.
func (spe *SubscriptionPolicyExemption) State() (*subscriptionPolicyExemptionState, bool) {
	return spe.state, spe.state != nil
}

// StateMust returns the state for [SubscriptionPolicyExemption]. Panics if the state is nil.
func (spe *SubscriptionPolicyExemption) StateMust() *subscriptionPolicyExemptionState {
	if spe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spe.Type(), spe.LocalName()))
	}
	return spe.state
}

// SubscriptionPolicyExemptionArgs contains the configurations for azurerm_subscription_policy_exemption.
type SubscriptionPolicyExemptionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// ExemptionCategory: string, required
	ExemptionCategory terra.StringValue `hcl:"exemption_category,attr" validate:"required"`
	// ExpiresOn: string, optional
	ExpiresOn terra.StringValue `hcl:"expires_on,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyAssignmentId: string, required
	PolicyAssignmentId terra.StringValue `hcl:"policy_assignment_id,attr" validate:"required"`
	// PolicyDefinitionReferenceIds: list of string, optional
	PolicyDefinitionReferenceIds terra.ListValue[terra.StringValue] `hcl:"policy_definition_reference_ids,attr"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subscriptionpolicyexemption.Timeouts `hcl:"timeouts,block"`
}
type subscriptionPolicyExemptionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("display_name"))
}

// ExemptionCategory returns a reference to field exemption_category of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) ExemptionCategory() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("exemption_category"))
}

// ExpiresOn returns a reference to field expires_on of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) ExpiresOn() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("expires_on"))
}

// Id returns a reference to field id of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("name"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionReferenceIds returns a reference to field policy_definition_reference_ids of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) PolicyDefinitionReferenceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spe.ref.Append("policy_definition_reference_ids"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription_policy_exemption.
func (spe subscriptionPolicyExemptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(spe.ref.Append("subscription_id"))
}

func (spe subscriptionPolicyExemptionAttributes) Timeouts() subscriptionpolicyexemption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscriptionpolicyexemption.TimeoutsAttributes](spe.ref.Append("timeouts"))
}

type subscriptionPolicyExemptionState struct {
	Description                  string                                     `json:"description"`
	DisplayName                  string                                     `json:"display_name"`
	ExemptionCategory            string                                     `json:"exemption_category"`
	ExpiresOn                    string                                     `json:"expires_on"`
	Id                           string                                     `json:"id"`
	Metadata                     string                                     `json:"metadata"`
	Name                         string                                     `json:"name"`
	PolicyAssignmentId           string                                     `json:"policy_assignment_id"`
	PolicyDefinitionReferenceIds []string                                   `json:"policy_definition_reference_ids"`
	SubscriptionId               string                                     `json:"subscription_id"`
	Timeouts                     *subscriptionpolicyexemption.TimeoutsState `json:"timeouts"`
}
