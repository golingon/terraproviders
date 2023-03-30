// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDbSnapshot(name string, args DataDbSnapshotArgs) *DataDbSnapshot {
	return &DataDbSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbSnapshot)(nil)

type DataDbSnapshot struct {
	Name string
	Args DataDbSnapshotArgs
}

func (ds *DataDbSnapshot) DataSource() string {
	return "aws_db_snapshot"
}

func (ds *DataDbSnapshot) LocalName() string {
	return ds.Name
}

func (ds *DataDbSnapshot) Configuration() interface{} {
	return ds.Args
}

func (ds *DataDbSnapshot) Attributes() dataDbSnapshotAttributes {
	return dataDbSnapshotAttributes{ref: terra.ReferenceDataResource(ds)}
}

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
}
type dataDbSnapshotAttributes struct {
	ref terra.Reference
}

func (ds dataDbSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceNumber(ds.ref.Append("allocated_storage"))
}

func (ds dataDbSnapshotAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("availability_zone"))
}

func (ds dataDbSnapshotAttributes) DbInstanceIdentifier() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("db_instance_identifier"))
}

func (ds dataDbSnapshotAttributes) DbSnapshotArn() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("db_snapshot_arn"))
}

func (ds dataDbSnapshotAttributes) DbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("db_snapshot_identifier"))
}

func (ds dataDbSnapshotAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(ds.ref.Append("encrypted"))
}

func (ds dataDbSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("engine"))
}

func (ds dataDbSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("engine_version"))
}

func (ds dataDbSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("id"))
}

func (ds dataDbSnapshotAttributes) IncludePublic() terra.BoolValue {
	return terra.ReferenceBool(ds.ref.Append("include_public"))
}

func (ds dataDbSnapshotAttributes) IncludeShared() terra.BoolValue {
	return terra.ReferenceBool(ds.ref.Append("include_shared"))
}

func (ds dataDbSnapshotAttributes) Iops() terra.NumberValue {
	return terra.ReferenceNumber(ds.ref.Append("iops"))
}

func (ds dataDbSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("kms_key_id"))
}

func (ds dataDbSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("license_model"))
}

func (ds dataDbSnapshotAttributes) MostRecent() terra.BoolValue {
	return terra.ReferenceBool(ds.ref.Append("most_recent"))
}

func (ds dataDbSnapshotAttributes) OptionGroupName() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("option_group_name"))
}

func (ds dataDbSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(ds.ref.Append("port"))
}

func (ds dataDbSnapshotAttributes) SnapshotCreateTime() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("snapshot_create_time"))
}

func (ds dataDbSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("snapshot_type"))
}

func (ds dataDbSnapshotAttributes) SourceDbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("source_db_snapshot_identifier"))
}

func (ds dataDbSnapshotAttributes) SourceRegion() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("source_region"))
}

func (ds dataDbSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("status"))
}

func (ds dataDbSnapshotAttributes) StorageType() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("storage_type"))
}

func (ds dataDbSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(ds.ref.Append("vpc_id"))
}
