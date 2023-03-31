// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package image

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DataDisk struct {
	// BlobUri: string, optional
	BlobUri terra.StringValue `hcl:"blob_uri,attr"`
	// Caching: string, optional
	Caching terra.StringValue `hcl:"caching,attr"`
	// Lun: number, optional
	Lun terra.NumberValue `hcl:"lun,attr"`
	// ManagedDiskId: string, optional
	ManagedDiskId terra.StringValue `hcl:"managed_disk_id,attr"`
	// SizeGb: number, optional
	SizeGb terra.NumberValue `hcl:"size_gb,attr"`
}

type OsDisk struct {
	// BlobUri: string, optional
	BlobUri terra.StringValue `hcl:"blob_uri,attr"`
	// Caching: string, optional
	Caching terra.StringValue `hcl:"caching,attr"`
	// ManagedDiskId: string, optional
	ManagedDiskId terra.StringValue `hcl:"managed_disk_id,attr"`
	// OsState: string, optional
	OsState terra.StringValue `hcl:"os_state,attr"`
	// OsType: string, optional
	OsType terra.StringValue `hcl:"os_type,attr"`
	// SizeGb: number, optional
	SizeGb terra.NumberValue `hcl:"size_gb,attr"`
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

type DataDiskAttributes struct {
	ref terra.Reference
}

func (dd DataDiskAttributes) InternalRef() terra.Reference {
	return dd.ref
}

func (dd DataDiskAttributes) InternalWithRef(ref terra.Reference) DataDiskAttributes {
	return DataDiskAttributes{ref: ref}
}

func (dd DataDiskAttributes) InternalTokens() hclwrite.Tokens {
	return dd.ref.InternalTokens()
}

func (dd DataDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceString(dd.ref.Append("blob_uri"))
}

func (dd DataDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceString(dd.ref.Append("caching"))
}

func (dd DataDiskAttributes) Lun() terra.NumberValue {
	return terra.ReferenceNumber(dd.ref.Append("lun"))
}

func (dd DataDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceString(dd.ref.Append("managed_disk_id"))
}

func (dd DataDiskAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceNumber(dd.ref.Append("size_gb"))
}

type OsDiskAttributes struct {
	ref terra.Reference
}

func (od OsDiskAttributes) InternalRef() terra.Reference {
	return od.ref
}

func (od OsDiskAttributes) InternalWithRef(ref terra.Reference) OsDiskAttributes {
	return OsDiskAttributes{ref: ref}
}

func (od OsDiskAttributes) InternalTokens() hclwrite.Tokens {
	return od.ref.InternalTokens()
}

func (od OsDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceString(od.ref.Append("blob_uri"))
}

func (od OsDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceString(od.ref.Append("caching"))
}

func (od OsDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceString(od.ref.Append("managed_disk_id"))
}

func (od OsDiskAttributes) OsState() terra.StringValue {
	return terra.ReferenceString(od.ref.Append("os_state"))
}

func (od OsDiskAttributes) OsType() terra.StringValue {
	return terra.ReferenceString(od.ref.Append("os_type"))
}

func (od OsDiskAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceNumber(od.ref.Append("size_gb"))
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

type DataDiskState struct {
	BlobUri       string  `json:"blob_uri"`
	Caching       string  `json:"caching"`
	Lun           float64 `json:"lun"`
	ManagedDiskId string  `json:"managed_disk_id"`
	SizeGb        float64 `json:"size_gb"`
}

type OsDiskState struct {
	BlobUri       string  `json:"blob_uri"`
	Caching       string  `json:"caching"`
	ManagedDiskId string  `json:"managed_disk_id"`
	OsState       string  `json:"os_state"`
	OsType        string  `json:"os_type"`
	SizeGb        float64 `json:"size_gb"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
