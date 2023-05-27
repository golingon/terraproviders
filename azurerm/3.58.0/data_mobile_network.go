// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetwork "github.com/golingon/terraproviders/azurerm/3.58.0/datamobilenetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetwork creates a new instance of [DataMobileNetwork].
func NewDataMobileNetwork(name string, args DataMobileNetworkArgs) *DataMobileNetwork {
	return &DataMobileNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetwork)(nil)

// DataMobileNetwork represents the Terraform data resource azurerm_mobile_network.
type DataMobileNetwork struct {
	Name string
	Args DataMobileNetworkArgs
}

// DataSource returns the Terraform object type for [DataMobileNetwork].
func (mn *DataMobileNetwork) DataSource() string {
	return "azurerm_mobile_network"
}

// LocalName returns the local name for [DataMobileNetwork].
func (mn *DataMobileNetwork) LocalName() string {
	return mn.Name
}

// Configuration returns the configuration (args) for [DataMobileNetwork].
func (mn *DataMobileNetwork) Configuration() interface{} {
	return mn.Args
}

// Attributes returns the attributes for [DataMobileNetwork].
func (mn *DataMobileNetwork) Attributes() dataMobileNetworkAttributes {
	return dataMobileNetworkAttributes{ref: terra.ReferenceDataResource(mn)}
}

// DataMobileNetworkArgs contains the configurations for azurerm_mobile_network.
type DataMobileNetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datamobilenetwork.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("location"))
}

// MobileCountryCode returns a reference to field mobile_country_code of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) MobileCountryCode() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("mobile_country_code"))
}

// MobileNetworkCode returns a reference to field mobile_network_code of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) MobileNetworkCode() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("mobile_network_code"))
}

// Name returns a reference to field name of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("resource_group_name"))
}

// ServiceKey returns a reference to field service_key of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) ServiceKey() terra.StringValue {
	return terra.ReferenceAsString(mn.ref.Append("service_key"))
}

// Tags returns a reference to field tags of azurerm_mobile_network.
func (mn dataMobileNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mn.ref.Append("tags"))
}

func (mn dataMobileNetworkAttributes) Timeouts() datamobilenetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetwork.TimeoutsAttributes](mn.ref.Append("timeouts"))
}
