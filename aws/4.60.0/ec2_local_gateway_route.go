// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEc2LocalGatewayRoute(name string, args Ec2LocalGatewayRouteArgs) *Ec2LocalGatewayRoute {
	return &Ec2LocalGatewayRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2LocalGatewayRoute)(nil)

type Ec2LocalGatewayRoute struct {
	Name  string
	Args  Ec2LocalGatewayRouteArgs
	state *ec2LocalGatewayRouteState
}

func (elgr *Ec2LocalGatewayRoute) Type() string {
	return "aws_ec2_local_gateway_route"
}

func (elgr *Ec2LocalGatewayRoute) LocalName() string {
	return elgr.Name
}

func (elgr *Ec2LocalGatewayRoute) Configuration() interface{} {
	return elgr.Args
}

func (elgr *Ec2LocalGatewayRoute) Attributes() ec2LocalGatewayRouteAttributes {
	return ec2LocalGatewayRouteAttributes{ref: terra.ReferenceResource(elgr)}
}

func (elgr *Ec2LocalGatewayRoute) ImportState(av io.Reader) error {
	elgr.state = &ec2LocalGatewayRouteState{}
	if err := json.NewDecoder(av).Decode(elgr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", elgr.Type(), elgr.LocalName(), err)
	}
	return nil
}

func (elgr *Ec2LocalGatewayRoute) State() (*ec2LocalGatewayRouteState, bool) {
	return elgr.state, elgr.state != nil
}

func (elgr *Ec2LocalGatewayRoute) StateMust() *ec2LocalGatewayRouteState {
	if elgr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", elgr.Type(), elgr.LocalName()))
	}
	return elgr.state
}

func (elgr *Ec2LocalGatewayRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(elgr)
}

type Ec2LocalGatewayRouteArgs struct {
	// DestinationCidrBlock: string, required
	DestinationCidrBlock terra.StringValue `hcl:"destination_cidr_block,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalGatewayRouteTableId: string, required
	LocalGatewayRouteTableId terra.StringValue `hcl:"local_gateway_route_table_id,attr" validate:"required"`
	// LocalGatewayVirtualInterfaceGroupId: string, required
	LocalGatewayVirtualInterfaceGroupId terra.StringValue `hcl:"local_gateway_virtual_interface_group_id,attr" validate:"required"`
	// DependsOn contains resources that Ec2LocalGatewayRoute depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ec2LocalGatewayRouteAttributes struct {
	ref terra.Reference
}

func (elgr ec2LocalGatewayRouteAttributes) DestinationCidrBlock() terra.StringValue {
	return terra.ReferenceString(elgr.ref.Append("destination_cidr_block"))
}

func (elgr ec2LocalGatewayRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceString(elgr.ref.Append("id"))
}

func (elgr ec2LocalGatewayRouteAttributes) LocalGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceString(elgr.ref.Append("local_gateway_route_table_id"))
}

func (elgr ec2LocalGatewayRouteAttributes) LocalGatewayVirtualInterfaceGroupId() terra.StringValue {
	return terra.ReferenceString(elgr.ref.Append("local_gateway_virtual_interface_group_id"))
}

type ec2LocalGatewayRouteState struct {
	DestinationCidrBlock                string `json:"destination_cidr_block"`
	Id                                  string `json:"id"`
	LocalGatewayRouteTableId            string `json:"local_gateway_route_table_id"`
	LocalGatewayVirtualInterfaceGroupId string `json:"local_gateway_virtual_interface_group_id"`
}
