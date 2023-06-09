// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computedisk "github.com/golingon/terraproviders/google/4.71.0/computedisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeDisk creates a new instance of [ComputeDisk].
func NewComputeDisk(name string, args ComputeDiskArgs) *ComputeDisk {
	return &ComputeDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeDisk)(nil)

// ComputeDisk represents the Terraform resource google_compute_disk.
type ComputeDisk struct {
	Name      string
	Args      ComputeDiskArgs
	state     *computeDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeDisk].
func (cd *ComputeDisk) Type() string {
	return "google_compute_disk"
}

// LocalName returns the local name for [ComputeDisk].
func (cd *ComputeDisk) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [ComputeDisk].
func (cd *ComputeDisk) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [ComputeDisk].
func (cd *ComputeDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [ComputeDisk] depends_on.
func (cd *ComputeDisk) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeDisk].
func (cd *ComputeDisk) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [ComputeDisk].
func (cd *ComputeDisk) Attributes() computeDiskAttributes {
	return computeDiskAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [ComputeDisk]'s state.
func (cd *ComputeDisk) ImportState(av io.Reader) error {
	cd.state = &computeDiskState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeDisk] has state.
func (cd *ComputeDisk) State() (*computeDiskState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [ComputeDisk]. Panics if the state is nil.
func (cd *ComputeDisk) StateMust() *computeDiskState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// ComputeDiskArgs contains the configurations for google_compute_disk.
type ComputeDiskArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, optional
	Image terra.StringValue `hcl:"image,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Licenses: list of string, optional
	Licenses terra.ListValue[terra.StringValue] `hcl:"licenses,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PhysicalBlockSizeBytes: number, optional
	PhysicalBlockSizeBytes terra.NumberValue `hcl:"physical_block_size_bytes,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProvisionedIops: number, optional
	ProvisionedIops terra.NumberValue `hcl:"provisioned_iops,attr"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// Snapshot: string, optional
	Snapshot terra.StringValue `hcl:"snapshot,attr"`
	// SourceDisk: string, optional
	SourceDisk terra.StringValue `hcl:"source_disk,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// DiskEncryptionKey: optional
	DiskEncryptionKey *computedisk.DiskEncryptionKey `hcl:"disk_encryption_key,block"`
	// GuestOsFeatures: min=0
	GuestOsFeatures []computedisk.GuestOsFeatures `hcl:"guest_os_features,block" validate:"min=0"`
	// SourceImageEncryptionKey: optional
	SourceImageEncryptionKey *computedisk.SourceImageEncryptionKey `hcl:"source_image_encryption_key,block"`
	// SourceSnapshotEncryptionKey: optional
	SourceSnapshotEncryptionKey *computedisk.SourceSnapshotEncryptionKey `hcl:"source_snapshot_encryption_key,block"`
	// Timeouts: optional
	Timeouts *computedisk.Timeouts `hcl:"timeouts,block"`
}
type computeDiskAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_disk.
func (cd computeDiskAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_disk.
func (cd computeDiskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_disk.
func (cd computeDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_disk.
func (cd computeDiskAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("image"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_disk.
func (cd computeDiskAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_disk.
func (cd computeDiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("labels"))
}

// LastAttachTimestamp returns a reference to field last_attach_timestamp of google_compute_disk.
func (cd computeDiskAttributes) LastAttachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_attach_timestamp"))
}

// LastDetachTimestamp returns a reference to field last_detach_timestamp of google_compute_disk.
func (cd computeDiskAttributes) LastDetachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_detach_timestamp"))
}

// Licenses returns a reference to field licenses of google_compute_disk.
func (cd computeDiskAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cd.ref.Append("licenses"))
}

// Name returns a reference to field name of google_compute_disk.
func (cd computeDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

// PhysicalBlockSizeBytes returns a reference to field physical_block_size_bytes of google_compute_disk.
func (cd computeDiskAttributes) PhysicalBlockSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("physical_block_size_bytes"))
}

// Project returns a reference to field project of google_compute_disk.
func (cd computeDiskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("project"))
}

// ProvisionedIops returns a reference to field provisioned_iops of google_compute_disk.
func (cd computeDiskAttributes) ProvisionedIops() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("provisioned_iops"))
}

// SelfLink returns a reference to field self_link of google_compute_disk.
func (cd computeDiskAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_disk.
func (cd computeDiskAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("size"))
}

// Snapshot returns a reference to field snapshot of google_compute_disk.
func (cd computeDiskAttributes) Snapshot() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("snapshot"))
}

// SourceDisk returns a reference to field source_disk of google_compute_disk.
func (cd computeDiskAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_disk"))
}

// SourceDiskId returns a reference to field source_disk_id of google_compute_disk.
func (cd computeDiskAttributes) SourceDiskId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_disk_id"))
}

// SourceImageId returns a reference to field source_image_id of google_compute_disk.
func (cd computeDiskAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_image_id"))
}

// SourceSnapshotId returns a reference to field source_snapshot_id of google_compute_disk.
func (cd computeDiskAttributes) SourceSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_snapshot_id"))
}

// Type returns a reference to field type of google_compute_disk.
func (cd computeDiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("type"))
}

// Users returns a reference to field users of google_compute_disk.
func (cd computeDiskAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cd.ref.Append("users"))
}

// Zone returns a reference to field zone of google_compute_disk.
func (cd computeDiskAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("zone"))
}

func (cd computeDiskAttributes) DiskEncryptionKey() terra.ListValue[computedisk.DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computedisk.DiskEncryptionKeyAttributes](cd.ref.Append("disk_encryption_key"))
}

func (cd computeDiskAttributes) GuestOsFeatures() terra.SetValue[computedisk.GuestOsFeaturesAttributes] {
	return terra.ReferenceAsSet[computedisk.GuestOsFeaturesAttributes](cd.ref.Append("guest_os_features"))
}

func (cd computeDiskAttributes) SourceImageEncryptionKey() terra.ListValue[computedisk.SourceImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computedisk.SourceImageEncryptionKeyAttributes](cd.ref.Append("source_image_encryption_key"))
}

func (cd computeDiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[computedisk.SourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computedisk.SourceSnapshotEncryptionKeyAttributes](cd.ref.Append("source_snapshot_encryption_key"))
}

func (cd computeDiskAttributes) Timeouts() computedisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computedisk.TimeoutsAttributes](cd.ref.Append("timeouts"))
}

type computeDiskState struct {
	CreationTimestamp           string                                         `json:"creation_timestamp"`
	Description                 string                                         `json:"description"`
	Id                          string                                         `json:"id"`
	Image                       string                                         `json:"image"`
	LabelFingerprint            string                                         `json:"label_fingerprint"`
	Labels                      map[string]string                              `json:"labels"`
	LastAttachTimestamp         string                                         `json:"last_attach_timestamp"`
	LastDetachTimestamp         string                                         `json:"last_detach_timestamp"`
	Licenses                    []string                                       `json:"licenses"`
	Name                        string                                         `json:"name"`
	PhysicalBlockSizeBytes      float64                                        `json:"physical_block_size_bytes"`
	Project                     string                                         `json:"project"`
	ProvisionedIops             float64                                        `json:"provisioned_iops"`
	SelfLink                    string                                         `json:"self_link"`
	Size                        float64                                        `json:"size"`
	Snapshot                    string                                         `json:"snapshot"`
	SourceDisk                  string                                         `json:"source_disk"`
	SourceDiskId                string                                         `json:"source_disk_id"`
	SourceImageId               string                                         `json:"source_image_id"`
	SourceSnapshotId            string                                         `json:"source_snapshot_id"`
	Type                        string                                         `json:"type"`
	Users                       []string                                       `json:"users"`
	Zone                        string                                         `json:"zone"`
	DiskEncryptionKey           []computedisk.DiskEncryptionKeyState           `json:"disk_encryption_key"`
	GuestOsFeatures             []computedisk.GuestOsFeaturesState             `json:"guest_os_features"`
	SourceImageEncryptionKey    []computedisk.SourceImageEncryptionKeyState    `json:"source_image_encryption_key"`
	SourceSnapshotEncryptionKey []computedisk.SourceSnapshotEncryptionKeyState `json:"source_snapshot_encryption_key"`
	Timeouts                    *computedisk.TimeoutsState                     `json:"timeouts"`
}
