// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkfirewallfirewallpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfiguration struct {
	// KeyId: string, optional
	KeyId terra.StringValue `hcl:"key_id,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type FirewallPolicy struct {
	// StatefulDefaultActions: set of string, optional
	StatefulDefaultActions terra.SetValue[terra.StringValue] `hcl:"stateful_default_actions,attr"`
	// StatelessDefaultActions: set of string, required
	StatelessDefaultActions terra.SetValue[terra.StringValue] `hcl:"stateless_default_actions,attr" validate:"required"`
	// StatelessFragmentDefaultActions: set of string, required
	StatelessFragmentDefaultActions terra.SetValue[terra.StringValue] `hcl:"stateless_fragment_default_actions,attr" validate:"required"`
	// StatefulEngineOptions: optional
	StatefulEngineOptions *StatefulEngineOptions `hcl:"stateful_engine_options,block"`
	// StatefulRuleGroupReference: min=0
	StatefulRuleGroupReference []StatefulRuleGroupReference `hcl:"stateful_rule_group_reference,block" validate:"min=0"`
	// StatelessCustomAction: min=0
	StatelessCustomAction []StatelessCustomAction `hcl:"stateless_custom_action,block" validate:"min=0"`
	// StatelessRuleGroupReference: min=0
	StatelessRuleGroupReference []StatelessRuleGroupReference `hcl:"stateless_rule_group_reference,block" validate:"min=0"`
}

type StatefulEngineOptions struct {
	// RuleOrder: string, optional
	RuleOrder terra.StringValue `hcl:"rule_order,attr"`
	// StreamExceptionPolicy: string, optional
	StreamExceptionPolicy terra.StringValue `hcl:"stream_exception_policy,attr"`
}

type StatefulRuleGroupReference struct {
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// Override: optional
	Override *Override `hcl:"override,block"`
}

type Override struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
}

type StatelessCustomAction struct {
	// ActionName: string, required
	ActionName terra.StringValue `hcl:"action_name,attr" validate:"required"`
	// ActionDefinition: required
	ActionDefinition *ActionDefinition `hcl:"action_definition,block" validate:"required"`
}

type ActionDefinition struct {
	// PublishMetricAction: required
	PublishMetricAction *PublishMetricAction `hcl:"publish_metric_action,block" validate:"required"`
}

type PublishMetricAction struct {
	// Dimension: min=1
	Dimension []Dimension `hcl:"dimension,block" validate:"min=1"`
}

type Dimension struct {
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type StatelessRuleGroupReference struct {
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("key_id"))
}

func (ec EncryptionConfigurationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("type"))
}

type FirewallPolicyAttributes struct {
	ref terra.Reference
}

func (fp FirewallPolicyAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FirewallPolicyAttributes) InternalWithRef(ref terra.Reference) FirewallPolicyAttributes {
	return FirewallPolicyAttributes{ref: ref}
}

func (fp FirewallPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FirewallPolicyAttributes) StatefulDefaultActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fp.ref.Append("stateful_default_actions"))
}

func (fp FirewallPolicyAttributes) StatelessDefaultActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fp.ref.Append("stateless_default_actions"))
}

func (fp FirewallPolicyAttributes) StatelessFragmentDefaultActions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fp.ref.Append("stateless_fragment_default_actions"))
}

func (fp FirewallPolicyAttributes) StatefulEngineOptions() terra.ListValue[StatefulEngineOptionsAttributes] {
	return terra.ReferenceAsList[StatefulEngineOptionsAttributes](fp.ref.Append("stateful_engine_options"))
}

func (fp FirewallPolicyAttributes) StatefulRuleGroupReference() terra.SetValue[StatefulRuleGroupReferenceAttributes] {
	return terra.ReferenceAsSet[StatefulRuleGroupReferenceAttributes](fp.ref.Append("stateful_rule_group_reference"))
}

func (fp FirewallPolicyAttributes) StatelessCustomAction() terra.SetValue[StatelessCustomActionAttributes] {
	return terra.ReferenceAsSet[StatelessCustomActionAttributes](fp.ref.Append("stateless_custom_action"))
}

func (fp FirewallPolicyAttributes) StatelessRuleGroupReference() terra.SetValue[StatelessRuleGroupReferenceAttributes] {
	return terra.ReferenceAsSet[StatelessRuleGroupReferenceAttributes](fp.ref.Append("stateless_rule_group_reference"))
}

type StatefulEngineOptionsAttributes struct {
	ref terra.Reference
}

func (seo StatefulEngineOptionsAttributes) InternalRef() (terra.Reference, error) {
	return seo.ref, nil
}

func (seo StatefulEngineOptionsAttributes) InternalWithRef(ref terra.Reference) StatefulEngineOptionsAttributes {
	return StatefulEngineOptionsAttributes{ref: ref}
}

func (seo StatefulEngineOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return seo.ref.InternalTokens()
}

func (seo StatefulEngineOptionsAttributes) RuleOrder() terra.StringValue {
	return terra.ReferenceAsString(seo.ref.Append("rule_order"))
}

func (seo StatefulEngineOptionsAttributes) StreamExceptionPolicy() terra.StringValue {
	return terra.ReferenceAsString(seo.ref.Append("stream_exception_policy"))
}

type StatefulRuleGroupReferenceAttributes struct {
	ref terra.Reference
}

func (srgr StatefulRuleGroupReferenceAttributes) InternalRef() (terra.Reference, error) {
	return srgr.ref, nil
}

func (srgr StatefulRuleGroupReferenceAttributes) InternalWithRef(ref terra.Reference) StatefulRuleGroupReferenceAttributes {
	return StatefulRuleGroupReferenceAttributes{ref: ref}
}

func (srgr StatefulRuleGroupReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return srgr.ref.InternalTokens()
}

func (srgr StatefulRuleGroupReferenceAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(srgr.ref.Append("priority"))
}

func (srgr StatefulRuleGroupReferenceAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(srgr.ref.Append("resource_arn"))
}

func (srgr StatefulRuleGroupReferenceAttributes) Override() terra.ListValue[OverrideAttributes] {
	return terra.ReferenceAsList[OverrideAttributes](srgr.ref.Append("override"))
}

type OverrideAttributes struct {
	ref terra.Reference
}

func (o OverrideAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OverrideAttributes) InternalWithRef(ref terra.Reference) OverrideAttributes {
	return OverrideAttributes{ref: ref}
}

func (o OverrideAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OverrideAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("action"))
}

type StatelessCustomActionAttributes struct {
	ref terra.Reference
}

func (sca StatelessCustomActionAttributes) InternalRef() (terra.Reference, error) {
	return sca.ref, nil
}

func (sca StatelessCustomActionAttributes) InternalWithRef(ref terra.Reference) StatelessCustomActionAttributes {
	return StatelessCustomActionAttributes{ref: ref}
}

func (sca StatelessCustomActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sca.ref.InternalTokens()
}

func (sca StatelessCustomActionAttributes) ActionName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("action_name"))
}

func (sca StatelessCustomActionAttributes) ActionDefinition() terra.ListValue[ActionDefinitionAttributes] {
	return terra.ReferenceAsList[ActionDefinitionAttributes](sca.ref.Append("action_definition"))
}

type ActionDefinitionAttributes struct {
	ref terra.Reference
}

func (ad ActionDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad ActionDefinitionAttributes) InternalWithRef(ref terra.Reference) ActionDefinitionAttributes {
	return ActionDefinitionAttributes{ref: ref}
}

func (ad ActionDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad ActionDefinitionAttributes) PublishMetricAction() terra.ListValue[PublishMetricActionAttributes] {
	return terra.ReferenceAsList[PublishMetricActionAttributes](ad.ref.Append("publish_metric_action"))
}

type PublishMetricActionAttributes struct {
	ref terra.Reference
}

func (pma PublishMetricActionAttributes) InternalRef() (terra.Reference, error) {
	return pma.ref, nil
}

func (pma PublishMetricActionAttributes) InternalWithRef(ref terra.Reference) PublishMetricActionAttributes {
	return PublishMetricActionAttributes{ref: ref}
}

func (pma PublishMetricActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pma.ref.InternalTokens()
}

func (pma PublishMetricActionAttributes) Dimension() terra.SetValue[DimensionAttributes] {
	return terra.ReferenceAsSet[DimensionAttributes](pma.ref.Append("dimension"))
}

type DimensionAttributes struct {
	ref terra.Reference
}

func (d DimensionAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DimensionAttributes) InternalWithRef(ref terra.Reference) DimensionAttributes {
	return DimensionAttributes{ref: ref}
}

func (d DimensionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DimensionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("value"))
}

type StatelessRuleGroupReferenceAttributes struct {
	ref terra.Reference
}

func (srgr StatelessRuleGroupReferenceAttributes) InternalRef() (terra.Reference, error) {
	return srgr.ref, nil
}

func (srgr StatelessRuleGroupReferenceAttributes) InternalWithRef(ref terra.Reference) StatelessRuleGroupReferenceAttributes {
	return StatelessRuleGroupReferenceAttributes{ref: ref}
}

func (srgr StatelessRuleGroupReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return srgr.ref.InternalTokens()
}

func (srgr StatelessRuleGroupReferenceAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(srgr.ref.Append("priority"))
}

func (srgr StatelessRuleGroupReferenceAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(srgr.ref.Append("resource_arn"))
}

type EncryptionConfigurationState struct {
	KeyId string `json:"key_id"`
	Type  string `json:"type"`
}

type FirewallPolicyState struct {
	StatefulDefaultActions          []string                           `json:"stateful_default_actions"`
	StatelessDefaultActions         []string                           `json:"stateless_default_actions"`
	StatelessFragmentDefaultActions []string                           `json:"stateless_fragment_default_actions"`
	StatefulEngineOptions           []StatefulEngineOptionsState       `json:"stateful_engine_options"`
	StatefulRuleGroupReference      []StatefulRuleGroupReferenceState  `json:"stateful_rule_group_reference"`
	StatelessCustomAction           []StatelessCustomActionState       `json:"stateless_custom_action"`
	StatelessRuleGroupReference     []StatelessRuleGroupReferenceState `json:"stateless_rule_group_reference"`
}

type StatefulEngineOptionsState struct {
	RuleOrder             string `json:"rule_order"`
	StreamExceptionPolicy string `json:"stream_exception_policy"`
}

type StatefulRuleGroupReferenceState struct {
	Priority    float64         `json:"priority"`
	ResourceArn string          `json:"resource_arn"`
	Override    []OverrideState `json:"override"`
}

type OverrideState struct {
	Action string `json:"action"`
}

type StatelessCustomActionState struct {
	ActionName       string                  `json:"action_name"`
	ActionDefinition []ActionDefinitionState `json:"action_definition"`
}

type ActionDefinitionState struct {
	PublishMetricAction []PublishMetricActionState `json:"publish_metric_action"`
}

type PublishMetricActionState struct {
	Dimension []DimensionState `json:"dimension"`
}

type DimensionState struct {
	Value string `json:"value"`
}

type StatelessRuleGroupReferenceState struct {
	Priority    float64 `json:"priority"`
	ResourceArn string  `json:"resource_arn"`
}
