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
	// NoWrapper: min=0
	NoWrapper []NoWrapper `hcl:"no_wrapper,block" validate:"min=0"`
	// OidcToken: min=0
	OidcToken []OidcToken `hcl:"oidc_token,block" validate:"min=0"`
}

type NoWrapper struct{}

type OidcToken struct{}

type RetryPolicy struct{}

type BigqueryConfigAttributes struct {
	ref terra.Reference
}

func (bc BigqueryConfigAttributes) InternalRef() (terra.Reference, error) {
	return bc.ref, nil
}

func (bc BigqueryConfigAttributes) InternalWithRef(ref terra.Reference) BigqueryConfigAttributes {
	return BigqueryConfigAttributes{ref: ref}
}

func (bc BigqueryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bc.ref.InternalTokens()
}

func (bc BigqueryConfigAttributes) DropUnknownFields() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("drop_unknown_fields"))
}

func (bc BigqueryConfigAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("table"))
}

func (bc BigqueryConfigAttributes) UseTopicSchema() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("use_topic_schema"))
}

func (bc BigqueryConfigAttributes) WriteMetadata() terra.BoolValue {
	return terra.ReferenceAsBool(bc.ref.Append("write_metadata"))
}

type DeadLetterPolicyAttributes struct {
	ref terra.Reference
}

func (dlp DeadLetterPolicyAttributes) InternalRef() (terra.Reference, error) {
	return dlp.ref, nil
}

func (dlp DeadLetterPolicyAttributes) InternalWithRef(ref terra.Reference) DeadLetterPolicyAttributes {
	return DeadLetterPolicyAttributes{ref: ref}
}

func (dlp DeadLetterPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dlp.ref.InternalTokens()
}

func (dlp DeadLetterPolicyAttributes) DeadLetterTopic() terra.StringValue {
	return terra.ReferenceAsString(dlp.ref.Append("dead_letter_topic"))
}

func (dlp DeadLetterPolicyAttributes) MaxDeliveryAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(dlp.ref.Append("max_delivery_attempts"))
}

type ExpirationPolicyAttributes struct {
	ref terra.Reference
}

func (ep ExpirationPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep ExpirationPolicyAttributes) InternalWithRef(ref terra.Reference) ExpirationPolicyAttributes {
	return ExpirationPolicyAttributes{ref: ref}
}

func (ep ExpirationPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep ExpirationPolicyAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("ttl"))
}

type PushConfigAttributes struct {
	ref terra.Reference
}

func (pc PushConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PushConfigAttributes) InternalWithRef(ref terra.Reference) PushConfigAttributes {
	return PushConfigAttributes{ref: ref}
}

func (pc PushConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc PushConfigAttributes) Attributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pc.ref.Append("attributes"))
}

func (pc PushConfigAttributes) PushEndpoint() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("push_endpoint"))
}

func (pc PushConfigAttributes) NoWrapper() terra.ListValue[NoWrapperAttributes] {
	return terra.ReferenceAsList[NoWrapperAttributes](pc.ref.Append("no_wrapper"))
}

func (pc PushConfigAttributes) OidcToken() terra.ListValue[OidcTokenAttributes] {
	return terra.ReferenceAsList[OidcTokenAttributes](pc.ref.Append("oidc_token"))
}

type NoWrapperAttributes struct {
	ref terra.Reference
}

func (nw NoWrapperAttributes) InternalRef() (terra.Reference, error) {
	return nw.ref, nil
}

func (nw NoWrapperAttributes) InternalWithRef(ref terra.Reference) NoWrapperAttributes {
	return NoWrapperAttributes{ref: ref}
}

func (nw NoWrapperAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nw.ref.InternalTokens()
}

func (nw NoWrapperAttributes) WriteMetadata() terra.BoolValue {
	return terra.ReferenceAsBool(nw.ref.Append("write_metadata"))
}

type OidcTokenAttributes struct {
	ref terra.Reference
}

func (ot OidcTokenAttributes) InternalRef() (terra.Reference, error) {
	return ot.ref, nil
}

func (ot OidcTokenAttributes) InternalWithRef(ref terra.Reference) OidcTokenAttributes {
	return OidcTokenAttributes{ref: ref}
}

func (ot OidcTokenAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ot.ref.InternalTokens()
}

func (ot OidcTokenAttributes) Audience() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("audience"))
}

func (ot OidcTokenAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("service_account_email"))
}

type RetryPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetryPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RetryPolicyAttributes) InternalWithRef(ref terra.Reference) RetryPolicyAttributes {
	return RetryPolicyAttributes{ref: ref}
}

func (rp RetryPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RetryPolicyAttributes) MaximumBackoff() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("maximum_backoff"))
}

func (rp RetryPolicyAttributes) MinimumBackoff() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("minimum_backoff"))
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
	NoWrapper    []NoWrapperState  `json:"no_wrapper"`
	OidcToken    []OidcTokenState  `json:"oidc_token"`
}

type NoWrapperState struct {
	WriteMetadata bool `json:"write_metadata"`
}

type OidcTokenState struct {
	Audience            string `json:"audience"`
	ServiceAccountEmail string `json:"service_account_email"`
}

type RetryPolicyState struct {
	MaximumBackoff string `json:"maximum_backoff"`
	MinimumBackoff string `json:"minimum_backoff"`
}
