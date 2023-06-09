// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	timestreamwritetable "github.com/golingon/terraproviders/aws/4.60.0/timestreamwritetable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTimestreamwriteTable creates a new instance of [TimestreamwriteTable].
func NewTimestreamwriteTable(name string, args TimestreamwriteTableArgs) *TimestreamwriteTable {
	return &TimestreamwriteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TimestreamwriteTable)(nil)

// TimestreamwriteTable represents the Terraform resource aws_timestreamwrite_table.
type TimestreamwriteTable struct {
	Name      string
	Args      TimestreamwriteTableArgs
	state     *timestreamwriteTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TimestreamwriteTable].
func (tt *TimestreamwriteTable) Type() string {
	return "aws_timestreamwrite_table"
}

// LocalName returns the local name for [TimestreamwriteTable].
func (tt *TimestreamwriteTable) LocalName() string {
	return tt.Name
}

// Configuration returns the configuration (args) for [TimestreamwriteTable].
func (tt *TimestreamwriteTable) Configuration() interface{} {
	return tt.Args
}

// DependOn is used for other resources to depend on [TimestreamwriteTable].
func (tt *TimestreamwriteTable) DependOn() terra.Reference {
	return terra.ReferenceResource(tt)
}

// Dependencies returns the list of resources [TimestreamwriteTable] depends_on.
func (tt *TimestreamwriteTable) Dependencies() terra.Dependencies {
	return tt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TimestreamwriteTable].
func (tt *TimestreamwriteTable) LifecycleManagement() *terra.Lifecycle {
	return tt.Lifecycle
}

// Attributes returns the attributes for [TimestreamwriteTable].
func (tt *TimestreamwriteTable) Attributes() timestreamwriteTableAttributes {
	return timestreamwriteTableAttributes{ref: terra.ReferenceResource(tt)}
}

// ImportState imports the given attribute values into [TimestreamwriteTable]'s state.
func (tt *TimestreamwriteTable) ImportState(av io.Reader) error {
	tt.state = &timestreamwriteTableState{}
	if err := json.NewDecoder(av).Decode(tt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tt.Type(), tt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TimestreamwriteTable] has state.
func (tt *TimestreamwriteTable) State() (*timestreamwriteTableState, bool) {
	return tt.state, tt.state != nil
}

// StateMust returns the state for [TimestreamwriteTable]. Panics if the state is nil.
func (tt *TimestreamwriteTable) StateMust() *timestreamwriteTableState {
	if tt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tt.Type(), tt.LocalName()))
	}
	return tt.state
}

// TimestreamwriteTableArgs contains the configurations for aws_timestreamwrite_table.
type TimestreamwriteTableArgs struct {
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// MagneticStoreWriteProperties: optional
	MagneticStoreWriteProperties *timestreamwritetable.MagneticStoreWriteProperties `hcl:"magnetic_store_write_properties,block"`
	// RetentionProperties: optional
	RetentionProperties *timestreamwritetable.RetentionProperties `hcl:"retention_properties,block"`
}
type timestreamwriteTableAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("arn"))
}

// DatabaseName returns a reference to field database_name of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("database_name"))
}

// Id returns a reference to field id of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("id"))
}

// TableName returns a reference to field table_name of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(tt.ref.Append("table_name"))
}

// Tags returns a reference to field tags of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_timestreamwrite_table.
func (tt timestreamwriteTableAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tt.ref.Append("tags_all"))
}

func (tt timestreamwriteTableAttributes) MagneticStoreWriteProperties() terra.ListValue[timestreamwritetable.MagneticStoreWritePropertiesAttributes] {
	return terra.ReferenceAsList[timestreamwritetable.MagneticStoreWritePropertiesAttributes](tt.ref.Append("magnetic_store_write_properties"))
}

func (tt timestreamwriteTableAttributes) RetentionProperties() terra.ListValue[timestreamwritetable.RetentionPropertiesAttributes] {
	return terra.ReferenceAsList[timestreamwritetable.RetentionPropertiesAttributes](tt.ref.Append("retention_properties"))
}

type timestreamwriteTableState struct {
	Arn                          string                                                   `json:"arn"`
	DatabaseName                 string                                                   `json:"database_name"`
	Id                           string                                                   `json:"id"`
	TableName                    string                                                   `json:"table_name"`
	Tags                         map[string]string                                        `json:"tags"`
	TagsAll                      map[string]string                                        `json:"tags_all"`
	MagneticStoreWriteProperties []timestreamwritetable.MagneticStoreWritePropertiesState `json:"magnetic_store_write_properties"`
	RetentionProperties          []timestreamwritetable.RetentionPropertiesState          `json:"retention_properties"`
}
