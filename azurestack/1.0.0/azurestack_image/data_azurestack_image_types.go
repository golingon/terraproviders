// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_image

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataDataDiskAttributes struct {
	ref terra.Reference
}

func (dd DataDataDiskAttributes) InternalRef() (terra.Reference, error) {
	return dd.ref, nil
}

func (dd DataDataDiskAttributes) InternalWithRef(ref terra.Reference) DataDataDiskAttributes {
	return DataDataDiskAttributes{ref: ref}
}

func (dd DataDataDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dd.ref.InternalTokens()
}

func (dd DataDataDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("blob_uri"))
}

func (dd DataDataDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("caching"))
}

func (dd DataDataDiskAttributes) Lun() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("lun"))
}

func (dd DataDataDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("managed_disk_id"))
}

func (dd DataDataDiskAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("size_gb"))
}

type DataOsDiskAttributes struct {
	ref terra.Reference
}

func (od DataOsDiskAttributes) InternalRef() (terra.Reference, error) {
	return od.ref, nil
}

func (od DataOsDiskAttributes) InternalWithRef(ref terra.Reference) DataOsDiskAttributes {
	return DataOsDiskAttributes{ref: ref}
}

func (od DataOsDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return od.ref.InternalTokens()
}

func (od DataOsDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("blob_uri"))
}

func (od DataOsDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("caching"))
}

func (od DataOsDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("managed_disk_id"))
}

func (od DataOsDiskAttributes) OsState() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("os_state"))
}

func (od DataOsDiskAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("os_type"))
}

func (od DataOsDiskAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(od.ref.Append("size_gb"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataDataDiskState struct {
	BlobUri       string  `json:"blob_uri"`
	Caching       string  `json:"caching"`
	Lun           float64 `json:"lun"`
	ManagedDiskId string  `json:"managed_disk_id"`
	SizeGb        float64 `json:"size_gb"`
}

type DataOsDiskState struct {
	BlobUri       string  `json:"blob_uri"`
	Caching       string  `json:"caching"`
	ManagedDiskId string  `json:"managed_disk_id"`
	OsState       string  `json:"os_state"`
	OsType        string  `json:"os_type"`
	SizeGb        float64 `json:"size_gb"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
