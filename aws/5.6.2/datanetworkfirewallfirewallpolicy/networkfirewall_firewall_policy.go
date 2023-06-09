// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanetworkfirewallfirewallpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FirewallPolicy struct {
	// StatefulEngineOptions: min=0
	StatefulEngineOptions []StatefulEngineOptions `hcl:"stateful_engine_options,block" validate:"min=0"`
	// StatefulRuleGroupReference: min=0
	StatefulRuleGroupReference []StatefulRuleGroupReference `hcl:"stateful_rule_group_reference,block" validate:"min=0"`
	// StatelessCustomAction: min=0
	StatelessCustomAction []StatelessCustomAction `hcl:"stateless_custom_action,block" validate:"min=0"`
	// StatelessRuleGroupReference: min=0
	StatelessRuleGroupReference []StatelessRuleGroupReference `hcl:"stateless_rule_group_reference,block" validate:"min=0"`
}

type StatefulEngineOptions struct{}

type StatefulRuleGroupReference struct {
	// Override: min=0
	Override []Override `hcl:"override,block" validate:"min=0"`
}

type Override struct{}

type StatelessCustomAction struct {
	// ActionDefinition: min=0
	ActionDefinition []ActionDefinition `hcl:"action_definition,block" validate:"min=0"`
}

type ActionDefinition struct {
	// PublishMetricAction: min=0
	PublishMetricAction []PublishMetricAction `hcl:"publish_metric_action,block" validate:"min=0"`
}

type PublishMetricAction struct {
	// Dimension: min=0
	Dimension []Dimension `hcl:"dimension,block" validate:"min=0"`
}

type Dimension struct{}

type StatelessRuleGroupReference struct{}

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
