// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_cluster_snapshot

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_db_cluster_snapshot.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adcs *DataSource) DataSource() string {
	return "aws_db_cluster_snapshot"
}

// LocalName returns the local name for [DataSource].
func (adcs *DataSource) LocalName() string {
	return adcs.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adcs *DataSource) Configuration() interface{} {
	return adcs.Args
}

// Attributes returns the attributes for [DataSource].
func (adcs *DataSource) Attributes() dataAwsDbClusterSnapshotAttributes {
	return dataAwsDbClusterSnapshotAttributes{ref: terra.ReferenceDataSource(adcs)}
}

// DataArgs contains the configurations for aws_db_cluster_snapshot.
type DataArgs struct {
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

type dataAwsDbClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(adcs.ref.Append("allocated_storage"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](adcs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("id"))
}

// IncludePublic returns a reference to field include_public of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) IncludePublic() terra.BoolValue {
	return terra.ReferenceAsBool(adcs.ref.Append("include_public"))
}

// IncludeShared returns a reference to field include_shared of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) IncludeShared() terra.BoolValue {
	return terra.ReferenceAsBool(adcs.ref.Append("include_shared"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("license_model"))
}

// MostRecent returns a reference to field most_recent of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(adcs.ref.Append("most_recent"))
}

// Port returns a reference to field port of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(adcs.ref.Append("port"))
}

// SnapshotCreateTime returns a reference to field snapshot_create_time of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) SnapshotCreateTime() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("snapshot_create_time"))
}

// SnapshotType returns a reference to field snapshot_type of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(adcs.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adcs.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_db_cluster_snapshot.
func (adcs dataAwsDbClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(adcs.ref.Append("vpc_id"))
}
