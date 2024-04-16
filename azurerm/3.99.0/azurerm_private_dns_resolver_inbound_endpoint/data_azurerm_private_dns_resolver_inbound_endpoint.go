// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver_inbound_endpoint

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_dns_resolver_inbound_endpoint.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apdrie *DataSource) DataSource() string {
	return "azurerm_private_dns_resolver_inbound_endpoint"
}

// LocalName returns the local name for [DataSource].
func (apdrie *DataSource) LocalName() string {
	return apdrie.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apdrie *DataSource) Configuration() interface{} {
	return apdrie.Args
}

// Attributes returns the attributes for [DataSource].
func (apdrie *DataSource) Attributes() dataAzurermPrivateDnsResolverInboundEndpointAttributes {
	return dataAzurermPrivateDnsResolverInboundEndpointAttributes{ref: terra.ReferenceDataSource(apdrie)}
}

// DataArgs contains the configurations for azurerm_private_dns_resolver_inbound_endpoint.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsResolverId: string, required
	PrivateDnsResolverId terra.StringValue `hcl:"private_dns_resolver_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPrivateDnsResolverInboundEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver_inbound_endpoint.
func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdrie.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver_inbound_endpoint.
func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(apdrie.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_inbound_endpoint.
func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdrie.ref.Append("name"))
}

// PrivateDnsResolverId returns a reference to field private_dns_resolver_id of azurerm_private_dns_resolver_inbound_endpoint.
func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) PrivateDnsResolverId() terra.StringValue {
	return terra.ReferenceAsString(apdrie.ref.Append("private_dns_resolver_id"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver_inbound_endpoint.
func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdrie.ref.Append("tags"))
}

func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) IpConfigurations() terra.ListValue[DataIpConfigurationsAttributes] {
	return terra.ReferenceAsList[DataIpConfigurationsAttributes](apdrie.ref.Append("ip_configurations"))
}

func (apdrie dataAzurermPrivateDnsResolverInboundEndpointAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apdrie.ref.Append("timeouts"))
}
