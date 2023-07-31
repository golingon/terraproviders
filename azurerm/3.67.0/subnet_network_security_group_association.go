// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subnetnetworksecuritygroupassociation "github.com/golingon/terraproviders/azurerm/3.67.0/subnetnetworksecuritygroupassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnetNetworkSecurityGroupAssociation creates a new instance of [SubnetNetworkSecurityGroupAssociation].
func NewSubnetNetworkSecurityGroupAssociation(name string, args SubnetNetworkSecurityGroupAssociationArgs) *SubnetNetworkSecurityGroupAssociation {
	return &SubnetNetworkSecurityGroupAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubnetNetworkSecurityGroupAssociation)(nil)

// SubnetNetworkSecurityGroupAssociation represents the Terraform resource azurerm_subnet_network_security_group_association.
type SubnetNetworkSecurityGroupAssociation struct {
	Name      string
	Args      SubnetNetworkSecurityGroupAssociationArgs
	state     *subnetNetworkSecurityGroupAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) Type() string {
	return "azurerm_subnet_network_security_group_association"
}

// LocalName returns the local name for [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) LocalName() string {
	return snsga.Name
}

// Configuration returns the configuration (args) for [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) Configuration() interface{} {
	return snsga.Args
}

// DependOn is used for other resources to depend on [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(snsga)
}

// Dependencies returns the list of resources [SubnetNetworkSecurityGroupAssociation] depends_on.
func (snsga *SubnetNetworkSecurityGroupAssociation) Dependencies() terra.Dependencies {
	return snsga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) LifecycleManagement() *terra.Lifecycle {
	return snsga.Lifecycle
}

// Attributes returns the attributes for [SubnetNetworkSecurityGroupAssociation].
func (snsga *SubnetNetworkSecurityGroupAssociation) Attributes() subnetNetworkSecurityGroupAssociationAttributes {
	return subnetNetworkSecurityGroupAssociationAttributes{ref: terra.ReferenceResource(snsga)}
}

// ImportState imports the given attribute values into [SubnetNetworkSecurityGroupAssociation]'s state.
func (snsga *SubnetNetworkSecurityGroupAssociation) ImportState(av io.Reader) error {
	snsga.state = &subnetNetworkSecurityGroupAssociationState{}
	if err := json.NewDecoder(av).Decode(snsga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snsga.Type(), snsga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubnetNetworkSecurityGroupAssociation] has state.
func (snsga *SubnetNetworkSecurityGroupAssociation) State() (*subnetNetworkSecurityGroupAssociationState, bool) {
	return snsga.state, snsga.state != nil
}

// StateMust returns the state for [SubnetNetworkSecurityGroupAssociation]. Panics if the state is nil.
func (snsga *SubnetNetworkSecurityGroupAssociation) StateMust() *subnetNetworkSecurityGroupAssociationState {
	if snsga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snsga.Type(), snsga.LocalName()))
	}
	return snsga.state
}

// SubnetNetworkSecurityGroupAssociationArgs contains the configurations for azurerm_subnet_network_security_group_association.
type SubnetNetworkSecurityGroupAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkSecurityGroupId: string, required
	NetworkSecurityGroupId terra.StringValue `hcl:"network_security_group_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subnetnetworksecuritygroupassociation.Timeouts `hcl:"timeouts,block"`
}
type subnetNetworkSecurityGroupAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_subnet_network_security_group_association.
func (snsga subnetNetworkSecurityGroupAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snsga.ref.Append("id"))
}

// NetworkSecurityGroupId returns a reference to field network_security_group_id of azurerm_subnet_network_security_group_association.
func (snsga subnetNetworkSecurityGroupAssociationAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(snsga.ref.Append("network_security_group_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_subnet_network_security_group_association.
func (snsga subnetNetworkSecurityGroupAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(snsga.ref.Append("subnet_id"))
}

func (snsga subnetNetworkSecurityGroupAssociationAttributes) Timeouts() subnetnetworksecuritygroupassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnetnetworksecuritygroupassociation.TimeoutsAttributes](snsga.ref.Append("timeouts"))
}

type subnetNetworkSecurityGroupAssociationState struct {
	Id                     string                                               `json:"id"`
	NetworkSecurityGroupId string                                               `json:"network_security_group_id"`
	SubnetId               string                                               `json:"subnet_id"`
	Timeouts               *subnetnetworksecuritygroupassociation.TimeoutsState `json:"timeouts"`
}
