// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	natgatewaypublicipassociation "github.com/golingon/terraproviders/azurerm/3.66.0/natgatewaypublicipassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNatGatewayPublicIpAssociation creates a new instance of [NatGatewayPublicIpAssociation].
func NewNatGatewayPublicIpAssociation(name string, args NatGatewayPublicIpAssociationArgs) *NatGatewayPublicIpAssociation {
	return &NatGatewayPublicIpAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NatGatewayPublicIpAssociation)(nil)

// NatGatewayPublicIpAssociation represents the Terraform resource azurerm_nat_gateway_public_ip_association.
type NatGatewayPublicIpAssociation struct {
	Name      string
	Args      NatGatewayPublicIpAssociationArgs
	state     *natGatewayPublicIpAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) Type() string {
	return "azurerm_nat_gateway_public_ip_association"
}

// LocalName returns the local name for [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) LocalName() string {
	return ngpia.Name
}

// Configuration returns the configuration (args) for [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) Configuration() interface{} {
	return ngpia.Args
}

// DependOn is used for other resources to depend on [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ngpia)
}

// Dependencies returns the list of resources [NatGatewayPublicIpAssociation] depends_on.
func (ngpia *NatGatewayPublicIpAssociation) Dependencies() terra.Dependencies {
	return ngpia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) LifecycleManagement() *terra.Lifecycle {
	return ngpia.Lifecycle
}

// Attributes returns the attributes for [NatGatewayPublicIpAssociation].
func (ngpia *NatGatewayPublicIpAssociation) Attributes() natGatewayPublicIpAssociationAttributes {
	return natGatewayPublicIpAssociationAttributes{ref: terra.ReferenceResource(ngpia)}
}

// ImportState imports the given attribute values into [NatGatewayPublicIpAssociation]'s state.
func (ngpia *NatGatewayPublicIpAssociation) ImportState(av io.Reader) error {
	ngpia.state = &natGatewayPublicIpAssociationState{}
	if err := json.NewDecoder(av).Decode(ngpia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ngpia.Type(), ngpia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NatGatewayPublicIpAssociation] has state.
func (ngpia *NatGatewayPublicIpAssociation) State() (*natGatewayPublicIpAssociationState, bool) {
	return ngpia.state, ngpia.state != nil
}

// StateMust returns the state for [NatGatewayPublicIpAssociation]. Panics if the state is nil.
func (ngpia *NatGatewayPublicIpAssociation) StateMust() *natGatewayPublicIpAssociationState {
	if ngpia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ngpia.Type(), ngpia.LocalName()))
	}
	return ngpia.state
}

// NatGatewayPublicIpAssociationArgs contains the configurations for azurerm_nat_gateway_public_ip_association.
type NatGatewayPublicIpAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NatGatewayId: string, required
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr" validate:"required"`
	// PublicIpAddressId: string, required
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *natgatewaypublicipassociation.Timeouts `hcl:"timeouts,block"`
}
type natGatewayPublicIpAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_nat_gateway_public_ip_association.
func (ngpia natGatewayPublicIpAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngpia.ref.Append("id"))
}

// NatGatewayId returns a reference to field nat_gateway_id of azurerm_nat_gateway_public_ip_association.
func (ngpia natGatewayPublicIpAssociationAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(ngpia.ref.Append("nat_gateway_id"))
}

// PublicIpAddressId returns a reference to field public_ip_address_id of azurerm_nat_gateway_public_ip_association.
func (ngpia natGatewayPublicIpAssociationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(ngpia.ref.Append("public_ip_address_id"))
}

func (ngpia natGatewayPublicIpAssociationAttributes) Timeouts() natgatewaypublicipassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[natgatewaypublicipassociation.TimeoutsAttributes](ngpia.ref.Append("timeouts"))
}

type natGatewayPublicIpAssociationState struct {
	Id                string                                       `json:"id"`
	NatGatewayId      string                                       `json:"nat_gateway_id"`
	PublicIpAddressId string                                       `json:"public_ip_address_id"`
	Timeouts          *natgatewaypublicipassociation.TimeoutsState `json:"timeouts"`
}
