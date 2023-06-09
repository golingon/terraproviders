// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamobilenetworkservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PccRule struct {
	// QosPolicy: min=0
	QosPolicy []QosPolicy `hcl:"qos_policy,block" validate:"min=0"`
	// ServiceDataFlowTemplate: min=0
	ServiceDataFlowTemplate []ServiceDataFlowTemplate `hcl:"service_data_flow_template,block" validate:"min=0"`
}

type QosPolicy struct {
	// GuaranteedBitRate: min=0
	GuaranteedBitRate []GuaranteedBitRate `hcl:"guaranteed_bit_rate,block" validate:"min=0"`
	// QosPolicyMaximumBitRate: min=0
	MaximumBitRate []QosPolicyMaximumBitRate `hcl:"maximum_bit_rate,block" validate:"min=0"`
}

type GuaranteedBitRate struct{}

type QosPolicyMaximumBitRate struct{}

type ServiceDataFlowTemplate struct{}

type ServiceQosPolicy struct {
	// ServiceQosPolicyMaximumBitRate: min=0
	MaximumBitRate []ServiceQosPolicyMaximumBitRate `hcl:"maximum_bit_rate,block" validate:"min=0"`
}

type ServiceQosPolicyMaximumBitRate struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type PccRuleAttributes struct {
	ref terra.Reference
}

func (pr PccRuleAttributes) InternalRef() (terra.Reference, error) {
	return pr.ref, nil
}

func (pr PccRuleAttributes) InternalWithRef(ref terra.Reference) PccRuleAttributes {
	return PccRuleAttributes{ref: ref}
}

func (pr PccRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pr.ref.InternalTokens()
}

func (pr PccRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pr.ref.Append("name"))
}

func (pr PccRuleAttributes) Precedence() terra.NumberValue {
	return terra.ReferenceAsNumber(pr.ref.Append("precedence"))
}

func (pr PccRuleAttributes) TrafficControlEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pr.ref.Append("traffic_control_enabled"))
}

func (pr PccRuleAttributes) QosPolicy() terra.ListValue[QosPolicyAttributes] {
	return terra.ReferenceAsList[QosPolicyAttributes](pr.ref.Append("qos_policy"))
}

func (pr PccRuleAttributes) ServiceDataFlowTemplate() terra.ListValue[ServiceDataFlowTemplateAttributes] {
	return terra.ReferenceAsList[ServiceDataFlowTemplateAttributes](pr.ref.Append("service_data_flow_template"))
}

type QosPolicyAttributes struct {
	ref terra.Reference
}

func (qp QosPolicyAttributes) InternalRef() (terra.Reference, error) {
	return qp.ref, nil
}

func (qp QosPolicyAttributes) InternalWithRef(ref terra.Reference) QosPolicyAttributes {
	return QosPolicyAttributes{ref: ref}
}

func (qp QosPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qp.ref.InternalTokens()
}

func (qp QosPolicyAttributes) AllocationAndRetentionPriorityLevel() terra.NumberValue {
	return terra.ReferenceAsNumber(qp.ref.Append("allocation_and_retention_priority_level"))
}

func (qp QosPolicyAttributes) PreemptionCapability() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("preemption_capability"))
}

func (qp QosPolicyAttributes) PreemptionVulnerability() terra.StringValue {
	return terra.ReferenceAsString(qp.ref.Append("preemption_vulnerability"))
}

func (qp QosPolicyAttributes) QosIndicator() terra.NumberValue {
	return terra.ReferenceAsNumber(qp.ref.Append("qos_indicator"))
}

func (qp QosPolicyAttributes) GuaranteedBitRate() terra.ListValue[GuaranteedBitRateAttributes] {
	return terra.ReferenceAsList[GuaranteedBitRateAttributes](qp.ref.Append("guaranteed_bit_rate"))
}

func (qp QosPolicyAttributes) MaximumBitRate() terra.ListValue[QosPolicyMaximumBitRateAttributes] {
	return terra.ReferenceAsList[QosPolicyMaximumBitRateAttributes](qp.ref.Append("maximum_bit_rate"))
}

type GuaranteedBitRateAttributes struct {
	ref terra.Reference
}

func (gbr GuaranteedBitRateAttributes) InternalRef() (terra.Reference, error) {
	return gbr.ref, nil
}

func (gbr GuaranteedBitRateAttributes) InternalWithRef(ref terra.Reference) GuaranteedBitRateAttributes {
	return GuaranteedBitRateAttributes{ref: ref}
}

func (gbr GuaranteedBitRateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gbr.ref.InternalTokens()
}

func (gbr GuaranteedBitRateAttributes) Downlink() terra.StringValue {
	return terra.ReferenceAsString(gbr.ref.Append("downlink"))
}

func (gbr GuaranteedBitRateAttributes) Uplink() terra.StringValue {
	return terra.ReferenceAsString(gbr.ref.Append("uplink"))
}

type QosPolicyMaximumBitRateAttributes struct {
	ref terra.Reference
}

func (mbr QosPolicyMaximumBitRateAttributes) InternalRef() (terra.Reference, error) {
	return mbr.ref, nil
}

func (mbr QosPolicyMaximumBitRateAttributes) InternalWithRef(ref terra.Reference) QosPolicyMaximumBitRateAttributes {
	return QosPolicyMaximumBitRateAttributes{ref: ref}
}

func (mbr QosPolicyMaximumBitRateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mbr.ref.InternalTokens()
}

func (mbr QosPolicyMaximumBitRateAttributes) Downlink() terra.StringValue {
	return terra.ReferenceAsString(mbr.ref.Append("downlink"))
}

func (mbr QosPolicyMaximumBitRateAttributes) Uplink() terra.StringValue {
	return terra.ReferenceAsString(mbr.ref.Append("uplink"))
}

type ServiceDataFlowTemplateAttributes struct {
	ref terra.Reference
}

func (sdft ServiceDataFlowTemplateAttributes) InternalRef() (terra.Reference, error) {
	return sdft.ref, nil
}

func (sdft ServiceDataFlowTemplateAttributes) InternalWithRef(ref terra.Reference) ServiceDataFlowTemplateAttributes {
	return ServiceDataFlowTemplateAttributes{ref: ref}
}

func (sdft ServiceDataFlowTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sdft.ref.InternalTokens()
}

func (sdft ServiceDataFlowTemplateAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(sdft.ref.Append("direction"))
}

func (sdft ServiceDataFlowTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdft.ref.Append("name"))
}

func (sdft ServiceDataFlowTemplateAttributes) Ports() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdft.ref.Append("ports"))
}

func (sdft ServiceDataFlowTemplateAttributes) Protocol() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdft.ref.Append("protocol"))
}

func (sdft ServiceDataFlowTemplateAttributes) RemoteIpList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sdft.ref.Append("remote_ip_list"))
}

type ServiceQosPolicyAttributes struct {
	ref terra.Reference
}

func (sqp ServiceQosPolicyAttributes) InternalRef() (terra.Reference, error) {
	return sqp.ref, nil
}

func (sqp ServiceQosPolicyAttributes) InternalWithRef(ref terra.Reference) ServiceQosPolicyAttributes {
	return ServiceQosPolicyAttributes{ref: ref}
}

func (sqp ServiceQosPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sqp.ref.InternalTokens()
}

func (sqp ServiceQosPolicyAttributes) AllocationAndRetentionPriorityLevel() terra.NumberValue {
	return terra.ReferenceAsNumber(sqp.ref.Append("allocation_and_retention_priority_level"))
}

func (sqp ServiceQosPolicyAttributes) PreemptionCapability() terra.StringValue {
	return terra.ReferenceAsString(sqp.ref.Append("preemption_capability"))
}

func (sqp ServiceQosPolicyAttributes) PreemptionVulnerability() terra.StringValue {
	return terra.ReferenceAsString(sqp.ref.Append("preemption_vulnerability"))
}

func (sqp ServiceQosPolicyAttributes) QosIndicator() terra.NumberValue {
	return terra.ReferenceAsNumber(sqp.ref.Append("qos_indicator"))
}

func (sqp ServiceQosPolicyAttributes) MaximumBitRate() terra.ListValue[ServiceQosPolicyMaximumBitRateAttributes] {
	return terra.ReferenceAsList[ServiceQosPolicyMaximumBitRateAttributes](sqp.ref.Append("maximum_bit_rate"))
}

type ServiceQosPolicyMaximumBitRateAttributes struct {
	ref terra.Reference
}

func (mbr ServiceQosPolicyMaximumBitRateAttributes) InternalRef() (terra.Reference, error) {
	return mbr.ref, nil
}

func (mbr ServiceQosPolicyMaximumBitRateAttributes) InternalWithRef(ref terra.Reference) ServiceQosPolicyMaximumBitRateAttributes {
	return ServiceQosPolicyMaximumBitRateAttributes{ref: ref}
}

func (mbr ServiceQosPolicyMaximumBitRateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mbr.ref.InternalTokens()
}

func (mbr ServiceQosPolicyMaximumBitRateAttributes) Downlink() terra.StringValue {
	return terra.ReferenceAsString(mbr.ref.Append("downlink"))
}

func (mbr ServiceQosPolicyMaximumBitRateAttributes) Uplink() terra.StringValue {
	return terra.ReferenceAsString(mbr.ref.Append("uplink"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type PccRuleState struct {
	Name                    string                         `json:"name"`
	Precedence              float64                        `json:"precedence"`
	TrafficControlEnabled   bool                           `json:"traffic_control_enabled"`
	QosPolicy               []QosPolicyState               `json:"qos_policy"`
	ServiceDataFlowTemplate []ServiceDataFlowTemplateState `json:"service_data_flow_template"`
}

type QosPolicyState struct {
	AllocationAndRetentionPriorityLevel float64                        `json:"allocation_and_retention_priority_level"`
	PreemptionCapability                string                         `json:"preemption_capability"`
	PreemptionVulnerability             string                         `json:"preemption_vulnerability"`
	QosIndicator                        float64                        `json:"qos_indicator"`
	GuaranteedBitRate                   []GuaranteedBitRateState       `json:"guaranteed_bit_rate"`
	MaximumBitRate                      []QosPolicyMaximumBitRateState `json:"maximum_bit_rate"`
}

type GuaranteedBitRateState struct {
	Downlink string `json:"downlink"`
	Uplink   string `json:"uplink"`
}

type QosPolicyMaximumBitRateState struct {
	Downlink string `json:"downlink"`
	Uplink   string `json:"uplink"`
}

type ServiceDataFlowTemplateState struct {
	Direction    string   `json:"direction"`
	Name         string   `json:"name"`
	Ports        []string `json:"ports"`
	Protocol     []string `json:"protocol"`
	RemoteIpList []string `json:"remote_ip_list"`
}

type ServiceQosPolicyState struct {
	AllocationAndRetentionPriorityLevel float64                               `json:"allocation_and_retention_priority_level"`
	PreemptionCapability                string                                `json:"preemption_capability"`
	PreemptionVulnerability             string                                `json:"preemption_vulnerability"`
	QosIndicator                        float64                               `json:"qos_indicator"`
	MaximumBitRate                      []ServiceQosPolicyMaximumBitRateState `json:"maximum_bit_rate"`
}

type ServiceQosPolicyMaximumBitRateState struct {
	Downlink string `json:"downlink"`
	Uplink   string `json:"uplink"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
