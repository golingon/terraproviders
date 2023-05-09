// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	frontdoorrulesengine "github.com/golingon/terraproviders/azurerm/3.55.0/frontdoorrulesengine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFrontdoorRulesEngine creates a new instance of [FrontdoorRulesEngine].
func NewFrontdoorRulesEngine(name string, args FrontdoorRulesEngineArgs) *FrontdoorRulesEngine {
	return &FrontdoorRulesEngine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FrontdoorRulesEngine)(nil)

// FrontdoorRulesEngine represents the Terraform resource azurerm_frontdoor_rules_engine.
type FrontdoorRulesEngine struct {
	Name      string
	Args      FrontdoorRulesEngineArgs
	state     *frontdoorRulesEngineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) Type() string {
	return "azurerm_frontdoor_rules_engine"
}

// LocalName returns the local name for [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) LocalName() string {
	return fre.Name
}

// Configuration returns the configuration (args) for [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) Configuration() interface{} {
	return fre.Args
}

// DependOn is used for other resources to depend on [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) DependOn() terra.Reference {
	return terra.ReferenceResource(fre)
}

// Dependencies returns the list of resources [FrontdoorRulesEngine] depends_on.
func (fre *FrontdoorRulesEngine) Dependencies() terra.Dependencies {
	return fre.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) LifecycleManagement() *terra.Lifecycle {
	return fre.Lifecycle
}

// Attributes returns the attributes for [FrontdoorRulesEngine].
func (fre *FrontdoorRulesEngine) Attributes() frontdoorRulesEngineAttributes {
	return frontdoorRulesEngineAttributes{ref: terra.ReferenceResource(fre)}
}

// ImportState imports the given attribute values into [FrontdoorRulesEngine]'s state.
func (fre *FrontdoorRulesEngine) ImportState(av io.Reader) error {
	fre.state = &frontdoorRulesEngineState{}
	if err := json.NewDecoder(av).Decode(fre.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fre.Type(), fre.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FrontdoorRulesEngine] has state.
func (fre *FrontdoorRulesEngine) State() (*frontdoorRulesEngineState, bool) {
	return fre.state, fre.state != nil
}

// StateMust returns the state for [FrontdoorRulesEngine]. Panics if the state is nil.
func (fre *FrontdoorRulesEngine) StateMust() *frontdoorRulesEngineState {
	if fre.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fre.Type(), fre.LocalName()))
	}
	return fre.state
}

// FrontdoorRulesEngineArgs contains the configurations for azurerm_frontdoor_rules_engine.
type FrontdoorRulesEngineArgs struct {
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
	Rule []frontdoorrulesengine.Rule `hcl:"rule,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *frontdoorrulesengine.Timeouts `hcl:"timeouts,block"`
}
type frontdoorRulesEngineAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(fre.ref.Append("enabled"))
}

// FrontdoorName returns a reference to field frontdoor_name of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) FrontdoorName() terra.StringValue {
	return terra.ReferenceAsString(fre.ref.Append("frontdoor_name"))
}

// Id returns a reference to field id of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fre.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fre.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fre.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_frontdoor_rules_engine.
func (fre frontdoorRulesEngineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fre.ref.Append("resource_group_name"))
}

func (fre frontdoorRulesEngineAttributes) Rule() terra.ListValue[frontdoorrulesengine.RuleAttributes] {
	return terra.ReferenceAsList[frontdoorrulesengine.RuleAttributes](fre.ref.Append("rule"))
}

func (fre frontdoorRulesEngineAttributes) Timeouts() frontdoorrulesengine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[frontdoorrulesengine.TimeoutsAttributes](fre.ref.Append("timeouts"))
}

type frontdoorRulesEngineState struct {
	Enabled           bool                                `json:"enabled"`
	FrontdoorName     string                              `json:"frontdoor_name"`
	Id                string                              `json:"id"`
	Location          string                              `json:"location"`
	Name              string                              `json:"name"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Rule              []frontdoorrulesengine.RuleState    `json:"rule"`
	Timeouts          *frontdoorrulesengine.TimeoutsState `json:"timeouts"`
}
