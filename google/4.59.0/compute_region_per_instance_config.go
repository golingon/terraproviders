// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregionperinstanceconfig "github.com/golingon/terraproviders/google/4.59.0/computeregionperinstanceconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeRegionPerInstanceConfig(name string, args ComputeRegionPerInstanceConfigArgs) *ComputeRegionPerInstanceConfig {
	return &ComputeRegionPerInstanceConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionPerInstanceConfig)(nil)

type ComputeRegionPerInstanceConfig struct {
	Name  string
	Args  ComputeRegionPerInstanceConfigArgs
	state *computeRegionPerInstanceConfigState
}

func (crpic *ComputeRegionPerInstanceConfig) Type() string {
	return "google_compute_region_per_instance_config"
}

func (crpic *ComputeRegionPerInstanceConfig) LocalName() string {
	return crpic.Name
}

func (crpic *ComputeRegionPerInstanceConfig) Configuration() interface{} {
	return crpic.Args
}

func (crpic *ComputeRegionPerInstanceConfig) Attributes() computeRegionPerInstanceConfigAttributes {
	return computeRegionPerInstanceConfigAttributes{ref: terra.ReferenceResource(crpic)}
}

func (crpic *ComputeRegionPerInstanceConfig) ImportState(av io.Reader) error {
	crpic.state = &computeRegionPerInstanceConfigState{}
	if err := json.NewDecoder(av).Decode(crpic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crpic.Type(), crpic.LocalName(), err)
	}
	return nil
}

func (crpic *ComputeRegionPerInstanceConfig) State() (*computeRegionPerInstanceConfigState, bool) {
	return crpic.state, crpic.state != nil
}

func (crpic *ComputeRegionPerInstanceConfig) StateMust() *computeRegionPerInstanceConfigState {
	if crpic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crpic.Type(), crpic.LocalName()))
	}
	return crpic.state
}

func (crpic *ComputeRegionPerInstanceConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(crpic)
}

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
	// DependsOn contains resources that ComputeRegionPerInstanceConfig depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeRegionPerInstanceConfigAttributes struct {
	ref terra.Reference
}

func (crpic computeRegionPerInstanceConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("id"))
}

func (crpic computeRegionPerInstanceConfigAttributes) MinimalAction() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("minimal_action"))
}

func (crpic computeRegionPerInstanceConfigAttributes) MostDisruptiveAllowedAction() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("most_disruptive_allowed_action"))
}

func (crpic computeRegionPerInstanceConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("name"))
}

func (crpic computeRegionPerInstanceConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("project"))
}

func (crpic computeRegionPerInstanceConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("region"))
}

func (crpic computeRegionPerInstanceConfigAttributes) RegionInstanceGroupManager() terra.StringValue {
	return terra.ReferenceString(crpic.ref.Append("region_instance_group_manager"))
}

func (crpic computeRegionPerInstanceConfigAttributes) RemoveInstanceStateOnDestroy() terra.BoolValue {
	return terra.ReferenceBool(crpic.ref.Append("remove_instance_state_on_destroy"))
}

func (crpic computeRegionPerInstanceConfigAttributes) PreservedState() terra.ListValue[computeregionperinstanceconfig.PreservedStateAttributes] {
	return terra.ReferenceList[computeregionperinstanceconfig.PreservedStateAttributes](crpic.ref.Append("preserved_state"))
}

func (crpic computeRegionPerInstanceConfigAttributes) Timeouts() computeregionperinstanceconfig.TimeoutsAttributes {
	return terra.ReferenceSingle[computeregionperinstanceconfig.TimeoutsAttributes](crpic.ref.Append("timeouts"))
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
