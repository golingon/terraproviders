// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	appenginefirewallrule "github.com/golingon/terraproviders/google/4.76.0/appenginefirewallrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineFirewallRule creates a new instance of [AppEngineFirewallRule].
func NewAppEngineFirewallRule(name string, args AppEngineFirewallRuleArgs) *AppEngineFirewallRule {
	return &AppEngineFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineFirewallRule)(nil)

// AppEngineFirewallRule represents the Terraform resource google_app_engine_firewall_rule.
type AppEngineFirewallRule struct {
	Name      string
	Args      AppEngineFirewallRuleArgs
	state     *appEngineFirewallRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) Type() string {
	return "google_app_engine_firewall_rule"
}

// LocalName returns the local name for [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) LocalName() string {
	return aefr.Name
}

// Configuration returns the configuration (args) for [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) Configuration() interface{} {
	return aefr.Args
}

// DependOn is used for other resources to depend on [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(aefr)
}

// Dependencies returns the list of resources [AppEngineFirewallRule] depends_on.
func (aefr *AppEngineFirewallRule) Dependencies() terra.Dependencies {
	return aefr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) LifecycleManagement() *terra.Lifecycle {
	return aefr.Lifecycle
}

// Attributes returns the attributes for [AppEngineFirewallRule].
func (aefr *AppEngineFirewallRule) Attributes() appEngineFirewallRuleAttributes {
	return appEngineFirewallRuleAttributes{ref: terra.ReferenceResource(aefr)}
}

// ImportState imports the given attribute values into [AppEngineFirewallRule]'s state.
func (aefr *AppEngineFirewallRule) ImportState(av io.Reader) error {
	aefr.state = &appEngineFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(aefr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aefr.Type(), aefr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineFirewallRule] has state.
func (aefr *AppEngineFirewallRule) State() (*appEngineFirewallRuleState, bool) {
	return aefr.state, aefr.state != nil
}

// StateMust returns the state for [AppEngineFirewallRule]. Panics if the state is nil.
func (aefr *AppEngineFirewallRule) StateMust() *appEngineFirewallRuleState {
	if aefr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aefr.Type(), aefr.LocalName()))
	}
	return aefr.state
}

// AppEngineFirewallRuleArgs contains the configurations for google_app_engine_firewall_rule.
type AppEngineFirewallRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceRange: string, required
	SourceRange terra.StringValue `hcl:"source_range,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *appenginefirewallrule.Timeouts `hcl:"timeouts,block"`
}
type appEngineFirewallRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(aefr.ref.Append("action"))
}

// Description returns a reference to field description of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aefr.ref.Append("description"))
}

// Id returns a reference to field id of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aefr.ref.Append("id"))
}

// Priority returns a reference to field priority of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(aefr.ref.Append("priority"))
}

// Project returns a reference to field project of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aefr.ref.Append("project"))
}

// SourceRange returns a reference to field source_range of google_app_engine_firewall_rule.
func (aefr appEngineFirewallRuleAttributes) SourceRange() terra.StringValue {
	return terra.ReferenceAsString(aefr.ref.Append("source_range"))
}

func (aefr appEngineFirewallRuleAttributes) Timeouts() appenginefirewallrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appenginefirewallrule.TimeoutsAttributes](aefr.ref.Append("timeouts"))
}

type appEngineFirewallRuleState struct {
	Action      string                               `json:"action"`
	Description string                               `json:"description"`
	Id          string                               `json:"id"`
	Priority    float64                              `json:"priority"`
	Project     string                               `json:"project"`
	SourceRange string                               `json:"source_range"`
	Timeouts    *appenginefirewallrule.TimeoutsState `json:"timeouts"`
}
