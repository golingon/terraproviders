// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailBucket creates a new instance of [LightsailBucket].
func NewLightsailBucket(name string, args LightsailBucketArgs) *LightsailBucket {
	return &LightsailBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailBucket)(nil)

// LightsailBucket represents the Terraform resource aws_lightsail_bucket.
type LightsailBucket struct {
	Name      string
	Args      LightsailBucketArgs
	state     *lightsailBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailBucket].
func (lb *LightsailBucket) Type() string {
	return "aws_lightsail_bucket"
}

// LocalName returns the local name for [LightsailBucket].
func (lb *LightsailBucket) LocalName() string {
	return lb.Name
}

// Configuration returns the configuration (args) for [LightsailBucket].
func (lb *LightsailBucket) Configuration() interface{} {
	return lb.Args
}

// DependOn is used for other resources to depend on [LightsailBucket].
func (lb *LightsailBucket) DependOn() terra.Reference {
	return terra.ReferenceResource(lb)
}

// Dependencies returns the list of resources [LightsailBucket] depends_on.
func (lb *LightsailBucket) Dependencies() terra.Dependencies {
	return lb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailBucket].
func (lb *LightsailBucket) LifecycleManagement() *terra.Lifecycle {
	return lb.Lifecycle
}

// Attributes returns the attributes for [LightsailBucket].
func (lb *LightsailBucket) Attributes() lightsailBucketAttributes {
	return lightsailBucketAttributes{ref: terra.ReferenceResource(lb)}
}

// ImportState imports the given attribute values into [LightsailBucket]'s state.
func (lb *LightsailBucket) ImportState(av io.Reader) error {
	lb.state = &lightsailBucketState{}
	if err := json.NewDecoder(av).Decode(lb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lb.Type(), lb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailBucket] has state.
func (lb *LightsailBucket) State() (*lightsailBucketState, bool) {
	return lb.state, lb.state != nil
}

// StateMust returns the state for [LightsailBucket]. Panics if the state is nil.
func (lb *LightsailBucket) StateMust() *lightsailBucketState {
	if lb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lb.Type(), lb.LocalName()))
	}
	return lb.state
}

// LightsailBucketArgs contains the configurations for aws_lightsail_bucket.
type LightsailBucketArgs struct {
	// BundleId: string, required
	BundleId terra.StringValue `hcl:"bundle_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type lightsailBucketAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("availability_zone"))
}

// BundleId returns a reference to field bundle_id of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("bundle_id"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("id"))
}

// Name returns a reference to field name of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("name"))
}

// Region returns a reference to field region of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("region"))
}

// SupportCode returns a reference to field support_code of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) SupportCode() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("support_code"))
}

// Tags returns a reference to field tags of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lb.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_lightsail_bucket.
func (lb lightsailBucketAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("url"))
}

type lightsailBucketState struct {
	Arn              string            `json:"arn"`
	AvailabilityZone string            `json:"availability_zone"`
	BundleId         string            `json:"bundle_id"`
	CreatedAt        string            `json:"created_at"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	Region           string            `json:"region"`
	SupportCode      string            `json:"support_code"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
	Url              string            `json:"url"`
}
