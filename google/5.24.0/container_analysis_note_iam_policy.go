// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewContainerAnalysisNoteIamPolicy creates a new instance of [ContainerAnalysisNoteIamPolicy].
func NewContainerAnalysisNoteIamPolicy(name string, args ContainerAnalysisNoteIamPolicyArgs) *ContainerAnalysisNoteIamPolicy {
	return &ContainerAnalysisNoteIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAnalysisNoteIamPolicy)(nil)

// ContainerAnalysisNoteIamPolicy represents the Terraform resource google_container_analysis_note_iam_policy.
type ContainerAnalysisNoteIamPolicy struct {
	Name      string
	Args      ContainerAnalysisNoteIamPolicyArgs
	state     *containerAnalysisNoteIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) Type() string {
	return "google_container_analysis_note_iam_policy"
}

// LocalName returns the local name for [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) LocalName() string {
	return canip.Name
}

// Configuration returns the configuration (args) for [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) Configuration() interface{} {
	return canip.Args
}

// DependOn is used for other resources to depend on [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(canip)
}

// Dependencies returns the list of resources [ContainerAnalysisNoteIamPolicy] depends_on.
func (canip *ContainerAnalysisNoteIamPolicy) Dependencies() terra.Dependencies {
	return canip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return canip.Lifecycle
}

// Attributes returns the attributes for [ContainerAnalysisNoteIamPolicy].
func (canip *ContainerAnalysisNoteIamPolicy) Attributes() containerAnalysisNoteIamPolicyAttributes {
	return containerAnalysisNoteIamPolicyAttributes{ref: terra.ReferenceResource(canip)}
}

// ImportState imports the given attribute values into [ContainerAnalysisNoteIamPolicy]'s state.
func (canip *ContainerAnalysisNoteIamPolicy) ImportState(av io.Reader) error {
	canip.state = &containerAnalysisNoteIamPolicyState{}
	if err := json.NewDecoder(av).Decode(canip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", canip.Type(), canip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAnalysisNoteIamPolicy] has state.
func (canip *ContainerAnalysisNoteIamPolicy) State() (*containerAnalysisNoteIamPolicyState, bool) {
	return canip.state, canip.state != nil
}

// StateMust returns the state for [ContainerAnalysisNoteIamPolicy]. Panics if the state is nil.
func (canip *ContainerAnalysisNoteIamPolicy) StateMust() *containerAnalysisNoteIamPolicyState {
	if canip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", canip.Type(), canip.LocalName()))
	}
	return canip.state
}

// ContainerAnalysisNoteIamPolicyArgs contains the configurations for google_container_analysis_note_iam_policy.
type ContainerAnalysisNoteIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Note: string, required
	Note terra.StringValue `hcl:"note,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type containerAnalysisNoteIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_container_analysis_note_iam_policy.
func (canip containerAnalysisNoteIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_analysis_note_iam_policy.
func (canip containerAnalysisNoteIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("id"))
}

// Note returns a reference to field note of google_container_analysis_note_iam_policy.
func (canip containerAnalysisNoteIamPolicyAttributes) Note() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("note"))
}

// PolicyData returns a reference to field policy_data of google_container_analysis_note_iam_policy.
func (canip containerAnalysisNoteIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_container_analysis_note_iam_policy.
func (canip containerAnalysisNoteIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(canip.ref.Append("project"))
}

type containerAnalysisNoteIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Note       string `json:"note"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
