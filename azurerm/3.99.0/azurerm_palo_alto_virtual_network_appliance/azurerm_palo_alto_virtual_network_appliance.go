// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_palo_alto_virtual_network_appliance

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

// Resource represents the Terraform resource azurerm_palo_alto_virtual_network_appliance.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPaloAltoVirtualNetworkApplianceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apavna *Resource) Type() string {
	return "azurerm_palo_alto_virtual_network_appliance"
}

// LocalName returns the local name for [Resource].
func (apavna *Resource) LocalName() string {
	return apavna.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apavna *Resource) Configuration() interface{} {
	return apavna.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apavna *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apavna)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apavna *Resource) Dependencies() terra.Dependencies {
	return apavna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apavna *Resource) LifecycleManagement() *terra.Lifecycle {
	return apavna.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apavna *Resource) Attributes() azurermPaloAltoVirtualNetworkApplianceAttributes {
	return azurermPaloAltoVirtualNetworkApplianceAttributes{ref: terra.ReferenceResource(apavna)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apavna *Resource) ImportState(state io.Reader) error {
	apavna.state = &azurermPaloAltoVirtualNetworkApplianceState{}
	if err := json.NewDecoder(state).Decode(apavna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apavna.Type(), apavna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apavna *Resource) State() (*azurermPaloAltoVirtualNetworkApplianceState, bool) {
	return apavna.state, apavna.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apavna *Resource) StateMust() *azurermPaloAltoVirtualNetworkApplianceState {
	if apavna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apavna.Type(), apavna.LocalName()))
	}
	return apavna.state
}

// Args contains the configurations for azurerm_palo_alto_virtual_network_appliance.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPaloAltoVirtualNetworkApplianceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_palo_alto_virtual_network_appliance.
func (apavna azurermPaloAltoVirtualNetworkApplianceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apavna.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_palo_alto_virtual_network_appliance.
func (apavna azurermPaloAltoVirtualNetworkApplianceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apavna.ref.Append("name"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_palo_alto_virtual_network_appliance.
func (apavna azurermPaloAltoVirtualNetworkApplianceAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(apavna.ref.Append("virtual_hub_id"))
}

func (apavna azurermPaloAltoVirtualNetworkApplianceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apavna.ref.Append("timeouts"))
}

type azurermPaloAltoVirtualNetworkApplianceState struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	VirtualHubId string         `json:"virtual_hub_id"`
	Timeouts     *TimeoutsState `json:"timeouts"`
}
