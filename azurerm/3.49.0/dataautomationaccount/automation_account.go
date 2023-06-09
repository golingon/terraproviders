// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataautomationaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct{}

type PrivateEndpointConnection struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type PrivateEndpointConnectionAttributes struct {
	ref terra.Reference
}

func (pec PrivateEndpointConnectionAttributes) InternalRef() (terra.Reference, error) {
	return pec.ref, nil
}

func (pec PrivateEndpointConnectionAttributes) InternalWithRef(ref terra.Reference) PrivateEndpointConnectionAttributes {
	return PrivateEndpointConnectionAttributes{ref: ref}
}

func (pec PrivateEndpointConnectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pec.ref.InternalTokens()
}

func (pec PrivateEndpointConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("id"))
}

func (pec PrivateEndpointConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pec.ref.Append("name"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type PrivateEndpointConnectionState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
