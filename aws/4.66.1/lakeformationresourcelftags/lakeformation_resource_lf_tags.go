// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lakeformationresourcelftags

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

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
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
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

func (lt LfTagAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("value"))
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

type DatabaseState struct {
	CatalogId string `json:"catalog_id"`
	Name      string `json:"name"`
}

type LfTagState struct {
	CatalogId string `json:"catalog_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
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

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
