// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_snapshot

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_compute_snapshot.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gcs *DataSource) DataSource() string {
	return "google_compute_snapshot"
}

// LocalName returns the local name for [DataSource].
func (gcs *DataSource) LocalName() string {
	return gcs.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gcs *DataSource) Configuration() interface{} {
	return gcs.Args
}

// Attributes returns the attributes for [DataSource].
func (gcs *DataSource) Attributes() dataGoogleComputeSnapshotAttributes {
	return dataGoogleComputeSnapshotAttributes{ref: terra.ReferenceDataSource(gcs)}
}

// DataArgs contains the configurations for google_compute_snapshot.
type DataArgs struct {
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
}

type dataGoogleComputeSnapshotAttributes struct {
	ref terra.Reference
}

// ChainName returns a reference to field chain_name of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) ChainName() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("chain_name"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("description"))
}

// DiskSizeGb returns a reference to field disk_size_gb of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(gcs.ref.Append("disk_size_gb"))
}

// EffectiveLabels returns a reference to field effective_labels of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcs.ref.Append("effective_labels"))
}

// Filter returns a reference to field filter of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("filter"))
}

// Id returns a reference to field id of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("id"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcs.ref.Append("labels"))
}

// Licenses returns a reference to field licenses of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Licenses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcs.ref.Append("licenses"))
}

// MostRecent returns a reference to field most_recent of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(gcs.ref.Append("most_recent"))
}

// Name returns a reference to field name of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("self_link"))
}

// SnapshotId returns a reference to field snapshot_id of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) SnapshotId() terra.NumberValue {
	return terra.ReferenceAsNumber(gcs.ref.Append("snapshot_id"))
}

// SourceDisk returns a reference to field source_disk of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) SourceDisk() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("source_disk"))
}

// StorageBytes returns a reference to field storage_bytes of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) StorageBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(gcs.ref.Append("storage_bytes"))
}

// StorageLocations returns a reference to field storage_locations of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) StorageLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcs.ref.Append("storage_locations"))
}

// TerraformLabels returns a reference to field terraform_labels of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcs.ref.Append("terraform_labels"))
}

// Zone returns a reference to field zone of google_compute_snapshot.
func (gcs dataGoogleComputeSnapshotAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcs.ref.Append("zone"))
}

func (gcs dataGoogleComputeSnapshotAttributes) SnapshotEncryptionKey() terra.ListValue[DataSnapshotEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DataSnapshotEncryptionKeyAttributes](gcs.ref.Append("snapshot_encryption_key"))
}

func (gcs dataGoogleComputeSnapshotAttributes) SourceDiskEncryptionKey() terra.ListValue[DataSourceDiskEncryptionKeyAttributes] {
	return terra.ReferenceAsList[DataSourceDiskEncryptionKeyAttributes](gcs.ref.Append("source_disk_encryption_key"))
}
