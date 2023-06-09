// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucket "github.com/golingon/terraproviders/aws/4.63.0/s3bucket"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3Bucket creates a new instance of [S3Bucket].
func NewS3Bucket(name string, args S3BucketArgs) *S3Bucket {
	return &S3Bucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3Bucket)(nil)

// S3Bucket represents the Terraform resource aws_s3_bucket.
type S3Bucket struct {
	Name      string
	Args      S3BucketArgs
	state     *s3BucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3Bucket].
func (sb *S3Bucket) Type() string {
	return "aws_s3_bucket"
}

// LocalName returns the local name for [S3Bucket].
func (sb *S3Bucket) LocalName() string {
	return sb.Name
}

// Configuration returns the configuration (args) for [S3Bucket].
func (sb *S3Bucket) Configuration() interface{} {
	return sb.Args
}

// DependOn is used for other resources to depend on [S3Bucket].
func (sb *S3Bucket) DependOn() terra.Reference {
	return terra.ReferenceResource(sb)
}

// Dependencies returns the list of resources [S3Bucket] depends_on.
func (sb *S3Bucket) Dependencies() terra.Dependencies {
	return sb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3Bucket].
func (sb *S3Bucket) LifecycleManagement() *terra.Lifecycle {
	return sb.Lifecycle
}

// Attributes returns the attributes for [S3Bucket].
func (sb *S3Bucket) Attributes() s3BucketAttributes {
	return s3BucketAttributes{ref: terra.ReferenceResource(sb)}
}

// ImportState imports the given attribute values into [S3Bucket]'s state.
func (sb *S3Bucket) ImportState(av io.Reader) error {
	sb.state = &s3BucketState{}
	if err := json.NewDecoder(av).Decode(sb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sb.Type(), sb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3Bucket] has state.
func (sb *S3Bucket) State() (*s3BucketState, bool) {
	return sb.state, sb.state != nil
}

// StateMust returns the state for [S3Bucket]. Panics if the state is nil.
func (sb *S3Bucket) StateMust() *s3BucketState {
	if sb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sb.Type(), sb.LocalName()))
	}
	return sb.state
}

// S3BucketArgs contains the configurations for aws_s3_bucket.
type S3BucketArgs struct {
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
	CorsRule []s3bucket.CorsRule `hcl:"cors_rule,block" validate:"min=0"`
	// Grant: min=0
	Grant []s3bucket.Grant `hcl:"grant,block" validate:"min=0"`
	// LifecycleRule: min=0
	LifecycleRule []s3bucket.LifecycleRule `hcl:"lifecycle_rule,block" validate:"min=0"`
	// Logging: optional
	Logging *s3bucket.Logging `hcl:"logging,block"`
	// ObjectLockConfiguration: optional
	ObjectLockConfiguration *s3bucket.ObjectLockConfiguration `hcl:"object_lock_configuration,block"`
	// ReplicationConfiguration: optional
	ReplicationConfiguration *s3bucket.ReplicationConfiguration `hcl:"replication_configuration,block"`
	// ServerSideEncryptionConfiguration: optional
	ServerSideEncryptionConfiguration *s3bucket.ServerSideEncryptionConfiguration `hcl:"server_side_encryption_configuration,block"`
	// Timeouts: optional
	Timeouts *s3bucket.Timeouts `hcl:"timeouts,block"`
	// Versioning: optional
	Versioning *s3bucket.Versioning `hcl:"versioning,block"`
	// Website: optional
	Website *s3bucket.Website `hcl:"website,block"`
}
type s3BucketAttributes struct {
	ref terra.Reference
}

// AccelerationStatus returns a reference to field acceleration_status of aws_s3_bucket.
func (sb s3BucketAttributes) AccelerationStatus() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("acceleration_status"))
}

// Acl returns a reference to field acl of aws_s3_bucket.
func (sb s3BucketAttributes) Acl() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("acl"))
}

// Arn returns a reference to field arn of aws_s3_bucket.
func (sb s3BucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket.
func (sb s3BucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket"))
}

// BucketDomainName returns a reference to field bucket_domain_name of aws_s3_bucket.
func (sb s3BucketAttributes) BucketDomainName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_domain_name"))
}

// BucketPrefix returns a reference to field bucket_prefix of aws_s3_bucket.
func (sb s3BucketAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_prefix"))
}

// BucketRegionalDomainName returns a reference to field bucket_regional_domain_name of aws_s3_bucket.
func (sb s3BucketAttributes) BucketRegionalDomainName() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("bucket_regional_domain_name"))
}

// ForceDestroy returns a reference to field force_destroy of aws_s3_bucket.
func (sb s3BucketAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("force_destroy"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_s3_bucket.
func (sb s3BucketAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_s3_bucket.
func (sb s3BucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("id"))
}

// ObjectLockEnabled returns a reference to field object_lock_enabled of aws_s3_bucket.
func (sb s3BucketAttributes) ObjectLockEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sb.ref.Append("object_lock_enabled"))
}

// Policy returns a reference to field policy of aws_s3_bucket.
func (sb s3BucketAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("policy"))
}

// Region returns a reference to field region of aws_s3_bucket.
func (sb s3BucketAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("region"))
}

// RequestPayer returns a reference to field request_payer of aws_s3_bucket.
func (sb s3BucketAttributes) RequestPayer() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("request_payer"))
}

// Tags returns a reference to field tags of aws_s3_bucket.
func (sb s3BucketAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_s3_bucket.
func (sb s3BucketAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sb.ref.Append("tags_all"))
}

// WebsiteDomain returns a reference to field website_domain of aws_s3_bucket.
func (sb s3BucketAttributes) WebsiteDomain() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("website_domain"))
}

// WebsiteEndpoint returns a reference to field website_endpoint of aws_s3_bucket.
func (sb s3BucketAttributes) WebsiteEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sb.ref.Append("website_endpoint"))
}

func (sb s3BucketAttributes) CorsRule() terra.ListValue[s3bucket.CorsRuleAttributes] {
	return terra.ReferenceAsList[s3bucket.CorsRuleAttributes](sb.ref.Append("cors_rule"))
}

func (sb s3BucketAttributes) Grant() terra.SetValue[s3bucket.GrantAttributes] {
	return terra.ReferenceAsSet[s3bucket.GrantAttributes](sb.ref.Append("grant"))
}

func (sb s3BucketAttributes) LifecycleRule() terra.ListValue[s3bucket.LifecycleRuleAttributes] {
	return terra.ReferenceAsList[s3bucket.LifecycleRuleAttributes](sb.ref.Append("lifecycle_rule"))
}

func (sb s3BucketAttributes) Logging() terra.ListValue[s3bucket.LoggingAttributes] {
	return terra.ReferenceAsList[s3bucket.LoggingAttributes](sb.ref.Append("logging"))
}

func (sb s3BucketAttributes) ObjectLockConfiguration() terra.ListValue[s3bucket.ObjectLockConfigurationAttributes] {
	return terra.ReferenceAsList[s3bucket.ObjectLockConfigurationAttributes](sb.ref.Append("object_lock_configuration"))
}

func (sb s3BucketAttributes) ReplicationConfiguration() terra.ListValue[s3bucket.ReplicationConfigurationAttributes] {
	return terra.ReferenceAsList[s3bucket.ReplicationConfigurationAttributes](sb.ref.Append("replication_configuration"))
}

func (sb s3BucketAttributes) ServerSideEncryptionConfiguration() terra.ListValue[s3bucket.ServerSideEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[s3bucket.ServerSideEncryptionConfigurationAttributes](sb.ref.Append("server_side_encryption_configuration"))
}

func (sb s3BucketAttributes) Timeouts() s3bucket.TimeoutsAttributes {
	return terra.ReferenceAsSingle[s3bucket.TimeoutsAttributes](sb.ref.Append("timeouts"))
}

func (sb s3BucketAttributes) Versioning() terra.ListValue[s3bucket.VersioningAttributes] {
	return terra.ReferenceAsList[s3bucket.VersioningAttributes](sb.ref.Append("versioning"))
}

func (sb s3BucketAttributes) Website() terra.ListValue[s3bucket.WebsiteAttributes] {
	return terra.ReferenceAsList[s3bucket.WebsiteAttributes](sb.ref.Append("website"))
}

type s3BucketState struct {
	AccelerationStatus                string                                            `json:"acceleration_status"`
	Acl                               string                                            `json:"acl"`
	Arn                               string                                            `json:"arn"`
	Bucket                            string                                            `json:"bucket"`
	BucketDomainName                  string                                            `json:"bucket_domain_name"`
	BucketPrefix                      string                                            `json:"bucket_prefix"`
	BucketRegionalDomainName          string                                            `json:"bucket_regional_domain_name"`
	ForceDestroy                      bool                                              `json:"force_destroy"`
	HostedZoneId                      string                                            `json:"hosted_zone_id"`
	Id                                string                                            `json:"id"`
	ObjectLockEnabled                 bool                                              `json:"object_lock_enabled"`
	Policy                            string                                            `json:"policy"`
	Region                            string                                            `json:"region"`
	RequestPayer                      string                                            `json:"request_payer"`
	Tags                              map[string]string                                 `json:"tags"`
	TagsAll                           map[string]string                                 `json:"tags_all"`
	WebsiteDomain                     string                                            `json:"website_domain"`
	WebsiteEndpoint                   string                                            `json:"website_endpoint"`
	CorsRule                          []s3bucket.CorsRuleState                          `json:"cors_rule"`
	Grant                             []s3bucket.GrantState                             `json:"grant"`
	LifecycleRule                     []s3bucket.LifecycleRuleState                     `json:"lifecycle_rule"`
	Logging                           []s3bucket.LoggingState                           `json:"logging"`
	ObjectLockConfiguration           []s3bucket.ObjectLockConfigurationState           `json:"object_lock_configuration"`
	ReplicationConfiguration          []s3bucket.ReplicationConfigurationState          `json:"replication_configuration"`
	ServerSideEncryptionConfiguration []s3bucket.ServerSideEncryptionConfigurationState `json:"server_side_encryption_configuration"`
	Timeouts                          *s3bucket.TimeoutsState                           `json:"timeouts"`
	Versioning                        []s3bucket.VersioningState                        `json:"versioning"`
	Website                           []s3bucket.WebsiteState                           `json:"website"`
}
