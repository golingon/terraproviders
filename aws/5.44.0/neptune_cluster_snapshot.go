// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	neptuneclustersnapshot "github.com/golingon/terraproviders/aws/5.44.0/neptuneclustersnapshot"
	"io"
)

// NewNeptuneClusterSnapshot creates a new instance of [NeptuneClusterSnapshot].
func NewNeptuneClusterSnapshot(name string, args NeptuneClusterSnapshotArgs) *NeptuneClusterSnapshot {
	return &NeptuneClusterSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneClusterSnapshot)(nil)

// NeptuneClusterSnapshot represents the Terraform resource aws_neptune_cluster_snapshot.
type NeptuneClusterSnapshot struct {
	Name      string
	Args      NeptuneClusterSnapshotArgs
	state     *neptuneClusterSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) Type() string {
	return "aws_neptune_cluster_snapshot"
}

// LocalName returns the local name for [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) LocalName() string {
	return ncs.Name
}

// Configuration returns the configuration (args) for [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) Configuration() interface{} {
	return ncs.Args
}

// DependOn is used for other resources to depend on [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(ncs)
}

// Dependencies returns the list of resources [NeptuneClusterSnapshot] depends_on.
func (ncs *NeptuneClusterSnapshot) Dependencies() terra.Dependencies {
	return ncs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) LifecycleManagement() *terra.Lifecycle {
	return ncs.Lifecycle
}

// Attributes returns the attributes for [NeptuneClusterSnapshot].
func (ncs *NeptuneClusterSnapshot) Attributes() neptuneClusterSnapshotAttributes {
	return neptuneClusterSnapshotAttributes{ref: terra.ReferenceResource(ncs)}
}

// ImportState imports the given attribute values into [NeptuneClusterSnapshot]'s state.
func (ncs *NeptuneClusterSnapshot) ImportState(av io.Reader) error {
	ncs.state = &neptuneClusterSnapshotState{}
	if err := json.NewDecoder(av).Decode(ncs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncs.Type(), ncs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneClusterSnapshot] has state.
func (ncs *NeptuneClusterSnapshot) State() (*neptuneClusterSnapshotState, bool) {
	return ncs.state, ncs.state != nil
}

// StateMust returns the state for [NeptuneClusterSnapshot]. Panics if the state is nil.
func (ncs *NeptuneClusterSnapshot) StateMust() *neptuneClusterSnapshotState {
	if ncs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncs.Type(), ncs.LocalName()))
	}
	return ncs.state
}

// NeptuneClusterSnapshotArgs contains the configurations for aws_neptune_cluster_snapshot.
type NeptuneClusterSnapshotArgs struct {
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// DbClusterSnapshotIdentifier: string, required
	DbClusterSnapshotIdentifier terra.StringValue `hcl:"db_cluster_snapshot_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *neptuneclustersnapshot.Timeouts `hcl:"timeouts,block"`
}
type neptuneClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(ncs.ref.Append("allocated_storage"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ncs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("license_model"))
}

// Port returns a reference to field port of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ncs.ref.Append("port"))
}

// SnapshotType returns a reference to field snapshot_type of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ncs.ref.Append("storage_encrypted"))
}

// VpcId returns a reference to field vpc_id of aws_neptune_cluster_snapshot.
func (ncs neptuneClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("vpc_id"))
}

func (ncs neptuneClusterSnapshotAttributes) Timeouts() neptuneclustersnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[neptuneclustersnapshot.TimeoutsAttributes](ncs.ref.Append("timeouts"))
}

type neptuneClusterSnapshotState struct {
	AllocatedStorage            float64                               `json:"allocated_storage"`
	AvailabilityZones           []string                              `json:"availability_zones"`
	DbClusterIdentifier         string                                `json:"db_cluster_identifier"`
	DbClusterSnapshotArn        string                                `json:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier string                                `json:"db_cluster_snapshot_identifier"`
	Engine                      string                                `json:"engine"`
	EngineVersion               string                                `json:"engine_version"`
	Id                          string                                `json:"id"`
	KmsKeyId                    string                                `json:"kms_key_id"`
	LicenseModel                string                                `json:"license_model"`
	Port                        float64                               `json:"port"`
	SnapshotType                string                                `json:"snapshot_type"`
	SourceDbClusterSnapshotArn  string                                `json:"source_db_cluster_snapshot_arn"`
	Status                      string                                `json:"status"`
	StorageEncrypted            bool                                  `json:"storage_encrypted"`
	VpcId                       string                                `json:"vpc_id"`
	Timeouts                    *neptuneclustersnapshot.TimeoutsState `json:"timeouts"`
}
