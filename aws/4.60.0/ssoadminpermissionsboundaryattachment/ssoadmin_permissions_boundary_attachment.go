// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssoadminpermissionsboundaryattachment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PermissionsBoundary struct {
	// ManagedPolicyArn: string, optional
	ManagedPolicyArn terra.StringValue `hcl:"managed_policy_arn,attr"`
	// CustomerManagedPolicyReference: optional
	CustomerManagedPolicyReference *CustomerManagedPolicyReference `hcl:"customer_managed_policy_reference,block"`
}

type CustomerManagedPolicyReference struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
}

type PermissionsBoundaryAttributes struct {
	ref terra.Reference
}

func (pb PermissionsBoundaryAttributes) InternalRef() terra.Reference {
	return pb.ref
}

func (pb PermissionsBoundaryAttributes) InternalWithRef(ref terra.Reference) PermissionsBoundaryAttributes {
	return PermissionsBoundaryAttributes{ref: ref}
}

func (pb PermissionsBoundaryAttributes) InternalTokens() hclwrite.Tokens {
	return pb.ref.InternalTokens()
}

func (pb PermissionsBoundaryAttributes) ManagedPolicyArn() terra.StringValue {
	return terra.ReferenceString(pb.ref.Append("managed_policy_arn"))
}

func (pb PermissionsBoundaryAttributes) CustomerManagedPolicyReference() terra.ListValue[CustomerManagedPolicyReferenceAttributes] {
	return terra.ReferenceList[CustomerManagedPolicyReferenceAttributes](pb.ref.Append("customer_managed_policy_reference"))
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
	return terra.ReferenceString(cmpr.ref.Append("name"))
}

func (cmpr CustomerManagedPolicyReferenceAttributes) Path() terra.StringValue {
	return terra.ReferenceString(cmpr.ref.Append("path"))
}

type PermissionsBoundaryState struct {
	ManagedPolicyArn               string                                `json:"managed_policy_arn"`
	CustomerManagedPolicyReference []CustomerManagedPolicyReferenceState `json:"customer_managed_policy_reference"`
}

type CustomerManagedPolicyReferenceState struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
