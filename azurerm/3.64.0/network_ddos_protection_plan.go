// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkddosprotectionplan "github.com/golingon/terraproviders/azurerm/3.64.0/networkddosprotectionplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkDdosProtectionPlan creates a new instance of [NetworkDdosProtectionPlan].
func NewNetworkDdosProtectionPlan(name string, args NetworkDdosProtectionPlanArgs) *NetworkDdosProtectionPlan {
	return &NetworkDdosProtectionPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkDdosProtectionPlan)(nil)

// NetworkDdosProtectionPlan represents the Terraform resource azurerm_network_ddos_protection_plan.
type NetworkDdosProtectionPlan struct {
	Name      string
	Args      NetworkDdosProtectionPlanArgs
	state     *networkDdosProtectionPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) Type() string {
	return "azurerm_network_ddos_protection_plan"
}

// LocalName returns the local name for [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) LocalName() string {
	return ndpp.Name
}

// Configuration returns the configuration (args) for [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) Configuration() interface{} {
	return ndpp.Args
}

// DependOn is used for other resources to depend on [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(ndpp)
}

// Dependencies returns the list of resources [NetworkDdosProtectionPlan] depends_on.
func (ndpp *NetworkDdosProtectionPlan) Dependencies() terra.Dependencies {
	return ndpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) LifecycleManagement() *terra.Lifecycle {
	return ndpp.Lifecycle
}

// Attributes returns the attributes for [NetworkDdosProtectionPlan].
func (ndpp *NetworkDdosProtectionPlan) Attributes() networkDdosProtectionPlanAttributes {
	return networkDdosProtectionPlanAttributes{ref: terra.ReferenceResource(ndpp)}
}

// ImportState imports the given attribute values into [NetworkDdosProtectionPlan]'s state.
func (ndpp *NetworkDdosProtectionPlan) ImportState(av io.Reader) error {
	ndpp.state = &networkDdosProtectionPlanState{}
	if err := json.NewDecoder(av).Decode(ndpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ndpp.Type(), ndpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkDdosProtectionPlan] has state.
func (ndpp *NetworkDdosProtectionPlan) State() (*networkDdosProtectionPlanState, bool) {
	return ndpp.state, ndpp.state != nil
}

// StateMust returns the state for [NetworkDdosProtectionPlan]. Panics if the state is nil.
func (ndpp *NetworkDdosProtectionPlan) StateMust() *networkDdosProtectionPlanState {
	if ndpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ndpp.Type(), ndpp.LocalName()))
	}
	return ndpp.state
}

// NetworkDdosProtectionPlanArgs contains the configurations for azurerm_network_ddos_protection_plan.
type NetworkDdosProtectionPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *networkddosprotectionplan.Timeouts `hcl:"timeouts,block"`
}
type networkDdosProtectionPlanAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ndpp.ref.Append("tags"))
}

// VirtualNetworkIds returns a reference to field virtual_network_ids of azurerm_network_ddos_protection_plan.
func (ndpp networkDdosProtectionPlanAttributes) VirtualNetworkIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ndpp.ref.Append("virtual_network_ids"))
}

func (ndpp networkDdosProtectionPlanAttributes) Timeouts() networkddosprotectionplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkddosprotectionplan.TimeoutsAttributes](ndpp.ref.Append("timeouts"))
}

type networkDdosProtectionPlanState struct {
	Id                string                                   `json:"id"`
	Location          string                                   `json:"location"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	Tags              map[string]string                        `json:"tags"`
	VirtualNetworkIds []string                                 `json:"virtual_network_ids"`
	Timeouts          *networkddosprotectionplan.TimeoutsState `json:"timeouts"`
}
