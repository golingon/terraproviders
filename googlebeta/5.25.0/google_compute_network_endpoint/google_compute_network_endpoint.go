// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_network_endpoint

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

// Resource represents the Terraform resource google_compute_network_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeNetworkEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcne *Resource) Type() string {
	return "google_compute_network_endpoint"
}

// LocalName returns the local name for [Resource].
func (gcne *Resource) LocalName() string {
	return gcne.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcne *Resource) Configuration() interface{} {
	return gcne.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcne *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcne)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcne *Resource) Dependencies() terra.Dependencies {
	return gcne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcne *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcne.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcne *Resource) Attributes() googleComputeNetworkEndpointAttributes {
	return googleComputeNetworkEndpointAttributes{ref: terra.ReferenceResource(gcne)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcne *Resource) ImportState(state io.Reader) error {
	gcne.state = &googleComputeNetworkEndpointState{}
	if err := json.NewDecoder(state).Decode(gcne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcne.Type(), gcne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcne *Resource) State() (*googleComputeNetworkEndpointState, bool) {
	return gcne.state, gcne.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcne *Resource) StateMust() *googleComputeNetworkEndpointState {
	if gcne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcne.Type(), gcne.LocalName()))
	}
	return gcne.state
}

// Args contains the configurations for google_compute_network_endpoint.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeNetworkEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("id"))
}

// Instance returns a reference to field instance of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("instance"))
}

// IpAddress returns a reference to field ip_address of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("ip_address"))
}

// NetworkEndpointGroup returns a reference to field network_endpoint_group of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) NetworkEndpointGroup() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("network_endpoint_group"))
}

// Port returns a reference to field port of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(gcne.ref.Append("port"))
}

// Project returns a reference to field project of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_network_endpoint.
func (gcne googleComputeNetworkEndpointAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcne.ref.Append("zone"))
}

func (gcne googleComputeNetworkEndpointAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcne.ref.Append("timeouts"))
}

type googleComputeNetworkEndpointState struct {
	Id                   string         `json:"id"`
	Instance             string         `json:"instance"`
	IpAddress            string         `json:"ip_address"`
	NetworkEndpointGroup string         `json:"network_endpoint_group"`
	Port                 float64        `json:"port"`
	Project              string         `json:"project"`
	Zone                 string         `json:"zone"`
	Timeouts             *TimeoutsState `json:"timeouts"`
}
