// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	containerattachedcluster "github.com/golingon/terraproviders/google/4.73.1/containerattachedcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAttachedCluster creates a new instance of [ContainerAttachedCluster].
func NewContainerAttachedCluster(name string, args ContainerAttachedClusterArgs) *ContainerAttachedCluster {
	return &ContainerAttachedCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAttachedCluster)(nil)

// ContainerAttachedCluster represents the Terraform resource google_container_attached_cluster.
type ContainerAttachedCluster struct {
	Name      string
	Args      ContainerAttachedClusterArgs
	state     *containerAttachedClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) Type() string {
	return "google_container_attached_cluster"
}

// LocalName returns the local name for [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) LocalName() string {
	return cac.Name
}

// Configuration returns the configuration (args) for [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) Configuration() interface{} {
	return cac.Args
}

// DependOn is used for other resources to depend on [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(cac)
}

// Dependencies returns the list of resources [ContainerAttachedCluster] depends_on.
func (cac *ContainerAttachedCluster) Dependencies() terra.Dependencies {
	return cac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) LifecycleManagement() *terra.Lifecycle {
	return cac.Lifecycle
}

// Attributes returns the attributes for [ContainerAttachedCluster].
func (cac *ContainerAttachedCluster) Attributes() containerAttachedClusterAttributes {
	return containerAttachedClusterAttributes{ref: terra.ReferenceResource(cac)}
}

// ImportState imports the given attribute values into [ContainerAttachedCluster]'s state.
func (cac *ContainerAttachedCluster) ImportState(av io.Reader) error {
	cac.state = &containerAttachedClusterState{}
	if err := json.NewDecoder(av).Decode(cac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cac.Type(), cac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAttachedCluster] has state.
func (cac *ContainerAttachedCluster) State() (*containerAttachedClusterState, bool) {
	return cac.state, cac.state != nil
}

// StateMust returns the state for [ContainerAttachedCluster]. Panics if the state is nil.
func (cac *ContainerAttachedCluster) StateMust() *containerAttachedClusterState {
	if cac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cac.Type(), cac.LocalName()))
	}
	return cac.state
}

// ContainerAttachedClusterArgs contains the configurations for google_container_attached_cluster.
type ContainerAttachedClusterArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Distribution: string, required
	Distribution terra.StringValue `hcl:"distribution,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformVersion: string, required
	PlatformVersion terra.StringValue `hcl:"platform_version,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Errors: min=0
	Errors []containerattachedcluster.Errors `hcl:"errors,block" validate:"min=0"`
	// WorkloadIdentityConfig: min=0
	WorkloadIdentityConfig []containerattachedcluster.WorkloadIdentityConfig `hcl:"workload_identity_config,block" validate:"min=0"`
	// Authorization: optional
	Authorization *containerattachedcluster.Authorization `hcl:"authorization,block"`
	// Fleet: required
	Fleet *containerattachedcluster.Fleet `hcl:"fleet,block" validate:"required"`
	// LoggingConfig: optional
	LoggingConfig *containerattachedcluster.LoggingConfig `hcl:"logging_config,block"`
	// MonitoringConfig: optional
	MonitoringConfig *containerattachedcluster.MonitoringConfig `hcl:"monitoring_config,block"`
	// OidcConfig: required
	OidcConfig *containerattachedcluster.OidcConfig `hcl:"oidc_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containerattachedcluster.Timeouts `hcl:"timeouts,block"`
}
type containerAttachedClusterAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cac.ref.Append("annotations"))
}

// ClusterRegion returns a reference to field cluster_region of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) ClusterRegion() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("cluster_region"))
}

// CreateTime returns a reference to field create_time of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("create_time"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("deletion_policy"))
}

// Description returns a reference to field description of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("description"))
}

// Distribution returns a reference to field distribution of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Distribution() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("distribution"))
}

// Id returns a reference to field id of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("id"))
}

// KubernetesVersion returns a reference to field kubernetes_version of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) KubernetesVersion() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("kubernetes_version"))
}

// Location returns a reference to field location of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("location"))
}

// Name returns a reference to field name of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("platform_version"))
}

// Project returns a reference to field project of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(cac.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("state"))
}

// Uid returns a reference to field uid of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_attached_cluster.
func (cac containerAttachedClusterAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("update_time"))
}

func (cac containerAttachedClusterAttributes) Errors() terra.ListValue[containerattachedcluster.ErrorsAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.ErrorsAttributes](cac.ref.Append("errors"))
}

func (cac containerAttachedClusterAttributes) WorkloadIdentityConfig() terra.ListValue[containerattachedcluster.WorkloadIdentityConfigAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.WorkloadIdentityConfigAttributes](cac.ref.Append("workload_identity_config"))
}

func (cac containerAttachedClusterAttributes) Authorization() terra.ListValue[containerattachedcluster.AuthorizationAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.AuthorizationAttributes](cac.ref.Append("authorization"))
}

func (cac containerAttachedClusterAttributes) Fleet() terra.ListValue[containerattachedcluster.FleetAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.FleetAttributes](cac.ref.Append("fleet"))
}

func (cac containerAttachedClusterAttributes) LoggingConfig() terra.ListValue[containerattachedcluster.LoggingConfigAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.LoggingConfigAttributes](cac.ref.Append("logging_config"))
}

func (cac containerAttachedClusterAttributes) MonitoringConfig() terra.ListValue[containerattachedcluster.MonitoringConfigAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.MonitoringConfigAttributes](cac.ref.Append("monitoring_config"))
}

func (cac containerAttachedClusterAttributes) OidcConfig() terra.ListValue[containerattachedcluster.OidcConfigAttributes] {
	return terra.ReferenceAsList[containerattachedcluster.OidcConfigAttributes](cac.ref.Append("oidc_config"))
}

func (cac containerAttachedClusterAttributes) Timeouts() containerattachedcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerattachedcluster.TimeoutsAttributes](cac.ref.Append("timeouts"))
}

type containerAttachedClusterState struct {
	Annotations            map[string]string                                      `json:"annotations"`
	ClusterRegion          string                                                 `json:"cluster_region"`
	CreateTime             string                                                 `json:"create_time"`
	DeletionPolicy         string                                                 `json:"deletion_policy"`
	Description            string                                                 `json:"description"`
	Distribution           string                                                 `json:"distribution"`
	Id                     string                                                 `json:"id"`
	KubernetesVersion      string                                                 `json:"kubernetes_version"`
	Location               string                                                 `json:"location"`
	Name                   string                                                 `json:"name"`
	PlatformVersion        string                                                 `json:"platform_version"`
	Project                string                                                 `json:"project"`
	Reconciling            bool                                                   `json:"reconciling"`
	State                  string                                                 `json:"state"`
	Uid                    string                                                 `json:"uid"`
	UpdateTime             string                                                 `json:"update_time"`
	Errors                 []containerattachedcluster.ErrorsState                 `json:"errors"`
	WorkloadIdentityConfig []containerattachedcluster.WorkloadIdentityConfigState `json:"workload_identity_config"`
	Authorization          []containerattachedcluster.AuthorizationState          `json:"authorization"`
	Fleet                  []containerattachedcluster.FleetState                  `json:"fleet"`
	LoggingConfig          []containerattachedcluster.LoggingConfigState          `json:"logging_config"`
	MonitoringConfig       []containerattachedcluster.MonitoringConfigState       `json:"monitoring_config"`
	OidcConfig             []containerattachedcluster.OidcConfigState             `json:"oidc_config"`
	Timeouts               *containerattachedcluster.TimeoutsState                `json:"timeouts"`
}
