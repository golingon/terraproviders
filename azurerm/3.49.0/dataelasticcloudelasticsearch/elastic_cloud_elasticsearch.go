// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataelasticcloudelasticsearch

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Logs struct {
	// FilteringTag: min=0
	FilteringTag []FilteringTag `hcl:"filtering_tag,block" validate:"min=0"`
}

type FilteringTag struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type LogsAttributes struct {
	ref terra.Reference
}

func (l LogsAttributes) InternalRef() terra.Reference {
	return l.ref
}

func (l LogsAttributes) InternalWithRef(ref terra.Reference) LogsAttributes {
	return LogsAttributes{ref: ref}
}

func (l LogsAttributes) InternalTokens() hclwrite.Tokens {
	return l.ref.InternalTokens()
}

func (l LogsAttributes) SendActivityLogs() terra.BoolValue {
	return terra.ReferenceBool(l.ref.Append("send_activity_logs"))
}

func (l LogsAttributes) SendAzureadLogs() terra.BoolValue {
	return terra.ReferenceBool(l.ref.Append("send_azuread_logs"))
}

func (l LogsAttributes) SendSubscriptionLogs() terra.BoolValue {
	return terra.ReferenceBool(l.ref.Append("send_subscription_logs"))
}

func (l LogsAttributes) FilteringTag() terra.ListValue[FilteringTagAttributes] {
	return terra.ReferenceList[FilteringTagAttributes](l.ref.Append("filtering_tag"))
}

type FilteringTagAttributes struct {
	ref terra.Reference
}

func (ft FilteringTagAttributes) InternalRef() terra.Reference {
	return ft.ref
}

func (ft FilteringTagAttributes) InternalWithRef(ref terra.Reference) FilteringTagAttributes {
	return FilteringTagAttributes{ref: ref}
}

func (ft FilteringTagAttributes) InternalTokens() hclwrite.Tokens {
	return ft.ref.InternalTokens()
}

func (ft FilteringTagAttributes) Action() terra.StringValue {
	return terra.ReferenceString(ft.ref.Append("action"))
}

func (ft FilteringTagAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ft.ref.Append("name"))
}

func (ft FilteringTagAttributes) Value() terra.StringValue {
	return terra.ReferenceString(ft.ref.Append("value"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type LogsState struct {
	SendActivityLogs     bool                `json:"send_activity_logs"`
	SendAzureadLogs      bool                `json:"send_azuread_logs"`
	SendSubscriptionLogs bool                `json:"send_subscription_logs"`
	FilteringTag         []FilteringTagState `json:"filtering_tag"`
}

type FilteringTagState struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
