// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	glueconnection "github.com/golingon/terraproviders/aws/5.9.0/glueconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueConnection creates a new instance of [GlueConnection].
func NewGlueConnection(name string, args GlueConnectionArgs) *GlueConnection {
	return &GlueConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueConnection)(nil)

// GlueConnection represents the Terraform resource aws_glue_connection.
type GlueConnection struct {
	Name      string
	Args      GlueConnectionArgs
	state     *glueConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueConnection].
func (gc *GlueConnection) Type() string {
	return "aws_glue_connection"
}

// LocalName returns the local name for [GlueConnection].
func (gc *GlueConnection) LocalName() string {
	return gc.Name
}

// Configuration returns the configuration (args) for [GlueConnection].
func (gc *GlueConnection) Configuration() interface{} {
	return gc.Args
}

// DependOn is used for other resources to depend on [GlueConnection].
func (gc *GlueConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(gc)
}

// Dependencies returns the list of resources [GlueConnection] depends_on.
func (gc *GlueConnection) Dependencies() terra.Dependencies {
	return gc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueConnection].
func (gc *GlueConnection) LifecycleManagement() *terra.Lifecycle {
	return gc.Lifecycle
}

// Attributes returns the attributes for [GlueConnection].
func (gc *GlueConnection) Attributes() glueConnectionAttributes {
	return glueConnectionAttributes{ref: terra.ReferenceResource(gc)}
}

// ImportState imports the given attribute values into [GlueConnection]'s state.
func (gc *GlueConnection) ImportState(av io.Reader) error {
	gc.state = &glueConnectionState{}
	if err := json.NewDecoder(av).Decode(gc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gc.Type(), gc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueConnection] has state.
func (gc *GlueConnection) State() (*glueConnectionState, bool) {
	return gc.state, gc.state != nil
}

// StateMust returns the state for [GlueConnection]. Panics if the state is nil.
func (gc *GlueConnection) StateMust() *glueConnectionState {
	if gc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gc.Type(), gc.LocalName()))
	}
	return gc.state
}

// GlueConnectionArgs contains the configurations for aws_glue_connection.
type GlueConnectionArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// ConnectionProperties: map of string, optional
	ConnectionProperties terra.MapValue[terra.StringValue] `hcl:"connection_properties,attr"`
	// ConnectionType: string, optional
	ConnectionType terra.StringValue `hcl:"connection_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MatchCriteria: list of string, optional
	MatchCriteria terra.ListValue[terra.StringValue] `hcl:"match_criteria,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// PhysicalConnectionRequirements: optional
	PhysicalConnectionRequirements *glueconnection.PhysicalConnectionRequirements `hcl:"physical_connection_requirements,block"`
}
type glueConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_connection.
func (gc glueConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("arn"))
}

// CatalogId returns a reference to field catalog_id of aws_glue_connection.
func (gc glueConnectionAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("catalog_id"))
}

// ConnectionProperties returns a reference to field connection_properties of aws_glue_connection.
func (gc glueConnectionAttributes) ConnectionProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gc.ref.Append("connection_properties"))
}

// ConnectionType returns a reference to field connection_type of aws_glue_connection.
func (gc glueConnectionAttributes) ConnectionType() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("connection_type"))
}

// Description returns a reference to field description of aws_glue_connection.
func (gc glueConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("description"))
}

// Id returns a reference to field id of aws_glue_connection.
func (gc glueConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("id"))
}

// MatchCriteria returns a reference to field match_criteria of aws_glue_connection.
func (gc glueConnectionAttributes) MatchCriteria() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gc.ref.Append("match_criteria"))
}

// Name returns a reference to field name of aws_glue_connection.
func (gc glueConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_glue_connection.
func (gc glueConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_glue_connection.
func (gc glueConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gc.ref.Append("tags_all"))
}

func (gc glueConnectionAttributes) PhysicalConnectionRequirements() terra.ListValue[glueconnection.PhysicalConnectionRequirementsAttributes] {
	return terra.ReferenceAsList[glueconnection.PhysicalConnectionRequirementsAttributes](gc.ref.Append("physical_connection_requirements"))
}

type glueConnectionState struct {
	Arn                            string                                               `json:"arn"`
	CatalogId                      string                                               `json:"catalog_id"`
	ConnectionProperties           map[string]string                                    `json:"connection_properties"`
	ConnectionType                 string                                               `json:"connection_type"`
	Description                    string                                               `json:"description"`
	Id                             string                                               `json:"id"`
	MatchCriteria                  []string                                             `json:"match_criteria"`
	Name                           string                                               `json:"name"`
	Tags                           map[string]string                                    `json:"tags"`
	TagsAll                        map[string]string                                    `json:"tags_all"`
	PhysicalConnectionRequirements []glueconnection.PhysicalConnectionRequirementsState `json:"physical_connection_requirements"`
}
