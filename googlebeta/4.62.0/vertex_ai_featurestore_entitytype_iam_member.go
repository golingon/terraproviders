// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestoreentitytypeiammember "github.com/golingon/terraproviders/googlebeta/4.62.0/vertexaifeaturestoreentitytypeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreEntitytypeIamMember creates a new instance of [VertexAiFeaturestoreEntitytypeIamMember].
func NewVertexAiFeaturestoreEntitytypeIamMember(name string, args VertexAiFeaturestoreEntitytypeIamMemberArgs) *VertexAiFeaturestoreEntitytypeIamMember {
	return &VertexAiFeaturestoreEntitytypeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreEntitytypeIamMember)(nil)

// VertexAiFeaturestoreEntitytypeIamMember represents the Terraform resource google_vertex_ai_featurestore_entitytype_iam_member.
type VertexAiFeaturestoreEntitytypeIamMember struct {
	Name      string
	Args      VertexAiFeaturestoreEntitytypeIamMemberArgs
	state     *vertexAiFeaturestoreEntitytypeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) Type() string {
	return "google_vertex_ai_featurestore_entitytype_iam_member"
}

// LocalName returns the local name for [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) LocalName() string {
	return vafeim.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) Configuration() interface{} {
	return vafeim.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(vafeim)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreEntitytypeIamMember] depends_on.
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) Dependencies() terra.Dependencies {
	return vafeim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) LifecycleManagement() *terra.Lifecycle {
	return vafeim.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreEntitytypeIamMember].
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) Attributes() vertexAiFeaturestoreEntitytypeIamMemberAttributes {
	return vertexAiFeaturestoreEntitytypeIamMemberAttributes{ref: terra.ReferenceResource(vafeim)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreEntitytypeIamMember]'s state.
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) ImportState(av io.Reader) error {
	vafeim.state = &vertexAiFeaturestoreEntitytypeIamMemberState{}
	if err := json.NewDecoder(av).Decode(vafeim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafeim.Type(), vafeim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreEntitytypeIamMember] has state.
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) State() (*vertexAiFeaturestoreEntitytypeIamMemberState, bool) {
	return vafeim.state, vafeim.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreEntitytypeIamMember]. Panics if the state is nil.
func (vafeim *VertexAiFeaturestoreEntitytypeIamMember) StateMust() *vertexAiFeaturestoreEntitytypeIamMemberState {
	if vafeim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafeim.Type(), vafeim.LocalName()))
	}
	return vafeim.state
}

// VertexAiFeaturestoreEntitytypeIamMemberArgs contains the configurations for google_vertex_ai_featurestore_entitytype_iam_member.
type VertexAiFeaturestoreEntitytypeIamMemberArgs struct {
	// Entitytype: string, required
	Entitytype terra.StringValue `hcl:"entitytype,attr" validate:"required"`
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *vertexaifeaturestoreentitytypeiammember.Condition `hcl:"condition,block"`
}
type vertexAiFeaturestoreEntitytypeIamMemberAttributes struct {
	ref terra.Reference
}

// Entitytype returns a reference to field entitytype of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Entitytype() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("entitytype"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("id"))
}

// Member returns a reference to field member of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("member"))
}

// Role returns a reference to field role of google_vertex_ai_featurestore_entitytype_iam_member.
func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(vafeim.ref.Append("role"))
}

func (vafeim vertexAiFeaturestoreEntitytypeIamMemberAttributes) Condition() terra.ListValue[vertexaifeaturestoreentitytypeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestoreentitytypeiammember.ConditionAttributes](vafeim.ref.Append("condition"))
}

type vertexAiFeaturestoreEntitytypeIamMemberState struct {
	Entitytype   string                                                   `json:"entitytype"`
	Etag         string                                                   `json:"etag"`
	Featurestore string                                                   `json:"featurestore"`
	Id           string                                                   `json:"id"`
	Member       string                                                   `json:"member"`
	Role         string                                                   `json:"role"`
	Condition    []vertexaifeaturestoreentitytypeiammember.ConditionState `json:"condition"`
}
