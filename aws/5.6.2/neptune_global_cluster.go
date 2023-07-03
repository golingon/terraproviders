// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptuneglobalcluster "github.com/golingon/terraproviders/aws/5.6.2/neptuneglobalcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNeptuneGlobalCluster creates a new instance of [NeptuneGlobalCluster].
func NewNeptuneGlobalCluster(name string, args NeptuneGlobalClusterArgs) *NeptuneGlobalCluster {
	return &NeptuneGlobalCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneGlobalCluster)(nil)

// NeptuneGlobalCluster represents the Terraform resource aws_neptune_global_cluster.
type NeptuneGlobalCluster struct {
	Name      string
	Args      NeptuneGlobalClusterArgs
	state     *neptuneGlobalClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) Type() string {
	return "aws_neptune_global_cluster"
}

// LocalName returns the local name for [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) LocalName() string {
	return ngc.Name
}

// Configuration returns the configuration (args) for [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) Configuration() interface{} {
	return ngc.Args
}

// DependOn is used for other resources to depend on [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ngc)
}

// Dependencies returns the list of resources [NeptuneGlobalCluster] depends_on.
func (ngc *NeptuneGlobalCluster) Dependencies() terra.Dependencies {
	return ngc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) LifecycleManagement() *terra.Lifecycle {
	return ngc.Lifecycle
}

// Attributes returns the attributes for [NeptuneGlobalCluster].
func (ngc *NeptuneGlobalCluster) Attributes() neptuneGlobalClusterAttributes {
	return neptuneGlobalClusterAttributes{ref: terra.ReferenceResource(ngc)}
}

// ImportState imports the given attribute values into [NeptuneGlobalCluster]'s state.
func (ngc *NeptuneGlobalCluster) ImportState(av io.Reader) error {
	ngc.state = &neptuneGlobalClusterState{}
	if err := json.NewDecoder(av).Decode(ngc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ngc.Type(), ngc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneGlobalCluster] has state.
func (ngc *NeptuneGlobalCluster) State() (*neptuneGlobalClusterState, bool) {
	return ngc.state, ngc.state != nil
}

// StateMust returns the state for [NeptuneGlobalCluster]. Panics if the state is nil.
func (ngc *NeptuneGlobalCluster) StateMust() *neptuneGlobalClusterState {
	if ngc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ngc.Type(), ngc.LocalName()))
	}
	return ngc.state
}

// NeptuneGlobalClusterArgs contains the configurations for aws_neptune_global_cluster.
type NeptuneGlobalClusterArgs struct {
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// GlobalClusterIdentifier: string, required
	GlobalClusterIdentifier terra.StringValue `hcl:"global_cluster_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SourceDbClusterIdentifier: string, optional
	SourceDbClusterIdentifier terra.StringValue `hcl:"source_db_cluster_identifier,attr"`
	// StorageEncrypted: bool, optional
	StorageEncrypted terra.BoolValue `hcl:"storage_encrypted,attr"`
	// GlobalClusterMembers: min=0
	GlobalClusterMembers []neptuneglobalcluster.GlobalClusterMembers `hcl:"global_cluster_members,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *neptuneglobalcluster.Timeouts `hcl:"timeouts,block"`
}
type neptuneGlobalClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("arn"))
}

// DeletionProtection returns a reference to field deletion_protection of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(ngc.ref.Append("deletion_protection"))
}

// Engine returns a reference to field engine of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("engine"))
}

// EngineVersion returns a reference to field engine_version of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("engine_version"))
}

// GlobalClusterIdentifier returns a reference to field global_cluster_identifier of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) GlobalClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("global_cluster_identifier"))
}

// GlobalClusterResourceId returns a reference to field global_cluster_resource_id of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) GlobalClusterResourceId() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("global_cluster_resource_id"))
}

// Id returns a reference to field id of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("id"))
}

// SourceDbClusterIdentifier returns a reference to field source_db_cluster_identifier of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) SourceDbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("source_db_cluster_identifier"))
}

// Status returns a reference to field status of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ngc.ref.Append("status"))
}

// StorageEncrypted returns a reference to field storage_encrypted of aws_neptune_global_cluster.
func (ngc neptuneGlobalClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ngc.ref.Append("storage_encrypted"))
}

func (ngc neptuneGlobalClusterAttributes) GlobalClusterMembers() terra.SetValue[neptuneglobalcluster.GlobalClusterMembersAttributes] {
	return terra.ReferenceAsSet[neptuneglobalcluster.GlobalClusterMembersAttributes](ngc.ref.Append("global_cluster_members"))
}

func (ngc neptuneGlobalClusterAttributes) Timeouts() neptuneglobalcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[neptuneglobalcluster.TimeoutsAttributes](ngc.ref.Append("timeouts"))
}

type neptuneGlobalClusterState struct {
	Arn                       string                                           `json:"arn"`
	DeletionProtection        bool                                             `json:"deletion_protection"`
	Engine                    string                                           `json:"engine"`
	EngineVersion             string                                           `json:"engine_version"`
	GlobalClusterIdentifier   string                                           `json:"global_cluster_identifier"`
	GlobalClusterResourceId   string                                           `json:"global_cluster_resource_id"`
	Id                        string                                           `json:"id"`
	SourceDbClusterIdentifier string                                           `json:"source_db_cluster_identifier"`
	Status                    string                                           `json:"status"`
	StorageEncrypted          bool                                             `json:"storage_encrypted"`
	GlobalClusterMembers      []neptuneglobalcluster.GlobalClusterMembersState `json:"global_cluster_members"`
	Timeouts                  *neptuneglobalcluster.TimeoutsState              `json:"timeouts"`
}
