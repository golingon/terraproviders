// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicevirtualnetworkswiftconnection "github.com/golingon/terraproviders/azurerm/3.68.0/appservicevirtualnetworkswiftconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceVirtualNetworkSwiftConnection creates a new instance of [AppServiceVirtualNetworkSwiftConnection].
func NewAppServiceVirtualNetworkSwiftConnection(name string, args AppServiceVirtualNetworkSwiftConnectionArgs) *AppServiceVirtualNetworkSwiftConnection {
	return &AppServiceVirtualNetworkSwiftConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceVirtualNetworkSwiftConnection)(nil)

// AppServiceVirtualNetworkSwiftConnection represents the Terraform resource azurerm_app_service_virtual_network_swift_connection.
type AppServiceVirtualNetworkSwiftConnection struct {
	Name      string
	Args      AppServiceVirtualNetworkSwiftConnectionArgs
	state     *appServiceVirtualNetworkSwiftConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) Type() string {
	return "azurerm_app_service_virtual_network_swift_connection"
}

// LocalName returns the local name for [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) LocalName() string {
	return asvnsc.Name
}

// Configuration returns the configuration (args) for [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) Configuration() interface{} {
	return asvnsc.Args
}

// DependOn is used for other resources to depend on [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(asvnsc)
}

// Dependencies returns the list of resources [AppServiceVirtualNetworkSwiftConnection] depends_on.
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) Dependencies() terra.Dependencies {
	return asvnsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) LifecycleManagement() *terra.Lifecycle {
	return asvnsc.Lifecycle
}

// Attributes returns the attributes for [AppServiceVirtualNetworkSwiftConnection].
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) Attributes() appServiceVirtualNetworkSwiftConnectionAttributes {
	return appServiceVirtualNetworkSwiftConnectionAttributes{ref: terra.ReferenceResource(asvnsc)}
}

// ImportState imports the given attribute values into [AppServiceVirtualNetworkSwiftConnection]'s state.
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) ImportState(av io.Reader) error {
	asvnsc.state = &appServiceVirtualNetworkSwiftConnectionState{}
	if err := json.NewDecoder(av).Decode(asvnsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asvnsc.Type(), asvnsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceVirtualNetworkSwiftConnection] has state.
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) State() (*appServiceVirtualNetworkSwiftConnectionState, bool) {
	return asvnsc.state, asvnsc.state != nil
}

// StateMust returns the state for [AppServiceVirtualNetworkSwiftConnection]. Panics if the state is nil.
func (asvnsc *AppServiceVirtualNetworkSwiftConnection) StateMust() *appServiceVirtualNetworkSwiftConnectionState {
	if asvnsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asvnsc.Type(), asvnsc.LocalName()))
	}
	return asvnsc.state
}

// AppServiceVirtualNetworkSwiftConnectionArgs contains the configurations for azurerm_app_service_virtual_network_swift_connection.
type AppServiceVirtualNetworkSwiftConnectionArgs struct {
	// AppServiceId: string, required
	AppServiceId terra.StringValue `hcl:"app_service_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *appservicevirtualnetworkswiftconnection.Timeouts `hcl:"timeouts,block"`
}
type appServiceVirtualNetworkSwiftConnectionAttributes struct {
	ref terra.Reference
}

// AppServiceId returns a reference to field app_service_id of azurerm_app_service_virtual_network_swift_connection.
func (asvnsc appServiceVirtualNetworkSwiftConnectionAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceAsString(asvnsc.ref.Append("app_service_id"))
}

// Id returns a reference to field id of azurerm_app_service_virtual_network_swift_connection.
func (asvnsc appServiceVirtualNetworkSwiftConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asvnsc.ref.Append("id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_app_service_virtual_network_swift_connection.
func (asvnsc appServiceVirtualNetworkSwiftConnectionAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(asvnsc.ref.Append("subnet_id"))
}

func (asvnsc appServiceVirtualNetworkSwiftConnectionAttributes) Timeouts() appservicevirtualnetworkswiftconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicevirtualnetworkswiftconnection.TimeoutsAttributes](asvnsc.ref.Append("timeouts"))
}

type appServiceVirtualNetworkSwiftConnectionState struct {
	AppServiceId string                                                 `json:"app_service_id"`
	Id           string                                                 `json:"id"`
	SubnetId     string                                                 `json:"subnet_id"`
	Timeouts     *appservicevirtualnetworkswiftconnection.TimeoutsState `json:"timeouts"`
}
