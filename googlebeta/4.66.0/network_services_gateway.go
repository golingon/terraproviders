// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicesgateway "github.com/golingon/terraproviders/googlebeta/4.66.0/networkservicesgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesGateway creates a new instance of [NetworkServicesGateway].
func NewNetworkServicesGateway(name string, args NetworkServicesGatewayArgs) *NetworkServicesGateway {
	return &NetworkServicesGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesGateway)(nil)

// NetworkServicesGateway represents the Terraform resource google_network_services_gateway.
type NetworkServicesGateway struct {
	Name      string
	Args      NetworkServicesGatewayArgs
	state     *networkServicesGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) Type() string {
	return "google_network_services_gateway"
}

// LocalName returns the local name for [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) LocalName() string {
	return nsg.Name
}

// Configuration returns the configuration (args) for [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) Configuration() interface{} {
	return nsg.Args
}

// DependOn is used for other resources to depend on [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(nsg)
}

// Dependencies returns the list of resources [NetworkServicesGateway] depends_on.
func (nsg *NetworkServicesGateway) Dependencies() terra.Dependencies {
	return nsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) LifecycleManagement() *terra.Lifecycle {
	return nsg.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesGateway].
func (nsg *NetworkServicesGateway) Attributes() networkServicesGatewayAttributes {
	return networkServicesGatewayAttributes{ref: terra.ReferenceResource(nsg)}
}

// ImportState imports the given attribute values into [NetworkServicesGateway]'s state.
func (nsg *NetworkServicesGateway) ImportState(av io.Reader) error {
	nsg.state = &networkServicesGatewayState{}
	if err := json.NewDecoder(av).Decode(nsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsg.Type(), nsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesGateway] has state.
func (nsg *NetworkServicesGateway) State() (*networkServicesGatewayState, bool) {
	return nsg.state, nsg.state != nil
}

// StateMust returns the state for [NetworkServicesGateway]. Panics if the state is nil.
func (nsg *NetworkServicesGateway) StateMust() *networkServicesGatewayState {
	if nsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsg.Type(), nsg.LocalName()))
	}
	return nsg.state
}

// NetworkServicesGatewayArgs contains the configurations for google_network_services_gateway.
type NetworkServicesGatewayArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Ports: list of number, required
	Ports terra.ListValue[terra.NumberValue] `hcl:"ports,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// ServerTlsPolicy: string, optional
	ServerTlsPolicy terra.StringValue `hcl:"server_tls_policy,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkservicesgateway.Timeouts `hcl:"timeouts,block"`
}
type networkServicesGatewayAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("description"))
}

// Id returns a reference to field id of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsg.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("location"))
}

// Name returns a reference to field name of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("name"))
}

// Ports returns a reference to field ports of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Ports() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](nsg.ref.Append("ports"))
}

// Project returns a reference to field project of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("project"))
}

// Scope returns a reference to field scope of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("scope"))
}

// SelfLink returns a reference to field self_link of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("self_link"))
}

// ServerTlsPolicy returns a reference to field server_tls_policy of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) ServerTlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("server_tls_policy"))
}

// Type returns a reference to field type of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("type"))
}

// UpdateTime returns a reference to field update_time of google_network_services_gateway.
func (nsg networkServicesGatewayAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsg.ref.Append("update_time"))
}

func (nsg networkServicesGatewayAttributes) Timeouts() networkservicesgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesgateway.TimeoutsAttributes](nsg.ref.Append("timeouts"))
}

type networkServicesGatewayState struct {
	CreateTime      string                                `json:"create_time"`
	Description     string                                `json:"description"`
	Id              string                                `json:"id"`
	Labels          map[string]string                     `json:"labels"`
	Location        string                                `json:"location"`
	Name            string                                `json:"name"`
	Ports           []float64                             `json:"ports"`
	Project         string                                `json:"project"`
	Scope           string                                `json:"scope"`
	SelfLink        string                                `json:"self_link"`
	ServerTlsPolicy string                                `json:"server_tls_policy"`
	Type            string                                `json:"type"`
	UpdateTime      string                                `json:"update_time"`
	Timeouts        *networkservicesgateway.TimeoutsState `json:"timeouts"`
}
