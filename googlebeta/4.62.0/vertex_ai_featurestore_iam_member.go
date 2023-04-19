// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestoreiammember "github.com/golingon/terraproviders/googlebeta/4.62.0/vertexaifeaturestoreiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreIamMember creates a new instance of [VertexAiFeaturestoreIamMember].
func NewVertexAiFeaturestoreIamMember(name string, args VertexAiFeaturestoreIamMemberArgs) *VertexAiFeaturestoreIamMember {
	return &VertexAiFeaturestoreIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreIamMember)(nil)

// VertexAiFeaturestoreIamMember represents the Terraform resource google_vertex_ai_featurestore_iam_member.
type VertexAiFeaturestoreIamMember struct {
	Name      string
	Args      VertexAiFeaturestoreIamMemberArgs
	state     *vertexAiFeaturestoreIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) Type() string {
	return "google_vertex_ai_featurestore_iam_member"
}

// LocalName returns the local name for [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) LocalName() string {
	return vafim.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) Configuration() interface{} {
	return vafim.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(vafim)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreIamMember] depends_on.
func (vafim *VertexAiFeaturestoreIamMember) Dependencies() terra.Dependencies {
	return vafim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) LifecycleManagement() *terra.Lifecycle {
	return vafim.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreIamMember].
func (vafim *VertexAiFeaturestoreIamMember) Attributes() vertexAiFeaturestoreIamMemberAttributes {
	return vertexAiFeaturestoreIamMemberAttributes{ref: terra.ReferenceResource(vafim)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreIamMember]'s state.
func (vafim *VertexAiFeaturestoreIamMember) ImportState(av io.Reader) error {
	vafim.state = &vertexAiFeaturestoreIamMemberState{}
	if err := json.NewDecoder(av).Decode(vafim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafim.Type(), vafim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreIamMember] has state.
func (vafim *VertexAiFeaturestoreIamMember) State() (*vertexAiFeaturestoreIamMemberState, bool) {
	return vafim.state, vafim.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreIamMember]. Panics if the state is nil.
func (vafim *VertexAiFeaturestoreIamMember) StateMust() *vertexAiFeaturestoreIamMemberState {
	if vafim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafim.Type(), vafim.LocalName()))
	}
	return vafim.state
}

// VertexAiFeaturestoreIamMemberArgs contains the configurations for google_vertex_ai_featurestore_iam_member.
type VertexAiFeaturestoreIamMemberArgs struct {
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *vertexaifeaturestoreiammember.Condition `hcl:"condition,block"`
}
type vertexAiFeaturestoreIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("id"))
}

// Member returns a reference to field member of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("member"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("region"))
}

// Role returns a reference to field role of google_vertex_ai_featurestore_iam_member.
func (vafim vertexAiFeaturestoreIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vafim.ref.Append("role"))
}

func (vafim vertexAiFeaturestoreIamMemberAttributes) Condition() terra.ListValue[vertexaifeaturestoreiammember.ConditionAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestoreiammember.ConditionAttributes](vafim.ref.Append("condition"))
}

type vertexAiFeaturestoreIamMemberState struct {
	Etag         string                                         `json:"etag"`
	Featurestore string                                         `json:"featurestore"`
	Id           string                                         `json:"id"`
	Member       string                                         `json:"member"`
	Project      string                                         `json:"project"`
	Region       string                                         `json:"region"`
	Role         string                                         `json:"role"`
	Condition    []vertexaifeaturestoreiammember.ConditionState `json:"condition"`
}
