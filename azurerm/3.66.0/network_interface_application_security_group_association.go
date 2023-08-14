// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkinterfaceapplicationsecuritygroupassociation "github.com/golingon/terraproviders/azurerm/3.66.0/networkinterfaceapplicationsecuritygroupassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceApplicationSecurityGroupAssociation creates a new instance of [NetworkInterfaceApplicationSecurityGroupAssociation].
func NewNetworkInterfaceApplicationSecurityGroupAssociation(name string, args NetworkInterfaceApplicationSecurityGroupAssociationArgs) *NetworkInterfaceApplicationSecurityGroupAssociation {
	return &NetworkInterfaceApplicationSecurityGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceApplicationSecurityGroupAssociation)(nil)

// NetworkInterfaceApplicationSecurityGroupAssociation represents the Terraform resource azurerm_network_interface_application_security_group_association.
type NetworkInterfaceApplicationSecurityGroupAssociation struct {
	Name      string
	Args      NetworkInterfaceApplicationSecurityGroupAssociationArgs
	state     *networkInterfaceApplicationSecurityGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) Type() string {
	return "azurerm_network_interface_application_security_group_association"
}

// LocalName returns the local name for [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) LocalName() string {
	return niasga.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) Configuration() interface{} {
	return niasga.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(niasga)
}

// Dependencies returns the list of resources [NetworkInterfaceApplicationSecurityGroupAssociation] depends_on.
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) Dependencies() terra.Dependencies {
	return niasga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return niasga.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceApplicationSecurityGroupAssociation].
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) Attributes() networkInterfaceApplicationSecurityGroupAssociationAttributes {
	return networkInterfaceApplicationSecurityGroupAssociationAttributes{ref: terra.ReferenceResource(niasga)}
}

// ImportState imports the given attribute values into [NetworkInterfaceApplicationSecurityGroupAssociation]'s state.
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) ImportState(av io.Reader) error {
	niasga.state = &networkInterfaceApplicationSecurityGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(niasga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", niasga.Type(), niasga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceApplicationSecurityGroupAssociation] has state.
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) State() (*networkInterfaceApplicationSecurityGroupAssociationState, bool) {
	return niasga.state, niasga.state != nil
}

// StateMust returns the state for [NetworkInterfaceApplicationSecurityGroupAssociation]. Panics if the state is nil.
func (niasga *NetworkInterfaceApplicationSecurityGroupAssociation) StateMust() *networkInterfaceApplicationSecurityGroupAssociationState {
	if niasga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", niasga.Type(), niasga.LocalName()))
	}
	return niasga.state
}

// NetworkInterfaceApplicationSecurityGroupAssociationArgs contains the configurations for azurerm_network_interface_application_security_group_association.
type NetworkInterfaceApplicationSecurityGroupAssociationArgs struct {
	// ApplicationSecurityGroupId: string, required
	ApplicationSecurityGroupId terra.StringValue `hcl:"application_security_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkinterfaceapplicationsecuritygroupassociation.Timeouts `hcl:"timeouts,block"`
}
type networkInterfaceApplicationSecurityGroupAssociationAttributes struct {
	ref terra.Reference
}

// ApplicationSecurityGroupId returns a reference to field application_security_group_id of azurerm_network_interface_application_security_group_association.
func (niasga networkInterfaceApplicationSecurityGroupAssociationAttributes) ApplicationSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(niasga.ref.Append("application_security_group_id"))
}

// Id returns a reference to field id of azurerm_network_interface_application_security_group_association.
func (niasga networkInterfaceApplicationSecurityGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(niasga.ref.Append("id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of azurerm_network_interface_application_security_group_association.
func (niasga networkInterfaceApplicationSecurityGroupAssociationAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(niasga.ref.Append("network_interface_id"))
}

func (niasga networkInterfaceApplicationSecurityGroupAssociationAttributes) Timeouts() networkinterfaceapplicationsecuritygroupassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkinterfaceapplicationsecuritygroupassociation.TimeoutsAttributes](niasga.ref.Append("timeouts"))
}

type networkInterfaceApplicationSecurityGroupAssociationState struct {
	ApplicationSecurityGroupId string                                                             `json:"application_security_group_id"`
	Id                         string                                                             `json:"id"`
	NetworkInterfaceId         string                                                             `json:"network_interface_id"`
	Timeouts                   *networkinterfaceapplicationsecuritygroupassociation.TimeoutsState `json:"timeouts"`
}