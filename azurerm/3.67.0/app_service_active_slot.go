// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceactiveslot "github.com/golingon/terraproviders/azurerm/3.67.0/appserviceactiveslot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceActiveSlot creates a new instance of [AppServiceActiveSlot].
func NewAppServiceActiveSlot(name string, args AppServiceActiveSlotArgs) *AppServiceActiveSlot {
	return &AppServiceActiveSlot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceActiveSlot)(nil)

// AppServiceActiveSlot represents the Terraform resource azurerm_app_service_active_slot.
type AppServiceActiveSlot struct {
	Name      string
	Args      AppServiceActiveSlotArgs
	state     *appServiceActiveSlotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) Type() string {
	return "azurerm_app_service_active_slot"
}

// LocalName returns the local name for [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) LocalName() string {
	return asas.Name
}

// Configuration returns the configuration (args) for [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) Configuration() interface{} {
	return asas.Args
}

// DependOn is used for other resources to depend on [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) DependOn() terra.Reference {
	return terra.ReferenceResource(asas)
}

// Dependencies returns the list of resources [AppServiceActiveSlot] depends_on.
func (asas *AppServiceActiveSlot) Dependencies() terra.Dependencies {
	return asas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) LifecycleManagement() *terra.Lifecycle {
	return asas.Lifecycle
}

// Attributes returns the attributes for [AppServiceActiveSlot].
func (asas *AppServiceActiveSlot) Attributes() appServiceActiveSlotAttributes {
	return appServiceActiveSlotAttributes{ref: terra.ReferenceResource(asas)}
}

// ImportState imports the given attribute values into [AppServiceActiveSlot]'s state.
func (asas *AppServiceActiveSlot) ImportState(av io.Reader) error {
	asas.state = &appServiceActiveSlotState{}
	if err := json.NewDecoder(av).Decode(asas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asas.Type(), asas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceActiveSlot] has state.
func (asas *AppServiceActiveSlot) State() (*appServiceActiveSlotState, bool) {
	return asas.state, asas.state != nil
}

// StateMust returns the state for [AppServiceActiveSlot]. Panics if the state is nil.
func (asas *AppServiceActiveSlot) StateMust() *appServiceActiveSlotState {
	if asas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asas.Type(), asas.LocalName()))
	}
	return asas.state
}

// AppServiceActiveSlotArgs contains the configurations for azurerm_app_service_active_slot.
type AppServiceActiveSlotArgs struct {
	// AppServiceName: string, required
	AppServiceName terra.StringValue `hcl:"app_service_name,attr" validate:"required"`
	// AppServiceSlotName: string, required
	AppServiceSlotName terra.StringValue `hcl:"app_service_slot_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *appserviceactiveslot.Timeouts `hcl:"timeouts,block"`
}
type appServiceActiveSlotAttributes struct {
	ref terra.Reference
}

// AppServiceName returns a reference to field app_service_name of azurerm_app_service_active_slot.
func (asas appServiceActiveSlotAttributes) AppServiceName() terra.StringValue {
	return terra.ReferenceAsString(asas.ref.Append("app_service_name"))
}

// AppServiceSlotName returns a reference to field app_service_slot_name of azurerm_app_service_active_slot.
func (asas appServiceActiveSlotAttributes) AppServiceSlotName() terra.StringValue {
	return terra.ReferenceAsString(asas.ref.Append("app_service_slot_name"))
}

// Id returns a reference to field id of azurerm_app_service_active_slot.
func (asas appServiceActiveSlotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asas.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_active_slot.
func (asas appServiceActiveSlotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asas.ref.Append("resource_group_name"))
}

func (asas appServiceActiveSlotAttributes) Timeouts() appserviceactiveslot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceactiveslot.TimeoutsAttributes](asas.ref.Append("timeouts"))
}

type appServiceActiveSlotState struct {
	AppServiceName     string                              `json:"app_service_name"`
	AppServiceSlotName string                              `json:"app_service_slot_name"`
	Id                 string                              `json:"id"`
	ResourceGroupName  string                              `json:"resource_group_name"`
	Timeouts           *appserviceactiveslot.TimeoutsState `json:"timeouts"`
}
