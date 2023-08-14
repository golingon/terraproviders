// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ebsvolume "github.com/golingon/terraproviders/aws/5.12.0/ebsvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEbsVolume creates a new instance of [EbsVolume].
func NewEbsVolume(name string, args EbsVolumeArgs) *EbsVolume {
	return &EbsVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EbsVolume)(nil)

// EbsVolume represents the Terraform resource aws_ebs_volume.
type EbsVolume struct {
	Name      string
	Args      EbsVolumeArgs
	state     *ebsVolumeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EbsVolume].
func (ev *EbsVolume) Type() string {
	return "aws_ebs_volume"
}

// LocalName returns the local name for [EbsVolume].
func (ev *EbsVolume) LocalName() string {
	return ev.Name
}

// Configuration returns the configuration (args) for [EbsVolume].
func (ev *EbsVolume) Configuration() interface{} {
	return ev.Args
}

// DependOn is used for other resources to depend on [EbsVolume].
func (ev *EbsVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(ev)
}

// Dependencies returns the list of resources [EbsVolume] depends_on.
func (ev *EbsVolume) Dependencies() terra.Dependencies {
	return ev.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EbsVolume].
func (ev *EbsVolume) LifecycleManagement() *terra.Lifecycle {
	return ev.Lifecycle
}

// Attributes returns the attributes for [EbsVolume].
func (ev *EbsVolume) Attributes() ebsVolumeAttributes {
	return ebsVolumeAttributes{ref: terra.ReferenceResource(ev)}
}

// ImportState imports the given attribute values into [EbsVolume]'s state.
func (ev *EbsVolume) ImportState(av io.Reader) error {
	ev.state = &ebsVolumeState{}
	if err := json.NewDecoder(av).Decode(ev.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ev.Type(), ev.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EbsVolume] has state.
func (ev *EbsVolume) State() (*ebsVolumeState, bool) {
	return ev.state, ev.state != nil
}

// StateMust returns the state for [EbsVolume]. Panics if the state is nil.
func (ev *EbsVolume) StateMust() *ebsVolumeState {
	if ev.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ev.Type(), ev.LocalName()))
	}
	return ev.state
}

// EbsVolumeArgs contains the configurations for aws_ebs_volume.
type EbsVolumeArgs struct {
	// AvailabilityZone: string, required
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr" validate:"required"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// FinalSnapshot: bool, optional
	FinalSnapshot terra.BoolValue `hcl:"final_snapshot,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// MultiAttachEnabled: bool, optional
	MultiAttachEnabled terra.BoolValue `hcl:"multi_attach_enabled,attr"`
	// OutpostArn: string, optional
	OutpostArn terra.StringValue `hcl:"outpost_arn,attr"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Throughput: number, optional
	Throughput terra.NumberValue `hcl:"throughput,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *ebsvolume.Timeouts `hcl:"timeouts,block"`
}
type ebsVolumeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ebs_volume.
func (ev ebsVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_ebs_volume.
func (ev ebsVolumeAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("availability_zone"))
}

// Encrypted returns a reference to field encrypted of aws_ebs_volume.
func (ev ebsVolumeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("encrypted"))
}

// FinalSnapshot returns a reference to field final_snapshot of aws_ebs_volume.
func (ev ebsVolumeAttributes) FinalSnapshot() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("final_snapshot"))
}

// Id returns a reference to field id of aws_ebs_volume.
func (ev ebsVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("id"))
}

// Iops returns a reference to field iops of aws_ebs_volume.
func (ev ebsVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("iops"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_ebs_volume.
func (ev ebsVolumeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("kms_key_id"))
}

// MultiAttachEnabled returns a reference to field multi_attach_enabled of aws_ebs_volume.
func (ev ebsVolumeAttributes) MultiAttachEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("multi_attach_enabled"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ebs_volume.
func (ev ebsVolumeAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("outpost_arn"))
}

// Size returns a reference to field size of aws_ebs_volume.
func (ev ebsVolumeAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("size"))
}

// SnapshotId returns a reference to field snapshot_id of aws_ebs_volume.
func (ev ebsVolumeAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("snapshot_id"))
}

// Tags returns a reference to field tags of aws_ebs_volume.
func (ev ebsVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ev.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ebs_volume.
func (ev ebsVolumeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ev.ref.Append("tags_all"))
}

// Throughput returns a reference to field throughput of aws_ebs_volume.
func (ev ebsVolumeAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("throughput"))
}

// Type returns a reference to field type of aws_ebs_volume.
func (ev ebsVolumeAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("type"))
}

func (ev ebsVolumeAttributes) Timeouts() ebsvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ebsvolume.TimeoutsAttributes](ev.ref.Append("timeouts"))
}

type ebsVolumeState struct {
	Arn                string                   `json:"arn"`
	AvailabilityZone   string                   `json:"availability_zone"`
	Encrypted          bool                     `json:"encrypted"`
	FinalSnapshot      bool                     `json:"final_snapshot"`
	Id                 string                   `json:"id"`
	Iops               float64                  `json:"iops"`
	KmsKeyId           string                   `json:"kms_key_id"`
	MultiAttachEnabled bool                     `json:"multi_attach_enabled"`
	OutpostArn         string                   `json:"outpost_arn"`
	Size               float64                  `json:"size"`
	SnapshotId         string                   `json:"snapshot_id"`
	Tags               map[string]string        `json:"tags"`
	TagsAll            map[string]string        `json:"tags_all"`
	Throughput         float64                  `json:"throughput"`
	Type               string                   `json:"type"`
	Timeouts           *ebsvolume.TimeoutsState `json:"timeouts"`
}
