// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gkehubmembershipiammember "github.com/golingon/terraproviders/google/4.71.0/gkehubmembershipiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubMembershipIamMember creates a new instance of [GkeHubMembershipIamMember].
func NewGkeHubMembershipIamMember(name string, args GkeHubMembershipIamMemberArgs) *GkeHubMembershipIamMember {
	return &GkeHubMembershipIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubMembershipIamMember)(nil)

// GkeHubMembershipIamMember represents the Terraform resource google_gke_hub_membership_iam_member.
type GkeHubMembershipIamMember struct {
	Name      string
	Args      GkeHubMembershipIamMemberArgs
	state     *gkeHubMembershipIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) Type() string {
	return "google_gke_hub_membership_iam_member"
}

// LocalName returns the local name for [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) LocalName() string {
	return ghmim.Name
}

// Configuration returns the configuration (args) for [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) Configuration() interface{} {
	return ghmim.Args
}

// DependOn is used for other resources to depend on [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ghmim)
}

// Dependencies returns the list of resources [GkeHubMembershipIamMember] depends_on.
func (ghmim *GkeHubMembershipIamMember) Dependencies() terra.Dependencies {
	return ghmim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) LifecycleManagement() *terra.Lifecycle {
	return ghmim.Lifecycle
}

// Attributes returns the attributes for [GkeHubMembershipIamMember].
func (ghmim *GkeHubMembershipIamMember) Attributes() gkeHubMembershipIamMemberAttributes {
	return gkeHubMembershipIamMemberAttributes{ref: terra.ReferenceResource(ghmim)}
}

// ImportState imports the given attribute values into [GkeHubMembershipIamMember]'s state.
func (ghmim *GkeHubMembershipIamMember) ImportState(av io.Reader) error {
	ghmim.state = &gkeHubMembershipIamMemberState{}
	if err := json.NewDecoder(av).Decode(ghmim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghmim.Type(), ghmim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubMembershipIamMember] has state.
func (ghmim *GkeHubMembershipIamMember) State() (*gkeHubMembershipIamMemberState, bool) {
	return ghmim.state, ghmim.state != nil
}

// StateMust returns the state for [GkeHubMembershipIamMember]. Panics if the state is nil.
func (ghmim *GkeHubMembershipIamMember) StateMust() *gkeHubMembershipIamMemberState {
	if ghmim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghmim.Type(), ghmim.LocalName()))
	}
	return ghmim.state
}

// GkeHubMembershipIamMemberArgs contains the configurations for google_gke_hub_membership_iam_member.
type GkeHubMembershipIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// MembershipId: string, required
	MembershipId terra.StringValue `hcl:"membership_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *gkehubmembershipiammember.Condition `hcl:"condition,block"`
}
type gkeHubMembershipIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("id"))
}

// Member returns a reference to field member of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("member"))
}

// MembershipId returns a reference to field membership_id of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) MembershipId() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("membership_id"))
}

// Project returns a reference to field project of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_membership_iam_member.
func (ghmim gkeHubMembershipIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghmim.ref.Append("role"))
}

func (ghmim gkeHubMembershipIamMemberAttributes) Condition() terra.ListValue[gkehubmembershipiammember.ConditionAttributes] {
	return terra.ReferenceAsList[gkehubmembershipiammember.ConditionAttributes](ghmim.ref.Append("condition"))
}

type gkeHubMembershipIamMemberState struct {
	Etag         string                                     `json:"etag"`
	Id           string                                     `json:"id"`
	Member       string                                     `json:"member"`
	MembershipId string                                     `json:"membership_id"`
	Project      string                                     `json:"project"`
	Role         string                                     `json:"role"`
	Condition    []gkehubmembershipiammember.ConditionState `json:"condition"`
}
