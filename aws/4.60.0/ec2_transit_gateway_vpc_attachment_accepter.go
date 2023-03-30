// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEc2TransitGatewayVpcAttachmentAccepter(name string, args Ec2TransitGatewayVpcAttachmentAccepterArgs) *Ec2TransitGatewayVpcAttachmentAccepter {
	return &Ec2TransitGatewayVpcAttachmentAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayVpcAttachmentAccepter)(nil)

type Ec2TransitGatewayVpcAttachmentAccepter struct {
	Name  string
	Args  Ec2TransitGatewayVpcAttachmentAccepterArgs
	state *ec2TransitGatewayVpcAttachmentAccepterState
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) Type() string {
	return "aws_ec2_transit_gateway_vpc_attachment_accepter"
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) LocalName() string {
	return etgvaa.Name
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) Configuration() interface{} {
	return etgvaa.Args
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) Attributes() ec2TransitGatewayVpcAttachmentAccepterAttributes {
	return ec2TransitGatewayVpcAttachmentAccepterAttributes{ref: terra.ReferenceResource(etgvaa)}
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) ImportState(av io.Reader) error {
	etgvaa.state = &ec2TransitGatewayVpcAttachmentAccepterState{}
	if err := json.NewDecoder(av).Decode(etgvaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgvaa.Type(), etgvaa.LocalName(), err)
	}
	return nil
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) State() (*ec2TransitGatewayVpcAttachmentAccepterState, bool) {
	return etgvaa.state, etgvaa.state != nil
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) StateMust() *ec2TransitGatewayVpcAttachmentAccepterState {
	if etgvaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgvaa.Type(), etgvaa.LocalName()))
	}
	return etgvaa.state
}

func (etgvaa *Ec2TransitGatewayVpcAttachmentAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(etgvaa)
}

type Ec2TransitGatewayVpcAttachmentAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayAttachmentId: string, required
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr" validate:"required"`
	// TransitGatewayDefaultRouteTableAssociation: bool, optional
	TransitGatewayDefaultRouteTableAssociation terra.BoolValue `hcl:"transit_gateway_default_route_table_association,attr"`
	// TransitGatewayDefaultRouteTablePropagation: bool, optional
	TransitGatewayDefaultRouteTablePropagation terra.BoolValue `hcl:"transit_gateway_default_route_table_propagation,attr"`
	// DependsOn contains resources that Ec2TransitGatewayVpcAttachmentAccepter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ec2TransitGatewayVpcAttachmentAccepterAttributes struct {
	ref terra.Reference
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) ApplianceModeSupport() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("appliance_mode_support"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) DnsSupport() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("dns_support"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("id"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) Ipv6Support() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("ipv6_support"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](etgvaa.ref.Append("subnet_ids"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](etgvaa.ref.Append("tags"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](etgvaa.ref.Append("tags_all"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("transit_gateway_attachment_id"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) TransitGatewayDefaultRouteTableAssociation() terra.BoolValue {
	return terra.ReferenceBool(etgvaa.ref.Append("transit_gateway_default_route_table_association"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) TransitGatewayDefaultRouteTablePropagation() terra.BoolValue {
	return terra.ReferenceBool(etgvaa.ref.Append("transit_gateway_default_route_table_propagation"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("transit_gateway_id"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("vpc_id"))
}

func (etgvaa ec2TransitGatewayVpcAttachmentAccepterAttributes) VpcOwnerId() terra.StringValue {
	return terra.ReferenceString(etgvaa.ref.Append("vpc_owner_id"))
}

type ec2TransitGatewayVpcAttachmentAccepterState struct {
	ApplianceModeSupport                       string            `json:"appliance_mode_support"`
	DnsSupport                                 string            `json:"dns_support"`
	Id                                         string            `json:"id"`
	Ipv6Support                                string            `json:"ipv6_support"`
	SubnetIds                                  []string          `json:"subnet_ids"`
	Tags                                       map[string]string `json:"tags"`
	TagsAll                                    map[string]string `json:"tags_all"`
	TransitGatewayAttachmentId                 string            `json:"transit_gateway_attachment_id"`
	TransitGatewayDefaultRouteTableAssociation bool              `json:"transit_gateway_default_route_table_association"`
	TransitGatewayDefaultRouteTablePropagation bool              `json:"transit_gateway_default_route_table_propagation"`
	TransitGatewayId                           string            `json:"transit_gateway_id"`
	VpcId                                      string            `json:"vpc_id"`
	VpcOwnerId                                 string            `json:"vpc_owner_id"`
}
