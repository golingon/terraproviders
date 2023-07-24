// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstationcluster "github.com/golingon/terraproviders/googlebeta/4.74.0/workstationsworkstationcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstationCluster creates a new instance of [WorkstationsWorkstationCluster].
func NewWorkstationsWorkstationCluster(name string, args WorkstationsWorkstationClusterArgs) *WorkstationsWorkstationCluster {
	return &WorkstationsWorkstationCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstationCluster)(nil)

// WorkstationsWorkstationCluster represents the Terraform resource google_workstations_workstation_cluster.
type WorkstationsWorkstationCluster struct {
	Name      string
	Args      WorkstationsWorkstationClusterArgs
	state     *workstationsWorkstationClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) Type() string {
	return "google_workstations_workstation_cluster"
}

// LocalName returns the local name for [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) LocalName() string {
	return wwc.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) Configuration() interface{} {
	return wwc.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(wwc)
}

// Dependencies returns the list of resources [WorkstationsWorkstationCluster] depends_on.
func (wwc *WorkstationsWorkstationCluster) Dependencies() terra.Dependencies {
	return wwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) LifecycleManagement() *terra.Lifecycle {
	return wwc.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstationCluster].
func (wwc *WorkstationsWorkstationCluster) Attributes() workstationsWorkstationClusterAttributes {
	return workstationsWorkstationClusterAttributes{ref: terra.ReferenceResource(wwc)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstationCluster]'s state.
func (wwc *WorkstationsWorkstationCluster) ImportState(av io.Reader) error {
	wwc.state = &workstationsWorkstationClusterState{}
	if err := json.NewDecoder(av).Decode(wwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwc.Type(), wwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstationCluster] has state.
func (wwc *WorkstationsWorkstationCluster) State() (*workstationsWorkstationClusterState, bool) {
	return wwc.state, wwc.state != nil
}

// StateMust returns the state for [WorkstationsWorkstationCluster]. Panics if the state is nil.
func (wwc *WorkstationsWorkstationCluster) StateMust() *workstationsWorkstationClusterState {
	if wwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwc.Type(), wwc.LocalName()))
	}
	return wwc.state
}

// WorkstationsWorkstationClusterArgs contains the configurations for google_workstations_workstation_cluster.
type WorkstationsWorkstationClusterArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Subnetwork: string, required
	Subnetwork terra.StringValue `hcl:"subnetwork,attr" validate:"required"`
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// Conditions: min=0
	Conditions []workstationsworkstationcluster.Conditions `hcl:"conditions,block" validate:"min=0"`
	// PrivateClusterConfig: optional
	PrivateClusterConfig *workstationsworkstationcluster.PrivateClusterConfig `hcl:"private_cluster_config,block"`
	// Timeouts: optional
	Timeouts *workstationsworkstationcluster.Timeouts `hcl:"timeouts,block"`
}
type workstationsWorkstationClusterAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwc.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("create_time"))
}

// Degraded returns a reference to field degraded of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Degraded() terra.BoolValue {
	return terra.ReferenceAsBool(wwc.ref.Append("degraded"))
}

// DisplayName returns a reference to field display_name of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wwc.ref.Append("labels"))
}

// Location returns a reference to field location of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("location"))
}

// Name returns a reference to field name of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("name"))
}

// Network returns a reference to field network of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("network"))
}

// Project returns a reference to field project of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("project"))
}

// Subnetwork returns a reference to field subnetwork of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("subnetwork"))
}

// Uid returns a reference to field uid of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("uid"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_cluster.
func (wwc workstationsWorkstationClusterAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwc.ref.Append("workstation_cluster_id"))
}

func (wwc workstationsWorkstationClusterAttributes) Conditions() terra.ListValue[workstationsworkstationcluster.ConditionsAttributes] {
	return terra.ReferenceAsList[workstationsworkstationcluster.ConditionsAttributes](wwc.ref.Append("conditions"))
}

func (wwc workstationsWorkstationClusterAttributes) PrivateClusterConfig() terra.ListValue[workstationsworkstationcluster.PrivateClusterConfigAttributes] {
	return terra.ReferenceAsList[workstationsworkstationcluster.PrivateClusterConfigAttributes](wwc.ref.Append("private_cluster_config"))
}

func (wwc workstationsWorkstationClusterAttributes) Timeouts() workstationsworkstationcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workstationsworkstationcluster.TimeoutsAttributes](wwc.ref.Append("timeouts"))
}

type workstationsWorkstationClusterState struct {
	Annotations          map[string]string                                          `json:"annotations"`
	CreateTime           string                                                     `json:"create_time"`
	Degraded             bool                                                       `json:"degraded"`
	DisplayName          string                                                     `json:"display_name"`
	Etag                 string                                                     `json:"etag"`
	Id                   string                                                     `json:"id"`
	Labels               map[string]string                                          `json:"labels"`
	Location             string                                                     `json:"location"`
	Name                 string                                                     `json:"name"`
	Network              string                                                     `json:"network"`
	Project              string                                                     `json:"project"`
	Subnetwork           string                                                     `json:"subnetwork"`
	Uid                  string                                                     `json:"uid"`
	WorkstationClusterId string                                                     `json:"workstation_cluster_id"`
	Conditions           []workstationsworkstationcluster.ConditionsState           `json:"conditions"`
	PrivateClusterConfig []workstationsworkstationcluster.PrivateClusterConfigState `json:"private_cluster_config"`
	Timeouts             *workstationsworkstationcluster.TimeoutsState              `json:"timeouts"`
}
