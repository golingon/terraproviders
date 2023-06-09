// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gkehubfeatureiammember "github.com/golingon/terraproviders/google/4.72.1/gkehubfeatureiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubFeatureIamMember creates a new instance of [GkeHubFeatureIamMember].
func NewGkeHubFeatureIamMember(name string, args GkeHubFeatureIamMemberArgs) *GkeHubFeatureIamMember {
	return &GkeHubFeatureIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubFeatureIamMember)(nil)

// GkeHubFeatureIamMember represents the Terraform resource google_gke_hub_feature_iam_member.
type GkeHubFeatureIamMember struct {
	Name      string
	Args      GkeHubFeatureIamMemberArgs
	state     *gkeHubFeatureIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) Type() string {
	return "google_gke_hub_feature_iam_member"
}

// LocalName returns the local name for [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) LocalName() string {
	return ghfim.Name
}

// Configuration returns the configuration (args) for [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) Configuration() interface{} {
	return ghfim.Args
}

// DependOn is used for other resources to depend on [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ghfim)
}

// Dependencies returns the list of resources [GkeHubFeatureIamMember] depends_on.
func (ghfim *GkeHubFeatureIamMember) Dependencies() terra.Dependencies {
	return ghfim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) LifecycleManagement() *terra.Lifecycle {
	return ghfim.Lifecycle
}

// Attributes returns the attributes for [GkeHubFeatureIamMember].
func (ghfim *GkeHubFeatureIamMember) Attributes() gkeHubFeatureIamMemberAttributes {
	return gkeHubFeatureIamMemberAttributes{ref: terra.ReferenceResource(ghfim)}
}

// ImportState imports the given attribute values into [GkeHubFeatureIamMember]'s state.
func (ghfim *GkeHubFeatureIamMember) ImportState(av io.Reader) error {
	ghfim.state = &gkeHubFeatureIamMemberState{}
	if err := json.NewDecoder(av).Decode(ghfim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghfim.Type(), ghfim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubFeatureIamMember] has state.
func (ghfim *GkeHubFeatureIamMember) State() (*gkeHubFeatureIamMemberState, bool) {
	return ghfim.state, ghfim.state != nil
}

// StateMust returns the state for [GkeHubFeatureIamMember]. Panics if the state is nil.
func (ghfim *GkeHubFeatureIamMember) StateMust() *gkeHubFeatureIamMemberState {
	if ghfim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghfim.Type(), ghfim.LocalName()))
	}
	return ghfim.state
}

// GkeHubFeatureIamMemberArgs contains the configurations for google_gke_hub_feature_iam_member.
type GkeHubFeatureIamMemberArgs struct {
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
	Condition *gkehubfeatureiammember.Condition `hcl:"condition,block"`
}
type gkeHubFeatureIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("location"))
}

// Member returns a reference to field member of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("member"))
}

// Name returns a reference to field name of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_feature_iam_member.
func (ghfim gkeHubFeatureIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghfim.ref.Append("role"))
}

func (ghfim gkeHubFeatureIamMemberAttributes) Condition() terra.ListValue[gkehubfeatureiammember.ConditionAttributes] {
	return terra.ReferenceAsList[gkehubfeatureiammember.ConditionAttributes](ghfim.ref.Append("condition"))
}

type gkeHubFeatureIamMemberState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Location  string                                  `json:"location"`
	Member    string                                  `json:"member"`
	Name      string                                  `json:"name"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	Condition []gkehubfeatureiammember.ConditionState `json:"condition"`
}
