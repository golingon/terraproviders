// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	natgatewaypublicipprefixassociation "github.com/golingon/terraproviders/azurerm/3.49.0/natgatewaypublicipprefixassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNatGatewayPublicIpPrefixAssociation(name string, args NatGatewayPublicIpPrefixAssociationArgs) *NatGatewayPublicIpPrefixAssociation {
	return &NatGatewayPublicIpPrefixAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NatGatewayPublicIpPrefixAssociation)(nil)

type NatGatewayPublicIpPrefixAssociation struct {
	Name  string
	Args  NatGatewayPublicIpPrefixAssociationArgs
	state *natGatewayPublicIpPrefixAssociationState
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) Type() string {
	return "azurerm_nat_gateway_public_ip_prefix_association"
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) LocalName() string {
	return ngpipa.Name
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) Configuration() interface{} {
	return ngpipa.Args
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) Attributes() natGatewayPublicIpPrefixAssociationAttributes {
	return natGatewayPublicIpPrefixAssociationAttributes{ref: terra.ReferenceResource(ngpipa)}
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) ImportState(av io.Reader) error {
	ngpipa.state = &natGatewayPublicIpPrefixAssociationState{}
	if err := json.NewDecoder(av).Decode(ngpipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ngpipa.Type(), ngpipa.LocalName(), err)
	}
	return nil
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) State() (*natGatewayPublicIpPrefixAssociationState, bool) {
	return ngpipa.state, ngpipa.state != nil
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) StateMust() *natGatewayPublicIpPrefixAssociationState {
	if ngpipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ngpipa.Type(), ngpipa.LocalName()))
	}
	return ngpipa.state
}

func (ngpipa *NatGatewayPublicIpPrefixAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ngpipa)
}

type NatGatewayPublicIpPrefixAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NatGatewayId: string, required
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr" validate:"required"`
	// PublicIpPrefixId: string, required
	PublicIpPrefixId terra.StringValue `hcl:"public_ip_prefix_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *natgatewaypublicipprefixassociation.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NatGatewayPublicIpPrefixAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type natGatewayPublicIpPrefixAssociationAttributes struct {
	ref terra.Reference
}

func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ngpipa.ref.Append("id"))
}

func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceString(ngpipa.ref.Append("nat_gateway_id"))
}

func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) PublicIpPrefixId() terra.StringValue {
	return terra.ReferenceString(ngpipa.ref.Append("public_ip_prefix_id"))
}

func (ngpipa natGatewayPublicIpPrefixAssociationAttributes) Timeouts() natgatewaypublicipprefixassociation.TimeoutsAttributes {
	return terra.ReferenceSingle[natgatewaypublicipprefixassociation.TimeoutsAttributes](ngpipa.ref.Append("timeouts"))
}

type natGatewayPublicIpPrefixAssociationState struct {
	Id               string                                             `json:"id"`
	NatGatewayId     string                                             `json:"nat_gateway_id"`
	PublicIpPrefixId string                                             `json:"public_ip_prefix_id"`
	Timeouts         *natgatewaypublicipprefixassociation.TimeoutsState `json:"timeouts"`
}
