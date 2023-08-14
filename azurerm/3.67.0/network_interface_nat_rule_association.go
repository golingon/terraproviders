// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkinterfacenatruleassociation "github.com/golingon/terraproviders/azurerm/3.67.0/networkinterfacenatruleassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceNatRuleAssociation creates a new instance of [NetworkInterfaceNatRuleAssociation].
func NewNetworkInterfaceNatRuleAssociation(name string, args NetworkInterfaceNatRuleAssociationArgs) *NetworkInterfaceNatRuleAssociation {
	return &NetworkInterfaceNatRuleAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceNatRuleAssociation)(nil)

// NetworkInterfaceNatRuleAssociation represents the Terraform resource azurerm_network_interface_nat_rule_association.
type NetworkInterfaceNatRuleAssociation struct {
	Name      string
	Args      NetworkInterfaceNatRuleAssociationArgs
	state     *networkInterfaceNatRuleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) Type() string {
	return "azurerm_network_interface_nat_rule_association"
}

// LocalName returns the local name for [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) LocalName() string {
	return ninra.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) Configuration() interface{} {
	return ninra.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ninra)
}

// Dependencies returns the list of resources [NetworkInterfaceNatRuleAssociation] depends_on.
func (ninra *NetworkInterfaceNatRuleAssociation) Dependencies() terra.Dependencies {
	return ninra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) LifecycleManagement() *terra.Lifecycle {
	return ninra.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceNatRuleAssociation].
func (ninra *NetworkInterfaceNatRuleAssociation) Attributes() networkInterfaceNatRuleAssociationAttributes {
	return networkInterfaceNatRuleAssociationAttributes{ref: terra.ReferenceResource(ninra)}
}

// ImportState imports the given attribute values into [NetworkInterfaceNatRuleAssociation]'s state.
func (ninra *NetworkInterfaceNatRuleAssociation) ImportState(av io.Reader) error {
	ninra.state = &networkInterfaceNatRuleAssociationState{}
	if err := json.NewDecoder(av).Decode(ninra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ninra.Type(), ninra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceNatRuleAssociation] has state.
func (ninra *NetworkInterfaceNatRuleAssociation) State() (*networkInterfaceNatRuleAssociationState, bool) {
	return ninra.state, ninra.state != nil
}

// StateMust returns the state for [NetworkInterfaceNatRuleAssociation]. Panics if the state is nil.
func (ninra *NetworkInterfaceNatRuleAssociation) StateMust() *networkInterfaceNatRuleAssociationState {
	if ninra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ninra.Type(), ninra.LocalName()))
	}
	return ninra.state
}

// NetworkInterfaceNatRuleAssociationArgs contains the configurations for azurerm_network_interface_nat_rule_association.
type NetworkInterfaceNatRuleAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConfigurationName: string, required
	IpConfigurationName terra.StringValue `hcl:"ip_configuration_name,attr" validate:"required"`
	// NatRuleId: string, required
	NatRuleId terra.StringValue `hcl:"nat_rule_id,attr" validate:"required"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkinterfacenatruleassociation.Timeouts `hcl:"timeouts,block"`
}
type networkInterfaceNatRuleAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_interface_nat_rule_association.
func (ninra networkInterfaceNatRuleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ninra.ref.Append("id"))
}

// IpConfigurationName returns a reference to field ip_configuration_name of azurerm_network_interface_nat_rule_association.
func (ninra networkInterfaceNatRuleAssociationAttributes) IpConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(ninra.ref.Append("ip_configuration_name"))
}

// NatRuleId returns a reference to field nat_rule_id of azurerm_network_interface_nat_rule_association.
func (ninra networkInterfaceNatRuleAssociationAttributes) NatRuleId() terra.StringValue {
	return terra.ReferenceAsString(ninra.ref.Append("nat_rule_id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_nat_rule_association.
func (ninra networkInterfaceNatRuleAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ninra.ref.Append("network_interface_id"))
}

func (ninra networkInterfaceNatRuleAssociationAttributes) Timeouts() networkinterfacenatruleassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkinterfacenatruleassociation.TimeoutsAttributes](ninra.ref.Append("timeouts"))
}

type networkInterfaceNatRuleAssociationState struct {
	Id                  string                                            `json:"id"`
	IpConfigurationName string                                            `json:"ip_configuration_name"`
	NatRuleId           string                                            `json:"nat_rule_id"`
	NetworkInterfaceId  string                                            `json:"network_interface_id"`
	Timeouts            *networkinterfacenatruleassociation.TimeoutsState `json:"timeouts"`
}