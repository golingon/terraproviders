// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketserversideencryptionconfiguration "github.com/golingon/terraproviders/aws/4.60.0/s3bucketserversideencryptionconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketServerSideEncryptionConfiguration creates a new instance of [S3BucketServerSideEncryptionConfiguration].
func NewS3BucketServerSideEncryptionConfiguration(name string, args S3BucketServerSideEncryptionConfigurationArgs) *S3BucketServerSideEncryptionConfiguration {
	return &S3BucketServerSideEncryptionConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketServerSideEncryptionConfiguration)(nil)

// S3BucketServerSideEncryptionConfiguration represents the Terraform resource aws_s3_bucket_server_side_encryption_configuration.
type S3BucketServerSideEncryptionConfiguration struct {
	Name      string
	Args      S3BucketServerSideEncryptionConfigurationArgs
	state     *s3BucketServerSideEncryptionConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) Type() string {
	return "aws_s3_bucket_server_side_encryption_configuration"
}

// LocalName returns the local name for [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) LocalName() string {
	return sbssec.Name
}

// Configuration returns the configuration (args) for [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) Configuration() interface{} {
	return sbssec.Args
}

// DependOn is used for other resources to depend on [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbssec)
}

// Dependencies returns the list of resources [S3BucketServerSideEncryptionConfiguration] depends_on.
func (sbssec *S3BucketServerSideEncryptionConfiguration) Dependencies() terra.Dependencies {
	return sbssec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbssec.Lifecycle
}

// Attributes returns the attributes for [S3BucketServerSideEncryptionConfiguration].
func (sbssec *S3BucketServerSideEncryptionConfiguration) Attributes() s3BucketServerSideEncryptionConfigurationAttributes {
	return s3BucketServerSideEncryptionConfigurationAttributes{ref: terra.ReferenceResource(sbssec)}
}

// ImportState imports the given attribute values into [S3BucketServerSideEncryptionConfiguration]'s state.
func (sbssec *S3BucketServerSideEncryptionConfiguration) ImportState(av io.Reader) error {
	sbssec.state = &s3BucketServerSideEncryptionConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbssec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbssec.Type(), sbssec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketServerSideEncryptionConfiguration] has state.
func (sbssec *S3BucketServerSideEncryptionConfiguration) State() (*s3BucketServerSideEncryptionConfigurationState, bool) {
	return sbssec.state, sbssec.state != nil
}

// StateMust returns the state for [S3BucketServerSideEncryptionConfiguration]. Panics if the state is nil.
func (sbssec *S3BucketServerSideEncryptionConfiguration) StateMust() *s3BucketServerSideEncryptionConfigurationState {
	if sbssec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbssec.Type(), sbssec.LocalName()))
	}
	return sbssec.state
}

// S3BucketServerSideEncryptionConfigurationArgs contains the configurations for aws_s3_bucket_server_side_encryption_configuration.
type S3BucketServerSideEncryptionConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Rule: min=1
	Rule []s3bucketserversideencryptionconfiguration.Rule `hcl:"rule,block" validate:"min=1"`
}
type s3BucketServerSideEncryptionConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_server_side_encryption_configuration.
func (sbssec s3BucketServerSideEncryptionConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbssec.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_server_side_encryption_configuration.
func (sbssec s3BucketServerSideEncryptionConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(sbssec.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_server_side_encryption_configuration.
func (sbssec s3BucketServerSideEncryptionConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbssec.ref.Append("id"))
}

func (sbssec s3BucketServerSideEncryptionConfigurationAttributes) Rule() terra.SetValue[s3bucketserversideencryptionconfiguration.RuleAttributes] {
	return terra.ReferenceAsSet[s3bucketserversideencryptionconfiguration.RuleAttributes](sbssec.ref.Append("rule"))
}

type s3BucketServerSideEncryptionConfigurationState struct {
	Bucket              string                                                `json:"bucket"`
	ExpectedBucketOwner string                                                `json:"expected_bucket_owner"`
	Id                  string                                                `json:"id"`
	Rule                []s3bucketserversideencryptionconfiguration.RuleState `json:"rule"`
}
