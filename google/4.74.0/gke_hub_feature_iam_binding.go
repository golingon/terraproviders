// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gkehubfeatureiambinding "github.com/golingon/terraproviders/google/4.74.0/gkehubfeatureiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubFeatureIamBinding creates a new instance of [GkeHubFeatureIamBinding].
func NewGkeHubFeatureIamBinding(name string, args GkeHubFeatureIamBindingArgs) *GkeHubFeatureIamBinding {
	return &GkeHubFeatureIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubFeatureIamBinding)(nil)

// GkeHubFeatureIamBinding represents the Terraform resource google_gke_hub_feature_iam_binding.
type GkeHubFeatureIamBinding struct {
	Name      string
	Args      GkeHubFeatureIamBindingArgs
	state     *gkeHubFeatureIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) Type() string {
	return "google_gke_hub_feature_iam_binding"
}

// LocalName returns the local name for [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) LocalName() string {
	return ghfib.Name
}

// Configuration returns the configuration (args) for [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) Configuration() interface{} {
	return ghfib.Args
}

// DependOn is used for other resources to depend on [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ghfib)
}

// Dependencies returns the list of resources [GkeHubFeatureIamBinding] depends_on.
func (ghfib *GkeHubFeatureIamBinding) Dependencies() terra.Dependencies {
	return ghfib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ghfib.Lifecycle
}

// Attributes returns the attributes for [GkeHubFeatureIamBinding].
func (ghfib *GkeHubFeatureIamBinding) Attributes() gkeHubFeatureIamBindingAttributes {
	return gkeHubFeatureIamBindingAttributes{ref: terra.ReferenceResource(ghfib)}
}

// ImportState imports the given attribute values into [GkeHubFeatureIamBinding]'s state.
func (ghfib *GkeHubFeatureIamBinding) ImportState(av io.Reader) error {
	ghfib.state = &gkeHubFeatureIamBindingState{}
	if err := json.NewDecoder(av).Decode(ghfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghfib.Type(), ghfib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubFeatureIamBinding] has state.
func (ghfib *GkeHubFeatureIamBinding) State() (*gkeHubFeatureIamBindingState, bool) {
	return ghfib.state, ghfib.state != nil
}

// StateMust returns the state for [GkeHubFeatureIamBinding]. Panics if the state is nil.
func (ghfib *GkeHubFeatureIamBinding) StateMust() *gkeHubFeatureIamBindingState {
	if ghfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghfib.Type(), ghfib.LocalName()))
	}
	return ghfib.state
}

// GkeHubFeatureIamBindingArgs contains the configurations for google_gke_hub_feature_iam_binding.
type GkeHubFeatureIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *gkehubfeatureiambinding.Condition `hcl:"condition,block"`
}
type gkeHubFeatureIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("location"))
}

// Members returns a reference to field members of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ghfib.ref.Append("members"))
}

// Name returns a reference to field name of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_feature_iam_binding.
func (ghfib gkeHubFeatureIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ghfib.ref.Append("role"))
}

func (ghfib gkeHubFeatureIamBindingAttributes) Condition() terra.ListValue[gkehubfeatureiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[gkehubfeatureiambinding.ConditionAttributes](ghfib.ref.Append("condition"))
}

type gkeHubFeatureIamBindingState struct {
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Location  string                                   `json:"location"`
	Members   []string                                 `json:"members"`
	Name      string                                   `json:"name"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	Condition []gkehubfeatureiambinding.ConditionState `json:"condition"`
}
