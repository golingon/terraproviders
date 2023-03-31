// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appenginedomainmapping

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ResourceRecords struct{}

type SslSettings struct {
	// CertificateId: string, optional
	CertificateId terra.StringValue `hcl:"certificate_id,attr"`
	// SslManagementType: string, required
	SslManagementType terra.StringValue `hcl:"ssl_management_type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ResourceRecordsAttributes struct {
	ref terra.Reference
}

func (rr ResourceRecordsAttributes) InternalRef() terra.Reference {
	return rr.ref
}

func (rr ResourceRecordsAttributes) InternalWithRef(ref terra.Reference) ResourceRecordsAttributes {
	return ResourceRecordsAttributes{ref: ref}
}

func (rr ResourceRecordsAttributes) InternalTokens() hclwrite.Tokens {
	return rr.ref.InternalTokens()
}

func (rr ResourceRecordsAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rr.ref.Append("name"))
}

func (rr ResourceRecordsAttributes) Rrdata() terra.StringValue {
	return terra.ReferenceString(rr.ref.Append("rrdata"))
}

func (rr ResourceRecordsAttributes) Type() terra.StringValue {
	return terra.ReferenceString(rr.ref.Append("type"))
}

type SslSettingsAttributes struct {
	ref terra.Reference
}

func (ss SslSettingsAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss SslSettingsAttributes) InternalWithRef(ref terra.Reference) SslSettingsAttributes {
	return SslSettingsAttributes{ref: ref}
}

func (ss SslSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss SslSettingsAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("certificate_id"))
}

func (ss SslSettingsAttributes) PendingManagedCertificateId() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("pending_managed_certificate_id"))
}

func (ss SslSettingsAttributes) SslManagementType() terra.StringValue {
	return terra.ReferenceString(ss.ref.Append("ssl_management_type"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type ResourceRecordsState struct {
	Name   string `json:"name"`
	Rrdata string `json:"rrdata"`
	Type   string `json:"type"`
}

type SslSettingsState struct {
	CertificateId               string `json:"certificate_id"`
	PendingManagedCertificateId string `json:"pending_managed_certificate_id"`
	SslManagementType           string `json:"ssl_management_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
