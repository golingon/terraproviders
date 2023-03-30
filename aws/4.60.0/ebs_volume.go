// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ebsvolume "github.com/golingon/terraproviders/aws/4.60.0/ebsvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEbsVolume(name string, args EbsVolumeArgs) *EbsVolume {
	return &EbsVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EbsVolume)(nil)

type EbsVolume struct {
	Name  string
	Args  EbsVolumeArgs
	state *ebsVolumeState
}

func (ev *EbsVolume) Type() string {
	return "aws_ebs_volume"
}

func (ev *EbsVolume) LocalName() string {
	return ev.Name
}

func (ev *EbsVolume) Configuration() interface{} {
	return ev.Args
}

func (ev *EbsVolume) Attributes() ebsVolumeAttributes {
	return ebsVolumeAttributes{ref: terra.ReferenceResource(ev)}
}

func (ev *EbsVolume) ImportState(av io.Reader) error {
	ev.state = &ebsVolumeState{}
	if err := json.NewDecoder(av).Decode(ev.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ev.Type(), ev.LocalName(), err)
	}
	return nil
}

func (ev *EbsVolume) State() (*ebsVolumeState, bool) {
	return ev.state, ev.state != nil
}

func (ev *EbsVolume) StateMust() *ebsVolumeState {
	if ev.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ev.Type(), ev.LocalName()))
	}
	return ev.state
}

func (ev *EbsVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(ev)
}

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
	// DependsOn contains resources that EbsVolume depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ebsVolumeAttributes struct {
	ref terra.Reference
}

func (ev ebsVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("arn"))
}

func (ev ebsVolumeAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("availability_zone"))
}

func (ev ebsVolumeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(ev.ref.Append("encrypted"))
}

func (ev ebsVolumeAttributes) FinalSnapshot() terra.BoolValue {
	return terra.ReferenceBool(ev.ref.Append("final_snapshot"))
}

func (ev ebsVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("id"))
}

func (ev ebsVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceNumber(ev.ref.Append("iops"))
}

func (ev ebsVolumeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("kms_key_id"))
}

func (ev ebsVolumeAttributes) MultiAttachEnabled() terra.BoolValue {
	return terra.ReferenceBool(ev.ref.Append("multi_attach_enabled"))
}

func (ev ebsVolumeAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("outpost_arn"))
}

func (ev ebsVolumeAttributes) Size() terra.NumberValue {
	return terra.ReferenceNumber(ev.ref.Append("size"))
}

func (ev ebsVolumeAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("snapshot_id"))
}

func (ev ebsVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ev.ref.Append("tags"))
}

func (ev ebsVolumeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ev.ref.Append("tags_all"))
}

func (ev ebsVolumeAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceNumber(ev.ref.Append("throughput"))
}

func (ev ebsVolumeAttributes) Type() terra.StringValue {
	return terra.ReferenceString(ev.ref.Append("type"))
}

func (ev ebsVolumeAttributes) Timeouts() ebsvolume.TimeoutsAttributes {
	return terra.ReferenceSingle[ebsvolume.TimeoutsAttributes](ev.ref.Append("timeouts"))
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
