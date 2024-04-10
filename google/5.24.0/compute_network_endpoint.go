// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computenetworkendpoint "github.com/golingon/terraproviders/google/5.24.0/computenetworkendpoint"
	"io"
)

// NewComputeNetworkEndpoint creates a new instance of [ComputeNetworkEndpoint].
func NewComputeNetworkEndpoint(name string, args ComputeNetworkEndpointArgs) *ComputeNetworkEndpoint {
	return &ComputeNetworkEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkEndpoint)(nil)

// ComputeNetworkEndpoint represents the Terraform resource google_compute_network_endpoint.
type ComputeNetworkEndpoint struct {
	Name      string
	Args      ComputeNetworkEndpointArgs
	state     *computeNetworkEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) Type() string {
	return "google_compute_network_endpoint"
}

// LocalName returns the local name for [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) LocalName() string {
	return cne.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) Configuration() interface{} {
	return cne.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(cne)
}

// Dependencies returns the list of resources [ComputeNetworkEndpoint] depends_on.
func (cne *ComputeNetworkEndpoint) Dependencies() terra.Dependencies {
	return cne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) LifecycleManagement() *terra.Lifecycle {
	return cne.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkEndpoint].
func (cne *ComputeNetworkEndpoint) Attributes() computeNetworkEndpointAttributes {
	return computeNetworkEndpointAttributes{ref: terra.ReferenceResource(cne)}
}

// ImportState imports the given attribute values into [ComputeNetworkEndpoint]'s state.
func (cne *ComputeNetworkEndpoint) ImportState(av io.Reader) error {
	cne.state = &computeNetworkEndpointState{}
	if err := json.NewDecoder(av).Decode(cne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cne.Type(), cne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkEndpoint] has state.
func (cne *ComputeNetworkEndpoint) State() (*computeNetworkEndpointState, bool) {
	return cne.state, cne.state != nil
}

// StateMust returns the state for [ComputeNetworkEndpoint]. Panics if the state is nil.
func (cne *ComputeNetworkEndpoint) StateMust() *computeNetworkEndpointState {
	if cne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cne.Type(), cne.LocalName()))
	}
	return cne.state
}

// ComputeNetworkEndpointArgs contains the configurations for google_compute_network_endpoint.
type ComputeNetworkEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// NetworkEndpointGroup: string, required
	NetworkEndpointGroup terra.StringValue `hcl:"network_endpoint_group,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computenetworkendpoint.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("instance"))
}

// IpAddress returns a reference to field ip_address of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("ip_address"))
}

// NetworkEndpointGroup returns a reference to field network_endpoint_group of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) NetworkEndpointGroup() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("network_endpoint_group"))
}

// Port returns a reference to field port of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cne.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_network_endpoint.
func (cne computeNetworkEndpointAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cne.ref.Append("zone"))
}

func (cne computeNetworkEndpointAttributes) Timeouts() computenetworkendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkendpoint.TimeoutsAttributes](cne.ref.Append("timeouts"))
}

type computeNetworkEndpointState struct {
	Id                   string                                `json:"id"`
	Instance             string                                `json:"instance"`
	IpAddress            string                                `json:"ip_address"`
	NetworkEndpointGroup string                                `json:"network_endpoint_group"`
	Port                 float64                               `json:"port"`
	Project              string                                `json:"project"`
	Zone                 string                                `json:"zone"`
	Timeouts             *computenetworkendpoint.TimeoutsState `json:"timeouts"`
}
