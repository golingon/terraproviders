// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computevpngateway "github.com/golingon/terraproviders/google/4.77.0/computevpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeVpnGateway creates a new instance of [ComputeVpnGateway].
func NewComputeVpnGateway(name string, args ComputeVpnGatewayArgs) *ComputeVpnGateway {
	return &ComputeVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeVpnGateway)(nil)

// ComputeVpnGateway represents the Terraform resource google_compute_vpn_gateway.
type ComputeVpnGateway struct {
	Name      string
	Args      ComputeVpnGatewayArgs
	state     *computeVpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) Type() string {
	return "google_compute_vpn_gateway"
}

// LocalName returns the local name for [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) LocalName() string {
	return cvg.Name
}

// Configuration returns the configuration (args) for [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) Configuration() interface{} {
	return cvg.Args
}

// DependOn is used for other resources to depend on [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(cvg)
}

// Dependencies returns the list of resources [ComputeVpnGateway] depends_on.
func (cvg *ComputeVpnGateway) Dependencies() terra.Dependencies {
	return cvg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) LifecycleManagement() *terra.Lifecycle {
	return cvg.Lifecycle
}

// Attributes returns the attributes for [ComputeVpnGateway].
func (cvg *ComputeVpnGateway) Attributes() computeVpnGatewayAttributes {
	return computeVpnGatewayAttributes{ref: terra.ReferenceResource(cvg)}
}

// ImportState imports the given attribute values into [ComputeVpnGateway]'s state.
func (cvg *ComputeVpnGateway) ImportState(av io.Reader) error {
	cvg.state = &computeVpnGatewayState{}
	if err := json.NewDecoder(av).Decode(cvg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvg.Type(), cvg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeVpnGateway] has state.
func (cvg *ComputeVpnGateway) State() (*computeVpnGatewayState, bool) {
	return cvg.state, cvg.state != nil
}

// StateMust returns the state for [ComputeVpnGateway]. Panics if the state is nil.
func (cvg *ComputeVpnGateway) StateMust() *computeVpnGatewayState {
	if cvg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvg.Type(), cvg.LocalName()))
	}
	return cvg.state
}

// ComputeVpnGatewayArgs contains the configurations for google_compute_vpn_gateway.
type ComputeVpnGatewayArgs struct {
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
	Timeouts *computevpngateway.Timeouts `hcl:"timeouts,block"`
}
type computeVpnGatewayAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("description"))
}

// GatewayId returns a reference to field gateway_id of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) GatewayId() terra.NumberValue {
	return terra.ReferenceAsNumber(cvg.ref.Append("gateway_id"))
}

// Id returns a reference to field id of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("network"))
}

// Project returns a reference to field project of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_vpn_gateway.
func (cvg computeVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cvg.ref.Append("self_link"))
}

func (cvg computeVpnGatewayAttributes) Timeouts() computevpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computevpngateway.TimeoutsAttributes](cvg.ref.Append("timeouts"))
}

type computeVpnGatewayState struct {
	CreationTimestamp string                           `json:"creation_timestamp"`
	Description       string                           `json:"description"`
	GatewayId         float64                          `json:"gateway_id"`
	Id                string                           `json:"id"`
	Name              string                           `json:"name"`
	Network           string                           `json:"network"`
	Project           string                           `json:"project"`
	Region            string                           `json:"region"`
	SelfLink          string                           `json:"self_link"`
	Timeouts          *computevpngateway.TimeoutsState `json:"timeouts"`
}
