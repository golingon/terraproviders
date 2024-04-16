// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_quick_connect

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataQuickConnectConfigAttributes struct {
	ref terra.Reference
}

func (qcc DataQuickConnectConfigAttributes) InternalRef() (terra.Reference, error) {
	return qcc.ref, nil
}

func (qcc DataQuickConnectConfigAttributes) InternalWithRef(ref terra.Reference) DataQuickConnectConfigAttributes {
	return DataQuickConnectConfigAttributes{ref: ref}
}

func (qcc DataQuickConnectConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qcc.ref.InternalTokens()
}

func (qcc DataQuickConnectConfigAttributes) QuickConnectType() terra.StringValue {
	return terra.ReferenceAsString(qcc.ref.Append("quick_connect_type"))
}

func (qcc DataQuickConnectConfigAttributes) PhoneConfig() terra.ListValue[DataQuickConnectConfigPhoneConfigAttributes] {
	return terra.ReferenceAsList[DataQuickConnectConfigPhoneConfigAttributes](qcc.ref.Append("phone_config"))
}

func (qcc DataQuickConnectConfigAttributes) QueueConfig() terra.ListValue[DataQuickConnectConfigQueueConfigAttributes] {
	return terra.ReferenceAsList[DataQuickConnectConfigQueueConfigAttributes](qcc.ref.Append("queue_config"))
}

func (qcc DataQuickConnectConfigAttributes) UserConfig() terra.ListValue[DataQuickConnectConfigUserConfigAttributes] {
	return terra.ReferenceAsList[DataQuickConnectConfigUserConfigAttributes](qcc.ref.Append("user_config"))
}

type DataQuickConnectConfigPhoneConfigAttributes struct {
	ref terra.Reference
}

func (pc DataQuickConnectConfigPhoneConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc DataQuickConnectConfigPhoneConfigAttributes) InternalWithRef(ref terra.Reference) DataQuickConnectConfigPhoneConfigAttributes {
	return DataQuickConnectConfigPhoneConfigAttributes{ref: ref}
}

func (pc DataQuickConnectConfigPhoneConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc DataQuickConnectConfigPhoneConfigAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("phone_number"))
}

type DataQuickConnectConfigQueueConfigAttributes struct {
	ref terra.Reference
}

func (qc DataQuickConnectConfigQueueConfigAttributes) InternalRef() (terra.Reference, error) {
	return qc.ref, nil
}

func (qc DataQuickConnectConfigQueueConfigAttributes) InternalWithRef(ref terra.Reference) DataQuickConnectConfigQueueConfigAttributes {
	return DataQuickConnectConfigQueueConfigAttributes{ref: ref}
}

func (qc DataQuickConnectConfigQueueConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qc.ref.InternalTokens()
}

func (qc DataQuickConnectConfigQueueConfigAttributes) ContactFlowId() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("contact_flow_id"))
}

func (qc DataQuickConnectConfigQueueConfigAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("queue_id"))
}

type DataQuickConnectConfigUserConfigAttributes struct {
	ref terra.Reference
}

func (uc DataQuickConnectConfigUserConfigAttributes) InternalRef() (terra.Reference, error) {
	return uc.ref, nil
}

func (uc DataQuickConnectConfigUserConfigAttributes) InternalWithRef(ref terra.Reference) DataQuickConnectConfigUserConfigAttributes {
	return DataQuickConnectConfigUserConfigAttributes{ref: ref}
}

func (uc DataQuickConnectConfigUserConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return uc.ref.InternalTokens()
}

func (uc DataQuickConnectConfigUserConfigAttributes) ContactFlowId() terra.StringValue {
	return terra.ReferenceAsString(uc.ref.Append("contact_flow_id"))
}

func (uc DataQuickConnectConfigUserConfigAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(uc.ref.Append("user_id"))
}

type DataQuickConnectConfigState struct {
	QuickConnectType string                                   `json:"quick_connect_type"`
	PhoneConfig      []DataQuickConnectConfigPhoneConfigState `json:"phone_config"`
	QueueConfig      []DataQuickConnectConfigQueueConfigState `json:"queue_config"`
	UserConfig       []DataQuickConnectConfigUserConfigState  `json:"user_config"`
}

type DataQuickConnectConfigPhoneConfigState struct {
	PhoneNumber string `json:"phone_number"`
}

type DataQuickConnectConfigQueueConfigState struct {
	ContactFlowId string `json:"contact_flow_id"`
	QueueId       string `json:"queue_id"`
}

type DataQuickConnectConfigUserConfigState struct {
	ContactFlowId string `json:"contact_flow_id"`
	UserId        string `json:"user_id"`
}
