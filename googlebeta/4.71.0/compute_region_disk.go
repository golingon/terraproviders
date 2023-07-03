// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregiondisk "github.com/golingon/terraproviders/googlebeta/4.71.0/computeregiondisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionDisk creates a new instance of [ComputeRegionDisk].
func NewComputeRegionDisk(name string, args ComputeRegionDiskArgs) *ComputeRegionDisk {
	return &ComputeRegionDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDisk)(nil)

// ComputeRegionDisk represents the Terraform resource google_compute_region_disk.
type ComputeRegionDisk struct {
	Name      string
	Args      ComputeRegionDiskArgs
	state     *computeRegionDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionDisk].
func (crd *ComputeRegionDisk) Type() string {
	return "google_compute_region_disk"
}

// LocalName returns the local name for [ComputeRegionDisk].
func (crd *ComputeRegionDisk) LocalName() string {
	return crd.Name
}

// Configuration returns the configuration (args) for [ComputeRegionDisk].
func (crd *ComputeRegionDisk) Configuration() interface{} {
	return crd.Args
}

// DependOn is used for other resources to depend on [ComputeRegionDisk].
func (crd *ComputeRegionDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(crd)
}

// Dependencies returns the list of resources [ComputeRegionDisk] depends_on.
func (crd *ComputeRegionDisk) Dependencies() terra.Dependencies {
	return crd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionDisk].
func (crd *ComputeRegionDisk) LifecycleManagement() *terra.Lifecycle {
	return crd.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionDisk].
func (crd *ComputeRegionDisk) Attributes() computeRegionDiskAttributes {
	return computeRegionDiskAttributes{ref: terra.ReferenceResource(crd)}
}

// ImportState imports the given attribute values into [ComputeRegionDisk]'s state.
func (crd *ComputeRegionDisk) ImportState(av io.Reader) error {
	crd.state = &computeRegionDiskState{}
	if err := json.NewDecoder(av).Decode(crd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crd.Type(), crd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionDisk] has state.
func (crd *ComputeRegionDisk) State() (*computeRegionDiskState, bool) {
	return crd.state, crd.state != nil
}

// StateMust returns the state for [ComputeRegionDisk]. Panics if the state is nil.
func (crd *ComputeRegionDisk) StateMust() *computeRegionDiskState {
	if crd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crd.Type(), crd.LocalName()))
	}
	return crd.state
}

// ComputeRegionDiskArgs contains the configurations for google_compute_region_disk.
type ComputeRegionDiskArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interface: string, optional
	Interface terra.StringValue `hcl:"interface,attr"`
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ReplicaZones: list of string, required
	ReplicaZones terra.ListValue[terra.StringValue] `hcl:"replica_zones,attr" validate:"required"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// Snapshot: string, optional
	Snapshot terra.StringValue `hcl:"snapshot,attr"`
	// SourceDisk: string, optional
	SourceDisk terra.StringValue `hcl:"source_disk,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// AsyncPrimaryDisk: optional
	AsyncPrimaryDisk *computeregiondisk.AsyncPrimaryDisk `hcl:"async_primary_disk,block"`
	// DiskEncryptionKey: optional
	DiskEncryptionKey *computeregiondisk.DiskEncryptionKey `hcl:"disk_encryption_key,block"`
	// GuestOsFeatures: min=0
	GuestOsFeatures []computeregiondisk.GuestOsFeatures `hcl:"guest_os_features,block" validate:"min=0"`
	// SourceSnapshotEncryptionKey: optional
	SourceSnapshotEncryptionKey *computeregiondisk.SourceSnapshotEncryptionKey `hcl:"source_snapshot_encryption_key,block"`
	// Timeouts: optional
	Timeouts *computeregiondisk.Timeouts `hcl:"timeouts,block"`
}
type computeRegionDiskAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_disk.
func (crd computeRegionDiskAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("id"))
}

// Interface returns a reference to field interface of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("interface"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_region_disk.
func (crd computeRegionDiskAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crd.ref.Append("labels"))
}

// LastAttachTimestamp returns a reference to field last_attach_timestamp of google_compute_region_disk.
func (crd computeRegionDiskAttributes) LastAttachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("last_attach_timestamp"))
}

// LastDetachTimestamp returns a reference to field last_detach_timestamp of google_compute_region_disk.
func (crd computeRegionDiskAttributes) LastDetachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("last_detach_timestamp"))
}

// Licenses returns a reference to field licenses of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crd.ref.Append("licenses"))
}

// Name returns a reference to field name of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("name"))
}

// PhysicalBlockSizeBytes returns a reference to field physical_block_size_bytes of google_compute_region_disk.
func (crd computeRegionDiskAttributes) PhysicalBlockSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(crd.ref.Append("physical_block_size_bytes"))
}

// Project returns a reference to field project of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("region"))
}

// ReplicaZones returns a reference to field replica_zones of google_compute_region_disk.
func (crd computeRegionDiskAttributes) ReplicaZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crd.ref.Append("replica_zones"))
}

// SelfLink returns a reference to field self_link of google_compute_region_disk.
func (crd computeRegionDiskAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(crd.ref.Append("size"))
}

// Snapshot returns a reference to field snapshot of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Snapshot() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("snapshot"))
}

// SourceDisk returns a reference to field source_disk of google_compute_region_disk.
func (crd computeRegionDiskAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("source_disk"))
}

// SourceDiskId returns a reference to field source_disk_id of google_compute_region_disk.
func (crd computeRegionDiskAttributes) SourceDiskId() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("source_disk_id"))
}

// SourceSnapshotId returns a reference to field source_snapshot_id of google_compute_region_disk.
func (crd computeRegionDiskAttributes) SourceSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("source_snapshot_id"))
}

// Type returns a reference to field type of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(crd.ref.Append("type"))
}

// Users returns a reference to field users of google_compute_region_disk.
func (crd computeRegionDiskAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crd.ref.Append("users"))
}

func (crd computeRegionDiskAttributes) AsyncPrimaryDisk() terra.ListValue[computeregiondisk.AsyncPrimaryDiskAttributes] {
	return terra.ReferenceAsList[computeregiondisk.AsyncPrimaryDiskAttributes](crd.ref.Append("async_primary_disk"))
}

func (crd computeRegionDiskAttributes) DiskEncryptionKey() terra.ListValue[computeregiondisk.DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computeregiondisk.DiskEncryptionKeyAttributes](crd.ref.Append("disk_encryption_key"))
}

func (crd computeRegionDiskAttributes) GuestOsFeatures() terra.SetValue[computeregiondisk.GuestOsFeaturesAttributes] {
	return terra.ReferenceAsSet[computeregiondisk.GuestOsFeaturesAttributes](crd.ref.Append("guest_os_features"))
}

func (crd computeRegionDiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[computeregiondisk.SourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[computeregiondisk.SourceSnapshotEncryptionKeyAttributes](crd.ref.Append("source_snapshot_encryption_key"))
}

func (crd computeRegionDiskAttributes) Timeouts() computeregiondisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregiondisk.TimeoutsAttributes](crd.ref.Append("timeouts"))
}

type computeRegionDiskState struct {
	CreationTimestamp           string                                               `json:"creation_timestamp"`
	Description                 string                                               `json:"description"`
	Id                          string                                               `json:"id"`
	Interface                   string                                               `json:"interface"`
	LabelFingerprint            string                                               `json:"label_fingerprint"`
	Labels                      map[string]string                                    `json:"labels"`
	LastAttachTimestamp         string                                               `json:"last_attach_timestamp"`
	LastDetachTimestamp         string                                               `json:"last_detach_timestamp"`
	Licenses                    []string                                             `json:"licenses"`
	Name                        string                                               `json:"name"`
	PhysicalBlockSizeBytes      float64                                              `json:"physical_block_size_bytes"`
	Project                     string                                               `json:"project"`
	Region                      string                                               `json:"region"`
	ReplicaZones                []string                                             `json:"replica_zones"`
	SelfLink                    string                                               `json:"self_link"`
	Size                        float64                                              `json:"size"`
	Snapshot                    string                                               `json:"snapshot"`
	SourceDisk                  string                                               `json:"source_disk"`
	SourceDiskId                string                                               `json:"source_disk_id"`
	SourceSnapshotId            string                                               `json:"source_snapshot_id"`
	Type                        string                                               `json:"type"`
	Users                       []string                                             `json:"users"`
	AsyncPrimaryDisk            []computeregiondisk.AsyncPrimaryDiskState            `json:"async_primary_disk"`
	DiskEncryptionKey           []computeregiondisk.DiskEncryptionKeyState           `json:"disk_encryption_key"`
	GuestOsFeatures             []computeregiondisk.GuestOsFeaturesState             `json:"guest_os_features"`
	SourceSnapshotEncryptionKey []computeregiondisk.SourceSnapshotEncryptionKeyState `json:"source_snapshot_encryption_key"`
	Timeouts                    *computeregiondisk.TimeoutsState                     `json:"timeouts"`
}
