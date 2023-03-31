// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketreplicationconfiguration "github.com/golingon/terraproviders/aws/4.60.0/s3bucketreplicationconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketReplicationConfiguration creates a new instance of [S3BucketReplicationConfiguration].
func NewS3BucketReplicationConfiguration(name string, args S3BucketReplicationConfigurationArgs) *S3BucketReplicationConfiguration {
	return &S3BucketReplicationConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketReplicationConfiguration)(nil)

// S3BucketReplicationConfiguration represents the Terraform resource aws_s3_bucket_replication_configuration.
type S3BucketReplicationConfiguration struct {
	Name      string
	Args      S3BucketReplicationConfigurationArgs
	state     *s3BucketReplicationConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) Type() string {
	return "aws_s3_bucket_replication_configuration"
}

// LocalName returns the local name for [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) LocalName() string {
	return sbrc.Name
}

// Configuration returns the configuration (args) for [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) Configuration() interface{} {
	return sbrc.Args
}

// DependOn is used for other resources to depend on [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(sbrc)
}

// Dependencies returns the list of resources [S3BucketReplicationConfiguration] depends_on.
func (sbrc *S3BucketReplicationConfiguration) Dependencies() terra.Dependencies {
	return sbrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) LifecycleManagement() *terra.Lifecycle {
	return sbrc.Lifecycle
}

// Attributes returns the attributes for [S3BucketReplicationConfiguration].
func (sbrc *S3BucketReplicationConfiguration) Attributes() s3BucketReplicationConfigurationAttributes {
	return s3BucketReplicationConfigurationAttributes{ref: terra.ReferenceResource(sbrc)}
}

// ImportState imports the given attribute values into [S3BucketReplicationConfiguration]'s state.
func (sbrc *S3BucketReplicationConfiguration) ImportState(av io.Reader) error {
	sbrc.state = &s3BucketReplicationConfigurationState{}
	if err := json.NewDecoder(av).Decode(sbrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbrc.Type(), sbrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketReplicationConfiguration] has state.
func (sbrc *S3BucketReplicationConfiguration) State() (*s3BucketReplicationConfigurationState, bool) {
	return sbrc.state, sbrc.state != nil
}

// StateMust returns the state for [S3BucketReplicationConfiguration]. Panics if the state is nil.
func (sbrc *S3BucketReplicationConfiguration) StateMust() *s3BucketReplicationConfigurationState {
	if sbrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbrc.Type(), sbrc.LocalName()))
	}
	return sbrc.state
}

// S3BucketReplicationConfigurationArgs contains the configurations for aws_s3_bucket_replication_configuration.
type S3BucketReplicationConfigurationArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Token: string, optional
	Token terra.StringValue `hcl:"token,attr"`
	// Rule: min=1,max=1000
	Rule []s3bucketreplicationconfiguration.Rule `hcl:"rule,block" validate:"min=1,max=1000"`
}
type s3BucketReplicationConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_replication_configuration.
func (sbrc s3BucketReplicationConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbrc.ref.Append("bucket"))
}

// Id returns a reference to field id of aws_s3_bucket_replication_configuration.
func (sbrc s3BucketReplicationConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbrc.ref.Append("id"))
}

// Role returns a reference to field role of aws_s3_bucket_replication_configuration.
func (sbrc s3BucketReplicationConfigurationAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(sbrc.ref.Append("role"))
}

// Token returns a reference to field token of aws_s3_bucket_replication_configuration.
func (sbrc s3BucketReplicationConfigurationAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(sbrc.ref.Append("token"))
}

func (sbrc s3BucketReplicationConfigurationAttributes) Rule() terra.ListValue[s3bucketreplicationconfiguration.RuleAttributes] {
	return terra.ReferenceAsList[s3bucketreplicationconfiguration.RuleAttributes](sbrc.ref.Append("rule"))
}

type s3BucketReplicationConfigurationState struct {
	Bucket string                                       `json:"bucket"`
	Id     string                                       `json:"id"`
	Role   string                                       `json:"role"`
	Token  string                                       `json:"token"`
	Rule   []s3bucketreplicationconfiguration.RuleState `json:"rule"`
}