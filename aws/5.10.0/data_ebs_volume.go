// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataebsvolume "github.com/golingon/terraproviders/aws/5.10.0/dataebsvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEbsVolume creates a new instance of [DataEbsVolume].
func NewDataEbsVolume(name string, args DataEbsVolumeArgs) *DataEbsVolume {
	return &DataEbsVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEbsVolume)(nil)

// DataEbsVolume represents the Terraform data resource aws_ebs_volume.
type DataEbsVolume struct {
	Name string
	Args DataEbsVolumeArgs
}

// DataSource returns the Terraform object type for [DataEbsVolume].
func (ev *DataEbsVolume) DataSource() string {
	return "aws_ebs_volume"
}

// LocalName returns the local name for [DataEbsVolume].
func (ev *DataEbsVolume) LocalName() string {
	return ev.Name
}

// Configuration returns the configuration (args) for [DataEbsVolume].
func (ev *DataEbsVolume) Configuration() interface{} {
	return ev.Args
}

// Attributes returns the attributes for [DataEbsVolume].
func (ev *DataEbsVolume) Attributes() dataEbsVolumeAttributes {
	return dataEbsVolumeAttributes{ref: terra.ReferenceDataResource(ev)}
}

// DataEbsVolumeArgs contains the configurations for aws_ebs_volume.
type DataEbsVolumeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataebsvolume.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataebsvolume.Timeouts `hcl:"timeouts,block"`
}
type dataEbsVolumeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("availability_zone"))
}

// Encrypted returns a reference to field encrypted of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("id"))
}

// Iops returns a reference to field iops of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("iops"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("kms_key_id"))
}

// MostRecent returns a reference to field most_recent of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("most_recent"))
}

// MultiAttachEnabled returns a reference to field multi_attach_enabled of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) MultiAttachEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("multi_attach_enabled"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("outpost_arn"))
}

// Size returns a reference to field size of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("size"))
}

// SnapshotId returns a reference to field snapshot_id of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("snapshot_id"))
}

// Tags returns a reference to field tags of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ev.ref.Append("tags"))
}

// Throughput returns a reference to field throughput of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) Throughput() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("throughput"))
}

// VolumeId returns a reference to field volume_id of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("volume_id"))
}

// VolumeType returns a reference to field volume_type of aws_ebs_volume.
func (ev dataEbsVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("volume_type"))
}

func (ev dataEbsVolumeAttributes) Filter() terra.SetValue[dataebsvolume.FilterAttributes] {
	return terra.ReferenceAsSet[dataebsvolume.FilterAttributes](ev.ref.Append("filter"))
}

func (ev dataEbsVolumeAttributes) Timeouts() dataebsvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataebsvolume.TimeoutsAttributes](ev.ref.Append("timeouts"))
}
