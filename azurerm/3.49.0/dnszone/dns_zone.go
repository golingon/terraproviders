// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dnszone

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SoaRecord struct {
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// ExpireTime: number, optional
	ExpireTime terra.NumberValue `hcl:"expire_time,attr"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// MinimumTtl: number, optional
	MinimumTtl terra.NumberValue `hcl:"minimum_ttl,attr"`
	// RefreshTime: number, optional
	RefreshTime terra.NumberValue `hcl:"refresh_time,attr"`
	// RetryTime: number, optional
	RetryTime terra.NumberValue `hcl:"retry_time,attr"`
	// SerialNumber: number, optional
	SerialNumber terra.NumberValue `hcl:"serial_number,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
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

type SoaRecordAttributes struct {
	ref terra.Reference
}

func (sr SoaRecordAttributes) InternalRef() terra.Reference {
	return sr.ref
}

func (sr SoaRecordAttributes) InternalWithRef(ref terra.Reference) SoaRecordAttributes {
	return SoaRecordAttributes{ref: ref}
}

func (sr SoaRecordAttributes) InternalTokens() hclwrite.Tokens {
	return sr.ref.InternalTokens()
}

func (sr SoaRecordAttributes) Email() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("email"))
}

func (sr SoaRecordAttributes) ExpireTime() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("expire_time"))
}

func (sr SoaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("fqdn"))
}

func (sr SoaRecordAttributes) HostName() terra.StringValue {
	return terra.ReferenceString(sr.ref.Append("host_name"))
}

func (sr SoaRecordAttributes) MinimumTtl() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("minimum_ttl"))
}

func (sr SoaRecordAttributes) RefreshTime() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("refresh_time"))
}

func (sr SoaRecordAttributes) RetryTime() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("retry_time"))
}

func (sr SoaRecordAttributes) SerialNumber() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("serial_number"))
}

func (sr SoaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sr.ref.Append("tags"))
}

func (sr SoaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(sr.ref.Append("ttl"))
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

type SoaRecordState struct {
	Email        string            `json:"email"`
	ExpireTime   float64           `json:"expire_time"`
	Fqdn         string            `json:"fqdn"`
	HostName     string            `json:"host_name"`
	MinimumTtl   float64           `json:"minimum_ttl"`
	RefreshTime  float64           `json:"refresh_time"`
	RetryTime    float64           `json:"retry_time"`
	SerialNumber float64           `json:"serial_number"`
	Tags         map[string]string `json:"tags"`
	Ttl          float64           `json:"ttl"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
