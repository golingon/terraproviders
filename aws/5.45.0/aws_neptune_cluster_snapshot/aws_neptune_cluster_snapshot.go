// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_neptune_cluster_snapshot

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_neptune_cluster_snapshot.
type Resource struct {
	Name      string
	Args      Args
	state     *awsNeptuneClusterSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ancs *Resource) Type() string {
	return "aws_neptune_cluster_snapshot"
}

// LocalName returns the local name for [Resource].
func (ancs *Resource) LocalName() string {
	return ancs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ancs *Resource) Configuration() interface{} {
	return ancs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ancs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ancs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ancs *Resource) Dependencies() terra.Dependencies {
	return ancs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ancs *Resource) LifecycleManagement() *terra.Lifecycle {
	return ancs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ancs *Resource) Attributes() awsNeptuneClusterSnapshotAttributes {
	return awsNeptuneClusterSnapshotAttributes{ref: terra.ReferenceResource(ancs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ancs *Resource) ImportState(state io.Reader) error {
	ancs.state = &awsNeptuneClusterSnapshotState{}
	if err := json.NewDecoder(state).Decode(ancs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ancs.Type(), ancs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ancs *Resource) State() (*awsNeptuneClusterSnapshotState, bool) {
	return ancs.state, ancs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ancs *Resource) StateMust() *awsNeptuneClusterSnapshotState {
	if ancs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ancs.Type(), ancs.LocalName()))
	}
	return ancs.state
}

// Args contains the configurations for aws_neptune_cluster_snapshot.
type Args struct {
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// DbClusterSnapshotIdentifier: string, required
	DbClusterSnapshotIdentifier terra.StringValue `hcl:"db_cluster_snapshot_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsNeptuneClusterSnapshotAttributes struct {
	ref terra.Reference
}

// AllocatedStorage returns a reference to field allocated_storage of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) AllocatedStorage() terra.NumberValue {
	return terra.ReferenceAsNumber(ancs.ref.Append("allocated_storage"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ancs.ref.Append("availability_zones"))
}

// DbClusterIdentifier returns a reference to field db_cluster_identifier of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("db_cluster_identifier"))
}

// DbClusterSnapshotArn returns a reference to field db_cluster_snapshot_arn of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) DbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("db_cluster_snapshot_arn"))
}

// DbClusterSnapshotIdentifier returns a reference to field db_cluster_snapshot_identifier of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) DbClusterSnapshotIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("db_cluster_snapshot_identifier"))
}

// Engine returns a reference to field engine of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("kms_key_id"))
}

// LicenseModel returns a reference to field license_model of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) LicenseModel() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("license_model"))
}

// Port returns a reference to field port of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ancs.ref.Append("port"))
}

// SnapshotType returns a reference to field snapshot_type of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) SnapshotType() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("snapshot_type"))
}

// SourceDbClusterSnapshotArn returns a reference to field source_db_cluster_snapshot_arn of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) SourceDbClusterSnapshotArn() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("source_db_cluster_snapshot_arn"))
}

// Status returns a reference to field status of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ancs.ref.Append("storage_encrypted"))
}

// VpcId returns a reference to field vpc_id of aws_neptune_cluster_snapshot.
func (ancs awsNeptuneClusterSnapshotAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ancs.ref.Append("vpc_id"))
}

func (ancs awsNeptuneClusterSnapshotAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ancs.ref.Append("timeouts"))
}

type awsNeptuneClusterSnapshotState struct {
	AllocatedStorage            float64        `json:"allocated_storage"`
	AvailabilityZones           []string       `json:"availability_zones"`
	DbClusterIdentifier         string         `json:"db_cluster_identifier"`
	DbClusterSnapshotArn        string         `json:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier string         `json:"db_cluster_snapshot_identifier"`
	Engine                      string         `json:"engine"`
	EngineVersion               string         `json:"engine_version"`
	Id                          string         `json:"id"`
	KmsKeyId                    string         `json:"kms_key_id"`
	LicenseModel                string         `json:"license_model"`
	Port                        float64        `json:"port"`
	SnapshotType                string         `json:"snapshot_type"`
	SourceDbClusterSnapshotArn  string         `json:"source_db_cluster_snapshot_arn"`
	Status                      string         `json:"status"`
	StorageEncrypted            bool           `json:"storage_encrypted"`
	VpcId                       string         `json:"vpc_id"`
	Timeouts                    *TimeoutsState `json:"timeouts"`
}
