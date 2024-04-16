// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lightsail_bucket_resource_access

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

// Resource represents the Terraform resource aws_lightsail_bucket_resource_access.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLightsailBucketResourceAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (albra *Resource) Type() string {
	return "aws_lightsail_bucket_resource_access"
}

// LocalName returns the local name for [Resource].
func (albra *Resource) LocalName() string {
	return albra.Name
}

// Configuration returns the configuration (args) for [Resource].
func (albra *Resource) Configuration() interface{} {
	return albra.Args
}

// DependOn is used for other resources to depend on [Resource].
func (albra *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(albra)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (albra *Resource) Dependencies() terra.Dependencies {
	return albra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (albra *Resource) LifecycleManagement() *terra.Lifecycle {
	return albra.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (albra *Resource) Attributes() awsLightsailBucketResourceAccessAttributes {
	return awsLightsailBucketResourceAccessAttributes{ref: terra.ReferenceResource(albra)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (albra *Resource) ImportState(state io.Reader) error {
	albra.state = &awsLightsailBucketResourceAccessState{}
	if err := json.NewDecoder(state).Decode(albra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", albra.Type(), albra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (albra *Resource) State() (*awsLightsailBucketResourceAccessState, bool) {
	return albra.state, albra.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (albra *Resource) StateMust() *awsLightsailBucketResourceAccessState {
	if albra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", albra.Type(), albra.LocalName()))
	}
	return albra.state
}

// Args contains the configurations for aws_lightsail_bucket_resource_access.
type Args struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceName: string, required
	ResourceName terra.StringValue `hcl:"resource_name,attr" validate:"required"`
}

type awsLightsailBucketResourceAccessAttributes struct {
	ref terra.Reference
}

// BucketName returns a reference to field bucket_name of aws_lightsail_bucket_resource_access.
func (albra awsLightsailBucketResourceAccessAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(albra.ref.Append("bucket_name"))
}

// Id returns a reference to field id of aws_lightsail_bucket_resource_access.
func (albra awsLightsailBucketResourceAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(albra.ref.Append("id"))
}

// ResourceName returns a reference to field resource_name of aws_lightsail_bucket_resource_access.
func (albra awsLightsailBucketResourceAccessAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceAsString(albra.ref.Append("resource_name"))
}

type awsLightsailBucketResourceAccessState struct {
	BucketName   string `json:"bucket_name"`
	Id           string `json:"id"`
	ResourceName string `json:"resource_name"`
}
