// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeperinstanceconfig "github.com/golingon/terraproviders/googlebeta/4.77.0/computeperinstanceconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputePerInstanceConfig creates a new instance of [ComputePerInstanceConfig].
func NewComputePerInstanceConfig(name string, args ComputePerInstanceConfigArgs) *ComputePerInstanceConfig {
	return &ComputePerInstanceConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputePerInstanceConfig)(nil)

// ComputePerInstanceConfig represents the Terraform resource google_compute_per_instance_config.
type ComputePerInstanceConfig struct {
	Name      string
	Args      ComputePerInstanceConfigArgs
	state     *computePerInstanceConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) Type() string {
	return "google_compute_per_instance_config"
}

// LocalName returns the local name for [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) LocalName() string {
	return cpic.Name
}

// Configuration returns the configuration (args) for [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) Configuration() interface{} {
	return cpic.Args
}

// DependOn is used for other resources to depend on [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cpic)
}

// Dependencies returns the list of resources [ComputePerInstanceConfig] depends_on.
func (cpic *ComputePerInstanceConfig) Dependencies() terra.Dependencies {
	return cpic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) LifecycleManagement() *terra.Lifecycle {
	return cpic.Lifecycle
}

// Attributes returns the attributes for [ComputePerInstanceConfig].
func (cpic *ComputePerInstanceConfig) Attributes() computePerInstanceConfigAttributes {
	return computePerInstanceConfigAttributes{ref: terra.ReferenceResource(cpic)}
}

// ImportState imports the given attribute values into [ComputePerInstanceConfig]'s state.
func (cpic *ComputePerInstanceConfig) ImportState(av io.Reader) error {
	cpic.state = &computePerInstanceConfigState{}
	if err := json.NewDecoder(av).Decode(cpic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpic.Type(), cpic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputePerInstanceConfig] has state.
func (cpic *ComputePerInstanceConfig) State() (*computePerInstanceConfigState, bool) {
	return cpic.state, cpic.state != nil
}

// StateMust returns the state for [ComputePerInstanceConfig]. Panics if the state is nil.
func (cpic *ComputePerInstanceConfig) StateMust() *computePerInstanceConfigState {
	if cpic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpic.Type(), cpic.LocalName()))
	}
	return cpic.state
}

// ComputePerInstanceConfigArgs contains the configurations for google_compute_per_instance_config.
type ComputePerInstanceConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceGroupManager: string, required
	InstanceGroupManager terra.StringValue `hcl:"instance_group_manager,attr" validate:"required"`
	// MinimalAction: string, optional
	MinimalAction terra.StringValue `hcl:"minimal_action,attr"`
	// MostDisruptiveAllowedAction: string, optional
	MostDisruptiveAllowedAction terra.StringValue `hcl:"most_disruptive_allowed_action,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RemoveInstanceStateOnDestroy: bool, optional
	RemoveInstanceStateOnDestroy terra.BoolValue `hcl:"remove_instance_state_on_destroy,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// PreservedState: optional
	PreservedState *computeperinstanceconfig.PreservedState `hcl:"preserved_state,block"`
	// Timeouts: optional
	Timeouts *computeperinstanceconfig.Timeouts `hcl:"timeouts,block"`
}
type computePerInstanceConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("id"))
}

// InstanceGroupManager returns a reference to field instance_group_manager of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) InstanceGroupManager() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("instance_group_manager"))
}

// MinimalAction returns a reference to field minimal_action of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) MinimalAction() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("minimal_action"))
}

// MostDisruptiveAllowedAction returns a reference to field most_disruptive_allowed_action of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) MostDisruptiveAllowedAction() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("most_disruptive_allowed_action"))
}

// Name returns a reference to field name of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("project"))
}

// RemoveInstanceStateOnDestroy returns a reference to field remove_instance_state_on_destroy of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) RemoveInstanceStateOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(cpic.ref.Append("remove_instance_state_on_destroy"))
}

// Zone returns a reference to field zone of google_compute_per_instance_config.
func (cpic computePerInstanceConfigAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cpic.ref.Append("zone"))
}

func (cpic computePerInstanceConfigAttributes) PreservedState() terra.ListValue[computeperinstanceconfig.PreservedStateAttributes] {
	return terra.ReferenceAsList[computeperinstanceconfig.PreservedStateAttributes](cpic.ref.Append("preserved_state"))
}

func (cpic computePerInstanceConfigAttributes) Timeouts() computeperinstanceconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeperinstanceconfig.TimeoutsAttributes](cpic.ref.Append("timeouts"))
}

type computePerInstanceConfigState struct {
	Id                           string                                         `json:"id"`
	InstanceGroupManager         string                                         `json:"instance_group_manager"`
	MinimalAction                string                                         `json:"minimal_action"`
	MostDisruptiveAllowedAction  string                                         `json:"most_disruptive_allowed_action"`
	Name                         string                                         `json:"name"`
	Project                      string                                         `json:"project"`
	RemoveInstanceStateOnDestroy bool                                           `json:"remove_instance_state_on_destroy"`
	Zone                         string                                         `json:"zone"`
	PreservedState               []computeperinstanceconfig.PreservedStateState `json:"preserved_state"`
	Timeouts                     *computeperinstanceconfig.TimeoutsState        `json:"timeouts"`
}
