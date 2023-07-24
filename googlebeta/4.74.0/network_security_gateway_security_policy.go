// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networksecuritygatewaysecuritypolicy "github.com/golingon/terraproviders/googlebeta/4.74.0/networksecuritygatewaysecuritypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityGatewaySecurityPolicy creates a new instance of [NetworkSecurityGatewaySecurityPolicy].
func NewNetworkSecurityGatewaySecurityPolicy(name string, args NetworkSecurityGatewaySecurityPolicyArgs) *NetworkSecurityGatewaySecurityPolicy {
	return &NetworkSecurityGatewaySecurityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityGatewaySecurityPolicy)(nil)

// NetworkSecurityGatewaySecurityPolicy represents the Terraform resource google_network_security_gateway_security_policy.
type NetworkSecurityGatewaySecurityPolicy struct {
	Name      string
	Args      NetworkSecurityGatewaySecurityPolicyArgs
	state     *networkSecurityGatewaySecurityPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) Type() string {
	return "google_network_security_gateway_security_policy"
}

// LocalName returns the local name for [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) LocalName() string {
	return nsgsp.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) Configuration() interface{} {
	return nsgsp.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(nsgsp)
}

// Dependencies returns the list of resources [NetworkSecurityGatewaySecurityPolicy] depends_on.
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) Dependencies() terra.Dependencies {
	return nsgsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) LifecycleManagement() *terra.Lifecycle {
	return nsgsp.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityGatewaySecurityPolicy].
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) Attributes() networkSecurityGatewaySecurityPolicyAttributes {
	return networkSecurityGatewaySecurityPolicyAttributes{ref: terra.ReferenceResource(nsgsp)}
}

// ImportState imports the given attribute values into [NetworkSecurityGatewaySecurityPolicy]'s state.
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) ImportState(av io.Reader) error {
	nsgsp.state = &networkSecurityGatewaySecurityPolicyState{}
	if err := json.NewDecoder(av).Decode(nsgsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsgsp.Type(), nsgsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityGatewaySecurityPolicy] has state.
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) State() (*networkSecurityGatewaySecurityPolicyState, bool) {
	return nsgsp.state, nsgsp.state != nil
}

// StateMust returns the state for [NetworkSecurityGatewaySecurityPolicy]. Panics if the state is nil.
func (nsgsp *NetworkSecurityGatewaySecurityPolicy) StateMust() *networkSecurityGatewaySecurityPolicyState {
	if nsgsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsgsp.Type(), nsgsp.LocalName()))
	}
	return nsgsp.state
}

// NetworkSecurityGatewaySecurityPolicyArgs contains the configurations for google_network_security_gateway_security_policy.
type NetworkSecurityGatewaySecurityPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TlsInspectionPolicy: string, optional
	TlsInspectionPolicy terra.StringValue `hcl:"tls_inspection_policy,attr"`
	// Timeouts: optional
	Timeouts *networksecuritygatewaysecuritypolicy.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityGatewaySecurityPolicyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("description"))
}

// Id returns a reference to field id of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("name"))
}

// Project returns a reference to field project of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("self_link"))
}

// TlsInspectionPolicy returns a reference to field tls_inspection_policy of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) TlsInspectionPolicy() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("tls_inspection_policy"))
}

// UpdateTime returns a reference to field update_time of google_network_security_gateway_security_policy.
func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgsp.ref.Append("update_time"))
}

func (nsgsp networkSecurityGatewaySecurityPolicyAttributes) Timeouts() networksecuritygatewaysecuritypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecuritygatewaysecuritypolicy.TimeoutsAttributes](nsgsp.ref.Append("timeouts"))
}

type networkSecurityGatewaySecurityPolicyState struct {
	CreateTime          string                                              `json:"create_time"`
	Description         string                                              `json:"description"`
	Id                  string                                              `json:"id"`
	Location            string                                              `json:"location"`
	Name                string                                              `json:"name"`
	Project             string                                              `json:"project"`
	SelfLink            string                                              `json:"self_link"`
	TlsInspectionPolicy string                                              `json:"tls_inspection_policy"`
	UpdateTime          string                                              `json:"update_time"`
	Timeouts            *networksecuritygatewaysecuritypolicy.TimeoutsState `json:"timeouts"`
}
