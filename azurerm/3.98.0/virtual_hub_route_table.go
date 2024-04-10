// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	virtualhubroutetable "github.com/golingon/terraproviders/azurerm/3.98.0/virtualhubroutetable"
	"io"
)

// NewVirtualHubRouteTable creates a new instance of [VirtualHubRouteTable].
func NewVirtualHubRouteTable(name string, args VirtualHubRouteTableArgs) *VirtualHubRouteTable {
	return &VirtualHubRouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualHubRouteTable)(nil)

// VirtualHubRouteTable represents the Terraform resource azurerm_virtual_hub_route_table.
type VirtualHubRouteTable struct {
	Name      string
	Args      VirtualHubRouteTableArgs
	state     *virtualHubRouteTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) Type() string {
	return "azurerm_virtual_hub_route_table"
}

// LocalName returns the local name for [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) LocalName() string {
	return vhrt.Name
}

// Configuration returns the configuration (args) for [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) Configuration() interface{} {
	return vhrt.Args
}

// DependOn is used for other resources to depend on [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) DependOn() terra.Reference {
	return terra.ReferenceResource(vhrt)
}

// Dependencies returns the list of resources [VirtualHubRouteTable] depends_on.
func (vhrt *VirtualHubRouteTable) Dependencies() terra.Dependencies {
	return vhrt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) LifecycleManagement() *terra.Lifecycle {
	return vhrt.Lifecycle
}

// Attributes returns the attributes for [VirtualHubRouteTable].
func (vhrt *VirtualHubRouteTable) Attributes() virtualHubRouteTableAttributes {
	return virtualHubRouteTableAttributes{ref: terra.ReferenceResource(vhrt)}
}

// ImportState imports the given attribute values into [VirtualHubRouteTable]'s state.
func (vhrt *VirtualHubRouteTable) ImportState(av io.Reader) error {
	vhrt.state = &virtualHubRouteTableState{}
	if err := json.NewDecoder(av).Decode(vhrt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vhrt.Type(), vhrt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualHubRouteTable] has state.
func (vhrt *VirtualHubRouteTable) State() (*virtualHubRouteTableState, bool) {
	return vhrt.state, vhrt.state != nil
}

// StateMust returns the state for [VirtualHubRouteTable]. Panics if the state is nil.
func (vhrt *VirtualHubRouteTable) StateMust() *virtualHubRouteTableState {
	if vhrt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vhrt.Type(), vhrt.LocalName()))
	}
	return vhrt.state
}

// VirtualHubRouteTableArgs contains the configurations for azurerm_virtual_hub_route_table.
type VirtualHubRouteTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: set of string, optional
	Labels terra.SetValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// Route: min=0
	Route []virtualhubroutetable.Route `hcl:"route,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *virtualhubroutetable.Timeouts `hcl:"timeouts,block"`
}
type virtualHubRouteTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_hub_route_table.
func (vhrt virtualHubRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("id"))
}

// Labels returns a reference to field labels of azurerm_virtual_hub_route_table.
func (vhrt virtualHubRouteTableAttributes) Labels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vhrt.ref.Append("labels"))
}

// Name returns a reference to field name of azurerm_virtual_hub_route_table.
func (vhrt virtualHubRouteTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("name"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_virtual_hub_route_table.
func (vhrt virtualHubRouteTableAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("virtual_hub_id"))
}

func (vhrt virtualHubRouteTableAttributes) Route() terra.SetValue[virtualhubroutetable.RouteAttributes] {
	return terra.ReferenceAsSet[virtualhubroutetable.RouteAttributes](vhrt.ref.Append("route"))
}

func (vhrt virtualHubRouteTableAttributes) Timeouts() virtualhubroutetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualhubroutetable.TimeoutsAttributes](vhrt.ref.Append("timeouts"))
}

type virtualHubRouteTableState struct {
	Id           string                              `json:"id"`
	Labels       []string                            `json:"labels"`
	Name         string                              `json:"name"`
	VirtualHubId string                              `json:"virtual_hub_id"`
	Route        []virtualhubroutetable.RouteState   `json:"route"`
	Timeouts     *virtualhubroutetable.TimeoutsState `json:"timeouts"`
}
