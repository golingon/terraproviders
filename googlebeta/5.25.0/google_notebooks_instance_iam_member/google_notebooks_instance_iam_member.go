// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_notebooks_instance_iam_member

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

// Resource represents the Terraform resource google_notebooks_instance_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNotebooksInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gniim *Resource) Type() string {
	return "google_notebooks_instance_iam_member"
}

// LocalName returns the local name for [Resource].
func (gniim *Resource) LocalName() string {
	return gniim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gniim *Resource) Configuration() interface{} {
	return gniim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gniim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gniim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gniim *Resource) Dependencies() terra.Dependencies {
	return gniim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gniim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gniim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gniim *Resource) Attributes() googleNotebooksInstanceIamMemberAttributes {
	return googleNotebooksInstanceIamMemberAttributes{ref: terra.ReferenceResource(gniim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gniim *Resource) ImportState(state io.Reader) error {
	gniim.state = &googleNotebooksInstanceIamMemberState{}
	if err := json.NewDecoder(state).Decode(gniim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gniim.Type(), gniim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gniim *Resource) State() (*googleNotebooksInstanceIamMemberState, bool) {
	return gniim.state, gniim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gniim *Resource) StateMust() *googleNotebooksInstanceIamMemberState {
	if gniim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gniim.Type(), gniim.LocalName()))
	}
	return gniim.state
}

// Args contains the configurations for google_notebooks_instance_iam_member.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleNotebooksInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("instance_name"))
}

// Location returns a reference to field location of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("location"))
}

// Member returns a reference to field member of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("member"))
}

// Project returns a reference to field project of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("project"))
}

// Role returns a reference to field role of google_notebooks_instance_iam_member.
func (gniim googleNotebooksInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gniim.ref.Append("role"))
}

func (gniim googleNotebooksInstanceIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gniim.ref.Append("condition"))
}

type googleNotebooksInstanceIamMemberState struct {
	Etag         string           `json:"etag"`
	Id           string           `json:"id"`
	InstanceName string           `json:"instance_name"`
	Location     string           `json:"location"`
	Member       string           `json:"member"`
	Project      string           `json:"project"`
	Role         string           `json:"role"`
	Condition    []ConditionState `json:"condition"`
}
