// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualdesktopscalingplan "github.com/golingon/terraproviders/azurerm/3.67.0/virtualdesktopscalingplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualDesktopScalingPlan creates a new instance of [VirtualDesktopScalingPlan].
func NewVirtualDesktopScalingPlan(name string, args VirtualDesktopScalingPlanArgs) *VirtualDesktopScalingPlan {
	return &VirtualDesktopScalingPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopScalingPlan)(nil)

// VirtualDesktopScalingPlan represents the Terraform resource azurerm_virtual_desktop_scaling_plan.
type VirtualDesktopScalingPlan struct {
	Name      string
	Args      VirtualDesktopScalingPlanArgs
	state     *virtualDesktopScalingPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) Type() string {
	return "azurerm_virtual_desktop_scaling_plan"
}

// LocalName returns the local name for [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) LocalName() string {
	return vdsp.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) Configuration() interface{} {
	return vdsp.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(vdsp)
}

// Dependencies returns the list of resources [VirtualDesktopScalingPlan] depends_on.
func (vdsp *VirtualDesktopScalingPlan) Dependencies() terra.Dependencies {
	return vdsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) LifecycleManagement() *terra.Lifecycle {
	return vdsp.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopScalingPlan].
func (vdsp *VirtualDesktopScalingPlan) Attributes() virtualDesktopScalingPlanAttributes {
	return virtualDesktopScalingPlanAttributes{ref: terra.ReferenceResource(vdsp)}
}

// ImportState imports the given attribute values into [VirtualDesktopScalingPlan]'s state.
func (vdsp *VirtualDesktopScalingPlan) ImportState(av io.Reader) error {
	vdsp.state = &virtualDesktopScalingPlanState{}
	if err := json.NewDecoder(av).Decode(vdsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdsp.Type(), vdsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopScalingPlan] has state.
func (vdsp *VirtualDesktopScalingPlan) State() (*virtualDesktopScalingPlanState, bool) {
	return vdsp.state, vdsp.state != nil
}

// StateMust returns the state for [VirtualDesktopScalingPlan]. Panics if the state is nil.
func (vdsp *VirtualDesktopScalingPlan) StateMust() *virtualDesktopScalingPlanState {
	if vdsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdsp.Type(), vdsp.LocalName()))
	}
	return vdsp.state
}

// VirtualDesktopScalingPlanArgs contains the configurations for azurerm_virtual_desktop_scaling_plan.
type VirtualDesktopScalingPlanArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExclusionTag: string, optional
	ExclusionTag terra.StringValue `hcl:"exclusion_tag,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// HostPool: min=0
	HostPool []virtualdesktopscalingplan.HostPool `hcl:"host_pool,block" validate:"min=0"`
	// Schedule: min=1
	Schedule []virtualdesktopscalingplan.Schedule `hcl:"schedule,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *virtualdesktopscalingplan.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopScalingPlanAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("description"))
}

// ExclusionTag returns a reference to field exclusion_tag of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) ExclusionTag() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("exclusion_tag"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("friendly_name"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdsp.ref.Append("tags"))
}

// TimeZone returns a reference to field time_zone of azurerm_virtual_desktop_scaling_plan.
func (vdsp virtualDesktopScalingPlanAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(vdsp.ref.Append("time_zone"))
}

func (vdsp virtualDesktopScalingPlanAttributes) HostPool() terra.ListValue[virtualdesktopscalingplan.HostPoolAttributes] {
	return terra.ReferenceAsList[virtualdesktopscalingplan.HostPoolAttributes](vdsp.ref.Append("host_pool"))
}

func (vdsp virtualDesktopScalingPlanAttributes) Schedule() terra.ListValue[virtualdesktopscalingplan.ScheduleAttributes] {
	return terra.ReferenceAsList[virtualdesktopscalingplan.ScheduleAttributes](vdsp.ref.Append("schedule"))
}

func (vdsp virtualDesktopScalingPlanAttributes) Timeouts() virtualdesktopscalingplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktopscalingplan.TimeoutsAttributes](vdsp.ref.Append("timeouts"))
}

type virtualDesktopScalingPlanState struct {
	Description       string                                    `json:"description"`
	ExclusionTag      string                                    `json:"exclusion_tag"`
	FriendlyName      string                                    `json:"friendly_name"`
	Id                string                                    `json:"id"`
	Location          string                                    `json:"location"`
	Name              string                                    `json:"name"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	Tags              map[string]string                         `json:"tags"`
	TimeZone          string                                    `json:"time_zone"`
	HostPool          []virtualdesktopscalingplan.HostPoolState `json:"host_pool"`
	Schedule          []virtualdesktopscalingplan.ScheduleState `json:"schedule"`
	Timeouts          *virtualdesktopscalingplan.TimeoutsState  `json:"timeouts"`
}
