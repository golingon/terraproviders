// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webappactiveslot "github.com/golingon/terraproviders/azurerm/3.67.0/webappactiveslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebAppActiveSlot creates a new instance of [WebAppActiveSlot].
func NewWebAppActiveSlot(name string, args WebAppActiveSlotArgs) *WebAppActiveSlot {
	return &WebAppActiveSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebAppActiveSlot)(nil)

// WebAppActiveSlot represents the Terraform resource azurerm_web_app_active_slot.
type WebAppActiveSlot struct {
	Name      string
	Args      WebAppActiveSlotArgs
	state     *webAppActiveSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebAppActiveSlot].
func (waas *WebAppActiveSlot) Type() string {
	return "azurerm_web_app_active_slot"
}

// LocalName returns the local name for [WebAppActiveSlot].
func (waas *WebAppActiveSlot) LocalName() string {
	return waas.Name
}

// Configuration returns the configuration (args) for [WebAppActiveSlot].
func (waas *WebAppActiveSlot) Configuration() interface{} {
	return waas.Args
}

// DependOn is used for other resources to depend on [WebAppActiveSlot].
func (waas *WebAppActiveSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(waas)
}

// Dependencies returns the list of resources [WebAppActiveSlot] depends_on.
func (waas *WebAppActiveSlot) Dependencies() terra.Dependencies {
	return waas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebAppActiveSlot].
func (waas *WebAppActiveSlot) LifecycleManagement() *terra.Lifecycle {
	return waas.Lifecycle
}

// Attributes returns the attributes for [WebAppActiveSlot].
func (waas *WebAppActiveSlot) Attributes() webAppActiveSlotAttributes {
	return webAppActiveSlotAttributes{ref: terra.ReferenceResource(waas)}
}

// ImportState imports the given attribute values into [WebAppActiveSlot]'s state.
func (waas *WebAppActiveSlot) ImportState(av io.Reader) error {
	waas.state = &webAppActiveSlotState{}
	if err := json.NewDecoder(av).Decode(waas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", waas.Type(), waas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebAppActiveSlot] has state.
func (waas *WebAppActiveSlot) State() (*webAppActiveSlotState, bool) {
	return waas.state, waas.state != nil
}

// StateMust returns the state for [WebAppActiveSlot]. Panics if the state is nil.
func (waas *WebAppActiveSlot) StateMust() *webAppActiveSlotState {
	if waas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", waas.Type(), waas.LocalName()))
	}
	return waas.state
}

// WebAppActiveSlotArgs contains the configurations for azurerm_web_app_active_slot.
type WebAppActiveSlotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OverwriteNetworkConfig: bool, optional
	OverwriteNetworkConfig terra.BoolValue `hcl:"overwrite_network_config,attr"`
	// SlotId: string, required
	SlotId terra.StringValue `hcl:"slot_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *webappactiveslot.Timeouts `hcl:"timeouts,block"`
}
type webAppActiveSlotAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_web_app_active_slot.
func (waas webAppActiveSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(waas.ref.Append("id"))
}

// LastSuccessfulSwap returns a reference to field last_successful_swap of azurerm_web_app_active_slot.
func (waas webAppActiveSlotAttributes) LastSuccessfulSwap() terra.StringValue {
	return terra.ReferenceAsString(waas.ref.Append("last_successful_swap"))
}

// OverwriteNetworkConfig returns a reference to field overwrite_network_config of azurerm_web_app_active_slot.
func (waas webAppActiveSlotAttributes) OverwriteNetworkConfig() terra.BoolValue {
	return terra.ReferenceAsBool(waas.ref.Append("overwrite_network_config"))
}

// SlotId returns a reference to field slot_id of azurerm_web_app_active_slot.
func (waas webAppActiveSlotAttributes) SlotId() terra.StringValue {
	return terra.ReferenceAsString(waas.ref.Append("slot_id"))
}

func (waas webAppActiveSlotAttributes) Timeouts() webappactiveslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webappactiveslot.TimeoutsAttributes](waas.ref.Append("timeouts"))
}

type webAppActiveSlotState struct {
	Id                     string                          `json:"id"`
	LastSuccessfulSwap     string                          `json:"last_successful_swap"`
	OverwriteNetworkConfig bool                            `json:"overwrite_network_config"`
	SlotId                 string                          `json:"slot_id"`
	Timeouts               *webappactiveslot.TimeoutsState `json:"timeouts"`
}
