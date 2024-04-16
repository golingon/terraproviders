// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_database_migration_service_private_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VpcPeeringConfig struct {
	// Subnet: string, required
	Subnet terra.StringValue `hcl:"subnet,attr" validate:"required"`
	// VpcName: string, required
	VpcName terra.StringValue `hcl:"vpc_name,attr" validate:"required"`
}

type ErrorAttributes struct {
	ref terra.Reference
}

func (e ErrorAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ErrorAttributes) InternalWithRef(ref terra.Reference) ErrorAttributes {
	return ErrorAttributes{ref: ref}
}

func (e ErrorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ErrorAttributes) Details() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("details"))
}

func (e ErrorAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("message"))
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

type VpcPeeringConfigAttributes struct {
	ref terra.Reference
}

func (vpc VpcPeeringConfigAttributes) InternalRef() (terra.Reference, error) {
	return vpc.ref, nil
}

func (vpc VpcPeeringConfigAttributes) InternalWithRef(ref terra.Reference) VpcPeeringConfigAttributes {
	return VpcPeeringConfigAttributes{ref: ref}
}

func (vpc VpcPeeringConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vpc.ref.InternalTokens()
}

func (vpc VpcPeeringConfigAttributes) Subnet() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("subnet"))
}

func (vpc VpcPeeringConfigAttributes) VpcName() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vpc_name"))
}

type ErrorState struct {
	Details map[string]string `json:"details"`
	Message string            `json:"message"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VpcPeeringConfigState struct {
	Subnet  string `json:"subnet"`
	VpcName string `json:"vpc_name"`
}
