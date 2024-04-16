// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataplex_task_iam_member

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

// Resource represents the Terraform resource google_dataplex_task_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataplexTaskIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdtim *Resource) Type() string {
	return "google_dataplex_task_iam_member"
}

// LocalName returns the local name for [Resource].
func (gdtim *Resource) LocalName() string {
	return gdtim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdtim *Resource) Configuration() interface{} {
	return gdtim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdtim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdtim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdtim *Resource) Dependencies() terra.Dependencies {
	return gdtim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdtim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdtim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdtim *Resource) Attributes() googleDataplexTaskIamMemberAttributes {
	return googleDataplexTaskIamMemberAttributes{ref: terra.ReferenceResource(gdtim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdtim *Resource) ImportState(state io.Reader) error {
	gdtim.state = &googleDataplexTaskIamMemberState{}
	if err := json.NewDecoder(state).Decode(gdtim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdtim.Type(), gdtim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdtim *Resource) State() (*googleDataplexTaskIamMemberState, bool) {
	return gdtim.state, gdtim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdtim *Resource) StateMust() *googleDataplexTaskIamMemberState {
	if gdtim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdtim.Type(), gdtim.LocalName()))
	}
	return gdtim.state
}

// Args contains the configurations for google_dataplex_task_iam_member.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleDataplexTaskIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("role"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_member.
func (gdtim googleDataplexTaskIamMemberAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(gdtim.ref.Append("task_id"))
}

func (gdtim googleDataplexTaskIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gdtim.ref.Append("condition"))
}

type googleDataplexTaskIamMemberState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Lake      string           `json:"lake"`
	Location  string           `json:"location"`
	Member    string           `json:"member"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	TaskId    string           `json:"task_id"`
	Condition []ConditionState `json:"condition"`
}
