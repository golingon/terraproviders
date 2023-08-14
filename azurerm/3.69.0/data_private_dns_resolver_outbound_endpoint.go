// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsresolveroutboundendpoint "github.com/golingon/terraproviders/azurerm/3.69.0/dataprivatednsresolveroutboundendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsResolverOutboundEndpoint creates a new instance of [DataPrivateDnsResolverOutboundEndpoint].
func NewDataPrivateDnsResolverOutboundEndpoint(name string, args DataPrivateDnsResolverOutboundEndpointArgs) *DataPrivateDnsResolverOutboundEndpoint {
	return &DataPrivateDnsResolverOutboundEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsResolverOutboundEndpoint)(nil)

// DataPrivateDnsResolverOutboundEndpoint represents the Terraform data resource azurerm_private_dns_resolver_outbound_endpoint.
type DataPrivateDnsResolverOutboundEndpoint struct {
	Name string
	Args DataPrivateDnsResolverOutboundEndpointArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsResolverOutboundEndpoint].
func (pdroe *DataPrivateDnsResolverOutboundEndpoint) DataSource() string {
	return "azurerm_private_dns_resolver_outbound_endpoint"
}

// LocalName returns the local name for [DataPrivateDnsResolverOutboundEndpoint].
func (pdroe *DataPrivateDnsResolverOutboundEndpoint) LocalName() string {
	return pdroe.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsResolverOutboundEndpoint].
func (pdroe *DataPrivateDnsResolverOutboundEndpoint) Configuration() interface{} {
	return pdroe.Args
}

// Attributes returns the attributes for [DataPrivateDnsResolverOutboundEndpoint].
func (pdroe *DataPrivateDnsResolverOutboundEndpoint) Attributes() dataPrivateDnsResolverOutboundEndpointAttributes {
	return dataPrivateDnsResolverOutboundEndpointAttributes{ref: terra.ReferenceDataResource(pdroe)}
}

// DataPrivateDnsResolverOutboundEndpointArgs contains the configurations for azurerm_private_dns_resolver_outbound_endpoint.
type DataPrivateDnsResolverOutboundEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsResolverId: string, required
	PrivateDnsResolverId terra.StringValue `hcl:"private_dns_resolver_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednsresolveroutboundendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsResolverOutboundEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("name"))
}

// PrivateDnsResolverId returns a reference to field private_dns_resolver_id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) PrivateDnsResolverId() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("private_dns_resolver_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(pdroe.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver_outbound_endpoint.
func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdroe.ref.Append("tags"))
}

func (pdroe dataPrivateDnsResolverOutboundEndpointAttributes) Timeouts() dataprivatednsresolveroutboundendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsresolveroutboundendpoint.TimeoutsAttributes](pdroe.ref.Append("timeouts"))
}
