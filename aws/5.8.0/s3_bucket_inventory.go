// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3bucketinventory "github.com/golingon/terraproviders/aws/5.8.0/s3bucketinventory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3BucketInventory creates a new instance of [S3BucketInventory].
func NewS3BucketInventory(name string, args S3BucketInventoryArgs) *S3BucketInventory {
	return &S3BucketInventory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3BucketInventory)(nil)

// S3BucketInventory represents the Terraform resource aws_s3_bucket_inventory.
type S3BucketInventory struct {
	Name      string
	Args      S3BucketInventoryArgs
	state     *s3BucketInventoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3BucketInventory].
func (sbi *S3BucketInventory) Type() string {
	return "aws_s3_bucket_inventory"
}

// LocalName returns the local name for [S3BucketInventory].
func (sbi *S3BucketInventory) LocalName() string {
	return sbi.Name
}

// Configuration returns the configuration (args) for [S3BucketInventory].
func (sbi *S3BucketInventory) Configuration() interface{} {
	return sbi.Args
}

// DependOn is used for other resources to depend on [S3BucketInventory].
func (sbi *S3BucketInventory) DependOn() terra.Reference {
	return terra.ReferenceResource(sbi)
}

// Dependencies returns the list of resources [S3BucketInventory] depends_on.
func (sbi *S3BucketInventory) Dependencies() terra.Dependencies {
	return sbi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3BucketInventory].
func (sbi *S3BucketInventory) LifecycleManagement() *terra.Lifecycle {
	return sbi.Lifecycle
}

// Attributes returns the attributes for [S3BucketInventory].
func (sbi *S3BucketInventory) Attributes() s3BucketInventoryAttributes {
	return s3BucketInventoryAttributes{ref: terra.ReferenceResource(sbi)}
}

// ImportState imports the given attribute values into [S3BucketInventory]'s state.
func (sbi *S3BucketInventory) ImportState(av io.Reader) error {
	sbi.state = &s3BucketInventoryState{}
	if err := json.NewDecoder(av).Decode(sbi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbi.Type(), sbi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3BucketInventory] has state.
func (sbi *S3BucketInventory) State() (*s3BucketInventoryState, bool) {
	return sbi.state, sbi.state != nil
}

// StateMust returns the state for [S3BucketInventory]. Panics if the state is nil.
func (sbi *S3BucketInventory) StateMust() *s3BucketInventoryState {
	if sbi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbi.Type(), sbi.LocalName()))
	}
	return sbi.state
}

// S3BucketInventoryArgs contains the configurations for aws_s3_bucket_inventory.
type S3BucketInventoryArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludedObjectVersions: string, required
	IncludedObjectVersions terra.StringValue `hcl:"included_object_versions,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OptionalFields: set of string, optional
	OptionalFields terra.SetValue[terra.StringValue] `hcl:"optional_fields,attr"`
	// Destination: required
	Destination *s3bucketinventory.Destination `hcl:"destination,block" validate:"required"`
	// Filter: optional
	Filter *s3bucketinventory.Filter `hcl:"filter,block"`
	// Schedule: required
	Schedule *s3bucketinventory.Schedule `hcl:"schedule,block" validate:"required"`
}
type s3BucketInventoryAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbi.ref.Append("bucket"))
}

// Enabled returns a reference to field enabled of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sbi.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbi.ref.Append("id"))
}

// IncludedObjectVersions returns a reference to field included_object_versions of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) IncludedObjectVersions() terra.StringValue {
	return terra.ReferenceAsString(sbi.ref.Append("included_object_versions"))
}

// Name returns a reference to field name of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbi.ref.Append("name"))
}

// OptionalFields returns a reference to field optional_fields of aws_s3_bucket_inventory.
func (sbi s3BucketInventoryAttributes) OptionalFields() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sbi.ref.Append("optional_fields"))
}

func (sbi s3BucketInventoryAttributes) Destination() terra.ListValue[s3bucketinventory.DestinationAttributes] {
	return terra.ReferenceAsList[s3bucketinventory.DestinationAttributes](sbi.ref.Append("destination"))
}

func (sbi s3BucketInventoryAttributes) Filter() terra.ListValue[s3bucketinventory.FilterAttributes] {
	return terra.ReferenceAsList[s3bucketinventory.FilterAttributes](sbi.ref.Append("filter"))
}

func (sbi s3BucketInventoryAttributes) Schedule() terra.ListValue[s3bucketinventory.ScheduleAttributes] {
	return terra.ReferenceAsList[s3bucketinventory.ScheduleAttributes](sbi.ref.Append("schedule"))
}

type s3BucketInventoryState struct {
	Bucket                 string                               `json:"bucket"`
	Enabled                bool                                 `json:"enabled"`
	Id                     string                               `json:"id"`
	IncludedObjectVersions string                               `json:"included_object_versions"`
	Name                   string                               `json:"name"`
	OptionalFields         []string                             `json:"optional_fields"`
	Destination            []s3bucketinventory.DestinationState `json:"destination"`
	Filter                 []s3bucketinventory.FilterState      `json:"filter"`
	Schedule               []s3bucketinventory.ScheduleState    `json:"schedule"`
}
