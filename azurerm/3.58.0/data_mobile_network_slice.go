// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetworkslice "github.com/golingon/terraproviders/azurerm/3.58.0/datamobilenetworkslice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetworkSlice creates a new instance of [DataMobileNetworkSlice].
func NewDataMobileNetworkSlice(name string, args DataMobileNetworkSliceArgs) *DataMobileNetworkSlice {
	return &DataMobileNetworkSlice{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkSlice)(nil)

// DataMobileNetworkSlice represents the Terraform data resource azurerm_mobile_network_slice.
type DataMobileNetworkSlice struct {
	Name string
	Args DataMobileNetworkSliceArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkSlice].
func (mns *DataMobileNetworkSlice) DataSource() string {
	return "azurerm_mobile_network_slice"
}

// LocalName returns the local name for [DataMobileNetworkSlice].
func (mns *DataMobileNetworkSlice) LocalName() string {
	return mns.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkSlice].
func (mns *DataMobileNetworkSlice) Configuration() interface{} {
	return mns.Args
}

// Attributes returns the attributes for [DataMobileNetworkSlice].
func (mns *DataMobileNetworkSlice) Attributes() dataMobileNetworkSliceAttributes {
	return dataMobileNetworkSliceAttributes{ref: terra.ReferenceDataResource(mns)}
}

// DataMobileNetworkSliceArgs contains the configurations for azurerm_mobile_network_slice.
type DataMobileNetworkSliceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SingleNetworkSliceSelectionAssistanceInformation: min=0
	SingleNetworkSliceSelectionAssistanceInformation []datamobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformation `hcl:"single_network_slice_selection_assistance_information,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamobilenetworkslice.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkSliceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_slice.
func (mns dataMobileNetworkSliceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mns.ref.Append("tags"))
}

func (mns dataMobileNetworkSliceAttributes) SingleNetworkSliceSelectionAssistanceInformation() terra.ListValue[datamobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformationAttributes] {
	return terra.ReferenceAsList[datamobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformationAttributes](mns.ref.Append("single_network_slice_selection_assistance_information"))
}

func (mns dataMobileNetworkSliceAttributes) Timeouts() datamobilenetworkslice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworkslice.TimeoutsAttributes](mns.ref.Append("timeouts"))
}
