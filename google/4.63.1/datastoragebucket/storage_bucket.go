// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datastoragebucket

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Autoclass struct{}

type Cors struct{}

type CustomPlacementConfig struct{}

type Encryption struct{}

type LifecycleRule struct {
	// Action: min=0
	Action []Action `hcl:"action,block" validate:"min=0"`
	// Condition: min=0
	Condition []Condition `hcl:"condition,block" validate:"min=0"`
}

type Action struct{}

type Condition struct{}

type Logging struct{}

type RetentionPolicy struct{}

type Versioning struct{}

type Website struct{}

type AutoclassAttributes struct {
	ref terra.Reference
}

func (a AutoclassAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AutoclassAttributes) InternalWithRef(ref terra.Reference) AutoclassAttributes {
	return AutoclassAttributes{ref: ref}
}

func (a AutoclassAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AutoclassAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enabled"))
}

type CorsAttributes struct {
	ref terra.Reference
}

func (c CorsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CorsAttributes) InternalWithRef(ref terra.Reference) CorsAttributes {
	return CorsAttributes{ref: ref}
}

func (c CorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CorsAttributes) MaxAgeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("max_age_seconds"))
}

func (c CorsAttributes) Method() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("method"))
}

func (c CorsAttributes) Origin() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("origin"))
}

func (c CorsAttributes) ResponseHeader() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("response_header"))
}

type CustomPlacementConfigAttributes struct {
	ref terra.Reference
}

func (cpc CustomPlacementConfigAttributes) InternalRef() (terra.Reference, error) {
	return cpc.ref, nil
}

func (cpc CustomPlacementConfigAttributes) InternalWithRef(ref terra.Reference) CustomPlacementConfigAttributes {
	return CustomPlacementConfigAttributes{ref: ref}
}

func (cpc CustomPlacementConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpc.ref.InternalTokens()
}

func (cpc CustomPlacementConfigAttributes) DataLocations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cpc.ref.Append("data_locations"))
}

type EncryptionAttributes struct {
	ref terra.Reference
}

func (e EncryptionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EncryptionAttributes) InternalWithRef(ref terra.Reference) EncryptionAttributes {
	return EncryptionAttributes{ref: ref}
}

func (e EncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EncryptionAttributes) DefaultKmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("default_kms_key_name"))
}

type LifecycleRuleAttributes struct {
	ref terra.Reference
}

func (lr LifecycleRuleAttributes) InternalRef() (terra.Reference, error) {
	return lr.ref, nil
}

func (lr LifecycleRuleAttributes) InternalWithRef(ref terra.Reference) LifecycleRuleAttributes {
	return LifecycleRuleAttributes{ref: ref}
}

func (lr LifecycleRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lr.ref.InternalTokens()
}

func (lr LifecycleRuleAttributes) Action() terra.SetValue[ActionAttributes] {
	return terra.ReferenceAsSet[ActionAttributes](lr.ref.Append("action"))
}

func (lr LifecycleRuleAttributes) Condition() terra.SetValue[ConditionAttributes] {
	return terra.ReferenceAsSet[ConditionAttributes](lr.ref.Append("condition"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("storage_class"))
}

func (a ActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) Age() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("age"))
}

func (c ConditionAttributes) CreatedBefore() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("created_before"))
}

func (c ConditionAttributes) CustomTimeBefore() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("custom_time_before"))
}

func (c ConditionAttributes) DaysSinceCustomTime() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("days_since_custom_time"))
}

func (c ConditionAttributes) DaysSinceNoncurrentTime() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("days_since_noncurrent_time"))
}

func (c ConditionAttributes) MatchesPrefix() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("matches_prefix"))
}

func (c ConditionAttributes) MatchesStorageClass() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("matches_storage_class"))
}

func (c ConditionAttributes) MatchesSuffix() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("matches_suffix"))
}

func (c ConditionAttributes) NoncurrentTimeBefore() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("noncurrent_time_before"))
}

func (c ConditionAttributes) NumNewerVersions() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("num_newer_versions"))
}

func (c ConditionAttributes) WithState() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("with_state"))
}

type LoggingAttributes struct {
	ref terra.Reference
}

func (l LoggingAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LoggingAttributes) InternalWithRef(ref terra.Reference) LoggingAttributes {
	return LoggingAttributes{ref: ref}
}

func (l LoggingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LoggingAttributes) LogBucket() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("log_bucket"))
}

func (l LoggingAttributes) LogObjectPrefix() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("log_object_prefix"))
}

type RetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RetentionPolicyAttributes) InternalWithRef(ref terra.Reference) RetentionPolicyAttributes {
	return RetentionPolicyAttributes{ref: ref}
}

func (rp RetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RetentionPolicyAttributes) IsLocked() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("is_locked"))
}

func (rp RetentionPolicyAttributes) RetentionPeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("retention_period"))
}

type VersioningAttributes struct {
	ref terra.Reference
}

func (v VersioningAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VersioningAttributes) InternalWithRef(ref terra.Reference) VersioningAttributes {
	return VersioningAttributes{ref: ref}
}

func (v VersioningAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VersioningAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("enabled"))
}

type WebsiteAttributes struct {
	ref terra.Reference
}

func (w WebsiteAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WebsiteAttributes) InternalWithRef(ref terra.Reference) WebsiteAttributes {
	return WebsiteAttributes{ref: ref}
}

func (w WebsiteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WebsiteAttributes) MainPageSuffix() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("main_page_suffix"))
}

func (w WebsiteAttributes) NotFoundPage() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("not_found_page"))
}

type AutoclassState struct {
	Enabled bool `json:"enabled"`
}

type CorsState struct {
	MaxAgeSeconds  float64  `json:"max_age_seconds"`
	Method         []string `json:"method"`
	Origin         []string `json:"origin"`
	ResponseHeader []string `json:"response_header"`
}

type CustomPlacementConfigState struct {
	DataLocations []string `json:"data_locations"`
}

type EncryptionState struct {
	DefaultKmsKeyName string `json:"default_kms_key_name"`
}

type LifecycleRuleState struct {
	Action    []ActionState    `json:"action"`
	Condition []ConditionState `json:"condition"`
}

type ActionState struct {
	StorageClass string `json:"storage_class"`
	Type         string `json:"type"`
}

type ConditionState struct {
	Age                     float64  `json:"age"`
	CreatedBefore           string   `json:"created_before"`
	CustomTimeBefore        string   `json:"custom_time_before"`
	DaysSinceCustomTime     float64  `json:"days_since_custom_time"`
	DaysSinceNoncurrentTime float64  `json:"days_since_noncurrent_time"`
	MatchesPrefix           []string `json:"matches_prefix"`
	MatchesStorageClass     []string `json:"matches_storage_class"`
	MatchesSuffix           []string `json:"matches_suffix"`
	NoncurrentTimeBefore    string   `json:"noncurrent_time_before"`
	NumNewerVersions        float64  `json:"num_newer_versions"`
	WithState               string   `json:"with_state"`
}

type LoggingState struct {
	LogBucket       string `json:"log_bucket"`
	LogObjectPrefix string `json:"log_object_prefix"`
}

type RetentionPolicyState struct {
	IsLocked        bool    `json:"is_locked"`
	RetentionPeriod float64 `json:"retention_period"`
}

type VersioningState struct {
	Enabled bool `json:"enabled"`
}

type WebsiteState struct {
	MainPageSuffix string `json:"main_page_suffix"`
	NotFoundPage   string `json:"not_found_page"`
}