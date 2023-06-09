// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package keyspacestable

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CapacitySpecification struct {
	// ReadCapacityUnits: number, optional
	ReadCapacityUnits terra.NumberValue `hcl:"read_capacity_units,attr"`
	// ThroughputMode: string, optional
	ThroughputMode terra.StringValue `hcl:"throughput_mode,attr"`
	// WriteCapacityUnits: number, optional
	WriteCapacityUnits terra.NumberValue `hcl:"write_capacity_units,attr"`
}

type Comment struct {
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
}

type EncryptionSpecification struct {
	// KmsKeyIdentifier: string, optional
	KmsKeyIdentifier terra.StringValue `hcl:"kms_key_identifier,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type PointInTimeRecovery struct {
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
}

type SchemaDefinition struct {
	// ClusteringKey: min=0
	ClusteringKey []ClusteringKey `hcl:"clustering_key,block" validate:"min=0"`
	// Column: min=1
	Column []Column `hcl:"column,block" validate:"min=1"`
	// PartitionKey: min=1
	PartitionKey []PartitionKey `hcl:"partition_key,block" validate:"min=1"`
	// StaticColumn: min=0
	StaticColumn []StaticColumn `hcl:"static_column,block" validate:"min=0"`
}

type ClusteringKey struct {
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

type StaticColumn struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Ttl struct {
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
}

type CapacitySpecificationAttributes struct {
	ref terra.Reference
}

func (cs CapacitySpecificationAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CapacitySpecificationAttributes) InternalWithRef(ref terra.Reference) CapacitySpecificationAttributes {
	return CapacitySpecificationAttributes{ref: ref}
}

func (cs CapacitySpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CapacitySpecificationAttributes) ReadCapacityUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("read_capacity_units"))
}

func (cs CapacitySpecificationAttributes) ThroughputMode() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("throughput_mode"))
}

func (cs CapacitySpecificationAttributes) WriteCapacityUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("write_capacity_units"))
}

type CommentAttributes struct {
	ref terra.Reference
}

func (c CommentAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CommentAttributes) InternalWithRef(ref terra.Reference) CommentAttributes {
	return CommentAttributes{ref: ref}
}

func (c CommentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CommentAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("message"))
}

type EncryptionSpecificationAttributes struct {
	ref terra.Reference
}

func (es EncryptionSpecificationAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EncryptionSpecificationAttributes) InternalWithRef(ref terra.Reference) EncryptionSpecificationAttributes {
	return EncryptionSpecificationAttributes{ref: ref}
}

func (es EncryptionSpecificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EncryptionSpecificationAttributes) KmsKeyIdentifier() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("kms_key_identifier"))
}

func (es EncryptionSpecificationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("type"))
}

type PointInTimeRecoveryAttributes struct {
	ref terra.Reference
}

func (pitr PointInTimeRecoveryAttributes) InternalRef() (terra.Reference, error) {
	return pitr.ref, nil
}

func (pitr PointInTimeRecoveryAttributes) InternalWithRef(ref terra.Reference) PointInTimeRecoveryAttributes {
	return PointInTimeRecoveryAttributes{ref: ref}
}

func (pitr PointInTimeRecoveryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pitr.ref.InternalTokens()
}

func (pitr PointInTimeRecoveryAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(pitr.ref.Append("status"))
}

type SchemaDefinitionAttributes struct {
	ref terra.Reference
}

func (sd SchemaDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd SchemaDefinitionAttributes) InternalWithRef(ref terra.Reference) SchemaDefinitionAttributes {
	return SchemaDefinitionAttributes{ref: ref}
}

func (sd SchemaDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd SchemaDefinitionAttributes) ClusteringKey() terra.ListValue[ClusteringKeyAttributes] {
	return terra.ReferenceAsList[ClusteringKeyAttributes](sd.ref.Append("clustering_key"))
}

func (sd SchemaDefinitionAttributes) Column() terra.SetValue[ColumnAttributes] {
	return terra.ReferenceAsSet[ColumnAttributes](sd.ref.Append("column"))
}

func (sd SchemaDefinitionAttributes) PartitionKey() terra.ListValue[PartitionKeyAttributes] {
	return terra.ReferenceAsList[PartitionKeyAttributes](sd.ref.Append("partition_key"))
}

func (sd SchemaDefinitionAttributes) StaticColumn() terra.SetValue[StaticColumnAttributes] {
	return terra.ReferenceAsSet[StaticColumnAttributes](sd.ref.Append("static_column"))
}

type ClusteringKeyAttributes struct {
	ref terra.Reference
}

func (ck ClusteringKeyAttributes) InternalRef() (terra.Reference, error) {
	return ck.ref, nil
}

func (ck ClusteringKeyAttributes) InternalWithRef(ref terra.Reference) ClusteringKeyAttributes {
	return ClusteringKeyAttributes{ref: ref}
}

func (ck ClusteringKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ck.ref.InternalTokens()
}

func (ck ClusteringKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ck.ref.Append("name"))
}

func (ck ClusteringKeyAttributes) OrderBy() terra.StringValue {
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

type StaticColumnAttributes struct {
	ref terra.Reference
}

func (sc StaticColumnAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc StaticColumnAttributes) InternalWithRef(ref terra.Reference) StaticColumnAttributes {
	return StaticColumnAttributes{ref: ref}
}

func (sc StaticColumnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc StaticColumnAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("name"))
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

type TtlAttributes struct {
	ref terra.Reference
}

func (t TtlAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TtlAttributes) InternalWithRef(ref terra.Reference) TtlAttributes {
	return TtlAttributes{ref: ref}
}

func (t TtlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TtlAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("status"))
}

type CapacitySpecificationState struct {
	ReadCapacityUnits  float64 `json:"read_capacity_units"`
	ThroughputMode     string  `json:"throughput_mode"`
	WriteCapacityUnits float64 `json:"write_capacity_units"`
}

type CommentState struct {
	Message string `json:"message"`
}

type EncryptionSpecificationState struct {
	KmsKeyIdentifier string `json:"kms_key_identifier"`
	Type             string `json:"type"`
}

type PointInTimeRecoveryState struct {
	Status string `json:"status"`
}

type SchemaDefinitionState struct {
	ClusteringKey []ClusteringKeyState `json:"clustering_key"`
	Column        []ColumnState        `json:"column"`
	PartitionKey  []PartitionKeyState  `json:"partition_key"`
	StaticColumn  []StaticColumnState  `json:"static_column"`
}

type ClusteringKeyState struct {
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

type StaticColumnState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TtlState struct {
	Status string `json:"status"`
}
