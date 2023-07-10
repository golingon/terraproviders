// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ebssnapshot "github.com/golingon/terraproviders/aws/5.7.0/ebssnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEbsSnapshot creates a new instance of [EbsSnapshot].
func NewEbsSnapshot(name string, args EbsSnapshotArgs) *EbsSnapshot {
	return &EbsSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EbsSnapshot)(nil)

// EbsSnapshot represents the Terraform resource aws_ebs_snapshot.
type EbsSnapshot struct {
	Name      string
	Args      EbsSnapshotArgs
	state     *ebsSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EbsSnapshot].
func (es *EbsSnapshot) Type() string {
	return "aws_ebs_snapshot"
}

// LocalName returns the local name for [EbsSnapshot].
func (es *EbsSnapshot) LocalName() string {
	return es.Name
}

// Configuration returns the configuration (args) for [EbsSnapshot].
func (es *EbsSnapshot) Configuration() interface{} {
	return es.Args
}

// DependOn is used for other resources to depend on [EbsSnapshot].
func (es *EbsSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(es)
}

// Dependencies returns the list of resources [EbsSnapshot] depends_on.
func (es *EbsSnapshot) Dependencies() terra.Dependencies {
	return es.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EbsSnapshot].
func (es *EbsSnapshot) LifecycleManagement() *terra.Lifecycle {
	return es.Lifecycle
}

// Attributes returns the attributes for [EbsSnapshot].
func (es *EbsSnapshot) Attributes() ebsSnapshotAttributes {
	return ebsSnapshotAttributes{ref: terra.ReferenceResource(es)}
}

// ImportState imports the given attribute values into [EbsSnapshot]'s state.
func (es *EbsSnapshot) ImportState(av io.Reader) error {
	es.state = &ebsSnapshotState{}
	if err := json.NewDecoder(av).Decode(es.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", es.Type(), es.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EbsSnapshot] has state.
func (es *EbsSnapshot) State() (*ebsSnapshotState, bool) {
	return es.state, es.state != nil
}

// StateMust returns the state for [EbsSnapshot]. Panics if the state is nil.
func (es *EbsSnapshot) StateMust() *ebsSnapshotState {
	if es.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", es.Type(), es.LocalName()))
	}
	return es.state
}

// EbsSnapshotArgs contains the configurations for aws_ebs_snapshot.
type EbsSnapshotArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OutpostArn: string, optional
	OutpostArn terra.StringValue `hcl:"outpost_arn,attr"`
	// PermanentRestore: bool, optional
	PermanentRestore terra.BoolValue `hcl:"permanent_restore,attr"`
	// StorageTier: string, optional
	StorageTier terra.StringValue `hcl:"storage_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemporaryRestoreDays: number, optional
	TemporaryRestoreDays terra.NumberValue `hcl:"temporary_restore_days,attr"`
	// VolumeId: string, required
	VolumeId terra.StringValue `hcl:"volume_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ebssnapshot.Timeouts `hcl:"timeouts,block"`
}
type ebsSnapshotAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("arn"))
}

// DataEncryptionKeyId returns a reference to field data_encryption_key_id of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) DataEncryptionKeyId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("data_encryption_key_id"))
}

// Description returns a reference to field description of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("description"))
}

// Encrypted returns a reference to field encrypted of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("kms_key_id"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("outpost_arn"))
}

// OwnerAlias returns a reference to field owner_alias of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) OwnerAlias() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("owner_alias"))
}

// OwnerId returns a reference to field owner_id of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("owner_id"))
}

// PermanentRestore returns a reference to field permanent_restore of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) PermanentRestore() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("permanent_restore"))
}

// StorageTier returns a reference to field storage_tier of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) StorageTier() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("storage_tier"))
}

// Tags returns a reference to field tags of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("tags_all"))
}

// TemporaryRestoreDays returns a reference to field temporary_restore_days of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) TemporaryRestoreDays() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("temporary_restore_days"))
}

// VolumeId returns a reference to field volume_id of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("volume_id"))
}

// VolumeSize returns a reference to field volume_size of aws_ebs_snapshot.
func (es ebsSnapshotAttributes) VolumeSize() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("volume_size"))
}

func (es ebsSnapshotAttributes) Timeouts() ebssnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ebssnapshot.TimeoutsAttributes](es.ref.Append("timeouts"))
}

type ebsSnapshotState struct {
	Arn                  string                     `json:"arn"`
	DataEncryptionKeyId  string                     `json:"data_encryption_key_id"`
	Description          string                     `json:"description"`
	Encrypted            bool                       `json:"encrypted"`
	Id                   string                     `json:"id"`
	KmsKeyId             string                     `json:"kms_key_id"`
	OutpostArn           string                     `json:"outpost_arn"`
	OwnerAlias           string                     `json:"owner_alias"`
	OwnerId              string                     `json:"owner_id"`
	PermanentRestore     bool                       `json:"permanent_restore"`
	StorageTier          string                     `json:"storage_tier"`
	Tags                 map[string]string          `json:"tags"`
	TagsAll              map[string]string          `json:"tags_all"`
	TemporaryRestoreDays float64                    `json:"temporary_restore_days"`
	VolumeId             string                     `json:"volume_id"`
	VolumeSize           float64                    `json:"volume_size"`
	Timeouts             *ebssnapshot.TimeoutsState `json:"timeouts"`
}
