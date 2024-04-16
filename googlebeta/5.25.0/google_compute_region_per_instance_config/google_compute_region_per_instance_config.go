// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_per_instance_config

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

// Resource represents the Terraform resource google_compute_region_per_instance_config.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionPerInstanceConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrpic *Resource) Type() string {
	return "google_compute_region_per_instance_config"
}

// LocalName returns the local name for [Resource].
func (gcrpic *Resource) LocalName() string {
	return gcrpic.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrpic *Resource) Configuration() interface{} {
	return gcrpic.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrpic *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrpic)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrpic *Resource) Dependencies() terra.Dependencies {
	return gcrpic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrpic *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrpic.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrpic *Resource) Attributes() googleComputeRegionPerInstanceConfigAttributes {
	return googleComputeRegionPerInstanceConfigAttributes{ref: terra.ReferenceResource(gcrpic)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrpic *Resource) ImportState(state io.Reader) error {
	gcrpic.state = &googleComputeRegionPerInstanceConfigState{}
	if err := json.NewDecoder(state).Decode(gcrpic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrpic.Type(), gcrpic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrpic *Resource) State() (*googleComputeRegionPerInstanceConfigState, bool) {
	return gcrpic.state, gcrpic.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrpic *Resource) StateMust() *googleComputeRegionPerInstanceConfigState {
	if gcrpic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrpic.Type(), gcrpic.LocalName()))
	}
	return gcrpic.state
}

// Args contains the configurations for google_compute_region_per_instance_config.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinimalAction: string, optional
	MinimalAction terra.StringValue `hcl:"minimal_action,attr"`
	// MostDisruptiveAllowedAction: string, optional
	MostDisruptiveAllowedAction terra.StringValue `hcl:"most_disruptive_allowed_action,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RegionInstanceGroupManager: string, required
	RegionInstanceGroupManager terra.StringValue `hcl:"region_instance_group_manager,attr" validate:"required"`
	// RemoveInstanceOnDestroy: bool, optional
	RemoveInstanceOnDestroy terra.BoolValue `hcl:"remove_instance_on_destroy,attr"`
	// RemoveInstanceStateOnDestroy: bool, optional
	RemoveInstanceStateOnDestroy terra.BoolValue `hcl:"remove_instance_state_on_destroy,attr"`
	// PreservedState: optional
	PreservedState *PreservedState `hcl:"preserved_state,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRegionPerInstanceConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("id"))
}

// MinimalAction returns a reference to field minimal_action of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) MinimalAction() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("minimal_action"))
}

// MostDisruptiveAllowedAction returns a reference to field most_disruptive_allowed_action of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) MostDisruptiveAllowedAction() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("most_disruptive_allowed_action"))
}

// Name returns a reference to field name of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("region"))
}

// RegionInstanceGroupManager returns a reference to field region_instance_group_manager of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) RegionInstanceGroupManager() terra.StringValue {
	return terra.ReferenceAsString(gcrpic.ref.Append("region_instance_group_manager"))
}

// RemoveInstanceOnDestroy returns a reference to field remove_instance_on_destroy of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) RemoveInstanceOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(gcrpic.ref.Append("remove_instance_on_destroy"))
}

// RemoveInstanceStateOnDestroy returns a reference to field remove_instance_state_on_destroy of google_compute_region_per_instance_config.
func (gcrpic googleComputeRegionPerInstanceConfigAttributes) RemoveInstanceStateOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(gcrpic.ref.Append("remove_instance_state_on_destroy"))
}

func (gcrpic googleComputeRegionPerInstanceConfigAttributes) PreservedState() terra.ListValue[PreservedStateAttributes] {
	return terra.ReferenceAsList[PreservedStateAttributes](gcrpic.ref.Append("preserved_state"))
}

func (gcrpic googleComputeRegionPerInstanceConfigAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrpic.ref.Append("timeouts"))
}

type googleComputeRegionPerInstanceConfigState struct {
	Id                           string                `json:"id"`
	MinimalAction                string                `json:"minimal_action"`
	MostDisruptiveAllowedAction  string                `json:"most_disruptive_allowed_action"`
	Name                         string                `json:"name"`
	Project                      string                `json:"project"`
	Region                       string                `json:"region"`
	RegionInstanceGroupManager   string                `json:"region_instance_group_manager"`
	RemoveInstanceOnDestroy      bool                  `json:"remove_instance_on_destroy"`
	RemoveInstanceStateOnDestroy bool                  `json:"remove_instance_state_on_destroy"`
	PreservedState               []PreservedStateState `json:"preserved_state"`
	Timeouts                     *TimeoutsState        `json:"timeouts"`
}
