// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2transitgateway "github.com/golingon/terraproviders/aws/5.11.0/ec2transitgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGateway creates a new instance of [Ec2TransitGateway].
func NewEc2TransitGateway(name string, args Ec2TransitGatewayArgs) *Ec2TransitGateway {
	return &Ec2TransitGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGateway)(nil)

// Ec2TransitGateway represents the Terraform resource aws_ec2_transit_gateway.
type Ec2TransitGateway struct {
	Name      string
	Args      Ec2TransitGatewayArgs
	state     *ec2TransitGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGateway].
func (etg *Ec2TransitGateway) Type() string {
	return "aws_ec2_transit_gateway"
}

// LocalName returns the local name for [Ec2TransitGateway].
func (etg *Ec2TransitGateway) LocalName() string {
	return etg.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGateway].
func (etg *Ec2TransitGateway) Configuration() interface{} {
	return etg.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGateway].
func (etg *Ec2TransitGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(etg)
}

// Dependencies returns the list of resources [Ec2TransitGateway] depends_on.
func (etg *Ec2TransitGateway) Dependencies() terra.Dependencies {
	return etg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGateway].
func (etg *Ec2TransitGateway) LifecycleManagement() *terra.Lifecycle {
	return etg.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGateway].
func (etg *Ec2TransitGateway) Attributes() ec2TransitGatewayAttributes {
	return ec2TransitGatewayAttributes{ref: terra.ReferenceResource(etg)}
}

// ImportState imports the given attribute values into [Ec2TransitGateway]'s state.
func (etg *Ec2TransitGateway) ImportState(av io.Reader) error {
	etg.state = &ec2TransitGatewayState{}
	if err := json.NewDecoder(av).Decode(etg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etg.Type(), etg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGateway] has state.
func (etg *Ec2TransitGateway) State() (*ec2TransitGatewayState, bool) {
	return etg.state, etg.state != nil
}

// StateMust returns the state for [Ec2TransitGateway]. Panics if the state is nil.
func (etg *Ec2TransitGateway) StateMust() *ec2TransitGatewayState {
	if etg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etg.Type(), etg.LocalName()))
	}
	return etg.state
}

// Ec2TransitGatewayArgs contains the configurations for aws_ec2_transit_gateway.
type Ec2TransitGatewayArgs struct {
	// AmazonSideAsn: number, optional
	AmazonSideAsn terra.NumberValue `hcl:"amazon_side_asn,attr"`
	// AutoAcceptSharedAttachments: string, optional
	AutoAcceptSharedAttachments terra.StringValue `hcl:"auto_accept_shared_attachments,attr"`
	// DefaultRouteTableAssociation: string, optional
	DefaultRouteTableAssociation terra.StringValue `hcl:"default_route_table_association,attr"`
	// DefaultRouteTablePropagation: string, optional
	DefaultRouteTablePropagation terra.StringValue `hcl:"default_route_table_propagation,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DnsSupport: string, optional
	DnsSupport terra.StringValue `hcl:"dns_support,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MulticastSupport: string, optional
	MulticastSupport terra.StringValue `hcl:"multicast_support,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayCidrBlocks: set of string, optional
	TransitGatewayCidrBlocks terra.SetValue[terra.StringValue] `hcl:"transit_gateway_cidr_blocks,attr"`
	// VpnEcmpSupport: string, optional
	VpnEcmpSupport terra.StringValue `hcl:"vpn_ecmp_support,attr"`
	// Timeouts: optional
	Timeouts *ec2transitgateway.Timeouts `hcl:"timeouts,block"`
}
type ec2TransitGatewayAttributes struct {
	ref terra.Reference
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) AmazonSideAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(etg.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("arn"))
}

// AssociationDefaultRouteTableId returns a reference to field association_default_route_table_id of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) AssociationDefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("association_default_route_table_id"))
}

// AutoAcceptSharedAttachments returns a reference to field auto_accept_shared_attachments of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) AutoAcceptSharedAttachments() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("auto_accept_shared_attachments"))
}

// DefaultRouteTableAssociation returns a reference to field default_route_table_association of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) DefaultRouteTableAssociation() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("default_route_table_association"))
}

// DefaultRouteTablePropagation returns a reference to field default_route_table_propagation of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) DefaultRouteTablePropagation() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("default_route_table_propagation"))
}

// Description returns a reference to field description of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("description"))
}

// DnsSupport returns a reference to field dns_support of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) DnsSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("dns_support"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("id"))
}

// MulticastSupport returns a reference to field multicast_support of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) MulticastSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("multicast_support"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("owner_id"))
}

// PropagationDefaultRouteTableId returns a reference to field propagation_default_route_table_id of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) PropagationDefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("propagation_default_route_table_id"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etg.ref.Append("tags_all"))
}

// TransitGatewayCidrBlocks returns a reference to field transit_gateway_cidr_blocks of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) TransitGatewayCidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](etg.ref.Append("transit_gateway_cidr_blocks"))
}

// VpnEcmpSupport returns a reference to field vpn_ecmp_support of aws_ec2_transit_gateway.
func (etg ec2TransitGatewayAttributes) VpnEcmpSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("vpn_ecmp_support"))
}

func (etg ec2TransitGatewayAttributes) Timeouts() ec2transitgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2transitgateway.TimeoutsAttributes](etg.ref.Append("timeouts"))
}

type ec2TransitGatewayState struct {
	AmazonSideAsn                  float64                          `json:"amazon_side_asn"`
	Arn                            string                           `json:"arn"`
	AssociationDefaultRouteTableId string                           `json:"association_default_route_table_id"`
	AutoAcceptSharedAttachments    string                           `json:"auto_accept_shared_attachments"`
	DefaultRouteTableAssociation   string                           `json:"default_route_table_association"`
	DefaultRouteTablePropagation   string                           `json:"default_route_table_propagation"`
	Description                    string                           `json:"description"`
	DnsSupport                     string                           `json:"dns_support"`
	Id                             string                           `json:"id"`
	MulticastSupport               string                           `json:"multicast_support"`
	OwnerId                        string                           `json:"owner_id"`
	PropagationDefaultRouteTableId string                           `json:"propagation_default_route_table_id"`
	Tags                           map[string]string                `json:"tags"`
	TagsAll                        map[string]string                `json:"tags_all"`
	TransitGatewayCidrBlocks       []string                         `json:"transit_gateway_cidr_blocks"`
	VpnEcmpSupport                 string                           `json:"vpn_ecmp_support"`
	Timeouts                       *ec2transitgateway.TimeoutsState `json:"timeouts"`
}