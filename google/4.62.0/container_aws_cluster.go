// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	containerawscluster "github.com/golingon/terraproviders/google/4.62.0/containerawscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAwsCluster creates a new instance of [ContainerAwsCluster].
func NewContainerAwsCluster(name string, args ContainerAwsClusterArgs) *ContainerAwsCluster {
	return &ContainerAwsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAwsCluster)(nil)

// ContainerAwsCluster represents the Terraform resource google_container_aws_cluster.
type ContainerAwsCluster struct {
	Name      string
	Args      ContainerAwsClusterArgs
	state     *containerAwsClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAwsCluster].
func (cac *ContainerAwsCluster) Type() string {
	return "google_container_aws_cluster"
}

// LocalName returns the local name for [ContainerAwsCluster].
func (cac *ContainerAwsCluster) LocalName() string {
	return cac.Name
}

// Configuration returns the configuration (args) for [ContainerAwsCluster].
func (cac *ContainerAwsCluster) Configuration() interface{} {
	return cac.Args
}

// DependOn is used for other resources to depend on [ContainerAwsCluster].
func (cac *ContainerAwsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(cac)
}

// Dependencies returns the list of resources [ContainerAwsCluster] depends_on.
func (cac *ContainerAwsCluster) Dependencies() terra.Dependencies {
	return cac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAwsCluster].
func (cac *ContainerAwsCluster) LifecycleManagement() *terra.Lifecycle {
	return cac.Lifecycle
}

// Attributes returns the attributes for [ContainerAwsCluster].
func (cac *ContainerAwsCluster) Attributes() containerAwsClusterAttributes {
	return containerAwsClusterAttributes{ref: terra.ReferenceResource(cac)}
}

// ImportState imports the given attribute values into [ContainerAwsCluster]'s state.
func (cac *ContainerAwsCluster) ImportState(av io.Reader) error {
	cac.state = &containerAwsClusterState{}
	if err := json.NewDecoder(av).Decode(cac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cac.Type(), cac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAwsCluster] has state.
func (cac *ContainerAwsCluster) State() (*containerAwsClusterState, bool) {
	return cac.state, cac.state != nil
}

// StateMust returns the state for [ContainerAwsCluster]. Panics if the state is nil.
func (cac *ContainerAwsCluster) StateMust() *containerAwsClusterState {
	if cac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cac.Type(), cac.LocalName()))
	}
	return cac.state
}

// ContainerAwsClusterArgs contains the configurations for google_container_aws_cluster.
type ContainerAwsClusterArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// AwsRegion: string, required
	AwsRegion terra.StringValue `hcl:"aws_region,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkloadIdentityConfig: min=0
	WorkloadIdentityConfig []containerawscluster.WorkloadIdentityConfig `hcl:"workload_identity_config,block" validate:"min=0"`
	// Authorization: required
	Authorization *containerawscluster.Authorization `hcl:"authorization,block" validate:"required"`
	// ControlPlane: required
	ControlPlane *containerawscluster.ControlPlane `hcl:"control_plane,block" validate:"required"`
	// Fleet: required
	Fleet *containerawscluster.Fleet `hcl:"fleet,block" validate:"required"`
	// Networking: required
	Networking *containerawscluster.Networking `hcl:"networking,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containerawscluster.Timeouts `hcl:"timeouts,block"`
}
type containerAwsClusterAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cac.ref.Append("annotations"))
}

// AwsRegion returns a reference to field aws_region of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) AwsRegion() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("aws_region"))
}

// CreateTime returns a reference to field create_time of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("create_time"))
}

// Description returns a reference to field description of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("description"))
}

// Endpoint returns a reference to field endpoint of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("endpoint"))
}

// Etag returns a reference to field etag of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("id"))
}

// Location returns a reference to field location of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("location"))
}

// Name returns a reference to field name of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("name"))
}

// Project returns a reference to field project of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(cac.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("state"))
}

// Uid returns a reference to field uid of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_aws_cluster.
func (cac containerAwsClusterAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("update_time"))
}

func (cac containerAwsClusterAttributes) WorkloadIdentityConfig() terra.ListValue[containerawscluster.WorkloadIdentityConfigAttributes] {
	return terra.ReferenceAsList[containerawscluster.WorkloadIdentityConfigAttributes](cac.ref.Append("workload_identity_config"))
}

func (cac containerAwsClusterAttributes) Authorization() terra.ListValue[containerawscluster.AuthorizationAttributes] {
	return terra.ReferenceAsList[containerawscluster.AuthorizationAttributes](cac.ref.Append("authorization"))
}

func (cac containerAwsClusterAttributes) ControlPlane() terra.ListValue[containerawscluster.ControlPlaneAttributes] {
	return terra.ReferenceAsList[containerawscluster.ControlPlaneAttributes](cac.ref.Append("control_plane"))
}

func (cac containerAwsClusterAttributes) Fleet() terra.ListValue[containerawscluster.FleetAttributes] {
	return terra.ReferenceAsList[containerawscluster.FleetAttributes](cac.ref.Append("fleet"))
}

func (cac containerAwsClusterAttributes) Networking() terra.ListValue[containerawscluster.NetworkingAttributes] {
	return terra.ReferenceAsList[containerawscluster.NetworkingAttributes](cac.ref.Append("networking"))
}

func (cac containerAwsClusterAttributes) Timeouts() containerawscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerawscluster.TimeoutsAttributes](cac.ref.Append("timeouts"))
}

type containerAwsClusterState struct {
	Annotations            map[string]string                                 `json:"annotations"`
	AwsRegion              string                                            `json:"aws_region"`
	CreateTime             string                                            `json:"create_time"`
	Description            string                                            `json:"description"`
	Endpoint               string                                            `json:"endpoint"`
	Etag                   string                                            `json:"etag"`
	Id                     string                                            `json:"id"`
	Location               string                                            `json:"location"`
	Name                   string                                            `json:"name"`
	Project                string                                            `json:"project"`
	Reconciling            bool                                              `json:"reconciling"`
	State                  string                                            `json:"state"`
	Uid                    string                                            `json:"uid"`
	UpdateTime             string                                            `json:"update_time"`
	WorkloadIdentityConfig []containerawscluster.WorkloadIdentityConfigState `json:"workload_identity_config"`
	Authorization          []containerawscluster.AuthorizationState          `json:"authorization"`
	ControlPlane           []containerawscluster.ControlPlaneState           `json:"control_plane"`
	Fleet                  []containerawscluster.FleetState                  `json:"fleet"`
	Networking             []containerawscluster.NetworkingState             `json:"networking"`
	Timeouts               *containerawscluster.TimeoutsState                `json:"timeouts"`
}
