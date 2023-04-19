// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gkehubmembership "github.com/golingon/terraproviders/google/4.62.0/gkehubmembership"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubMembership creates a new instance of [GkeHubMembership].
func NewGkeHubMembership(name string, args GkeHubMembershipArgs) *GkeHubMembership {
	return &GkeHubMembership{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubMembership)(nil)

// GkeHubMembership represents the Terraform resource google_gke_hub_membership.
type GkeHubMembership struct {
	Name      string
	Args      GkeHubMembershipArgs
	state     *gkeHubMembershipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubMembership].
func (ghm *GkeHubMembership) Type() string {
	return "google_gke_hub_membership"
}

// LocalName returns the local name for [GkeHubMembership].
func (ghm *GkeHubMembership) LocalName() string {
	return ghm.Name
}

// Configuration returns the configuration (args) for [GkeHubMembership].
func (ghm *GkeHubMembership) Configuration() interface{} {
	return ghm.Args
}

// DependOn is used for other resources to depend on [GkeHubMembership].
func (ghm *GkeHubMembership) DependOn() terra.Reference {
	return terra.ReferenceResource(ghm)
}

// Dependencies returns the list of resources [GkeHubMembership] depends_on.
func (ghm *GkeHubMembership) Dependencies() terra.Dependencies {
	return ghm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubMembership].
func (ghm *GkeHubMembership) LifecycleManagement() *terra.Lifecycle {
	return ghm.Lifecycle
}

// Attributes returns the attributes for [GkeHubMembership].
func (ghm *GkeHubMembership) Attributes() gkeHubMembershipAttributes {
	return gkeHubMembershipAttributes{ref: terra.ReferenceResource(ghm)}
}

// ImportState imports the given attribute values into [GkeHubMembership]'s state.
func (ghm *GkeHubMembership) ImportState(av io.Reader) error {
	ghm.state = &gkeHubMembershipState{}
	if err := json.NewDecoder(av).Decode(ghm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghm.Type(), ghm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubMembership] has state.
func (ghm *GkeHubMembership) State() (*gkeHubMembershipState, bool) {
	return ghm.state, ghm.state != nil
}

// StateMust returns the state for [GkeHubMembership]. Panics if the state is nil.
func (ghm *GkeHubMembership) StateMust() *gkeHubMembershipState {
	if ghm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghm.Type(), ghm.LocalName()))
	}
	return ghm.state
}

// GkeHubMembershipArgs contains the configurations for google_gke_hub_membership.
type GkeHubMembershipArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MembershipId: string, required
	MembershipId terra.StringValue `hcl:"membership_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Authority: optional
	Authority *gkehubmembership.Authority `hcl:"authority,block"`
	// Endpoint: optional
	Endpoint *gkehubmembership.Endpoint `hcl:"endpoint,block"`
	// Timeouts: optional
	Timeouts *gkehubmembership.Timeouts `hcl:"timeouts,block"`
}
type gkeHubMembershipAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_gke_hub_membership.
func (ghm gkeHubMembershipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghm.ref.Append("id"))
}

// Labels returns a reference to field labels of google_gke_hub_membership.
func (ghm gkeHubMembershipAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ghm.ref.Append("labels"))
}

// MembershipId returns a reference to field membership_id of google_gke_hub_membership.
func (ghm gkeHubMembershipAttributes) MembershipId() terra.StringValue {
	return terra.ReferenceAsString(ghm.ref.Append("membership_id"))
}

// Name returns a reference to field name of google_gke_hub_membership.
func (ghm gkeHubMembershipAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghm.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_hub_membership.
func (ghm gkeHubMembershipAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghm.ref.Append("project"))
}

func (ghm gkeHubMembershipAttributes) Authority() terra.ListValue[gkehubmembership.AuthorityAttributes] {
	return terra.ReferenceAsList[gkehubmembership.AuthorityAttributes](ghm.ref.Append("authority"))
}

func (ghm gkeHubMembershipAttributes) Endpoint() terra.ListValue[gkehubmembership.EndpointAttributes] {
	return terra.ReferenceAsList[gkehubmembership.EndpointAttributes](ghm.ref.Append("endpoint"))
}

func (ghm gkeHubMembershipAttributes) Timeouts() gkehubmembership.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkehubmembership.TimeoutsAttributes](ghm.ref.Append("timeouts"))
}

type gkeHubMembershipState struct {
	Id           string                            `json:"id"`
	Labels       map[string]string                 `json:"labels"`
	MembershipId string                            `json:"membership_id"`
	Name         string                            `json:"name"`
	Project      string                            `json:"project"`
	Authority    []gkehubmembership.AuthorityState `json:"authority"`
	Endpoint     []gkehubmembership.EndpointState  `json:"endpoint"`
	Timeouts     *gkehubmembership.TimeoutsState   `json:"timeouts"`
}
