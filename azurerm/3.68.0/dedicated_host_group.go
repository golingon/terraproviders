// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dedicatedhostgroup "github.com/golingon/terraproviders/azurerm/3.68.0/dedicatedhostgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDedicatedHostGroup creates a new instance of [DedicatedHostGroup].
func NewDedicatedHostGroup(name string, args DedicatedHostGroupArgs) *DedicatedHostGroup {
	return &DedicatedHostGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DedicatedHostGroup)(nil)

// DedicatedHostGroup represents the Terraform resource azurerm_dedicated_host_group.
type DedicatedHostGroup struct {
	Name      string
	Args      DedicatedHostGroupArgs
	state     *dedicatedHostGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) Type() string {
	return "azurerm_dedicated_host_group"
}

// LocalName returns the local name for [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) LocalName() string {
	return dhg.Name
}

// Configuration returns the configuration (args) for [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) Configuration() interface{} {
	return dhg.Args
}

// DependOn is used for other resources to depend on [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(dhg)
}

// Dependencies returns the list of resources [DedicatedHostGroup] depends_on.
func (dhg *DedicatedHostGroup) Dependencies() terra.Dependencies {
	return dhg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) LifecycleManagement() *terra.Lifecycle {
	return dhg.Lifecycle
}

// Attributes returns the attributes for [DedicatedHostGroup].
func (dhg *DedicatedHostGroup) Attributes() dedicatedHostGroupAttributes {
	return dedicatedHostGroupAttributes{ref: terra.ReferenceResource(dhg)}
}

// ImportState imports the given attribute values into [DedicatedHostGroup]'s state.
func (dhg *DedicatedHostGroup) ImportState(av io.Reader) error {
	dhg.state = &dedicatedHostGroupState{}
	if err := json.NewDecoder(av).Decode(dhg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhg.Type(), dhg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DedicatedHostGroup] has state.
func (dhg *DedicatedHostGroup) State() (*dedicatedHostGroupState, bool) {
	return dhg.state, dhg.state != nil
}

// StateMust returns the state for [DedicatedHostGroup]. Panics if the state is nil.
func (dhg *DedicatedHostGroup) StateMust() *dedicatedHostGroupState {
	if dhg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhg.Type(), dhg.LocalName()))
	}
	return dhg.state
}

// DedicatedHostGroupArgs contains the configurations for azurerm_dedicated_host_group.
type DedicatedHostGroupArgs struct {
	// AutomaticPlacementEnabled: bool, optional
	AutomaticPlacementEnabled terra.BoolValue `hcl:"automatic_placement_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformFaultDomainCount: number, required
	PlatformFaultDomainCount terra.NumberValue `hcl:"platform_fault_domain_count,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *dedicatedhostgroup.Timeouts `hcl:"timeouts,block"`
}
type dedicatedHostGroupAttributes struct {
	ref terra.Reference
}

// AutomaticPlacementEnabled returns a reference to field automatic_placement_enabled of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) AutomaticPlacementEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dhg.ref.Append("automatic_placement_enabled"))
}

// Id returns a reference to field id of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dhg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dhg.ref.Append("name"))
}

// PlatformFaultDomainCount returns a reference to field platform_fault_domain_count of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) PlatformFaultDomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dhg.ref.Append("platform_fault_domain_count"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dhg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dhg.ref.Append("tags"))
}

// Zone returns a reference to field zone of azurerm_dedicated_host_group.
func (dhg dedicatedHostGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(dhg.ref.Append("zone"))
}

func (dhg dedicatedHostGroupAttributes) Timeouts() dedicatedhostgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dedicatedhostgroup.TimeoutsAttributes](dhg.ref.Append("timeouts"))
}

type dedicatedHostGroupState struct {
	AutomaticPlacementEnabled bool                              `json:"automatic_placement_enabled"`
	Id                        string                            `json:"id"`
	Location                  string                            `json:"location"`
	Name                      string                            `json:"name"`
	PlatformFaultDomainCount  float64                           `json:"platform_fault_domain_count"`
	ResourceGroupName         string                            `json:"resource_group_name"`
	Tags                      map[string]string                 `json:"tags"`
	Zone                      string                            `json:"zone"`
	Timeouts                  *dedicatedhostgroup.TimeoutsState `json:"timeouts"`
}
