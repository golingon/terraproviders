// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_glue_schema

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_glue_schema.
type Resource struct {
	Name      string
	Args      Args
	state     *awsGlueSchemaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ags *Resource) Type() string {
	return "aws_glue_schema"
}

// LocalName returns the local name for [Resource].
func (ags *Resource) LocalName() string {
	return ags.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ags *Resource) Configuration() interface{} {
	return ags.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ags *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ags)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ags *Resource) Dependencies() terra.Dependencies {
	return ags.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ags *Resource) LifecycleManagement() *terra.Lifecycle {
	return ags.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ags *Resource) Attributes() awsGlueSchemaAttributes {
	return awsGlueSchemaAttributes{ref: terra.ReferenceResource(ags)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ags *Resource) ImportState(state io.Reader) error {
	ags.state = &awsGlueSchemaState{}
	if err := json.NewDecoder(state).Decode(ags.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ags.Type(), ags.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ags *Resource) State() (*awsGlueSchemaState, bool) {
	return ags.state, ags.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ags *Resource) StateMust() *awsGlueSchemaState {
	if ags.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ags.Type(), ags.LocalName()))
	}
	return ags.state
}

// Args contains the configurations for aws_glue_schema.
type Args struct {
	// Compatibility: string, required
	Compatibility terra.StringValue `hcl:"compatibility,attr" validate:"required"`
	// DataFormat: string, required
	DataFormat terra.StringValue `hcl:"data_format,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RegistryArn: string, optional
	RegistryArn terra.StringValue `hcl:"registry_arn,attr"`
	// SchemaDefinition: string, required
	SchemaDefinition terra.StringValue `hcl:"schema_definition,attr" validate:"required"`
	// SchemaName: string, required
	SchemaName terra.StringValue `hcl:"schema_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsGlueSchemaAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_schema.
func (ags awsGlueSchemaAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("arn"))
}

// Compatibility returns a reference to field compatibility of aws_glue_schema.
func (ags awsGlueSchemaAttributes) Compatibility() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("compatibility"))
}

// DataFormat returns a reference to field data_format of aws_glue_schema.
func (ags awsGlueSchemaAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("data_format"))
}

// Description returns a reference to field description of aws_glue_schema.
func (ags awsGlueSchemaAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("description"))
}

// Id returns a reference to field id of aws_glue_schema.
func (ags awsGlueSchemaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("id"))
}

// LatestSchemaVersion returns a reference to field latest_schema_version of aws_glue_schema.
func (ags awsGlueSchemaAttributes) LatestSchemaVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(ags.ref.Append("latest_schema_version"))
}

// NextSchemaVersion returns a reference to field next_schema_version of aws_glue_schema.
func (ags awsGlueSchemaAttributes) NextSchemaVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(ags.ref.Append("next_schema_version"))
}

// RegistryArn returns a reference to field registry_arn of aws_glue_schema.
func (ags awsGlueSchemaAttributes) RegistryArn() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("registry_arn"))
}

// RegistryName returns a reference to field registry_name of aws_glue_schema.
func (ags awsGlueSchemaAttributes) RegistryName() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("registry_name"))
}

// SchemaCheckpoint returns a reference to field schema_checkpoint of aws_glue_schema.
func (ags awsGlueSchemaAttributes) SchemaCheckpoint() terra.NumberValue {
	return terra.ReferenceAsNumber(ags.ref.Append("schema_checkpoint"))
}

// SchemaDefinition returns a reference to field schema_definition of aws_glue_schema.
func (ags awsGlueSchemaAttributes) SchemaDefinition() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("schema_definition"))
}

// SchemaName returns a reference to field schema_name of aws_glue_schema.
func (ags awsGlueSchemaAttributes) SchemaName() terra.StringValue {
	return terra.ReferenceAsString(ags.ref.Append("schema_name"))
}

// Tags returns a reference to field tags of aws_glue_schema.
func (ags awsGlueSchemaAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ags.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_glue_schema.
func (ags awsGlueSchemaAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ags.ref.Append("tags_all"))
}

type awsGlueSchemaState struct {
	Arn                 string            `json:"arn"`
	Compatibility       string            `json:"compatibility"`
	DataFormat          string            `json:"data_format"`
	Description         string            `json:"description"`
	Id                  string            `json:"id"`
	LatestSchemaVersion float64           `json:"latest_schema_version"`
	NextSchemaVersion   float64           `json:"next_schema_version"`
	RegistryArn         string            `json:"registry_arn"`
	RegistryName        string            `json:"registry_name"`
	SchemaCheckpoint    float64           `json:"schema_checkpoint"`
	SchemaDefinition    string            `json:"schema_definition"`
	SchemaName          string            `json:"schema_name"`
	Tags                map[string]string `json:"tags"`
	TagsAll             map[string]string `json:"tags_all"`
}
