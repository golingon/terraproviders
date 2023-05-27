// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketlifecycleconfiguration "github.com/golingon/terraproviders/aws/5.0.1/s3bucketlifecycleconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketLifecycleConfiguration creates a new instance of [S3BucketLifecycleConfiguration].
func NewS3BucketLifecycleConfiguration(name string, args S3BucketLifecycleConfigurationArgs) *S3BucketLifecycleConfiguration {
	return &S3BucketLifecycleConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketLifecycleConfiguration)(nil)

// S3BucketLifecycleConfiguration represents the Terraform resource aws_s3_bucket_lifecycle_configuration.
type S3BucketLifecycleConfiguration struct {
	Name      string
	Args      S3BucketLifecycleConfigurationArgs
	state     *s3BucketLifecycleConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) Type() string {
	return "aws_s3_bucket_lifecycle_configuration"
}

// LocalName returns the local name for [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) LocalName() string {
	return sblc.Name
}

// Configuration returns the configuration (args) for [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) Configuration() interface{} {
	return sblc.Args
}

// DependOn is used for other resources to depend on [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sblc)
}

// Dependencies returns the list of resources [S3BucketLifecycleConfiguration] depends_on.
func (sblc *S3BucketLifecycleConfiguration) Dependencies() terra.Dependencies {
	return sblc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sblc.Lifecycle
}

// Attributes returns the attributes for [S3BucketLifecycleConfiguration].
func (sblc *S3BucketLifecycleConfiguration) Attributes() s3BucketLifecycleConfigurationAttributes {
	return s3BucketLifecycleConfigurationAttributes{ref: terra.ReferenceResource(sblc)}
}

// ImportState imports the given attribute values into [S3BucketLifecycleConfiguration]'s state.
func (sblc *S3BucketLifecycleConfiguration) ImportState(av io.Reader) error {
	sblc.state = &s3BucketLifecycleConfigurationState{}
	if err := json.NewDecoder(av).Decode(sblc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sblc.Type(), sblc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketLifecycleConfiguration] has state.
func (sblc *S3BucketLifecycleConfiguration) State() (*s3BucketLifecycleConfigurationState, bool) {
	return sblc.state, sblc.state != nil
}

// StateMust returns the state for [S3BucketLifecycleConfiguration]. Panics if the state is nil.
func (sblc *S3BucketLifecycleConfiguration) StateMust() *s3BucketLifecycleConfigurationState {
	if sblc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sblc.Type(), sblc.LocalName()))
	}
	return sblc.state
}

// S3BucketLifecycleConfigurationArgs contains the configurations for aws_s3_bucket_lifecycle_configuration.
type S3BucketLifecycleConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Rule: min=1
	Rule []s3bucketlifecycleconfiguration.Rule `hcl:"rule,block" validate:"min=1"`
}
type s3BucketLifecycleConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_lifecycle_configuration.
func (sblc s3BucketLifecycleConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sblc.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_lifecycle_configuration.
func (sblc s3BucketLifecycleConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sblc.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_lifecycle_configuration.
func (sblc s3BucketLifecycleConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sblc.ref.Append("id"))
}

func (sblc s3BucketLifecycleConfigurationAttributes) Rule() terra.ListValue[s3bucketlifecycleconfiguration.RuleAttributes] {
	return terra.ReferenceAsList[s3bucketlifecycleconfiguration.RuleAttributes](sblc.ref.Append("rule"))
}

type s3BucketLifecycleConfigurationState struct {
	Bucket              string                                     `json:"bucket"`
	ExpectedBucketOwner string                                     `json:"expected_bucket_owner"`
	Id                  string                                     `json:"id"`
	Rule                []s3bucketlifecycleconfiguration.RuleState `json:"rule"`
}
