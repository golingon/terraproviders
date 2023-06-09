// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkehubfeaturemembership "github.com/golingon/terraproviders/googlebeta/4.63.1/gkehubfeaturemembership"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubFeatureMembership creates a new instance of [GkeHubFeatureMembership].
func NewGkeHubFeatureMembership(name string, args GkeHubFeatureMembershipArgs) *GkeHubFeatureMembership {
	return &GkeHubFeatureMembership{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubFeatureMembership)(nil)

// GkeHubFeatureMembership represents the Terraform resource google_gke_hub_feature_membership.
type GkeHubFeatureMembership struct {
	Name      string
	Args      GkeHubFeatureMembershipArgs
	state     *gkeHubFeatureMembershipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) Type() string {
	return "google_gke_hub_feature_membership"
}

// LocalName returns the local name for [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) LocalName() string {
	return ghfm.Name
}

// Configuration returns the configuration (args) for [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) Configuration() interface{} {
	return ghfm.Args
}

// DependOn is used for other resources to depend on [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) DependOn() terra.Reference {
	return terra.ReferenceResource(ghfm)
}

// Dependencies returns the list of resources [GkeHubFeatureMembership] depends_on.
func (ghfm *GkeHubFeatureMembership) Dependencies() terra.Dependencies {
	return ghfm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) LifecycleManagement() *terra.Lifecycle {
	return ghfm.Lifecycle
}

// Attributes returns the attributes for [GkeHubFeatureMembership].
func (ghfm *GkeHubFeatureMembership) Attributes() gkeHubFeatureMembershipAttributes {
	return gkeHubFeatureMembershipAttributes{ref: terra.ReferenceResource(ghfm)}
}

// ImportState imports the given attribute values into [GkeHubFeatureMembership]'s state.
func (ghfm *GkeHubFeatureMembership) ImportState(av io.Reader) error {
	ghfm.state = &gkeHubFeatureMembershipState{}
	if err := json.NewDecoder(av).Decode(ghfm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghfm.Type(), ghfm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubFeatureMembership] has state.
func (ghfm *GkeHubFeatureMembership) State() (*gkeHubFeatureMembershipState, bool) {
	return ghfm.state, ghfm.state != nil
}

// StateMust returns the state for [GkeHubFeatureMembership]. Panics if the state is nil.
func (ghfm *GkeHubFeatureMembership) StateMust() *gkeHubFeatureMembershipState {
	if ghfm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghfm.Type(), ghfm.LocalName()))
	}
	return ghfm.state
}

// GkeHubFeatureMembershipArgs contains the configurations for google_gke_hub_feature_membership.
type GkeHubFeatureMembershipArgs struct {
	// Feature: string, required
	Feature terra.StringValue `hcl:"feature,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Membership: string, required
	Membership terra.StringValue `hcl:"membership,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Configmanagement: optional
	Configmanagement *gkehubfeaturemembership.Configmanagement `hcl:"configmanagement,block"`
	// Mesh: optional
	Mesh *gkehubfeaturemembership.Mesh `hcl:"mesh,block"`
	// Timeouts: optional
	Timeouts *gkehubfeaturemembership.Timeouts `hcl:"timeouts,block"`
}
type gkeHubFeatureMembershipAttributes struct {
	ref terra.Reference
}

// Feature returns a reference to field feature of google_gke_hub_feature_membership.
func (ghfm gkeHubFeatureMembershipAttributes) Feature() terra.StringValue {
	return terra.ReferenceAsString(ghfm.ref.Append("feature"))
}

// Id returns a reference to field id of google_gke_hub_feature_membership.
func (ghfm gkeHubFeatureMembershipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghfm.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_hub_feature_membership.
func (ghfm gkeHubFeatureMembershipAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ghfm.ref.Append("location"))
}

// Membership returns a reference to field membership of google_gke_hub_feature_membership.
func (ghfm gkeHubFeatureMembershipAttributes) Membership() terra.StringValue {
	return terra.ReferenceAsString(ghfm.ref.Append("membership"))
}

// Project returns a reference to field project of google_gke_hub_feature_membership.
func (ghfm gkeHubFeatureMembershipAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghfm.ref.Append("project"))
}

func (ghfm gkeHubFeatureMembershipAttributes) Configmanagement() terra.ListValue[gkehubfeaturemembership.ConfigmanagementAttributes] {
	return terra.ReferenceAsList[gkehubfeaturemembership.ConfigmanagementAttributes](ghfm.ref.Append("configmanagement"))
}

func (ghfm gkeHubFeatureMembershipAttributes) Mesh() terra.ListValue[gkehubfeaturemembership.MeshAttributes] {
	return terra.ReferenceAsList[gkehubfeaturemembership.MeshAttributes](ghfm.ref.Append("mesh"))
}

func (ghfm gkeHubFeatureMembershipAttributes) Timeouts() gkehubfeaturemembership.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkehubfeaturemembership.TimeoutsAttributes](ghfm.ref.Append("timeouts"))
}

type gkeHubFeatureMembershipState struct {
	Feature          string                                          `json:"feature"`
	Id               string                                          `json:"id"`
	Location         string                                          `json:"location"`
	Membership       string                                          `json:"membership"`
	Project          string                                          `json:"project"`
	Configmanagement []gkehubfeaturemembership.ConfigmanagementState `json:"configmanagement"`
	Mesh             []gkehubfeaturemembership.MeshState             `json:"mesh"`
	Timeouts         *gkehubfeaturemembership.TimeoutsState          `json:"timeouts"`
}
