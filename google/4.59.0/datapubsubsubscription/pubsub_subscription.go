// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapubsubsubscription

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BigqueryConfig struct{}

type DeadLetterPolicy struct{}

type ExpirationPolicy struct{}

type PushConfig struct {
	// OidcToken: min=0
	OidcToken []OidcToken `hcl:"oidc_token,block" validate:"min=0"`
}

type OidcToken struct{}

type RetryPolicy struct{}

type BigqueryConfigAttributes struct {
	ref terra.Reference
}

func (bc BigqueryConfigAttributes) InternalRef() terra.Reference {
	return bc.ref
}

func (bc BigqueryConfigAttributes) InternalWithRef(ref terra.Reference) BigqueryConfigAttributes {
	return BigqueryConfigAttributes{ref: ref}
}

func (bc BigqueryConfigAttributes) InternalTokens() hclwrite.Tokens {
	return bc.ref.InternalTokens()
}

func (bc BigqueryConfigAttributes) DropUnknownFields() terra.BoolValue {
	return terra.ReferenceBool(bc.ref.Append("drop_unknown_fields"))
}

func (bc BigqueryConfigAttributes) Table() terra.StringValue {
	return terra.ReferenceString(bc.ref.Append("table"))
}

func (bc BigqueryConfigAttributes) UseTopicSchema() terra.BoolValue {
	return terra.ReferenceBool(bc.ref.Append("use_topic_schema"))
}

func (bc BigqueryConfigAttributes) WriteMetadata() terra.BoolValue {
	return terra.ReferenceBool(bc.ref.Append("write_metadata"))
}

type DeadLetterPolicyAttributes struct {
	ref terra.Reference
}

func (dlp DeadLetterPolicyAttributes) InternalRef() terra.Reference {
	return dlp.ref
}

func (dlp DeadLetterPolicyAttributes) InternalWithRef(ref terra.Reference) DeadLetterPolicyAttributes {
	return DeadLetterPolicyAttributes{ref: ref}
}

func (dlp DeadLetterPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return dlp.ref.InternalTokens()
}

func (dlp DeadLetterPolicyAttributes) DeadLetterTopic() terra.StringValue {
	return terra.ReferenceString(dlp.ref.Append("dead_letter_topic"))
}

func (dlp DeadLetterPolicyAttributes) MaxDeliveryAttempts() terra.NumberValue {
	return terra.ReferenceNumber(dlp.ref.Append("max_delivery_attempts"))
}

type ExpirationPolicyAttributes struct {
	ref terra.Reference
}

func (ep ExpirationPolicyAttributes) InternalRef() terra.Reference {
	return ep.ref
}

func (ep ExpirationPolicyAttributes) InternalWithRef(ref terra.Reference) ExpirationPolicyAttributes {
	return ExpirationPolicyAttributes{ref: ref}
}

func (ep ExpirationPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return ep.ref.InternalTokens()
}

func (ep ExpirationPolicyAttributes) Ttl() terra.StringValue {
	return terra.ReferenceString(ep.ref.Append("ttl"))
}

type PushConfigAttributes struct {
	ref terra.Reference
}

func (pc PushConfigAttributes) InternalRef() terra.Reference {
	return pc.ref
}

func (pc PushConfigAttributes) InternalWithRef(ref terra.Reference) PushConfigAttributes {
	return PushConfigAttributes{ref: ref}
}

func (pc PushConfigAttributes) InternalTokens() hclwrite.Tokens {
	return pc.ref.InternalTokens()
}

func (pc PushConfigAttributes) Attributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pc.ref.Append("attributes"))
}

func (pc PushConfigAttributes) PushEndpoint() terra.StringValue {
	return terra.ReferenceString(pc.ref.Append("push_endpoint"))
}

func (pc PushConfigAttributes) OidcToken() terra.ListValue[OidcTokenAttributes] {
	return terra.ReferenceList[OidcTokenAttributes](pc.ref.Append("oidc_token"))
}

type OidcTokenAttributes struct {
	ref terra.Reference
}

func (ot OidcTokenAttributes) InternalRef() terra.Reference {
	return ot.ref
}

func (ot OidcTokenAttributes) InternalWithRef(ref terra.Reference) OidcTokenAttributes {
	return OidcTokenAttributes{ref: ref}
}

func (ot OidcTokenAttributes) InternalTokens() hclwrite.Tokens {
	return ot.ref.InternalTokens()
}

func (ot OidcTokenAttributes) Audience() terra.StringValue {
	return terra.ReferenceString(ot.ref.Append("audience"))
}

func (ot OidcTokenAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceString(ot.ref.Append("service_account_email"))
}

type RetryPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetryPolicyAttributes) InternalRef() terra.Reference {
	return rp.ref
}

func (rp RetryPolicyAttributes) InternalWithRef(ref terra.Reference) RetryPolicyAttributes {
	return RetryPolicyAttributes{ref: ref}
}

func (rp RetryPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return rp.ref.InternalTokens()
}

func (rp RetryPolicyAttributes) MaximumBackoff() terra.StringValue {
	return terra.ReferenceString(rp.ref.Append("maximum_backoff"))
}

func (rp RetryPolicyAttributes) MinimumBackoff() terra.StringValue {
	return terra.ReferenceString(rp.ref.Append("minimum_backoff"))
}

type BigqueryConfigState struct {
	DropUnknownFields bool   `json:"drop_unknown_fields"`
	Table             string `json:"table"`
	UseTopicSchema    bool   `json:"use_topic_schema"`
	WriteMetadata     bool   `json:"write_metadata"`
}

type DeadLetterPolicyState struct {
	DeadLetterTopic     string  `json:"dead_letter_topic"`
	MaxDeliveryAttempts float64 `json:"max_delivery_attempts"`
}

type ExpirationPolicyState struct {
	Ttl string `json:"ttl"`
}

type PushConfigState struct {
	Attributes   map[string]string `json:"attributes"`
	PushEndpoint string            `json:"push_endpoint"`
	OidcToken    []OidcTokenState  `json:"oidc_token"`
}

type OidcTokenState struct {
	Audience            string `json:"audience"`
	ServiceAccountEmail string `json:"service_account_email"`
}

type RetryPolicyState struct {
	MaximumBackoff string `json:"maximum_backoff"`
	MinimumBackoff string `json:"minimum_backoff"`
}
