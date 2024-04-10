// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cosmosdbcassandratable

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AutoscaleSettings struct {
	// MaxThroughput: number, optional
	MaxThroughput terra.NumberValue `hcl:"max_throughput,attr"`
}

type Schema struct {
	// ClusterKey: min=0
	ClusterKey []ClusterKey `hcl:"cluster_key,block" validate:"min=0"`
	// Column: min=1
	Column []Column `hcl:"column,block" validate:"min=1"`
	// PartitionKey: min=1
	PartitionKey []PartitionKey `hcl:"partition_key,block" validate:"min=1"`
}

type ClusterKey struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrderBy: string, required
	OrderBy terra.StringValue `hcl:"order_by,attr" validate:"required"`
}

type Column struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type PartitionKey struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type AutoscaleSettingsAttributes struct {
	ref terra.Reference
}

func (as AutoscaleSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AutoscaleSettingsAttributes) InternalWithRef(ref terra.Reference) AutoscaleSettingsAttributes {
	return AutoscaleSettingsAttributes{ref: ref}
}

func (as AutoscaleSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AutoscaleSettingsAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("max_throughput"))
}

type SchemaAttributes struct {
	ref terra.Reference
}

func (s SchemaAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SchemaAttributes) InternalWithRef(ref terra.Reference) SchemaAttributes {
	return SchemaAttributes{ref: ref}
}

func (s SchemaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SchemaAttributes) ClusterKey() terra.ListValue[ClusterKeyAttributes] {
	return terra.ReferenceAsList[ClusterKeyAttributes](s.ref.Append("cluster_key"))
}

func (s SchemaAttributes) Column() terra.ListValue[ColumnAttributes] {
	return terra.ReferenceAsList[ColumnAttributes](s.ref.Append("column"))
}

func (s SchemaAttributes) PartitionKey() terra.ListValue[PartitionKeyAttributes] {
	return terra.ReferenceAsList[PartitionKeyAttributes](s.ref.Append("partition_key"))
}

type ClusterKeyAttributes struct {
	ref terra.Reference
}

func (ck ClusterKeyAttributes) InternalRef() (terra.Reference, error) {
	return ck.ref, nil
}

func (ck ClusterKeyAttributes) InternalWithRef(ref terra.Reference) ClusterKeyAttributes {
	return ClusterKeyAttributes{ref: ref}
}

func (ck ClusterKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ck.ref.InternalTokens()
}

func (ck ClusterKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ck.ref.Append("name"))
}

func (ck ClusterKeyAttributes) OrderBy() terra.StringValue {
	return terra.ReferenceAsString(ck.ref.Append("order_by"))
}

type ColumnAttributes struct {
	ref terra.Reference
}

func (c ColumnAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ColumnAttributes) InternalWithRef(ref terra.Reference) ColumnAttributes {
	return ColumnAttributes{ref: ref}
}

func (c ColumnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ColumnAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ColumnAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type PartitionKeyAttributes struct {
	ref terra.Reference
}

func (pk PartitionKeyAttributes) InternalRef() (terra.Reference, error) {
	return pk.ref, nil
}

func (pk PartitionKeyAttributes) InternalWithRef(ref terra.Reference) PartitionKeyAttributes {
	return PartitionKeyAttributes{ref: ref}
}

func (pk PartitionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pk.ref.InternalTokens()
}

func (pk PartitionKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("name"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AutoscaleSettingsState struct {
	MaxThroughput float64 `json:"max_throughput"`
}

type SchemaState struct {
	ClusterKey   []ClusterKeyState   `json:"cluster_key"`
	Column       []ColumnState       `json:"column"`
	PartitionKey []PartitionKeyState `json:"partition_key"`
}

type ClusterKeyState struct {
	Name    string `json:"name"`
	OrderBy string `json:"order_by"`
}

type ColumnState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PartitionKeyState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
