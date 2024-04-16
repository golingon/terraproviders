// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mssql_elasticpool

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataSkuAttributes struct {
	ref terra.Reference
}

func (s DataSkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DataSkuAttributes) InternalWithRef(ref terra.Reference) DataSkuAttributes {
	return DataSkuAttributes{ref: ref}
}

func (s DataSkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DataSkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s DataSkuAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("family"))
}

func (s DataSkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s DataSkuAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
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

type DataSkuState struct {
	Capacity float64 `json:"capacity"`
	Family   string  `json:"family"`
	Name     string  `json:"name"`
	Tier     string  `json:"tier"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
