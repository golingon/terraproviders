// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamobilenetworksimpolicy "github.com/golingon/terraproviders/azurerm/3.52.0/datamobilenetworksimpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMobileNetworkSimPolicy creates a new instance of [DataMobileNetworkSimPolicy].
func NewDataMobileNetworkSimPolicy(name string, args DataMobileNetworkSimPolicyArgs) *DataMobileNetworkSimPolicy {
	return &DataMobileNetworkSimPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMobileNetworkSimPolicy)(nil)

// DataMobileNetworkSimPolicy represents the Terraform data resource azurerm_mobile_network_sim_policy.
type DataMobileNetworkSimPolicy struct {
	Name string
	Args DataMobileNetworkSimPolicyArgs
}

// DataSource returns the Terraform object type for [DataMobileNetworkSimPolicy].
func (mnsp *DataMobileNetworkSimPolicy) DataSource() string {
	return "azurerm_mobile_network_sim_policy"
}

// LocalName returns the local name for [DataMobileNetworkSimPolicy].
func (mnsp *DataMobileNetworkSimPolicy) LocalName() string {
	return mnsp.Name
}

// Configuration returns the configuration (args) for [DataMobileNetworkSimPolicy].
func (mnsp *DataMobileNetworkSimPolicy) Configuration() interface{} {
	return mnsp.Args
}

// Attributes returns the attributes for [DataMobileNetworkSimPolicy].
func (mnsp *DataMobileNetworkSimPolicy) Attributes() dataMobileNetworkSimPolicyAttributes {
	return dataMobileNetworkSimPolicyAttributes{ref: terra.ReferenceDataResource(mnsp)}
}

// DataMobileNetworkSimPolicyArgs contains the configurations for azurerm_mobile_network_sim_policy.
type DataMobileNetworkSimPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Slice: min=0
	Slice []datamobilenetworksimpolicy.Slice `hcl:"slice,block" validate:"min=0"`
	// UserEquipmentAggregateMaximumBitRate: min=0
	UserEquipmentAggregateMaximumBitRate []datamobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRate `hcl:"user_equipment_aggregate_maximum_bit_rate,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamobilenetworksimpolicy.Timeouts `hcl:"timeouts,block"`
}
type dataMobileNetworkSimPolicyAttributes struct {
	ref terra.Reference
}

// DefaultSliceId returns a reference to field default_slice_id of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) DefaultSliceId() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("default_slice_id"))
}

// Id returns a reference to field id of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("name"))
}

// RatFrequencySelectionPriorityIndex returns a reference to field rat_frequency_selection_priority_index of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) RatFrequencySelectionPriorityIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(mnsp.ref.Append("rat_frequency_selection_priority_index"))
}

// RegistrationTimerInSeconds returns a reference to field registration_timer_in_seconds of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) RegistrationTimerInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mnsp.ref.Append("registration_timer_in_seconds"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_sim_policy.
func (mnsp dataMobileNetworkSimPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnsp.ref.Append("tags"))
}

func (mnsp dataMobileNetworkSimPolicyAttributes) Slice() terra.ListValue[datamobilenetworksimpolicy.SliceAttributes] {
	return terra.ReferenceAsList[datamobilenetworksimpolicy.SliceAttributes](mnsp.ref.Append("slice"))
}

func (mnsp dataMobileNetworkSimPolicyAttributes) UserEquipmentAggregateMaximumBitRate() terra.ListValue[datamobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRateAttributes] {
	return terra.ReferenceAsList[datamobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRateAttributes](mnsp.ref.Append("user_equipment_aggregate_maximum_bit_rate"))
}

func (mnsp dataMobileNetworkSimPolicyAttributes) Timeouts() datamobilenetworksimpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamobilenetworksimpolicy.TimeoutsAttributes](mnsp.ref.Append("timeouts"))
}
