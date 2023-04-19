// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapublicipprefix "github.com/golingon/terraproviders/azurerm/3.52.0/datapublicipprefix"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPublicIpPrefix creates a new instance of [DataPublicIpPrefix].
func NewDataPublicIpPrefix(name string, args DataPublicIpPrefixArgs) *DataPublicIpPrefix {
	return &DataPublicIpPrefix{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPublicIpPrefix)(nil)

// DataPublicIpPrefix represents the Terraform data resource azurerm_public_ip_prefix.
type DataPublicIpPrefix struct {
	Name string
	Args DataPublicIpPrefixArgs
}

// DataSource returns the Terraform object type for [DataPublicIpPrefix].
func (pip *DataPublicIpPrefix) DataSource() string {
	return "azurerm_public_ip_prefix"
}

// LocalName returns the local name for [DataPublicIpPrefix].
func (pip *DataPublicIpPrefix) LocalName() string {
	return pip.Name
}

// Configuration returns the configuration (args) for [DataPublicIpPrefix].
func (pip *DataPublicIpPrefix) Configuration() interface{} {
	return pip.Args
}

// Attributes returns the attributes for [DataPublicIpPrefix].
func (pip *DataPublicIpPrefix) Attributes() dataPublicIpPrefixAttributes {
	return dataPublicIpPrefixAttributes{ref: terra.ReferenceDataResource(pip)}
}

// DataPublicIpPrefixArgs contains the configurations for azurerm_public_ip_prefix.
type DataPublicIpPrefixArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datapublicipprefix.Timeouts `hcl:"timeouts,block"`
}
type dataPublicIpPrefixAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("id"))
}

// IpPrefix returns a reference to field ip_prefix of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) IpPrefix() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("ip_prefix"))
}

// Location returns a reference to field location of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("name"))
}

// PrefixLength returns a reference to field prefix_length of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(pip.ref.Append("prefix_length"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pip.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_public_ip_prefix.
func (pip dataPublicIpPrefixAttributes) Zones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pip.ref.Append("zones"))
}

func (pip dataPublicIpPrefixAttributes) Timeouts() datapublicipprefix.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapublicipprefix.TimeoutsAttributes](pip.ref.Append("timeouts"))
}
