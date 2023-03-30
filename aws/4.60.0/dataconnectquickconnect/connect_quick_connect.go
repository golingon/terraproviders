// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataconnectquickconnect

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type QuickConnectConfig struct {
	// PhoneConfig: min=0
	PhoneConfig []PhoneConfig `hcl:"phone_config,block" validate:"min=0"`
	// QueueConfig: min=0
	QueueConfig []QueueConfig `hcl:"queue_config,block" validate:"min=0"`
	// UserConfig: min=0
	UserConfig []UserConfig `hcl:"user_config,block" validate:"min=0"`
}

type PhoneConfig struct{}

type QueueConfig struct{}

type UserConfig struct{}

type QuickConnectConfigAttributes struct {
	ref terra.Reference
}

func (qcc QuickConnectConfigAttributes) InternalRef() terra.Reference {
	return qcc.ref
}

func (qcc QuickConnectConfigAttributes) InternalWithRef(ref terra.Reference) QuickConnectConfigAttributes {
	return QuickConnectConfigAttributes{ref: ref}
}

func (qcc QuickConnectConfigAttributes) InternalTokens() hclwrite.Tokens {
	return qcc.ref.InternalTokens()
}

func (qcc QuickConnectConfigAttributes) QuickConnectType() terra.StringValue {
	return terra.ReferenceString(qcc.ref.Append("quick_connect_type"))
}

func (qcc QuickConnectConfigAttributes) PhoneConfig() terra.ListValue[PhoneConfigAttributes] {
	return terra.ReferenceList[PhoneConfigAttributes](qcc.ref.Append("phone_config"))
}

func (qcc QuickConnectConfigAttributes) QueueConfig() terra.ListValue[QueueConfigAttributes] {
	return terra.ReferenceList[QueueConfigAttributes](qcc.ref.Append("queue_config"))
}

func (qcc QuickConnectConfigAttributes) UserConfig() terra.ListValue[UserConfigAttributes] {
	return terra.ReferenceList[UserConfigAttributes](qcc.ref.Append("user_config"))
}

type PhoneConfigAttributes struct {
	ref terra.Reference
}

func (pc PhoneConfigAttributes) InternalRef() terra.Reference {
	return pc.ref
}

func (pc PhoneConfigAttributes) InternalWithRef(ref terra.Reference) PhoneConfigAttributes {
	return PhoneConfigAttributes{ref: ref}
}

func (pc PhoneConfigAttributes) InternalTokens() hclwrite.Tokens {
	return pc.ref.InternalTokens()
}

func (pc PhoneConfigAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(pc.ref.Append("phone_number"))
}

type QueueConfigAttributes struct {
	ref terra.Reference
}

func (qc QueueConfigAttributes) InternalRef() terra.Reference {
	return qc.ref
}

func (qc QueueConfigAttributes) InternalWithRef(ref terra.Reference) QueueConfigAttributes {
	return QueueConfigAttributes{ref: ref}
}

func (qc QueueConfigAttributes) InternalTokens() hclwrite.Tokens {
	return qc.ref.InternalTokens()
}

func (qc QueueConfigAttributes) ContactFlowId() terra.StringValue {
	return terra.ReferenceString(qc.ref.Append("contact_flow_id"))
}

func (qc QueueConfigAttributes) QueueId() terra.StringValue {
	return terra.ReferenceString(qc.ref.Append("queue_id"))
}

type UserConfigAttributes struct {
	ref terra.Reference
}

func (uc UserConfigAttributes) InternalRef() terra.Reference {
	return uc.ref
}

func (uc UserConfigAttributes) InternalWithRef(ref terra.Reference) UserConfigAttributes {
	return UserConfigAttributes{ref: ref}
}

func (uc UserConfigAttributes) InternalTokens() hclwrite.Tokens {
	return uc.ref.InternalTokens()
}

func (uc UserConfigAttributes) ContactFlowId() terra.StringValue {
	return terra.ReferenceString(uc.ref.Append("contact_flow_id"))
}

func (uc UserConfigAttributes) UserId() terra.StringValue {
	return terra.ReferenceString(uc.ref.Append("user_id"))
}

type QuickConnectConfigState struct {
	QuickConnectType string             `json:"quick_connect_type"`
	PhoneConfig      []PhoneConfigState `json:"phone_config"`
	QueueConfig      []QueueConfigState `json:"queue_config"`
	UserConfig       []UserConfigState  `json:"user_config"`
}

type PhoneConfigState struct {
	PhoneNumber string `json:"phone_number"`
}

type QueueConfigState struct {
	ContactFlowId string `json:"contact_flow_id"`
	QueueId       string `json:"queue_id"`
}

type UserConfigState struct {
	ContactFlowId string `json:"contact_flow_id"`
	UserId        string `json:"user_id"`
}
