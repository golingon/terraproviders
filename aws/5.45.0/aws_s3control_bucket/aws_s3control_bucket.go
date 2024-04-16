// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3control_bucket

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

// Resource represents the Terraform resource aws_s3control_bucket.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3ControlBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asb *Resource) Type() string {
	return "aws_s3control_bucket"
}

// LocalName returns the local name for [Resource].
func (asb *Resource) LocalName() string {
	return asb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asb *Resource) Configuration() interface{} {
	return asb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asb *Resource) Dependencies() terra.Dependencies {
	return asb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asb *Resource) LifecycleManagement() *terra.Lifecycle {
	return asb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asb *Resource) Attributes() awsS3ControlBucketAttributes {
	return awsS3ControlBucketAttributes{ref: terra.ReferenceResource(asb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asb *Resource) ImportState(state io.Reader) error {
	asb.state = &awsS3ControlBucketState{}
	if err := json.NewDecoder(state).Decode(asb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asb.Type(), asb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asb *Resource) State() (*awsS3ControlBucketState, bool) {
	return asb.state, asb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asb *Resource) StateMust() *awsS3ControlBucketState {
	if asb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asb.Type(), asb.LocalName()))
	}
	return asb.state
}

// Args contains the configurations for aws_s3control_bucket.
type Args struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OutpostId: string, required
	OutpostId terra.StringValue `hcl:"outpost_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsS3ControlBucketAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("bucket"))
}

// CreationDate returns a reference to field creation_date of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("creation_date"))
}

// Id returns a reference to field id of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("id"))
}

// OutpostId returns a reference to field outpost_id of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) OutpostId() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("outpost_id"))
}

// PublicAccessBlockEnabled returns a reference to field public_access_block_enabled of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) PublicAccessBlockEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asb.ref.Append("public_access_block_enabled"))
}

// Tags returns a reference to field tags of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_s3control_bucket.
func (asb awsS3ControlBucketAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asb.ref.Append("tags_all"))
}

type awsS3ControlBucketState struct {
	Arn                      string            `json:"arn"`
	Bucket                   string            `json:"bucket"`
	CreationDate             string            `json:"creation_date"`
	Id                       string            `json:"id"`
	OutpostId                string            `json:"outpost_id"`
	PublicAccessBlockEnabled bool              `json:"public_access_block_enabled"`
	Tags                     map[string]string `json:"tags"`
	TagsAll                  map[string]string `json:"tags_all"`
}
