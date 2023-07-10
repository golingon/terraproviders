// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorruleset "github.com/golingon/terraproviders/azurerm/3.64.0/cdnfrontdoorruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorRuleSet creates a new instance of [CdnFrontdoorRuleSet].
func NewCdnFrontdoorRuleSet(name string, args CdnFrontdoorRuleSetArgs) *CdnFrontdoorRuleSet {
	return &CdnFrontdoorRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorRuleSet)(nil)

// CdnFrontdoorRuleSet represents the Terraform resource azurerm_cdn_frontdoor_rule_set.
type CdnFrontdoorRuleSet struct {
	Name      string
	Args      CdnFrontdoorRuleSetArgs
	state     *cdnFrontdoorRuleSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) Type() string {
	return "azurerm_cdn_frontdoor_rule_set"
}

// LocalName returns the local name for [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) LocalName() string {
	return cfrs.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) Configuration() interface{} {
	return cfrs.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) DependOn() terra.Reference {
	return terra.ReferenceResource(cfrs)
}

// Dependencies returns the list of resources [CdnFrontdoorRuleSet] depends_on.
func (cfrs *CdnFrontdoorRuleSet) Dependencies() terra.Dependencies {
	return cfrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) LifecycleManagement() *terra.Lifecycle {
	return cfrs.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorRuleSet].
func (cfrs *CdnFrontdoorRuleSet) Attributes() cdnFrontdoorRuleSetAttributes {
	return cdnFrontdoorRuleSetAttributes{ref: terra.ReferenceResource(cfrs)}
}

// ImportState imports the given attribute values into [CdnFrontdoorRuleSet]'s state.
func (cfrs *CdnFrontdoorRuleSet) ImportState(av io.Reader) error {
	cfrs.state = &cdnFrontdoorRuleSetState{}
	if err := json.NewDecoder(av).Decode(cfrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfrs.Type(), cfrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorRuleSet] has state.
func (cfrs *CdnFrontdoorRuleSet) State() (*cdnFrontdoorRuleSetState, bool) {
	return cfrs.state, cfrs.state != nil
}

// StateMust returns the state for [CdnFrontdoorRuleSet]. Panics if the state is nil.
func (cfrs *CdnFrontdoorRuleSet) StateMust() *cdnFrontdoorRuleSetState {
	if cfrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfrs.Type(), cfrs.LocalName()))
	}
	return cfrs.state
}

// CdnFrontdoorRuleSetArgs contains the configurations for azurerm_cdn_frontdoor_rule_set.
type CdnFrontdoorRuleSetArgs struct {
	// CdnFrontdoorProfileId: string, required
	CdnFrontdoorProfileId terra.StringValue `hcl:"cdn_frontdoor_profile_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorruleset.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorRuleSetAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorProfileId returns a reference to field cdn_frontdoor_profile_id of azurerm_cdn_frontdoor_rule_set.
func (cfrs cdnFrontdoorRuleSetAttributes) CdnFrontdoorProfileId() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("cdn_frontdoor_profile_id"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_rule_set.
func (cfrs cdnFrontdoorRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_rule_set.
func (cfrs cdnFrontdoorRuleSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfrs.ref.Append("name"))
}

func (cfrs cdnFrontdoorRuleSetAttributes) Timeouts() cdnfrontdoorruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorruleset.TimeoutsAttributes](cfrs.ref.Append("timeouts"))
}

type cdnFrontdoorRuleSetState struct {
	CdnFrontdoorProfileId string                             `json:"cdn_frontdoor_profile_id"`
	Id                    string                             `json:"id"`
	Name                  string                             `json:"name"`
	Timeouts              *cdnfrontdoorruleset.TimeoutsState `json:"timeouts"`
}
