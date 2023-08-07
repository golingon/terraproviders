// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworksimpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/mobilenetworksimpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkSimPolicy creates a new instance of [MobileNetworkSimPolicy].
func NewMobileNetworkSimPolicy(name string, args MobileNetworkSimPolicyArgs) *MobileNetworkSimPolicy {
	return &MobileNetworkSimPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkSimPolicy)(nil)

// MobileNetworkSimPolicy represents the Terraform resource azurerm_mobile_network_sim_policy.
type MobileNetworkSimPolicy struct {
	Name      string
	Args      MobileNetworkSimPolicyArgs
	state     *mobileNetworkSimPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) Type() string {
	return "azurerm_mobile_network_sim_policy"
}

// LocalName returns the local name for [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) LocalName() string {
	return mnsp.Name
}

// Configuration returns the configuration (args) for [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) Configuration() interface{} {
	return mnsp.Args
}

// DependOn is used for other resources to depend on [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(mnsp)
}

// Dependencies returns the list of resources [MobileNetworkSimPolicy] depends_on.
func (mnsp *MobileNetworkSimPolicy) Dependencies() terra.Dependencies {
	return mnsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) LifecycleManagement() *terra.Lifecycle {
	return mnsp.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkSimPolicy].
func (mnsp *MobileNetworkSimPolicy) Attributes() mobileNetworkSimPolicyAttributes {
	return mobileNetworkSimPolicyAttributes{ref: terra.ReferenceResource(mnsp)}
}

// ImportState imports the given attribute values into [MobileNetworkSimPolicy]'s state.
func (mnsp *MobileNetworkSimPolicy) ImportState(av io.Reader) error {
	mnsp.state = &mobileNetworkSimPolicyState{}
	if err := json.NewDecoder(av).Decode(mnsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mnsp.Type(), mnsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkSimPolicy] has state.
func (mnsp *MobileNetworkSimPolicy) State() (*mobileNetworkSimPolicyState, bool) {
	return mnsp.state, mnsp.state != nil
}

// StateMust returns the state for [MobileNetworkSimPolicy]. Panics if the state is nil.
func (mnsp *MobileNetworkSimPolicy) StateMust() *mobileNetworkSimPolicyState {
	if mnsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mnsp.Type(), mnsp.LocalName()))
	}
	return mnsp.state
}

// MobileNetworkSimPolicyArgs contains the configurations for azurerm_mobile_network_sim_policy.
type MobileNetworkSimPolicyArgs struct {
	// DefaultSliceId: string, required
	DefaultSliceId terra.StringValue `hcl:"default_slice_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RatFrequencySelectionPriorityIndex: number, optional
	RatFrequencySelectionPriorityIndex terra.NumberValue `hcl:"rat_frequency_selection_priority_index,attr"`
	// RegistrationTimerInSeconds: number, optional
	RegistrationTimerInSeconds terra.NumberValue `hcl:"registration_timer_in_seconds,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Slice: min=1
	Slice []mobilenetworksimpolicy.Slice `hcl:"slice,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *mobilenetworksimpolicy.Timeouts `hcl:"timeouts,block"`
	// UserEquipmentAggregateMaximumBitRate: required
	UserEquipmentAggregateMaximumBitRate *mobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRate `hcl:"user_equipment_aggregate_maximum_bit_rate,block" validate:"required"`
}
type mobileNetworkSimPolicyAttributes struct {
	ref terra.Reference
}

// DefaultSliceId returns a reference to field default_slice_id of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) DefaultSliceId() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("default_slice_id"))
}

// Id returns a reference to field id of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnsp.ref.Append("name"))
}

// RatFrequencySelectionPriorityIndex returns a reference to field rat_frequency_selection_priority_index of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) RatFrequencySelectionPriorityIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(mnsp.ref.Append("rat_frequency_selection_priority_index"))
}

// RegistrationTimerInSeconds returns a reference to field registration_timer_in_seconds of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) RegistrationTimerInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(mnsp.ref.Append("registration_timer_in_seconds"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_sim_policy.
func (mnsp mobileNetworkSimPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnsp.ref.Append("tags"))
}

func (mnsp mobileNetworkSimPolicyAttributes) Slice() terra.ListValue[mobilenetworksimpolicy.SliceAttributes] {
	return terra.ReferenceAsList[mobilenetworksimpolicy.SliceAttributes](mnsp.ref.Append("slice"))
}

func (mnsp mobileNetworkSimPolicyAttributes) Timeouts() mobilenetworksimpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworksimpolicy.TimeoutsAttributes](mnsp.ref.Append("timeouts"))
}

func (mnsp mobileNetworkSimPolicyAttributes) UserEquipmentAggregateMaximumBitRate() terra.ListValue[mobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRateAttributes] {
	return terra.ReferenceAsList[mobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRateAttributes](mnsp.ref.Append("user_equipment_aggregate_maximum_bit_rate"))
}

type mobileNetworkSimPolicyState struct {
	DefaultSliceId                       string                                                             `json:"default_slice_id"`
	Id                                   string                                                             `json:"id"`
	Location                             string                                                             `json:"location"`
	MobileNetworkId                      string                                                             `json:"mobile_network_id"`
	Name                                 string                                                             `json:"name"`
	RatFrequencySelectionPriorityIndex   float64                                                            `json:"rat_frequency_selection_priority_index"`
	RegistrationTimerInSeconds           float64                                                            `json:"registration_timer_in_seconds"`
	Tags                                 map[string]string                                                  `json:"tags"`
	Slice                                []mobilenetworksimpolicy.SliceState                                `json:"slice"`
	Timeouts                             *mobilenetworksimpolicy.TimeoutsState                              `json:"timeouts"`
	UserEquipmentAggregateMaximumBitRate []mobilenetworksimpolicy.UserEquipmentAggregateMaximumBitRateState `json:"user_equipment_aggregate_maximum_bit_rate"`
}
