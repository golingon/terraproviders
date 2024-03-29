// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptionpolicyassignment "github.com/golingon/terraproviders/azurerm/3.66.0/subscriptionpolicyassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscriptionPolicyAssignment creates a new instance of [SubscriptionPolicyAssignment].
func NewSubscriptionPolicyAssignment(name string, args SubscriptionPolicyAssignmentArgs) *SubscriptionPolicyAssignment {
	return &SubscriptionPolicyAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionPolicyAssignment)(nil)

// SubscriptionPolicyAssignment represents the Terraform resource azurerm_subscription_policy_assignment.
type SubscriptionPolicyAssignment struct {
	Name      string
	Args      SubscriptionPolicyAssignmentArgs
	state     *subscriptionPolicyAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) Type() string {
	return "azurerm_subscription_policy_assignment"
}

// LocalName returns the local name for [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) LocalName() string {
	return spa.Name
}

// Configuration returns the configuration (args) for [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) Configuration() interface{} {
	return spa.Args
}

// DependOn is used for other resources to depend on [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(spa)
}

// Dependencies returns the list of resources [SubscriptionPolicyAssignment] depends_on.
func (spa *SubscriptionPolicyAssignment) Dependencies() terra.Dependencies {
	return spa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) LifecycleManagement() *terra.Lifecycle {
	return spa.Lifecycle
}

// Attributes returns the attributes for [SubscriptionPolicyAssignment].
func (spa *SubscriptionPolicyAssignment) Attributes() subscriptionPolicyAssignmentAttributes {
	return subscriptionPolicyAssignmentAttributes{ref: terra.ReferenceResource(spa)}
}

// ImportState imports the given attribute values into [SubscriptionPolicyAssignment]'s state.
func (spa *SubscriptionPolicyAssignment) ImportState(av io.Reader) error {
	spa.state = &subscriptionPolicyAssignmentState{}
	if err := json.NewDecoder(av).Decode(spa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spa.Type(), spa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubscriptionPolicyAssignment] has state.
func (spa *SubscriptionPolicyAssignment) State() (*subscriptionPolicyAssignmentState, bool) {
	return spa.state, spa.state != nil
}

// StateMust returns the state for [SubscriptionPolicyAssignment]. Panics if the state is nil.
func (spa *SubscriptionPolicyAssignment) StateMust() *subscriptionPolicyAssignmentState {
	if spa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spa.Type(), spa.LocalName()))
	}
	return spa.state
}

// SubscriptionPolicyAssignmentArgs contains the configurations for azurerm_subscription_policy_assignment.
type SubscriptionPolicyAssignmentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Enforce: bool, optional
	Enforce terra.BoolValue `hcl:"enforce,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotScopes: list of string, optional
	NotScopes terra.ListValue[terra.StringValue] `hcl:"not_scopes,attr"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// PolicyDefinitionId: string, required
	PolicyDefinitionId terra.StringValue `hcl:"policy_definition_id,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Identity: optional
	Identity *subscriptionpolicyassignment.Identity `hcl:"identity,block"`
	// NonComplianceMessage: min=0
	NonComplianceMessage []subscriptionpolicyassignment.NonComplianceMessage `hcl:"non_compliance_message,block" validate:"min=0"`
	// Overrides: min=0
	Overrides []subscriptionpolicyassignment.Overrides `hcl:"overrides,block" validate:"min=0"`
	// ResourceSelectors: min=0
	ResourceSelectors []subscriptionpolicyassignment.ResourceSelectors `hcl:"resource_selectors,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *subscriptionpolicyassignment.Timeouts `hcl:"timeouts,block"`
}
type subscriptionPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("display_name"))
}

// Enforce returns a reference to field enforce of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Enforce() terra.BoolValue {
	return terra.ReferenceAsBool(spa.ref.Append("enforce"))
}

// Id returns a reference to field id of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("location"))
}

// Metadata returns a reference to field metadata of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("name"))
}

// NotScopes returns a reference to field not_scopes of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) NotScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spa.ref.Append("not_scopes"))
}

// Parameters returns a reference to field parameters of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("parameters"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("policy_definition_id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription_policy_assignment.
func (spa subscriptionPolicyAssignmentAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("subscription_id"))
}

func (spa subscriptionPolicyAssignmentAttributes) Identity() terra.ListValue[subscriptionpolicyassignment.IdentityAttributes] {
	return terra.ReferenceAsList[subscriptionpolicyassignment.IdentityAttributes](spa.ref.Append("identity"))
}

func (spa subscriptionPolicyAssignmentAttributes) NonComplianceMessage() terra.ListValue[subscriptionpolicyassignment.NonComplianceMessageAttributes] {
	return terra.ReferenceAsList[subscriptionpolicyassignment.NonComplianceMessageAttributes](spa.ref.Append("non_compliance_message"))
}

func (spa subscriptionPolicyAssignmentAttributes) Overrides() terra.ListValue[subscriptionpolicyassignment.OverridesAttributes] {
	return terra.ReferenceAsList[subscriptionpolicyassignment.OverridesAttributes](spa.ref.Append("overrides"))
}

func (spa subscriptionPolicyAssignmentAttributes) ResourceSelectors() terra.ListValue[subscriptionpolicyassignment.ResourceSelectorsAttributes] {
	return terra.ReferenceAsList[subscriptionpolicyassignment.ResourceSelectorsAttributes](spa.ref.Append("resource_selectors"))
}

func (spa subscriptionPolicyAssignmentAttributes) Timeouts() subscriptionpolicyassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscriptionpolicyassignment.TimeoutsAttributes](spa.ref.Append("timeouts"))
}

type subscriptionPolicyAssignmentState struct {
	Description          string                                                   `json:"description"`
	DisplayName          string                                                   `json:"display_name"`
	Enforce              bool                                                     `json:"enforce"`
	Id                   string                                                   `json:"id"`
	Location             string                                                   `json:"location"`
	Metadata             string                                                   `json:"metadata"`
	Name                 string                                                   `json:"name"`
	NotScopes            []string                                                 `json:"not_scopes"`
	Parameters           string                                                   `json:"parameters"`
	PolicyDefinitionId   string                                                   `json:"policy_definition_id"`
	SubscriptionId       string                                                   `json:"subscription_id"`
	Identity             []subscriptionpolicyassignment.IdentityState             `json:"identity"`
	NonComplianceMessage []subscriptionpolicyassignment.NonComplianceMessageState `json:"non_compliance_message"`
	Overrides            []subscriptionpolicyassignment.OverridesState            `json:"overrides"`
	ResourceSelectors    []subscriptionpolicyassignment.ResourceSelectorsState    `json:"resource_selectors"`
	Timeouts             *subscriptionpolicyassignment.TimeoutsState              `json:"timeouts"`
}
