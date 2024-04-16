// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_container_attached_cluster

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

// Resource represents the Terraform resource google_container_attached_cluster.
type Resource struct {
	Name      string
	Args      Args
	state     *googleContainerAttachedClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcac *Resource) Type() string {
	return "google_container_attached_cluster"
}

// LocalName returns the local name for [Resource].
func (gcac *Resource) LocalName() string {
	return gcac.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcac *Resource) Configuration() interface{} {
	return gcac.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcac *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcac)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcac *Resource) Dependencies() terra.Dependencies {
	return gcac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcac *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcac.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcac *Resource) Attributes() googleContainerAttachedClusterAttributes {
	return googleContainerAttachedClusterAttributes{ref: terra.ReferenceResource(gcac)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcac *Resource) ImportState(state io.Reader) error {
	gcac.state = &googleContainerAttachedClusterState{}
	if err := json.NewDecoder(state).Decode(gcac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcac.Type(), gcac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcac *Resource) State() (*googleContainerAttachedClusterState, bool) {
	return gcac.state, gcac.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcac *Resource) StateMust() *googleContainerAttachedClusterState {
	if gcac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcac.Type(), gcac.LocalName()))
	}
	return gcac.state
}

// Args contains the configurations for google_container_attached_cluster.
type Args struct {
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
	// Authorization: optional
	Authorization *Authorization `hcl:"authorization,block"`
	// BinaryAuthorization: optional
	BinaryAuthorization *BinaryAuthorization `hcl:"binary_authorization,block"`
	// Fleet: required
	Fleet *Fleet `hcl:"fleet,block" validate:"required"`
	// LoggingConfig: optional
	LoggingConfig *LoggingConfig `hcl:"logging_config,block"`
	// MonitoringConfig: optional
	MonitoringConfig *MonitoringConfig `hcl:"monitoring_config,block"`
	// OidcConfig: required
	OidcConfig *OidcConfig `hcl:"oidc_config,block" validate:"required"`
	// ProxyConfig: optional
	ProxyConfig *ProxyConfig `hcl:"proxy_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleContainerAttachedClusterAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcac.ref.Append("annotations"))
}

// ClusterRegion returns a reference to field cluster_region of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) ClusterRegion() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("cluster_region"))
}

// CreateTime returns a reference to field create_time of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("create_time"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("deletion_policy"))
}

// Description returns a reference to field description of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("description"))
}

// Distribution returns a reference to field distribution of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Distribution() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("distribution"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcac.ref.Append("effective_annotations"))
}

// Id returns a reference to field id of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("id"))
}

// KubernetesVersion returns a reference to field kubernetes_version of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) KubernetesVersion() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("kubernetes_version"))
}

// Location returns a reference to field location of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("location"))
}

// Name returns a reference to field name of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("platform_version"))
}

// Project returns a reference to field project of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(gcac.ref.Append("reconciling"))
}

// State returns a reference to field state of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("state"))
}

// Uid returns a reference to field uid of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_container_attached_cluster.
func (gcac googleContainerAttachedClusterAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gcac.ref.Append("update_time"))
}

func (gcac googleContainerAttachedClusterAttributes) Errors() terra.ListValue[ErrorsAttributes] {
	return terra.ReferenceAsList[ErrorsAttributes](gcac.ref.Append("errors"))
}

func (gcac googleContainerAttachedClusterAttributes) WorkloadIdentityConfig() terra.ListValue[WorkloadIdentityConfigAttributes] {
	return terra.ReferenceAsList[WorkloadIdentityConfigAttributes](gcac.ref.Append("workload_identity_config"))
}

func (gcac googleContainerAttachedClusterAttributes) Authorization() terra.ListValue[AuthorizationAttributes] {
	return terra.ReferenceAsList[AuthorizationAttributes](gcac.ref.Append("authorization"))
}

func (gcac googleContainerAttachedClusterAttributes) BinaryAuthorization() terra.ListValue[BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[BinaryAuthorizationAttributes](gcac.ref.Append("binary_authorization"))
}

func (gcac googleContainerAttachedClusterAttributes) Fleet() terra.ListValue[FleetAttributes] {
	return terra.ReferenceAsList[FleetAttributes](gcac.ref.Append("fleet"))
}

func (gcac googleContainerAttachedClusterAttributes) LoggingConfig() terra.ListValue[LoggingConfigAttributes] {
	return terra.ReferenceAsList[LoggingConfigAttributes](gcac.ref.Append("logging_config"))
}

func (gcac googleContainerAttachedClusterAttributes) MonitoringConfig() terra.ListValue[MonitoringConfigAttributes] {
	return terra.ReferenceAsList[MonitoringConfigAttributes](gcac.ref.Append("monitoring_config"))
}

func (gcac googleContainerAttachedClusterAttributes) OidcConfig() terra.ListValue[OidcConfigAttributes] {
	return terra.ReferenceAsList[OidcConfigAttributes](gcac.ref.Append("oidc_config"))
}

func (gcac googleContainerAttachedClusterAttributes) ProxyConfig() terra.ListValue[ProxyConfigAttributes] {
	return terra.ReferenceAsList[ProxyConfigAttributes](gcac.ref.Append("proxy_config"))
}

func (gcac googleContainerAttachedClusterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcac.ref.Append("timeouts"))
}

type googleContainerAttachedClusterState struct {
	Annotations            map[string]string             `json:"annotations"`
	ClusterRegion          string                        `json:"cluster_region"`
	CreateTime             string                        `json:"create_time"`
	DeletionPolicy         string                        `json:"deletion_policy"`
	Description            string                        `json:"description"`
	Distribution           string                        `json:"distribution"`
	EffectiveAnnotations   map[string]string             `json:"effective_annotations"`
	Id                     string                        `json:"id"`
	KubernetesVersion      string                        `json:"kubernetes_version"`
	Location               string                        `json:"location"`
	Name                   string                        `json:"name"`
	PlatformVersion        string                        `json:"platform_version"`
	Project                string                        `json:"project"`
	Reconciling            bool                          `json:"reconciling"`
	State                  string                        `json:"state"`
	Uid                    string                        `json:"uid"`
	UpdateTime             string                        `json:"update_time"`
	Errors                 []ErrorsState                 `json:"errors"`
	WorkloadIdentityConfig []WorkloadIdentityConfigState `json:"workload_identity_config"`
	Authorization          []AuthorizationState          `json:"authorization"`
	BinaryAuthorization    []BinaryAuthorizationState    `json:"binary_authorization"`
	Fleet                  []FleetState                  `json:"fleet"`
	LoggingConfig          []LoggingConfigState          `json:"logging_config"`
	MonitoringConfig       []MonitoringConfigState       `json:"monitoring_config"`
	OidcConfig             []OidcConfigState             `json:"oidc_config"`
	ProxyConfig            []ProxyConfigState            `json:"proxy_config"`
	Timeouts               *TimeoutsState                `json:"timeouts"`
}
