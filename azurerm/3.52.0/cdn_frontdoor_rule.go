// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorrule "github.com/golingon/terraproviders/azurerm/3.52.0/cdnfrontdoorrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorRule creates a new instance of [CdnFrontdoorRule].
func NewCdnFrontdoorRule(name string, args CdnFrontdoorRuleArgs) *CdnFrontdoorRule {
	return &CdnFrontdoorRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorRule)(nil)

// CdnFrontdoorRule represents the Terraform resource azurerm_cdn_frontdoor_rule.
type CdnFrontdoorRule struct {
	Name      string
	Args      CdnFrontdoorRuleArgs
	state     *cdnFrontdoorRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) Type() string {
	return "azurerm_cdn_frontdoor_rule"
}

// LocalName returns the local name for [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) LocalName() string {
	return cfr.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) Configuration() interface{} {
	return cfr.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cfr)
}

// Dependencies returns the list of resources [CdnFrontdoorRule] depends_on.
func (cfr *CdnFrontdoorRule) Dependencies() terra.Dependencies {
	return cfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) LifecycleManagement() *terra.Lifecycle {
	return cfr.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorRule].
func (cfr *CdnFrontdoorRule) Attributes() cdnFrontdoorRuleAttributes {
	return cdnFrontdoorRuleAttributes{ref: terra.ReferenceResource(cfr)}
}

// ImportState imports the given attribute values into [CdnFrontdoorRule]'s state.
func (cfr *CdnFrontdoorRule) ImportState(av io.Reader) error {
	cfr.state = &cdnFrontdoorRuleState{}
	if err := json.NewDecoder(av).Decode(cfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfr.Type(), cfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorRule] has state.
func (cfr *CdnFrontdoorRule) State() (*cdnFrontdoorRuleState, bool) {
	return cfr.state, cfr.state != nil
}

// StateMust returns the state for [CdnFrontdoorRule]. Panics if the state is nil.
func (cfr *CdnFrontdoorRule) StateMust() *cdnFrontdoorRuleState {
	if cfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfr.Type(), cfr.LocalName()))
	}
	return cfr.state
}

// CdnFrontdoorRuleArgs contains the configurations for azurerm_cdn_frontdoor_rule.
type CdnFrontdoorRuleArgs struct {
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
	Actions *cdnfrontdoorrule.Actions `hcl:"actions,block" validate:"required"`
	// Conditions: optional
	Conditions *cdnfrontdoorrule.Conditions `hcl:"conditions,block"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorrule.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorRuleAttributes struct {
	ref terra.Reference
}

// BehaviorOnMatch returns a reference to field behavior_on_match of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) BehaviorOnMatch() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("behavior_on_match"))
}

// CdnFrontdoorRuleSetId returns a reference to field cdn_frontdoor_rule_set_id of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) CdnFrontdoorRuleSetId() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("cdn_frontdoor_rule_set_id"))
}

// CdnFrontdoorRuleSetName returns a reference to field cdn_frontdoor_rule_set_name of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) CdnFrontdoorRuleSetName() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("cdn_frontdoor_rule_set_name"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("name"))
}

// Order returns a reference to field order of azurerm_cdn_frontdoor_rule.
func (cfr cdnFrontdoorRuleAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(cfr.ref.Append("order"))
}

func (cfr cdnFrontdoorRuleAttributes) Actions() terra.ListValue[cdnfrontdoorrule.ActionsAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorrule.ActionsAttributes](cfr.ref.Append("actions"))
}

func (cfr cdnFrontdoorRuleAttributes) Conditions() terra.ListValue[cdnfrontdoorrule.ConditionsAttributes] {
	return terra.ReferenceAsList[cdnfrontdoorrule.ConditionsAttributes](cfr.ref.Append("conditions"))
}

func (cfr cdnFrontdoorRuleAttributes) Timeouts() cdnfrontdoorrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorrule.TimeoutsAttributes](cfr.ref.Append("timeouts"))
}

type cdnFrontdoorRuleState struct {
	BehaviorOnMatch         string                             `json:"behavior_on_match"`
	CdnFrontdoorRuleSetId   string                             `json:"cdn_frontdoor_rule_set_id"`
	CdnFrontdoorRuleSetName string                             `json:"cdn_frontdoor_rule_set_name"`
	Id                      string                             `json:"id"`
	Name                    string                             `json:"name"`
	Order                   float64                            `json:"order"`
	Actions                 []cdnfrontdoorrule.ActionsState    `json:"actions"`
	Conditions              []cdnfrontdoorrule.ConditionsState `json:"conditions"`
	Timeouts                *cdnfrontdoorrule.TimeoutsState    `json:"timeouts"`
}
