// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_bastion_host

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

// Resource represents the Terraform resource azurerm_bastion_host.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermBastionHostState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (abh *Resource) Type() string {
	return "azurerm_bastion_host"
}

// LocalName returns the local name for [Resource].
func (abh *Resource) LocalName() string {
	return abh.Name
}

// Configuration returns the configuration (args) for [Resource].
func (abh *Resource) Configuration() interface{} {
	return abh.Args
}

// DependOn is used for other resources to depend on [Resource].
func (abh *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(abh)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (abh *Resource) Dependencies() terra.Dependencies {
	return abh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (abh *Resource) LifecycleManagement() *terra.Lifecycle {
	return abh.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (abh *Resource) Attributes() azurermBastionHostAttributes {
	return azurermBastionHostAttributes{ref: terra.ReferenceResource(abh)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (abh *Resource) ImportState(state io.Reader) error {
	abh.state = &azurermBastionHostState{}
	if err := json.NewDecoder(state).Decode(abh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", abh.Type(), abh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (abh *Resource) State() (*azurermBastionHostState, bool) {
	return abh.state, abh.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (abh *Resource) StateMust() *azurermBastionHostState {
	if abh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", abh.Type(), abh.LocalName()))
	}
	return abh.state
}

// Args contains the configurations for azurerm_bastion_host.
type Args struct {
	// CopyPasteEnabled: bool, optional
	CopyPasteEnabled terra.BoolValue `hcl:"copy_paste_enabled,attr"`
	// FileCopyEnabled: bool, optional
	FileCopyEnabled terra.BoolValue `hcl:"file_copy_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpConnectEnabled: bool, optional
	IpConnectEnabled terra.BoolValue `hcl:"ip_connect_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScaleUnits: number, optional
	ScaleUnits terra.NumberValue `hcl:"scale_units,attr"`
	// ShareableLinkEnabled: bool, optional
	ShareableLinkEnabled terra.BoolValue `hcl:"shareable_link_enabled,attr"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TunnelingEnabled: bool, optional
	TunnelingEnabled terra.BoolValue `hcl:"tunneling_enabled,attr"`
	// IpConfiguration: required
	IpConfiguration *IpConfiguration `hcl:"ip_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermBastionHostAttributes struct {
	ref terra.Reference
}

// CopyPasteEnabled returns a reference to field copy_paste_enabled of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) CopyPasteEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("copy_paste_enabled"))
}

// DnsName returns a reference to field dns_name of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("dns_name"))
}

// FileCopyEnabled returns a reference to field file_copy_enabled of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) FileCopyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("file_copy_enabled"))
}

// Id returns a reference to field id of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("id"))
}

// IpConnectEnabled returns a reference to field ip_connect_enabled of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) IpConnectEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("ip_connect_enabled"))
}

// Location returns a reference to field location of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("resource_group_name"))
}

// ScaleUnits returns a reference to field scale_units of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) ScaleUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(abh.ref.Append("scale_units"))
}

// ShareableLinkEnabled returns a reference to field shareable_link_enabled of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) ShareableLinkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("shareable_link_enabled"))
}

// Sku returns a reference to field sku of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(abh.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abh.ref.Append("tags"))
}

// TunnelingEnabled returns a reference to field tunneling_enabled of azurerm_bastion_host.
func (abh azurermBastionHostAttributes) TunnelingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(abh.ref.Append("tunneling_enabled"))
}

func (abh azurermBastionHostAttributes) IpConfiguration() terra.ListValue[IpConfigurationAttributes] {
	return terra.ReferenceAsList[IpConfigurationAttributes](abh.ref.Append("ip_configuration"))
}

func (abh azurermBastionHostAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](abh.ref.Append("timeouts"))
}

type azurermBastionHostState struct {
	CopyPasteEnabled     bool                   `json:"copy_paste_enabled"`
	DnsName              string                 `json:"dns_name"`
	FileCopyEnabled      bool                   `json:"file_copy_enabled"`
	Id                   string                 `json:"id"`
	IpConnectEnabled     bool                   `json:"ip_connect_enabled"`
	Location             string                 `json:"location"`
	Name                 string                 `json:"name"`
	ResourceGroupName    string                 `json:"resource_group_name"`
	ScaleUnits           float64                `json:"scale_units"`
	ShareableLinkEnabled bool                   `json:"shareable_link_enabled"`
	Sku                  string                 `json:"sku"`
	Tags                 map[string]string      `json:"tags"`
	TunnelingEnabled     bool                   `json:"tunneling_enabled"`
	IpConfiguration      []IpConfigurationState `json:"ip_configuration"`
	Timeouts             *TimeoutsState         `json:"timeouts"`
}
