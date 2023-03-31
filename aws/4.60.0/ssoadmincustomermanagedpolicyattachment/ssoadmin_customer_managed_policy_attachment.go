// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssoadmincustomermanagedpolicyattachment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomerManagedPolicyReference struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
}

type CustomerManagedPolicyReferenceAttributes struct {
	ref terra.Reference
}

func (cmpr CustomerManagedPolicyReferenceAttributes) InternalRef() terra.Reference {
	return cmpr.ref
}

func (cmpr CustomerManagedPolicyReferenceAttributes) InternalWithRef(ref terra.Reference) CustomerManagedPolicyReferenceAttributes {
	return CustomerManagedPolicyReferenceAttributes{ref: ref}
}

func (cmpr CustomerManagedPolicyReferenceAttributes) InternalTokens() hclwrite.Tokens {
	return cmpr.ref.InternalTokens()
}

func (cmpr CustomerManagedPolicyReferenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmpr.ref.Append("name"))
}

func (cmpr CustomerManagedPolicyReferenceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(cmpr.ref.Append("path"))
}

type CustomerManagedPolicyReferenceState struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
