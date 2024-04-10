// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cloudformationstacksetinstance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type StackInstanceSummaries struct{}

type DeploymentTargets struct {
	// OrganizationalUnitIds: set of string, optional
	OrganizationalUnitIds terra.SetValue[terra.StringValue] `hcl:"organizational_unit_ids,attr"`
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
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type StackInstanceSummariesAttributes struct {
	ref terra.Reference
}

func (sis StackInstanceSummariesAttributes) InternalRef() (terra.Reference, error) {
	return sis.ref, nil
}

func (sis StackInstanceSummariesAttributes) InternalWithRef(ref terra.Reference) StackInstanceSummariesAttributes {
	return StackInstanceSummariesAttributes{ref: ref}
}

func (sis StackInstanceSummariesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sis.ref.InternalTokens()
}

func (sis StackInstanceSummariesAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(sis.ref.Append("account_id"))
}

func (sis StackInstanceSummariesAttributes) OrganizationalUnitId() terra.StringValue {
	return terra.ReferenceAsString(sis.ref.Append("organizational_unit_id"))
}

func (sis StackInstanceSummariesAttributes) StackId() terra.StringValue {
	return terra.ReferenceAsString(sis.ref.Append("stack_id"))
}

type DeploymentTargetsAttributes struct {
	ref terra.Reference
}

func (dt DeploymentTargetsAttributes) InternalRef() (terra.Reference, error) {
	return dt.ref, nil
}

func (dt DeploymentTargetsAttributes) InternalWithRef(ref terra.Reference) DeploymentTargetsAttributes {
	return DeploymentTargetsAttributes{ref: ref}
}

func (dt DeploymentTargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dt.ref.InternalTokens()
}

func (dt DeploymentTargetsAttributes) OrganizationalUnitIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dt.ref.Append("organizational_unit_ids"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type StackInstanceSummariesState struct {
	AccountId            string `json:"account_id"`
	OrganizationalUnitId string `json:"organizational_unit_id"`
	StackId              string `json:"stack_id"`
}

type DeploymentTargetsState struct {
	OrganizationalUnitIds []string `json:"organizational_unit_ids"`
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
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
