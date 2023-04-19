// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dynamodbtable "github.com/golingon/terraproviders/aws/4.63.0/dynamodbtable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDynamodbTable creates a new instance of [DynamodbTable].
func NewDynamodbTable(name string, args DynamodbTableArgs) *DynamodbTable {
	return &DynamodbTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DynamodbTable)(nil)

// DynamodbTable represents the Terraform resource aws_dynamodb_table.
type DynamodbTable struct {
	Name      string
	Args      DynamodbTableArgs
	state     *dynamodbTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DynamodbTable].
func (dt *DynamodbTable) Type() string {
	return "aws_dynamodb_table"
}

// LocalName returns the local name for [DynamodbTable].
func (dt *DynamodbTable) LocalName() string {
	return dt.Name
}

// Configuration returns the configuration (args) for [DynamodbTable].
func (dt *DynamodbTable) Configuration() interface{} {
	return dt.Args
}

// DependOn is used for other resources to depend on [DynamodbTable].
func (dt *DynamodbTable) DependOn() terra.Reference {
	return terra.ReferenceResource(dt)
}

// Dependencies returns the list of resources [DynamodbTable] depends_on.
func (dt *DynamodbTable) Dependencies() terra.Dependencies {
	return dt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DynamodbTable].
func (dt *DynamodbTable) LifecycleManagement() *terra.Lifecycle {
	return dt.Lifecycle
}

// Attributes returns the attributes for [DynamodbTable].
func (dt *DynamodbTable) Attributes() dynamodbTableAttributes {
	return dynamodbTableAttributes{ref: terra.ReferenceResource(dt)}
}

// ImportState imports the given attribute values into [DynamodbTable]'s state.
func (dt *DynamodbTable) ImportState(av io.Reader) error {
	dt.state = &dynamodbTableState{}
	if err := json.NewDecoder(av).Decode(dt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dt.Type(), dt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DynamodbTable] has state.
func (dt *DynamodbTable) State() (*dynamodbTableState, bool) {
	return dt.state, dt.state != nil
}

// StateMust returns the state for [DynamodbTable]. Panics if the state is nil.
func (dt *DynamodbTable) StateMust() *dynamodbTableState {
	if dt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dt.Type(), dt.LocalName()))
	}
	return dt.state
}

// DynamodbTableArgs contains the configurations for aws_dynamodb_table.
type DynamodbTableArgs struct {
	// BillingMode: string, optional
	BillingMode terra.StringValue `hcl:"billing_mode,attr"`
	// DeletionProtectionEnabled: bool, optional
	DeletionProtectionEnabled terra.BoolValue `hcl:"deletion_protection_enabled,attr"`
	// HashKey: string, optional
	HashKey terra.StringValue `hcl:"hash_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RangeKey: string, optional
	RangeKey terra.StringValue `hcl:"range_key,attr"`
	// ReadCapacity: number, optional
	ReadCapacity terra.NumberValue `hcl:"read_capacity,attr"`
	// RestoreDateTime: string, optional
	RestoreDateTime terra.StringValue `hcl:"restore_date_time,attr"`
	// RestoreSourceName: string, optional
	RestoreSourceName terra.StringValue `hcl:"restore_source_name,attr"`
	// RestoreToLatestTime: bool, optional
	RestoreToLatestTime terra.BoolValue `hcl:"restore_to_latest_time,attr"`
	// StreamEnabled: bool, optional
	StreamEnabled terra.BoolValue `hcl:"stream_enabled,attr"`
	// StreamViewType: string, optional
	StreamViewType terra.StringValue `hcl:"stream_view_type,attr"`
	// TableClass: string, optional
	TableClass terra.StringValue `hcl:"table_class,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WriteCapacity: number, optional
	WriteCapacity terra.NumberValue `hcl:"write_capacity,attr"`
	// Attribute: min=0
	Attribute []dynamodbtable.Attribute `hcl:"attribute,block" validate:"min=0"`
	// GlobalSecondaryIndex: min=0
	GlobalSecondaryIndex []dynamodbtable.GlobalSecondaryIndex `hcl:"global_secondary_index,block" validate:"min=0"`
	// LocalSecondaryIndex: min=0
	LocalSecondaryIndex []dynamodbtable.LocalSecondaryIndex `hcl:"local_secondary_index,block" validate:"min=0"`
	// PointInTimeRecovery: optional
	PointInTimeRecovery *dynamodbtable.PointInTimeRecovery `hcl:"point_in_time_recovery,block"`
	// Replica: min=0
	Replica []dynamodbtable.Replica `hcl:"replica,block" validate:"min=0"`
	// ServerSideEncryption: optional
	ServerSideEncryption *dynamodbtable.ServerSideEncryption `hcl:"server_side_encryption,block"`
	// Timeouts: optional
	Timeouts *dynamodbtable.Timeouts `hcl:"timeouts,block"`
	// Ttl: optional
	Ttl *dynamodbtable.Ttl `hcl:"ttl,block"`
}
type dynamodbTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dynamodb_table.
func (dt dynamodbTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("arn"))
}

// BillingMode returns a reference to field billing_mode of aws_dynamodb_table.
func (dt dynamodbTableAttributes) BillingMode() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("billing_mode"))
}

// DeletionProtectionEnabled returns a reference to field deletion_protection_enabled of aws_dynamodb_table.
func (dt dynamodbTableAttributes) DeletionProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dt.ref.Append("deletion_protection_enabled"))
}

// HashKey returns a reference to field hash_key of aws_dynamodb_table.
func (dt dynamodbTableAttributes) HashKey() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("hash_key"))
}

// Id returns a reference to field id of aws_dynamodb_table.
func (dt dynamodbTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("id"))
}

// Name returns a reference to field name of aws_dynamodb_table.
func (dt dynamodbTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("name"))
}

// RangeKey returns a reference to field range_key of aws_dynamodb_table.
func (dt dynamodbTableAttributes) RangeKey() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("range_key"))
}

// ReadCapacity returns a reference to field read_capacity of aws_dynamodb_table.
func (dt dynamodbTableAttributes) ReadCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dt.ref.Append("read_capacity"))
}

// RestoreDateTime returns a reference to field restore_date_time of aws_dynamodb_table.
func (dt dynamodbTableAttributes) RestoreDateTime() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("restore_date_time"))
}

// RestoreSourceName returns a reference to field restore_source_name of aws_dynamodb_table.
func (dt dynamodbTableAttributes) RestoreSourceName() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("restore_source_name"))
}

// RestoreToLatestTime returns a reference to field restore_to_latest_time of aws_dynamodb_table.
func (dt dynamodbTableAttributes) RestoreToLatestTime() terra.BoolValue {
	return terra.ReferenceAsBool(dt.ref.Append("restore_to_latest_time"))
}

// StreamArn returns a reference to field stream_arn of aws_dynamodb_table.
func (dt dynamodbTableAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_arn"))
}

// StreamEnabled returns a reference to field stream_enabled of aws_dynamodb_table.
func (dt dynamodbTableAttributes) StreamEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dt.ref.Append("stream_enabled"))
}

// StreamLabel returns a reference to field stream_label of aws_dynamodb_table.
func (dt dynamodbTableAttributes) StreamLabel() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_label"))
}

// StreamViewType returns a reference to field stream_view_type of aws_dynamodb_table.
func (dt dynamodbTableAttributes) StreamViewType() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("stream_view_type"))
}

// TableClass returns a reference to field table_class of aws_dynamodb_table.
func (dt dynamodbTableAttributes) TableClass() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("table_class"))
}

// Tags returns a reference to field tags of aws_dynamodb_table.
func (dt dynamodbTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dynamodb_table.
func (dt dynamodbTableAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("tags_all"))
}

// WriteCapacity returns a reference to field write_capacity of aws_dynamodb_table.
func (dt dynamodbTableAttributes) WriteCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dt.ref.Append("write_capacity"))
}

func (dt dynamodbTableAttributes) Attribute() terra.SetValue[dynamodbtable.AttributeAttributes] {
	return terra.ReferenceAsSet[dynamodbtable.AttributeAttributes](dt.ref.Append("attribute"))
}

func (dt dynamodbTableAttributes) GlobalSecondaryIndex() terra.SetValue[dynamodbtable.GlobalSecondaryIndexAttributes] {
	return terra.ReferenceAsSet[dynamodbtable.GlobalSecondaryIndexAttributes](dt.ref.Append("global_secondary_index"))
}

func (dt dynamodbTableAttributes) LocalSecondaryIndex() terra.SetValue[dynamodbtable.LocalSecondaryIndexAttributes] {
	return terra.ReferenceAsSet[dynamodbtable.LocalSecondaryIndexAttributes](dt.ref.Append("local_secondary_index"))
}

func (dt dynamodbTableAttributes) PointInTimeRecovery() terra.ListValue[dynamodbtable.PointInTimeRecoveryAttributes] {
	return terra.ReferenceAsList[dynamodbtable.PointInTimeRecoveryAttributes](dt.ref.Append("point_in_time_recovery"))
}

func (dt dynamodbTableAttributes) Replica() terra.SetValue[dynamodbtable.ReplicaAttributes] {
	return terra.ReferenceAsSet[dynamodbtable.ReplicaAttributes](dt.ref.Append("replica"))
}

func (dt dynamodbTableAttributes) ServerSideEncryption() terra.ListValue[dynamodbtable.ServerSideEncryptionAttributes] {
	return terra.ReferenceAsList[dynamodbtable.ServerSideEncryptionAttributes](dt.ref.Append("server_side_encryption"))
}

func (dt dynamodbTableAttributes) Timeouts() dynamodbtable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dynamodbtable.TimeoutsAttributes](dt.ref.Append("timeouts"))
}

func (dt dynamodbTableAttributes) Ttl() terra.ListValue[dynamodbtable.TtlAttributes] {
	return terra.ReferenceAsList[dynamodbtable.TtlAttributes](dt.ref.Append("ttl"))
}

type dynamodbTableState struct {
	Arn                       string                                    `json:"arn"`
	BillingMode               string                                    `json:"billing_mode"`
	DeletionProtectionEnabled bool                                      `json:"deletion_protection_enabled"`
	HashKey                   string                                    `json:"hash_key"`
	Id                        string                                    `json:"id"`
	Name                      string                                    `json:"name"`
	RangeKey                  string                                    `json:"range_key"`
	ReadCapacity              float64                                   `json:"read_capacity"`
	RestoreDateTime           string                                    `json:"restore_date_time"`
	RestoreSourceName         string                                    `json:"restore_source_name"`
	RestoreToLatestTime       bool                                      `json:"restore_to_latest_time"`
	StreamArn                 string                                    `json:"stream_arn"`
	StreamEnabled             bool                                      `json:"stream_enabled"`
	StreamLabel               string                                    `json:"stream_label"`
	StreamViewType            string                                    `json:"stream_view_type"`
	TableClass                string                                    `json:"table_class"`
	Tags                      map[string]string                         `json:"tags"`
	TagsAll                   map[string]string                         `json:"tags_all"`
	WriteCapacity             float64                                   `json:"write_capacity"`
	Attribute                 []dynamodbtable.AttributeState            `json:"attribute"`
	GlobalSecondaryIndex      []dynamodbtable.GlobalSecondaryIndexState `json:"global_secondary_index"`
	LocalSecondaryIndex       []dynamodbtable.LocalSecondaryIndexState  `json:"local_secondary_index"`
	PointInTimeRecovery       []dynamodbtable.PointInTimeRecoveryState  `json:"point_in_time_recovery"`
	Replica                   []dynamodbtable.ReplicaState              `json:"replica"`
	ServerSideEncryption      []dynamodbtable.ServerSideEncryptionState `json:"server_side_encryption"`
	Timeouts                  *dynamodbtable.TimeoutsState              `json:"timeouts"`
	Ttl                       []dynamodbtable.TtlState                  `json:"ttl"`
}
