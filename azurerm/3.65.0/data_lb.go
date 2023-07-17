// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalb "github.com/golingon/terraproviders/azurerm/3.65.0/datalb"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLb creates a new instance of [DataLb].
func NewDataLb(name string, args DataLbArgs) *DataLb {
	return &DataLb{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLb)(nil)

// DataLb represents the Terraform data resource azurerm_lb.
type DataLb struct {
	Name string
	Args DataLbArgs
}

// DataSource returns the Terraform object type for [DataLb].
func (l *DataLb) DataSource() string {
	return "azurerm_lb"
}

// LocalName returns the local name for [DataLb].
func (l *DataLb) LocalName() string {
	return l.Name
}

// Configuration returns the configuration (args) for [DataLb].
func (l *DataLb) Configuration() interface{} {
	return l.Args
}

// Attributes returns the attributes for [DataLb].
func (l *DataLb) Attributes() dataLbAttributes {
	return dataLbAttributes{ref: terra.ReferenceDataResource(l)}
}

// DataLbArgs contains the configurations for azurerm_lb.
type DataLbArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// FrontendIpConfiguration: min=0
	FrontendIpConfiguration []datalb.FrontendIpConfiguration `hcl:"frontend_ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalb.Timeouts `hcl:"timeouts,block"`
}
type dataLbAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_lb.
func (l dataLbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_lb.
func (l dataLbAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_lb.
func (l dataLbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of azurerm_lb.
func (l dataLbAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("private_ip_address"))
}

// PrivateIpAddresses returns a reference to field private_ip_addresses of azurerm_lb.
func (l dataLbAttributes) PrivateIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("private_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lb.
func (l dataLbAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_lb.
func (l dataLbAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_lb.
func (l dataLbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags"))
}

func (l dataLbAttributes) FrontendIpConfiguration() terra.ListValue[datalb.FrontendIpConfigurationAttributes] {
	return terra.ReferenceAsList[datalb.FrontendIpConfigurationAttributes](l.ref.Append("frontend_ip_configuration"))
}

func (l dataLbAttributes) Timeouts() datalb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalb.TimeoutsAttributes](l.ref.Append("timeouts"))
}
