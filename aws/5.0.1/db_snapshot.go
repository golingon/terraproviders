// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dbsnapshot "github.com/golingon/terraproviders/aws/5.0.1/dbsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbSnapshot creates a new instance of [DbSnapshot].
func NewDbSnapshot(name string, args DbSnapshotArgs) *DbSnapshot {
	return &DbSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbSnapshot)(nil)

// DbSnapshot represents the Terraform resource aws_db_snapshot.
type DbSnapshot struct {
	Name      string
	Args      DbSnapshotArgs
	state     *dbSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbSnapshot].
func (ds *DbSnapshot) Type() string {
	return "aws_db_snapshot"
}

// LocalName returns the local name for [DbSnapshot].
func (ds *DbSnapshot) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DbSnapshot].
func (ds *DbSnapshot) Configuration() interface{} {
	return ds.Args
}

// DependOn is used for other resources to depend on [DbSnapshot].
func (ds *DbSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(ds)
}

// Dependencies returns the list of resources [DbSnapshot] depends_on.
func (ds *DbSnapshot) Dependencies() terra.Dependencies {
	return ds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbSnapshot].
func (ds *DbSnapshot) LifecycleManagement() *terra.Lifecycle {
	return ds.Lifecycle
}

// Attributes returns the attributes for [DbSnapshot].
func (ds *DbSnapshot) Attributes() dbSnapshotAttributes {
	return dbSnapshotAttributes{ref: terra.ReferenceResource(ds)}
}

// ImportState imports the given attribute values into [DbSnapshot]'s state.
func (ds *DbSnapshot) ImportState(av io.Reader) error {
	ds.state = &dbSnapshotState{}
	if err := json.NewDecoder(av).Decode(ds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ds.Type(), ds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbSnapshot] has state.
func (ds *DbSnapshot) State() (*dbSnapshotState, bool) {
	return ds.state, ds.state != nil
}

// StateMust returns the state for [DbSnapshot]. Panics if the state is nil.
func (ds *DbSnapshot) StateMust() *dbSnapshotState {
	if ds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ds.Type(), ds.LocalName()))
	}
	return ds.state
}

// DbSnapshotArgs contains the configurations for aws_db_snapshot.
type DbSnapshotArgs struct {
	// DbInstanceIdentifier: string, required
	DbInstanceIdentifier terra.StringValue `hcl:"db_instance_identifier,attr" validate:"required"`
	// DbSnapshotIdentifier: string, required
	DbSnapshotIdentifier terra.StringValue `hcl:"db_snapshot_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SharedAccounts: set of string, optional
	SharedAccounts terra.SetValue[terra.StringValue] `hcl:"shared_accounts,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *dbsnapshot.Timeouts `hcl:"timeouts,block"`
}
type dbSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_db_snapshot.
func (ds dbSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("allocated_storage"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_db_snapshot.
func (ds dbSnapshotAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("availability_zone"))
}

// DbInstanceIdentifier returns a reference to field db_instance_identifier of aws_db_snapshot.
func (ds dbSnapshotAttributes) DbInstanceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_instance_identifier"))
}

// DbSnapshotArn returns a reference to field db_snapshot_arn of aws_db_snapshot.
func (ds dbSnapshotAttributes) DbSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_snapshot_arn"))
}

// DbSnapshotIdentifier returns a reference to field db_snapshot_identifier of aws_db_snapshot.
func (ds dbSnapshotAttributes) DbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("db_snapshot_identifier"))
}

// Encrypted returns a reference to field encrypted of aws_db_snapshot.
func (ds dbSnapshotAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("encrypted"))
}

// Engine returns a reference to field engine of aws_db_snapshot.
func (ds dbSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_db_snapshot.
func (ds dbSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_db_snapshot.
func (ds dbSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// Iops returns a reference to field iops of aws_db_snapshot.
func (ds dbSnapshotAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("iops"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_snapshot.
func (ds dbSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_db_snapshot.
func (ds dbSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("license_model"))
}

// OptionGroupName returns a reference to field option_group_name of aws_db_snapshot.
func (ds dbSnapshotAttributes) OptionGroupName() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("option_group_name"))
}

// Port returns a reference to field port of aws_db_snapshot.
func (ds dbSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("port"))
}

// SharedAccounts returns a reference to field shared_accounts of aws_db_snapshot.
func (ds dbSnapshotAttributes) SharedAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ds.ref.Append("shared_accounts"))
}

// SnapshotType returns a reference to field snapshot_type of aws_db_snapshot.
func (ds dbSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("snapshot_type"))
}

// SourceDbSnapshotIdentifier returns a reference to field source_db_snapshot_identifier of aws_db_snapshot.
func (ds dbSnapshotAttributes) SourceDbSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("source_db_snapshot_identifier"))
}

// SourceRegion returns a reference to field source_region of aws_db_snapshot.
func (ds dbSnapshotAttributes) SourceRegion() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("source_region"))
}

// Status returns a reference to field status of aws_db_snapshot.
func (ds dbSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("status"))
}

// StorageType returns a reference to field storage_type of aws_db_snapshot.
func (ds dbSnapshotAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of aws_db_snapshot.
func (ds dbSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_snapshot.
func (ds dbSnapshotAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_db_snapshot.
func (ds dbSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("vpc_id"))
}

func (ds dbSnapshotAttributes) Timeouts() dbsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dbsnapshot.TimeoutsAttributes](ds.ref.Append("timeouts"))
}

type dbSnapshotState struct {
	AllocatedStorage           float64                   `json:"allocated_storage"`
	AvailabilityZone           string                    `json:"availability_zone"`
	DbInstanceIdentifier       string                    `json:"db_instance_identifier"`
	DbSnapshotArn              string                    `json:"db_snapshot_arn"`
	DbSnapshotIdentifier       string                    `json:"db_snapshot_identifier"`
	Encrypted                  bool                      `json:"encrypted"`
	Engine                     string                    `json:"engine"`
	EngineVersion              string                    `json:"engine_version"`
	Id                         string                    `json:"id"`
	Iops                       float64                   `json:"iops"`
	KmsKeyId                   string                    `json:"kms_key_id"`
	LicenseModel               string                    `json:"license_model"`
	OptionGroupName            string                    `json:"option_group_name"`
	Port                       float64                   `json:"port"`
	SharedAccounts             []string                  `json:"shared_accounts"`
	SnapshotType               string                    `json:"snapshot_type"`
	SourceDbSnapshotIdentifier string                    `json:"source_db_snapshot_identifier"`
	SourceRegion               string                    `json:"source_region"`
	Status                     string                    `json:"status"`
	StorageType                string                    `json:"storage_type"`
	Tags                       map[string]string         `json:"tags"`
	TagsAll                    map[string]string         `json:"tags_all"`
	VpcId                      string                    `json:"vpc_id"`
	Timeouts                   *dbsnapshot.TimeoutsState `json:"timeouts"`
}
