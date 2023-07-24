// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenetworkendpoints "github.com/golingon/terraproviders/google/4.74.0/computenetworkendpoints"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkEndpoints creates a new instance of [ComputeNetworkEndpoints].
func NewComputeNetworkEndpoints(name string, args ComputeNetworkEndpointsArgs) *ComputeNetworkEndpoints {
	return &ComputeNetworkEndpoints{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkEndpoints)(nil)

// ComputeNetworkEndpoints represents the Terraform resource google_compute_network_endpoints.
type ComputeNetworkEndpoints struct {
	Name      string
	Args      ComputeNetworkEndpointsArgs
	state     *computeNetworkEndpointsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) Type() string {
	return "google_compute_network_endpoints"
}

// LocalName returns the local name for [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) LocalName() string {
	return cne.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) Configuration() interface{} {
	return cne.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) DependOn() terra.Reference {
	return terra.ReferenceResource(cne)
}

// Dependencies returns the list of resources [ComputeNetworkEndpoints] depends_on.
func (cne *ComputeNetworkEndpoints) Dependencies() terra.Dependencies {
	return cne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) LifecycleManagement() *terra.Lifecycle {
	return cne.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkEndpoints].
func (cne *ComputeNetworkEndpoints) Attributes() computeNetworkEndpointsAttributes {
	return computeNetworkEndpointsAttributes{ref: terra.ReferenceResource(cne)}
}

// ImportState imports the given attribute values into [ComputeNetworkEndpoints]'s state.
func (cne *ComputeNetworkEndpoints) ImportState(av io.Reader) error {
	cne.state = &computeNetworkEndpointsState{}
	if err := json.NewDecoder(av).Decode(cne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cne.Type(), cne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkEndpoints] has state.
func (cne *ComputeNetworkEndpoints) State() (*computeNetworkEndpointsState, bool) {
	return cne.state, cne.state != nil
}

// StateMust returns the state for [ComputeNetworkEndpoints]. Panics if the state is nil.
func (cne *ComputeNetworkEndpoints) StateMust() *computeNetworkEndpointsState {
	if cne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cne.Type(), cne.LocalName()))
	}
	return cne.state
}

// ComputeNetworkEndpointsArgs contains the configurations for google_compute_network_endpoints.
type ComputeNetworkEndpointsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkEndpointGroup: string, required
	NetworkEndpointGroup terra.StringValue `hcl:"network_endpoint_group,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// NetworkEndpoints: min=0
	NetworkEndpoints []computenetworkendpoints.NetworkEndpoints `hcl:"network_endpoints,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computenetworkendpoints.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkEndpointsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_network_endpoints.
func (cne computeNetworkEndpointsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("id"))
}

// NetworkEndpointGroup returns a reference to field network_endpoint_group of google_compute_network_endpoints.
func (cne computeNetworkEndpointsAttributes) NetworkEndpointGroup() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("network_endpoint_group"))
}

// Project returns a reference to field project of google_compute_network_endpoints.
func (cne computeNetworkEndpointsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_network_endpoints.
func (cne computeNetworkEndpointsAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("zone"))
}

func (cne computeNetworkEndpointsAttributes) NetworkEndpoints() terra.SetValue[computenetworkendpoints.NetworkEndpointsAttributes] {
	return terra.ReferenceAsSet[computenetworkendpoints.NetworkEndpointsAttributes](cne.ref.Append("network_endpoints"))
}

func (cne computeNetworkEndpointsAttributes) Timeouts() computenetworkendpoints.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkendpoints.TimeoutsAttributes](cne.ref.Append("timeouts"))
}

type computeNetworkEndpointsState struct {
	Id                   string                                          `json:"id"`
	NetworkEndpointGroup string                                          `json:"network_endpoint_group"`
	Project              string                                          `json:"project"`
	Zone                 string                                          `json:"zone"`
	NetworkEndpoints     []computenetworkendpoints.NetworkEndpointsState `json:"network_endpoints"`
	Timeouts             *computenetworkendpoints.TimeoutsState          `json:"timeouts"`
}
