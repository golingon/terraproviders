// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	relayhybridconnectionauthorizationrule "github.com/golingon/terraproviders/azurerm/3.52.0/relayhybridconnectionauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRelayHybridConnectionAuthorizationRule creates a new instance of [RelayHybridConnectionAuthorizationRule].
func NewRelayHybridConnectionAuthorizationRule(name string, args RelayHybridConnectionAuthorizationRuleArgs) *RelayHybridConnectionAuthorizationRule {
	return &RelayHybridConnectionAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RelayHybridConnectionAuthorizationRule)(nil)

// RelayHybridConnectionAuthorizationRule represents the Terraform resource azurerm_relay_hybrid_connection_authorization_rule.
type RelayHybridConnectionAuthorizationRule struct {
	Name      string
	Args      RelayHybridConnectionAuthorizationRuleArgs
	state     *relayHybridConnectionAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) Type() string {
	return "azurerm_relay_hybrid_connection_authorization_rule"
}

// LocalName returns the local name for [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) LocalName() string {
	return rhcar.Name
}

// Configuration returns the configuration (args) for [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) Configuration() interface{} {
	return rhcar.Args
}

// DependOn is used for other resources to depend on [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rhcar)
}

// Dependencies returns the list of resources [RelayHybridConnectionAuthorizationRule] depends_on.
func (rhcar *RelayHybridConnectionAuthorizationRule) Dependencies() terra.Dependencies {
	return rhcar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return rhcar.Lifecycle
}

// Attributes returns the attributes for [RelayHybridConnectionAuthorizationRule].
func (rhcar *RelayHybridConnectionAuthorizationRule) Attributes() relayHybridConnectionAuthorizationRuleAttributes {
	return relayHybridConnectionAuthorizationRuleAttributes{ref: terra.ReferenceResource(rhcar)}
}

// ImportState imports the given attribute values into [RelayHybridConnectionAuthorizationRule]'s state.
func (rhcar *RelayHybridConnectionAuthorizationRule) ImportState(av io.Reader) error {
	rhcar.state = &relayHybridConnectionAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(rhcar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rhcar.Type(), rhcar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RelayHybridConnectionAuthorizationRule] has state.
func (rhcar *RelayHybridConnectionAuthorizationRule) State() (*relayHybridConnectionAuthorizationRuleState, bool) {
	return rhcar.state, rhcar.state != nil
}

// StateMust returns the state for [RelayHybridConnectionAuthorizationRule]. Panics if the state is nil.
func (rhcar *RelayHybridConnectionAuthorizationRule) StateMust() *relayHybridConnectionAuthorizationRuleState {
	if rhcar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rhcar.Type(), rhcar.LocalName()))
	}
	return rhcar.state
}

// RelayHybridConnectionAuthorizationRuleArgs contains the configurations for azurerm_relay_hybrid_connection_authorization_rule.
type RelayHybridConnectionAuthorizationRuleArgs struct {
	// HybridConnectionName: string, required
	HybridConnectionName terra.StringValue `hcl:"hybrid_connection_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *relayhybridconnectionauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type relayHybridConnectionAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// HybridConnectionName returns a reference to field hybrid_connection_name of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) HybridConnectionName() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("hybrid_connection_name"))
}

// Id returns a reference to field id of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(rhcar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(rhcar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(rhcar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_relay_hybrid_connection_authorization_rule.
func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(rhcar.ref.Append("send"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Timeouts() relayhybridconnectionauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[relayhybridconnectionauthorizationrule.TimeoutsAttributes](rhcar.ref.Append("timeouts"))
}

type relayHybridConnectionAuthorizationRuleState struct {
	HybridConnectionName      string                                                `json:"hybrid_connection_name"`
	Id                        string                                                `json:"id"`
	Listen                    bool                                                  `json:"listen"`
	Manage                    bool                                                  `json:"manage"`
	Name                      string                                                `json:"name"`
	NamespaceName             string                                                `json:"namespace_name"`
	PrimaryConnectionString   string                                                `json:"primary_connection_string"`
	PrimaryKey                string                                                `json:"primary_key"`
	ResourceGroupName         string                                                `json:"resource_group_name"`
	SecondaryConnectionString string                                                `json:"secondary_connection_string"`
	SecondaryKey              string                                                `json:"secondary_key"`
	Send                      bool                                                  `json:"send"`
	Timeouts                  *relayhybridconnectionauthorizationrule.TimeoutsState `json:"timeouts"`
}
