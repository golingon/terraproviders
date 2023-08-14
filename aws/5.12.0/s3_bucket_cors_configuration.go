// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketcorsconfiguration "github.com/golingon/terraproviders/aws/5.12.0/s3bucketcorsconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketCorsConfiguration creates a new instance of [S3BucketCorsConfiguration].
func NewS3BucketCorsConfiguration(name string, args S3BucketCorsConfigurationArgs) *S3BucketCorsConfiguration {
	return &S3BucketCorsConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketCorsConfiguration)(nil)

// S3BucketCorsConfiguration represents the Terraform resource aws_s3_bucket_cors_configuration.
type S3BucketCorsConfiguration struct {
	Name      string
	Args      S3BucketCorsConfigurationArgs
	state     *s3BucketCorsConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) Type() string {
	return "aws_s3_bucket_cors_configuration"
}

// LocalName returns the local name for [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) LocalName() string {
	return sbcc.Name
}

// Configuration returns the configuration (args) for [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) Configuration() interface{} {
	return sbcc.Args
}

// DependOn is used for other resources to depend on [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbcc)
}

// Dependencies returns the list of resources [S3BucketCorsConfiguration] depends_on.
func (sbcc *S3BucketCorsConfiguration) Dependencies() terra.Dependencies {
	return sbcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbcc.Lifecycle
}

// Attributes returns the attributes for [S3BucketCorsConfiguration].
func (sbcc *S3BucketCorsConfiguration) Attributes() s3BucketCorsConfigurationAttributes {
	return s3BucketCorsConfigurationAttributes{ref: terra.ReferenceResource(sbcc)}
}

// ImportState imports the given attribute values into [S3BucketCorsConfiguration]'s state.
func (sbcc *S3BucketCorsConfiguration) ImportState(av io.Reader) error {
	sbcc.state = &s3BucketCorsConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbcc.Type(), sbcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketCorsConfiguration] has state.
func (sbcc *S3BucketCorsConfiguration) State() (*s3BucketCorsConfigurationState, bool) {
	return sbcc.state, sbcc.state != nil
}

// StateMust returns the state for [S3BucketCorsConfiguration]. Panics if the state is nil.
func (sbcc *S3BucketCorsConfiguration) StateMust() *s3BucketCorsConfigurationState {
	if sbcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbcc.Type(), sbcc.LocalName()))
	}
	return sbcc.state
}

// S3BucketCorsConfigurationArgs contains the configurations for aws_s3_bucket_cors_configuration.
type S3BucketCorsConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// CorsRule: min=1,max=100
	CorsRule []s3bucketcorsconfiguration.CorsRule `hcl:"cors_rule,block" validate:"min=1,max=100"`
}
type s3BucketCorsConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_cors_configuration.
func (sbcc s3BucketCorsConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbcc.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_cors_configuration.
func (sbcc s3BucketCorsConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sbcc.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_cors_configuration.
func (sbcc s3BucketCorsConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbcc.ref.Append("id"))
}

func (sbcc s3BucketCorsConfigurationAttributes) CorsRule() terra.SetValue[s3bucketcorsconfiguration.CorsRuleAttributes] {
	return terra.ReferenceAsSet[s3bucketcorsconfiguration.CorsRuleAttributes](sbcc.ref.Append("cors_rule"))
}

type s3BucketCorsConfigurationState struct {
	Bucket              string                                    `json:"bucket"`
	ExpectedBucketOwner string                                    `json:"expected_bucket_owner"`
	Id                  string                                    `json:"id"`
	CorsRule            []s3bucketcorsconfiguration.CorsRuleState `json:"cors_rule"`
}
