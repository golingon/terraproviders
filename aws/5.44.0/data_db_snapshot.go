// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataDbSnapshot creates a new instance of [DataDbSnapshot].
func NewDataDbSnapshot(name string, args DataDbSnapshotArgs) *DataDbSnapshot {
	return &DataDbSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbSnapshot)(nil)

// DataDbSnapshot represents the Terraform data resource aws_db_snapshot.
type DataDbSnapshot struct {
	Name string
	Args DataDbSnapshotArgs
}

// DataSource returns the Terraform object type for [DataDbSnapshot].
func (ds *DataDbSnapshot) DataSource() string {
	return "aws_db_snapshot"
}

// LocalName returns the local name for [DataDbSnapshot].
func (ds *DataDbSnapshot) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DataDbSnapshot].
func (ds *DataDbSnapshot) Configuration() interface{} {
	return ds.Args
}

// Attributes returns the attributes for [DataDbSnapshot].
func (ds *DataDbSnapshot) Attributes() dataDbSnapshotAttributes {
	return dataDbSnapshotAttributes{ref: terra.ReferenceDataResource(ds)}
}

// DataDbSnapshotArgs contains the configurations for aws_db_snapshot.
type DataDbSnapshotArgs struct {
	// DbInstanceIdentifier: string, optional
	DbInstanceIdentifier terra.StringValue `hcl:"db_instance_identifier,attr"`
	// DbSnapshotIdentifier: string, optional
	DbSnapshotIdentifier terra.StringValue `hcl:"db_snapshot_identifier,attr"`
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
type dataDbSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("allocated_storage"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("availability_zone"))
}

// DbInstanceIdentifier returns a reference to field db_instance_identifier of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) DbInstanceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_instance_identifier"))
}

// DbSnapshotArn returns a reference to field db_snapshot_arn of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) DbSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_snapshot_arn"))
}

// DbSnapshotIdentifier returns a reference to field db_snapshot_identifier of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) DbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_snapshot_identifier"))
}

// Encrypted returns a reference to field encrypted of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("encrypted"))
}

// Engine returns a reference to field engine of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// IncludePublic returns a reference to field include_public of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) IncludePublic() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("include_public"))
}

// IncludeShared returns a reference to field include_shared of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) IncludeShared() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("include_shared"))
}

// Iops returns a reference to field iops of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("iops"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("license_model"))
}

// MostRecent returns a reference to field most_recent of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("most_recent"))
}

// OptionGroupName returns a reference to field option_group_name of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) OptionGroupName() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("option_group_name"))
}

// OriginalSnapshotCreateTime returns a reference to field original_snapshot_create_time of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) OriginalSnapshotCreateTime() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("original_snapshot_create_time"))
}

// Port returns a reference to field port of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("port"))
}

// SnapshotCreateTime returns a reference to field snapshot_create_time of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) SnapshotCreateTime() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("snapshot_create_time"))
}

// SnapshotType returns a reference to field snapshot_type of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("snapshot_type"))
}

// SourceDbSnapshotIdentifier returns a reference to field source_db_snapshot_identifier of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) SourceDbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("source_db_snapshot_identifier"))
}

// SourceRegion returns a reference to field source_region of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) SourceRegion() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("source_region"))
}

// Status returns a reference to field status of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("status"))
}

// StorageType returns a reference to field storage_type of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_db_snapshot.
func (ds dataDbSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("vpc_id"))
}
