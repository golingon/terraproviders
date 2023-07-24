// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ebssnapshotcopy "github.com/golingon/terraproviders/aws/5.9.0/ebssnapshotcopy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEbsSnapshotCopy creates a new instance of [EbsSnapshotCopy].
func NewEbsSnapshotCopy(name string, args EbsSnapshotCopyArgs) *EbsSnapshotCopy {
	return &EbsSnapshotCopy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EbsSnapshotCopy)(nil)

// EbsSnapshotCopy represents the Terraform resource aws_ebs_snapshot_copy.
type EbsSnapshotCopy struct {
	Name      string
	Args      EbsSnapshotCopyArgs
	state     *ebsSnapshotCopyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) Type() string {
	return "aws_ebs_snapshot_copy"
}

// LocalName returns the local name for [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) LocalName() string {
	return esc.Name
}

// Configuration returns the configuration (args) for [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) Configuration() interface{} {
	return esc.Args
}

// DependOn is used for other resources to depend on [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) DependOn() terra.Reference {
	return terra.ReferenceResource(esc)
}

// Dependencies returns the list of resources [EbsSnapshotCopy] depends_on.
func (esc *EbsSnapshotCopy) Dependencies() terra.Dependencies {
	return esc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) LifecycleManagement() *terra.Lifecycle {
	return esc.Lifecycle
}

// Attributes returns the attributes for [EbsSnapshotCopy].
func (esc *EbsSnapshotCopy) Attributes() ebsSnapshotCopyAttributes {
	return ebsSnapshotCopyAttributes{ref: terra.ReferenceResource(esc)}
}

// ImportState imports the given attribute values into [EbsSnapshotCopy]'s state.
func (esc *EbsSnapshotCopy) ImportState(av io.Reader) error {
	esc.state = &ebsSnapshotCopyState{}
	if err := json.NewDecoder(av).Decode(esc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", esc.Type(), esc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EbsSnapshotCopy] has state.
func (esc *EbsSnapshotCopy) State() (*ebsSnapshotCopyState, bool) {
	return esc.state, esc.state != nil
}

// StateMust returns the state for [EbsSnapshotCopy]. Panics if the state is nil.
func (esc *EbsSnapshotCopy) StateMust() *ebsSnapshotCopyState {
	if esc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", esc.Type(), esc.LocalName()))
	}
	return esc.state
}

// EbsSnapshotCopyArgs contains the configurations for aws_ebs_snapshot_copy.
type EbsSnapshotCopyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PermanentRestore: bool, optional
	PermanentRestore terra.BoolValue `hcl:"permanent_restore,attr"`
	// SourceRegion: string, required
	SourceRegion terra.StringValue `hcl:"source_region,attr" validate:"required"`
	// SourceSnapshotId: string, required
	SourceSnapshotId terra.StringValue `hcl:"source_snapshot_id,attr" validate:"required"`
	// StorageTier: string, optional
	StorageTier terra.StringValue `hcl:"storage_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemporaryRestoreDays: number, optional
	TemporaryRestoreDays terra.NumberValue `hcl:"temporary_restore_days,attr"`
	// Timeouts: optional
	Timeouts *ebssnapshotcopy.Timeouts `hcl:"timeouts,block"`
}
type ebsSnapshotCopyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("arn"))
}

// DataEncryptionKeyId returns a reference to field data_encryption_key_id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) DataEncryptionKeyId() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("data_encryption_key_id"))
}

// Description returns a reference to field description of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(esc.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("kms_key_id"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("outpost_arn"))
}

// OwnerAlias returns a reference to field owner_alias of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) OwnerAlias() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("owner_alias"))
}

// OwnerId returns a reference to field owner_id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("owner_id"))
}

// PermanentRestore returns a reference to field permanent_restore of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) PermanentRestore() terra.BoolValue {
	return terra.ReferenceAsBool(esc.ref.Append("permanent_restore"))
}

// SourceRegion returns a reference to field source_region of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) SourceRegion() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("source_region"))
}

// SourceSnapshotId returns a reference to field source_snapshot_id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) SourceSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("source_snapshot_id"))
}

// StorageTier returns a reference to field storage_tier of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) StorageTier() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("storage_tier"))
}

// Tags returns a reference to field tags of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](esc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](esc.ref.Append("tags_all"))
}

// TemporaryRestoreDays returns a reference to field temporary_restore_days of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) TemporaryRestoreDays() terra.NumberValue {
	return terra.ReferenceAsNumber(esc.ref.Append("temporary_restore_days"))
}

// VolumeId returns a reference to field volume_id of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(esc.ref.Append("volume_id"))
}

// VolumeSize returns a reference to field volume_size of aws_ebs_snapshot_copy.
func (esc ebsSnapshotCopyAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(esc.ref.Append("volume_size"))
}

func (esc ebsSnapshotCopyAttributes) Timeouts() ebssnapshotcopy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ebssnapshotcopy.TimeoutsAttributes](esc.ref.Append("timeouts"))
}

type ebsSnapshotCopyState struct {
	Arn                  string                         `json:"arn"`
	DataEncryptionKeyId  string                         `json:"data_encryption_key_id"`
	Description          string                         `json:"description"`
	Encrypted            bool                           `json:"encrypted"`
	Id                   string                         `json:"id"`
	KmsKeyId             string                         `json:"kms_key_id"`
	OutpostArn           string                         `json:"outpost_arn"`
	OwnerAlias           string                         `json:"owner_alias"`
	OwnerId              string                         `json:"owner_id"`
	PermanentRestore     bool                           `json:"permanent_restore"`
	SourceRegion         string                         `json:"source_region"`
	SourceSnapshotId     string                         `json:"source_snapshot_id"`
	StorageTier          string                         `json:"storage_tier"`
	Tags                 map[string]string              `json:"tags"`
	TagsAll              map[string]string              `json:"tags_all"`
	TemporaryRestoreDays float64                        `json:"temporary_restore_days"`
	VolumeId             string                         `json:"volume_id"`
	VolumeSize           float64                        `json:"volume_size"`
	Timeouts             *ebssnapshotcopy.TimeoutsState `json:"timeouts"`
}
