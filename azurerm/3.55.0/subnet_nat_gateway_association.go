// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subnetnatgatewayassociation "github.com/golingon/terraproviders/azurerm/3.55.0/subnetnatgatewayassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnetNatGatewayAssociation creates a new instance of [SubnetNatGatewayAssociation].
func NewSubnetNatGatewayAssociation(name string, args SubnetNatGatewayAssociationArgs) *SubnetNatGatewayAssociation {
	return &SubnetNatGatewayAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubnetNatGatewayAssociation)(nil)

// SubnetNatGatewayAssociation represents the Terraform resource azurerm_subnet_nat_gateway_association.
type SubnetNatGatewayAssociation struct {
	Name      string
	Args      SubnetNatGatewayAssociationArgs
	state     *subnetNatGatewayAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) Type() string {
	return "azurerm_subnet_nat_gateway_association"
}

// LocalName returns the local name for [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) LocalName() string {
	return snga.Name
}

// Configuration returns the configuration (args) for [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) Configuration() interface{} {
	return snga.Args
}

// DependOn is used for other resources to depend on [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(snga)
}

// Dependencies returns the list of resources [SubnetNatGatewayAssociation] depends_on.
func (snga *SubnetNatGatewayAssociation) Dependencies() terra.Dependencies {
	return snga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) LifecycleManagement() *terra.Lifecycle {
	return snga.Lifecycle
}

// Attributes returns the attributes for [SubnetNatGatewayAssociation].
func (snga *SubnetNatGatewayAssociation) Attributes() subnetNatGatewayAssociationAttributes {
	return subnetNatGatewayAssociationAttributes{ref: terra.ReferenceResource(snga)}
}

// ImportState imports the given attribute values into [SubnetNatGatewayAssociation]'s state.
func (snga *SubnetNatGatewayAssociation) ImportState(av io.Reader) error {
	snga.state = &subnetNatGatewayAssociationState{}
	if err := json.NewDecoder(av).Decode(snga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snga.Type(), snga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubnetNatGatewayAssociation] has state.
func (snga *SubnetNatGatewayAssociation) State() (*subnetNatGatewayAssociationState, bool) {
	return snga.state, snga.state != nil
}

// StateMust returns the state for [SubnetNatGatewayAssociation]. Panics if the state is nil.
func (snga *SubnetNatGatewayAssociation) StateMust() *subnetNatGatewayAssociationState {
	if snga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snga.Type(), snga.LocalName()))
	}
	return snga.state
}

// SubnetNatGatewayAssociationArgs contains the configurations for azurerm_subnet_nat_gateway_association.
type SubnetNatGatewayAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NatGatewayId: string, required
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subnetnatgatewayassociation.Timeouts `hcl:"timeouts,block"`
}
type subnetNatGatewayAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_subnet_nat_gateway_association.
func (snga subnetNatGatewayAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snga.ref.Append("id"))
}

// NatGatewayId returns a reference to field nat_gateway_id of azurerm_subnet_nat_gateway_association.
func (snga subnetNatGatewayAssociationAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(snga.ref.Append("nat_gateway_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_subnet_nat_gateway_association.
func (snga subnetNatGatewayAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(snga.ref.Append("subnet_id"))
}

func (snga subnetNatGatewayAssociationAttributes) Timeouts() subnetnatgatewayassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnetnatgatewayassociation.TimeoutsAttributes](snga.ref.Append("timeouts"))
}

type subnetNatGatewayAssociationState struct {
	Id           string                                     `json:"id"`
	NatGatewayId string                                     `json:"nat_gateway_id"`
	SubnetId     string                                     `json:"subnet_id"`
	Timeouts     *subnetnatgatewayassociation.TimeoutsState `json:"timeouts"`
}
