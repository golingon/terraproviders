// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacomputesnapshot "github.com/golingon/terraproviders/googlebeta/4.77.0/datacomputesnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeSnapshot creates a new instance of [DataComputeSnapshot].
func NewDataComputeSnapshot(name string, args DataComputeSnapshotArgs) *DataComputeSnapshot {
	return &DataComputeSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeSnapshot)(nil)

// DataComputeSnapshot represents the Terraform data resource google_compute_snapshot.
type DataComputeSnapshot struct {
	Name string
	Args DataComputeSnapshotArgs
}

// DataSource returns the Terraform object type for [DataComputeSnapshot].
func (cs *DataComputeSnapshot) DataSource() string {
	return "google_compute_snapshot"
}

// LocalName returns the local name for [DataComputeSnapshot].
func (cs *DataComputeSnapshot) LocalName() string {
	return cs.Name
}

// Configuration returns the configuration (args) for [DataComputeSnapshot].
func (cs *DataComputeSnapshot) Configuration() interface{} {
	return cs.Args
}

// Attributes returns the attributes for [DataComputeSnapshot].
func (cs *DataComputeSnapshot) Attributes() dataComputeSnapshotAttributes {
	return dataComputeSnapshotAttributes{ref: terra.ReferenceDataResource(cs)}
}

// DataComputeSnapshotArgs contains the configurations for google_compute_snapshot.
type DataComputeSnapshotArgs struct {
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SnapshotEncryptionKey: min=0
	SnapshotEncryptionKey []datacomputesnapshot.SnapshotEncryptionKey `hcl:"snapshot_encryption_key,block" validate:"min=0"`
	// SourceDiskEncryptionKey: min=0
	SourceDiskEncryptionKey []datacomputesnapshot.SourceDiskEncryptionKey `hcl:"source_disk_encryption_key,block" validate:"min=0"`
}
type dataComputeSnapshotAttributes struct {
	ref terra.Reference
}

// ChainName returns a reference to field chain_name of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) ChainName() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("chain_name"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("description"))
}

// DiskSizeGb returns a reference to field disk_size_gb of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("disk_size_gb"))
}

// Filter returns a reference to field filter of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("labels"))
}

// Licenses returns a reference to field licenses of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cs.ref.Append("licenses"))
}

// MostRecent returns a reference to field most_recent of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("most_recent"))
}

// Name returns a reference to field name of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("self_link"))
}

// SnapshotId returns a reference to field snapshot_id of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) SnapshotId() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("snapshot_id"))
}

// SourceDisk returns a reference to field source_disk of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("source_disk"))
}

// StorageBytes returns a reference to field storage_bytes of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) StorageBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("storage_bytes"))
}

// StorageLocations returns a reference to field storage_locations of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) StorageLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cs.ref.Append("storage_locations"))
}

// Zone returns a reference to field zone of google_compute_snapshot.
func (cs dataComputeSnapshotAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("zone"))
}

func (cs dataComputeSnapshotAttributes) SnapshotEncryptionKey() terra.ListValue[datacomputesnapshot.SnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[datacomputesnapshot.SnapshotEncryptionKeyAttributes](cs.ref.Append("snapshot_encryption_key"))
}

func (cs dataComputeSnapshotAttributes) SourceDiskEncryptionKey() terra.ListValue[datacomputesnapshot.SourceDiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[datacomputesnapshot.SourceDiskEncryptionKeyAttributes](cs.ref.Append("source_disk_encryption_key"))
}
