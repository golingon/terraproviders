// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagercustomergatewayassociation "github.com/golingon/terraproviders/aws/4.60.0/networkmanagercustomergatewayassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkmanagerCustomerGatewayAssociation(name string, args NetworkmanagerCustomerGatewayAssociationArgs) *NetworkmanagerCustomerGatewayAssociation {
	return &NetworkmanagerCustomerGatewayAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerCustomerGatewayAssociation)(nil)

type NetworkmanagerCustomerGatewayAssociation struct {
	Name  string
	Args  NetworkmanagerCustomerGatewayAssociationArgs
	state *networkmanagerCustomerGatewayAssociationState
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) Type() string {
	return "aws_networkmanager_customer_gateway_association"
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) LocalName() string {
	return ncga.Name
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) Configuration() interface{} {
	return ncga.Args
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) Attributes() networkmanagerCustomerGatewayAssociationAttributes {
	return networkmanagerCustomerGatewayAssociationAttributes{ref: terra.ReferenceResource(ncga)}
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) ImportState(av io.Reader) error {
	ncga.state = &networkmanagerCustomerGatewayAssociationState{}
	if err := json.NewDecoder(av).Decode(ncga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncga.Type(), ncga.LocalName(), err)
	}
	return nil
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) State() (*networkmanagerCustomerGatewayAssociationState, bool) {
	return ncga.state, ncga.state != nil
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) StateMust() *networkmanagerCustomerGatewayAssociationState {
	if ncga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncga.Type(), ncga.LocalName()))
	}
	return ncga.state
}

func (ncga *NetworkmanagerCustomerGatewayAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ncga)
}

type NetworkmanagerCustomerGatewayAssociationArgs struct {
	// CustomerGatewayArn: string, required
	CustomerGatewayArn terra.StringValue `hcl:"customer_gateway_arn,attr" validate:"required"`
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, optional
	LinkId terra.StringValue `hcl:"link_id,attr"`
	// Timeouts: optional
	Timeouts *networkmanagercustomergatewayassociation.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NetworkmanagerCustomerGatewayAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkmanagerCustomerGatewayAssociationAttributes struct {
	ref terra.Reference
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) CustomerGatewayArn() terra.StringValue {
	return terra.ReferenceString(ncga.ref.Append("customer_gateway_arn"))
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceString(ncga.ref.Append("device_id"))
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceString(ncga.ref.Append("global_network_id"))
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ncga.ref.Append("id"))
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) LinkId() terra.StringValue {
	return terra.ReferenceString(ncga.ref.Append("link_id"))
}

func (ncga networkmanagerCustomerGatewayAssociationAttributes) Timeouts() networkmanagercustomergatewayassociation.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanagercustomergatewayassociation.TimeoutsAttributes](ncga.ref.Append("timeouts"))
}

type networkmanagerCustomerGatewayAssociationState struct {
	CustomerGatewayArn string                                                  `json:"customer_gateway_arn"`
	DeviceId           string                                                  `json:"device_id"`
	GlobalNetworkId    string                                                  `json:"global_network_id"`
	Id                 string                                                  `json:"id"`
	LinkId             string                                                  `json:"link_id"`
	Timeouts           *networkmanagercustomergatewayassociation.TimeoutsState `json:"timeouts"`
}
