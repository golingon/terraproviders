// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_rule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_cdn_frontdoor_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermCdnFrontdoorRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acfr *Resource) Type() string {
	return "azurerm_cdn_frontdoor_rule"
}

// LocalName returns the local name for [Resource].
func (acfr *Resource) LocalName() string {
	return acfr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acfr *Resource) Configuration() interface{} {
	return acfr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acfr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acfr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acfr *Resource) Dependencies() terra.Dependencies {
	return acfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acfr *Resource) LifecycleManagement() *terra.Lifecycle {
	return acfr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acfr *Resource) Attributes() azurermCdnFrontdoorRuleAttributes {
	return azurermCdnFrontdoorRuleAttributes{ref: terra.ReferenceResource(acfr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acfr *Resource) ImportState(state io.Reader) error {
	acfr.state = &azurermCdnFrontdoorRuleState{}
	if err := json.NewDecoder(state).Decode(acfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acfr.Type(), acfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acfr *Resource) State() (*azurermCdnFrontdoorRuleState, bool) {
	return acfr.state, acfr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acfr *Resource) StateMust() *azurermCdnFrontdoorRuleState {
	if acfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acfr.Type(), acfr.LocalName()))
	}
	return acfr.state
}

// Args contains the configurations for azurerm_cdn_frontdoor_rule.
type Args struct {
	// BehaviorOnMatch: string, optional
	BehaviorOnMatch terra.StringValue `hcl:"behavior_on_match,attr"`
	// CdnFrontdoorRuleSetId: string, required
	CdnFrontdoorRuleSetId terra.StringValue `hcl:"cdn_frontdoor_rule_set_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Order: number, required
	Order terra.NumberValue `hcl:"order,attr" validate:"required"`
	// Actions: required
	Actions *Actions `hcl:"actions,block" validate:"required"`
	// Conditions: optional
	Conditions *Conditions `hcl:"conditions,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermCdnFrontdoorRuleAttributes struct {
	ref terra.Reference
}

// BehaviorOnMatch returns a reference to field behavior_on_match of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) BehaviorOnMatch() terra.StringValue {
	return terra.ReferenceAsString(acfr.ref.Append("behavior_on_match"))
}

// CdnFrontdoorRuleSetId returns a reference to field cdn_frontdoor_rule_set_id of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) CdnFrontdoorRuleSetId() terra.StringValue {
	return terra.ReferenceAsString(acfr.ref.Append("cdn_frontdoor_rule_set_id"))
}

// CdnFrontdoorRuleSetName returns a reference to field cdn_frontdoor_rule_set_name of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) CdnFrontdoorRuleSetName() terra.StringValue {
	return terra.ReferenceAsString(acfr.ref.Append("cdn_frontdoor_rule_set_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acfr.ref.Append("name"))
}

// Order returns a reference to field order of azurerm_cdn_frontdoor_rule.
func (acfr azurermCdnFrontdoorRuleAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(acfr.ref.Append("order"))
}

func (acfr azurermCdnFrontdoorRuleAttributes) Actions() terra.ListValue[ActionsAttributes] {
	return terra.ReferenceAsList[ActionsAttributes](acfr.ref.Append("actions"))
}

func (acfr azurermCdnFrontdoorRuleAttributes) Conditions() terra.ListValue[ConditionsAttributes] {
	return terra.ReferenceAsList[ConditionsAttributes](acfr.ref.Append("conditions"))
}

func (acfr azurermCdnFrontdoorRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acfr.ref.Append("timeouts"))
}

type azurermCdnFrontdoorRuleState struct {
	BehaviorOnMatch         string            `json:"behavior_on_match"`
	CdnFrontdoorRuleSetId   string            `json:"cdn_frontdoor_rule_set_id"`
	CdnFrontdoorRuleSetName string            `json:"cdn_frontdoor_rule_set_name"`
	Id                      string            `json:"id"`
	Name                    string            `json:"name"`
	Order                   float64           `json:"order"`
	Actions                 []ActionsState    `json:"actions"`
	Conditions              []ConditionsState `json:"conditions"`
	Timeouts                *TimeoutsState    `json:"timeouts"`
}
