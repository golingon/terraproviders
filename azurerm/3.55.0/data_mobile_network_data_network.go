// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetworkdatanetwork "github.com/golingon/terraproviders/azurerm/3.55.0/datamobilenetworkdatanetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetworkDataNetwork creates a new instance of [DataMobileNetworkDataNetwork].
func NewDataMobileNetworkDataNetwork(name string, args DataMobileNetworkDataNetworkArgs) *DataMobileNetworkDataNetwork {
	return &DataMobileNetworkDataNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkDataNetwork)(nil)

// DataMobileNetworkDataNetwork represents the Terraform data resource azurerm_mobile_network_data_network.
type DataMobileNetworkDataNetwork struct {
	Name string
	Args DataMobileNetworkDataNetworkArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkDataNetwork].
func (mndn *DataMobileNetworkDataNetwork) DataSource() string {
	return "azurerm_mobile_network_data_network"
}

// LocalName returns the local name for [DataMobileNetworkDataNetwork].
func (mndn *DataMobileNetworkDataNetwork) LocalName() string {
	return mndn.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkDataNetwork].
func (mndn *DataMobileNetworkDataNetwork) Configuration() interface{} {
	return mndn.Args
}

// Attributes returns the attributes for [DataMobileNetworkDataNetwork].
func (mndn *DataMobileNetworkDataNetwork) Attributes() dataMobileNetworkDataNetworkAttributes {
	return dataMobileNetworkDataNetworkAttributes{ref: terra.ReferenceDataResource(mndn)}
}

// DataMobileNetworkDataNetworkArgs contains the configurations for azurerm_mobile_network_data_network.
type DataMobileNetworkDataNetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamobilenetworkdatanetwork.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkDataNetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mndn.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mndn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mndn.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mndn.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mndn.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_data_network.
func (mndn dataMobileNetworkDataNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mndn.ref.Append("tags"))
}

func (mndn dataMobileNetworkDataNetworkAttributes) Timeouts() datamobilenetworkdatanetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworkdatanetwork.TimeoutsAttributes](mndn.ref.Append("timeouts"))
}
