// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucket

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CorsRule struct {
	// AllowedHeaders: list of string, optional
	AllowedHeaders terra.ListValue[terra.StringValue] `hcl:"allowed_headers,attr"`
	// AllowedMethods: list of string, required
	AllowedMethods terra.ListValue[terra.StringValue] `hcl:"allowed_methods,attr" validate:"required"`
	// AllowedOrigins: list of string, required
	AllowedOrigins terra.ListValue[terra.StringValue] `hcl:"allowed_origins,attr" validate:"required"`
	// ExposeHeaders: list of string, optional
	ExposeHeaders terra.ListValue[terra.StringValue] `hcl:"expose_headers,attr"`
	// MaxAgeSeconds: number, optional
	MaxAgeSeconds terra.NumberValue `hcl:"max_age_seconds,attr"`
}

type Grant struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Permissions: set of string, required
	Permissions terra.SetValue[terra.StringValue] `hcl:"permissions,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
}

type LifecycleRule struct {
	// AbortIncompleteMultipartUploadDays: number, optional
	AbortIncompleteMultipartUploadDays terra.NumberValue `hcl:"abort_incomplete_multipart_upload_days,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Expiration: optional
	Expiration *Expiration `hcl:"expiration,block"`
	// NoncurrentVersionExpiration: optional
	NoncurrentVersionExpiration *NoncurrentVersionExpiration `hcl:"noncurrent_version_expiration,block"`
	// NoncurrentVersionTransition: min=0
	NoncurrentVersionTransition []NoncurrentVersionTransition `hcl:"noncurrent_version_transition,block" validate:"min=0"`
	// Transition: min=0
	Transition []Transition `hcl:"transition,block" validate:"min=0"`
}

type Expiration struct {
	// Date: string, optional
	Date terra.StringValue `hcl:"date,attr"`
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// ExpiredObjectDeleteMarker: bool, optional
	ExpiredObjectDeleteMarker terra.BoolValue `hcl:"expired_object_delete_marker,attr"`
}

type NoncurrentVersionExpiration struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
}

type NoncurrentVersionTransition struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// StorageClass: string, required
	StorageClass terra.StringValue `hcl:"storage_class,attr" validate:"required"`
}

type Transition struct {
	// Date: string, optional
	Date terra.StringValue `hcl:"date,attr"`
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// StorageClass: string, required
	StorageClass terra.StringValue `hcl:"storage_class,attr" validate:"required"`
}

type Logging struct {
	// TargetBucket: string, required
	TargetBucket terra.StringValue `hcl:"target_bucket,attr" validate:"required"`
	// TargetPrefix: string, optional
	TargetPrefix terra.StringValue `hcl:"target_prefix,attr"`
}

type ObjectLockConfiguration struct {
	// ObjectLockEnabled: string, optional
	ObjectLockEnabled terra.StringValue `hcl:"object_lock_enabled,attr"`
	// ObjectLockConfigurationRule: optional
	Rule *ObjectLockConfigurationRule `hcl:"rule,block"`
}

type ObjectLockConfigurationRule struct {
	// DefaultRetention: required
	DefaultRetention *DefaultRetention `hcl:"default_retention,block" validate:"required"`
}

type DefaultRetention struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Years: number, optional
	Years terra.NumberValue `hcl:"years,attr"`
}

type ReplicationConfiguration struct {
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Rules: min=1
	Rules []Rules `hcl:"rules,block" validate:"min=1"`
}

type Rules struct {
	// DeleteMarkerReplicationStatus: string, optional
	DeleteMarkerReplicationStatus terra.StringValue `hcl:"delete_marker_replication_status,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
	// Destination: required
	Destination *Destination `hcl:"destination,block" validate:"required"`
	// Filter: optional
	Filter *Filter `hcl:"filter,block"`
	// SourceSelectionCriteria: optional
	SourceSelectionCriteria *SourceSelectionCriteria `hcl:"source_selection_criteria,block"`
}

type Destination struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ReplicaKmsKeyId: string, optional
	ReplicaKmsKeyId terra.StringValue `hcl:"replica_kms_key_id,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
	// AccessControlTranslation: optional
	AccessControlTranslation *AccessControlTranslation `hcl:"access_control_translation,block"`
	// Metrics: optional
	Metrics *Metrics `hcl:"metrics,block"`
	// ReplicationTime: optional
	ReplicationTime *ReplicationTime `hcl:"replication_time,block"`
}

type AccessControlTranslation struct {
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
}

type Metrics struct {
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
}

type ReplicationTime struct {
	// Minutes: number, optional
	Minutes terra.NumberValue `hcl:"minutes,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
}

type Filter struct {
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type SourceSelectionCriteria struct {
	// SseKmsEncryptedObjects: optional
	SseKmsEncryptedObjects *SseKmsEncryptedObjects `hcl:"sse_kms_encrypted_objects,block"`
}

type SseKmsEncryptedObjects struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type ServerSideEncryptionConfiguration struct {
	// ServerSideEncryptionConfigurationRule: required
	Rule *ServerSideEncryptionConfigurationRule `hcl:"rule,block" validate:"required"`
}

type ServerSideEncryptionConfigurationRule struct {
	// BucketKeyEnabled: bool, optional
	BucketKeyEnabled terra.BoolValue `hcl:"bucket_key_enabled,attr"`
	// ApplyServerSideEncryptionByDefault: required
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefault `hcl:"apply_server_side_encryption_by_default,block" validate:"required"`
}

type ApplyServerSideEncryptionByDefault struct {
	// KmsMasterKeyId: string, optional
	KmsMasterKeyId terra.StringValue `hcl:"kms_master_key_id,attr"`
	// SseAlgorithm: string, required
	SseAlgorithm terra.StringValue `hcl:"sse_algorithm,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Versioning struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// MfaDelete: bool, optional
	MfaDelete terra.BoolValue `hcl:"mfa_delete,attr"`
}

type Website struct {
	// ErrorDocument: string, optional
	ErrorDocument terra.StringValue `hcl:"error_document,attr"`
	// IndexDocument: string, optional
	IndexDocument terra.StringValue `hcl:"index_document,attr"`
	// RedirectAllRequestsTo: string, optional
	RedirectAllRequestsTo terra.StringValue `hcl:"redirect_all_requests_to,attr"`
	// RoutingRules: string, optional
	RoutingRules terra.StringValue `hcl:"routing_rules,attr"`
}

type CorsRuleAttributes struct {
	ref terra.Reference
}

func (cr CorsRuleAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CorsRuleAttributes) InternalWithRef(ref terra.Reference) CorsRuleAttributes {
	return CorsRuleAttributes{ref: ref}
}

func (cr CorsRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CorsRuleAttributes) AllowedHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_headers"))
}

func (cr CorsRuleAttributes) AllowedMethods() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_methods"))
}

func (cr CorsRuleAttributes) AllowedOrigins() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("allowed_origins"))
}

func (cr CorsRuleAttributes) ExposeHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cr.ref.Append("expose_headers"))
}

func (cr CorsRuleAttributes) MaxAgeSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cr.ref.Append("max_age_seconds"))
}

type GrantAttributes struct {
	ref terra.Reference
}

func (g GrantAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GrantAttributes) InternalWithRef(ref terra.Reference) GrantAttributes {
	return GrantAttributes{ref: ref}
}

func (g GrantAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GrantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("id"))
}

func (g GrantAttributes) Permissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](g.ref.Append("permissions"))
}

func (g GrantAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("type"))
}

func (g GrantAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("uri"))
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

func (lr LifecycleRuleAttributes) AbortIncompleteMultipartUploadDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lr.ref.Append("abort_incomplete_multipart_upload_days"))
}

func (lr LifecycleRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lr.ref.Append("enabled"))
}

func (lr LifecycleRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("id"))
}

func (lr LifecycleRuleAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(lr.ref.Append("prefix"))
}

func (lr LifecycleRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lr.ref.Append("tags"))
}

func (lr LifecycleRuleAttributes) Expiration() terra.ListValue[ExpirationAttributes] {
	return terra.ReferenceAsList[ExpirationAttributes](lr.ref.Append("expiration"))
}

func (lr LifecycleRuleAttributes) NoncurrentVersionExpiration() terra.ListValue[NoncurrentVersionExpirationAttributes] {
	return terra.ReferenceAsList[NoncurrentVersionExpirationAttributes](lr.ref.Append("noncurrent_version_expiration"))
}

func (lr LifecycleRuleAttributes) NoncurrentVersionTransition() terra.SetValue[NoncurrentVersionTransitionAttributes] {
	return terra.ReferenceAsSet[NoncurrentVersionTransitionAttributes](lr.ref.Append("noncurrent_version_transition"))
}

func (lr LifecycleRuleAttributes) Transition() terra.SetValue[TransitionAttributes] {
	return terra.ReferenceAsSet[TransitionAttributes](lr.ref.Append("transition"))
}

type ExpirationAttributes struct {
	ref terra.Reference
}

func (e ExpirationAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExpirationAttributes) InternalWithRef(ref terra.Reference) ExpirationAttributes {
	return ExpirationAttributes{ref: ref}
}

func (e ExpirationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExpirationAttributes) Date() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("date"))
}

func (e ExpirationAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("days"))
}

func (e ExpirationAttributes) ExpiredObjectDeleteMarker() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("expired_object_delete_marker"))
}

type NoncurrentVersionExpirationAttributes struct {
	ref terra.Reference
}

func (nve NoncurrentVersionExpirationAttributes) InternalRef() (terra.Reference, error) {
	return nve.ref, nil
}

func (nve NoncurrentVersionExpirationAttributes) InternalWithRef(ref terra.Reference) NoncurrentVersionExpirationAttributes {
	return NoncurrentVersionExpirationAttributes{ref: ref}
}

func (nve NoncurrentVersionExpirationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nve.ref.InternalTokens()
}

func (nve NoncurrentVersionExpirationAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(nve.ref.Append("days"))
}

type NoncurrentVersionTransitionAttributes struct {
	ref terra.Reference
}

func (nvt NoncurrentVersionTransitionAttributes) InternalRef() (terra.Reference, error) {
	return nvt.ref, nil
}

func (nvt NoncurrentVersionTransitionAttributes) InternalWithRef(ref terra.Reference) NoncurrentVersionTransitionAttributes {
	return NoncurrentVersionTransitionAttributes{ref: ref}
}

func (nvt NoncurrentVersionTransitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nvt.ref.InternalTokens()
}

func (nvt NoncurrentVersionTransitionAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(nvt.ref.Append("days"))
}

func (nvt NoncurrentVersionTransitionAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(nvt.ref.Append("storage_class"))
}

type TransitionAttributes struct {
	ref terra.Reference
}

func (t TransitionAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TransitionAttributes) InternalWithRef(ref terra.Reference) TransitionAttributes {
	return TransitionAttributes{ref: ref}
}

func (t TransitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TransitionAttributes) Date() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("date"))
}

func (t TransitionAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("days"))
}

func (t TransitionAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("storage_class"))
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

func (l LoggingAttributes) TargetBucket() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("target_bucket"))
}

func (l LoggingAttributes) TargetPrefix() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("target_prefix"))
}

type ObjectLockConfigurationAttributes struct {
	ref terra.Reference
}

func (olc ObjectLockConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return olc.ref, nil
}

func (olc ObjectLockConfigurationAttributes) InternalWithRef(ref terra.Reference) ObjectLockConfigurationAttributes {
	return ObjectLockConfigurationAttributes{ref: ref}
}

func (olc ObjectLockConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return olc.ref.InternalTokens()
}

func (olc ObjectLockConfigurationAttributes) ObjectLockEnabled() terra.StringValue {
	return terra.ReferenceAsString(olc.ref.Append("object_lock_enabled"))
}

func (olc ObjectLockConfigurationAttributes) Rule() terra.ListValue[ObjectLockConfigurationRuleAttributes] {
	return terra.ReferenceAsList[ObjectLockConfigurationRuleAttributes](olc.ref.Append("rule"))
}

type ObjectLockConfigurationRuleAttributes struct {
	ref terra.Reference
}

func (r ObjectLockConfigurationRuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ObjectLockConfigurationRuleAttributes) InternalWithRef(ref terra.Reference) ObjectLockConfigurationRuleAttributes {
	return ObjectLockConfigurationRuleAttributes{ref: ref}
}

func (r ObjectLockConfigurationRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ObjectLockConfigurationRuleAttributes) DefaultRetention() terra.ListValue[DefaultRetentionAttributes] {
	return terra.ReferenceAsList[DefaultRetentionAttributes](r.ref.Append("default_retention"))
}

type DefaultRetentionAttributes struct {
	ref terra.Reference
}

func (dr DefaultRetentionAttributes) InternalRef() (terra.Reference, error) {
	return dr.ref, nil
}

func (dr DefaultRetentionAttributes) InternalWithRef(ref terra.Reference) DefaultRetentionAttributes {
	return DefaultRetentionAttributes{ref: ref}
}

func (dr DefaultRetentionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dr.ref.InternalTokens()
}

func (dr DefaultRetentionAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(dr.ref.Append("days"))
}

func (dr DefaultRetentionAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("mode"))
}

func (dr DefaultRetentionAttributes) Years() terra.NumberValue {
	return terra.ReferenceAsNumber(dr.ref.Append("years"))
}

type ReplicationConfigurationAttributes struct {
	ref terra.Reference
}

func (rc ReplicationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReplicationConfigurationAttributes) InternalWithRef(ref terra.Reference) ReplicationConfigurationAttributes {
	return ReplicationConfigurationAttributes{ref: ref}
}

func (rc ReplicationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReplicationConfigurationAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("role"))
}

func (rc ReplicationConfigurationAttributes) Rules() terra.SetValue[RulesAttributes] {
	return terra.ReferenceAsSet[RulesAttributes](rc.ref.Append("rules"))
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) DeleteMarkerReplicationStatus() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("delete_marker_replication_status"))
}

func (r RulesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

func (r RulesAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("prefix"))
}

func (r RulesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("priority"))
}

func (r RulesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status"))
}

func (r RulesAttributes) Destination() terra.ListValue[DestinationAttributes] {
	return terra.ReferenceAsList[DestinationAttributes](r.ref.Append("destination"))
}

func (r RulesAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](r.ref.Append("filter"))
}

func (r RulesAttributes) SourceSelectionCriteria() terra.ListValue[SourceSelectionCriteriaAttributes] {
	return terra.ReferenceAsList[SourceSelectionCriteriaAttributes](r.ref.Append("source_selection_criteria"))
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("account_id"))
}

func (d DestinationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("bucket"))
}

func (d DestinationAttributes) ReplicaKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("replica_kms_key_id"))
}

func (d DestinationAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("storage_class"))
}

func (d DestinationAttributes) AccessControlTranslation() terra.ListValue[AccessControlTranslationAttributes] {
	return terra.ReferenceAsList[AccessControlTranslationAttributes](d.ref.Append("access_control_translation"))
}

func (d DestinationAttributes) Metrics() terra.ListValue[MetricsAttributes] {
	return terra.ReferenceAsList[MetricsAttributes](d.ref.Append("metrics"))
}

func (d DestinationAttributes) ReplicationTime() terra.ListValue[ReplicationTimeAttributes] {
	return terra.ReferenceAsList[ReplicationTimeAttributes](d.ref.Append("replication_time"))
}

type AccessControlTranslationAttributes struct {
	ref terra.Reference
}

func (act AccessControlTranslationAttributes) InternalRef() (terra.Reference, error) {
	return act.ref, nil
}

func (act AccessControlTranslationAttributes) InternalWithRef(ref terra.Reference) AccessControlTranslationAttributes {
	return AccessControlTranslationAttributes{ref: ref}
}

func (act AccessControlTranslationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return act.ref.InternalTokens()
}

func (act AccessControlTranslationAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(act.ref.Append("owner"))
}

type MetricsAttributes struct {
	ref terra.Reference
}

func (m MetricsAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetricsAttributes) InternalWithRef(ref terra.Reference) MetricsAttributes {
	return MetricsAttributes{ref: ref}
}

func (m MetricsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetricsAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("minutes"))
}

func (m MetricsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("status"))
}

type ReplicationTimeAttributes struct {
	ref terra.Reference
}

func (rt ReplicationTimeAttributes) InternalRef() (terra.Reference, error) {
	return rt.ref, nil
}

func (rt ReplicationTimeAttributes) InternalWithRef(ref terra.Reference) ReplicationTimeAttributes {
	return ReplicationTimeAttributes{ref: ref}
}

func (rt ReplicationTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rt.ref.InternalTokens()
}

func (rt ReplicationTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(rt.ref.Append("minutes"))
}

func (rt ReplicationTimeAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("status"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("prefix"))
}

func (f FilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("tags"))
}

type SourceSelectionCriteriaAttributes struct {
	ref terra.Reference
}

func (ssc SourceSelectionCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return ssc.ref, nil
}

func (ssc SourceSelectionCriteriaAttributes) InternalWithRef(ref terra.Reference) SourceSelectionCriteriaAttributes {
	return SourceSelectionCriteriaAttributes{ref: ref}
}

func (ssc SourceSelectionCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssc.ref.InternalTokens()
}

func (ssc SourceSelectionCriteriaAttributes) SseKmsEncryptedObjects() terra.ListValue[SseKmsEncryptedObjectsAttributes] {
	return terra.ReferenceAsList[SseKmsEncryptedObjectsAttributes](ssc.ref.Append("sse_kms_encrypted_objects"))
}

type SseKmsEncryptedObjectsAttributes struct {
	ref terra.Reference
}

func (skeo SseKmsEncryptedObjectsAttributes) InternalRef() (terra.Reference, error) {
	return skeo.ref, nil
}

func (skeo SseKmsEncryptedObjectsAttributes) InternalWithRef(ref terra.Reference) SseKmsEncryptedObjectsAttributes {
	return SseKmsEncryptedObjectsAttributes{ref: ref}
}

func (skeo SseKmsEncryptedObjectsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return skeo.ref.InternalTokens()
}

func (skeo SseKmsEncryptedObjectsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(skeo.ref.Append("enabled"))
}

type ServerSideEncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ssec.ref, nil
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) ServerSideEncryptionConfigurationAttributes {
	return ServerSideEncryptionConfigurationAttributes{ref: ref}
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssec.ref.InternalTokens()
}

func (ssec ServerSideEncryptionConfigurationAttributes) Rule() terra.ListValue[ServerSideEncryptionConfigurationRuleAttributes] {
	return terra.ReferenceAsList[ServerSideEncryptionConfigurationRuleAttributes](ssec.ref.Append("rule"))
}

type ServerSideEncryptionConfigurationRuleAttributes struct {
	ref terra.Reference
}

func (r ServerSideEncryptionConfigurationRuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ServerSideEncryptionConfigurationRuleAttributes) InternalWithRef(ref terra.Reference) ServerSideEncryptionConfigurationRuleAttributes {
	return ServerSideEncryptionConfigurationRuleAttributes{ref: ref}
}

func (r ServerSideEncryptionConfigurationRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ServerSideEncryptionConfigurationRuleAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("bucket_key_enabled"))
}

func (r ServerSideEncryptionConfigurationRuleAttributes) ApplyServerSideEncryptionByDefault() terra.ListValue[ApplyServerSideEncryptionByDefaultAttributes] {
	return terra.ReferenceAsList[ApplyServerSideEncryptionByDefaultAttributes](r.ref.Append("apply_server_side_encryption_by_default"))
}

type ApplyServerSideEncryptionByDefaultAttributes struct {
	ref terra.Reference
}

func (assebd ApplyServerSideEncryptionByDefaultAttributes) InternalRef() (terra.Reference, error) {
	return assebd.ref, nil
}

func (assebd ApplyServerSideEncryptionByDefaultAttributes) InternalWithRef(ref terra.Reference) ApplyServerSideEncryptionByDefaultAttributes {
	return ApplyServerSideEncryptionByDefaultAttributes{ref: ref}
}

func (assebd ApplyServerSideEncryptionByDefaultAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return assebd.ref.InternalTokens()
}

func (assebd ApplyServerSideEncryptionByDefaultAttributes) KmsMasterKeyId() terra.StringValue {
	return terra.ReferenceAsString(assebd.ref.Append("kms_master_key_id"))
}

func (assebd ApplyServerSideEncryptionByDefaultAttributes) SseAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(assebd.ref.Append("sse_algorithm"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
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

func (v VersioningAttributes) MfaDelete() terra.BoolValue {
	return terra.ReferenceAsBool(v.ref.Append("mfa_delete"))
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

func (w WebsiteAttributes) ErrorDocument() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("error_document"))
}

func (w WebsiteAttributes) IndexDocument() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("index_document"))
}

func (w WebsiteAttributes) RedirectAllRequestsTo() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("redirect_all_requests_to"))
}

func (w WebsiteAttributes) RoutingRules() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("routing_rules"))
}

type CorsRuleState struct {
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers"`
	MaxAgeSeconds  float64  `json:"max_age_seconds"`
}

type GrantState struct {
	Id          string   `json:"id"`
	Permissions []string `json:"permissions"`
	Type        string   `json:"type"`
	Uri         string   `json:"uri"`
}

type LifecycleRuleState struct {
	AbortIncompleteMultipartUploadDays float64                            `json:"abort_incomplete_multipart_upload_days"`
	Enabled                            bool                               `json:"enabled"`
	Id                                 string                             `json:"id"`
	Prefix                             string                             `json:"prefix"`
	Tags                               map[string]string                  `json:"tags"`
	Expiration                         []ExpirationState                  `json:"expiration"`
	NoncurrentVersionExpiration        []NoncurrentVersionExpirationState `json:"noncurrent_version_expiration"`
	NoncurrentVersionTransition        []NoncurrentVersionTransitionState `json:"noncurrent_version_transition"`
	Transition                         []TransitionState                  `json:"transition"`
}

type ExpirationState struct {
	Date                      string  `json:"date"`
	Days                      float64 `json:"days"`
	ExpiredObjectDeleteMarker bool    `json:"expired_object_delete_marker"`
}

type NoncurrentVersionExpirationState struct {
	Days float64 `json:"days"`
}

type NoncurrentVersionTransitionState struct {
	Days         float64 `json:"days"`
	StorageClass string  `json:"storage_class"`
}

type TransitionState struct {
	Date         string  `json:"date"`
	Days         float64 `json:"days"`
	StorageClass string  `json:"storage_class"`
}

type LoggingState struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
}

type ObjectLockConfigurationState struct {
	ObjectLockEnabled string                             `json:"object_lock_enabled"`
	Rule              []ObjectLockConfigurationRuleState `json:"rule"`
}

type ObjectLockConfigurationRuleState struct {
	DefaultRetention []DefaultRetentionState `json:"default_retention"`
}

type DefaultRetentionState struct {
	Days  float64 `json:"days"`
	Mode  string  `json:"mode"`
	Years float64 `json:"years"`
}

type ReplicationConfigurationState struct {
	Role  string       `json:"role"`
	Rules []RulesState `json:"rules"`
}

type RulesState struct {
	DeleteMarkerReplicationStatus string                         `json:"delete_marker_replication_status"`
	Id                            string                         `json:"id"`
	Prefix                        string                         `json:"prefix"`
	Priority                      float64                        `json:"priority"`
	Status                        string                         `json:"status"`
	Destination                   []DestinationState             `json:"destination"`
	Filter                        []FilterState                  `json:"filter"`
	SourceSelectionCriteria       []SourceSelectionCriteriaState `json:"source_selection_criteria"`
}

type DestinationState struct {
	AccountId                string                          `json:"account_id"`
	Bucket                   string                          `json:"bucket"`
	ReplicaKmsKeyId          string                          `json:"replica_kms_key_id"`
	StorageClass             string                          `json:"storage_class"`
	AccessControlTranslation []AccessControlTranslationState `json:"access_control_translation"`
	Metrics                  []MetricsState                  `json:"metrics"`
	ReplicationTime          []ReplicationTimeState          `json:"replication_time"`
}

type AccessControlTranslationState struct {
	Owner string `json:"owner"`
}

type MetricsState struct {
	Minutes float64 `json:"minutes"`
	Status  string  `json:"status"`
}

type ReplicationTimeState struct {
	Minutes float64 `json:"minutes"`
	Status  string  `json:"status"`
}

type FilterState struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type SourceSelectionCriteriaState struct {
	SseKmsEncryptedObjects []SseKmsEncryptedObjectsState `json:"sse_kms_encrypted_objects"`
}

type SseKmsEncryptedObjectsState struct {
	Enabled bool `json:"enabled"`
}

type ServerSideEncryptionConfigurationState struct {
	Rule []ServerSideEncryptionConfigurationRuleState `json:"rule"`
}

type ServerSideEncryptionConfigurationRuleState struct {
	BucketKeyEnabled                   bool                                      `json:"bucket_key_enabled"`
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultState `json:"apply_server_side_encryption_by_default"`
}

type ApplyServerSideEncryptionByDefaultState struct {
	KmsMasterKeyId string `json:"kms_master_key_id"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type VersioningState struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

type WebsiteState struct {
	ErrorDocument         string `json:"error_document"`
	IndexDocument         string `json:"index_document"`
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
	RoutingRules          string `json:"routing_rules"`
}