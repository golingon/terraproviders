// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDbClusterSnapshot creates a new instance of [DataDbClusterSnapshot].
func NewDataDbClusterSnapshot(name string, args DataDbClusterSnapshotArgs) *DataDbClusterSnapshot {
	return &DataDbClusterSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbClusterSnapshot)(nil)

// DataDbClusterSnapshot represents the Terraform data resource aws_db_cluster_snapshot.
type DataDbClusterSnapshot struct {
	Name string
	Args DataDbClusterSnapshotArgs
}

// DataSource returns the Terraform object type for [DataDbClusterSnapshot].
func (dcs *DataDbClusterSnapshot) DataSource() string {
	return "aws_db_cluster_snapshot"
}

// LocalName returns the local name for [DataDbClusterSnapshot].
func (dcs *DataDbClusterSnapshot) LocalName() string {
	return dcs.Name
}

// Configuration returns the configuration (args) for [DataDbClusterSnapshot].
func (dcs *DataDbClusterSnapshot) Configuration() interface{} {
	return dcs.Args
}

// Attributes returns the attributes for [DataDbClusterSnapshot].
func (dcs *DataDbClusterSnapshot) Attributes() dataDbClusterSnapshotAttributes {
	return dataDbClusterSnapshotAttributes{ref: terra.ReferenceDataResource(dcs)}
}

// DataDbClusterSnapshotArgs contains the configurations for aws_db_cluster_snapshot.
type DataDbClusterSnapshotArgs struct {
	// DbClusterIdentifier: string, optional
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr"`
	// DbClusterSnapshotIdentifier: string, optional
	DbClusterSnapshotIdentifier terra.StringValue `hcl:"db_cluster_snapshot_identifier,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludePublic: bool, optional
	IncludePublic terra.BoolValue `hcl:"include_public,attr"`
	// IncludeShared: bool, optional
	IncludeShared terra.BoolValue `hcl:"include_shared,attr"`
	// MostRecent: bool, optional
	MostRecent terra.BoolValue `hcl:"most_recent,attr"`
	// SnapshotType: string, optional
	SnapshotType terra.StringValue `hcl:"snapshot_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataDbClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(dcs.ref.Append("allocated_storage"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("id"))
}

// IncludePublic returns a reference to field include_public of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) IncludePublic() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("include_public"))
}

// IncludeShared returns a reference to field include_shared of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) IncludeShared() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("include_shared"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("license_model"))
}

// MostRecent returns a reference to field most_recent of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("most_recent"))
}

// Port returns a reference to field port of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dcs.ref.Append("port"))
}

// SnapshotCreateTime returns a reference to field snapshot_create_time of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) SnapshotCreateTime() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("snapshot_create_time"))
}

// SnapshotType returns a reference to field snapshot_type of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcs.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_db_cluster_snapshot.
func (dcs dataDbClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("vpc_id"))
}
