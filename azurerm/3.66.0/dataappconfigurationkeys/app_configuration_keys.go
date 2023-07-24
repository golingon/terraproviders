// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataappconfigurationkeys

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Items struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ItemsAttributes struct {
	ref terra.Reference
}

func (i ItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ItemsAttributes) InternalWithRef(ref terra.Reference) ItemsAttributes {
	return ItemsAttributes{ref: ref}
}

func (i ItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ItemsAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("content_type"))
}

func (i ItemsAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("etag"))
}

func (i ItemsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("key"))
}

func (i ItemsAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("label"))
}

func (i ItemsAttributes) Locked() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("locked"))
}

func (i ItemsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

func (i ItemsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

func (i ItemsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("value"))
}

func (i ItemsAttributes) VaultKeyReference() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("vault_key_reference"))
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

type ItemsState struct {
	ContentType       string            `json:"content_type"`
	Etag              string            `json:"etag"`
	Key               string            `json:"key"`
	Label             string            `json:"label"`
	Locked            bool              `json:"locked"`
	Tags              map[string]string `json:"tags"`
	Type              string            `json:"type"`
	Value             string            `json:"value"`
	VaultKeyReference string            `json:"vault_key_reference"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
