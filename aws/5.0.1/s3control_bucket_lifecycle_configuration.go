// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3controlbucketlifecycleconfiguration "github.com/golingon/terraproviders/aws/5.0.1/s3controlbucketlifecycleconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3ControlBucketLifecycleConfiguration creates a new instance of [S3ControlBucketLifecycleConfiguration].
func NewS3ControlBucketLifecycleConfiguration(name string, args S3ControlBucketLifecycleConfigurationArgs) *S3ControlBucketLifecycleConfiguration {
	return &S3ControlBucketLifecycleConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3ControlBucketLifecycleConfiguration)(nil)

// S3ControlBucketLifecycleConfiguration represents the Terraform resource aws_s3control_bucket_lifecycle_configuration.
type S3ControlBucketLifecycleConfiguration struct {
	Name      string
	Args      S3ControlBucketLifecycleConfigurationArgs
	state     *s3ControlBucketLifecycleConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) Type() string {
	return "aws_s3control_bucket_lifecycle_configuration"
}

// LocalName returns the local name for [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) LocalName() string {
	return sblc.Name
}

// Configuration returns the configuration (args) for [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) Configuration() interface{} {
	return sblc.Args
}

// DependOn is used for other resources to depend on [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sblc)
}

// Dependencies returns the list of resources [S3ControlBucketLifecycleConfiguration] depends_on.
func (sblc *S3ControlBucketLifecycleConfiguration) Dependencies() terra.Dependencies {
	return sblc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sblc.Lifecycle
}

// Attributes returns the attributes for [S3ControlBucketLifecycleConfiguration].
func (sblc *S3ControlBucketLifecycleConfiguration) Attributes() s3ControlBucketLifecycleConfigurationAttributes {
	return s3ControlBucketLifecycleConfigurationAttributes{ref: terra.ReferenceResource(sblc)}
}

// ImportState imports the given attribute values into [S3ControlBucketLifecycleConfiguration]'s state.
func (sblc *S3ControlBucketLifecycleConfiguration) ImportState(av io.Reader) error {
	sblc.state = &s3ControlBucketLifecycleConfigurationState{}
	if err := json.NewDecoder(av).Decode(sblc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sblc.Type(), sblc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3ControlBucketLifecycleConfiguration] has state.
func (sblc *S3ControlBucketLifecycleConfiguration) State() (*s3ControlBucketLifecycleConfigurationState, bool) {
	return sblc.state, sblc.state != nil
}

// StateMust returns the state for [S3ControlBucketLifecycleConfiguration]. Panics if the state is nil.
func (sblc *S3ControlBucketLifecycleConfiguration) StateMust() *s3ControlBucketLifecycleConfigurationState {
	if sblc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sblc.Type(), sblc.LocalName()))
	}
	return sblc.state
}

// S3ControlBucketLifecycleConfigurationArgs contains the configurations for aws_s3control_bucket_lifecycle_configuration.
type S3ControlBucketLifecycleConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Rule: min=1
	Rule []s3controlbucketlifecycleconfiguration.Rule `hcl:"rule,block" validate:"min=1"`
}
type s3ControlBucketLifecycleConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3control_bucket_lifecycle_configuration.
func (sblc s3ControlBucketLifecycleConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sblc.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3control_bucket_lifecycle_configuration.
func (sblc s3ControlBucketLifecycleConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sblc.ref.Append("id"))
}

func (sblc s3ControlBucketLifecycleConfigurationAttributes) Rule() terra.SetValue[s3controlbucketlifecycleconfiguration.RuleAttributes] {
	return terra.ReferenceAsSet[s3controlbucketlifecycleconfiguration.RuleAttributes](sblc.ref.Append("rule"))
}

type s3ControlBucketLifecycleConfigurationState struct {
	Bucket string                                            `json:"bucket"`
	Id     string                                            `json:"id"`
	Rule   []s3controlbucketlifecycleconfiguration.RuleState `json:"rule"`
}
