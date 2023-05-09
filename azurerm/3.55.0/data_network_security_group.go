// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetworksecuritygroup "github.com/golingon/terraproviders/azurerm/3.55.0/datanetworksecuritygroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkSecurityGroup creates a new instance of [DataNetworkSecurityGroup].
func NewDataNetworkSecurityGroup(name string, args DataNetworkSecurityGroupArgs) *DataNetworkSecurityGroup {
	return &DataNetworkSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkSecurityGroup)(nil)

// DataNetworkSecurityGroup represents the Terraform data resource azurerm_network_security_group.
type DataNetworkSecurityGroup struct {
	Name string
	Args DataNetworkSecurityGroupArgs
}

// DataSource returns the Terraform object type for [DataNetworkSecurityGroup].
func (nsg *DataNetworkSecurityGroup) DataSource() string {
	return "azurerm_network_security_group"
}

// LocalName returns the local name for [DataNetworkSecurityGroup].
func (nsg *DataNetworkSecurityGroup) LocalName() string {
	return nsg.Name
}

// Configuration returns the configuration (args) for [DataNetworkSecurityGroup].
func (nsg *DataNetworkSecurityGroup) Configuration() interface{} {
	return nsg.Args
}

// Attributes returns the attributes for [DataNetworkSecurityGroup].
func (nsg *DataNetworkSecurityGroup) Attributes() dataNetworkSecurityGroupAttributes {
	return dataNetworkSecurityGroupAttributes{ref: terra.ReferenceDataResource(nsg)}
}

// DataNetworkSecurityGroupArgs contains the configurations for azurerm_network_security_group.
type DataNetworkSecurityGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SecurityRule: min=0
	SecurityRule []datanetworksecuritygroup.SecurityRule `hcl:"security_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetworksecuritygroup.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkSecurityGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_security_group.
func (nsg dataNetworkSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_security_group.
func (nsg dataNetworkSecurityGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_security_group.
func (nsg dataNetworkSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_security_group.
func (nsg dataNetworkSecurityGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_security_group.
func (nsg dataNetworkSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsg.ref.Append("tags"))
}

func (nsg dataNetworkSecurityGroupAttributes) SecurityRule() terra.ListValue[datanetworksecuritygroup.SecurityRuleAttributes] {
	return terra.ReferenceAsList[datanetworksecuritygroup.SecurityRuleAttributes](nsg.ref.Append("security_rule"))
}

func (nsg dataNetworkSecurityGroupAttributes) Timeouts() datanetworksecuritygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworksecuritygroup.TimeoutsAttributes](nsg.ref.Append("timeouts"))
}
