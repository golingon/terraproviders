// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkehubfeature "github.com/golingon/terraproviders/googlebeta/4.76.0/gkehubfeature"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeHubFeature creates a new instance of [GkeHubFeature].
func NewGkeHubFeature(name string, args GkeHubFeatureArgs) *GkeHubFeature {
	return &GkeHubFeature{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeHubFeature)(nil)

// GkeHubFeature represents the Terraform resource google_gke_hub_feature.
type GkeHubFeature struct {
	Name      string
	Args      GkeHubFeatureArgs
	state     *gkeHubFeatureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeHubFeature].
func (ghf *GkeHubFeature) Type() string {
	return "google_gke_hub_feature"
}

// LocalName returns the local name for [GkeHubFeature].
func (ghf *GkeHubFeature) LocalName() string {
	return ghf.Name
}

// Configuration returns the configuration (args) for [GkeHubFeature].
func (ghf *GkeHubFeature) Configuration() interface{} {
	return ghf.Args
}

// DependOn is used for other resources to depend on [GkeHubFeature].
func (ghf *GkeHubFeature) DependOn() terra.Reference {
	return terra.ReferenceResource(ghf)
}

// Dependencies returns the list of resources [GkeHubFeature] depends_on.
func (ghf *GkeHubFeature) Dependencies() terra.Dependencies {
	return ghf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeHubFeature].
func (ghf *GkeHubFeature) LifecycleManagement() *terra.Lifecycle {
	return ghf.Lifecycle
}

// Attributes returns the attributes for [GkeHubFeature].
func (ghf *GkeHubFeature) Attributes() gkeHubFeatureAttributes {
	return gkeHubFeatureAttributes{ref: terra.ReferenceResource(ghf)}
}

// ImportState imports the given attribute values into [GkeHubFeature]'s state.
func (ghf *GkeHubFeature) ImportState(av io.Reader) error {
	ghf.state = &gkeHubFeatureState{}
	if err := json.NewDecoder(av).Decode(ghf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghf.Type(), ghf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeHubFeature] has state.
func (ghf *GkeHubFeature) State() (*gkeHubFeatureState, bool) {
	return ghf.state, ghf.state != nil
}

// StateMust returns the state for [GkeHubFeature]. Panics if the state is nil.
func (ghf *GkeHubFeature) StateMust() *gkeHubFeatureState {
	if ghf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghf.Type(), ghf.LocalName()))
	}
	return ghf.state
}

// GkeHubFeatureArgs contains the configurations for google_gke_hub_feature.
type GkeHubFeatureArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResourceState: min=0
	ResourceState []gkehubfeature.ResourceState `hcl:"resource_state,block" validate:"min=0"`
	// State: min=0
	State []gkehubfeature.State `hcl:"state,block" validate:"min=0"`
	// Spec: optional
	Spec *gkehubfeature.Spec `hcl:"spec,block"`
	// Timeouts: optional
	Timeouts *gkehubfeature.Timeouts `hcl:"timeouts,block"`
}
type gkeHubFeatureAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("delete_time"))
}

// Id returns a reference to field id of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ghf.ref.Append("labels"))
}

// Location returns a reference to field location of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("location"))
}

// Name returns a reference to field name of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_gke_hub_feature.
func (ghf gkeHubFeatureAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ghf.ref.Append("update_time"))
}

func (ghf gkeHubFeatureAttributes) ResourceState() terra.ListValue[gkehubfeature.ResourceStateAttributes] {
	return terra.ReferenceAsList[gkehubfeature.ResourceStateAttributes](ghf.ref.Append("resource_state"))
}

func (ghf gkeHubFeatureAttributes) State() terra.ListValue[gkehubfeature.StateAttributes] {
	return terra.ReferenceAsList[gkehubfeature.StateAttributes](ghf.ref.Append("state"))
}

func (ghf gkeHubFeatureAttributes) Spec() terra.ListValue[gkehubfeature.SpecAttributes] {
	return terra.ReferenceAsList[gkehubfeature.SpecAttributes](ghf.ref.Append("spec"))
}

func (ghf gkeHubFeatureAttributes) Timeouts() gkehubfeature.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkehubfeature.TimeoutsAttributes](ghf.ref.Append("timeouts"))
}

type gkeHubFeatureState struct {
	CreateTime    string                             `json:"create_time"`
	DeleteTime    string                             `json:"delete_time"`
	Id            string                             `json:"id"`
	Labels        map[string]string                  `json:"labels"`
	Location      string                             `json:"location"`
	Name          string                             `json:"name"`
	Project       string                             `json:"project"`
	UpdateTime    string                             `json:"update_time"`
	ResourceState []gkehubfeature.ResourceStateState `json:"resource_state"`
	State         []gkehubfeature.StateState         `json:"state"`
	Spec          []gkehubfeature.SpecState          `json:"spec"`
	Timeouts      *gkehubfeature.TimeoutsState       `json:"timeouts"`
}
