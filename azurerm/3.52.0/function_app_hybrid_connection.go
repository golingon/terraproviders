// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	functionapphybridconnection "github.com/golingon/terraproviders/azurerm/3.52.0/functionapphybridconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFunctionAppHybridConnection creates a new instance of [FunctionAppHybridConnection].
func NewFunctionAppHybridConnection(name string, args FunctionAppHybridConnectionArgs) *FunctionAppHybridConnection {
	return &FunctionAppHybridConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FunctionAppHybridConnection)(nil)

// FunctionAppHybridConnection represents the Terraform resource azurerm_function_app_hybrid_connection.
type FunctionAppHybridConnection struct {
	Name      string
	Args      FunctionAppHybridConnectionArgs
	state     *functionAppHybridConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) Type() string {
	return "azurerm_function_app_hybrid_connection"
}

// LocalName returns the local name for [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) LocalName() string {
	return fahc.Name
}

// Configuration returns the configuration (args) for [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) Configuration() interface{} {
	return fahc.Args
}

// DependOn is used for other resources to depend on [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(fahc)
}

// Dependencies returns the list of resources [FunctionAppHybridConnection] depends_on.
func (fahc *FunctionAppHybridConnection) Dependencies() terra.Dependencies {
	return fahc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) LifecycleManagement() *terra.Lifecycle {
	return fahc.Lifecycle
}

// Attributes returns the attributes for [FunctionAppHybridConnection].
func (fahc *FunctionAppHybridConnection) Attributes() functionAppHybridConnectionAttributes {
	return functionAppHybridConnectionAttributes{ref: terra.ReferenceResource(fahc)}
}

// ImportState imports the given attribute values into [FunctionAppHybridConnection]'s state.
func (fahc *FunctionAppHybridConnection) ImportState(av io.Reader) error {
	fahc.state = &functionAppHybridConnectionState{}
	if err := json.NewDecoder(av).Decode(fahc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fahc.Type(), fahc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FunctionAppHybridConnection] has state.
func (fahc *FunctionAppHybridConnection) State() (*functionAppHybridConnectionState, bool) {
	return fahc.state, fahc.state != nil
}

// StateMust returns the state for [FunctionAppHybridConnection]. Panics if the state is nil.
func (fahc *FunctionAppHybridConnection) StateMust() *functionAppHybridConnectionState {
	if fahc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fahc.Type(), fahc.LocalName()))
	}
	return fahc.state
}

// FunctionAppHybridConnectionArgs contains the configurations for azurerm_function_app_hybrid_connection.
type FunctionAppHybridConnectionArgs struct {
	// FunctionAppId: string, required
	FunctionAppId terra.StringValue `hcl:"function_app_id,attr" validate:"required"`
	// Hostname: string, required
	Hostname terra.StringValue `hcl:"hostname,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// RelayId: string, required
	RelayId terra.StringValue `hcl:"relay_id,attr" validate:"required"`
	// SendKeyName: string, optional
	SendKeyName terra.StringValue `hcl:"send_key_name,attr"`
	// Timeouts: optional
	Timeouts *functionapphybridconnection.Timeouts `hcl:"timeouts,block"`
}
type functionAppHybridConnectionAttributes struct {
	ref terra.Reference
}

// FunctionAppId returns a reference to field function_app_id of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) FunctionAppId() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("function_app_id"))
}

// Hostname returns a reference to field hostname of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("namespace_name"))
}

// Port returns a reference to field port of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(fahc.ref.Append("port"))
}

// RelayId returns a reference to field relay_id of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) RelayId() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("relay_id"))
}

// RelayName returns a reference to field relay_name of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) RelayName() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("relay_name"))
}

// SendKeyName returns a reference to field send_key_name of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) SendKeyName() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("send_key_name"))
}

// SendKeyValue returns a reference to field send_key_value of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) SendKeyValue() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("send_key_value"))
}

// ServiceBusNamespace returns a reference to field service_bus_namespace of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) ServiceBusNamespace() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("service_bus_namespace"))
}

// ServiceBusSuffix returns a reference to field service_bus_suffix of azurerm_function_app_hybrid_connection.
func (fahc functionAppHybridConnectionAttributes) ServiceBusSuffix() terra.StringValue {
	return terra.ReferenceAsString(fahc.ref.Append("service_bus_suffix"))
}

func (fahc functionAppHybridConnectionAttributes) Timeouts() functionapphybridconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[functionapphybridconnection.TimeoutsAttributes](fahc.ref.Append("timeouts"))
}

type functionAppHybridConnectionState struct {
	FunctionAppId       string                                     `json:"function_app_id"`
	Hostname            string                                     `json:"hostname"`
	Id                  string                                     `json:"id"`
	NamespaceName       string                                     `json:"namespace_name"`
	Port                float64                                    `json:"port"`
	RelayId             string                                     `json:"relay_id"`
	RelayName           string                                     `json:"relay_name"`
	SendKeyName         string                                     `json:"send_key_name"`
	SendKeyValue        string                                     `json:"send_key_value"`
	ServiceBusNamespace string                                     `json:"service_bus_namespace"`
	ServiceBusSuffix    string                                     `json:"service_bus_suffix"`
	Timeouts            *functionapphybridconnection.TimeoutsState `json:"timeouts"`
}