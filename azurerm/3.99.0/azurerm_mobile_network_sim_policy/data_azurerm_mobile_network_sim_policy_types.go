// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mobile_network_sim_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataSliceAttributes struct {
	ref terra.Reference
}

func (s DataSliceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataSliceAttributes) InternalWithRef(ref terra.Reference) DataSliceAttributes {
	return DataSliceAttributes{ref: ref}
}

func (s DataSliceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataSliceAttributes) DefaultDataNetworkId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("default_data_network_id"))
}

func (s DataSliceAttributes) SliceId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("slice_id"))
}

func (s DataSliceAttributes) DataNetwork() terra.ListValue[DataSliceDataNetworkAttributes] {
	return terra.ReferenceAsList[DataSliceDataNetworkAttributes](s.ref.Append("data_network"))
}

type DataSliceDataNetworkAttributes struct {
	ref terra.Reference
}

func (dn DataSliceDataNetworkAttributes) InternalRef() (terra.Reference, error) {
	return dn.ref, nil
}

func (dn DataSliceDataNetworkAttributes) InternalWithRef(ref terra.Reference) DataSliceDataNetworkAttributes {
	return DataSliceDataNetworkAttributes{ref: ref}
}

func (dn DataSliceDataNetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dn.ref.InternalTokens()
}

func (dn DataSliceDataNetworkAttributes) AdditionalAllowedSessionTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dn.ref.Append("additional_allowed_session_types"))
}

func (dn DataSliceDataNetworkAttributes) AllocationAndRetentionPriorityLevel() terra.NumberValue {
	return terra.ReferenceAsNumber(dn.ref.Append("allocation_and_retention_priority_level"))
}

func (dn DataSliceDataNetworkAttributes) AllowedServicesIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dn.ref.Append("allowed_services_ids"))
}

func (dn DataSliceDataNetworkAttributes) DataNetworkId() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("data_network_id"))
}

func (dn DataSliceDataNetworkAttributes) DefaultSessionType() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("default_session_type"))
}

func (dn DataSliceDataNetworkAttributes) MaxBufferedPackets() terra.NumberValue {
	return terra.ReferenceAsNumber(dn.ref.Append("max_buffered_packets"))
}

func (dn DataSliceDataNetworkAttributes) PreemptionCapability() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("preemption_capability"))
}

func (dn DataSliceDataNetworkAttributes) PreemptionVulnerability() terra.StringValue {
	return terra.ReferenceAsString(dn.ref.Append("preemption_vulnerability"))
}

func (dn DataSliceDataNetworkAttributes) QosIndicator() terra.NumberValue {
	return terra.ReferenceAsNumber(dn.ref.Append("qos_indicator"))
}

func (dn DataSliceDataNetworkAttributes) SessionAggregateMaximumBitRate() terra.ListValue[DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes] {
	return terra.ReferenceAsList[DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes](dn.ref.Append("session_aggregate_maximum_bit_rate"))
}

type DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes struct {
	ref terra.Reference
}

func (sambr DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes) InternalRef() (terra.Reference, error) {
	return sambr.ref, nil
}

func (sambr DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes) InternalWithRef(ref terra.Reference) DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes {
	return DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes{ref: ref}
}

func (sambr DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sambr.ref.InternalTokens()
}

func (sambr DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes) Downlink() terra.StringValue {
	return terra.ReferenceAsString(sambr.ref.Append("downlink"))
}

func (sambr DataSliceDataNetworkSessionAggregateMaximumBitRateAttributes) Uplink() terra.StringValue {
	return terra.ReferenceAsString(sambr.ref.Append("uplink"))
}

type DataUserEquipmentAggregateMaximumBitRateAttributes struct {
	ref terra.Reference
}

func (ueambr DataUserEquipmentAggregateMaximumBitRateAttributes) InternalRef() (terra.Reference, error) {
	return ueambr.ref, nil
}

func (ueambr DataUserEquipmentAggregateMaximumBitRateAttributes) InternalWithRef(ref terra.Reference) DataUserEquipmentAggregateMaximumBitRateAttributes {
	return DataUserEquipmentAggregateMaximumBitRateAttributes{ref: ref}
}

func (ueambr DataUserEquipmentAggregateMaximumBitRateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ueambr.ref.InternalTokens()
}

func (ueambr DataUserEquipmentAggregateMaximumBitRateAttributes) Downlink() terra.StringValue {
	return terra.ReferenceAsString(ueambr.ref.Append("downlink"))
}

func (ueambr DataUserEquipmentAggregateMaximumBitRateAttributes) Uplink() terra.StringValue {
	return terra.ReferenceAsString(ueambr.ref.Append("uplink"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataSliceState struct {
	DefaultDataNetworkId string                      `json:"default_data_network_id"`
	SliceId              string                      `json:"slice_id"`
	DataNetwork          []DataSliceDataNetworkState `json:"data_network"`
}

type DataSliceDataNetworkState struct {
	AdditionalAllowedSessionTypes       []string                                                  `json:"additional_allowed_session_types"`
	AllocationAndRetentionPriorityLevel float64                                                   `json:"allocation_and_retention_priority_level"`
	AllowedServicesIds                  []string                                                  `json:"allowed_services_ids"`
	DataNetworkId                       string                                                    `json:"data_network_id"`
	DefaultSessionType                  string                                                    `json:"default_session_type"`
	MaxBufferedPackets                  float64                                                   `json:"max_buffered_packets"`
	PreemptionCapability                string                                                    `json:"preemption_capability"`
	PreemptionVulnerability             string                                                    `json:"preemption_vulnerability"`
	QosIndicator                        float64                                                   `json:"qos_indicator"`
	SessionAggregateMaximumBitRate      []DataSliceDataNetworkSessionAggregateMaximumBitRateState `json:"session_aggregate_maximum_bit_rate"`
}

type DataSliceDataNetworkSessionAggregateMaximumBitRateState struct {
	Downlink string `json:"downlink"`
	Uplink   string `json:"uplink"`
}

type DataUserEquipmentAggregateMaximumBitRateState struct {
	Downlink string `json:"downlink"`
	Uplink   string `json:"uplink"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
