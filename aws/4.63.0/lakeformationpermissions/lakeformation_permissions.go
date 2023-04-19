// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lakeformationpermissions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DataLocation struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
}

type Database struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type LfTag struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type LfTagPolicy struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// Expression: min=1,max=5
	Expression []Expression `hcl:"expression,block" validate:"min=1,max=5"`
}

type Expression struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Table struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Wildcard: bool, optional
	Wildcard terra.BoolValue `hcl:"wildcard,attr"`
}

type TableWithColumns struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ColumnNames: set of string, optional
	ColumnNames terra.SetValue[terra.StringValue] `hcl:"column_names,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// ExcludedColumnNames: set of string, optional
	ExcludedColumnNames terra.SetValue[terra.StringValue] `hcl:"excluded_column_names,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Wildcard: bool, optional
	Wildcard terra.BoolValue `hcl:"wildcard,attr"`
}

type DataLocationAttributes struct {
	ref terra.Reference
}

func (dl DataLocationAttributes) InternalRef() (terra.Reference, error) {
	return dl.ref, nil
}

func (dl DataLocationAttributes) InternalWithRef(ref terra.Reference) DataLocationAttributes {
	return DataLocationAttributes{ref: ref}
}

func (dl DataLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dl.ref.InternalTokens()
}

func (dl DataLocationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("arn"))
}

func (dl DataLocationAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("catalog_id"))
}

type DatabaseAttributes struct {
	ref terra.Reference
}

func (d DatabaseAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DatabaseAttributes) InternalWithRef(ref terra.Reference) DatabaseAttributes {
	return DatabaseAttributes{ref: ref}
}

func (d DatabaseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DatabaseAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("catalog_id"))
}

func (d DatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

type LfTagAttributes struct {
	ref terra.Reference
}

func (lt LfTagAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt LfTagAttributes) InternalWithRef(ref terra.Reference) LfTagAttributes {
	return LfTagAttributes{ref: ref}
}

func (lt LfTagAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt LfTagAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("catalog_id"))
}

func (lt LfTagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("key"))
}

func (lt LfTagAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lt.ref.Append("values"))
}

type LfTagPolicyAttributes struct {
	ref terra.Reference
}

func (ltp LfTagPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ltp.ref, nil
}

func (ltp LfTagPolicyAttributes) InternalWithRef(ref terra.Reference) LfTagPolicyAttributes {
	return LfTagPolicyAttributes{ref: ref}
}

func (ltp LfTagPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ltp.ref.InternalTokens()
}

func (ltp LfTagPolicyAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(ltp.ref.Append("catalog_id"))
}

func (ltp LfTagPolicyAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(ltp.ref.Append("resource_type"))
}

func (ltp LfTagPolicyAttributes) Expression() terra.ListValue[ExpressionAttributes] {
	return terra.ReferenceAsList[ExpressionAttributes](ltp.ref.Append("expression"))
}

type ExpressionAttributes struct {
	ref terra.Reference
}

func (e ExpressionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExpressionAttributes) InternalWithRef(ref terra.Reference) ExpressionAttributes {
	return ExpressionAttributes{ref: ref}
}

func (e ExpressionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExpressionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key"))
}

func (e ExpressionAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("values"))
}

type TableAttributes struct {
	ref terra.Reference
}

func (t TableAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TableAttributes) InternalWithRef(ref terra.Reference) TableAttributes {
	return TableAttributes{ref: ref}
}

func (t TableAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TableAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("catalog_id"))
}

func (t TableAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("database_name"))
}

func (t TableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("name"))
}

func (t TableAttributes) Wildcard() terra.BoolValue {
	return terra.ReferenceAsBool(t.ref.Append("wildcard"))
}

type TableWithColumnsAttributes struct {
	ref terra.Reference
}

func (twc TableWithColumnsAttributes) InternalRef() (terra.Reference, error) {
	return twc.ref, nil
}

func (twc TableWithColumnsAttributes) InternalWithRef(ref terra.Reference) TableWithColumnsAttributes {
	return TableWithColumnsAttributes{ref: ref}
}

func (twc TableWithColumnsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return twc.ref.InternalTokens()
}

func (twc TableWithColumnsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(twc.ref.Append("catalog_id"))
}

func (twc TableWithColumnsAttributes) ColumnNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](twc.ref.Append("column_names"))
}

func (twc TableWithColumnsAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(twc.ref.Append("database_name"))
}

func (twc TableWithColumnsAttributes) ExcludedColumnNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](twc.ref.Append("excluded_column_names"))
}

func (twc TableWithColumnsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(twc.ref.Append("name"))
}

func (twc TableWithColumnsAttributes) Wildcard() terra.BoolValue {
	return terra.ReferenceAsBool(twc.ref.Append("wildcard"))
}

type DataLocationState struct {
	Arn       string `json:"arn"`
	CatalogId string `json:"catalog_id"`
}

type DatabaseState struct {
	CatalogId string `json:"catalog_id"`
	Name      string `json:"name"`
}

type LfTagState struct {
	CatalogId string   `json:"catalog_id"`
	Key       string   `json:"key"`
	Values    []string `json:"values"`
}

type LfTagPolicyState struct {
	CatalogId    string            `json:"catalog_id"`
	ResourceType string            `json:"resource_type"`
	Expression   []ExpressionState `json:"expression"`
}

type ExpressionState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type TableState struct {
	CatalogId    string `json:"catalog_id"`
	DatabaseName string `json:"database_name"`
	Name         string `json:"name"`
	Wildcard     bool   `json:"wildcard"`
}

type TableWithColumnsState struct {
	CatalogId           string   `json:"catalog_id"`
	ColumnNames         []string `json:"column_names"`
	DatabaseName        string   `json:"database_name"`
	ExcludedColumnNames []string `json:"excluded_column_names"`
	Name                string   `json:"name"`
	Wildcard            bool     `json:"wildcard"`
}
