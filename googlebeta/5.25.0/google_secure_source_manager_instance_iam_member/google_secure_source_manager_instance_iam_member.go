// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_secure_source_manager_instance_iam_member

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

// Resource represents the Terraform resource google_secure_source_manager_instance_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSecureSourceManagerInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gssmiim *Resource) Type() string {
	return "google_secure_source_manager_instance_iam_member"
}

// LocalName returns the local name for [Resource].
func (gssmiim *Resource) LocalName() string {
	return gssmiim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gssmiim *Resource) Configuration() interface{} {
	return gssmiim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gssmiim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gssmiim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gssmiim *Resource) Dependencies() terra.Dependencies {
	return gssmiim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gssmiim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gssmiim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gssmiim *Resource) Attributes() googleSecureSourceManagerInstanceIamMemberAttributes {
	return googleSecureSourceManagerInstanceIamMemberAttributes{ref: terra.ReferenceResource(gssmiim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gssmiim *Resource) ImportState(state io.Reader) error {
	gssmiim.state = &googleSecureSourceManagerInstanceIamMemberState{}
	if err := json.NewDecoder(state).Decode(gssmiim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gssmiim.Type(), gssmiim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gssmiim *Resource) State() (*googleSecureSourceManagerInstanceIamMemberState, bool) {
	return gssmiim.state, gssmiim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gssmiim *Resource) StateMust() *googleSecureSourceManagerInstanceIamMemberState {
	if gssmiim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gssmiim.Type(), gssmiim.LocalName()))
	}
	return gssmiim.state
}

// Args contains the configurations for google_secure_source_manager_instance_iam_member.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
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

type googleSecureSourceManagerInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("etag"))
}

// Id returns a reference to field id of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("instance_id"))
}

// Location returns a reference to field location of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("location"))
}

// Member returns a reference to field member of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("member"))
}

// Project returns a reference to field project of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("project"))
}

// Role returns a reference to field role of google_secure_source_manager_instance_iam_member.
func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gssmiim.ref.Append("role"))
}

func (gssmiim googleSecureSourceManagerInstanceIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gssmiim.ref.Append("condition"))
}

type googleSecureSourceManagerInstanceIamMemberState struct {
	Etag       string           `json:"etag"`
	Id         string           `json:"id"`
	InstanceId string           `json:"instance_id"`
	Location   string           `json:"location"`
	Member     string           `json:"member"`
	Project    string           `json:"project"`
	Role       string           `json:"role"`
	Condition  []ConditionState `json:"condition"`
}
