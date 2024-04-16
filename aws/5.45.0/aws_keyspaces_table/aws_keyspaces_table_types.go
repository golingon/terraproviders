// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_keyspaces_table

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CapacitySpecification struct {
	// ReadCapacityUnits: number, optional
	ReadCapacityUnits terra.NumberValue `hcl:"read_capacity_units,attr"`
	// ThroughputMode: string, optional
	ThroughputMode terra.StringValue `hcl:"throughput_mode,attr"`
	// WriteCapacityUnits: number, optional
	WriteCapacityUnits terra.NumberValue `hcl:"write_capacity_units,attr"`
}

type ClientSideTimestamps struct {
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
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
	// SchemaDefinitionClusteringKey: min=0
	ClusteringKey []SchemaDefinitionClusteringKey `hcl:"clustering_key,block" validate:"min=0"`
	// SchemaDefinitionColumn: min=1
	Column []SchemaDefinitionColumn `hcl:"column,block" validate:"min=1"`
	// SchemaDefinitionPartitionKey: min=1
	PartitionKey []SchemaDefinitionPartitionKey `hcl:"partition_key,block" validate:"min=1"`
	// SchemaDefinitionStaticColumn: min=0
	StaticColumn []SchemaDefinitionStaticColumn `hcl:"static_column,block" validate:"min=0"`
}

type SchemaDefinitionClusteringKey struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrderBy: string, required
	OrderBy terra.StringValue `hcl:"order_by,attr" validate:"required"`
}

type SchemaDefinitionColumn struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type SchemaDefinitionPartitionKey struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type SchemaDefinitionStaticColumn struct {
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

type ClientSideTimestampsAttributes struct {
	ref terra.Reference
}

func (cst ClientSideTimestampsAttributes) InternalRef() (terra.Reference, error) {
	return cst.ref, nil
}

func (cst ClientSideTimestampsAttributes) InternalWithRef(ref terra.Reference) ClientSideTimestampsAttributes {
	return ClientSideTimestampsAttributes{ref: ref}
}

func (cst ClientSideTimestampsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cst.ref.InternalTokens()
}

func (cst ClientSideTimestampsAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cst.ref.Append("status"))
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

func (sd SchemaDefinitionAttributes) ClusteringKey() terra.ListValue[SchemaDefinitionClusteringKeyAttributes] {
	return terra.ReferenceAsList[SchemaDefinitionClusteringKeyAttributes](sd.ref.Append("clustering_key"))
}

func (sd SchemaDefinitionAttributes) Column() terra.SetValue[SchemaDefinitionColumnAttributes] {
	return terra.ReferenceAsSet[SchemaDefinitionColumnAttributes](sd.ref.Append("column"))
}

func (sd SchemaDefinitionAttributes) PartitionKey() terra.ListValue[SchemaDefinitionPartitionKeyAttributes] {
	return terra.ReferenceAsList[SchemaDefinitionPartitionKeyAttributes](sd.ref.Append("partition_key"))
}

func (sd SchemaDefinitionAttributes) StaticColumn() terra.SetValue[SchemaDefinitionStaticColumnAttributes] {
	return terra.ReferenceAsSet[SchemaDefinitionStaticColumnAttributes](sd.ref.Append("static_column"))
}

type SchemaDefinitionClusteringKeyAttributes struct {
	ref terra.Reference
}

func (ck SchemaDefinitionClusteringKeyAttributes) InternalRef() (terra.Reference, error) {
	return ck.ref, nil
}

func (ck SchemaDefinitionClusteringKeyAttributes) InternalWithRef(ref terra.Reference) SchemaDefinitionClusteringKeyAttributes {
	return SchemaDefinitionClusteringKeyAttributes{ref: ref}
}

func (ck SchemaDefinitionClusteringKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ck.ref.InternalTokens()
}

func (ck SchemaDefinitionClusteringKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ck.ref.Append("name"))
}

func (ck SchemaDefinitionClusteringKeyAttributes) OrderBy() terra.StringValue {
	return terra.ReferenceAsString(ck.ref.Append("order_by"))
}

type SchemaDefinitionColumnAttributes struct {
	ref terra.Reference
}

func (c SchemaDefinitionColumnAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c SchemaDefinitionColumnAttributes) InternalWithRef(ref terra.Reference) SchemaDefinitionColumnAttributes {
	return SchemaDefinitionColumnAttributes{ref: ref}
}

func (c SchemaDefinitionColumnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c SchemaDefinitionColumnAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c SchemaDefinitionColumnAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type SchemaDefinitionPartitionKeyAttributes struct {
	ref terra.Reference
}

func (pk SchemaDefinitionPartitionKeyAttributes) InternalRef() (terra.Reference, error) {
	return pk.ref, nil
}

func (pk SchemaDefinitionPartitionKeyAttributes) InternalWithRef(ref terra.Reference) SchemaDefinitionPartitionKeyAttributes {
	return SchemaDefinitionPartitionKeyAttributes{ref: ref}
}

func (pk SchemaDefinitionPartitionKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pk.ref.InternalTokens()
}

func (pk SchemaDefinitionPartitionKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("name"))
}

type SchemaDefinitionStaticColumnAttributes struct {
	ref terra.Reference
}

func (sc SchemaDefinitionStaticColumnAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SchemaDefinitionStaticColumnAttributes) InternalWithRef(ref terra.Reference) SchemaDefinitionStaticColumnAttributes {
	return SchemaDefinitionStaticColumnAttributes{ref: ref}
}

func (sc SchemaDefinitionStaticColumnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SchemaDefinitionStaticColumnAttributes) Name() terra.StringValue {
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

type ClientSideTimestampsState struct {
	Status string `json:"status"`
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
	ClusteringKey []SchemaDefinitionClusteringKeyState `json:"clustering_key"`
	Column        []SchemaDefinitionColumnState        `json:"column"`
	PartitionKey  []SchemaDefinitionPartitionKeyState  `json:"partition_key"`
	StaticColumn  []SchemaDefinitionStaticColumnState  `json:"static_column"`
}

type SchemaDefinitionClusteringKeyState struct {
	Name    string `json:"name"`
	OrderBy string `json:"order_by"`
}

type SchemaDefinitionColumnState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SchemaDefinitionPartitionKeyState struct {
	Name string `json:"name"`
}

type SchemaDefinitionStaticColumnState struct {
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
