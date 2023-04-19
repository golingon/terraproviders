// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionperinstanceconfig "github.com/golingon/terraproviders/googlebeta/4.62.0/computeregionperinstanceconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionPerInstanceConfig creates a new instance of [ComputeRegionPerInstanceConfig].
func NewComputeRegionPerInstanceConfig(name string, args ComputeRegionPerInstanceConfigArgs) *ComputeRegionPerInstanceConfig {
	return &ComputeRegionPerInstanceConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionPerInstanceConfig)(nil)

// ComputeRegionPerInstanceConfig represents the Terraform resource google_compute_region_per_instance_config.
type ComputeRegionPerInstanceConfig struct {
	Name      string
	Args      ComputeRegionPerInstanceConfigArgs
	state     *computeRegionPerInstanceConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) Type() string {
	return "google_compute_region_per_instance_config"
}

// LocalName returns the local name for [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) LocalName() string {
	return crpic.Name
}

// Configuration returns the configuration (args) for [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) Configuration() interface{} {
	return crpic.Args
}

// DependOn is used for other resources to depend on [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(crpic)
}

// Dependencies returns the list of resources [ComputeRegionPerInstanceConfig] depends_on.
func (crpic *ComputeRegionPerInstanceConfig) Dependencies() terra.Dependencies {
	return crpic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) LifecycleManagement() *terra.Lifecycle {
	return crpic.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionPerInstanceConfig].
func (crpic *ComputeRegionPerInstanceConfig) Attributes() computeRegionPerInstanceConfigAttributes {
	return computeRegionPerInstanceConfigAttributes{ref: terra.ReferenceResource(crpic)}
}

// ImportState imports the given attribute values into [ComputeRegionPerInstanceConfig]'s state.
func (crpic *ComputeRegionPerInstanceConfig) ImportState(av io.Reader) error {
	crpic.state = &computeRegionPerInstanceConfigState{}
	if err := json.NewDecoder(av).Decode(crpic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crpic.Type(), crpic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionPerInstanceConfig] has state.
func (crpic *ComputeRegionPerInstanceConfig) State() (*computeRegionPerInstanceConfigState, bool) {
	return crpic.state, crpic.state != nil
}

// StateMust returns the state for [ComputeRegionPerInstanceConfig]. Panics if the state is nil.
func (crpic *ComputeRegionPerInstanceConfig) StateMust() *computeRegionPerInstanceConfigState {
	if crpic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crpic.Type(), crpic.LocalName()))
	}
	return crpic.state
}

// ComputeRegionPerInstanceConfigArgs contains the configurations for google_compute_region_per_instance_config.
type ComputeRegionPerInstanceConfigArgs struct {
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
	// RemoveInstanceStateOnDestroy: bool, optional
	RemoveInstanceStateOnDestroy terra.BoolValue `hcl:"remove_instance_state_on_destroy,attr"`
	// PreservedState: optional
	PreservedState *computeregionperinstanceconfig.PreservedState `hcl:"preserved_state,block"`
	// Timeouts: optional
	Timeouts *computeregionperinstanceconfig.Timeouts `hcl:"timeouts,block"`
}
type computeRegionPerInstanceConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("id"))
}

// MinimalAction returns a reference to field minimal_action of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) MinimalAction() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("minimal_action"))
}

// MostDisruptiveAllowedAction returns a reference to field most_disruptive_allowed_action of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) MostDisruptiveAllowedAction() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("most_disruptive_allowed_action"))
}

// Name returns a reference to field name of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("region"))
}

// RegionInstanceGroupManager returns a reference to field region_instance_group_manager of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) RegionInstanceGroupManager() terra.StringValue {
	return terra.ReferenceAsString(crpic.ref.Append("region_instance_group_manager"))
}

// RemoveInstanceStateOnDestroy returns a reference to field remove_instance_state_on_destroy of google_compute_region_per_instance_config.
func (crpic computeRegionPerInstanceConfigAttributes) RemoveInstanceStateOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(crpic.ref.Append("remove_instance_state_on_destroy"))
}

func (crpic computeRegionPerInstanceConfigAttributes) PreservedState() terra.ListValue[computeregionperinstanceconfig.PreservedStateAttributes] {
	return terra.ReferenceAsList[computeregionperinstanceconfig.PreservedStateAttributes](crpic.ref.Append("preserved_state"))
}

func (crpic computeRegionPerInstanceConfigAttributes) Timeouts() computeregionperinstanceconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionperinstanceconfig.TimeoutsAttributes](crpic.ref.Append("timeouts"))
}

type computeRegionPerInstanceConfigState struct {
	Id                           string                                               `json:"id"`
	MinimalAction                string                                               `json:"minimal_action"`
	MostDisruptiveAllowedAction  string                                               `json:"most_disruptive_allowed_action"`
	Name                         string                                               `json:"name"`
	Project                      string                                               `json:"project"`
	Region                       string                                               `json:"region"`
	RegionInstanceGroupManager   string                                               `json:"region_instance_group_manager"`
	RemoveInstanceStateOnDestroy bool                                                 `json:"remove_instance_state_on_destroy"`
	PreservedState               []computeregionperinstanceconfig.PreservedStateState `json:"preserved_state"`
	Timeouts                     *computeregionperinstanceconfig.TimeoutsState        `json:"timeouts"`
}
