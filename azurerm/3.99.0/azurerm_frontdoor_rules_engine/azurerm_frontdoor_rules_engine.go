// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_frontdoor_rules_engine

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

// Resource represents the Terraform resource azurerm_frontdoor_rules_engine.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermFrontdoorRulesEngineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afre *Resource) Type() string {
	return "azurerm_frontdoor_rules_engine"
}

// LocalName returns the local name for [Resource].
func (afre *Resource) LocalName() string {
	return afre.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afre *Resource) Configuration() interface{} {
	return afre.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afre *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afre)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afre *Resource) Dependencies() terra.Dependencies {
	return afre.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afre *Resource) LifecycleManagement() *terra.Lifecycle {
	return afre.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afre *Resource) Attributes() azurermFrontdoorRulesEngineAttributes {
	return azurermFrontdoorRulesEngineAttributes{ref: terra.ReferenceResource(afre)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afre *Resource) ImportState(state io.Reader) error {
	afre.state = &azurermFrontdoorRulesEngineState{}
	if err := json.NewDecoder(state).Decode(afre.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afre.Type(), afre.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afre *Resource) State() (*azurermFrontdoorRulesEngineState, bool) {
	return afre.state, afre.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afre *Resource) StateMust() *azurermFrontdoorRulesEngineState {
	if afre.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afre.Type(), afre.LocalName()))
	}
	return afre.state
}

// Args contains the configurations for azurerm_frontdoor_rules_engine.
type Args struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FrontdoorName: string, required
	FrontdoorName terra.StringValue `hcl:"frontdoor_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Rule: min=0,max=100
	Rule []Rule `hcl:"rule,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermFrontdoorRulesEngineAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(afre.ref.Append("enabled"))
}

// FrontdoorName returns a reference to field frontdoor_name of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) FrontdoorName() terra.StringValue {
	return terra.ReferenceAsString(afre.ref.Append("frontdoor_name"))
}

// Id returns a reference to field id of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afre.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(afre.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(afre.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_frontdoor_rules_engine.
func (afre azurermFrontdoorRulesEngineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(afre.ref.Append("resource_group_name"))
}

func (afre azurermFrontdoorRulesEngineAttributes) Rule() terra.ListValue[RuleAttributes] {
	return terra.ReferenceAsList[RuleAttributes](afre.ref.Append("rule"))
}

func (afre azurermFrontdoorRulesEngineAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](afre.ref.Append("timeouts"))
}

type azurermFrontdoorRulesEngineState struct {
	Enabled           bool           `json:"enabled"`
	FrontdoorName     string         `json:"frontdoor_name"`
	Id                string         `json:"id"`
	Location          string         `json:"location"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	Rule              []RuleState    `json:"rule"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
