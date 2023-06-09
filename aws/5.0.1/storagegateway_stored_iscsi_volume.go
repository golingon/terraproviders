// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStoragegatewayStoredIscsiVolume creates a new instance of [StoragegatewayStoredIscsiVolume].
func NewStoragegatewayStoredIscsiVolume(name string, args StoragegatewayStoredIscsiVolumeArgs) *StoragegatewayStoredIscsiVolume {
	return &StoragegatewayStoredIscsiVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewayStoredIscsiVolume)(nil)

// StoragegatewayStoredIscsiVolume represents the Terraform resource aws_storagegateway_stored_iscsi_volume.
type StoragegatewayStoredIscsiVolume struct {
	Name      string
	Args      StoragegatewayStoredIscsiVolumeArgs
	state     *storagegatewayStoredIscsiVolumeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) Type() string {
	return "aws_storagegateway_stored_iscsi_volume"
}

// LocalName returns the local name for [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) LocalName() string {
	return ssiv.Name
}

// Configuration returns the configuration (args) for [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) Configuration() interface{} {
	return ssiv.Args
}

// DependOn is used for other resources to depend on [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(ssiv)
}

// Dependencies returns the list of resources [StoragegatewayStoredIscsiVolume] depends_on.
func (ssiv *StoragegatewayStoredIscsiVolume) Dependencies() terra.Dependencies {
	return ssiv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) LifecycleManagement() *terra.Lifecycle {
	return ssiv.Lifecycle
}

// Attributes returns the attributes for [StoragegatewayStoredIscsiVolume].
func (ssiv *StoragegatewayStoredIscsiVolume) Attributes() storagegatewayStoredIscsiVolumeAttributes {
	return storagegatewayStoredIscsiVolumeAttributes{ref: terra.ReferenceResource(ssiv)}
}

// ImportState imports the given attribute values into [StoragegatewayStoredIscsiVolume]'s state.
func (ssiv *StoragegatewayStoredIscsiVolume) ImportState(av io.Reader) error {
	ssiv.state = &storagegatewayStoredIscsiVolumeState{}
	if err := json.NewDecoder(av).Decode(ssiv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssiv.Type(), ssiv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StoragegatewayStoredIscsiVolume] has state.
func (ssiv *StoragegatewayStoredIscsiVolume) State() (*storagegatewayStoredIscsiVolumeState, bool) {
	return ssiv.state, ssiv.state != nil
}

// StateMust returns the state for [StoragegatewayStoredIscsiVolume]. Panics if the state is nil.
func (ssiv *StoragegatewayStoredIscsiVolume) StateMust() *storagegatewayStoredIscsiVolumeState {
	if ssiv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssiv.Type(), ssiv.LocalName()))
	}
	return ssiv.state
}

// StoragegatewayStoredIscsiVolumeArgs contains the configurations for aws_storagegateway_stored_iscsi_volume.
type StoragegatewayStoredIscsiVolumeArgs struct {
	// DiskId: string, required
	DiskId terra.StringValue `hcl:"disk_id,attr" validate:"required"`
	// GatewayArn: string, required
	GatewayArn terra.StringValue `hcl:"gateway_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsEncrypted: bool, optional
	KmsEncrypted terra.BoolValue `hcl:"kms_encrypted,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// PreserveExistingData: bool, required
	PreserveExistingData terra.BoolValue `hcl:"preserve_existing_data,attr" validate:"required"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetName: string, required
	TargetName terra.StringValue `hcl:"target_name,attr" validate:"required"`
}
type storagegatewayStoredIscsiVolumeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("arn"))
}

// ChapEnabled returns a reference to field chap_enabled of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) ChapEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssiv.ref.Append("chap_enabled"))
}

// DiskId returns a reference to field disk_id of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) DiskId() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("disk_id"))
}

// GatewayArn returns a reference to field gateway_arn of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("gateway_arn"))
}

// Id returns a reference to field id of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("id"))
}

// KmsEncrypted returns a reference to field kms_encrypted of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) KmsEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ssiv.ref.Append("kms_encrypted"))
}

// KmsKey returns a reference to field kms_key of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("kms_key"))
}

// LunNumber returns a reference to field lun_number of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) LunNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ssiv.ref.Append("lun_number"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("network_interface_id"))
}

// NetworkInterfacePort returns a reference to field network_interface_port of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) NetworkInterfacePort() terra.NumberValue {
	return terra.ReferenceAsNumber(ssiv.ref.Append("network_interface_port"))
}

// PreserveExistingData returns a reference to field preserve_existing_data of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) PreserveExistingData() terra.BoolValue {
	return terra.ReferenceAsBool(ssiv.ref.Append("preserve_existing_data"))
}

// SnapshotId returns a reference to field snapshot_id of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("snapshot_id"))
}

// Tags returns a reference to field tags of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssiv.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssiv.ref.Append("tags_all"))
}

// TargetArn returns a reference to field target_arn of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) TargetArn() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("target_arn"))
}

// TargetName returns a reference to field target_name of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) TargetName() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("target_name"))
}

// VolumeAttachmentStatus returns a reference to field volume_attachment_status of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) VolumeAttachmentStatus() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("volume_attachment_status"))
}

// VolumeId returns a reference to field volume_id of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("volume_id"))
}

// VolumeSizeInBytes returns a reference to field volume_size_in_bytes of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) VolumeSizeInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ssiv.ref.Append("volume_size_in_bytes"))
}

// VolumeStatus returns a reference to field volume_status of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) VolumeStatus() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("volume_status"))
}

// VolumeType returns a reference to field volume_type of aws_storagegateway_stored_iscsi_volume.
func (ssiv storagegatewayStoredIscsiVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(ssiv.ref.Append("volume_type"))
}

type storagegatewayStoredIscsiVolumeState struct {
	Arn                    string            `json:"arn"`
	ChapEnabled            bool              `json:"chap_enabled"`
	DiskId                 string            `json:"disk_id"`
	GatewayArn             string            `json:"gateway_arn"`
	Id                     string            `json:"id"`
	KmsEncrypted           bool              `json:"kms_encrypted"`
	KmsKey                 string            `json:"kms_key"`
	LunNumber              float64           `json:"lun_number"`
	NetworkInterfaceId     string            `json:"network_interface_id"`
	NetworkInterfacePort   float64           `json:"network_interface_port"`
	PreserveExistingData   bool              `json:"preserve_existing_data"`
	SnapshotId             string            `json:"snapshot_id"`
	Tags                   map[string]string `json:"tags"`
	TagsAll                map[string]string `json:"tags_all"`
	TargetArn              string            `json:"target_arn"`
	TargetName             string            `json:"target_name"`
	VolumeAttachmentStatus string            `json:"volume_attachment_status"`
	VolumeId               string            `json:"volume_id"`
	VolumeSizeInBytes      float64           `json:"volume_size_in_bytes"`
	VolumeStatus           string            `json:"volume_status"`
	VolumeType             string            `json:"volume_type"`
}
