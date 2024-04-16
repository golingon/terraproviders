// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_network_endpoint_group

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

// Resource represents the Terraform resource google_compute_region_network_endpoint_group.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionNetworkEndpointGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrneg *Resource) Type() string {
	return "google_compute_region_network_endpoint_group"
}

// LocalName returns the local name for [Resource].
func (gcrneg *Resource) LocalName() string {
	return gcrneg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrneg *Resource) Configuration() interface{} {
	return gcrneg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrneg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrneg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrneg *Resource) Dependencies() terra.Dependencies {
	return gcrneg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrneg *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrneg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrneg *Resource) Attributes() googleComputeRegionNetworkEndpointGroupAttributes {
	return googleComputeRegionNetworkEndpointGroupAttributes{ref: terra.ReferenceResource(gcrneg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrneg *Resource) ImportState(state io.Reader) error {
	gcrneg.state = &googleComputeRegionNetworkEndpointGroupState{}
	if err := json.NewDecoder(state).Decode(gcrneg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrneg.Type(), gcrneg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrneg *Resource) State() (*googleComputeRegionNetworkEndpointGroupState, bool) {
	return gcrneg.state, gcrneg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrneg *Resource) StateMust() *googleComputeRegionNetworkEndpointGroupState {
	if gcrneg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrneg.Type(), gcrneg.LocalName()))
	}
	return gcrneg.state
}

// Args contains the configurations for google_compute_region_network_endpoint_group.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkEndpointType: string, optional
	NetworkEndpointType terra.StringValue `hcl:"network_endpoint_type,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PscTargetService: string, optional
	PscTargetService terra.StringValue `hcl:"psc_target_service,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// AppEngine: optional
	AppEngine *AppEngine `hcl:"app_engine,block"`
	// CloudFunction: optional
	CloudFunction *CloudFunction `hcl:"cloud_function,block"`
	// CloudRun: optional
	CloudRun *CloudRun `hcl:"cloud_run,block"`
	// ServerlessDeployment: optional
	ServerlessDeployment *ServerlessDeployment `hcl:"serverless_deployment,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRegionNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("network"))
}

// NetworkEndpointType returns a reference to field network_endpoint_type of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("network_endpoint_type"))
}

// Project returns a reference to field project of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("project"))
}

// PscTargetService returns a reference to field psc_target_service of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) PscTargetService() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("psc_target_service"))
}

// Region returns a reference to field region of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("self_link"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_region_network_endpoint_group.
func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(gcrneg.ref.Append("subnetwork"))
}

func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) AppEngine() terra.ListValue[AppEngineAttributes] {
	return terra.ReferenceAsList[AppEngineAttributes](gcrneg.ref.Append("app_engine"))
}

func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) CloudFunction() terra.ListValue[CloudFunctionAttributes] {
	return terra.ReferenceAsList[CloudFunctionAttributes](gcrneg.ref.Append("cloud_function"))
}

func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) CloudRun() terra.ListValue[CloudRunAttributes] {
	return terra.ReferenceAsList[CloudRunAttributes](gcrneg.ref.Append("cloud_run"))
}

func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) ServerlessDeployment() terra.ListValue[ServerlessDeploymentAttributes] {
	return terra.ReferenceAsList[ServerlessDeploymentAttributes](gcrneg.ref.Append("serverless_deployment"))
}

func (gcrneg googleComputeRegionNetworkEndpointGroupAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrneg.ref.Append("timeouts"))
}

type googleComputeRegionNetworkEndpointGroupState struct {
	Description          string                      `json:"description"`
	Id                   string                      `json:"id"`
	Name                 string                      `json:"name"`
	Network              string                      `json:"network"`
	NetworkEndpointType  string                      `json:"network_endpoint_type"`
	Project              string                      `json:"project"`
	PscTargetService     string                      `json:"psc_target_service"`
	Region               string                      `json:"region"`
	SelfLink             string                      `json:"self_link"`
	Subnetwork           string                      `json:"subnetwork"`
	AppEngine            []AppEngineState            `json:"app_engine"`
	CloudFunction        []CloudFunctionState        `json:"cloud_function"`
	CloudRun             []CloudRunState             `json:"cloud_run"`
	ServerlessDeployment []ServerlessDeploymentState `json:"serverless_deployment"`
	Timeouts             *TimeoutsState              `json:"timeouts"`
}
