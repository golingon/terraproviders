// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudidentitygroupmembership

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MemberKey struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}

type PreferredMemberKey struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}

type Roles struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MemberKeyAttributes struct {
	ref terra.Reference
}

func (mk MemberKeyAttributes) InternalRef() (terra.Reference, error) {
	return mk.ref, nil
}

func (mk MemberKeyAttributes) InternalWithRef(ref terra.Reference) MemberKeyAttributes {
	return MemberKeyAttributes{ref: ref}
}

func (mk MemberKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mk.ref.InternalTokens()
}

func (mk MemberKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mk.ref.Append("id"))
}

func (mk MemberKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(mk.ref.Append("namespace"))
}

type PreferredMemberKeyAttributes struct {
	ref terra.Reference
}

func (pmk PreferredMemberKeyAttributes) InternalRef() (terra.Reference, error) {
	return pmk.ref, nil
}

func (pmk PreferredMemberKeyAttributes) InternalWithRef(ref terra.Reference) PreferredMemberKeyAttributes {
	return PreferredMemberKeyAttributes{ref: ref}
}

func (pmk PreferredMemberKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pmk.ref.InternalTokens()
}

func (pmk PreferredMemberKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pmk.ref.Append("id"))
}

func (pmk PreferredMemberKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(pmk.ref.Append("namespace"))
}

type RolesAttributes struct {
	ref terra.Reference
}

func (r RolesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RolesAttributes) InternalWithRef(ref terra.Reference) RolesAttributes {
	return RolesAttributes{ref: ref}
}

func (r RolesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RolesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
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

type MemberKeyState struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

type PreferredMemberKeyState struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

type RolesState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}