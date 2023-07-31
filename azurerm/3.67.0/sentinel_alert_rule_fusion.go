// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertrulefusion "github.com/golingon/terraproviders/azurerm/3.67.0/sentinelalertrulefusion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleFusion creates a new instance of [SentinelAlertRuleFusion].
func NewSentinelAlertRuleFusion(name string, args SentinelAlertRuleFusionArgs) *SentinelAlertRuleFusion {
	return &SentinelAlertRuleFusion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleFusion)(nil)

// SentinelAlertRuleFusion represents the Terraform resource azurerm_sentinel_alert_rule_fusion.
type SentinelAlertRuleFusion struct {
	Name      string
	Args      SentinelAlertRuleFusionArgs
	state     *sentinelAlertRuleFusionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) Type() string {
	return "azurerm_sentinel_alert_rule_fusion"
}

// LocalName returns the local name for [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) LocalName() string {
	return sarf.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) Configuration() interface{} {
	return sarf.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) DependOn() terra.Reference {
	return terra.ReferenceResource(sarf)
}

// Dependencies returns the list of resources [SentinelAlertRuleFusion] depends_on.
func (sarf *SentinelAlertRuleFusion) Dependencies() terra.Dependencies {
	return sarf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) LifecycleManagement() *terra.Lifecycle {
	return sarf.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleFusion].
func (sarf *SentinelAlertRuleFusion) Attributes() sentinelAlertRuleFusionAttributes {
	return sentinelAlertRuleFusionAttributes{ref: terra.ReferenceResource(sarf)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleFusion]'s state.
func (sarf *SentinelAlertRuleFusion) ImportState(av io.Reader) error {
	sarf.state = &sentinelAlertRuleFusionState{}
	if err := json.NewDecoder(av).Decode(sarf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarf.Type(), sarf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleFusion] has state.
func (sarf *SentinelAlertRuleFusion) State() (*sentinelAlertRuleFusionState, bool) {
	return sarf.state, sarf.state != nil
}

// StateMust returns the state for [SentinelAlertRuleFusion]. Panics if the state is nil.
func (sarf *SentinelAlertRuleFusion) StateMust() *sentinelAlertRuleFusionState {
	if sarf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarf.Type(), sarf.LocalName()))
	}
	return sarf.state
}

// SentinelAlertRuleFusionArgs contains the configurations for azurerm_sentinel_alert_rule_fusion.
type SentinelAlertRuleFusionArgs struct {
	// AlertRuleTemplateGuid: string, required
	AlertRuleTemplateGuid terra.StringValue `hcl:"alert_rule_template_guid,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Source: min=0
	Source []sentinelalertrulefusion.Source `hcl:"source,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *sentinelalertrulefusion.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleFusionAttributes struct {
	ref terra.Reference
}

// AlertRuleTemplateGuid returns a reference to field alert_rule_template_guid of azurerm_sentinel_alert_rule_fusion.
func (sarf sentinelAlertRuleFusionAttributes) AlertRuleTemplateGuid() terra.StringValue {
	return terra.ReferenceAsString(sarf.ref.Append("alert_rule_template_guid"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_fusion.
func (sarf sentinelAlertRuleFusionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarf.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_fusion.
func (sarf sentinelAlertRuleFusionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarf.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_fusion.
func (sarf sentinelAlertRuleFusionAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarf.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_fusion.
func (sarf sentinelAlertRuleFusionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarf.ref.Append("name"))
}

func (sarf sentinelAlertRuleFusionAttributes) Source() terra.ListValue[sentinelalertrulefusion.SourceAttributes] {
	return terra.ReferenceAsList[sentinelalertrulefusion.SourceAttributes](sarf.ref.Append("source"))
}

func (sarf sentinelAlertRuleFusionAttributes) Timeouts() sentinelalertrulefusion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertrulefusion.TimeoutsAttributes](sarf.ref.Append("timeouts"))
}

type sentinelAlertRuleFusionState struct {
	AlertRuleTemplateGuid   string                                 `json:"alert_rule_template_guid"`
	Enabled                 bool                                   `json:"enabled"`
	Id                      string                                 `json:"id"`
	LogAnalyticsWorkspaceId string                                 `json:"log_analytics_workspace_id"`
	Name                    string                                 `json:"name"`
	Source                  []sentinelalertrulefusion.SourceState  `json:"source"`
	Timeouts                *sentinelalertrulefusion.TimeoutsState `json:"timeouts"`
}
