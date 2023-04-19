// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetworksimgroup "github.com/golingon/terraproviders/azurerm/3.52.0/datamobilenetworksimgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetworkSimGroup creates a new instance of [DataMobileNetworkSimGroup].
func NewDataMobileNetworkSimGroup(name string, args DataMobileNetworkSimGroupArgs) *DataMobileNetworkSimGroup {
	return &DataMobileNetworkSimGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkSimGroup)(nil)

// DataMobileNetworkSimGroup represents the Terraform data resource azurerm_mobile_network_sim_group.
type DataMobileNetworkSimGroup struct {
	Name string
	Args DataMobileNetworkSimGroupArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkSimGroup].
func (mnsg *DataMobileNetworkSimGroup) DataSource() string {
	return "azurerm_mobile_network_sim_group"
}

// LocalName returns the local name for [DataMobileNetworkSimGroup].
func (mnsg *DataMobileNetworkSimGroup) LocalName() string {
	return mnsg.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkSimGroup].
func (mnsg *DataMobileNetworkSimGroup) Configuration() interface{} {
	return mnsg.Args
}

// Attributes returns the attributes for [DataMobileNetworkSimGroup].
func (mnsg *DataMobileNetworkSimGroup) Attributes() dataMobileNetworkSimGroupAttributes {
	return dataMobileNetworkSimGroupAttributes{ref: terra.ReferenceDataResource(mnsg)}
}

// DataMobileNetworkSimGroupArgs contains the configurations for azurerm_mobile_network_sim_group.
type DataMobileNetworkSimGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Identity: min=0
	Identity []datamobilenetworksimgroup.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamobilenetworksimgroup.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkSimGroupAttributes struct {
	ref terra.Reference
}

// EncryptionKeyUrl returns a reference to field encryption_key_url of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) EncryptionKeyUrl() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("encryption_key_url"))
}

// Id returns a reference to field id of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_sim_group.
func (mnsg dataMobileNetworkSimGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnsg.ref.Append("tags"))
}

func (mnsg dataMobileNetworkSimGroupAttributes) Identity() terra.ListValue[datamobilenetworksimgroup.IdentityAttributes] {
	return terra.ReferenceAsList[datamobilenetworksimgroup.IdentityAttributes](mnsg.ref.Append("identity"))
}

func (mnsg dataMobileNetworkSimGroupAttributes) Timeouts() datamobilenetworksimgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworksimgroup.TimeoutsAttributes](mnsg.ref.Append("timeouts"))
}
