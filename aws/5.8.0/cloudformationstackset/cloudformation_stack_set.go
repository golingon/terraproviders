// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudformationstackset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoDeployment struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RetainStacksOnAccountRemoval: bool, optional
	RetainStacksOnAccountRemoval terra.BoolValue `hcl:"retain_stacks_on_account_removal,attr"`
}

type ManagedExecution struct {
	// Active: bool, optional
	Active terra.BoolValue `hcl:"active,attr"`
}

type OperationPreferences struct {
	// FailureToleranceCount: number, optional
	FailureToleranceCount terra.NumberValue `hcl:"failure_tolerance_count,attr"`
	// FailureTolerancePercentage: number, optional
	FailureTolerancePercentage terra.NumberValue `hcl:"failure_tolerance_percentage,attr"`
	// MaxConcurrentCount: number, optional
	MaxConcurrentCount terra.NumberValue `hcl:"max_concurrent_count,attr"`
	// MaxConcurrentPercentage: number, optional
	MaxConcurrentPercentage terra.NumberValue `hcl:"max_concurrent_percentage,attr"`
	// RegionConcurrencyType: string, optional
	RegionConcurrencyType terra.StringValue `hcl:"region_concurrency_type,attr"`
	// RegionOrder: list of string, optional
	RegionOrder terra.ListValue[terra.StringValue] `hcl:"region_order,attr"`
}

type Timeouts struct {
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AutoDeploymentAttributes struct {
	ref terra.Reference
}

func (ad AutoDeploymentAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad AutoDeploymentAttributes) InternalWithRef(ref terra.Reference) AutoDeploymentAttributes {
	return AutoDeploymentAttributes{ref: ref}
}

func (ad AutoDeploymentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad AutoDeploymentAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ad.ref.Append("enabled"))
}

func (ad AutoDeploymentAttributes) RetainStacksOnAccountRemoval() terra.BoolValue {
	return terra.ReferenceAsBool(ad.ref.Append("retain_stacks_on_account_removal"))
}

type ManagedExecutionAttributes struct {
	ref terra.Reference
}

func (me ManagedExecutionAttributes) InternalRef() (terra.Reference, error) {
	return me.ref, nil
}

func (me ManagedExecutionAttributes) InternalWithRef(ref terra.Reference) ManagedExecutionAttributes {
	return ManagedExecutionAttributes{ref: ref}
}

func (me ManagedExecutionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return me.ref.InternalTokens()
}

func (me ManagedExecutionAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(me.ref.Append("active"))
}

type OperationPreferencesAttributes struct {
	ref terra.Reference
}

func (op OperationPreferencesAttributes) InternalRef() (terra.Reference, error) {
	return op.ref, nil
}

func (op OperationPreferencesAttributes) InternalWithRef(ref terra.Reference) OperationPreferencesAttributes {
	return OperationPreferencesAttributes{ref: ref}
}

func (op OperationPreferencesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return op.ref.InternalTokens()
}

func (op OperationPreferencesAttributes) FailureToleranceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(op.ref.Append("failure_tolerance_count"))
}

func (op OperationPreferencesAttributes) FailureTolerancePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(op.ref.Append("failure_tolerance_percentage"))
}

func (op OperationPreferencesAttributes) MaxConcurrentCount() terra.NumberValue {
	return terra.ReferenceAsNumber(op.ref.Append("max_concurrent_count"))
}

func (op OperationPreferencesAttributes) MaxConcurrentPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(op.ref.Append("max_concurrent_percentage"))
}

func (op OperationPreferencesAttributes) RegionConcurrencyType() terra.StringValue {
	return terra.ReferenceAsString(op.ref.Append("region_concurrency_type"))
}

func (op OperationPreferencesAttributes) RegionOrder() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](op.ref.Append("region_order"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AutoDeploymentState struct {
	Enabled                      bool `json:"enabled"`
	RetainStacksOnAccountRemoval bool `json:"retain_stacks_on_account_removal"`
}

type ManagedExecutionState struct {
	Active bool `json:"active"`
}

type OperationPreferencesState struct {
	FailureToleranceCount      float64  `json:"failure_tolerance_count"`
	FailureTolerancePercentage float64  `json:"failure_tolerance_percentage"`
	MaxConcurrentCount         float64  `json:"max_concurrent_count"`
	MaxConcurrentPercentage    float64  `json:"max_concurrent_percentage"`
	RegionConcurrencyType      string   `json:"region_concurrency_type"`
	RegionOrder                []string `json:"region_order"`
}

type TimeoutsState struct {
	Update string `json:"update"`
}
