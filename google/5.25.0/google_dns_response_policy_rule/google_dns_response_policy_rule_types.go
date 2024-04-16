// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dns_response_policy_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type LocalData struct {
	// LocalDataLocalDatas: min=1
	LocalDatas []LocalDataLocalDatas `hcl:"local_datas,block" validate:"min=1"`
}

type LocalDataLocalDatas struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Rrdatas: list of string, optional
	Rrdatas terra.ListValue[terra.StringValue] `hcl:"rrdatas,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type LocalDataAttributes struct {
	ref terra.Reference
}

func (ld LocalDataAttributes) InternalRef() (terra.Reference, error) {
	return ld.ref, nil
}

func (ld LocalDataAttributes) InternalWithRef(ref terra.Reference) LocalDataAttributes {
	return LocalDataAttributes{ref: ref}
}

func (ld LocalDataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ld.ref.InternalTokens()
}

func (ld LocalDataAttributes) LocalDatas() terra.ListValue[LocalDataLocalDatasAttributes] {
	return terra.ReferenceAsList[LocalDataLocalDatasAttributes](ld.ref.Append("local_datas"))
}

type LocalDataLocalDatasAttributes struct {
	ref terra.Reference
}

func (ld LocalDataLocalDatasAttributes) InternalRef() (terra.Reference, error) {
	return ld.ref, nil
}

func (ld LocalDataLocalDatasAttributes) InternalWithRef(ref terra.Reference) LocalDataLocalDatasAttributes {
	return LocalDataLocalDatasAttributes{ref: ref}
}

func (ld LocalDataLocalDatasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ld.ref.InternalTokens()
}

func (ld LocalDataLocalDatasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("name"))
}

func (ld LocalDataLocalDatasAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ld.ref.Append("rrdatas"))
}

func (ld LocalDataLocalDatasAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(ld.ref.Append("ttl"))
}

func (ld LocalDataLocalDatasAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("type"))
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

type LocalDataState struct {
	LocalDatas []LocalDataLocalDatasState `json:"local_datas"`
}

type LocalDataLocalDatasState struct {
	Name    string   `json:"name"`
	Rrdatas []string `json:"rrdatas"`
	Ttl     float64  `json:"ttl"`
	Type    string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
