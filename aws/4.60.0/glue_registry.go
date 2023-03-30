// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueRegistry(name string, args GlueRegistryArgs) *GlueRegistry {
	return &GlueRegistry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueRegistry)(nil)

type GlueRegistry struct {
	Name  string
	Args  GlueRegistryArgs
	state *glueRegistryState
}

func (gr *GlueRegistry) Type() string {
	return "aws_glue_registry"
}

func (gr *GlueRegistry) LocalName() string {
	return gr.Name
}

func (gr *GlueRegistry) Configuration() interface{} {
	return gr.Args
}

func (gr *GlueRegistry) Attributes() glueRegistryAttributes {
	return glueRegistryAttributes{ref: terra.ReferenceResource(gr)}
}

func (gr *GlueRegistry) ImportState(av io.Reader) error {
	gr.state = &glueRegistryState{}
	if err := json.NewDecoder(av).Decode(gr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gr.Type(), gr.LocalName(), err)
	}
	return nil
}

func (gr *GlueRegistry) State() (*glueRegistryState, bool) {
	return gr.state, gr.state != nil
}

func (gr *GlueRegistry) StateMust() *glueRegistryState {
	if gr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gr.Type(), gr.LocalName()))
	}
	return gr.state
}

func (gr *GlueRegistry) DependOn() terra.Reference {
	return terra.ReferenceResource(gr)
}

type GlueRegistryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RegistryName: string, required
	RegistryName terra.StringValue `hcl:"registry_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that GlueRegistry depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueRegistryAttributes struct {
	ref terra.Reference
}

func (gr glueRegistryAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gr.ref.Append("arn"))
}

func (gr glueRegistryAttributes) Description() terra.StringValue {
	return terra.ReferenceString(gr.ref.Append("description"))
}

func (gr glueRegistryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gr.ref.Append("id"))
}

func (gr glueRegistryAttributes) RegistryName() terra.StringValue {
	return terra.ReferenceString(gr.ref.Append("registry_name"))
}

func (gr glueRegistryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gr.ref.Append("tags"))
}

func (gr glueRegistryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gr.ref.Append("tags_all"))
}

type glueRegistryState struct {
	Arn          string            `json:"arn"`
	Description  string            `json:"description"`
	Id           string            `json:"id"`
	RegistryName string            `json:"registry_name"`
	Tags         map[string]string `json:"tags"`
	TagsAll      map[string]string `json:"tags_all"`
}
