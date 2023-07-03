// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpnsite "github.com/golingon/terraproviders/azurerm/3.63.0/vpnsite"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnSite creates a new instance of [VpnSite].
func NewVpnSite(name string, args VpnSiteArgs) *VpnSite {
	return &VpnSite{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnSite)(nil)

// VpnSite represents the Terraform resource azurerm_vpn_site.
type VpnSite struct {
	Name      string
	Args      VpnSiteArgs
	state     *vpnSiteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnSite].
func (vs *VpnSite) Type() string {
	return "azurerm_vpn_site"
}

// LocalName returns the local name for [VpnSite].
func (vs *VpnSite) LocalName() string {
	return vs.Name
}

// Configuration returns the configuration (args) for [VpnSite].
func (vs *VpnSite) Configuration() interface{} {
	return vs.Args
}

// DependOn is used for other resources to depend on [VpnSite].
func (vs *VpnSite) DependOn() terra.Reference {
	return terra.ReferenceResource(vs)
}

// Dependencies returns the list of resources [VpnSite] depends_on.
func (vs *VpnSite) Dependencies() terra.Dependencies {
	return vs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnSite].
func (vs *VpnSite) LifecycleManagement() *terra.Lifecycle {
	return vs.Lifecycle
}

// Attributes returns the attributes for [VpnSite].
func (vs *VpnSite) Attributes() vpnSiteAttributes {
	return vpnSiteAttributes{ref: terra.ReferenceResource(vs)}
}

// ImportState imports the given attribute values into [VpnSite]'s state.
func (vs *VpnSite) ImportState(av io.Reader) error {
	vs.state = &vpnSiteState{}
	if err := json.NewDecoder(av).Decode(vs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vs.Type(), vs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnSite] has state.
func (vs *VpnSite) State() (*vpnSiteState, bool) {
	return vs.state, vs.state != nil
}

// StateMust returns the state for [VpnSite]. Panics if the state is nil.
func (vs *VpnSite) StateMust() *vpnSiteState {
	if vs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vs.Type(), vs.LocalName()))
	}
	return vs.state
}

// VpnSiteArgs contains the configurations for azurerm_vpn_site.
type VpnSiteArgs struct {
	// AddressCidrs: set of string, optional
	AddressCidrs terra.SetValue[terra.StringValue] `hcl:"address_cidrs,attr"`
	// DeviceModel: string, optional
	DeviceModel terra.StringValue `hcl:"device_model,attr"`
	// DeviceVendor: string, optional
	DeviceVendor terra.StringValue `hcl:"device_vendor,attr"`
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
	// VirtualWanId: string, required
	VirtualWanId terra.StringValue `hcl:"virtual_wan_id,attr" validate:"required"`
	// Link: min=0
	Link []vpnsite.Link `hcl:"link,block" validate:"min=0"`
	// O365Policy: optional
	O365Policy *vpnsite.O365Policy `hcl:"o365_policy,block"`
	// Timeouts: optional
	Timeouts *vpnsite.Timeouts `hcl:"timeouts,block"`
}
type vpnSiteAttributes struct {
	ref terra.Reference
}

// AddressCidrs returns a reference to field address_cidrs of azurerm_vpn_site.
func (vs vpnSiteAttributes) AddressCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vs.ref.Append("address_cidrs"))
}

// DeviceModel returns a reference to field device_model of azurerm_vpn_site.
func (vs vpnSiteAttributes) DeviceModel() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("device_model"))
}

// DeviceVendor returns a reference to field device_vendor of azurerm_vpn_site.
func (vs vpnSiteAttributes) DeviceVendor() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("device_vendor"))
}

// Id returns a reference to field id of azurerm_vpn_site.
func (vs vpnSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_vpn_site.
func (vs vpnSiteAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_vpn_site.
func (vs vpnSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vpn_site.
func (vs vpnSiteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_vpn_site.
func (vs vpnSiteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vs.ref.Append("tags"))
}

// VirtualWanId returns a reference to field virtual_wan_id of azurerm_vpn_site.
func (vs vpnSiteAttributes) VirtualWanId() terra.StringValue {
	return terra.ReferenceAsString(vs.ref.Append("virtual_wan_id"))
}

func (vs vpnSiteAttributes) Link() terra.ListValue[vpnsite.LinkAttributes] {
	return terra.ReferenceAsList[vpnsite.LinkAttributes](vs.ref.Append("link"))
}

func (vs vpnSiteAttributes) O365Policy() terra.ListValue[vpnsite.O365PolicyAttributes] {
	return terra.ReferenceAsList[vpnsite.O365PolicyAttributes](vs.ref.Append("o365_policy"))
}

func (vs vpnSiteAttributes) Timeouts() vpnsite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpnsite.TimeoutsAttributes](vs.ref.Append("timeouts"))
}

type vpnSiteState struct {
	AddressCidrs      []string                  `json:"address_cidrs"`
	DeviceModel       string                    `json:"device_model"`
	DeviceVendor      string                    `json:"device_vendor"`
	Id                string                    `json:"id"`
	Location          string                    `json:"location"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	Tags              map[string]string         `json:"tags"`
	VirtualWanId      string                    `json:"virtual_wan_id"`
	Link              []vpnsite.LinkState       `json:"link"`
	O365Policy        []vpnsite.O365PolicyState `json:"o365_policy"`
	Timeouts          *vpnsite.TimeoutsState    `json:"timeouts"`
}
