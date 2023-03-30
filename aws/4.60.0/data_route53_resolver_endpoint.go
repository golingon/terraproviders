// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataroute53resolverendpoint "github.com/golingon/terraproviders/aws/4.60.0/dataroute53resolverendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataRoute53ResolverEndpoint(name string, args DataRoute53ResolverEndpointArgs) *DataRoute53ResolverEndpoint {
	return &DataRoute53ResolverEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53ResolverEndpoint)(nil)

type DataRoute53ResolverEndpoint struct {
	Name string
	Args DataRoute53ResolverEndpointArgs
}

func (rre *DataRoute53ResolverEndpoint) DataSource() string {
	return "aws_route53_resolver_endpoint"
}

func (rre *DataRoute53ResolverEndpoint) LocalName() string {
	return rre.Name
}

func (rre *DataRoute53ResolverEndpoint) Configuration() interface{} {
	return rre.Args
}

func (rre *DataRoute53ResolverEndpoint) Attributes() dataRoute53ResolverEndpointAttributes {
	return dataRoute53ResolverEndpointAttributes{ref: terra.ReferenceDataResource(rre)}
}

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

func (rre dataRoute53ResolverEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("arn"))
}

func (rre dataRoute53ResolverEndpointAttributes) Direction() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("direction"))
}

func (rre dataRoute53ResolverEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("id"))
}

func (rre dataRoute53ResolverEndpointAttributes) IpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rre.ref.Append("ip_addresses"))
}

func (rre dataRoute53ResolverEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("name"))
}

func (rre dataRoute53ResolverEndpointAttributes) ResolverEndpointId() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("resolver_endpoint_id"))
}

func (rre dataRoute53ResolverEndpointAttributes) Status() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("status"))
}

func (rre dataRoute53ResolverEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(rre.ref.Append("vpc_id"))
}

func (rre dataRoute53ResolverEndpointAttributes) Filter() terra.SetValue[dataroute53resolverendpoint.FilterAttributes] {
	return terra.ReferenceSet[dataroute53resolverendpoint.FilterAttributes](rre.ref.Append("filter"))
}
