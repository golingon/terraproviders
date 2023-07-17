// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkehubmembershipiambinding "github.com/golingon/terraproviders/googlebeta/4.73.1/gkehubmembershipiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubMembershipIamBinding creates a new instance of [GkeHubMembershipIamBinding].
func NewGkeHubMembershipIamBinding(name string, args GkeHubMembershipIamBindingArgs) *GkeHubMembershipIamBinding {
	return &GkeHubMembershipIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubMembershipIamBinding)(nil)

// GkeHubMembershipIamBinding represents the Terraform resource google_gke_hub_membership_iam_binding.
type GkeHubMembershipIamBinding struct {
	Name      string
	Args      GkeHubMembershipIamBindingArgs
	state     *gkeHubMembershipIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) Type() string {
	return "google_gke_hub_membership_iam_binding"
}

// LocalName returns the local name for [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) LocalName() string {
	return ghmib.Name
}

// Configuration returns the configuration (args) for [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) Configuration() interface{} {
	return ghmib.Args
}

// DependOn is used for other resources to depend on [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ghmib)
}

// Dependencies returns the list of resources [GkeHubMembershipIamBinding] depends_on.
func (ghmib *GkeHubMembershipIamBinding) Dependencies() terra.Dependencies {
	return ghmib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ghmib.Lifecycle
}

// Attributes returns the attributes for [GkeHubMembershipIamBinding].
func (ghmib *GkeHubMembershipIamBinding) Attributes() gkeHubMembershipIamBindingAttributes {
	return gkeHubMembershipIamBindingAttributes{ref: terra.ReferenceResource(ghmib)}
}

// ImportState imports the given attribute values into [GkeHubMembershipIamBinding]'s state.
func (ghmib *GkeHubMembershipIamBinding) ImportState(av io.Reader) error {
	ghmib.state = &gkeHubMembershipIamBindingState{}
	if err := json.NewDecoder(av).Decode(ghmib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghmib.Type(), ghmib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubMembershipIamBinding] has state.
func (ghmib *GkeHubMembershipIamBinding) State() (*gkeHubMembershipIamBindingState, bool) {
	return ghmib.state, ghmib.state != nil
}

// StateMust returns the state for [GkeHubMembershipIamBinding]. Panics if the state is nil.
func (ghmib *GkeHubMembershipIamBinding) StateMust() *gkeHubMembershipIamBindingState {
	if ghmib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghmib.Type(), ghmib.LocalName()))
	}
	return ghmib.state
}

// GkeHubMembershipIamBindingArgs contains the configurations for google_gke_hub_membership_iam_binding.
type GkeHubMembershipIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// MembershipId: string, required
	MembershipId terra.StringValue `hcl:"membership_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *gkehubmembershipiambinding.Condition `hcl:"condition,block"`
}
type gkeHubMembershipIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghmib.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghmib.ref.Append("id"))
}

// Members returns a reference to field members of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ghmib.ref.Append("members"))
}

// MembershipId returns a reference to field membership_id of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) MembershipId() terra.StringValue {
	return terra.ReferenceAsString(ghmib.ref.Append("membership_id"))
}

// Project returns a reference to field project of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghmib.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_membership_iam_binding.
func (ghmib gkeHubMembershipIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghmib.ref.Append("role"))
}

func (ghmib gkeHubMembershipIamBindingAttributes) Condition() terra.ListValue[gkehubmembershipiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[gkehubmembershipiambinding.ConditionAttributes](ghmib.ref.Append("condition"))
}

type gkeHubMembershipIamBindingState struct {
	Etag         string                                      `json:"etag"`
	Id           string                                      `json:"id"`
	Members      []string                                    `json:"members"`
	MembershipId string                                      `json:"membership_id"`
	Project      string                                      `json:"project"`
	Role         string                                      `json:"role"`
	Condition    []gkehubmembershipiambinding.ConditionState `json:"condition"`
}
