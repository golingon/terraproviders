// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	paloaltolocalrulestackrule "github.com/golingon/terraproviders/azurerm/3.98.0/paloaltolocalrulestackrule"
	"io"
)

// NewPaloAltoLocalRulestackRule creates a new instance of [PaloAltoLocalRulestackRule].
func NewPaloAltoLocalRulestackRule(name string, args PaloAltoLocalRulestackRuleArgs) *PaloAltoLocalRulestackRule {
	return &PaloAltoLocalRulestackRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PaloAltoLocalRulestackRule)(nil)

// PaloAltoLocalRulestackRule represents the Terraform resource azurerm_palo_alto_local_rulestack_rule.
type PaloAltoLocalRulestackRule struct {
	Name      string
	Args      PaloAltoLocalRulestackRuleArgs
	state     *paloAltoLocalRulestackRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) Type() string {
	return "azurerm_palo_alto_local_rulestack_rule"
}

// LocalName returns the local name for [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) LocalName() string {
	return palrr.Name
}

// Configuration returns the configuration (args) for [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) Configuration() interface{} {
	return palrr.Args
}

// DependOn is used for other resources to depend on [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) DependOn() terra.Reference {
	return terra.ReferenceResource(palrr)
}

// Dependencies returns the list of resources [PaloAltoLocalRulestackRule] depends_on.
func (palrr *PaloAltoLocalRulestackRule) Dependencies() terra.Dependencies {
	return palrr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) LifecycleManagement() *terra.Lifecycle {
	return palrr.Lifecycle
}

// Attributes returns the attributes for [PaloAltoLocalRulestackRule].
func (palrr *PaloAltoLocalRulestackRule) Attributes() paloAltoLocalRulestackRuleAttributes {
	return paloAltoLocalRulestackRuleAttributes{ref: terra.ReferenceResource(palrr)}
}

// ImportState imports the given attribute values into [PaloAltoLocalRulestackRule]'s state.
func (palrr *PaloAltoLocalRulestackRule) ImportState(av io.Reader) error {
	palrr.state = &paloAltoLocalRulestackRuleState{}
	if err := json.NewDecoder(av).Decode(palrr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", palrr.Type(), palrr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PaloAltoLocalRulestackRule] has state.
func (palrr *PaloAltoLocalRulestackRule) State() (*paloAltoLocalRulestackRuleState, bool) {
	return palrr.state, palrr.state != nil
}

// StateMust returns the state for [PaloAltoLocalRulestackRule]. Panics if the state is nil.
func (palrr *PaloAltoLocalRulestackRule) StateMust() *paloAltoLocalRulestackRuleState {
	if palrr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", palrr.Type(), palrr.LocalName()))
	}
	return palrr.state
}

// PaloAltoLocalRulestackRuleArgs contains the configurations for azurerm_palo_alto_local_rulestack_rule.
type PaloAltoLocalRulestackRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Applications: list of string, required
	Applications terra.ListValue[terra.StringValue] `hcl:"applications,attr" validate:"required"`
	// AuditComment: string, optional
	AuditComment terra.StringValue `hcl:"audit_comment,attr"`
	// DecryptionRuleType: string, optional
	DecryptionRuleType terra.StringValue `hcl:"decryption_rule_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InspectionCertificateId: string, optional
	InspectionCertificateId terra.StringValue `hcl:"inspection_certificate_id,attr"`
	// LoggingEnabled: bool, optional
	LoggingEnabled terra.BoolValue `hcl:"logging_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NegateDestination: bool, optional
	NegateDestination terra.BoolValue `hcl:"negate_destination,attr"`
	// NegateSource: bool, optional
	NegateSource terra.BoolValue `hcl:"negate_source,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// ProtocolPorts: list of string, optional
	ProtocolPorts terra.ListValue[terra.StringValue] `hcl:"protocol_ports,attr"`
	// RulestackId: string, required
	RulestackId terra.StringValue `hcl:"rulestack_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Category: optional
	Category *paloaltolocalrulestackrule.Category `hcl:"category,block"`
	// Destination: required
	Destination *paloaltolocalrulestackrule.Destination `hcl:"destination,block" validate:"required"`
	// Source: required
	Source *paloaltolocalrulestackrule.Source `hcl:"source,block" validate:"required"`
	// Timeouts: optional
	Timeouts *paloaltolocalrulestackrule.Timeouts `hcl:"timeouts,block"`
}
type paloAltoLocalRulestackRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("action"))
}

// Applications returns a reference to field applications of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Applications() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](palrr.ref.Append("applications"))
}

// AuditComment returns a reference to field audit_comment of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) AuditComment() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("audit_comment"))
}

// DecryptionRuleType returns a reference to field decryption_rule_type of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) DecryptionRuleType() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("decryption_rule_type"))
}

// Description returns a reference to field description of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(palrr.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("id"))
}

// InspectionCertificateId returns a reference to field inspection_certificate_id of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) InspectionCertificateId() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("inspection_certificate_id"))
}

// LoggingEnabled returns a reference to field logging_enabled of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) LoggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(palrr.ref.Append("logging_enabled"))
}

// Name returns a reference to field name of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("name"))
}

// NegateDestination returns a reference to field negate_destination of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) NegateDestination() terra.BoolValue {
	return terra.ReferenceAsBool(palrr.ref.Append("negate_destination"))
}

// NegateSource returns a reference to field negate_source of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) NegateSource() terra.BoolValue {
	return terra.ReferenceAsBool(palrr.ref.Append("negate_source"))
}

// Priority returns a reference to field priority of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(palrr.ref.Append("priority"))
}

// Protocol returns a reference to field protocol of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("protocol"))
}

// ProtocolPorts returns a reference to field protocol_ports of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) ProtocolPorts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](palrr.ref.Append("protocol_ports"))
}

// RulestackId returns a reference to field rulestack_id of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) RulestackId() terra.StringValue {
	return terra.ReferenceAsString(palrr.ref.Append("rulestack_id"))
}

// Tags returns a reference to field tags of azurerm_palo_alto_local_rulestack_rule.
func (palrr paloAltoLocalRulestackRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](palrr.ref.Append("tags"))
}

func (palrr paloAltoLocalRulestackRuleAttributes) Category() terra.ListValue[paloaltolocalrulestackrule.CategoryAttributes] {
	return terra.ReferenceAsList[paloaltolocalrulestackrule.CategoryAttributes](palrr.ref.Append("category"))
}

func (palrr paloAltoLocalRulestackRuleAttributes) Destination() terra.ListValue[paloaltolocalrulestackrule.DestinationAttributes] {
	return terra.ReferenceAsList[paloaltolocalrulestackrule.DestinationAttributes](palrr.ref.Append("destination"))
}

func (palrr paloAltoLocalRulestackRuleAttributes) Source() terra.ListValue[paloaltolocalrulestackrule.SourceAttributes] {
	return terra.ReferenceAsList[paloaltolocalrulestackrule.SourceAttributes](palrr.ref.Append("source"))
}

func (palrr paloAltoLocalRulestackRuleAttributes) Timeouts() paloaltolocalrulestackrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[paloaltolocalrulestackrule.TimeoutsAttributes](palrr.ref.Append("timeouts"))
}

type paloAltoLocalRulestackRuleState struct {
	Action                  string                                        `json:"action"`
	Applications            []string                                      `json:"applications"`
	AuditComment            string                                        `json:"audit_comment"`
	DecryptionRuleType      string                                        `json:"decryption_rule_type"`
	Description             string                                        `json:"description"`
	Enabled                 bool                                          `json:"enabled"`
	Id                      string                                        `json:"id"`
	InspectionCertificateId string                                        `json:"inspection_certificate_id"`
	LoggingEnabled          bool                                          `json:"logging_enabled"`
	Name                    string                                        `json:"name"`
	NegateDestination       bool                                          `json:"negate_destination"`
	NegateSource            bool                                          `json:"negate_source"`
	Priority                float64                                       `json:"priority"`
	Protocol                string                                        `json:"protocol"`
	ProtocolPorts           []string                                      `json:"protocol_ports"`
	RulestackId             string                                        `json:"rulestack_id"`
	Tags                    map[string]string                             `json:"tags"`
	Category                []paloaltolocalrulestackrule.CategoryState    `json:"category"`
	Destination             []paloaltolocalrulestackrule.DestinationState `json:"destination"`
	Source                  []paloaltolocalrulestackrule.SourceState      `json:"source"`
	Timeouts                *paloaltolocalrulestackrule.TimeoutsState     `json:"timeouts"`
}
