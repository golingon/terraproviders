// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	keyspacestable "github.com/golingon/terraproviders/aws/5.0.1/keyspacestable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyspacesTable creates a new instance of [KeyspacesTable].
func NewKeyspacesTable(name string, args KeyspacesTableArgs) *KeyspacesTable {
	return &KeyspacesTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyspacesTable)(nil)

// KeyspacesTable represents the Terraform resource aws_keyspaces_table.
type KeyspacesTable struct {
	Name      string
	Args      KeyspacesTableArgs
	state     *keyspacesTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyspacesTable].
func (kt *KeyspacesTable) Type() string {
	return "aws_keyspaces_table"
}

// LocalName returns the local name for [KeyspacesTable].
func (kt *KeyspacesTable) LocalName() string {
	return kt.Name
}

// Configuration returns the configuration (args) for [KeyspacesTable].
func (kt *KeyspacesTable) Configuration() interface{} {
	return kt.Args
}

// DependOn is used for other resources to depend on [KeyspacesTable].
func (kt *KeyspacesTable) DependOn() terra.Reference {
	return terra.ReferenceResource(kt)
}

// Dependencies returns the list of resources [KeyspacesTable] depends_on.
func (kt *KeyspacesTable) Dependencies() terra.Dependencies {
	return kt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyspacesTable].
func (kt *KeyspacesTable) LifecycleManagement() *terra.Lifecycle {
	return kt.Lifecycle
}

// Attributes returns the attributes for [KeyspacesTable].
func (kt *KeyspacesTable) Attributes() keyspacesTableAttributes {
	return keyspacesTableAttributes{ref: terra.ReferenceResource(kt)}
}

// ImportState imports the given attribute values into [KeyspacesTable]'s state.
func (kt *KeyspacesTable) ImportState(av io.Reader) error {
	kt.state = &keyspacesTableState{}
	if err := json.NewDecoder(av).Decode(kt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kt.Type(), kt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyspacesTable] has state.
func (kt *KeyspacesTable) State() (*keyspacesTableState, bool) {
	return kt.state, kt.state != nil
}

// StateMust returns the state for [KeyspacesTable]. Panics if the state is nil.
func (kt *KeyspacesTable) StateMust() *keyspacesTableState {
	if kt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kt.Type(), kt.LocalName()))
	}
	return kt.state
}

// KeyspacesTableArgs contains the configurations for aws_keyspaces_table.
type KeyspacesTableArgs struct {
	// DefaultTimeToLive: number, optional
	DefaultTimeToLive terra.NumberValue `hcl:"default_time_to_live,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyspaceName: string, required
	KeyspaceName terra.StringValue `hcl:"keyspace_name,attr" validate:"required"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CapacitySpecification: optional
	CapacitySpecification *keyspacestable.CapacitySpecification `hcl:"capacity_specification,block"`
	// Comment: optional
	Comment *keyspacestable.Comment `hcl:"comment,block"`
	// EncryptionSpecification: optional
	EncryptionSpecification *keyspacestable.EncryptionSpecification `hcl:"encryption_specification,block"`
	// PointInTimeRecovery: optional
	PointInTimeRecovery *keyspacestable.PointInTimeRecovery `hcl:"point_in_time_recovery,block"`
	// SchemaDefinition: required
	SchemaDefinition *keyspacestable.SchemaDefinition `hcl:"schema_definition,block" validate:"required"`
	// Timeouts: optional
	Timeouts *keyspacestable.Timeouts `hcl:"timeouts,block"`
	// Ttl: optional
	Ttl *keyspacestable.Ttl `hcl:"ttl,block"`
}
type keyspacesTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_keyspaces_table.
func (kt keyspacesTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("arn"))
}

// DefaultTimeToLive returns a reference to field default_time_to_live of aws_keyspaces_table.
func (kt keyspacesTableAttributes) DefaultTimeToLive() terra.NumberValue {
	return terra.ReferenceAsNumber(kt.ref.Append("default_time_to_live"))
}

// Id returns a reference to field id of aws_keyspaces_table.
func (kt keyspacesTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("id"))
}

// KeyspaceName returns a reference to field keyspace_name of aws_keyspaces_table.
func (kt keyspacesTableAttributes) KeyspaceName() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("keyspace_name"))
}

// TableName returns a reference to field table_name of aws_keyspaces_table.
func (kt keyspacesTableAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(kt.ref.Append("table_name"))
}

// Tags returns a reference to field tags of aws_keyspaces_table.
func (kt keyspacesTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_keyspaces_table.
func (kt keyspacesTableAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kt.ref.Append("tags_all"))
}

func (kt keyspacesTableAttributes) CapacitySpecification() terra.ListValue[keyspacestable.CapacitySpecificationAttributes] {
	return terra.ReferenceAsList[keyspacestable.CapacitySpecificationAttributes](kt.ref.Append("capacity_specification"))
}

func (kt keyspacesTableAttributes) Comment() terra.ListValue[keyspacestable.CommentAttributes] {
	return terra.ReferenceAsList[keyspacestable.CommentAttributes](kt.ref.Append("comment"))
}

func (kt keyspacesTableAttributes) EncryptionSpecification() terra.ListValue[keyspacestable.EncryptionSpecificationAttributes] {
	return terra.ReferenceAsList[keyspacestable.EncryptionSpecificationAttributes](kt.ref.Append("encryption_specification"))
}

func (kt keyspacesTableAttributes) PointInTimeRecovery() terra.ListValue[keyspacestable.PointInTimeRecoveryAttributes] {
	return terra.ReferenceAsList[keyspacestable.PointInTimeRecoveryAttributes](kt.ref.Append("point_in_time_recovery"))
}

func (kt keyspacesTableAttributes) SchemaDefinition() terra.ListValue[keyspacestable.SchemaDefinitionAttributes] {
	return terra.ReferenceAsList[keyspacestable.SchemaDefinitionAttributes](kt.ref.Append("schema_definition"))
}

func (kt keyspacesTableAttributes) Timeouts() keyspacestable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyspacestable.TimeoutsAttributes](kt.ref.Append("timeouts"))
}

func (kt keyspacesTableAttributes) Ttl() terra.ListValue[keyspacestable.TtlAttributes] {
	return terra.ReferenceAsList[keyspacestable.TtlAttributes](kt.ref.Append("ttl"))
}

type keyspacesTableState struct {
	Arn                     string                                        `json:"arn"`
	DefaultTimeToLive       float64                                       `json:"default_time_to_live"`
	Id                      string                                        `json:"id"`
	KeyspaceName            string                                        `json:"keyspace_name"`
	TableName               string                                        `json:"table_name"`
	Tags                    map[string]string                             `json:"tags"`
	TagsAll                 map[string]string                             `json:"tags_all"`
	CapacitySpecification   []keyspacestable.CapacitySpecificationState   `json:"capacity_specification"`
	Comment                 []keyspacestable.CommentState                 `json:"comment"`
	EncryptionSpecification []keyspacestable.EncryptionSpecificationState `json:"encryption_specification"`
	PointInTimeRecovery     []keyspacestable.PointInTimeRecoveryState     `json:"point_in_time_recovery"`
	SchemaDefinition        []keyspacestable.SchemaDefinitionState        `json:"schema_definition"`
	Timeouts                *keyspacestable.TimeoutsState                 `json:"timeouts"`
	Ttl                     []keyspacestable.TtlState                     `json:"ttl"`
}
