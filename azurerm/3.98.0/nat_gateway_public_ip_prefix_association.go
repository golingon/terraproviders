// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	natgatewaypublicipprefixassociation "github.com/golingon/terraproviders/azurerm/3.98.0/natgatewaypublicipprefixassociation"
	"io"
)

// NewNatGatewayPublicIpPrefixAssociation creates a new instance of [NatGatewayPublicIpPrefixAssociation].
func NewNatGatewayPublicIpPrefixAssociation(name string, args NatGatewayPublicIpPrefixAssociationArgs) *NatGatewayPublicIpPrefixAssociation {
	return &NatGatewayPublicIpPrefixAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NatGatewayPublicIpPrefixAssociation)(nil)

// NatGatewayPublicIpPrefixAssociation represents the Terraform resource azurerm_nat_gateway_public_ip_prefix_association.
type NatGatewayPublicIpPrefixAssociation struct {
	Name      string
	Args      NatGatewayPublicIpPrefixAssociationArgs
	state     *natGatewayPublicIpPrefixAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) Type() string {
	return "azurerm_nat_gateway_public_ip_prefix_association"
}

// LocalName returns the local name for [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) LocalName() string {
	return ngpipa.Name
}

// Configuration returns the configuration (args) for [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) Configuration() interface{} {
	return ngpipa.Args
}

// DependOn is used for other resources to depend on [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ngpipa)
}

// Dependencies returns the list of resources [NatGatewayPublicIpPrefixAssociation] depends_on.
func (ngpipa *NatGatewayPublicIpPrefixAssociation) Dependencies() terra.Dependencies {
	return ngpipa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) LifecycleManagement() *terra.Lifecycle {
	return ngpipa.Lifecycle
}

// Attributes returns the attributes for [NatGatewayPublicIpPrefixAssociation].
func (ngpipa *NatGatewayPublicIpPrefixAssociation) Attributes() natGatewayPublicIpPrefixAssociationAttributes {
	return natGatewayPublicIpPrefixAssociationAttributes{ref: terra.ReferenceResource(ngpipa)}
}

// ImportState imports the given attribute values into [NatGatewayPublicIpPrefixAssociation]'s state.
func (ngpipa *NatGatewayPublicIpPrefixAssociation) ImportState(av io.Reader) error {
	ngpipa.state = &natGatewayPublicIpPrefixAssociationState{}
	if err := json.NewDecoder(av).Decode(ngpipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ngpipa.Type(), ngpipa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NatGatewayPublicIpPrefixAssociation] has state.
func (ngpipa *NatGatewayPublicIpPrefixAssociation) State() (*natGatewayPublicIpPrefixAssociationState, bool) {
	return ngpipa.state, ngpipa.state != nil
}

// StateMust returns the state for [NatGatewayPublicIpPrefixAssociation]. Panics if the state is nil.
func (ngpipa *NatGatewayPublicIpPrefixAssociation) StateMust() *natGatewayPublicIpPrefixAssociationState {
	if ngpipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ngpipa.Type(), ngpipa.LocalName()))
	}
	return ngpipa.state
}

// NatGatewayPublicIpPrefixAssociationArgs contains the configurations for azurerm_nat_gateway_public_ip_prefix_association.
type NatGatewayPublicIpPrefixAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NatGatewayId: string, required
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr" validate:"required"`
	// PublicIpPrefixId: string, required
	PublicIpPrefixId terra.StringValue `hcl:"public_ip_prefix_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *natgatewaypublicipprefixassociation.Timeouts `hcl:"timeouts,block"`
}
type natGatewayPublicIpPrefixAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_nat_gateway_public_ip_prefix_association.
func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngpipa.ref.Append("id"))
}

// NatGatewayId returns a reference to field nat_gateway_id of azurerm_nat_gateway_public_ip_prefix_association.
func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(ngpipa.ref.Append("nat_gateway_id"))
}

// PublicIpPrefixId returns a reference to field public_ip_prefix_id of azurerm_nat_gateway_public_ip_prefix_association.
func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) PublicIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(ngpipa.ref.Append("public_ip_prefix_id"))
}

func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) Timeouts() natgatewaypublicipprefixassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[natgatewaypublicipprefixassociation.TimeoutsAttributes](ngpipa.ref.Append("timeouts"))
}

type natGatewayPublicIpPrefixAssociationState struct {
	Id               string                                             `json:"id"`
	NatGatewayId     string                                             `json:"nat_gateway_id"`
	PublicIpPrefixId string                                             `json:"public_ip_prefix_id"`
	Timeouts         *natgatewaypublicipprefixassociation.TimeoutsState `json:"timeouts"`
}
