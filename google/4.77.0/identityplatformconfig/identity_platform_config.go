// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package identityplatformconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BlockingFunctions struct {
	// ForwardInboundCredentials: optional
	ForwardInboundCredentials *ForwardInboundCredentials `hcl:"forward_inbound_credentials,block"`
	// Triggers: min=1
	Triggers []Triggers `hcl:"triggers,block" validate:"min=1"`
}

type ForwardInboundCredentials struct {
	// AccessToken: bool, optional
	AccessToken terra.BoolValue `hcl:"access_token,attr"`
	// IdToken: bool, optional
	IdToken terra.BoolValue `hcl:"id_token,attr"`
	// RefreshToken: bool, optional
	RefreshToken terra.BoolValue `hcl:"refresh_token,attr"`
}

type Triggers struct {
	// EventType: string, required
	EventType terra.StringValue `hcl:"event_type,attr" validate:"required"`
	// FunctionUri: string, required
	FunctionUri terra.StringValue `hcl:"function_uri,attr" validate:"required"`
}

type Quota struct {
	// SignUpQuotaConfig: optional
	SignUpQuotaConfig *SignUpQuotaConfig `hcl:"sign_up_quota_config,block"`
}

type SignUpQuotaConfig struct {
	// Quota: number, optional
	Quota terra.NumberValue `hcl:"quota,attr"`
	// QuotaDuration: string, optional
	QuotaDuration terra.StringValue `hcl:"quota_duration,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type BlockingFunctionsAttributes struct {
	ref terra.Reference
}

func (bf BlockingFunctionsAttributes) InternalRef() (terra.Reference, error) {
	return bf.ref, nil
}

func (bf BlockingFunctionsAttributes) InternalWithRef(ref terra.Reference) BlockingFunctionsAttributes {
	return BlockingFunctionsAttributes{ref: ref}
}

func (bf BlockingFunctionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bf.ref.InternalTokens()
}

func (bf BlockingFunctionsAttributes) ForwardInboundCredentials() terra.ListValue[ForwardInboundCredentialsAttributes] {
	return terra.ReferenceAsList[ForwardInboundCredentialsAttributes](bf.ref.Append("forward_inbound_credentials"))
}

func (bf BlockingFunctionsAttributes) Triggers() terra.SetValue[TriggersAttributes] {
	return terra.ReferenceAsSet[TriggersAttributes](bf.ref.Append("triggers"))
}

type ForwardInboundCredentialsAttributes struct {
	ref terra.Reference
}

func (fic ForwardInboundCredentialsAttributes) InternalRef() (terra.Reference, error) {
	return fic.ref, nil
}

func (fic ForwardInboundCredentialsAttributes) InternalWithRef(ref terra.Reference) ForwardInboundCredentialsAttributes {
	return ForwardInboundCredentialsAttributes{ref: ref}
}

func (fic ForwardInboundCredentialsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fic.ref.InternalTokens()
}

func (fic ForwardInboundCredentialsAttributes) AccessToken() terra.BoolValue {
	return terra.ReferenceAsBool(fic.ref.Append("access_token"))
}

func (fic ForwardInboundCredentialsAttributes) IdToken() terra.BoolValue {
	return terra.ReferenceAsBool(fic.ref.Append("id_token"))
}

func (fic ForwardInboundCredentialsAttributes) RefreshToken() terra.BoolValue {
	return terra.ReferenceAsBool(fic.ref.Append("refresh_token"))
}

type TriggersAttributes struct {
	ref terra.Reference
}

func (t TriggersAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TriggersAttributes) InternalWithRef(ref terra.Reference) TriggersAttributes {
	return TriggersAttributes{ref: ref}
}

func (t TriggersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TriggersAttributes) EventType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("event_type"))
}

func (t TriggersAttributes) FunctionUri() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("function_uri"))
}

func (t TriggersAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update_time"))
}

type QuotaAttributes struct {
	ref terra.Reference
}

func (q QuotaAttributes) InternalRef() (terra.Reference, error) {
	return q.ref, nil
}

func (q QuotaAttributes) InternalWithRef(ref terra.Reference) QuotaAttributes {
	return QuotaAttributes{ref: ref}
}

func (q QuotaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return q.ref.InternalTokens()
}

func (q QuotaAttributes) SignUpQuotaConfig() terra.ListValue[SignUpQuotaConfigAttributes] {
	return terra.ReferenceAsList[SignUpQuotaConfigAttributes](q.ref.Append("sign_up_quota_config"))
}

type SignUpQuotaConfigAttributes struct {
	ref terra.Reference
}

func (suqc SignUpQuotaConfigAttributes) InternalRef() (terra.Reference, error) {
	return suqc.ref, nil
}

func (suqc SignUpQuotaConfigAttributes) InternalWithRef(ref terra.Reference) SignUpQuotaConfigAttributes {
	return SignUpQuotaConfigAttributes{ref: ref}
}

func (suqc SignUpQuotaConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return suqc.ref.InternalTokens()
}

func (suqc SignUpQuotaConfigAttributes) Quota() terra.NumberValue {
	return terra.ReferenceAsNumber(suqc.ref.Append("quota"))
}

func (suqc SignUpQuotaConfigAttributes) QuotaDuration() terra.StringValue {
	return terra.ReferenceAsString(suqc.ref.Append("quota_duration"))
}

func (suqc SignUpQuotaConfigAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(suqc.ref.Append("start_time"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type BlockingFunctionsState struct {
	ForwardInboundCredentials []ForwardInboundCredentialsState `json:"forward_inbound_credentials"`
	Triggers                  []TriggersState                  `json:"triggers"`
}

type ForwardInboundCredentialsState struct {
	AccessToken  bool `json:"access_token"`
	IdToken      bool `json:"id_token"`
	RefreshToken bool `json:"refresh_token"`
}

type TriggersState struct {
	EventType   string `json:"event_type"`
	FunctionUri string `json:"function_uri"`
	UpdateTime  string `json:"update_time"`
}

type QuotaState struct {
	SignUpQuotaConfig []SignUpQuotaConfigState `json:"sign_up_quota_config"`
}

type SignUpQuotaConfigState struct {
	Quota         float64 `json:"quota"`
	QuotaDuration string  `json:"quota_duration"`
	StartTime     string  `json:"start_time"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}