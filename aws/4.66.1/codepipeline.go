// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codepipeline "github.com/golingon/terraproviders/aws/4.66.1/codepipeline"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodepipeline creates a new instance of [Codepipeline].
func NewCodepipeline(name string, args CodepipelineArgs) *Codepipeline {
	return &Codepipeline{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Codepipeline)(nil)

// Codepipeline represents the Terraform resource aws_codepipeline.
type Codepipeline struct {
	Name      string
	Args      CodepipelineArgs
	state     *codepipelineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Codepipeline].
func (c *Codepipeline) Type() string {
	return "aws_codepipeline"
}

// LocalName returns the local name for [Codepipeline].
func (c *Codepipeline) LocalName() string {
	return c.Name
}

// Configuration returns the configuration (args) for [Codepipeline].
func (c *Codepipeline) Configuration() interface{} {
	return c.Args
}

// DependOn is used for other resources to depend on [Codepipeline].
func (c *Codepipeline) DependOn() terra.Reference {
	return terra.ReferenceResource(c)
}

// Dependencies returns the list of resources [Codepipeline] depends_on.
func (c *Codepipeline) Dependencies() terra.Dependencies {
	return c.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Codepipeline].
func (c *Codepipeline) LifecycleManagement() *terra.Lifecycle {
	return c.Lifecycle
}

// Attributes returns the attributes for [Codepipeline].
func (c *Codepipeline) Attributes() codepipelineAttributes {
	return codepipelineAttributes{ref: terra.ReferenceResource(c)}
}

// ImportState imports the given attribute values into [Codepipeline]'s state.
func (c *Codepipeline) ImportState(av io.Reader) error {
	c.state = &codepipelineState{}
	if err := json.NewDecoder(av).Decode(c.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", c.Type(), c.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Codepipeline] has state.
func (c *Codepipeline) State() (*codepipelineState, bool) {
	return c.state, c.state != nil
}

// StateMust returns the state for [Codepipeline]. Panics if the state is nil.
func (c *Codepipeline) StateMust() *codepipelineState {
	if c.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", c.Type(), c.LocalName()))
	}
	return c.state
}

// CodepipelineArgs contains the configurations for aws_codepipeline.
type CodepipelineArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ArtifactStore: min=1
	ArtifactStore []codepipeline.ArtifactStore `hcl:"artifact_store,block" validate:"min=1"`
	// Stage: min=2
	Stage []codepipeline.Stage `hcl:"stage,block" validate:"min=2"`
}
type codepipelineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codepipeline.
func (c codepipelineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("arn"))
}

// Id returns a reference to field id of aws_codepipeline.
func (c codepipelineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("id"))
}

// Name returns a reference to field name of aws_codepipeline.
func (c codepipelineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_codepipeline.
func (c codepipelineAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_codepipeline.
func (c codepipelineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codepipeline.
func (c codepipelineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags_all"))
}

func (c codepipelineAttributes) ArtifactStore() terra.SetValue[codepipeline.ArtifactStoreAttributes] {
	return terra.ReferenceAsSet[codepipeline.ArtifactStoreAttributes](c.ref.Append("artifact_store"))
}

func (c codepipelineAttributes) Stage() terra.ListValue[codepipeline.StageAttributes] {
	return terra.ReferenceAsList[codepipeline.StageAttributes](c.ref.Append("stage"))
}

type codepipelineState struct {
	Arn           string                            `json:"arn"`
	Id            string                            `json:"id"`
	Name          string                            `json:"name"`
	RoleArn       string                            `json:"role_arn"`
	Tags          map[string]string                 `json:"tags"`
	TagsAll       map[string]string                 `json:"tags_all"`
	ArtifactStore []codepipeline.ArtifactStoreState `json:"artifact_store"`
	Stage         []codepipeline.StageState         `json:"stage"`
}
