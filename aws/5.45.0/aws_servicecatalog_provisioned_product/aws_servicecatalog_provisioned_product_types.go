// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_servicecatalog_provisioned_product

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ProvisioningParameters struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// UsePreviousValue: bool, optional
	UsePreviousValue terra.BoolValue `hcl:"use_previous_value,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type StackSetProvisioningPreferences struct {
	// Accounts: list of string, optional
	Accounts terra.ListValue[terra.StringValue] `hcl:"accounts,attr"`
	// FailureToleranceCount: number, optional
	FailureToleranceCount terra.NumberValue `hcl:"failure_tolerance_count,attr"`
	// FailureTolerancePercentage: number, optional
	FailureTolerancePercentage terra.NumberValue `hcl:"failure_tolerance_percentage,attr"`
	// MaxConcurrencyCount: number, optional
	MaxConcurrencyCount terra.NumberValue `hcl:"max_concurrency_count,attr"`
	// MaxConcurrencyPercentage: number, optional
	MaxConcurrencyPercentage terra.NumberValue `hcl:"max_concurrency_percentage,attr"`
	// Regions: list of string, optional
	Regions terra.ListValue[terra.StringValue] `hcl:"regions,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type OutputsAttributes struct {
	ref terra.Reference
}

func (o OutputsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OutputsAttributes) InternalWithRef(ref terra.Reference) OutputsAttributes {
	return OutputsAttributes{ref: ref}
}

func (o OutputsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OutputsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("description"))
}

func (o OutputsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("key"))
}

func (o OutputsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("value"))
}

type ProvisioningParametersAttributes struct {
	ref terra.Reference
}

func (pp ProvisioningParametersAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp ProvisioningParametersAttributes) InternalWithRef(ref terra.Reference) ProvisioningParametersAttributes {
	return ProvisioningParametersAttributes{ref: ref}
}

func (pp ProvisioningParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp ProvisioningParametersAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("key"))
}

func (pp ProvisioningParametersAttributes) UsePreviousValue() terra.BoolValue {
	return terra.ReferenceAsBool(pp.ref.Append("use_previous_value"))
}

func (pp ProvisioningParametersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("value"))
}

type StackSetProvisioningPreferencesAttributes struct {
	ref terra.Reference
}

func (sspp StackSetProvisioningPreferencesAttributes) InternalRef() (terra.Reference, error) {
	return sspp.ref, nil
}

func (sspp StackSetProvisioningPreferencesAttributes) InternalWithRef(ref terra.Reference) StackSetProvisioningPreferencesAttributes {
	return StackSetProvisioningPreferencesAttributes{ref: ref}
}

func (sspp StackSetProvisioningPreferencesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sspp.ref.InternalTokens()
}

func (sspp StackSetProvisioningPreferencesAttributes) Accounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sspp.ref.Append("accounts"))
}

func (sspp StackSetProvisioningPreferencesAttributes) FailureToleranceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sspp.ref.Append("failure_tolerance_count"))
}

func (sspp StackSetProvisioningPreferencesAttributes) FailureTolerancePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sspp.ref.Append("failure_tolerance_percentage"))
}

func (sspp StackSetProvisioningPreferencesAttributes) MaxConcurrencyCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sspp.ref.Append("max_concurrency_count"))
}

func (sspp StackSetProvisioningPreferencesAttributes) MaxConcurrencyPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(sspp.ref.Append("max_concurrency_percentage"))
}

func (sspp StackSetProvisioningPreferencesAttributes) Regions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sspp.ref.Append("regions"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type OutputsState struct {
	Description string `json:"description"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

type ProvisioningParametersState struct {
	Key              string `json:"key"`
	UsePreviousValue bool   `json:"use_previous_value"`
	Value            string `json:"value"`
}

type StackSetProvisioningPreferencesState struct {
	Accounts                   []string `json:"accounts"`
	FailureToleranceCount      float64  `json:"failure_tolerance_count"`
	FailureTolerancePercentage float64  `json:"failure_tolerance_percentage"`
	MaxConcurrencyCount        float64  `json:"max_concurrency_count"`
	MaxConcurrencyPercentage   float64  `json:"max_concurrency_percentage"`
	Regions                    []string `json:"regions"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
