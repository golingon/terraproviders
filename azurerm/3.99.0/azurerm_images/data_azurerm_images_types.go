// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_images

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataImagesAttributes struct {
	ref terra.Reference
}

func (i DataImagesAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataImagesAttributes) InternalWithRef(ref terra.Reference) DataImagesAttributes {
	return DataImagesAttributes{ref: ref}
}

func (i DataImagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataImagesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("location"))
}

func (i DataImagesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name"))
}

func (i DataImagesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

func (i DataImagesAttributes) ZoneResilient() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("zone_resilient"))
}

func (i DataImagesAttributes) DataDisk() terra.ListValue[DataImagesDataDiskAttributes] {
	return terra.ReferenceAsList[DataImagesDataDiskAttributes](i.ref.Append("data_disk"))
}

func (i DataImagesAttributes) OsDisk() terra.ListValue[DataImagesOsDiskAttributes] {
	return terra.ReferenceAsList[DataImagesOsDiskAttributes](i.ref.Append("os_disk"))
}

type DataImagesDataDiskAttributes struct {
	ref terra.Reference
}

func (dd DataImagesDataDiskAttributes) InternalRef() (terra.Reference, error) {
	return dd.ref, nil
}

func (dd DataImagesDataDiskAttributes) InternalWithRef(ref terra.Reference) DataImagesDataDiskAttributes {
	return DataImagesDataDiskAttributes{ref: ref}
}

func (dd DataImagesDataDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dd.ref.InternalTokens()
}

func (dd DataImagesDataDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("blob_uri"))
}

func (dd DataImagesDataDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("caching"))
}

func (dd DataImagesDataDiskAttributes) Lun() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("lun"))
}

func (dd DataImagesDataDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(dd.ref.Append("managed_disk_id"))
}

func (dd DataImagesDataDiskAttributes) SizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(dd.ref.Append("size_gb"))
}

type DataImagesOsDiskAttributes struct {
	ref terra.Reference
}

func (od DataImagesOsDiskAttributes) InternalRef() (terra.Reference, error) {
	return od.ref, nil
}

func (od DataImagesOsDiskAttributes) InternalWithRef(ref terra.Reference) DataImagesOsDiskAttributes {
	return DataImagesOsDiskAttributes{ref: ref}
}

func (od DataImagesOsDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return od.ref.InternalTokens()
}

func (od DataImagesOsDiskAttributes) BlobUri() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("blob_uri"))
}

func (od DataImagesOsDiskAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("caching"))
}

func (od DataImagesOsDiskAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("disk_encryption_set_id"))
}

func (od DataImagesOsDiskAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("managed_disk_id"))
}

func (od DataImagesOsDiskAttributes) OsState() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("os_state"))
}

func (od DataImagesOsDiskAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("os_type"))
}

func (od DataImagesOsDiskAttributes) SizeGb() terra.NumberValue {
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

type DataImagesState struct {
	Location      string                    `json:"location"`
	Name          string                    `json:"name"`
	Tags          map[string]string         `json:"tags"`
	ZoneResilient bool                      `json:"zone_resilient"`
	DataDisk      []DataImagesDataDiskState `json:"data_disk"`
	OsDisk        []DataImagesOsDiskState   `json:"os_disk"`
}

type DataImagesDataDiskState struct {
	BlobUri       string  `json:"blob_uri"`
	Caching       string  `json:"caching"`
	Lun           float64 `json:"lun"`
	ManagedDiskId string  `json:"managed_disk_id"`
	SizeGb        float64 `json:"size_gb"`
}

type DataImagesOsDiskState struct {
	BlobUri             string  `json:"blob_uri"`
	Caching             string  `json:"caching"`
	DiskEncryptionSetId string  `json:"disk_encryption_set_id"`
	ManagedDiskId       string  `json:"managed_disk_id"`
	OsState             string  `json:"os_state"`
	OsType              string  `json:"os_type"`
	SizeGb              float64 `json:"size_gb"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
