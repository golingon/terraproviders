// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computehavpngateway "github.com/golingon/terraproviders/googlebeta/4.66.0/computehavpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeHaVpnGateway creates a new instance of [ComputeHaVpnGateway].
func NewComputeHaVpnGateway(name string, args ComputeHaVpnGatewayArgs) *ComputeHaVpnGateway {
	return &ComputeHaVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeHaVpnGateway)(nil)

// ComputeHaVpnGateway represents the Terraform resource google_compute_ha_vpn_gateway.
type ComputeHaVpnGateway struct {
	Name      string
	Args      ComputeHaVpnGatewayArgs
	state     *computeHaVpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) Type() string {
	return "google_compute_ha_vpn_gateway"
}

// LocalName returns the local name for [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) LocalName() string {
	return chvg.Name
}

// Configuration returns the configuration (args) for [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) Configuration() interface{} {
	return chvg.Args
}

// DependOn is used for other resources to depend on [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(chvg)
}

// Dependencies returns the list of resources [ComputeHaVpnGateway] depends_on.
func (chvg *ComputeHaVpnGateway) Dependencies() terra.Dependencies {
	return chvg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) LifecycleManagement() *terra.Lifecycle {
	return chvg.Lifecycle
}

// Attributes returns the attributes for [ComputeHaVpnGateway].
func (chvg *ComputeHaVpnGateway) Attributes() computeHaVpnGatewayAttributes {
	return computeHaVpnGatewayAttributes{ref: terra.ReferenceResource(chvg)}
}

// ImportState imports the given attribute values into [ComputeHaVpnGateway]'s state.
func (chvg *ComputeHaVpnGateway) ImportState(av io.Reader) error {
	chvg.state = &computeHaVpnGatewayState{}
	if err := json.NewDecoder(av).Decode(chvg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", chvg.Type(), chvg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeHaVpnGateway] has state.
func (chvg *ComputeHaVpnGateway) State() (*computeHaVpnGatewayState, bool) {
	return chvg.state, chvg.state != nil
}

// StateMust returns the state for [ComputeHaVpnGateway]. Panics if the state is nil.
func (chvg *ComputeHaVpnGateway) StateMust() *computeHaVpnGatewayState {
	if chvg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", chvg.Type(), chvg.LocalName()))
	}
	return chvg.state
}

// ComputeHaVpnGatewayArgs contains the configurations for google_compute_ha_vpn_gateway.
type ComputeHaVpnGatewayArgs struct {
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
	// StackType: string, optional
	StackType terra.StringValue `hcl:"stack_type,attr"`
	// Timeouts: optional
	Timeouts *computehavpngateway.Timeouts `hcl:"timeouts,block"`
	// VpnInterfaces: min=0
	VpnInterfaces []computehavpngateway.VpnInterfaces `hcl:"vpn_interfaces,block" validate:"min=0"`
}
type computeHaVpnGatewayAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("self_link"))
}

// StackType returns a reference to field stack_type of google_compute_ha_vpn_gateway.
func (chvg computeHaVpnGatewayAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(chvg.ref.Append("stack_type"))
}

func (chvg computeHaVpnGatewayAttributes) Timeouts() computehavpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computehavpngateway.TimeoutsAttributes](chvg.ref.Append("timeouts"))
}

func (chvg computeHaVpnGatewayAttributes) VpnInterfaces() terra.ListValue[computehavpngateway.VpnInterfacesAttributes] {
	return terra.ReferenceAsList[computehavpngateway.VpnInterfacesAttributes](chvg.ref.Append("vpn_interfaces"))
}

type computeHaVpnGatewayState struct {
	Description   string                                   `json:"description"`
	Id            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	Network       string                                   `json:"network"`
	Project       string                                   `json:"project"`
	Region        string                                   `json:"region"`
	SelfLink      string                                   `json:"self_link"`
	StackType     string                                   `json:"stack_type"`
	Timeouts      *computehavpngateway.TimeoutsState       `json:"timeouts"`
	VpnInterfaces []computehavpngateway.VpnInterfacesState `json:"vpn_interfaces"`
}
