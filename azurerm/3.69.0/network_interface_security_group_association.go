// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkinterfacesecuritygroupassociation "github.com/golingon/terraproviders/azurerm/3.69.0/networkinterfacesecuritygroupassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceSecurityGroupAssociation creates a new instance of [NetworkInterfaceSecurityGroupAssociation].
func NewNetworkInterfaceSecurityGroupAssociation(name string, args NetworkInterfaceSecurityGroupAssociationArgs) *NetworkInterfaceSecurityGroupAssociation {
	return &NetworkInterfaceSecurityGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceSecurityGroupAssociation)(nil)

// NetworkInterfaceSecurityGroupAssociation represents the Terraform resource azurerm_network_interface_security_group_association.
type NetworkInterfaceSecurityGroupAssociation struct {
	Name      string
	Args      NetworkInterfaceSecurityGroupAssociationArgs
	state     *networkInterfaceSecurityGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) Type() string {
	return "azurerm_network_interface_security_group_association"
}

// LocalName returns the local name for [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) LocalName() string {
	return nisga.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) Configuration() interface{} {
	return nisga.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(nisga)
}

// Dependencies returns the list of resources [NetworkInterfaceSecurityGroupAssociation] depends_on.
func (nisga *NetworkInterfaceSecurityGroupAssociation) Dependencies() terra.Dependencies {
	return nisga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return nisga.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceSecurityGroupAssociation].
func (nisga *NetworkInterfaceSecurityGroupAssociation) Attributes() networkInterfaceSecurityGroupAssociationAttributes {
	return networkInterfaceSecurityGroupAssociationAttributes{ref: terra.ReferenceResource(nisga)}
}

// ImportState imports the given attribute values into [NetworkInterfaceSecurityGroupAssociation]'s state.
func (nisga *NetworkInterfaceSecurityGroupAssociation) ImportState(av io.Reader) error {
	nisga.state = &networkInterfaceSecurityGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(nisga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nisga.Type(), nisga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceSecurityGroupAssociation] has state.
func (nisga *NetworkInterfaceSecurityGroupAssociation) State() (*networkInterfaceSecurityGroupAssociationState, bool) {
	return nisga.state, nisga.state != nil
}

// StateMust returns the state for [NetworkInterfaceSecurityGroupAssociation]. Panics if the state is nil.
func (nisga *NetworkInterfaceSecurityGroupAssociation) StateMust() *networkInterfaceSecurityGroupAssociationState {
	if nisga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nisga.Type(), nisga.LocalName()))
	}
	return nisga.state
}

// NetworkInterfaceSecurityGroupAssociationArgs contains the configurations for azurerm_network_interface_security_group_association.
type NetworkInterfaceSecurityGroupAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// NetworkSecurityGroupId: string, required
	NetworkSecurityGroupId terra.StringValue `hcl:"network_security_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkinterfacesecuritygroupassociation.Timeouts `hcl:"timeouts,block"`
}
type networkInterfaceSecurityGroupAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_interface_security_group_association.
func (nisga networkInterfaceSecurityGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nisga.ref.Append("id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_security_group_association.
func (nisga networkInterfaceSecurityGroupAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(nisga.ref.Append("network_interface_id"))
}

// NetworkSecurityGroupId returns a reference to field network_security_group_id of azurerm_network_interface_security_group_association.
func (nisga networkInterfaceSecurityGroupAssociationAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(nisga.ref.Append("network_security_group_id"))
}

func (nisga networkInterfaceSecurityGroupAssociationAttributes) Timeouts() networkinterfacesecuritygroupassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkinterfacesecuritygroupassociation.TimeoutsAttributes](nisga.ref.Append("timeouts"))
}

type networkInterfaceSecurityGroupAssociationState struct {
	Id                     string                                                  `json:"id"`
	NetworkInterfaceId     string                                                  `json:"network_interface_id"`
	NetworkSecurityGroupId string                                                  `json:"network_security_group_id"`
	Timeouts               *networkinterfacesecuritygroupassociation.TimeoutsState `json:"timeouts"`
}
