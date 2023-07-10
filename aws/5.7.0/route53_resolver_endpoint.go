// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53resolverendpoint "github.com/golingon/terraproviders/aws/5.7.0/route53resolverendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53ResolverEndpoint creates a new instance of [Route53ResolverEndpoint].
func NewRoute53ResolverEndpoint(name string, args Route53ResolverEndpointArgs) *Route53ResolverEndpoint {
	return &Route53ResolverEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverEndpoint)(nil)

// Route53ResolverEndpoint represents the Terraform resource aws_route53_resolver_endpoint.
type Route53ResolverEndpoint struct {
	Name      string
	Args      Route53ResolverEndpointArgs
	state     *route53ResolverEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) Type() string {
	return "aws_route53_resolver_endpoint"
}

// LocalName returns the local name for [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) LocalName() string {
	return rre.Name
}

// Configuration returns the configuration (args) for [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) Configuration() interface{} {
	return rre.Args
}

// DependOn is used for other resources to depend on [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(rre)
}

// Dependencies returns the list of resources [Route53ResolverEndpoint] depends_on.
func (rre *Route53ResolverEndpoint) Dependencies() terra.Dependencies {
	return rre.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) LifecycleManagement() *terra.Lifecycle {
	return rre.Lifecycle
}

// Attributes returns the attributes for [Route53ResolverEndpoint].
func (rre *Route53ResolverEndpoint) Attributes() route53ResolverEndpointAttributes {
	return route53ResolverEndpointAttributes{ref: terra.ReferenceResource(rre)}
}

// ImportState imports the given attribute values into [Route53ResolverEndpoint]'s state.
func (rre *Route53ResolverEndpoint) ImportState(av io.Reader) error {
	rre.state = &route53ResolverEndpointState{}
	if err := json.NewDecoder(av).Decode(rre.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rre.Type(), rre.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53ResolverEndpoint] has state.
func (rre *Route53ResolverEndpoint) State() (*route53ResolverEndpointState, bool) {
	return rre.state, rre.state != nil
}

// StateMust returns the state for [Route53ResolverEndpoint]. Panics if the state is nil.
func (rre *Route53ResolverEndpoint) StateMust() *route53ResolverEndpointState {
	if rre.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rre.Type(), rre.LocalName()))
	}
	return rre.state
}

// Route53ResolverEndpointArgs contains the configurations for aws_route53_resolver_endpoint.
type Route53ResolverEndpointArgs struct {
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// SecurityGroupIds: set of string, required
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// IpAddress: min=2,max=10
	IpAddress []route53resolverendpoint.IpAddress `hcl:"ip_address,block" validate:"min=2,max=10"`
	// Timeouts: optional
	Timeouts *route53resolverendpoint.Timeouts `hcl:"timeouts,block"`
}
type route53ResolverEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("arn"))
}

// Direction returns a reference to field direction of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("direction"))
}

// HostVpcId returns a reference to field host_vpc_id of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) HostVpcId() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("host_vpc_id"))
}

// Id returns a reference to field id of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("name"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rre.ref.Append("security_group_ids"))
}

// Tags returns a reference to field tags of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rre.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53_resolver_endpoint.
func (rre route53ResolverEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rre.ref.Append("tags_all"))
}

func (rre route53ResolverEndpointAttributes) IpAddress() terra.SetValue[route53resolverendpoint.IpAddressAttributes] {
	return terra.ReferenceAsSet[route53resolverendpoint.IpAddressAttributes](rre.ref.Append("ip_address"))
}

func (rre route53ResolverEndpointAttributes) Timeouts() route53resolverendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53resolverendpoint.TimeoutsAttributes](rre.ref.Append("timeouts"))
}

type route53ResolverEndpointState struct {
	Arn              string                                   `json:"arn"`
	Direction        string                                   `json:"direction"`
	HostVpcId        string                                   `json:"host_vpc_id"`
	Id               string                                   `json:"id"`
	Name             string                                   `json:"name"`
	SecurityGroupIds []string                                 `json:"security_group_ids"`
	Tags             map[string]string                        `json:"tags"`
	TagsAll          map[string]string                        `json:"tags_all"`
	IpAddress        []route53resolverendpoint.IpAddressState `json:"ip_address"`
	Timeouts         *route53resolverendpoint.TimeoutsState   `json:"timeouts"`
}
