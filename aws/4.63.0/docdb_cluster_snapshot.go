// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	docdbclustersnapshot "github.com/golingon/terraproviders/aws/4.63.0/docdbclustersnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocdbClusterSnapshot creates a new instance of [DocdbClusterSnapshot].
func NewDocdbClusterSnapshot(name string, args DocdbClusterSnapshotArgs) *DocdbClusterSnapshot {
	return &DocdbClusterSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocdbClusterSnapshot)(nil)

// DocdbClusterSnapshot represents the Terraform resource aws_docdb_cluster_snapshot.
type DocdbClusterSnapshot struct {
	Name      string
	Args      DocdbClusterSnapshotArgs
	state     *docdbClusterSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) Type() string {
	return "aws_docdb_cluster_snapshot"
}

// LocalName returns the local name for [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) LocalName() string {
	return dcs.Name
}

// Configuration returns the configuration (args) for [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) Configuration() interface{} {
	return dcs.Args
}

// DependOn is used for other resources to depend on [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(dcs)
}

// Dependencies returns the list of resources [DocdbClusterSnapshot] depends_on.
func (dcs *DocdbClusterSnapshot) Dependencies() terra.Dependencies {
	return dcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) LifecycleManagement() *terra.Lifecycle {
	return dcs.Lifecycle
}

// Attributes returns the attributes for [DocdbClusterSnapshot].
func (dcs *DocdbClusterSnapshot) Attributes() docdbClusterSnapshotAttributes {
	return docdbClusterSnapshotAttributes{ref: terra.ReferenceResource(dcs)}
}

// ImportState imports the given attribute values into [DocdbClusterSnapshot]'s state.
func (dcs *DocdbClusterSnapshot) ImportState(av io.Reader) error {
	dcs.state = &docdbClusterSnapshotState{}
	if err := json.NewDecoder(av).Decode(dcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcs.Type(), dcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocdbClusterSnapshot] has state.
func (dcs *DocdbClusterSnapshot) State() (*docdbClusterSnapshotState, bool) {
	return dcs.state, dcs.state != nil
}

// StateMust returns the state for [DocdbClusterSnapshot]. Panics if the state is nil.
func (dcs *DocdbClusterSnapshot) StateMust() *docdbClusterSnapshotState {
	if dcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcs.Type(), dcs.LocalName()))
	}
	return dcs.state
}

// DocdbClusterSnapshotArgs contains the configurations for aws_docdb_cluster_snapshot.
type DocdbClusterSnapshotArgs struct {
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// DbClusterSnapshotIdentifier: string, required
	DbClusterSnapshotIdentifier terra.StringValue `hcl:"db_cluster_snapshot_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *docdbclustersnapshot.Timeouts `hcl:"timeouts,block"`
}
type docdbClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AvailabilityZones returns a reference to field availability_zones of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("kms_key_id"))
}

// Port returns a reference to field port of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dcs.ref.Append("port"))
}

// SnapshotType returns a reference to field snapshot_type of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(dcs.ref.Append("storage_encrypted"))
}

// VpcId returns a reference to field vpc_id of aws_docdb_cluster_snapshot.
func (dcs docdbClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(dcs.ref.Append("vpc_id"))
}

func (dcs docdbClusterSnapshotAttributes) Timeouts() docdbclustersnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[docdbclustersnapshot.TimeoutsAttributes](dcs.ref.Append("timeouts"))
}

type docdbClusterSnapshotState struct {
	AvailabilityZones           []string                            `json:"availability_zones"`
	DbClusterIdentifier         string                              `json:"db_cluster_identifier"`
	DbClusterSnapshotArn        string                              `json:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier string                              `json:"db_cluster_snapshot_identifier"`
	Engine                      string                              `json:"engine"`
	EngineVersion               string                              `json:"engine_version"`
	Id                          string                              `json:"id"`
	KmsKeyId                    string                              `json:"kms_key_id"`
	Port                        float64                             `json:"port"`
	SnapshotType                string                              `json:"snapshot_type"`
	SourceDbClusterSnapshotArn  string                              `json:"source_db_cluster_snapshot_arn"`
	Status                      string                              `json:"status"`
	StorageEncrypted            bool                                `json:"storage_encrypted"`
	VpcId                       string                              `json:"vpc_id"`
	Timeouts                    *docdbclustersnapshot.TimeoutsState `json:"timeouts"`
}
