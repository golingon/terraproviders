// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	relayhybridconnectionauthorizationrule "github.com/golingon/terraproviders/azurerm/3.49.0/relayhybridconnectionauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRelayHybridConnectionAuthorizationRule(name string, args RelayHybridConnectionAuthorizationRuleArgs) *RelayHybridConnectionAuthorizationRule {
	return &RelayHybridConnectionAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RelayHybridConnectionAuthorizationRule)(nil)

type RelayHybridConnectionAuthorizationRule struct {
	Name  string
	Args  RelayHybridConnectionAuthorizationRuleArgs
	state *relayHybridConnectionAuthorizationRuleState
}

func (rhcar *RelayHybridConnectionAuthorizationRule) Type() string {
	return "azurerm_relay_hybrid_connection_authorization_rule"
}

func (rhcar *RelayHybridConnectionAuthorizationRule) LocalName() string {
	return rhcar.Name
}

func (rhcar *RelayHybridConnectionAuthorizationRule) Configuration() interface{} {
	return rhcar.Args
}

func (rhcar *RelayHybridConnectionAuthorizationRule) Attributes() relayHybridConnectionAuthorizationRuleAttributes {
	return relayHybridConnectionAuthorizationRuleAttributes{ref: terra.ReferenceResource(rhcar)}
}

func (rhcar *RelayHybridConnectionAuthorizationRule) ImportState(av io.Reader) error {
	rhcar.state = &relayHybridConnectionAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(rhcar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rhcar.Type(), rhcar.LocalName(), err)
	}
	return nil
}

func (rhcar *RelayHybridConnectionAuthorizationRule) State() (*relayHybridConnectionAuthorizationRuleState, bool) {
	return rhcar.state, rhcar.state != nil
}

func (rhcar *RelayHybridConnectionAuthorizationRule) StateMust() *relayHybridConnectionAuthorizationRuleState {
	if rhcar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rhcar.Type(), rhcar.LocalName()))
	}
	return rhcar.state
}

func (rhcar *RelayHybridConnectionAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rhcar)
}

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
	// DependsOn contains resources that RelayHybridConnectionAuthorizationRule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type relayHybridConnectionAuthorizationRuleAttributes struct {
	ref terra.Reference
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) HybridConnectionName() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("hybrid_connection_name"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("id"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceBool(rhcar.ref.Append("listen"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceBool(rhcar.ref.Append("manage"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("name"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("namespace_name"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("primary_connection_string"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("primary_key"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("resource_group_name"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("secondary_connection_string"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceString(rhcar.ref.Append("secondary_key"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceBool(rhcar.ref.Append("send"))
}

func (rhcar relayHybridConnectionAuthorizationRuleAttributes) Timeouts() relayhybridconnectionauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceSingle[relayhybridconnectionauthorizationrule.TimeoutsAttributes](rhcar.ref.Append("timeouts"))
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
