// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_vpn_gateway

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

// Resource represents the Terraform resource google_compute_vpn_gateway.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeVpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcvg *Resource) Type() string {
	return "google_compute_vpn_gateway"
}

// LocalName returns the local name for [Resource].
func (gcvg *Resource) LocalName() string {
	return gcvg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcvg *Resource) Configuration() interface{} {
	return gcvg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcvg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcvg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcvg *Resource) Dependencies() terra.Dependencies {
	return gcvg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcvg *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcvg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcvg *Resource) Attributes() googleComputeVpnGatewayAttributes {
	return googleComputeVpnGatewayAttributes{ref: terra.ReferenceResource(gcvg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcvg *Resource) ImportState(state io.Reader) error {
	gcvg.state = &googleComputeVpnGatewayState{}
	if err := json.NewDecoder(state).Decode(gcvg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcvg.Type(), gcvg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcvg *Resource) State() (*googleComputeVpnGatewayState, bool) {
	return gcvg.state, gcvg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcvg *Resource) StateMust() *googleComputeVpnGatewayState {
	if gcvg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcvg.Type(), gcvg.LocalName()))
	}
	return gcvg.state
}

// Args contains the configurations for google_compute_vpn_gateway.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeVpnGatewayAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("description"))
}

// GatewayId returns a reference to field gateway_id of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) GatewayId() terra.NumberValue {
	return terra.ReferenceAsNumber(gcvg.ref.Append("gateway_id"))
}

// Id returns a reference to field id of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_vpn_gateway.
func (gcvg googleComputeVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcvg.ref.Append("self_link"))
}

func (gcvg googleComputeVpnGatewayAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcvg.ref.Append("timeouts"))
}

type googleComputeVpnGatewayState struct {
	CreationTimestamp string         `json:"creation_timestamp"`
	Description       string         `json:"description"`
	GatewayId         float64        `json:"gateway_id"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	Network           string         `json:"network"`
	Project           string         `json:"project"`
	Region            string         `json:"region"`
	SelfLink          string         `json:"self_link"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
