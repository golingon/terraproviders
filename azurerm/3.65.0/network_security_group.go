// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networksecuritygroup "github.com/golingon/terraproviders/azurerm/3.65.0/networksecuritygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityGroup creates a new instance of [NetworkSecurityGroup].
func NewNetworkSecurityGroup(name string, args NetworkSecurityGroupArgs) *NetworkSecurityGroup {
	return &NetworkSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityGroup)(nil)

// NetworkSecurityGroup represents the Terraform resource azurerm_network_security_group.
type NetworkSecurityGroup struct {
	Name      string
	Args      NetworkSecurityGroupArgs
	state     *networkSecurityGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) Type() string {
	return "azurerm_network_security_group"
}

// LocalName returns the local name for [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) LocalName() string {
	return nsg.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) Configuration() interface{} {
	return nsg.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(nsg)
}

// Dependencies returns the list of resources [NetworkSecurityGroup] depends_on.
func (nsg *NetworkSecurityGroup) Dependencies() terra.Dependencies {
	return nsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) LifecycleManagement() *terra.Lifecycle {
	return nsg.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityGroup].
func (nsg *NetworkSecurityGroup) Attributes() networkSecurityGroupAttributes {
	return networkSecurityGroupAttributes{ref: terra.ReferenceResource(nsg)}
}

// ImportState imports the given attribute values into [NetworkSecurityGroup]'s state.
func (nsg *NetworkSecurityGroup) ImportState(av io.Reader) error {
	nsg.state = &networkSecurityGroupState{}
	if err := json.NewDecoder(av).Decode(nsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsg.Type(), nsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityGroup] has state.
func (nsg *NetworkSecurityGroup) State() (*networkSecurityGroupState, bool) {
	return nsg.state, nsg.state != nil
}

// StateMust returns the state for [NetworkSecurityGroup]. Panics if the state is nil.
func (nsg *NetworkSecurityGroup) StateMust() *networkSecurityGroupState {
	if nsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsg.Type(), nsg.LocalName()))
	}
	return nsg.state
}

// NetworkSecurityGroupArgs contains the configurations for azurerm_network_security_group.
type NetworkSecurityGroupArgs struct {
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
	// SecurityRule: min=0
	SecurityRule []networksecuritygroup.SecurityRule `hcl:"security_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *networksecuritygroup.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_security_group.
func (nsg networkSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_security_group.
func (nsg networkSecurityGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_security_group.
func (nsg networkSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_security_group.
func (nsg networkSecurityGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_security_group.
func (nsg networkSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsg.ref.Append("tags"))
}

func (nsg networkSecurityGroupAttributes) SecurityRule() terra.SetValue[networksecuritygroup.SecurityRuleAttributes] {
	return terra.ReferenceAsSet[networksecuritygroup.SecurityRuleAttributes](nsg.ref.Append("security_rule"))
}

func (nsg networkSecurityGroupAttributes) Timeouts() networksecuritygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecuritygroup.TimeoutsAttributes](nsg.ref.Append("timeouts"))
}

type networkSecurityGroupState struct {
	Id                string                                   `json:"id"`
	Location          string                                   `json:"location"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	Tags              map[string]string                        `json:"tags"`
	SecurityRule      []networksecuritygroup.SecurityRuleState `json:"security_rule"`
	Timeouts          *networksecuritygroup.TimeoutsState      `json:"timeouts"`
}
