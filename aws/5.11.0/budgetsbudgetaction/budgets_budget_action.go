// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package budgetsbudgetaction

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ActionThreshold struct {
	// ActionThresholdType: string, required
	ActionThresholdType terra.StringValue `hcl:"action_threshold_type,attr" validate:"required"`
	// ActionThresholdValue: number, required
	ActionThresholdValue terra.NumberValue `hcl:"action_threshold_value,attr" validate:"required"`
}

type Definition struct {
	// IamActionDefinition: optional
	IamActionDefinition *IamActionDefinition `hcl:"iam_action_definition,block"`
	// ScpActionDefinition: optional
	ScpActionDefinition *ScpActionDefinition `hcl:"scp_action_definition,block"`
	// SsmActionDefinition: optional
	SsmActionDefinition *SsmActionDefinition `hcl:"ssm_action_definition,block"`
}

type IamActionDefinition struct {
	// Groups: set of string, optional
	Groups terra.SetValue[terra.StringValue] `hcl:"groups,attr"`
	// PolicyArn: string, required
	PolicyArn terra.StringValue `hcl:"policy_arn,attr" validate:"required"`
	// Roles: set of string, optional
	Roles terra.SetValue[terra.StringValue] `hcl:"roles,attr"`
	// Users: set of string, optional
	Users terra.SetValue[terra.StringValue] `hcl:"users,attr"`
}

type ScpActionDefinition struct {
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// TargetIds: set of string, required
	TargetIds terra.SetValue[terra.StringValue] `hcl:"target_ids,attr" validate:"required"`
}

type SsmActionDefinition struct {
	// ActionSubType: string, required
	ActionSubType terra.StringValue `hcl:"action_sub_type,attr" validate:"required"`
	// InstanceIds: set of string, required
	InstanceIds terra.SetValue[terra.StringValue] `hcl:"instance_ids,attr" validate:"required"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type Subscriber struct {
	// Address: string, required
	Address terra.StringValue `hcl:"address,attr" validate:"required"`
	// SubscriptionType: string, required
	SubscriptionType terra.StringValue `hcl:"subscription_type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ActionThresholdAttributes struct {
	ref terra.Reference
}

func (at ActionThresholdAttributes) InternalRef() (terra.Reference, error) {
	return at.ref, nil
}

func (at ActionThresholdAttributes) InternalWithRef(ref terra.Reference) ActionThresholdAttributes {
	return ActionThresholdAttributes{ref: ref}
}

func (at ActionThresholdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return at.ref.InternalTokens()
}

func (at ActionThresholdAttributes) ActionThresholdType() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("action_threshold_type"))
}

func (at ActionThresholdAttributes) ActionThresholdValue() terra.NumberValue {
	return terra.ReferenceAsNumber(at.ref.Append("action_threshold_value"))
}

type DefinitionAttributes struct {
	ref terra.Reference
}

func (d DefinitionAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DefinitionAttributes) InternalWithRef(ref terra.Reference) DefinitionAttributes {
	return DefinitionAttributes{ref: ref}
}

func (d DefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DefinitionAttributes) IamActionDefinition() terra.ListValue[IamActionDefinitionAttributes] {
	return terra.ReferenceAsList[IamActionDefinitionAttributes](d.ref.Append("iam_action_definition"))
}

func (d DefinitionAttributes) ScpActionDefinition() terra.ListValue[ScpActionDefinitionAttributes] {
	return terra.ReferenceAsList[ScpActionDefinitionAttributes](d.ref.Append("scp_action_definition"))
}

func (d DefinitionAttributes) SsmActionDefinition() terra.ListValue[SsmActionDefinitionAttributes] {
	return terra.ReferenceAsList[SsmActionDefinitionAttributes](d.ref.Append("ssm_action_definition"))
}

type IamActionDefinitionAttributes struct {
	ref terra.Reference
}

func (iad IamActionDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return iad.ref, nil
}

func (iad IamActionDefinitionAttributes) InternalWithRef(ref terra.Reference) IamActionDefinitionAttributes {
	return IamActionDefinitionAttributes{ref: ref}
}

func (iad IamActionDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iad.ref.InternalTokens()
}

func (iad IamActionDefinitionAttributes) Groups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iad.ref.Append("groups"))
}

func (iad IamActionDefinitionAttributes) PolicyArn() terra.StringValue {
	return terra.ReferenceAsString(iad.ref.Append("policy_arn"))
}

func (iad IamActionDefinitionAttributes) Roles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iad.ref.Append("roles"))
}

func (iad IamActionDefinitionAttributes) Users() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iad.ref.Append("users"))
}

type ScpActionDefinitionAttributes struct {
	ref terra.Reference
}

func (sad ScpActionDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return sad.ref, nil
}

func (sad ScpActionDefinitionAttributes) InternalWithRef(ref terra.Reference) ScpActionDefinitionAttributes {
	return ScpActionDefinitionAttributes{ref: ref}
}

func (sad ScpActionDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sad.ref.InternalTokens()
}

func (sad ScpActionDefinitionAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(sad.ref.Append("policy_id"))
}

func (sad ScpActionDefinitionAttributes) TargetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sad.ref.Append("target_ids"))
}

type SsmActionDefinitionAttributes struct {
	ref terra.Reference
}

func (sad SsmActionDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return sad.ref, nil
}

func (sad SsmActionDefinitionAttributes) InternalWithRef(ref terra.Reference) SsmActionDefinitionAttributes {
	return SsmActionDefinitionAttributes{ref: ref}
}

func (sad SsmActionDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sad.ref.InternalTokens()
}

func (sad SsmActionDefinitionAttributes) ActionSubType() terra.StringValue {
	return terra.ReferenceAsString(sad.ref.Append("action_sub_type"))
}

func (sad SsmActionDefinitionAttributes) InstanceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sad.ref.Append("instance_ids"))
}

func (sad SsmActionDefinitionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sad.ref.Append("region"))
}

type SubscriberAttributes struct {
	ref terra.Reference
}

func (s SubscriberAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubscriberAttributes) InternalWithRef(ref terra.Reference) SubscriberAttributes {
	return SubscriberAttributes{ref: ref}
}

func (s SubscriberAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubscriberAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("address"))
}

func (s SubscriberAttributes) SubscriptionType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("subscription_type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ActionThresholdState struct {
	ActionThresholdType  string  `json:"action_threshold_type"`
	ActionThresholdValue float64 `json:"action_threshold_value"`
}

type DefinitionState struct {
	IamActionDefinition []IamActionDefinitionState `json:"iam_action_definition"`
	ScpActionDefinition []ScpActionDefinitionState `json:"scp_action_definition"`
	SsmActionDefinition []SsmActionDefinitionState `json:"ssm_action_definition"`
}

type IamActionDefinitionState struct {
	Groups    []string `json:"groups"`
	PolicyArn string   `json:"policy_arn"`
	Roles     []string `json:"roles"`
	Users     []string `json:"users"`
}

type ScpActionDefinitionState struct {
	PolicyId  string   `json:"policy_id"`
	TargetIds []string `json:"target_ids"`
}

type SsmActionDefinitionState struct {
	ActionSubType string   `json:"action_sub_type"`
	InstanceIds   []string `json:"instance_ids"`
	Region        string   `json:"region"`
}

type SubscriberState struct {
	Address          string `json:"address"`
	SubscriptionType string `json:"subscription_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
