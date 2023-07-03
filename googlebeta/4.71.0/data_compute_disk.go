// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputedisk "github.com/golingon/terraproviders/googlebeta/4.71.0/datacomputedisk"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeDisk creates a new instance of [DataComputeDisk].
func NewDataComputeDisk(name string, args DataComputeDiskArgs) *DataComputeDisk {
	return &DataComputeDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeDisk)(nil)

// DataComputeDisk represents the Terraform data resource google_compute_disk.
type DataComputeDisk struct {
	Name string
	Args DataComputeDiskArgs
}

// DataSource returns the Terraform object type for [DataComputeDisk].
func (cd *DataComputeDisk) DataSource() string {
	return "google_compute_disk"
}

// LocalName returns the local name for [DataComputeDisk].
func (cd *DataComputeDisk) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [DataComputeDisk].
func (cd *DataComputeDisk) Configuration() interface{} {
	return cd.Args
}

// Attributes returns the attributes for [DataComputeDisk].
func (cd *DataComputeDisk) Attributes() dataComputeDiskAttributes {
	return dataComputeDiskAttributes{ref: terra.ReferenceDataResource(cd)}
}

// DataComputeDiskArgs contains the configurations for google_compute_disk.
type DataComputeDiskArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AsyncPrimaryDisk: min=0
	AsyncPrimaryDisk []datacomputedisk.AsyncPrimaryDisk `hcl:"async_primary_disk,block" validate:"min=0"`
	// DiskEncryptionKey: min=0
	DiskEncryptionKey []datacomputedisk.DiskEncryptionKey `hcl:"disk_encryption_key,block" validate:"min=0"`
	// GuestOsFeatures: min=0
	GuestOsFeatures []datacomputedisk.GuestOsFeatures `hcl:"guest_os_features,block" validate:"min=0"`
	// SourceImageEncryptionKey: min=0
	SourceImageEncryptionKey []datacomputedisk.SourceImageEncryptionKey `hcl:"source_image_encryption_key,block" validate:"min=0"`
	// SourceSnapshotEncryptionKey: min=0
	SourceSnapshotEncryptionKey []datacomputedisk.SourceSnapshotEncryptionKey `hcl:"source_snapshot_encryption_key,block" validate:"min=0"`
}
type dataComputeDiskAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_disk.
func (cd dataComputeDiskAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_disk.
func (cd dataComputeDiskAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_disk.
func (cd dataComputeDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_disk.
func (cd dataComputeDiskAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("image"))
}

// Interface returns a reference to field interface of google_compute_disk.
func (cd dataComputeDiskAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("interface"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_disk.
func (cd dataComputeDiskAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_disk.
func (cd dataComputeDiskAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("labels"))
}

// LastAttachTimestamp returns a reference to field last_attach_timestamp of google_compute_disk.
func (cd dataComputeDiskAttributes) LastAttachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_attach_timestamp"))
}

// LastDetachTimestamp returns a reference to field last_detach_timestamp of google_compute_disk.
func (cd dataComputeDiskAttributes) LastDetachTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_detach_timestamp"))
}

// Licenses returns a reference to field licenses of google_compute_disk.
func (cd dataComputeDiskAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cd.ref.Append("licenses"))
}

// MultiWriter returns a reference to field multi_writer of google_compute_disk.
func (cd dataComputeDiskAttributes) MultiWriter() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("multi_writer"))
}

// Name returns a reference to field name of google_compute_disk.
func (cd dataComputeDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

// PhysicalBlockSizeBytes returns a reference to field physical_block_size_bytes of google_compute_disk.
func (cd dataComputeDiskAttributes) PhysicalBlockSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("physical_block_size_bytes"))
}

// Project returns a reference to field project of google_compute_disk.
func (cd dataComputeDiskAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("project"))
}

// ProvisionedIops returns a reference to field provisioned_iops of google_compute_disk.
func (cd dataComputeDiskAttributes) ProvisionedIops() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("provisioned_iops"))
}

// ResourcePolicies returns a reference to field resource_policies of google_compute_disk.
func (cd dataComputeDiskAttributes) ResourcePolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cd.ref.Append("resource_policies"))
}

// SelfLink returns a reference to field self_link of google_compute_disk.
func (cd dataComputeDiskAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_disk.
func (cd dataComputeDiskAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("size"))
}

// Snapshot returns a reference to field snapshot of google_compute_disk.
func (cd dataComputeDiskAttributes) Snapshot() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("snapshot"))
}

// SourceDisk returns a reference to field source_disk of google_compute_disk.
func (cd dataComputeDiskAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_disk"))
}

// SourceDiskId returns a reference to field source_disk_id of google_compute_disk.
func (cd dataComputeDiskAttributes) SourceDiskId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_disk_id"))
}

// SourceImageId returns a reference to field source_image_id of google_compute_disk.
func (cd dataComputeDiskAttributes) SourceImageId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_image_id"))
}

// SourceSnapshotId returns a reference to field source_snapshot_id of google_compute_disk.
func (cd dataComputeDiskAttributes) SourceSnapshotId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("source_snapshot_id"))
}

// Type returns a reference to field type of google_compute_disk.
func (cd dataComputeDiskAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("type"))
}

// Users returns a reference to field users of google_compute_disk.
func (cd dataComputeDiskAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cd.ref.Append("users"))
}

// Zone returns a reference to field zone of google_compute_disk.
func (cd dataComputeDiskAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("zone"))
}

func (cd dataComputeDiskAttributes) AsyncPrimaryDisk() terra.ListValue[datacomputedisk.AsyncPrimaryDiskAttributes] {
	return terra.ReferenceAsList[datacomputedisk.AsyncPrimaryDiskAttributes](cd.ref.Append("async_primary_disk"))
}

func (cd dataComputeDiskAttributes) DiskEncryptionKey() terra.ListValue[datacomputedisk.DiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[datacomputedisk.DiskEncryptionKeyAttributes](cd.ref.Append("disk_encryption_key"))
}

func (cd dataComputeDiskAttributes) GuestOsFeatures() terra.SetValue[datacomputedisk.GuestOsFeaturesAttributes] {
	return terra.ReferenceAsSet[datacomputedisk.GuestOsFeaturesAttributes](cd.ref.Append("guest_os_features"))
}

func (cd dataComputeDiskAttributes) SourceImageEncryptionKey() terra.ListValue[datacomputedisk.SourceImageEncryptionKeyAttributes] {
	return terra.ReferenceAsList[datacomputedisk.SourceImageEncryptionKeyAttributes](cd.ref.Append("source_image_encryption_key"))
}

func (cd dataComputeDiskAttributes) SourceSnapshotEncryptionKey() terra.ListValue[datacomputedisk.SourceSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[datacomputedisk.SourceSnapshotEncryptionKeyAttributes](cd.ref.Append("source_snapshot_encryption_key"))
}
