// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeglobalnetworkendpointgroup "github.com/golingon/terraproviders/google/4.59.0/computeglobalnetworkendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeGlobalNetworkEndpointGroup creates a new instance of [ComputeGlobalNetworkEndpointGroup].
func NewComputeGlobalNetworkEndpointGroup(name string, args ComputeGlobalNetworkEndpointGroupArgs) *ComputeGlobalNetworkEndpointGroup {
	return &ComputeGlobalNetworkEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeGlobalNetworkEndpointGroup)(nil)

// ComputeGlobalNetworkEndpointGroup represents the Terraform resource google_compute_global_network_endpoint_group.
type ComputeGlobalNetworkEndpointGroup struct {
	Name      string
	Args      ComputeGlobalNetworkEndpointGroupArgs
	state     *computeGlobalNetworkEndpointGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) Type() string {
	return "google_compute_global_network_endpoint_group"
}

// LocalName returns the local name for [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) LocalName() string {
	return cgneg.Name
}

// Configuration returns the configuration (args) for [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) Configuration() interface{} {
	return cgneg.Args
}

// DependOn is used for other resources to depend on [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cgneg)
}

// Dependencies returns the list of resources [ComputeGlobalNetworkEndpointGroup] depends_on.
func (cgneg *ComputeGlobalNetworkEndpointGroup) Dependencies() terra.Dependencies {
	return cgneg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) LifecycleManagement() *terra.Lifecycle {
	return cgneg.Lifecycle
}

// Attributes returns the attributes for [ComputeGlobalNetworkEndpointGroup].
func (cgneg *ComputeGlobalNetworkEndpointGroup) Attributes() computeGlobalNetworkEndpointGroupAttributes {
	return computeGlobalNetworkEndpointGroupAttributes{ref: terra.ReferenceResource(cgneg)}
}

// ImportState imports the given attribute values into [ComputeGlobalNetworkEndpointGroup]'s state.
func (cgneg *ComputeGlobalNetworkEndpointGroup) ImportState(av io.Reader) error {
	cgneg.state = &computeGlobalNetworkEndpointGroupState{}
	if err := json.NewDecoder(av).Decode(cgneg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgneg.Type(), cgneg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeGlobalNetworkEndpointGroup] has state.
func (cgneg *ComputeGlobalNetworkEndpointGroup) State() (*computeGlobalNetworkEndpointGroupState, bool) {
	return cgneg.state, cgneg.state != nil
}

// StateMust returns the state for [ComputeGlobalNetworkEndpointGroup]. Panics if the state is nil.
func (cgneg *ComputeGlobalNetworkEndpointGroup) StateMust() *computeGlobalNetworkEndpointGroupState {
	if cgneg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgneg.Type(), cgneg.LocalName()))
	}
	return cgneg.state
}

// ComputeGlobalNetworkEndpointGroupArgs contains the configurations for google_compute_global_network_endpoint_group.
type ComputeGlobalNetworkEndpointGroupArgs struct {
	// DefaultPort: number, optional
	DefaultPort terra.NumberValue `hcl:"default_port,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkEndpointType: string, required
	NetworkEndpointType terra.StringValue `hcl:"network_endpoint_type,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computeglobalnetworkendpointgroup.Timeouts `hcl:"timeouts,block"`
}
type computeGlobalNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

// DefaultPort returns a reference to field default_port of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) DefaultPort() terra.NumberValue {
	return terra.ReferenceAsNumber(cgneg.ref.Append("default_port"))
}

// Description returns a reference to field description of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("name"))
}

// NetworkEndpointType returns a reference to field network_endpoint_type of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("network_endpoint_type"))
}

// Project returns a reference to field project of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_global_network_endpoint_group.
func (cgneg computeGlobalNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cgneg.ref.Append("self_link"))
}

func (cgneg computeGlobalNetworkEndpointGroupAttributes) Timeouts() computeglobalnetworkendpointgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeglobalnetworkendpointgroup.TimeoutsAttributes](cgneg.ref.Append("timeouts"))
}

type computeGlobalNetworkEndpointGroupState struct {
	DefaultPort         float64                                          `json:"default_port"`
	Description         string                                           `json:"description"`
	Id                  string                                           `json:"id"`
	Name                string                                           `json:"name"`
	NetworkEndpointType string                                           `json:"network_endpoint_type"`
	Project             string                                           `json:"project"`
	SelfLink            string                                           `json:"self_link"`
	Timeouts            *computeglobalnetworkendpointgroup.TimeoutsState `json:"timeouts"`
}
