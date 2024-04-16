// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_bucket_accelerate_configuration

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

// Resource represents the Terraform resource aws_s3_bucket_accelerate_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3BucketAccelerateConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asbac *Resource) Type() string {
	return "aws_s3_bucket_accelerate_configuration"
}

// LocalName returns the local name for [Resource].
func (asbac *Resource) LocalName() string {
	return asbac.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asbac *Resource) Configuration() interface{} {
	return asbac.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asbac *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asbac)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asbac *Resource) Dependencies() terra.Dependencies {
	return asbac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asbac *Resource) LifecycleManagement() *terra.Lifecycle {
	return asbac.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asbac *Resource) Attributes() awsS3BucketAccelerateConfigurationAttributes {
	return awsS3BucketAccelerateConfigurationAttributes{ref: terra.ReferenceResource(asbac)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asbac *Resource) ImportState(state io.Reader) error {
	asbac.state = &awsS3BucketAccelerateConfigurationState{}
	if err := json.NewDecoder(state).Decode(asbac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asbac.Type(), asbac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asbac *Resource) State() (*awsS3BucketAccelerateConfigurationState, bool) {
	return asbac.state, asbac.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asbac *Resource) StateMust() *awsS3BucketAccelerateConfigurationState {
	if asbac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asbac.Type(), asbac.LocalName()))
	}
	return asbac.state
}

// Args contains the configurations for aws_s3_bucket_accelerate_configuration.
type Args struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
}

type awsS3BucketAccelerateConfigurationAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_accelerate_configuration.
func (asbac awsS3BucketAccelerateConfigurationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asbac.ref.Append("bucket"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_bucket_accelerate_configuration.
func (asbac awsS3BucketAccelerateConfigurationAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(asbac.ref.Append("expected_bucket_owner"))
}

// Id returns a reference to field id of aws_s3_bucket_accelerate_configuration.
func (asbac awsS3BucketAccelerateConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asbac.ref.Append("id"))
}

// Status returns a reference to field status of aws_s3_bucket_accelerate_configuration.
func (asbac awsS3BucketAccelerateConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(asbac.ref.Append("status"))
}

type awsS3BucketAccelerateConfigurationState struct {
	Bucket              string `json:"bucket"`
	ExpectedBucketOwner string `json:"expected_bucket_owner"`
	Id                  string `json:"id"`
	Status              string `json:"status"`
}
