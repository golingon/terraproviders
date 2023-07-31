// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeglobalnetworkendpoint "github.com/golingon/terraproviders/google/4.75.1/computeglobalnetworkendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeGlobalNetworkEndpoint creates a new instance of [ComputeGlobalNetworkEndpoint].
func NewComputeGlobalNetworkEndpoint(name string, args ComputeGlobalNetworkEndpointArgs) *ComputeGlobalNetworkEndpoint {
	return &ComputeGlobalNetworkEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeGlobalNetworkEndpoint)(nil)

// ComputeGlobalNetworkEndpoint represents the Terraform resource google_compute_global_network_endpoint.
type ComputeGlobalNetworkEndpoint struct {
	Name      string
	Args      ComputeGlobalNetworkEndpointArgs
	state     *computeGlobalNetworkEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) Type() string {
	return "google_compute_global_network_endpoint"
}

// LocalName returns the local name for [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) LocalName() string {
	return cgne.Name
}

// Configuration returns the configuration (args) for [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) Configuration() interface{} {
	return cgne.Args
}

// DependOn is used for other resources to depend on [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(cgne)
}

// Dependencies returns the list of resources [ComputeGlobalNetworkEndpoint] depends_on.
func (cgne *ComputeGlobalNetworkEndpoint) Dependencies() terra.Dependencies {
	return cgne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) LifecycleManagement() *terra.Lifecycle {
	return cgne.Lifecycle
}

// Attributes returns the attributes for [ComputeGlobalNetworkEndpoint].
func (cgne *ComputeGlobalNetworkEndpoint) Attributes() computeGlobalNetworkEndpointAttributes {
	return computeGlobalNetworkEndpointAttributes{ref: terra.ReferenceResource(cgne)}
}

// ImportState imports the given attribute values into [ComputeGlobalNetworkEndpoint]'s state.
func (cgne *ComputeGlobalNetworkEndpoint) ImportState(av io.Reader) error {
	cgne.state = &computeGlobalNetworkEndpointState{}
	if err := json.NewDecoder(av).Decode(cgne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgne.Type(), cgne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeGlobalNetworkEndpoint] has state.
func (cgne *ComputeGlobalNetworkEndpoint) State() (*computeGlobalNetworkEndpointState, bool) {
	return cgne.state, cgne.state != nil
}

// StateMust returns the state for [ComputeGlobalNetworkEndpoint]. Panics if the state is nil.
func (cgne *ComputeGlobalNetworkEndpoint) StateMust() *computeGlobalNetworkEndpointState {
	if cgne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgne.Type(), cgne.LocalName()))
	}
	return cgne.state
}

// ComputeGlobalNetworkEndpointArgs contains the configurations for google_compute_global_network_endpoint.
type ComputeGlobalNetworkEndpointArgs struct {
	// Fqdn: string, optional
	Fqdn terra.StringValue `hcl:"fqdn,attr"`
	// GlobalNetworkEndpointGroup: string, required
	GlobalNetworkEndpointGroup terra.StringValue `hcl:"global_network_endpoint_group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computeglobalnetworkendpoint.Timeouts `hcl:"timeouts,block"`
}
type computeGlobalNetworkEndpointAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(cgne.ref.Append("fqdn"))
}

// GlobalNetworkEndpointGroup returns a reference to field global_network_endpoint_group of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) GlobalNetworkEndpointGroup() terra.StringValue {
	return terra.ReferenceAsString(cgne.ref.Append("global_network_endpoint_group"))
}

// Id returns a reference to field id of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgne.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cgne.ref.Append("ip_address"))
}

// Port returns a reference to field port of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cgne.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_global_network_endpoint.
func (cgne computeGlobalNetworkEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cgne.ref.Append("project"))
}

func (cgne computeGlobalNetworkEndpointAttributes) Timeouts() computeglobalnetworkendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeglobalnetworkendpoint.TimeoutsAttributes](cgne.ref.Append("timeouts"))
}

type computeGlobalNetworkEndpointState struct {
	Fqdn                       string                                      `json:"fqdn"`
	GlobalNetworkEndpointGroup string                                      `json:"global_network_endpoint_group"`
	Id                         string                                      `json:"id"`
	IpAddress                  string                                      `json:"ip_address"`
	Port                       float64                                     `json:"port"`
	Project                    string                                      `json:"project"`
	Timeouts                   *computeglobalnetworkendpoint.TimeoutsState `json:"timeouts"`
}
