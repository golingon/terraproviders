// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketobjectlockconfiguration "github.com/golingon/terraproviders/aws/5.9.0/s3bucketobjectlockconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketObjectLockConfiguration creates a new instance of [S3BucketObjectLockConfiguration].
func NewS3BucketObjectLockConfiguration(name string, args S3BucketObjectLockConfigurationArgs) *S3BucketObjectLockConfiguration {
	return &S3BucketObjectLockConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketObjectLockConfiguration)(nil)

// S3BucketObjectLockConfiguration represents the Terraform resource aws_s3_bucket_object_lock_configuration.
type S3BucketObjectLockConfiguration struct {
	Name      string
	Args      S3BucketObjectLockConfigurationArgs
	state     *s3BucketObjectLockConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) Type() string {
	return "aws_s3_bucket_object_lock_configuration"
}

// LocalName returns the local name for [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) LocalName() string {
	return sbolc.Name
}

// Configuration returns the configuration (args) for [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) Configuration() interface{} {
	return sbolc.Args
}

// DependOn is used for other resources to depend on [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbolc)
}

// Dependencies returns the list of resources [S3BucketObjectLockConfiguration] depends_on.
func (sbolc *S3BucketObjectLockConfiguration) Dependencies() terra.Dependencies {
	return sbolc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbolc.Lifecycle
}

// Attributes returns the attributes for [S3BucketObjectLockConfiguration].
func (sbolc *S3BucketObjectLockConfiguration) Attributes() s3BucketObjectLockConfigurationAttributes {
	return s3BucketObjectLockConfigurationAttributes{ref: terra.ReferenceResource(sbolc)}
}

// ImportState imports the given attribute values into [S3BucketObjectLockConfiguration]'s state.
func (sbolc *S3BucketObjectLockConfiguration) ImportState(av io.Reader) error {
	sbolc.state = &s3BucketObjectLockConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbolc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbolc.Type(), sbolc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketObjectLockConfiguration] has state.
func (sbolc *S3BucketObjectLockConfiguration) State() (*s3BucketObjectLockConfigurationState, bool) {
	return sbolc.state, sbolc.state != nil
}

// StateMust returns the state for [S3BucketObjectLockConfiguration]. Panics if the state is nil.
func (sbolc *S3BucketObjectLockConfiguration) StateMust() *s3BucketObjectLockConfigurationState {
	if sbolc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbolc.Type(), sbolc.LocalName()))
	}
	return sbolc.state
}

// S3BucketObjectLockConfigurationArgs contains the configurations for aws_s3_bucket_object_lock_configuration.
type S3BucketObjectLockConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectLockEnabled: string, optional
	ObjectLockEnabled terra.StringValue `hcl:"object_lock_enabled,attr"`
	// Token: string, optional
	Token terra.StringValue `hcl:"token,attr"`
	// Rule: optional
	Rule *s3bucketobjectlockconfiguration.Rule `hcl:"rule,block"`
}
type s3BucketObjectLockConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_object_lock_configuration.
func (sbolc s3BucketObjectLockConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbolc.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_object_lock_configuration.
func (sbolc s3BucketObjectLockConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sbolc.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_object_lock_configuration.
func (sbolc s3BucketObjectLockConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbolc.ref.Append("id"))
}

// ObjectLockEnabled returns a reference to field object_lock_enabled of aws_s3_bucket_object_lock_configuration.
func (sbolc s3BucketObjectLockConfigurationAttributes) ObjectLockEnabled() terra.StringValue {
	return terra.ReferenceAsString(sbolc.ref.Append("object_lock_enabled"))
}

// Token returns a reference to field token of aws_s3_bucket_object_lock_configuration.
func (sbolc s3BucketObjectLockConfigurationAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(sbolc.ref.Append("token"))
}

func (sbolc s3BucketObjectLockConfigurationAttributes) Rule() terra.ListValue[s3bucketobjectlockconfiguration.RuleAttributes] {
	return terra.ReferenceAsList[s3bucketobjectlockconfiguration.RuleAttributes](sbolc.ref.Append("rule"))
}

type s3BucketObjectLockConfigurationState struct {
	Bucket              string                                      `json:"bucket"`
	ExpectedBucketOwner string                                      `json:"expected_bucket_owner"`
	Id                  string                                      `json:"id"`
	ObjectLockEnabled   string                                      `json:"object_lock_enabled"`
	Token               string                                      `json:"token"`
	Rule                []s3bucketobjectlockconfiguration.RuleState `json:"rule"`
}
