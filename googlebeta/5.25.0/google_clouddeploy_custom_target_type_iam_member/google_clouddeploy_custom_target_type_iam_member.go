// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_clouddeploy_custom_target_type_iam_member

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_clouddeploy_custom_target_type_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleClouddeployCustomTargetTypeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gccttim *Resource) Type() string {
	return "google_clouddeploy_custom_target_type_iam_member"
}

// LocalName returns the local name for [Resource].
func (gccttim *Resource) LocalName() string {
	return gccttim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gccttim *Resource) Configuration() interface{} {
	return gccttim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gccttim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gccttim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gccttim *Resource) Dependencies() terra.Dependencies {
	return gccttim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gccttim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gccttim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gccttim *Resource) Attributes() googleClouddeployCustomTargetTypeIamMemberAttributes {
	return googleClouddeployCustomTargetTypeIamMemberAttributes{ref: terra.ReferenceResource(gccttim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gccttim *Resource) ImportState(state io.Reader) error {
	gccttim.state = &googleClouddeployCustomTargetTypeIamMemberState{}
	if err := json.NewDecoder(state).Decode(gccttim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gccttim.Type(), gccttim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gccttim *Resource) State() (*googleClouddeployCustomTargetTypeIamMemberState, bool) {
	return gccttim.state, gccttim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gccttim *Resource) StateMust() *googleClouddeployCustomTargetTypeIamMemberState {
	if gccttim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gccttim.Type(), gccttim.LocalName()))
	}
	return gccttim.state
}

// Args contains the configurations for google_clouddeploy_custom_target_type_iam_member.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleClouddeployCustomTargetTypeIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("location"))
}

// Member returns a reference to field member of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("member"))
}

// Name returns a reference to field name of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("name"))
}

// Project returns a reference to field project of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("project"))
}

// Role returns a reference to field role of google_clouddeploy_custom_target_type_iam_member.
func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gccttim.ref.Append("role"))
}

func (gccttim googleClouddeployCustomTargetTypeIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gccttim.ref.Append("condition"))
}

type googleClouddeployCustomTargetTypeIamMemberState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Location  string           `json:"location"`
	Member    string           `json:"member"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
