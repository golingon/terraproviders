// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketlogging "github.com/golingon/terraproviders/aws/4.66.1/s3bucketlogging"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketLogging creates a new instance of [S3BucketLogging].
func NewS3BucketLogging(name string, args S3BucketLoggingArgs) *S3BucketLogging {
	return &S3BucketLogging{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketLogging)(nil)

// S3BucketLogging represents the Terraform resource aws_s3_bucket_logging.
type S3BucketLogging struct {
	Name      string
	Args      S3BucketLoggingArgs
	state     *s3BucketLoggingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketLogging].
func (sbl *S3BucketLogging) Type() string {
	return "aws_s3_bucket_logging"
}

// LocalName returns the local name for [S3BucketLogging].
func (sbl *S3BucketLogging) LocalName() string {
	return sbl.Name
}

// Configuration returns the configuration (args) for [S3BucketLogging].
func (sbl *S3BucketLogging) Configuration() interface{} {
	return sbl.Args
}

// DependOn is used for other resources to depend on [S3BucketLogging].
func (sbl *S3BucketLogging) DependOn() terra.Reference {
	return terra.ReferenceResource(sbl)
}

// Dependencies returns the list of resources [S3BucketLogging] depends_on.
func (sbl *S3BucketLogging) Dependencies() terra.Dependencies {
	return sbl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketLogging].
func (sbl *S3BucketLogging) LifecycleManagement() *terra.Lifecycle {
	return sbl.Lifecycle
}

// Attributes returns the attributes for [S3BucketLogging].
func (sbl *S3BucketLogging) Attributes() s3BucketLoggingAttributes {
	return s3BucketLoggingAttributes{ref: terra.ReferenceResource(sbl)}
}

// ImportState imports the given attribute values into [S3BucketLogging]'s state.
func (sbl *S3BucketLogging) ImportState(av io.Reader) error {
	sbl.state = &s3BucketLoggingState{}
	if err := json.NewDecoder(av).Decode(sbl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbl.Type(), sbl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketLogging] has state.
func (sbl *S3BucketLogging) State() (*s3BucketLoggingState, bool) {
	return sbl.state, sbl.state != nil
}

// StateMust returns the state for [S3BucketLogging]. Panics if the state is nil.
func (sbl *S3BucketLogging) StateMust() *s3BucketLoggingState {
	if sbl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbl.Type(), sbl.LocalName()))
	}
	return sbl.state
}

// S3BucketLoggingArgs contains the configurations for aws_s3_bucket_logging.
type S3BucketLoggingArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetBucket: string, required
	TargetBucket terra.StringValue `hcl:"target_bucket,attr" validate:"required"`
	// TargetPrefix: string, required
	TargetPrefix terra.StringValue `hcl:"target_prefix,attr" validate:"required"`
	// TargetGrant: min=0
	TargetGrant []s3bucketlogging.TargetGrant `hcl:"target_grant,block" validate:"min=0"`
}
type s3BucketLoggingAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_logging.
func (sbl s3BucketLoggingAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbl.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_logging.
func (sbl s3BucketLoggingAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sbl.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_logging.
func (sbl s3BucketLoggingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbl.ref.Append("id"))
}

// TargetBucket returns a reference to field target_bucket of aws_s3_bucket_logging.
func (sbl s3BucketLoggingAttributes) TargetBucket() terra.StringValue {
	return terra.ReferenceAsString(sbl.ref.Append("target_bucket"))
}

// TargetPrefix returns a reference to field target_prefix of aws_s3_bucket_logging.
func (sbl s3BucketLoggingAttributes) TargetPrefix() terra.StringValue {
	return terra.ReferenceAsString(sbl.ref.Append("target_prefix"))
}

func (sbl s3BucketLoggingAttributes) TargetGrant() terra.SetValue[s3bucketlogging.TargetGrantAttributes] {
	return terra.ReferenceAsSet[s3bucketlogging.TargetGrantAttributes](sbl.ref.Append("target_grant"))
}

type s3BucketLoggingState struct {
	Bucket              string                             `json:"bucket"`
	ExpectedBucketOwner string                             `json:"expected_bucket_owner"`
	Id                  string                             `json:"id"`
	TargetBucket        string                             `json:"target_bucket"`
	TargetPrefix        string                             `json:"target_prefix"`
	TargetGrant         []s3bucketlogging.TargetGrantState `json:"target_grant"`
}
