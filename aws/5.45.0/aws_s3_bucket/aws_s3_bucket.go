// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_bucket

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_s3_bucket.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3BucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asb *Resource) Type() string {
	return "aws_s3_bucket"
}

// LocalName returns the local name for [Resource].
func (asb *Resource) LocalName() string {
	return asb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asb *Resource) Configuration() interface{} {
	return asb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asb *Resource) Dependencies() terra.Dependencies {
	return asb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asb *Resource) LifecycleManagement() *terra.Lifecycle {
	return asb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asb *Resource) Attributes() awsS3BucketAttributes {
	return awsS3BucketAttributes{ref: terra.ReferenceResource(asb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asb *Resource) ImportState(state io.Reader) error {
	asb.state = &awsS3BucketState{}
	if err := json.NewDecoder(state).Decode(asb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asb.Type(), asb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asb *Resource) State() (*awsS3BucketState, bool) {
	return asb.state, asb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asb *Resource) StateMust() *awsS3BucketState {
	if asb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asb.Type(), asb.LocalName()))
	}
	return asb.state
}

// Args contains the configurations for aws_s3_bucket.
type Args struct {
	// AccelerationStatus: string, optional
	AccelerationStatus terra.StringValue `hcl:"acceleration_status,attr"`
	// Acl: string, optional
	Acl terra.StringValue `hcl:"acl,attr"`
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// BucketPrefix: string, optional
	BucketPrefix terra.StringValue `hcl:"bucket_prefix,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectLockEnabled: bool, optional
	ObjectLockEnabled terra.BoolValue `hcl:"object_lock_enabled,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// RequestPayer: string, optional
	RequestPayer terra.StringValue `hcl:"request_payer,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CorsRule: min=0
	CorsRule []CorsRule `hcl:"cors_rule,block" validate:"min=0"`
	// Grant: min=0
	Grant []Grant `hcl:"grant,block" validate:"min=0"`
	// LifecycleRule: min=0
	LifecycleRule []LifecycleRule `hcl:"lifecycle_rule,block" validate:"min=0"`
	// Logging: optional
	Logging *Logging `hcl:"logging,block"`
	// ObjectLockConfiguration: optional
	ObjectLockConfiguration *ObjectLockConfiguration `hcl:"object_lock_configuration,block"`
	// ReplicationConfiguration: optional
	ReplicationConfiguration *ReplicationConfiguration `hcl:"replication_configuration,block"`
	// ServerSideEncryptionConfiguration: optional
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `hcl:"server_side_encryption_configuration,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// Versioning: optional
	Versioning *Versioning `hcl:"versioning,block"`
	// Website: optional
	Website *Website `hcl:"website,block"`
}

type awsS3BucketAttributes struct {
	ref terra.Reference
}

// AccelerationStatus returns a reference to field acceleration_status of aws_s3_bucket.
func (asb awsS3BucketAttributes) AccelerationStatus() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("acceleration_status"))
}

// Acl returns a reference to field acl of aws_s3_bucket.
func (asb awsS3BucketAttributes) Acl() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("acl"))
}

// Arn returns a reference to field arn of aws_s3_bucket.
func (asb awsS3BucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket.
func (asb awsS3BucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("bucket"))
}

// BucketDomainName returns a reference to field bucket_domain_name of aws_s3_bucket.
func (asb awsS3BucketAttributes) BucketDomainName() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("bucket_domain_name"))
}

// BucketPrefix returns a reference to field bucket_prefix of aws_s3_bucket.
func (asb awsS3BucketAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("bucket_prefix"))
}

// BucketRegionalDomainName returns a reference to field bucket_regional_domain_name of aws_s3_bucket.
func (asb awsS3BucketAttributes) BucketRegionalDomainName() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("bucket_regional_domain_name"))
}

// ForceDestroy returns a reference to field force_destroy of aws_s3_bucket.
func (asb awsS3BucketAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(asb.ref.Append("force_destroy"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_s3_bucket.
func (asb awsS3BucketAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_s3_bucket.
func (asb awsS3BucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("id"))
}

// ObjectLockEnabled returns a reference to field object_lock_enabled of aws_s3_bucket.
func (asb awsS3BucketAttributes) ObjectLockEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asb.ref.Append("object_lock_enabled"))
}

// Policy returns a reference to field policy of aws_s3_bucket.
func (asb awsS3BucketAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("policy"))
}

// Region returns a reference to field region of aws_s3_bucket.
func (asb awsS3BucketAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("region"))
}

// RequestPayer returns a reference to field request_payer of aws_s3_bucket.
func (asb awsS3BucketAttributes) RequestPayer() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("request_payer"))
}

// Tags returns a reference to field tags of aws_s3_bucket.
func (asb awsS3BucketAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_s3_bucket.
func (asb awsS3BucketAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asb.ref.Append("tags_all"))
}

// WebsiteDomain returns a reference to field website_domain of aws_s3_bucket.
func (asb awsS3BucketAttributes) WebsiteDomain() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("website_domain"))
}

// WebsiteEndpoint returns a reference to field website_endpoint of aws_s3_bucket.
func (asb awsS3BucketAttributes) WebsiteEndpoint() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("website_endpoint"))
}

func (asb awsS3BucketAttributes) CorsRule() terra.ListValue[CorsRuleAttributes] {
	return terra.ReferenceAsList[CorsRuleAttributes](asb.ref.Append("cors_rule"))
}

func (asb awsS3BucketAttributes) Grant() terra.SetValue[GrantAttributes] {
	return terra.ReferenceAsSet[GrantAttributes](asb.ref.Append("grant"))
}

func (asb awsS3BucketAttributes) LifecycleRule() terra.ListValue[LifecycleRuleAttributes] {
	return terra.ReferenceAsList[LifecycleRuleAttributes](asb.ref.Append("lifecycle_rule"))
}

func (asb awsS3BucketAttributes) Logging() terra.ListValue[LoggingAttributes] {
	return terra.ReferenceAsList[LoggingAttributes](asb.ref.Append("logging"))
}

func (asb awsS3BucketAttributes) ObjectLockConfiguration() terra.ListValue[ObjectLockConfigurationAttributes] {
	return terra.ReferenceAsList[ObjectLockConfigurationAttributes](asb.ref.Append("object_lock_configuration"))
}

func (asb awsS3BucketAttributes) ReplicationConfiguration() terra.ListValue[ReplicationConfigurationAttributes] {
	return terra.ReferenceAsList[ReplicationConfigurationAttributes](asb.ref.Append("replication_configuration"))
}

func (asb awsS3BucketAttributes) ServerSideEncryptionConfiguration() terra.ListValue[ServerSideEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[ServerSideEncryptionConfigurationAttributes](asb.ref.Append("server_side_encryption_configuration"))
}

func (asb awsS3BucketAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asb.ref.Append("timeouts"))
}

func (asb awsS3BucketAttributes) Versioning() terra.ListValue[VersioningAttributes] {
	return terra.ReferenceAsList[VersioningAttributes](asb.ref.Append("versioning"))
}

func (asb awsS3BucketAttributes) Website() terra.ListValue[WebsiteAttributes] {
	return terra.ReferenceAsList[WebsiteAttributes](asb.ref.Append("website"))
}

type awsS3BucketState struct {
	AccelerationStatus                string                                   `json:"acceleration_status"`
	Acl                               string                                   `json:"acl"`
	Arn                               string                                   `json:"arn"`
	Bucket                            string                                   `json:"bucket"`
	BucketDomainName                  string                                   `json:"bucket_domain_name"`
	BucketPrefix                      string                                   `json:"bucket_prefix"`
	BucketRegionalDomainName          string                                   `json:"bucket_regional_domain_name"`
	ForceDestroy                      bool                                     `json:"force_destroy"`
	HostedZoneId                      string                                   `json:"hosted_zone_id"`
	Id                                string                                   `json:"id"`
	ObjectLockEnabled                 bool                                     `json:"object_lock_enabled"`
	Policy                            string                                   `json:"policy"`
	Region                            string                                   `json:"region"`
	RequestPayer                      string                                   `json:"request_payer"`
	Tags                              map[string]string                        `json:"tags"`
	TagsAll                           map[string]string                        `json:"tags_all"`
	WebsiteDomain                     string                                   `json:"website_domain"`
	WebsiteEndpoint                   string                                   `json:"website_endpoint"`
	CorsRule                          []CorsRuleState                          `json:"cors_rule"`
	Grant                             []GrantState                             `json:"grant"`
	LifecycleRule                     []LifecycleRuleState                     `json:"lifecycle_rule"`
	Logging                           []LoggingState                           `json:"logging"`
	ObjectLockConfiguration           []ObjectLockConfigurationState           `json:"object_lock_configuration"`
	ReplicationConfiguration          []ReplicationConfigurationState          `json:"replication_configuration"`
	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationState `json:"server_side_encryption_configuration"`
	Timeouts                          *TimeoutsState                           `json:"timeouts"`
	Versioning                        []VersioningState                        `json:"versioning"`
	Website                           []WebsiteState                           `json:"website"`
}
