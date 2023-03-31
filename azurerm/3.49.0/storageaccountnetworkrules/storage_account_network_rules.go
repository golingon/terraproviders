// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storageaccountnetworkrules

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PrivateLinkAccess struct {
	// EndpointResourceId: string, required
	EndpointResourceId terra.StringValue `hcl:"endpoint_resource_id,attr" validate:"required"`
	// EndpointTenantId: string, optional
	EndpointTenantId terra.StringValue `hcl:"endpoint_tenant_id,attr"`
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

type PrivateLinkAccessAttributes struct {
	ref terra.Reference
}

func (pla PrivateLinkAccessAttributes) InternalRef() terra.Reference {
	return pla.ref
}

func (pla PrivateLinkAccessAttributes) InternalWithRef(ref terra.Reference) PrivateLinkAccessAttributes {
	return PrivateLinkAccessAttributes{ref: ref}
}

func (pla PrivateLinkAccessAttributes) InternalTokens() hclwrite.Tokens {
	return pla.ref.InternalTokens()
}

func (pla PrivateLinkAccessAttributes) EndpointResourceId() terra.StringValue {
	return terra.ReferenceString(pla.ref.Append("endpoint_resource_id"))
}

func (pla PrivateLinkAccessAttributes) EndpointTenantId() terra.StringValue {
	return terra.ReferenceString(pla.ref.Append("endpoint_tenant_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type PrivateLinkAccessState struct {
	EndpointResourceId string `json:"endpoint_resource_id"`
	EndpointTenantId   string `json:"endpoint_tenant_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
