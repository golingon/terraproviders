// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestoreIamPolicy creates a new instance of [VertexAiFeaturestoreIamPolicy].
func NewVertexAiFeaturestoreIamPolicy(name string, args VertexAiFeaturestoreIamPolicyArgs) *VertexAiFeaturestoreIamPolicy {
	return &VertexAiFeaturestoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestoreIamPolicy)(nil)

// VertexAiFeaturestoreIamPolicy represents the Terraform resource google_vertex_ai_featurestore_iam_policy.
type VertexAiFeaturestoreIamPolicy struct {
	Name      string
	Args      VertexAiFeaturestoreIamPolicyArgs
	state     *vertexAiFeaturestoreIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) Type() string {
	return "google_vertex_ai_featurestore_iam_policy"
}

// LocalName returns the local name for [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) LocalName() string {
	return vafip.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) Configuration() interface{} {
	return vafip.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(vafip)
}

// Dependencies returns the list of resources [VertexAiFeaturestoreIamPolicy] depends_on.
func (vafip *VertexAiFeaturestoreIamPolicy) Dependencies() terra.Dependencies {
	return vafip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return vafip.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestoreIamPolicy].
func (vafip *VertexAiFeaturestoreIamPolicy) Attributes() vertexAiFeaturestoreIamPolicyAttributes {
	return vertexAiFeaturestoreIamPolicyAttributes{ref: terra.ReferenceResource(vafip)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestoreIamPolicy]'s state.
func (vafip *VertexAiFeaturestoreIamPolicy) ImportState(av io.Reader) error {
	vafip.state = &vertexAiFeaturestoreIamPolicyState{}
	if err := json.NewDecoder(av).Decode(vafip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vafip.Type(), vafip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestoreIamPolicy] has state.
func (vafip *VertexAiFeaturestoreIamPolicy) State() (*vertexAiFeaturestoreIamPolicyState, bool) {
	return vafip.state, vafip.state != nil
}

// StateMust returns the state for [VertexAiFeaturestoreIamPolicy]. Panics if the state is nil.
func (vafip *VertexAiFeaturestoreIamPolicy) StateMust() *vertexAiFeaturestoreIamPolicyState {
	if vafip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vafip.Type(), vafip.LocalName()))
	}
	return vafip.state
}

// VertexAiFeaturestoreIamPolicyArgs contains the configurations for google_vertex_ai_featurestore_iam_policy.
type VertexAiFeaturestoreIamPolicyArgs struct {
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type vertexAiFeaturestoreIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_iam_policy.
func (vafip vertexAiFeaturestoreIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("region"))
}

type vertexAiFeaturestoreIamPolicyState struct {
	Etag         string `json:"etag"`
	Featurestore string `json:"featurestore"`
	Id           string `json:"id"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
	Region       string `json:"region"`
}