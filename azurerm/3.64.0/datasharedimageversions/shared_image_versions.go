// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasharedimageversions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Images struct {
	// TargetRegion: min=0
	TargetRegion []TargetRegion `hcl:"target_region,block" validate:"min=0"`
}

type TargetRegion struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ImagesAttributes struct {
	ref terra.Reference
}

func (i ImagesAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ImagesAttributes) InternalWithRef(ref terra.Reference) ImagesAttributes {
	return ImagesAttributes{ref: ref}
}

func (i ImagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ImagesAttributes) ExcludeFromLatest() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("exclude_from_latest"))
}

func (i ImagesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

func (i ImagesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("location"))
}

func (i ImagesAttributes) ManagedImageId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("managed_image_id"))
}

func (i ImagesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name"))
}

func (i ImagesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

func (i ImagesAttributes) TargetRegion() terra.ListValue[TargetRegionAttributes] {
	return terra.ReferenceAsList[TargetRegionAttributes](i.ref.Append("target_region"))
}

type TargetRegionAttributes struct {
	ref terra.Reference
}

func (tr TargetRegionAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetRegionAttributes) InternalWithRef(ref terra.Reference) TargetRegionAttributes {
	return TargetRegionAttributes{ref: ref}
}

func (tr TargetRegionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetRegionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("name"))
}

func (tr TargetRegionAttributes) RegionalReplicaCount() terra.NumberValue {
	return terra.ReferenceAsNumber(tr.ref.Append("regional_replica_count"))
}

func (tr TargetRegionAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("storage_account_type"))
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

type ImagesState struct {
	ExcludeFromLatest bool                `json:"exclude_from_latest"`
	Id                string              `json:"id"`
	Location          string              `json:"location"`
	ManagedImageId    string              `json:"managed_image_id"`
	Name              string              `json:"name"`
	Tags              map[string]string   `json:"tags"`
	TargetRegion      []TargetRegionState `json:"target_region"`
}

type TargetRegionState struct {
	Name                 string  `json:"name"`
	RegionalReplicaCount float64 `json:"regional_replica_count"`
	StorageAccountType   string  `json:"storage_account_type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}