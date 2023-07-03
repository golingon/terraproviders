// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsresolver "github.com/golingon/terraproviders/azurerm/3.63.0/dataprivatednsresolver"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsResolver creates a new instance of [DataPrivateDnsResolver].
func NewDataPrivateDnsResolver(name string, args DataPrivateDnsResolverArgs) *DataPrivateDnsResolver {
	return &DataPrivateDnsResolver{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsResolver)(nil)

// DataPrivateDnsResolver represents the Terraform data resource azurerm_private_dns_resolver.
type DataPrivateDnsResolver struct {
	Name string
	Args DataPrivateDnsResolverArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsResolver].
func (pdr *DataPrivateDnsResolver) DataSource() string {
	return "azurerm_private_dns_resolver"
}

// LocalName returns the local name for [DataPrivateDnsResolver].
func (pdr *DataPrivateDnsResolver) LocalName() string {
	return pdr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsResolver].
func (pdr *DataPrivateDnsResolver) Configuration() interface{} {
	return pdr.Args
}

// Attributes returns the attributes for [DataPrivateDnsResolver].
func (pdr *DataPrivateDnsResolver) Attributes() dataPrivateDnsResolverAttributes {
	return dataPrivateDnsResolverAttributes{ref: terra.ReferenceDataResource(pdr)}
}

// DataPrivateDnsResolverArgs contains the configurations for azurerm_private_dns_resolver.
type DataPrivateDnsResolverArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednsresolver.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsResolverAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdr.ref.Append("tags"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver.
func (pdr dataPrivateDnsResolverAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdr.ref.Append("virtual_network_id"))
}

func (pdr dataPrivateDnsResolverAttributes) Timeouts() dataprivatednsresolver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsresolver.TimeoutsAttributes](pdr.ref.Append("timeouts"))
}
