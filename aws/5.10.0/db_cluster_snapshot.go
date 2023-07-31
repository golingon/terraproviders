// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dbclustersnapshot "github.com/golingon/terraproviders/aws/5.10.0/dbclustersnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDbClusterSnapshot creates a new instance of [DbClusterSnapshot].
func NewDbClusterSnapshot(name string, args DbClusterSnapshotArgs) *DbClusterSnapshot {
	return &DbClusterSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DbClusterSnapshot)(nil)

// DbClusterSnapshot represents the Terraform resource aws_db_cluster_snapshot.
type DbClusterSnapshot struct {
	Name      string
	Args      DbClusterSnapshotArgs
	state     *dbClusterSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) Type() string {
	return "aws_db_cluster_snapshot"
}

// LocalName returns the local name for [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) LocalName() string {
	return dcs.Name
}

// Configuration returns the configuration (args) for [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) Configuration() interface{} {
	return dcs.Args
}

// DependOn is used for other resources to depend on [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(dcs)
}

// Dependencies returns the list of resources [DbClusterSnapshot] depends_on.
func (dcs *DbClusterSnapshot) Dependencies() terra.Dependencies {
	return dcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) LifecycleManagement() *terra.Lifecycle {
	return dcs.Lifecycle
}

// Attributes returns the attributes for [DbClusterSnapshot].
func (dcs *DbClusterSnapshot) Attributes() dbClusterSnapshotAttributes {
	return dbClusterSnapshotAttributes{ref: terra.ReferenceResource(dcs)}
}

// ImportState imports the given attribute values into [DbClusterSnapshot]'s state.
func (dcs *DbClusterSnapshot) ImportState(av io.Reader) error {
	dcs.state = &dbClusterSnapshotState{}
	if err := json.NewDecoder(av).Decode(dcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcs.Type(), dcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DbClusterSnapshot] has state.
func (dcs *DbClusterSnapshot) State() (*dbClusterSnapshotState, bool) {
	return dcs.state, dcs.state != nil
}

// StateMust returns the state for [DbClusterSnapshot]. Panics if the state is nil.
func (dcs *DbClusterSnapshot) StateMust() *dbClusterSnapshotState {
	if dcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcs.Type(), dcs.LocalName()))
	}
	return dcs.state
}

// DbClusterSnapshotArgs contains the configurations for aws_db_cluster_snapshot.
type DbClusterSnapshotArgs struct {
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// DbClusterSnapshotIdentifier: string, required
	DbClusterSnapshotIdentifier terra.StringValue `hcl:"db_cluster_snapshot_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *dbclustersnapshot.Timeouts `hcl:"timeouts,block"`
}
type dbClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(dcs.ref.Append("allocated_storage"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("license_model"))
}

// Port returns a reference to field port of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dcs.ref.Append("port"))
}

// SnapshotType returns a reference to field snapshot_type of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("storage_encrypted"))
}

// Tags returns a reference to field tags of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcs.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_db_cluster_snapshot.
func (dcs dbClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("vpc_id"))
}

func (dcs dbClusterSnapshotAttributes) Timeouts() dbclustersnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dbclustersnapshot.TimeoutsAttributes](dcs.ref.Append("timeouts"))
}

type dbClusterSnapshotState struct {
	AllocatedStorage            float64                          `json:"allocated_storage"`
	AvailabilityZones           []string                         `json:"availability_zones"`
	DbClusterIdentifier         string                           `json:"db_cluster_identifier"`
	DbClusterSnapshotArn        string                           `json:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier string                           `json:"db_cluster_snapshot_identifier"`
	Engine                      string                           `json:"engine"`
	EngineVersion               string                           `json:"engine_version"`
	Id                          string                           `json:"id"`
	KmsKeyId                    string                           `json:"kms_key_id"`
	LicenseModel                string                           `json:"license_model"`
	Port                        float64                          `json:"port"`
	SnapshotType                string                           `json:"snapshot_type"`
	SourceDbClusterSnapshotArn  string                           `json:"source_db_cluster_snapshot_arn"`
	Status                      string                           `json:"status"`
	StorageEncrypted            bool                             `json:"storage_encrypted"`
	Tags                        map[string]string                `json:"tags"`
	TagsAll                     map[string]string                `json:"tags_all"`
	VpcId                       string                           `json:"vpc_id"`
	Timeouts                    *dbclustersnapshot.TimeoutsState `json:"timeouts"`
}
