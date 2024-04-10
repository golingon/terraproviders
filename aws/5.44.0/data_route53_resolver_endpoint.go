// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	dataroute53resolverendpoint "github.com/golingon/terraproviders/aws/5.44.0/dataroute53resolverendpoint"
)

// NewDataRoute53ResolverEndpoint creates a new instance of [DataRoute53ResolverEndpoint].
func NewDataRoute53ResolverEndpoint(name string, args DataRoute53ResolverEndpointArgs) *DataRoute53ResolverEndpoint {
	return &DataRoute53ResolverEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53ResolverEndpoint)(nil)

// DataRoute53ResolverEndpoint represents the Terraform data resource aws_route53_resolver_endpoint.
type DataRoute53ResolverEndpoint struct {
	Name string
	Args DataRoute53ResolverEndpointArgs
}

// DataSource returns the Terraform object type for [DataRoute53ResolverEndpoint].
func (rre *DataRoute53ResolverEndpoint) DataSource() string {
	return "aws_route53_resolver_endpoint"
}

// LocalName returns the local name for [DataRoute53ResolverEndpoint].
func (rre *DataRoute53ResolverEndpoint) LocalName() string {
	return rre.Name
}

// Configuration returns the configuration (args) for [DataRoute53ResolverEndpoint].
func (rre *DataRoute53ResolverEndpoint) Configuration() interface{} {
	return rre.Args
}

// Attributes returns the attributes for [DataRoute53ResolverEndpoint].
func (rre *DataRoute53ResolverEndpoint) Attributes() dataRoute53ResolverEndpointAttributes {
	return dataRoute53ResolverEndpointAttributes{ref: terra.ReferenceDataResource(rre)}
}

// DataRoute53ResolverEndpointArgs contains the configurations for aws_route53_resolver_endpoint.
type DataRoute53ResolverEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResolverEndpointId: string, optional
	ResolverEndpointId terra.StringValue `hcl:"resolver_endpoint_id,attr"`
	// Filter: min=0
	Filter []dataroute53resolverendpoint.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataRoute53ResolverEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("arn"))
}

// Direction returns a reference to field direction of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("direction"))
}

// Id returns a reference to field id of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("id"))
}

// IpAddresses returns a reference to field ip_addresses of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rre.ref.Append("ip_addresses"))
}

// Name returns a reference to field name of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("name"))
}

// Protocols returns a reference to field protocols of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rre.ref.Append("protocols"))
}

// ResolverEndpointId returns a reference to field resolver_endpoint_id of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) ResolverEndpointId() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("resolver_endpoint_id"))
}

// ResolverEndpointType returns a reference to field resolver_endpoint_type of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) ResolverEndpointType() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("resolver_endpoint_type"))
}

// Status returns a reference to field status of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("status"))
}

// VpcId returns a reference to field vpc_id of aws_route53_resolver_endpoint.
func (rre dataRoute53ResolverEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(rre.ref.Append("vpc_id"))
}

func (rre dataRoute53ResolverEndpointAttributes) Filter() terra.SetValue[dataroute53resolverendpoint.FilterAttributes] {
	return terra.ReferenceAsSet[dataroute53resolverendpoint.FilterAttributes](rre.ref.Append("filter"))
}
