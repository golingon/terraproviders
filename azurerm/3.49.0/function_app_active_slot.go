// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	functionappactiveslot "github.com/golingon/terraproviders/azurerm/3.49.0/functionappactiveslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFunctionAppActiveSlot creates a new instance of [FunctionAppActiveSlot].
func NewFunctionAppActiveSlot(name string, args FunctionAppActiveSlotArgs) *FunctionAppActiveSlot {
	return &FunctionAppActiveSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FunctionAppActiveSlot)(nil)

// FunctionAppActiveSlot represents the Terraform resource azurerm_function_app_active_slot.
type FunctionAppActiveSlot struct {
	Name      string
	Args      FunctionAppActiveSlotArgs
	state     *functionAppActiveSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) Type() string {
	return "azurerm_function_app_active_slot"
}

// LocalName returns the local name for [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) LocalName() string {
	return faas.Name
}

// Configuration returns the configuration (args) for [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) Configuration() interface{} {
	return faas.Args
}

// DependOn is used for other resources to depend on [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(faas)
}

// Dependencies returns the list of resources [FunctionAppActiveSlot] depends_on.
func (faas *FunctionAppActiveSlot) Dependencies() terra.Dependencies {
	return faas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) LifecycleManagement() *terra.Lifecycle {
	return faas.Lifecycle
}

// Attributes returns the attributes for [FunctionAppActiveSlot].
func (faas *FunctionAppActiveSlot) Attributes() functionAppActiveSlotAttributes {
	return functionAppActiveSlotAttributes{ref: terra.ReferenceResource(faas)}
}

// ImportState imports the given attribute values into [FunctionAppActiveSlot]'s state.
func (faas *FunctionAppActiveSlot) ImportState(av io.Reader) error {
	faas.state = &functionAppActiveSlotState{}
	if err := json.NewDecoder(av).Decode(faas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", faas.Type(), faas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FunctionAppActiveSlot] has state.
func (faas *FunctionAppActiveSlot) State() (*functionAppActiveSlotState, bool) {
	return faas.state, faas.state != nil
}

// StateMust returns the state for [FunctionAppActiveSlot]. Panics if the state is nil.
func (faas *FunctionAppActiveSlot) StateMust() *functionAppActiveSlotState {
	if faas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", faas.Type(), faas.LocalName()))
	}
	return faas.state
}

// FunctionAppActiveSlotArgs contains the configurations for azurerm_function_app_active_slot.
type FunctionAppActiveSlotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OverwriteNetworkConfig: bool, optional
	OverwriteNetworkConfig terra.BoolValue `hcl:"overwrite_network_config,attr"`
	// SlotId: string, required
	SlotId terra.StringValue `hcl:"slot_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *functionappactiveslot.Timeouts `hcl:"timeouts,block"`
}
type functionAppActiveSlotAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_function_app_active_slot.
func (faas functionAppActiveSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("id"))
}

// LastSuccessfulSwap returns a reference to field last_successful_swap of azurerm_function_app_active_slot.
func (faas functionAppActiveSlotAttributes) LastSuccessfulSwap() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("last_successful_swap"))
}

// OverwriteNetworkConfig returns a reference to field overwrite_network_config of azurerm_function_app_active_slot.
func (faas functionAppActiveSlotAttributes) OverwriteNetworkConfig() terra.BoolValue {
	return terra.ReferenceAsBool(faas.ref.Append("overwrite_network_config"))
}

// SlotId returns a reference to field slot_id of azurerm_function_app_active_slot.
func (faas functionAppActiveSlotAttributes) SlotId() terra.StringValue {
	return terra.ReferenceAsString(faas.ref.Append("slot_id"))
}

func (faas functionAppActiveSlotAttributes) Timeouts() functionappactiveslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[functionappactiveslot.TimeoutsAttributes](faas.ref.Append("timeouts"))
}

type functionAppActiveSlotState struct {
	Id                     string                               `json:"id"`
	LastSuccessfulSwap     string                               `json:"last_successful_swap"`
	OverwriteNetworkConfig bool                                 `json:"overwrite_network_config"`
	SlotId                 string                               `json:"slot_id"`
	Timeouts               *functionappactiveslot.TimeoutsState `json:"timeouts"`
}
