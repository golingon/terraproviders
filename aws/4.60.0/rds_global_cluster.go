// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	rdsglobalcluster "github.com/golingon/terraproviders/aws/4.60.0/rdsglobalcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRdsGlobalCluster(name string, args RdsGlobalClusterArgs) *RdsGlobalCluster {
	return &RdsGlobalCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RdsGlobalCluster)(nil)

type RdsGlobalCluster struct {
	Name  string
	Args  RdsGlobalClusterArgs
	state *rdsGlobalClusterState
}

func (rgc *RdsGlobalCluster) Type() string {
	return "aws_rds_global_cluster"
}

func (rgc *RdsGlobalCluster) LocalName() string {
	return rgc.Name
}

func (rgc *RdsGlobalCluster) Configuration() interface{} {
	return rgc.Args
}

func (rgc *RdsGlobalCluster) Attributes() rdsGlobalClusterAttributes {
	return rdsGlobalClusterAttributes{ref: terra.ReferenceResource(rgc)}
}

func (rgc *RdsGlobalCluster) ImportState(av io.Reader) error {
	rgc.state = &rdsGlobalClusterState{}
	if err := json.NewDecoder(av).Decode(rgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgc.Type(), rgc.LocalName(), err)
	}
	return nil
}

func (rgc *RdsGlobalCluster) State() (*rdsGlobalClusterState, bool) {
	return rgc.state, rgc.state != nil
}

func (rgc *RdsGlobalCluster) StateMust() *rdsGlobalClusterState {
	if rgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgc.Type(), rgc.LocalName()))
	}
	return rgc.state
}

func (rgc *RdsGlobalCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(rgc)
}

type RdsGlobalClusterArgs struct {
	// DatabaseName: string, optional
	DatabaseName terra.StringValue `hcl:"database_name,attr"`
	// DeletionProtection: bool, optional
	DeletionProtection terra.BoolValue `hcl:"deletion_protection,attr"`
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// GlobalClusterIdentifier: string, required
	GlobalClusterIdentifier terra.StringValue `hcl:"global_cluster_identifier,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SourceDbClusterIdentifier: string, optional
	SourceDbClusterIdentifier terra.StringValue `hcl:"source_db_cluster_identifier,attr"`
	// StorageEncrypted: bool, optional
	StorageEncrypted terra.BoolValue `hcl:"storage_encrypted,attr"`
	// GlobalClusterMembers: min=0
	GlobalClusterMembers []rdsglobalcluster.GlobalClusterMembers `hcl:"global_cluster_members,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *rdsglobalcluster.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that RdsGlobalCluster depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type rdsGlobalClusterAttributes struct {
	ref terra.Reference
}

func (rgc rdsGlobalClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("arn"))
}

func (rgc rdsGlobalClusterAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("database_name"))
}

func (rgc rdsGlobalClusterAttributes) DeletionProtection() terra.BoolValue {
	return terra.ReferenceBool(rgc.ref.Append("deletion_protection"))
}

func (rgc rdsGlobalClusterAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("engine"))
}

func (rgc rdsGlobalClusterAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("engine_version"))
}

func (rgc rdsGlobalClusterAttributes) EngineVersionActual() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("engine_version_actual"))
}

func (rgc rdsGlobalClusterAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceBool(rgc.ref.Append("force_destroy"))
}

func (rgc rdsGlobalClusterAttributes) GlobalClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("global_cluster_identifier"))
}

func (rgc rdsGlobalClusterAttributes) GlobalClusterResourceId() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("global_cluster_resource_id"))
}

func (rgc rdsGlobalClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("id"))
}

func (rgc rdsGlobalClusterAttributes) SourceDbClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rgc.ref.Append("source_db_cluster_identifier"))
}

func (rgc rdsGlobalClusterAttributes) StorageEncrypted() terra.BoolValue {
	return terra.ReferenceBool(rgc.ref.Append("storage_encrypted"))
}

func (rgc rdsGlobalClusterAttributes) GlobalClusterMembers() terra.SetValue[rdsglobalcluster.GlobalClusterMembersAttributes] {
	return terra.ReferenceSet[rdsglobalcluster.GlobalClusterMembersAttributes](rgc.ref.Append("global_cluster_members"))
}

func (rgc rdsGlobalClusterAttributes) Timeouts() rdsglobalcluster.TimeoutsAttributes {
	return terra.ReferenceSingle[rdsglobalcluster.TimeoutsAttributes](rgc.ref.Append("timeouts"))
}

type rdsGlobalClusterState struct {
	Arn                       string                                       `json:"arn"`
	DatabaseName              string                                       `json:"database_name"`
	DeletionProtection        bool                                         `json:"deletion_protection"`
	Engine                    string                                       `json:"engine"`
	EngineVersion             string                                       `json:"engine_version"`
	EngineVersionActual       string                                       `json:"engine_version_actual"`
	ForceDestroy              bool                                         `json:"force_destroy"`
	GlobalClusterIdentifier   string                                       `json:"global_cluster_identifier"`
	GlobalClusterResourceId   string                                       `json:"global_cluster_resource_id"`
	Id                        string                                       `json:"id"`
	SourceDbClusterIdentifier string                                       `json:"source_db_cluster_identifier"`
	StorageEncrypted          bool                                         `json:"storage_encrypted"`
	GlobalClusterMembers      []rdsglobalcluster.GlobalClusterMembersState `json:"global_cluster_members"`
	Timeouts                  *rdsglobalcluster.TimeoutsState              `json:"timeouts"`
}
