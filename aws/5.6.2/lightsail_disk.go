// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailDisk creates a new instance of [LightsailDisk].
func NewLightsailDisk(name string, args LightsailDiskArgs) *LightsailDisk {
	return &LightsailDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailDisk)(nil)

// LightsailDisk represents the Terraform resource aws_lightsail_disk.
type LightsailDisk struct {
	Name      string
	Args      LightsailDiskArgs
	state     *lightsailDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailDisk].
func (ld *LightsailDisk) Type() string {
	return "aws_lightsail_disk"
}

// LocalName returns the local name for [LightsailDisk].
func (ld *LightsailDisk) LocalName() string {
	return ld.Name
}

// Configuration returns the configuration (args) for [LightsailDisk].
func (ld *LightsailDisk) Configuration() interface{} {
	return ld.Args
}

// DependOn is used for other resources to depend on [LightsailDisk].
func (ld *LightsailDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(ld)
}

// Dependencies returns the list of resources [LightsailDisk] depends_on.
func (ld *LightsailDisk) Dependencies() terra.Dependencies {
	return ld.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailDisk].
func (ld *LightsailDisk) LifecycleManagement() *terra.Lifecycle {
	return ld.Lifecycle
}

// Attributes returns the attributes for [LightsailDisk].
func (ld *LightsailDisk) Attributes() lightsailDiskAttributes {
	return lightsailDiskAttributes{ref: terra.ReferenceResource(ld)}
}

// ImportState imports the given attribute values into [LightsailDisk]'s state.
func (ld *LightsailDisk) ImportState(av io.Reader) error {
	ld.state = &lightsailDiskState{}
	if err := json.NewDecoder(av).Decode(ld.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ld.Type(), ld.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailDisk] has state.
func (ld *LightsailDisk) State() (*lightsailDiskState, bool) {
	return ld.state, ld.state != nil
}

// StateMust returns the state for [LightsailDisk]. Panics if the state is nil.
func (ld *LightsailDisk) StateMust() *lightsailDiskState {
	if ld.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ld.Type(), ld.LocalName()))
	}
	return ld.state
}

// LightsailDiskArgs contains the configurations for aws_lightsail_disk.
type LightsailDiskArgs struct {
	// AvailabilityZone: string, required
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SizeInGb: number, required
	SizeInGb terra.NumberValue `hcl:"size_in_gb,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type lightsailDiskAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_disk.
func (ld lightsailDiskAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_lightsail_disk.
func (ld lightsailDiskAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("availability_zone"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_disk.
func (ld lightsailDiskAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_disk.
func (ld lightsailDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("id"))
}

// Name returns a reference to field name of aws_lightsail_disk.
func (ld lightsailDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("name"))
}

// SizeInGb returns a reference to field size_in_gb of aws_lightsail_disk.
func (ld lightsailDiskAttributes) SizeInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ld.ref.Append("size_in_gb"))
}

// SupportCode returns a reference to field support_code of aws_lightsail_disk.
func (ld lightsailDiskAttributes) SupportCode() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("support_code"))
}

// Tags returns a reference to field tags of aws_lightsail_disk.
func (ld lightsailDiskAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ld.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_disk.
func (ld lightsailDiskAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ld.ref.Append("tags_all"))
}

type lightsailDiskState struct {
	Arn              string            `json:"arn"`
	AvailabilityZone string            `json:"availability_zone"`
	CreatedAt        string            `json:"created_at"`
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	SizeInGb         float64           `json:"size_in_gb"`
	SupportCode      string            `json:"support_code"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
}
