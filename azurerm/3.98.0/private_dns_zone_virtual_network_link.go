// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	privatednszonevirtualnetworklink "github.com/golingon/terraproviders/azurerm/3.98.0/privatednszonevirtualnetworklink"
	"io"
)

// NewPrivateDnsZoneVirtualNetworkLink creates a new instance of [PrivateDnsZoneVirtualNetworkLink].
func NewPrivateDnsZoneVirtualNetworkLink(name string, args PrivateDnsZoneVirtualNetworkLinkArgs) *PrivateDnsZoneVirtualNetworkLink {
	return &PrivateDnsZoneVirtualNetworkLink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsZoneVirtualNetworkLink)(nil)

// PrivateDnsZoneVirtualNetworkLink represents the Terraform resource azurerm_private_dns_zone_virtual_network_link.
type PrivateDnsZoneVirtualNetworkLink struct {
	Name      string
	Args      PrivateDnsZoneVirtualNetworkLinkArgs
	state     *privateDnsZoneVirtualNetworkLinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) Type() string {
	return "azurerm_private_dns_zone_virtual_network_link"
}

// LocalName returns the local name for [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) LocalName() string {
	return pdzvnl.Name
}

// Configuration returns the configuration (args) for [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) Configuration() interface{} {
	return pdzvnl.Args
}

// DependOn is used for other resources to depend on [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) DependOn() terra.Reference {
	return terra.ReferenceResource(pdzvnl)
}

// Dependencies returns the list of resources [PrivateDnsZoneVirtualNetworkLink] depends_on.
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) Dependencies() terra.Dependencies {
	return pdzvnl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) LifecycleManagement() *terra.Lifecycle {
	return pdzvnl.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) Attributes() privateDnsZoneVirtualNetworkLinkAttributes {
	return privateDnsZoneVirtualNetworkLinkAttributes{ref: terra.ReferenceResource(pdzvnl)}
}

// ImportState imports the given attribute values into [PrivateDnsZoneVirtualNetworkLink]'s state.
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) ImportState(av io.Reader) error {
	pdzvnl.state = &privateDnsZoneVirtualNetworkLinkState{}
	if err := json.NewDecoder(av).Decode(pdzvnl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdzvnl.Type(), pdzvnl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsZoneVirtualNetworkLink] has state.
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) State() (*privateDnsZoneVirtualNetworkLinkState, bool) {
	return pdzvnl.state, pdzvnl.state != nil
}

// StateMust returns the state for [PrivateDnsZoneVirtualNetworkLink]. Panics if the state is nil.
func (pdzvnl *PrivateDnsZoneVirtualNetworkLink) StateMust() *privateDnsZoneVirtualNetworkLinkState {
	if pdzvnl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdzvnl.Type(), pdzvnl.LocalName()))
	}
	return pdzvnl.state
}

// PrivateDnsZoneVirtualNetworkLinkArgs contains the configurations for azurerm_private_dns_zone_virtual_network_link.
type PrivateDnsZoneVirtualNetworkLinkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsZoneName: string, required
	PrivateDnsZoneName terra.StringValue `hcl:"private_dns_zone_name,attr" validate:"required"`
	// RegistrationEnabled: bool, optional
	RegistrationEnabled terra.BoolValue `hcl:"registration_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualNetworkId: string, required
	VirtualNetworkId terra.StringValue `hcl:"virtual_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *privatednszonevirtualnetworklink.Timeouts `hcl:"timeouts,block"`
}
type privateDnsZoneVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("name"))
}

// PrivateDnsZoneName returns a reference to field private_dns_zone_name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) PrivateDnsZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("private_dns_zone_name"))
}

// RegistrationEnabled returns a reference to field registration_enabled of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) RegistrationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pdzvnl.ref.Append("registration_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdzvnl.ref.Append("tags"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("virtual_network_id"))
}

func (pdzvnl privateDnsZoneVirtualNetworkLinkAttributes) Timeouts() privatednszonevirtualnetworklink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednszonevirtualnetworklink.TimeoutsAttributes](pdzvnl.ref.Append("timeouts"))
}

type privateDnsZoneVirtualNetworkLinkState struct {
	Id                  string                                          `json:"id"`
	Name                string                                          `json:"name"`
	PrivateDnsZoneName  string                                          `json:"private_dns_zone_name"`
	RegistrationEnabled bool                                            `json:"registration_enabled"`
	ResourceGroupName   string                                          `json:"resource_group_name"`
	Tags                map[string]string                               `json:"tags"`
	VirtualNetworkId    string                                          `json:"virtual_network_id"`
	Timeouts            *privatednszonevirtualnetworklink.TimeoutsState `json:"timeouts"`
}
