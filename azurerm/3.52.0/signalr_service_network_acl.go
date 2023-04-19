// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	signalrservicenetworkacl "github.com/golingon/terraproviders/azurerm/3.52.0/signalrservicenetworkacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSignalrServiceNetworkAcl creates a new instance of [SignalrServiceNetworkAcl].
func NewSignalrServiceNetworkAcl(name string, args SignalrServiceNetworkAclArgs) *SignalrServiceNetworkAcl {
	return &SignalrServiceNetworkAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SignalrServiceNetworkAcl)(nil)

// SignalrServiceNetworkAcl represents the Terraform resource azurerm_signalr_service_network_acl.
type SignalrServiceNetworkAcl struct {
	Name      string
	Args      SignalrServiceNetworkAclArgs
	state     *signalrServiceNetworkAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) Type() string {
	return "azurerm_signalr_service_network_acl"
}

// LocalName returns the local name for [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) LocalName() string {
	return ssna.Name
}

// Configuration returns the configuration (args) for [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) Configuration() interface{} {
	return ssna.Args
}

// DependOn is used for other resources to depend on [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(ssna)
}

// Dependencies returns the list of resources [SignalrServiceNetworkAcl] depends_on.
func (ssna *SignalrServiceNetworkAcl) Dependencies() terra.Dependencies {
	return ssna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) LifecycleManagement() *terra.Lifecycle {
	return ssna.Lifecycle
}

// Attributes returns the attributes for [SignalrServiceNetworkAcl].
func (ssna *SignalrServiceNetworkAcl) Attributes() signalrServiceNetworkAclAttributes {
	return signalrServiceNetworkAclAttributes{ref: terra.ReferenceResource(ssna)}
}

// ImportState imports the given attribute values into [SignalrServiceNetworkAcl]'s state.
func (ssna *SignalrServiceNetworkAcl) ImportState(av io.Reader) error {
	ssna.state = &signalrServiceNetworkAclState{}
	if err := json.NewDecoder(av).Decode(ssna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssna.Type(), ssna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SignalrServiceNetworkAcl] has state.
func (ssna *SignalrServiceNetworkAcl) State() (*signalrServiceNetworkAclState, bool) {
	return ssna.state, ssna.state != nil
}

// StateMust returns the state for [SignalrServiceNetworkAcl]. Panics if the state is nil.
func (ssna *SignalrServiceNetworkAcl) StateMust() *signalrServiceNetworkAclState {
	if ssna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssna.Type(), ssna.LocalName()))
	}
	return ssna.state
}

// SignalrServiceNetworkAclArgs contains the configurations for azurerm_signalr_service_network_acl.
type SignalrServiceNetworkAclArgs struct {
	// DefaultAction: string, required
	DefaultAction terra.StringValue `hcl:"default_action,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SignalrServiceId: string, required
	SignalrServiceId terra.StringValue `hcl:"signalr_service_id,attr" validate:"required"`
	// PrivateEndpoint: min=0
	PrivateEndpoint []signalrservicenetworkacl.PrivateEndpoint `hcl:"private_endpoint,block" validate:"min=0"`
	// PublicNetwork: required
	PublicNetwork *signalrservicenetworkacl.PublicNetwork `hcl:"public_network,block" validate:"required"`
	// Timeouts: optional
	Timeouts *signalrservicenetworkacl.Timeouts `hcl:"timeouts,block"`
}
type signalrServiceNetworkAclAttributes struct {
	ref terra.Reference
}

// DefaultAction returns a reference to field default_action of azurerm_signalr_service_network_acl.
func (ssna signalrServiceNetworkAclAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(ssna.ref.Append("default_action"))
}

// Id returns a reference to field id of azurerm_signalr_service_network_acl.
func (ssna signalrServiceNetworkAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssna.ref.Append("id"))
}

// SignalrServiceId returns a reference to field signalr_service_id of azurerm_signalr_service_network_acl.
func (ssna signalrServiceNetworkAclAttributes) SignalrServiceId() terra.StringValue {
	return terra.ReferenceAsString(ssna.ref.Append("signalr_service_id"))
}

func (ssna signalrServiceNetworkAclAttributes) PrivateEndpoint() terra.SetValue[signalrservicenetworkacl.PrivateEndpointAttributes] {
	return terra.ReferenceAsSet[signalrservicenetworkacl.PrivateEndpointAttributes](ssna.ref.Append("private_endpoint"))
}

func (ssna signalrServiceNetworkAclAttributes) PublicNetwork() terra.ListValue[signalrservicenetworkacl.PublicNetworkAttributes] {
	return terra.ReferenceAsList[signalrservicenetworkacl.PublicNetworkAttributes](ssna.ref.Append("public_network"))
}

func (ssna signalrServiceNetworkAclAttributes) Timeouts() signalrservicenetworkacl.TimeoutsAttributes {
	return terra.ReferenceAsSingle[signalrservicenetworkacl.TimeoutsAttributes](ssna.ref.Append("timeouts"))
}

type signalrServiceNetworkAclState struct {
	DefaultAction    string                                          `json:"default_action"`
	Id               string                                          `json:"id"`
	SignalrServiceId string                                          `json:"signalr_service_id"`
	PrivateEndpoint  []signalrservicenetworkacl.PrivateEndpointState `json:"private_endpoint"`
	PublicNetwork    []signalrservicenetworkacl.PublicNetworkState   `json:"public_network"`
	Timeouts         *signalrservicenetworkacl.TimeoutsState         `json:"timeouts"`
}
