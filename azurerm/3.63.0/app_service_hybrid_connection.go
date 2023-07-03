// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicehybridconnection "github.com/golingon/terraproviders/azurerm/3.63.0/appservicehybridconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceHybridConnection creates a new instance of [AppServiceHybridConnection].
func NewAppServiceHybridConnection(name string, args AppServiceHybridConnectionArgs) *AppServiceHybridConnection {
	return &AppServiceHybridConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceHybridConnection)(nil)

// AppServiceHybridConnection represents the Terraform resource azurerm_app_service_hybrid_connection.
type AppServiceHybridConnection struct {
	Name      string
	Args      AppServiceHybridConnectionArgs
	state     *appServiceHybridConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) Type() string {
	return "azurerm_app_service_hybrid_connection"
}

// LocalName returns the local name for [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) LocalName() string {
	return ashc.Name
}

// Configuration returns the configuration (args) for [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) Configuration() interface{} {
	return ashc.Args
}

// DependOn is used for other resources to depend on [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ashc)
}

// Dependencies returns the list of resources [AppServiceHybridConnection] depends_on.
func (ashc *AppServiceHybridConnection) Dependencies() terra.Dependencies {
	return ashc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) LifecycleManagement() *terra.Lifecycle {
	return ashc.Lifecycle
}

// Attributes returns the attributes for [AppServiceHybridConnection].
func (ashc *AppServiceHybridConnection) Attributes() appServiceHybridConnectionAttributes {
	return appServiceHybridConnectionAttributes{ref: terra.ReferenceResource(ashc)}
}

// ImportState imports the given attribute values into [AppServiceHybridConnection]'s state.
func (ashc *AppServiceHybridConnection) ImportState(av io.Reader) error {
	ashc.state = &appServiceHybridConnectionState{}
	if err := json.NewDecoder(av).Decode(ashc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ashc.Type(), ashc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceHybridConnection] has state.
func (ashc *AppServiceHybridConnection) State() (*appServiceHybridConnectionState, bool) {
	return ashc.state, ashc.state != nil
}

// StateMust returns the state for [AppServiceHybridConnection]. Panics if the state is nil.
func (ashc *AppServiceHybridConnection) StateMust() *appServiceHybridConnectionState {
	if ashc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ashc.Type(), ashc.LocalName()))
	}
	return ashc.state
}

// AppServiceHybridConnectionArgs contains the configurations for azurerm_app_service_hybrid_connection.
type AppServiceHybridConnectionArgs struct {
	// AppServiceName: string, required
	AppServiceName terra.StringValue `hcl:"app_service_name,attr" validate:"required"`
	// Hostname: string, required
	Hostname terra.StringValue `hcl:"hostname,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// RelayId: string, required
	RelayId terra.StringValue `hcl:"relay_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SendKeyName: string, optional
	SendKeyName terra.StringValue `hcl:"send_key_name,attr"`
	// Timeouts: optional
	Timeouts *appservicehybridconnection.Timeouts `hcl:"timeouts,block"`
}
type appServiceHybridConnectionAttributes struct {
	ref terra.Reference
}

// AppServiceName returns a reference to field app_service_name of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) AppServiceName() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("app_service_name"))
}

// Hostname returns a reference to field hostname of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("namespace_name"))
}

// Port returns a reference to field port of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ashc.ref.Append("port"))
}

// RelayId returns a reference to field relay_id of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) RelayId() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("relay_id"))
}

// RelayName returns a reference to field relay_name of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) RelayName() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("relay_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("resource_group_name"))
}

// SendKeyName returns a reference to field send_key_name of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) SendKeyName() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("send_key_name"))
}

// SendKeyValue returns a reference to field send_key_value of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) SendKeyValue() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("send_key_value"))
}

// ServiceBusNamespace returns a reference to field service_bus_namespace of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) ServiceBusNamespace() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("service_bus_namespace"))
}

// ServiceBusSuffix returns a reference to field service_bus_suffix of azurerm_app_service_hybrid_connection.
func (ashc appServiceHybridConnectionAttributes) ServiceBusSuffix() terra.StringValue {
	return terra.ReferenceAsString(ashc.ref.Append("service_bus_suffix"))
}

func (ashc appServiceHybridConnectionAttributes) Timeouts() appservicehybridconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicehybridconnection.TimeoutsAttributes](ashc.ref.Append("timeouts"))
}

type appServiceHybridConnectionState struct {
	AppServiceName      string                                    `json:"app_service_name"`
	Hostname            string                                    `json:"hostname"`
	Id                  string                                    `json:"id"`
	NamespaceName       string                                    `json:"namespace_name"`
	Port                float64                                   `json:"port"`
	RelayId             string                                    `json:"relay_id"`
	RelayName           string                                    `json:"relay_name"`
	ResourceGroupName   string                                    `json:"resource_group_name"`
	SendKeyName         string                                    `json:"send_key_name"`
	SendKeyValue        string                                    `json:"send_key_value"`
	ServiceBusNamespace string                                    `json:"service_bus_namespace"`
	ServiceBusSuffix    string                                    `json:"service_bus_suffix"`
	Timeouts            *appservicehybridconnection.TimeoutsState `json:"timeouts"`
}
