// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_forwarding_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ServiceDirectoryRegistrations struct {
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Service: string, optional
	Service terra.StringValue `hcl:"service,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ServiceDirectoryRegistrationsAttributes struct {
	ref terra.Reference
}

func (sdr ServiceDirectoryRegistrationsAttributes) InternalRef() (terra.Reference, error) {
	return sdr.ref, nil
}

func (sdr ServiceDirectoryRegistrationsAttributes) InternalWithRef(ref terra.Reference) ServiceDirectoryRegistrationsAttributes {
	return ServiceDirectoryRegistrationsAttributes{ref: ref}
}

func (sdr ServiceDirectoryRegistrationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sdr.ref.InternalTokens()
}

func (sdr ServiceDirectoryRegistrationsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(sdr.ref.Append("namespace"))
}

func (sdr ServiceDirectoryRegistrationsAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(sdr.ref.Append("service"))
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

type ServiceDirectoryRegistrationsState struct {
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
