// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestoreentitytypeiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/vertexaifeaturestoreentitytypeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreEntitytypeIamBinding creates a new instance of [VertexAiFeaturestoreEntitytypeIamBinding].
func NewVertexAiFeaturestoreEntitytypeIamBinding(name string, args VertexAiFeaturestoreEntitytypeIamBindingArgs) *VertexAiFeaturestoreEntitytypeIamBinding {
	return &VertexAiFeaturestoreEntitytypeIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreEntitytypeIamBinding)(nil)

// VertexAiFeaturestoreEntitytypeIamBinding represents the Terraform resource google_vertex_ai_featurestore_entitytype_iam_binding.
type VertexAiFeaturestoreEntitytypeIamBinding struct {
	Name      string
	Args      VertexAiFeaturestoreEntitytypeIamBindingArgs
	state     *vertexAiFeaturestoreEntitytypeIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) Type() string {
	return "google_vertex_ai_featurestore_entitytype_iam_binding"
}

// LocalName returns the local name for [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) LocalName() string {
	return vafeib.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) Configuration() interface{} {
	return vafeib.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(vafeib)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreEntitytypeIamBinding] depends_on.
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) Dependencies() terra.Dependencies {
	return vafeib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) LifecycleManagement() *terra.Lifecycle {
	return vafeib.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreEntitytypeIamBinding].
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) Attributes() vertexAiFeaturestoreEntitytypeIamBindingAttributes {
	return vertexAiFeaturestoreEntitytypeIamBindingAttributes{ref: terra.ReferenceResource(vafeib)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreEntitytypeIamBinding]'s state.
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) ImportState(av io.Reader) error {
	vafeib.state = &vertexAiFeaturestoreEntitytypeIamBindingState{}
	if err := json.NewDecoder(av).Decode(vafeib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafeib.Type(), vafeib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreEntitytypeIamBinding] has state.
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) State() (*vertexAiFeaturestoreEntitytypeIamBindingState, bool) {
	return vafeib.state, vafeib.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreEntitytypeIamBinding]. Panics if the state is nil.
func (vafeib *VertexAiFeaturestoreEntitytypeIamBinding) StateMust() *vertexAiFeaturestoreEntitytypeIamBindingState {
	if vafeib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafeib.Type(), vafeib.LocalName()))
	}
	return vafeib.state
}

// VertexAiFeaturestoreEntitytypeIamBindingArgs contains the configurations for google_vertex_ai_featurestore_entitytype_iam_binding.
type VertexAiFeaturestoreEntitytypeIamBindingArgs struct {
	// Entitytype: string, required
	Entitytype terra.StringValue `hcl:"entitytype,attr" validate:"required"`
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *vertexaifeaturestoreentitytypeiambinding.Condition `hcl:"condition,block"`
}
type vertexAiFeaturestoreEntitytypeIamBindingAttributes struct {
	ref terra.Reference
}

// Entitytype returns a reference to field entitytype of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Entitytype() terra.StringValue {
	return terra.ReferenceAsString(vafeib.ref.Append("entitytype"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafeib.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafeib.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafeib.ref.Append("id"))
}

// Members returns a reference to field members of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vafeib.ref.Append("members"))
}

// Role returns a reference to field role of google_vertex_ai_featurestore_entitytype_iam_binding.
func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vafeib.ref.Append("role"))
}

func (vafeib vertexAiFeaturestoreEntitytypeIamBindingAttributes) Condition() terra.ListValue[vertexaifeaturestoreentitytypeiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestoreentitytypeiambinding.ConditionAttributes](vafeib.ref.Append("condition"))
}

type vertexAiFeaturestoreEntitytypeIamBindingState struct {
	Entitytype   string                                                    `json:"entitytype"`
	Etag         string                                                    `json:"etag"`
	Featurestore string                                                    `json:"featurestore"`
	Id           string                                                    `json:"id"`
	Members      []string                                                  `json:"members"`
	Role         string                                                    `json:"role"`
	Condition    []vertexaifeaturestoreentitytypeiambinding.ConditionState `json:"condition"`
}
