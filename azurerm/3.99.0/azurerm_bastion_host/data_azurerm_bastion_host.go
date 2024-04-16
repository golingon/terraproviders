// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_bastion_host

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_bastion_host.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (abh *DataSource) DataSource() string {
	return "azurerm_bastion_host"
}

// LocalName returns the local name for [DataSource].
func (abh *DataSource) LocalName() string {
	return abh.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (abh *DataSource) Configuration() interface{} {
	return abh.Args
}

// Attributes returns the attributes for [DataSource].
func (abh *DataSource) Attributes() dataAzurermBastionHostAttributes {
	return dataAzurermBastionHostAttributes{ref: terra.ReferenceDataSource(abh)}
}

// DataArgs contains the configurations for azurerm_bastion_host.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermBastionHostAttributes struct {
	ref terra.Reference
}

// CopyPasteEnabled returns a reference to field copy_paste_enabled of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) CopyPasteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("copy_paste_enabled"))
}

// DnsName returns a reference to field dns_name of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("dns_name"))
}

// FileCopyEnabled returns a reference to field file_copy_enabled of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) FileCopyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("file_copy_enabled"))
}

// Id returns a reference to field id of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("id"))
}

// IpConnectEnabled returns a reference to field ip_connect_enabled of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) IpConnectEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("ip_connect_enabled"))
}

// Location returns a reference to field location of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("resource_group_name"))
}

// ScaleUnits returns a reference to field scale_units of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) ScaleUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(abh.ref.Append("scale_units"))
}

// ShareableLinkEnabled returns a reference to field shareable_link_enabled of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) ShareableLinkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("shareable_link_enabled"))
}

// Sku returns a reference to field sku of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abh.ref.Append("tags"))
}

// TunnelingEnabled returns a reference to field tunneling_enabled of azurerm_bastion_host.
func (abh dataAzurermBastionHostAttributes) TunnelingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("tunneling_enabled"))
}

func (abh dataAzurermBastionHostAttributes) IpConfiguration() terra.ListValue[DataIpConfigurationAttributes] {
	return terra.ReferenceAsList[DataIpConfigurationAttributes](abh.ref.Append("ip_configuration"))
}

func (abh dataAzurermBastionHostAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](abh.ref.Append("timeouts"))
}
