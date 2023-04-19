// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeexternalvpngateway "github.com/golingon/terraproviders/google/4.62.0/computeexternalvpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeExternalVpnGateway creates a new instance of [ComputeExternalVpnGateway].
func NewComputeExternalVpnGateway(name string, args ComputeExternalVpnGatewayArgs) *ComputeExternalVpnGateway {
	return &ComputeExternalVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeExternalVpnGateway)(nil)

// ComputeExternalVpnGateway represents the Terraform resource google_compute_external_vpn_gateway.
type ComputeExternalVpnGateway struct {
	Name      string
	Args      ComputeExternalVpnGatewayArgs
	state     *computeExternalVpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) Type() string {
	return "google_compute_external_vpn_gateway"
}

// LocalName returns the local name for [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) LocalName() string {
	return cevg.Name
}

// Configuration returns the configuration (args) for [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) Configuration() interface{} {
	return cevg.Args
}

// DependOn is used for other resources to depend on [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(cevg)
}

// Dependencies returns the list of resources [ComputeExternalVpnGateway] depends_on.
func (cevg *ComputeExternalVpnGateway) Dependencies() terra.Dependencies {
	return cevg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) LifecycleManagement() *terra.Lifecycle {
	return cevg.Lifecycle
}

// Attributes returns the attributes for [ComputeExternalVpnGateway].
func (cevg *ComputeExternalVpnGateway) Attributes() computeExternalVpnGatewayAttributes {
	return computeExternalVpnGatewayAttributes{ref: terra.ReferenceResource(cevg)}
}

// ImportState imports the given attribute values into [ComputeExternalVpnGateway]'s state.
func (cevg *ComputeExternalVpnGateway) ImportState(av io.Reader) error {
	cevg.state = &computeExternalVpnGatewayState{}
	if err := json.NewDecoder(av).Decode(cevg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cevg.Type(), cevg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeExternalVpnGateway] has state.
func (cevg *ComputeExternalVpnGateway) State() (*computeExternalVpnGatewayState, bool) {
	return cevg.state, cevg.state != nil
}

// StateMust returns the state for [ComputeExternalVpnGateway]. Panics if the state is nil.
func (cevg *ComputeExternalVpnGateway) StateMust() *computeExternalVpnGatewayState {
	if cevg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cevg.Type(), cevg.LocalName()))
	}
	return cevg.state
}

// ComputeExternalVpnGatewayArgs contains the configurations for google_compute_external_vpn_gateway.
type ComputeExternalVpnGatewayArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RedundancyType: string, optional
	RedundancyType terra.StringValue `hcl:"redundancy_type,attr"`
	// Interface: min=0
	Interface []computeexternalvpngateway.Interface `hcl:"interface,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeexternalvpngateway.Timeouts `hcl:"timeouts,block"`
}
type computeExternalVpnGatewayAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("id"))
}

// Labels returns a reference to field labels of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cevg.ref.Append("labels"))
}

// Name returns a reference to field name of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("project"))
}

// RedundancyType returns a reference to field redundancy_type of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) RedundancyType() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("redundancy_type"))
}

// SelfLink returns a reference to field self_link of google_compute_external_vpn_gateway.
func (cevg computeExternalVpnGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cevg.ref.Append("self_link"))
}

func (cevg computeExternalVpnGatewayAttributes) Interface() terra.ListValue[computeexternalvpngateway.InterfaceAttributes] {
	return terra.ReferenceAsList[computeexternalvpngateway.InterfaceAttributes](cevg.ref.Append("interface"))
}

func (cevg computeExternalVpnGatewayAttributes) Timeouts() computeexternalvpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeexternalvpngateway.TimeoutsAttributes](cevg.ref.Append("timeouts"))
}

type computeExternalVpnGatewayState struct {
	Description    string                                     `json:"description"`
	Id             string                                     `json:"id"`
	Labels         map[string]string                          `json:"labels"`
	Name           string                                     `json:"name"`
	Project        string                                     `json:"project"`
	RedundancyType string                                     `json:"redundancy_type"`
	SelfLink       string                                     `json:"self_link"`
	Interface      []computeexternalvpngateway.InterfaceState `json:"interface"`
	Timeouts       *computeexternalvpngateway.TimeoutsState   `json:"timeouts"`
}
