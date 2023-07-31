// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceslotvirtualnetworkswiftconnection "github.com/golingon/terraproviders/azurerm/3.67.0/appserviceslotvirtualnetworkswiftconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSlotVirtualNetworkSwiftConnection creates a new instance of [AppServiceSlotVirtualNetworkSwiftConnection].
func NewAppServiceSlotVirtualNetworkSwiftConnection(name string, args AppServiceSlotVirtualNetworkSwiftConnectionArgs) *AppServiceSlotVirtualNetworkSwiftConnection {
	return &AppServiceSlotVirtualNetworkSwiftConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSlotVirtualNetworkSwiftConnection)(nil)

// AppServiceSlotVirtualNetworkSwiftConnection represents the Terraform resource azurerm_app_service_slot_virtual_network_swift_connection.
type AppServiceSlotVirtualNetworkSwiftConnection struct {
	Name      string
	Args      AppServiceSlotVirtualNetworkSwiftConnectionArgs
	state     *appServiceSlotVirtualNetworkSwiftConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) Type() string {
	return "azurerm_app_service_slot_virtual_network_swift_connection"
}

// LocalName returns the local name for [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) LocalName() string {
	return assvnsc.Name
}

// Configuration returns the configuration (args) for [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) Configuration() interface{} {
	return assvnsc.Args
}

// DependOn is used for other resources to depend on [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(assvnsc)
}

// Dependencies returns the list of resources [AppServiceSlotVirtualNetworkSwiftConnection] depends_on.
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) Dependencies() terra.Dependencies {
	return assvnsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) LifecycleManagement() *terra.Lifecycle {
	return assvnsc.Lifecycle
}

// Attributes returns the attributes for [AppServiceSlotVirtualNetworkSwiftConnection].
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) Attributes() appServiceSlotVirtualNetworkSwiftConnectionAttributes {
	return appServiceSlotVirtualNetworkSwiftConnectionAttributes{ref: terra.ReferenceResource(assvnsc)}
}

// ImportState imports the given attribute values into [AppServiceSlotVirtualNetworkSwiftConnection]'s state.
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) ImportState(av io.Reader) error {
	assvnsc.state = &appServiceSlotVirtualNetworkSwiftConnectionState{}
	if err := json.NewDecoder(av).Decode(assvnsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", assvnsc.Type(), assvnsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSlotVirtualNetworkSwiftConnection] has state.
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) State() (*appServiceSlotVirtualNetworkSwiftConnectionState, bool) {
	return assvnsc.state, assvnsc.state != nil
}

// StateMust returns the state for [AppServiceSlotVirtualNetworkSwiftConnection]. Panics if the state is nil.
func (assvnsc *AppServiceSlotVirtualNetworkSwiftConnection) StateMust() *appServiceSlotVirtualNetworkSwiftConnectionState {
	if assvnsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", assvnsc.Type(), assvnsc.LocalName()))
	}
	return assvnsc.state
}

// AppServiceSlotVirtualNetworkSwiftConnectionArgs contains the configurations for azurerm_app_service_slot_virtual_network_swift_connection.
type AppServiceSlotVirtualNetworkSwiftConnectionArgs struct {
	// AppServiceId: string, required
	AppServiceId terra.StringValue `hcl:"app_service_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SlotName: string, required
	SlotName terra.StringValue `hcl:"slot_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *appserviceslotvirtualnetworkswiftconnection.Timeouts `hcl:"timeouts,block"`
}
type appServiceSlotVirtualNetworkSwiftConnectionAttributes struct {
	ref terra.Reference
}

// AppServiceId returns a reference to field app_service_id of azurerm_app_service_slot_virtual_network_swift_connection.
func (assvnsc appServiceSlotVirtualNetworkSwiftConnectionAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceAsString(assvnsc.ref.Append("app_service_id"))
}

// Id returns a reference to field id of azurerm_app_service_slot_virtual_network_swift_connection.
func (assvnsc appServiceSlotVirtualNetworkSwiftConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assvnsc.ref.Append("id"))
}

// SlotName returns a reference to field slot_name of azurerm_app_service_slot_virtual_network_swift_connection.
func (assvnsc appServiceSlotVirtualNetworkSwiftConnectionAttributes) SlotName() terra.StringValue {
	return terra.ReferenceAsString(assvnsc.ref.Append("slot_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_slot_virtual_network_swift_connection.
func (assvnsc appServiceSlotVirtualNetworkSwiftConnectionAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(assvnsc.ref.Append("subnet_id"))
}

func (assvnsc appServiceSlotVirtualNetworkSwiftConnectionAttributes) Timeouts() appserviceslotvirtualnetworkswiftconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceslotvirtualnetworkswiftconnection.TimeoutsAttributes](assvnsc.ref.Append("timeouts"))
}

type appServiceSlotVirtualNetworkSwiftConnectionState struct {
	AppServiceId string                                                     `json:"app_service_id"`
	Id           string                                                     `json:"id"`
	SlotName     string                                                     `json:"slot_name"`
	SubnetId     string                                                     `json:"subnet_id"`
	Timeouts     *appserviceslotvirtualnetworkswiftconnection.TimeoutsState `json:"timeouts"`
}
