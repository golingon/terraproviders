// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestoreiambinding "github.com/golingon/terraproviders/googlebeta/4.73.1/vertexaifeaturestoreiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreIamBinding creates a new instance of [VertexAiFeaturestoreIamBinding].
func NewVertexAiFeaturestoreIamBinding(name string, args VertexAiFeaturestoreIamBindingArgs) *VertexAiFeaturestoreIamBinding {
	return &VertexAiFeaturestoreIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreIamBinding)(nil)

// VertexAiFeaturestoreIamBinding represents the Terraform resource google_vertex_ai_featurestore_iam_binding.
type VertexAiFeaturestoreIamBinding struct {
	Name      string
	Args      VertexAiFeaturestoreIamBindingArgs
	state     *vertexAiFeaturestoreIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) Type() string {
	return "google_vertex_ai_featurestore_iam_binding"
}

// LocalName returns the local name for [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) LocalName() string {
	return vafib.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) Configuration() interface{} {
	return vafib.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(vafib)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreIamBinding] depends_on.
func (vafib *VertexAiFeaturestoreIamBinding) Dependencies() terra.Dependencies {
	return vafib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) LifecycleManagement() *terra.Lifecycle {
	return vafib.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreIamBinding].
func (vafib *VertexAiFeaturestoreIamBinding) Attributes() vertexAiFeaturestoreIamBindingAttributes {
	return vertexAiFeaturestoreIamBindingAttributes{ref: terra.ReferenceResource(vafib)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreIamBinding]'s state.
func (vafib *VertexAiFeaturestoreIamBinding) ImportState(av io.Reader) error {
	vafib.state = &vertexAiFeaturestoreIamBindingState{}
	if err := json.NewDecoder(av).Decode(vafib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafib.Type(), vafib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreIamBinding] has state.
func (vafib *VertexAiFeaturestoreIamBinding) State() (*vertexAiFeaturestoreIamBindingState, bool) {
	return vafib.state, vafib.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreIamBinding]. Panics if the state is nil.
func (vafib *VertexAiFeaturestoreIamBinding) StateMust() *vertexAiFeaturestoreIamBindingState {
	if vafib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafib.Type(), vafib.LocalName()))
	}
	return vafib.state
}

// VertexAiFeaturestoreIamBindingArgs contains the configurations for google_vertex_ai_featurestore_iam_binding.
type VertexAiFeaturestoreIamBindingArgs struct {
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *vertexaifeaturestoreiambinding.Condition `hcl:"condition,block"`
}
type vertexAiFeaturestoreIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("id"))
}

// Members returns a reference to field members of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vafib.ref.Append("members"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("region"))
}

// Role returns a reference to field role of google_vertex_ai_featurestore_iam_binding.
func (vafib vertexAiFeaturestoreIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vafib.ref.Append("role"))
}

func (vafib vertexAiFeaturestoreIamBindingAttributes) Condition() terra.ListValue[vertexaifeaturestoreiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestoreiambinding.ConditionAttributes](vafib.ref.Append("condition"))
}

type vertexAiFeaturestoreIamBindingState struct {
	Etag         string                                          `json:"etag"`
	Featurestore string                                          `json:"featurestore"`
	Id           string                                          `json:"id"`
	Members      []string                                        `json:"members"`
	Project      string                                          `json:"project"`
	Region       string                                          `json:"region"`
	Role         string                                          `json:"role"`
	Condition    []vertexaifeaturestoreiambinding.ConditionState `json:"condition"`
}
