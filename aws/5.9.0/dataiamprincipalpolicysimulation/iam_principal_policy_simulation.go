// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataiamprincipalpolicysimulation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Results struct {
	// MatchedStatements: min=0
	MatchedStatements []MatchedStatements `hcl:"matched_statements,block" validate:"min=0"`
}

type MatchedStatements struct{}

type Context struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ResultsAttributes struct {
	ref terra.Reference
}

func (r ResultsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResultsAttributes) InternalWithRef(ref terra.Reference) ResultsAttributes {
	return ResultsAttributes{ref: ref}
}

func (r ResultsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResultsAttributes) ActionName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("action_name"))
}

func (r ResultsAttributes) Allowed() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("allowed"))
}

func (r ResultsAttributes) Decision() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("decision"))
}

func (r ResultsAttributes) DecisionDetails() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("decision_details"))
}

func (r ResultsAttributes) MissingContextKeys() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](r.ref.Append("missing_context_keys"))
}

func (r ResultsAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("resource_arn"))
}

func (r ResultsAttributes) MatchedStatements() terra.SetValue[MatchedStatementsAttributes] {
	return terra.ReferenceAsSet[MatchedStatementsAttributes](r.ref.Append("matched_statements"))
}

type MatchedStatementsAttributes struct {
	ref terra.Reference
}

func (ms MatchedStatementsAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MatchedStatementsAttributes) InternalWithRef(ref terra.Reference) MatchedStatementsAttributes {
	return MatchedStatementsAttributes{ref: ref}
}

func (ms MatchedStatementsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MatchedStatementsAttributes) SourcePolicyId() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("source_policy_id"))
}

func (ms MatchedStatementsAttributes) SourcePolicyType() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("source_policy_type"))
}

type ContextAttributes struct {
	ref terra.Reference
}

func (c ContextAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContextAttributes) InternalWithRef(ref terra.Reference) ContextAttributes {
	return ContextAttributes{ref: ref}
}

func (c ContextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContextAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("key"))
}

func (c ContextAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

func (c ContextAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("values"))
}

type ResultsState struct {
	ActionName         string                   `json:"action_name"`
	Allowed            bool                     `json:"allowed"`
	Decision           string                   `json:"decision"`
	DecisionDetails    map[string]string        `json:"decision_details"`
	MissingContextKeys []string                 `json:"missing_context_keys"`
	ResourceArn        string                   `json:"resource_arn"`
	MatchedStatements  []MatchedStatementsState `json:"matched_statements"`
}

type MatchedStatementsState struct {
	SourcePolicyId   string `json:"source_policy_id"`
	SourcePolicyType string `json:"source_policy_type"`
}

type ContextState struct {
	Key    string   `json:"key"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
}