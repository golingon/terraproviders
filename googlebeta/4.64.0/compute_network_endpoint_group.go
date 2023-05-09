// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computenetworkendpointgroup "github.com/golingon/terraproviders/googlebeta/4.64.0/computenetworkendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkEndpointGroup creates a new instance of [ComputeNetworkEndpointGroup].
func NewComputeNetworkEndpointGroup(name string, args ComputeNetworkEndpointGroupArgs) *ComputeNetworkEndpointGroup {
	return &ComputeNetworkEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkEndpointGroup)(nil)

// ComputeNetworkEndpointGroup represents the Terraform resource google_compute_network_endpoint_group.
type ComputeNetworkEndpointGroup struct {
	Name      string
	Args      ComputeNetworkEndpointGroupArgs
	state     *computeNetworkEndpointGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) Type() string {
	return "google_compute_network_endpoint_group"
}

// LocalName returns the local name for [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) LocalName() string {
	return cneg.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) Configuration() interface{} {
	return cneg.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cneg)
}

// Dependencies returns the list of resources [ComputeNetworkEndpointGroup] depends_on.
func (cneg *ComputeNetworkEndpointGroup) Dependencies() terra.Dependencies {
	return cneg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) LifecycleManagement() *terra.Lifecycle {
	return cneg.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkEndpointGroup].
func (cneg *ComputeNetworkEndpointGroup) Attributes() computeNetworkEndpointGroupAttributes {
	return computeNetworkEndpointGroupAttributes{ref: terra.ReferenceResource(cneg)}
}

// ImportState imports the given attribute values into [ComputeNetworkEndpointGroup]'s state.
func (cneg *ComputeNetworkEndpointGroup) ImportState(av io.Reader) error {
	cneg.state = &computeNetworkEndpointGroupState{}
	if err := json.NewDecoder(av).Decode(cneg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cneg.Type(), cneg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkEndpointGroup] has state.
func (cneg *ComputeNetworkEndpointGroup) State() (*computeNetworkEndpointGroupState, bool) {
	return cneg.state, cneg.state != nil
}

// StateMust returns the state for [ComputeNetworkEndpointGroup]. Panics if the state is nil.
func (cneg *ComputeNetworkEndpointGroup) StateMust() *computeNetworkEndpointGroupState {
	if cneg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cneg.Type(), cneg.LocalName()))
	}
	return cneg.state
}

// ComputeNetworkEndpointGroupArgs contains the configurations for google_compute_network_endpoint_group.
type ComputeNetworkEndpointGroupArgs struct {
	// DefaultPort: number, optional
	DefaultPort terra.NumberValue `hcl:"default_port,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// NetworkEndpointType: string, optional
	NetworkEndpointType terra.StringValue `hcl:"network_endpoint_type,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computenetworkendpointgroup.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

// DefaultPort returns a reference to field default_port of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) DefaultPort() terra.NumberValue {
	return terra.ReferenceAsNumber(cneg.ref.Append("default_port"))
}

// Description returns a reference to field description of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("network"))
}

// NetworkEndpointType returns a reference to field network_endpoint_type of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("network_endpoint_type"))
}

// Project returns a reference to field project of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cneg.ref.Append("size"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("subnetwork"))
}

// Zone returns a reference to field zone of google_compute_network_endpoint_group.
func (cneg computeNetworkEndpointGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cneg.ref.Append("zone"))
}

func (cneg computeNetworkEndpointGroupAttributes) Timeouts() computenetworkendpointgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkendpointgroup.TimeoutsAttributes](cneg.ref.Append("timeouts"))
}

type computeNetworkEndpointGroupState struct {
	DefaultPort         float64                                    `json:"default_port"`
	Description         string                                     `json:"description"`
	Id                  string                                     `json:"id"`
	Name                string                                     `json:"name"`
	Network             string                                     `json:"network"`
	NetworkEndpointType string                                     `json:"network_endpoint_type"`
	Project             string                                     `json:"project"`
	SelfLink            string                                     `json:"self_link"`
	Size                float64                                    `json:"size"`
	Subnetwork          string                                     `json:"subnetwork"`
	Zone                string                                     `json:"zone"`
	Timeouts            *computenetworkendpointgroup.TimeoutsState `json:"timeouts"`
}
