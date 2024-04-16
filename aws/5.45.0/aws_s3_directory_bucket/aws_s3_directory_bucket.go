// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_directory_bucket

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

// Resource represents the Terraform resource aws_s3_directory_bucket.
type Resource struct {
	Name      string
	Args      Args
	state     *awsS3DirectoryBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asdb *Resource) Type() string {
	return "aws_s3_directory_bucket"
}

// LocalName returns the local name for [Resource].
func (asdb *Resource) LocalName() string {
	return asdb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asdb *Resource) Configuration() interface{} {
	return asdb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asdb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asdb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asdb *Resource) Dependencies() terra.Dependencies {
	return asdb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asdb *Resource) LifecycleManagement() *terra.Lifecycle {
	return asdb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asdb *Resource) Attributes() awsS3DirectoryBucketAttributes {
	return awsS3DirectoryBucketAttributes{ref: terra.ReferenceResource(asdb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asdb *Resource) ImportState(state io.Reader) error {
	asdb.state = &awsS3DirectoryBucketState{}
	if err := json.NewDecoder(state).Decode(asdb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asdb.Type(), asdb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asdb *Resource) State() (*awsS3DirectoryBucketState, bool) {
	return asdb.state, asdb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asdb *Resource) StateMust() *awsS3DirectoryBucketState {
	if asdb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asdb.Type(), asdb.LocalName()))
	}
	return asdb.state
}

// Args contains the configurations for aws_s3_directory_bucket.
type Args struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// DataRedundancy: string, optional
	DataRedundancy terra.StringValue `hcl:"data_redundancy,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Location: min=0
	Location []Location `hcl:"location,block" validate:"min=0"`
}

type awsS3DirectoryBucketAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asdb.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asdb.ref.Append("bucket"))
}

// DataRedundancy returns a reference to field data_redundancy of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) DataRedundancy() terra.StringValue {
	return terra.ReferenceAsString(asdb.ref.Append("data_redundancy"))
}

// ForceDestroy returns a reference to field force_destroy of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(asdb.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asdb.ref.Append("id"))
}

// Type returns a reference to field type of aws_s3_directory_bucket.
func (asdb awsS3DirectoryBucketAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(asdb.ref.Append("type"))
}

func (asdb awsS3DirectoryBucketAttributes) Location() terra.ListValue[LocationAttributes] {
	return terra.ReferenceAsList[LocationAttributes](asdb.ref.Append("location"))
}

type awsS3DirectoryBucketState struct {
	Arn            string          `json:"arn"`
	Bucket         string          `json:"bucket"`
	DataRedundancy string          `json:"data_redundancy"`
	ForceDestroy   bool            `json:"force_destroy"`
	Id             string          `json:"id"`
	Type           string          `json:"type"`
	Location       []LocationState `json:"location"`
}
