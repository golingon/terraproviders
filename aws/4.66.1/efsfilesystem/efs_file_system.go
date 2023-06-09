// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package efsfilesystem

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SizeInBytes struct{}

type LifecyclePolicy struct {
	// TransitionToIa: string, optional
	TransitionToIa terra.StringValue `hcl:"transition_to_ia,attr"`
	// TransitionToPrimaryStorageClass: string, optional
	TransitionToPrimaryStorageClass terra.StringValue `hcl:"transition_to_primary_storage_class,attr"`
}

type SizeInBytesAttributes struct {
	ref terra.Reference
}

func (sib SizeInBytesAttributes) InternalRef() (terra.Reference, error) {
	return sib.ref, nil
}

func (sib SizeInBytesAttributes) InternalWithRef(ref terra.Reference) SizeInBytesAttributes {
	return SizeInBytesAttributes{ref: ref}
}

func (sib SizeInBytesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sib.ref.InternalTokens()
}

func (sib SizeInBytesAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(sib.ref.Append("value"))
}

func (sib SizeInBytesAttributes) ValueInIa() terra.NumberValue {
	return terra.ReferenceAsNumber(sib.ref.Append("value_in_ia"))
}

func (sib SizeInBytesAttributes) ValueInStandard() terra.NumberValue {
	return terra.ReferenceAsNumber(sib.ref.Append("value_in_standard"))
}

type LifecyclePolicyAttributes struct {
	ref terra.Reference
}

func (lp LifecyclePolicyAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp LifecyclePolicyAttributes) InternalWithRef(ref terra.Reference) LifecyclePolicyAttributes {
	return LifecyclePolicyAttributes{ref: ref}
}

func (lp LifecyclePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp LifecyclePolicyAttributes) TransitionToIa() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("transition_to_ia"))
}

func (lp LifecyclePolicyAttributes) TransitionToPrimaryStorageClass() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("transition_to_primary_storage_class"))
}

type SizeInBytesState struct {
	Value           float64 `json:"value"`
	ValueInIa       float64 `json:"value_in_ia"`
	ValueInStandard float64 `json:"value_in_standard"`
}

type LifecyclePolicyState struct {
	TransitionToIa                  string `json:"transition_to_ia"`
	TransitionToPrimaryStorageClass string `json:"transition_to_primary_storage_class"`
}
