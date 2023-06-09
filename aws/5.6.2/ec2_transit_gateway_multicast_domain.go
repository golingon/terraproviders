// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2transitgatewaymulticastdomain "github.com/golingon/terraproviders/aws/5.6.2/ec2transitgatewaymulticastdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayMulticastDomain creates a new instance of [Ec2TransitGatewayMulticastDomain].
func NewEc2TransitGatewayMulticastDomain(name string, args Ec2TransitGatewayMulticastDomainArgs) *Ec2TransitGatewayMulticastDomain {
	return &Ec2TransitGatewayMulticastDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayMulticastDomain)(nil)

// Ec2TransitGatewayMulticastDomain represents the Terraform resource aws_ec2_transit_gateway_multicast_domain.
type Ec2TransitGatewayMulticastDomain struct {
	Name      string
	Args      Ec2TransitGatewayMulticastDomainArgs
	state     *ec2TransitGatewayMulticastDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) Type() string {
	return "aws_ec2_transit_gateway_multicast_domain"
}

// LocalName returns the local name for [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) LocalName() string {
	return etgmd.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) Configuration() interface{} {
	return etgmd.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(etgmd)
}

// Dependencies returns the list of resources [Ec2TransitGatewayMulticastDomain] depends_on.
func (etgmd *Ec2TransitGatewayMulticastDomain) Dependencies() terra.Dependencies {
	return etgmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) LifecycleManagement() *terra.Lifecycle {
	return etgmd.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayMulticastDomain].
func (etgmd *Ec2TransitGatewayMulticastDomain) Attributes() ec2TransitGatewayMulticastDomainAttributes {
	return ec2TransitGatewayMulticastDomainAttributes{ref: terra.ReferenceResource(etgmd)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayMulticastDomain]'s state.
func (etgmd *Ec2TransitGatewayMulticastDomain) ImportState(av io.Reader) error {
	etgmd.state = &ec2TransitGatewayMulticastDomainState{}
	if err := json.NewDecoder(av).Decode(etgmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgmd.Type(), etgmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayMulticastDomain] has state.
func (etgmd *Ec2TransitGatewayMulticastDomain) State() (*ec2TransitGatewayMulticastDomainState, bool) {
	return etgmd.state, etgmd.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayMulticastDomain]. Panics if the state is nil.
func (etgmd *Ec2TransitGatewayMulticastDomain) StateMust() *ec2TransitGatewayMulticastDomainState {
	if etgmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgmd.Type(), etgmd.LocalName()))
	}
	return etgmd.state
}

// Ec2TransitGatewayMulticastDomainArgs contains the configurations for aws_ec2_transit_gateway_multicast_domain.
type Ec2TransitGatewayMulticastDomainArgs struct {
	// AutoAcceptSharedAssociations: string, optional
	AutoAcceptSharedAssociations terra.StringValue `hcl:"auto_accept_shared_associations,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Igmpv2Support: string, optional
	Igmpv2Support terra.StringValue `hcl:"igmpv2_support,attr"`
	// StaticSourcesSupport: string, optional
	StaticSourcesSupport terra.StringValue `hcl:"static_sources_support,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayId: string, required
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2transitgatewaymulticastdomain.Timeouts `hcl:"timeouts,block"`
}
type ec2TransitGatewayMulticastDomainAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("arn"))
}

// AutoAcceptSharedAssociations returns a reference to field auto_accept_shared_associations of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) AutoAcceptSharedAssociations() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("auto_accept_shared_associations"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("id"))
}

// Igmpv2Support returns a reference to field igmpv2_support of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) Igmpv2Support() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("igmpv2_support"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("owner_id"))
}

// StaticSourcesSupport returns a reference to field static_sources_support of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) StaticSourcesSupport() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("static_sources_support"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgmd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgmd.ref.Append("tags_all"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_multicast_domain.
func (etgmd ec2TransitGatewayMulticastDomainAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgmd.ref.Append("transit_gateway_id"))
}

func (etgmd ec2TransitGatewayMulticastDomainAttributes) Timeouts() ec2transitgatewaymulticastdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2transitgatewaymulticastdomain.TimeoutsAttributes](etgmd.ref.Append("timeouts"))
}

type ec2TransitGatewayMulticastDomainState struct {
	Arn                          string                                          `json:"arn"`
	AutoAcceptSharedAssociations string                                          `json:"auto_accept_shared_associations"`
	Id                           string                                          `json:"id"`
	Igmpv2Support                string                                          `json:"igmpv2_support"`
	OwnerId                      string                                          `json:"owner_id"`
	StaticSourcesSupport         string                                          `json:"static_sources_support"`
	Tags                         map[string]string                               `json:"tags"`
	TagsAll                      map[string]string                               `json:"tags_all"`
	TransitGatewayId             string                                          `json:"transit_gateway_id"`
	Timeouts                     *ec2transitgatewaymulticastdomain.TimeoutsState `json:"timeouts"`
}
