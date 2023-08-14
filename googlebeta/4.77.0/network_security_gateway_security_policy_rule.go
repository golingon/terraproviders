// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networksecuritygatewaysecuritypolicyrule "github.com/golingon/terraproviders/googlebeta/4.77.0/networksecuritygatewaysecuritypolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkSecurityGatewaySecurityPolicyRule creates a new instance of [NetworkSecurityGatewaySecurityPolicyRule].
func NewNetworkSecurityGatewaySecurityPolicyRule(name string, args NetworkSecurityGatewaySecurityPolicyRuleArgs) *NetworkSecurityGatewaySecurityPolicyRule {
	return &NetworkSecurityGatewaySecurityPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkSecurityGatewaySecurityPolicyRule)(nil)

// NetworkSecurityGatewaySecurityPolicyRule represents the Terraform resource google_network_security_gateway_security_policy_rule.
type NetworkSecurityGatewaySecurityPolicyRule struct {
	Name      string
	Args      NetworkSecurityGatewaySecurityPolicyRuleArgs
	state     *networkSecurityGatewaySecurityPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) Type() string {
	return "google_network_security_gateway_security_policy_rule"
}

// LocalName returns the local name for [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) LocalName() string {
	return nsgspr.Name
}

// Configuration returns the configuration (args) for [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) Configuration() interface{} {
	return nsgspr.Args
}

// DependOn is used for other resources to depend on [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(nsgspr)
}

// Dependencies returns the list of resources [NetworkSecurityGatewaySecurityPolicyRule] depends_on.
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) Dependencies() terra.Dependencies {
	return nsgspr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return nsgspr.Lifecycle
}

// Attributes returns the attributes for [NetworkSecurityGatewaySecurityPolicyRule].
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) Attributes() networkSecurityGatewaySecurityPolicyRuleAttributes {
	return networkSecurityGatewaySecurityPolicyRuleAttributes{ref: terra.ReferenceResource(nsgspr)}
}

// ImportState imports the given attribute values into [NetworkSecurityGatewaySecurityPolicyRule]'s state.
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) ImportState(av io.Reader) error {
	nsgspr.state = &networkSecurityGatewaySecurityPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(nsgspr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsgspr.Type(), nsgspr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkSecurityGatewaySecurityPolicyRule] has state.
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) State() (*networkSecurityGatewaySecurityPolicyRuleState, bool) {
	return nsgspr.state, nsgspr.state != nil
}

// StateMust returns the state for [NetworkSecurityGatewaySecurityPolicyRule]. Panics if the state is nil.
func (nsgspr *NetworkSecurityGatewaySecurityPolicyRule) StateMust() *networkSecurityGatewaySecurityPolicyRuleState {
	if nsgspr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsgspr.Type(), nsgspr.LocalName()))
	}
	return nsgspr.state
}

// NetworkSecurityGatewaySecurityPolicyRuleArgs contains the configurations for google_network_security_gateway_security_policy_rule.
type NetworkSecurityGatewaySecurityPolicyRuleArgs struct {
	// ApplicationMatcher: string, optional
	ApplicationMatcher terra.StringValue `hcl:"application_matcher,attr"`
	// BasicProfile: string, required
	BasicProfile terra.StringValue `hcl:"basic_profile,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// GatewaySecurityPolicy: string, required
	GatewaySecurityPolicy terra.StringValue `hcl:"gateway_security_policy,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SessionMatcher: string, required
	SessionMatcher terra.StringValue `hcl:"session_matcher,attr" validate:"required"`
	// TlsInspectionEnabled: bool, optional
	TlsInspectionEnabled terra.BoolValue `hcl:"tls_inspection_enabled,attr"`
	// Timeouts: optional
	Timeouts *networksecuritygatewaysecuritypolicyrule.Timeouts `hcl:"timeouts,block"`
}
type networkSecurityGatewaySecurityPolicyRuleAttributes struct {
	ref terra.Reference
}

// ApplicationMatcher returns a reference to field application_matcher of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) ApplicationMatcher() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("application_matcher"))
}

// BasicProfile returns a reference to field basic_profile of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) BasicProfile() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("basic_profile"))
}

// CreateTime returns a reference to field create_time of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("description"))
}

// Enabled returns a reference to field enabled of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(nsgspr.ref.Append("enabled"))
}

// GatewaySecurityPolicy returns a reference to field gateway_security_policy of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) GatewaySecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("gateway_security_policy"))
}

// Id returns a reference to field id of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("id"))
}

// Location returns a reference to field location of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("location"))
}

// Name returns a reference to field name of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("name"))
}

// Priority returns a reference to field priority of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(nsgspr.ref.Append("priority"))
}

// Project returns a reference to field project of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("self_link"))
}

// SessionMatcher returns a reference to field session_matcher of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) SessionMatcher() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("session_matcher"))
}

// TlsInspectionEnabled returns a reference to field tls_inspection_enabled of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) TlsInspectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nsgspr.ref.Append("tls_inspection_enabled"))
}

// UpdateTime returns a reference to field update_time of google_network_security_gateway_security_policy_rule.
func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgspr.ref.Append("update_time"))
}

func (nsgspr networkSecurityGatewaySecurityPolicyRuleAttributes) Timeouts() networksecuritygatewaysecuritypolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networksecuritygatewaysecuritypolicyrule.TimeoutsAttributes](nsgspr.ref.Append("timeouts"))
}

type networkSecurityGatewaySecurityPolicyRuleState struct {
	ApplicationMatcher    string                                                  `json:"application_matcher"`
	BasicProfile          string                                                  `json:"basic_profile"`
	CreateTime            string                                                  `json:"create_time"`
	Description           string                                                  `json:"description"`
	Enabled               bool                                                    `json:"enabled"`
	GatewaySecurityPolicy string                                                  `json:"gateway_security_policy"`
	Id                    string                                                  `json:"id"`
	Location              string                                                  `json:"location"`
	Name                  string                                                  `json:"name"`
	Priority              float64                                                 `json:"priority"`
	Project               string                                                  `json:"project"`
	SelfLink              string                                                  `json:"self_link"`
	SessionMatcher        string                                                  `json:"session_matcher"`
	TlsInspectionEnabled  bool                                                    `json:"tls_inspection_enabled"`
	UpdateTime            string                                                  `json:"update_time"`
	Timeouts              *networksecuritygatewaysecuritypolicyrule.TimeoutsState `json:"timeouts"`
}
