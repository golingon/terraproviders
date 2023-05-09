// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxopenzfsvolume "github.com/golingon/terraproviders/aws/4.66.1/fsxopenzfsvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxOpenzfsVolume creates a new instance of [FsxOpenzfsVolume].
func NewFsxOpenzfsVolume(name string, args FsxOpenzfsVolumeArgs) *FsxOpenzfsVolume {
	return &FsxOpenzfsVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxOpenzfsVolume)(nil)

// FsxOpenzfsVolume represents the Terraform resource aws_fsx_openzfs_volume.
type FsxOpenzfsVolume struct {
	Name      string
	Args      FsxOpenzfsVolumeArgs
	state     *fsxOpenzfsVolumeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) Type() string {
	return "aws_fsx_openzfs_volume"
}

// LocalName returns the local name for [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) LocalName() string {
	return fov.Name
}

// Configuration returns the configuration (args) for [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) Configuration() interface{} {
	return fov.Args
}

// DependOn is used for other resources to depend on [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(fov)
}

// Dependencies returns the list of resources [FsxOpenzfsVolume] depends_on.
func (fov *FsxOpenzfsVolume) Dependencies() terra.Dependencies {
	return fov.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) LifecycleManagement() *terra.Lifecycle {
	return fov.Lifecycle
}

// Attributes returns the attributes for [FsxOpenzfsVolume].
func (fov *FsxOpenzfsVolume) Attributes() fsxOpenzfsVolumeAttributes {
	return fsxOpenzfsVolumeAttributes{ref: terra.ReferenceResource(fov)}
}

// ImportState imports the given attribute values into [FsxOpenzfsVolume]'s state.
func (fov *FsxOpenzfsVolume) ImportState(av io.Reader) error {
	fov.state = &fsxOpenzfsVolumeState{}
	if err := json.NewDecoder(av).Decode(fov.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fov.Type(), fov.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxOpenzfsVolume] has state.
func (fov *FsxOpenzfsVolume) State() (*fsxOpenzfsVolumeState, bool) {
	return fov.state, fov.state != nil
}

// StateMust returns the state for [FsxOpenzfsVolume]. Panics if the state is nil.
func (fov *FsxOpenzfsVolume) StateMust() *fsxOpenzfsVolumeState {
	if fov.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fov.Type(), fov.LocalName()))
	}
	return fov.state
}

// FsxOpenzfsVolumeArgs contains the configurations for aws_fsx_openzfs_volume.
type FsxOpenzfsVolumeArgs struct {
	// CopyTagsToSnapshots: bool, optional
	CopyTagsToSnapshots terra.BoolValue `hcl:"copy_tags_to_snapshots,attr"`
	// DataCompressionType: string, optional
	DataCompressionType terra.StringValue `hcl:"data_compression_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentVolumeId: string, required
	ParentVolumeId terra.StringValue `hcl:"parent_volume_id,attr" validate:"required"`
	// ReadOnly: bool, optional
	ReadOnly terra.BoolValue `hcl:"read_only,attr"`
	// RecordSizeKib: number, optional
	RecordSizeKib terra.NumberValue `hcl:"record_size_kib,attr"`
	// StorageCapacityQuotaGib: number, optional
	StorageCapacityQuotaGib terra.NumberValue `hcl:"storage_capacity_quota_gib,attr"`
	// StorageCapacityReservationGib: number, optional
	StorageCapacityReservationGib terra.NumberValue `hcl:"storage_capacity_reservation_gib,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
	// NfsExports: optional
	NfsExports *fsxopenzfsvolume.NfsExports `hcl:"nfs_exports,block"`
	// OriginSnapshot: optional
	OriginSnapshot *fsxopenzfsvolume.OriginSnapshot `hcl:"origin_snapshot,block"`
	// Timeouts: optional
	Timeouts *fsxopenzfsvolume.Timeouts `hcl:"timeouts,block"`
	// UserAndGroupQuotas: min=0,max=100
	UserAndGroupQuotas []fsxopenzfsvolume.UserAndGroupQuotas `hcl:"user_and_group_quotas,block" validate:"min=0,max=100"`
}
type fsxOpenzfsVolumeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("arn"))
}

// CopyTagsToSnapshots returns a reference to field copy_tags_to_snapshots of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) CopyTagsToSnapshots() terra.BoolValue {
	return terra.ReferenceAsBool(fov.ref.Append("copy_tags_to_snapshots"))
}

// DataCompressionType returns a reference to field data_compression_type of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) DataCompressionType() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("data_compression_type"))
}

// Id returns a reference to field id of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("id"))
}

// Name returns a reference to field name of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("name"))
}

// ParentVolumeId returns a reference to field parent_volume_id of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) ParentVolumeId() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("parent_volume_id"))
}

// ReadOnly returns a reference to field read_only of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) ReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(fov.ref.Append("read_only"))
}

// RecordSizeKib returns a reference to field record_size_kib of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) RecordSizeKib() terra.NumberValue {
	return terra.ReferenceAsNumber(fov.ref.Append("record_size_kib"))
}

// StorageCapacityQuotaGib returns a reference to field storage_capacity_quota_gib of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) StorageCapacityQuotaGib() terra.NumberValue {
	return terra.ReferenceAsNumber(fov.ref.Append("storage_capacity_quota_gib"))
}

// StorageCapacityReservationGib returns a reference to field storage_capacity_reservation_gib of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) StorageCapacityReservationGib() terra.NumberValue {
	return terra.ReferenceAsNumber(fov.ref.Append("storage_capacity_reservation_gib"))
}

// Tags returns a reference to field tags of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fov.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fov.ref.Append("tags_all"))
}

// VolumeType returns a reference to field volume_type of aws_fsx_openzfs_volume.
func (fov fsxOpenzfsVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("volume_type"))
}

func (fov fsxOpenzfsVolumeAttributes) NfsExports() terra.ListValue[fsxopenzfsvolume.NfsExportsAttributes] {
	return terra.ReferenceAsList[fsxopenzfsvolume.NfsExportsAttributes](fov.ref.Append("nfs_exports"))
}

func (fov fsxOpenzfsVolumeAttributes) OriginSnapshot() terra.ListValue[fsxopenzfsvolume.OriginSnapshotAttributes] {
	return terra.ReferenceAsList[fsxopenzfsvolume.OriginSnapshotAttributes](fov.ref.Append("origin_snapshot"))
}

func (fov fsxOpenzfsVolumeAttributes) Timeouts() fsxopenzfsvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxopenzfsvolume.TimeoutsAttributes](fov.ref.Append("timeouts"))
}

func (fov fsxOpenzfsVolumeAttributes) UserAndGroupQuotas() terra.SetValue[fsxopenzfsvolume.UserAndGroupQuotasAttributes] {
	return terra.ReferenceAsSet[fsxopenzfsvolume.UserAndGroupQuotasAttributes](fov.ref.Append("user_and_group_quotas"))
}

type fsxOpenzfsVolumeState struct {
	Arn                           string                                     `json:"arn"`
	CopyTagsToSnapshots           bool                                       `json:"copy_tags_to_snapshots"`
	DataCompressionType           string                                     `json:"data_compression_type"`
	Id                            string                                     `json:"id"`
	Name                          string                                     `json:"name"`
	ParentVolumeId                string                                     `json:"parent_volume_id"`
	ReadOnly                      bool                                       `json:"read_only"`
	RecordSizeKib                 float64                                    `json:"record_size_kib"`
	StorageCapacityQuotaGib       float64                                    `json:"storage_capacity_quota_gib"`
	StorageCapacityReservationGib float64                                    `json:"storage_capacity_reservation_gib"`
	Tags                          map[string]string                          `json:"tags"`
	TagsAll                       map[string]string                          `json:"tags_all"`
	VolumeType                    string                                     `json:"volume_type"`
	NfsExports                    []fsxopenzfsvolume.NfsExportsState         `json:"nfs_exports"`
	OriginSnapshot                []fsxopenzfsvolume.OriginSnapshotState     `json:"origin_snapshot"`
	Timeouts                      *fsxopenzfsvolume.TimeoutsState            `json:"timeouts"`
	UserAndGroupQuotas            []fsxopenzfsvolume.UserAndGroupQuotasState `json:"user_and_group_quotas"`
}
