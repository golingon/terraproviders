// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewaymulticastdomain "github.com/golingon/terraproviders/aws/5.6.2/dataec2transitgatewaymulticastdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayMulticastDomain creates a new instance of [DataEc2TransitGatewayMulticastDomain].
func NewDataEc2TransitGatewayMulticastDomain(name string, args DataEc2TransitGatewayMulticastDomainArgs) *DataEc2TransitGatewayMulticastDomain {
	return &DataEc2TransitGatewayMulticastDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayMulticastDomain)(nil)

// DataEc2TransitGatewayMulticastDomain represents the Terraform data resource aws_ec2_transit_gateway_multicast_domain.
type DataEc2TransitGatewayMulticastDomain struct {
	Name string
	Args DataEc2TransitGatewayMulticastDomainArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayMulticastDomain].
func (etgmd *DataEc2TransitGatewayMulticastDomain) DataSource() string {
	return "aws_ec2_transit_gateway_multicast_domain"
}

// LocalName returns the local name for [DataEc2TransitGatewayMulticastDomain].
func (etgmd *DataEc2TransitGatewayMulticastDomain) LocalName() string {
	return etgmd.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayMulticastDomain].
func (etgmd *DataEc2TransitGatewayMulticastDomain) Configuration() interface{} {
	return etgmd.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayMulticastDomain].
func (etgmd *DataEc2TransitGatewayMulticastDomain) Attributes() dataEc2TransitGatewayMulticastDomainAttributes {
	return dataEc2TransitGatewayMulticastDomainAttributes{ref: terra.ReferenceDataResource(etgmd)}
}

// DataEc2TransitGatewayMulticastDomainArgs contains the configurations for aws_ec2_transit_gateway_multicast_domain.
type DataEc2TransitGatewayMulticastDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransitGatewayMulticastDomainId: string, optional
	TransitGatewayMulticastDomainId terra.StringValue `hcl:"transit_gateway_multicast_domain_id,attr"`
	// Associations: min=0
	Associations []dataec2transitgatewaymulticastdomain.Associations `hcl:"associations,block" validate:"min=0"`
	// Members: min=0
	Members []dataec2transitgatewaymulticastdomain.Members `hcl:"members,block" validate:"min=0"`
	// Sources: min=0
	Sources []dataec2transitgatewaymulticastdomain.Sources `hcl:"sources,block" validate:"min=0"`
	// Filter: min=0
	Filter []dataec2transitgatewaymulticastdomain.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewaymulticastdomain.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayMulticastDomainAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("arn"))
}

// AutoAcceptSharedAssociations returns a reference to field auto_accept_shared_associations of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) AutoAcceptSharedAssociations() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("auto_accept_shared_associations"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("id"))
}

// Igmpv2Support returns a reference to field igmpv2_support of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Igmpv2Support() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("igmpv2_support"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("owner_id"))
}

// State returns a reference to field state of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("state"))
}

// StaticSourcesSupport returns a reference to field static_sources_support of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) StaticSourcesSupport() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("static_sources_support"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgmd.ref.Append("tags"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("transit_gateway_id"))
}

// TransitGatewayMulticastDomainId returns a reference to field transit_gateway_multicast_domain_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) TransitGatewayMulticastDomainId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("transit_gateway_multicast_domain_id"))
}

func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Associations() terra.ListValue[dataec2transitgatewaymulticastdomain.AssociationsAttributes] {
	return terra.ReferenceAsList[dataec2transitgatewaymulticastdomain.AssociationsAttributes](etgmd.ref.Append("associations"))
}

func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Members() terra.ListValue[dataec2transitgatewaymulticastdomain.MembersAttributes] {
	return terra.ReferenceAsList[dataec2transitgatewaymulticastdomain.MembersAttributes](etgmd.ref.Append("members"))
}

func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Sources() terra.ListValue[dataec2transitgatewaymulticastdomain.SourcesAttributes] {
	return terra.ReferenceAsList[dataec2transitgatewaymulticastdomain.SourcesAttributes](etgmd.ref.Append("sources"))
}

func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Filter() terra.SetValue[dataec2transitgatewaymulticastdomain.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewaymulticastdomain.FilterAttributes](etgmd.ref.Append("filter"))
}

func (etgmd dataEc2TransitGatewayMulticastDomainAttributes) Timeouts() dataec2transitgatewaymulticastdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewaymulticastdomain.TimeoutsAttributes](etgmd.ref.Append("timeouts"))
}
