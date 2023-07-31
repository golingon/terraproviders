// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databastionhost "github.com/golingon/terraproviders/azurerm/3.67.0/databastionhost"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBastionHost creates a new instance of [DataBastionHost].
func NewDataBastionHost(name string, args DataBastionHostArgs) *DataBastionHost {
	return &DataBastionHost{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBastionHost)(nil)

// DataBastionHost represents the Terraform data resource azurerm_bastion_host.
type DataBastionHost struct {
	Name string
	Args DataBastionHostArgs
}

// DataSource returns the Terraform object type for [DataBastionHost].
func (bh *DataBastionHost) DataSource() string {
	return "azurerm_bastion_host"
}

// LocalName returns the local name for [DataBastionHost].
func (bh *DataBastionHost) LocalName() string {
	return bh.Name
}

// Configuration returns the configuration (args) for [DataBastionHost].
func (bh *DataBastionHost) Configuration() interface{} {
	return bh.Args
}

// Attributes returns the attributes for [DataBastionHost].
func (bh *DataBastionHost) Attributes() dataBastionHostAttributes {
	return dataBastionHostAttributes{ref: terra.ReferenceDataResource(bh)}
}

// DataBastionHostArgs contains the configurations for azurerm_bastion_host.
type DataBastionHostArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// IpConfiguration: min=0
	IpConfiguration []databastionhost.IpConfiguration `hcl:"ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *databastionhost.Timeouts `hcl:"timeouts,block"`
}
type dataBastionHostAttributes struct {
	ref terra.Reference
}

// CopyPasteEnabled returns a reference to field copy_paste_enabled of azurerm_bastion_host.
func (bh dataBastionHostAttributes) CopyPasteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bh.ref.Append("copy_paste_enabled"))
}

// DnsName returns a reference to field dns_name of azurerm_bastion_host.
func (bh dataBastionHostAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("dns_name"))
}

// FileCopyEnabled returns a reference to field file_copy_enabled of azurerm_bastion_host.
func (bh dataBastionHostAttributes) FileCopyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bh.ref.Append("file_copy_enabled"))
}

// Id returns a reference to field id of azurerm_bastion_host.
func (bh dataBastionHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("id"))
}

// IpConnectEnabled returns a reference to field ip_connect_enabled of azurerm_bastion_host.
func (bh dataBastionHostAttributes) IpConnectEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bh.ref.Append("ip_connect_enabled"))
}

// Location returns a reference to field location of azurerm_bastion_host.
func (bh dataBastionHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_bastion_host.
func (bh dataBastionHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bastion_host.
func (bh dataBastionHostAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("resource_group_name"))
}

// ScaleUnits returns a reference to field scale_units of azurerm_bastion_host.
func (bh dataBastionHostAttributes) ScaleUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(bh.ref.Append("scale_units"))
}

// ShareableLinkEnabled returns a reference to field shareable_link_enabled of azurerm_bastion_host.
func (bh dataBastionHostAttributes) ShareableLinkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bh.ref.Append("shareable_link_enabled"))
}

// Sku returns a reference to field sku of azurerm_bastion_host.
func (bh dataBastionHostAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(bh.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_bastion_host.
func (bh dataBastionHostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bh.ref.Append("tags"))
}

// TunnelingEnabled returns a reference to field tunneling_enabled of azurerm_bastion_host.
func (bh dataBastionHostAttributes) TunnelingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bh.ref.Append("tunneling_enabled"))
}

func (bh dataBastionHostAttributes) IpConfiguration() terra.ListValue[databastionhost.IpConfigurationAttributes] {
	return terra.ReferenceAsList[databastionhost.IpConfigurationAttributes](bh.ref.Append("ip_configuration"))
}

func (bh dataBastionHostAttributes) Timeouts() databastionhost.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databastionhost.TimeoutsAttributes](bh.ref.Append("timeouts"))
}
