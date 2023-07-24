// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionnetworkendpointgroup "github.com/golingon/terraproviders/googlebeta/4.74.0/computeregionnetworkendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionNetworkEndpointGroup creates a new instance of [ComputeRegionNetworkEndpointGroup].
func NewComputeRegionNetworkEndpointGroup(name string, args ComputeRegionNetworkEndpointGroupArgs) *ComputeRegionNetworkEndpointGroup {
	return &ComputeRegionNetworkEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkEndpointGroup)(nil)

// ComputeRegionNetworkEndpointGroup represents the Terraform resource google_compute_region_network_endpoint_group.
type ComputeRegionNetworkEndpointGroup struct {
	Name      string
	Args      ComputeRegionNetworkEndpointGroupArgs
	state     *computeRegionNetworkEndpointGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) Type() string {
	return "google_compute_region_network_endpoint_group"
}

// LocalName returns the local name for [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) LocalName() string {
	return crneg.Name
}

// Configuration returns the configuration (args) for [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) Configuration() interface{} {
	return crneg.Args
}

// DependOn is used for other resources to depend on [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(crneg)
}

// Dependencies returns the list of resources [ComputeRegionNetworkEndpointGroup] depends_on.
func (crneg *ComputeRegionNetworkEndpointGroup) Dependencies() terra.Dependencies {
	return crneg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) LifecycleManagement() *terra.Lifecycle {
	return crneg.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionNetworkEndpointGroup].
func (crneg *ComputeRegionNetworkEndpointGroup) Attributes() computeRegionNetworkEndpointGroupAttributes {
	return computeRegionNetworkEndpointGroupAttributes{ref: terra.ReferenceResource(crneg)}
}

// ImportState imports the given attribute values into [ComputeRegionNetworkEndpointGroup]'s state.
func (crneg *ComputeRegionNetworkEndpointGroup) ImportState(av io.Reader) error {
	crneg.state = &computeRegionNetworkEndpointGroupState{}
	if err := json.NewDecoder(av).Decode(crneg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crneg.Type(), crneg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionNetworkEndpointGroup] has state.
func (crneg *ComputeRegionNetworkEndpointGroup) State() (*computeRegionNetworkEndpointGroupState, bool) {
	return crneg.state, crneg.state != nil
}

// StateMust returns the state for [ComputeRegionNetworkEndpointGroup]. Panics if the state is nil.
func (crneg *ComputeRegionNetworkEndpointGroup) StateMust() *computeRegionNetworkEndpointGroupState {
	if crneg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crneg.Type(), crneg.LocalName()))
	}
	return crneg.state
}

// ComputeRegionNetworkEndpointGroupArgs contains the configurations for google_compute_region_network_endpoint_group.
type ComputeRegionNetworkEndpointGroupArgs struct {
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
	AppEngine *computeregionnetworkendpointgroup.AppEngine `hcl:"app_engine,block"`
	// CloudFunction: optional
	CloudFunction *computeregionnetworkendpointgroup.CloudFunction `hcl:"cloud_function,block"`
	// CloudRun: optional
	CloudRun *computeregionnetworkendpointgroup.CloudRun `hcl:"cloud_run,block"`
	// ServerlessDeployment: optional
	ServerlessDeployment *computeregionnetworkendpointgroup.ServerlessDeployment `hcl:"serverless_deployment,block"`
	// Timeouts: optional
	Timeouts *computeregionnetworkendpointgroup.Timeouts `hcl:"timeouts,block"`
}
type computeRegionNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("network"))
}

// NetworkEndpointType returns a reference to field network_endpoint_type of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("network_endpoint_type"))
}

// Project returns a reference to field project of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("project"))
}

// PscTargetService returns a reference to field psc_target_service of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) PscTargetService() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("psc_target_service"))
}

// Region returns a reference to field region of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("self_link"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_region_network_endpoint_group.
func (crneg computeRegionNetworkEndpointGroupAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(crneg.ref.Append("subnetwork"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) AppEngine() terra.ListValue[computeregionnetworkendpointgroup.AppEngineAttributes] {
	return terra.ReferenceAsList[computeregionnetworkendpointgroup.AppEngineAttributes](crneg.ref.Append("app_engine"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) CloudFunction() terra.ListValue[computeregionnetworkendpointgroup.CloudFunctionAttributes] {
	return terra.ReferenceAsList[computeregionnetworkendpointgroup.CloudFunctionAttributes](crneg.ref.Append("cloud_function"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) CloudRun() terra.ListValue[computeregionnetworkendpointgroup.CloudRunAttributes] {
	return terra.ReferenceAsList[computeregionnetworkendpointgroup.CloudRunAttributes](crneg.ref.Append("cloud_run"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) ServerlessDeployment() terra.ListValue[computeregionnetworkendpointgroup.ServerlessDeploymentAttributes] {
	return terra.ReferenceAsList[computeregionnetworkendpointgroup.ServerlessDeploymentAttributes](crneg.ref.Append("serverless_deployment"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Timeouts() computeregionnetworkendpointgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionnetworkendpointgroup.TimeoutsAttributes](crneg.ref.Append("timeouts"))
}

type computeRegionNetworkEndpointGroupState struct {
	Description          string                                                        `json:"description"`
	Id                   string                                                        `json:"id"`
	Name                 string                                                        `json:"name"`
	Network              string                                                        `json:"network"`
	NetworkEndpointType  string                                                        `json:"network_endpoint_type"`
	Project              string                                                        `json:"project"`
	PscTargetService     string                                                        `json:"psc_target_service"`
	Region               string                                                        `json:"region"`
	SelfLink             string                                                        `json:"self_link"`
	Subnetwork           string                                                        `json:"subnetwork"`
	AppEngine            []computeregionnetworkendpointgroup.AppEngineState            `json:"app_engine"`
	CloudFunction        []computeregionnetworkendpointgroup.CloudFunctionState        `json:"cloud_function"`
	CloudRun             []computeregionnetworkendpointgroup.CloudRunState             `json:"cloud_run"`
	ServerlessDeployment []computeregionnetworkendpointgroup.ServerlessDeploymentState `json:"serverless_deployment"`
	Timeouts             *computeregionnetworkendpointgroup.TimeoutsState              `json:"timeouts"`
}
