// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadynamodbtable "github.com/golingon/terraproviders/aws/4.60.0/datadynamodbtable"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDynamodbTable creates a new instance of [DataDynamodbTable].
func NewDataDynamodbTable(name string, args DataDynamodbTableArgs) *DataDynamodbTable {
	return &DataDynamodbTable{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDynamodbTable)(nil)

// DataDynamodbTable represents the Terraform data resource aws_dynamodb_table.
type DataDynamodbTable struct {
	Name string
	Args DataDynamodbTableArgs
}

// DataSource returns the Terraform object type for [DataDynamodbTable].
func (dt *DataDynamodbTable) DataSource() string {
	return "aws_dynamodb_table"
}

// LocalName returns the local name for [DataDynamodbTable].
func (dt *DataDynamodbTable) LocalName() string {
	return dt.Name
}

// Configuration returns the configuration (args) for [DataDynamodbTable].
func (dt *DataDynamodbTable) Configuration() interface{} {
	return dt.Args
}

// Attributes returns the attributes for [DataDynamodbTable].
func (dt *DataDynamodbTable) Attributes() dataDynamodbTableAttributes {
	return dataDynamodbTableAttributes{ref: terra.ReferenceDataResource(dt)}
}

// DataDynamodbTableArgs contains the configurations for aws_dynamodb_table.
type DataDynamodbTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Attribute: min=0
	Attribute []datadynamodbtable.Attribute `hcl:"attribute,block" validate:"min=0"`
	// GlobalSecondaryIndex: min=0
	GlobalSecondaryIndex []datadynamodbtable.GlobalSecondaryIndex `hcl:"global_secondary_index,block" validate:"min=0"`
	// LocalSecondaryIndex: min=0
	LocalSecondaryIndex []datadynamodbtable.LocalSecondaryIndex `hcl:"local_secondary_index,block" validate:"min=0"`
	// PointInTimeRecovery: min=0
	PointInTimeRecovery []datadynamodbtable.PointInTimeRecovery `hcl:"point_in_time_recovery,block" validate:"min=0"`
	// Replica: min=0
	Replica []datadynamodbtable.Replica `hcl:"replica,block" validate:"min=0"`
	// Ttl: min=0
	Ttl []datadynamodbtable.Ttl `hcl:"ttl,block" validate:"min=0"`
	// ServerSideEncryption: optional
	ServerSideEncryption *datadynamodbtable.ServerSideEncryption `hcl:"server_side_encryption,block"`
}
type dataDynamodbTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("arn"))
}

// BillingMode returns a reference to field billing_mode of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) BillingMode() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("billing_mode"))
}

// DeletionProtectionEnabled returns a reference to field deletion_protection_enabled of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) DeletionProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dt.ref.Append("deletion_protection_enabled"))
}

// HashKey returns a reference to field hash_key of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) HashKey() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("hash_key"))
}

// Id returns a reference to field id of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("id"))
}

// Name returns a reference to field name of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("name"))
}

// RangeKey returns a reference to field range_key of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) RangeKey() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("range_key"))
}

// ReadCapacity returns a reference to field read_capacity of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) ReadCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dt.ref.Append("read_capacity"))
}

// StreamArn returns a reference to field stream_arn of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_arn"))
}

// StreamEnabled returns a reference to field stream_enabled of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) StreamEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dt.ref.Append("stream_enabled"))
}

// StreamLabel returns a reference to field stream_label of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) StreamLabel() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_label"))
}

// StreamViewType returns a reference to field stream_view_type of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) StreamViewType() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_view_type"))
}

// TableClass returns a reference to field table_class of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) TableClass() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("table_class"))
}

// Tags returns a reference to field tags of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("tags"))
}

// WriteCapacity returns a reference to field write_capacity of aws_dynamodb_table.
func (dt dataDynamodbTableAttributes) WriteCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dt.ref.Append("write_capacity"))
}

func (dt dataDynamodbTableAttributes) Attribute() terra.SetValue[datadynamodbtable.AttributeAttributes] {
	return terra.ReferenceAsSet[datadynamodbtable.AttributeAttributes](dt.ref.Append("attribute"))
}

func (dt dataDynamodbTableAttributes) GlobalSecondaryIndex() terra.SetValue[datadynamodbtable.GlobalSecondaryIndexAttributes] {
	return terra.ReferenceAsSet[datadynamodbtable.GlobalSecondaryIndexAttributes](dt.ref.Append("global_secondary_index"))
}

func (dt dataDynamodbTableAttributes) LocalSecondaryIndex() terra.SetValue[datadynamodbtable.LocalSecondaryIndexAttributes] {
	return terra.ReferenceAsSet[datadynamodbtable.LocalSecondaryIndexAttributes](dt.ref.Append("local_secondary_index"))
}

func (dt dataDynamodbTableAttributes) PointInTimeRecovery() terra.ListValue[datadynamodbtable.PointInTimeRecoveryAttributes] {
	return terra.ReferenceAsList[datadynamodbtable.PointInTimeRecoveryAttributes](dt.ref.Append("point_in_time_recovery"))
}

func (dt dataDynamodbTableAttributes) Replica() terra.SetValue[datadynamodbtable.ReplicaAttributes] {
	return terra.ReferenceAsSet[datadynamodbtable.ReplicaAttributes](dt.ref.Append("replica"))
}

func (dt dataDynamodbTableAttributes) Ttl() terra.SetValue[datadynamodbtable.TtlAttributes] {
	return terra.ReferenceAsSet[datadynamodbtable.TtlAttributes](dt.ref.Append("ttl"))
}

func (dt dataDynamodbTableAttributes) ServerSideEncryption() terra.ListValue[datadynamodbtable.ServerSideEncryptionAttributes] {
	return terra.ReferenceAsList[datadynamodbtable.ServerSideEncryptionAttributes](dt.ref.Append("server_side_encryption"))
}
