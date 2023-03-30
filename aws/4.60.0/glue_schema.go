// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueSchema(name string, args GlueSchemaArgs) *GlueSchema {
	return &GlueSchema{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueSchema)(nil)

type GlueSchema struct {
	Name  string
	Args  GlueSchemaArgs
	state *glueSchemaState
}

func (gs *GlueSchema) Type() string {
	return "aws_glue_schema"
}

func (gs *GlueSchema) LocalName() string {
	return gs.Name
}

func (gs *GlueSchema) Configuration() interface{} {
	return gs.Args
}

func (gs *GlueSchema) Attributes() glueSchemaAttributes {
	return glueSchemaAttributes{ref: terra.ReferenceResource(gs)}
}

func (gs *GlueSchema) ImportState(av io.Reader) error {
	gs.state = &glueSchemaState{}
	if err := json.NewDecoder(av).Decode(gs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gs.Type(), gs.LocalName(), err)
	}
	return nil
}

func (gs *GlueSchema) State() (*glueSchemaState, bool) {
	return gs.state, gs.state != nil
}

func (gs *GlueSchema) StateMust() *glueSchemaState {
	if gs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gs.Type(), gs.LocalName()))
	}
	return gs.state
}

func (gs *GlueSchema) DependOn() terra.Reference {
	return terra.ReferenceResource(gs)
}

type GlueSchemaArgs struct {
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
	// DependsOn contains resources that GlueSchema depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueSchemaAttributes struct {
	ref terra.Reference
}

func (gs glueSchemaAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("arn"))
}

func (gs glueSchemaAttributes) Compatibility() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("compatibility"))
}

func (gs glueSchemaAttributes) DataFormat() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("data_format"))
}

func (gs glueSchemaAttributes) Description() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("description"))
}

func (gs glueSchemaAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("id"))
}

func (gs glueSchemaAttributes) LatestSchemaVersion() terra.NumberValue {
	return terra.ReferenceNumber(gs.ref.Append("latest_schema_version"))
}

func (gs glueSchemaAttributes) NextSchemaVersion() terra.NumberValue {
	return terra.ReferenceNumber(gs.ref.Append("next_schema_version"))
}

func (gs glueSchemaAttributes) RegistryArn() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("registry_arn"))
}

func (gs glueSchemaAttributes) RegistryName() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("registry_name"))
}

func (gs glueSchemaAttributes) SchemaCheckpoint() terra.NumberValue {
	return terra.ReferenceNumber(gs.ref.Append("schema_checkpoint"))
}

func (gs glueSchemaAttributes) SchemaDefinition() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("schema_definition"))
}

func (gs glueSchemaAttributes) SchemaName() terra.StringValue {
	return terra.ReferenceString(gs.ref.Append("schema_name"))
}

func (gs glueSchemaAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gs.ref.Append("tags"))
}

func (gs glueSchemaAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gs.ref.Append("tags_all"))
}

type glueSchemaState struct {
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
