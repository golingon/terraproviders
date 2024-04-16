// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_web_app_hybrid_connection

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

// Resource represents the Terraform resource azurerm_web_app_hybrid_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermWebAppHybridConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awahc *Resource) Type() string {
	return "azurerm_web_app_hybrid_connection"
}

// LocalName returns the local name for [Resource].
func (awahc *Resource) LocalName() string {
	return awahc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awahc *Resource) Configuration() interface{} {
	return awahc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awahc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awahc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awahc *Resource) Dependencies() terra.Dependencies {
	return awahc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awahc *Resource) LifecycleManagement() *terra.Lifecycle {
	return awahc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awahc *Resource) Attributes() azurermWebAppHybridConnectionAttributes {
	return azurermWebAppHybridConnectionAttributes{ref: terra.ReferenceResource(awahc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awahc *Resource) ImportState(state io.Reader) error {
	awahc.state = &azurermWebAppHybridConnectionState{}
	if err := json.NewDecoder(state).Decode(awahc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awahc.Type(), awahc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awahc *Resource) State() (*azurermWebAppHybridConnectionState, bool) {
	return awahc.state, awahc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awahc *Resource) StateMust() *azurermWebAppHybridConnectionState {
	if awahc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awahc.Type(), awahc.LocalName()))
	}
	return awahc.state
}

// Args contains the configurations for azurerm_web_app_hybrid_connection.
type Args struct {
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
	// WebAppId: string, required
	WebAppId terra.StringValue `hcl:"web_app_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermWebAppHybridConnectionAttributes struct {
	ref terra.Reference
}

// Hostname returns a reference to field hostname of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("namespace_name"))
}

// Port returns a reference to field port of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(awahc.ref.Append("port"))
}

// RelayId returns a reference to field relay_id of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) RelayId() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("relay_id"))
}

// RelayName returns a reference to field relay_name of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) RelayName() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("relay_name"))
}

// SendKeyName returns a reference to field send_key_name of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) SendKeyName() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("send_key_name"))
}

// SendKeyValue returns a reference to field send_key_value of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) SendKeyValue() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("send_key_value"))
}

// ServiceBusNamespace returns a reference to field service_bus_namespace of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) ServiceBusNamespace() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("service_bus_namespace"))
}

// ServiceBusSuffix returns a reference to field service_bus_suffix of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) ServiceBusSuffix() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("service_bus_suffix"))
}

// WebAppId returns a reference to field web_app_id of azurerm_web_app_hybrid_connection.
func (awahc azurermWebAppHybridConnectionAttributes) WebAppId() terra.StringValue {
	return terra.ReferenceAsString(awahc.ref.Append("web_app_id"))
}

func (awahc azurermWebAppHybridConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](awahc.ref.Append("timeouts"))
}

type azurermWebAppHybridConnectionState struct {
	Hostname            string         `json:"hostname"`
	Id                  string         `json:"id"`
	NamespaceName       string         `json:"namespace_name"`
	Port                float64        `json:"port"`
	RelayId             string         `json:"relay_id"`
	RelayName           string         `json:"relay_name"`
	SendKeyName         string         `json:"send_key_name"`
	SendKeyValue        string         `json:"send_key_value"`
	ServiceBusNamespace string         `json:"service_bus_namespace"`
	ServiceBusSuffix    string         `json:"service_bus_suffix"`
	WebAppId            string         `json:"web_app_id"`
	Timeouts            *TimeoutsState `json:"timeouts"`
}
