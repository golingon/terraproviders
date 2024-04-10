// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	gkehubscopeiammember "github.com/golingon/terraproviders/googlebeta/5.24.0/gkehubscopeiammember"
	"io"
)

// NewGkeHubScopeIamMember creates a new instance of [GkeHubScopeIamMember].
func NewGkeHubScopeIamMember(name string, args GkeHubScopeIamMemberArgs) *GkeHubScopeIamMember {
	return &GkeHubScopeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubScopeIamMember)(nil)

// GkeHubScopeIamMember represents the Terraform resource google_gke_hub_scope_iam_member.
type GkeHubScopeIamMember struct {
	Name      string
	Args      GkeHubScopeIamMemberArgs
	state     *gkeHubScopeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) Type() string {
	return "google_gke_hub_scope_iam_member"
}

// LocalName returns the local name for [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) LocalName() string {
	return ghsim.Name
}

// Configuration returns the configuration (args) for [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) Configuration() interface{} {
	return ghsim.Args
}

// DependOn is used for other resources to depend on [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ghsim)
}

// Dependencies returns the list of resources [GkeHubScopeIamMember] depends_on.
func (ghsim *GkeHubScopeIamMember) Dependencies() terra.Dependencies {
	return ghsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) LifecycleManagement() *terra.Lifecycle {
	return ghsim.Lifecycle
}

// Attributes returns the attributes for [GkeHubScopeIamMember].
func (ghsim *GkeHubScopeIamMember) Attributes() gkeHubScopeIamMemberAttributes {
	return gkeHubScopeIamMemberAttributes{ref: terra.ReferenceResource(ghsim)}
}

// ImportState imports the given attribute values into [GkeHubScopeIamMember]'s state.
func (ghsim *GkeHubScopeIamMember) ImportState(av io.Reader) error {
	ghsim.state = &gkeHubScopeIamMemberState{}
	if err := json.NewDecoder(av).Decode(ghsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghsim.Type(), ghsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubScopeIamMember] has state.
func (ghsim *GkeHubScopeIamMember) State() (*gkeHubScopeIamMemberState, bool) {
	return ghsim.state, ghsim.state != nil
}

// StateMust returns the state for [GkeHubScopeIamMember]. Panics if the state is nil.
func (ghsim *GkeHubScopeIamMember) StateMust() *gkeHubScopeIamMemberState {
	if ghsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghsim.Type(), ghsim.LocalName()))
	}
	return ghsim.state
}

// GkeHubScopeIamMemberArgs contains the configurations for google_gke_hub_scope_iam_member.
type GkeHubScopeIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ScopeId: string, required
	ScopeId terra.StringValue `hcl:"scope_id,attr" validate:"required"`
	// Condition: optional
	Condition *gkehubscopeiammember.Condition `hcl:"condition,block"`
}
type gkeHubScopeIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("id"))
}

// Member returns a reference to field member of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("member"))
}

// Project returns a reference to field project of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("role"))
}

// ScopeId returns a reference to field scope_id of google_gke_hub_scope_iam_member.
func (ghsim gkeHubScopeIamMemberAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceAsString(ghsim.ref.Append("scope_id"))
}

func (ghsim gkeHubScopeIamMemberAttributes) Condition() terra.ListValue[gkehubscopeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[gkehubscopeiammember.ConditionAttributes](ghsim.ref.Append("condition"))
}

type gkeHubScopeIamMemberState struct {
	Etag      string                                `json:"etag"`
	Id        string                                `json:"id"`
	Member    string                                `json:"member"`
	Project   string                                `json:"project"`
	Role      string                                `json:"role"`
	ScopeId   string                                `json:"scope_id"`
	Condition []gkehubscopeiammember.ConditionState `json:"condition"`
}
