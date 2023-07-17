// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webapphybridconnection "github.com/golingon/terraproviders/azurerm/3.65.0/webapphybridconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebAppHybridConnection creates a new instance of [WebAppHybridConnection].
func NewWebAppHybridConnection(name string, args WebAppHybridConnectionArgs) *WebAppHybridConnection {
	return &WebAppHybridConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebAppHybridConnection)(nil)

// WebAppHybridConnection represents the Terraform resource azurerm_web_app_hybrid_connection.
type WebAppHybridConnection struct {
	Name      string
	Args      WebAppHybridConnectionArgs
	state     *webAppHybridConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) Type() string {
	return "azurerm_web_app_hybrid_connection"
}

// LocalName returns the local name for [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) LocalName() string {
	return wahc.Name
}

// Configuration returns the configuration (args) for [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) Configuration() interface{} {
	return wahc.Args
}

// DependOn is used for other resources to depend on [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(wahc)
}

// Dependencies returns the list of resources [WebAppHybridConnection] depends_on.
func (wahc *WebAppHybridConnection) Dependencies() terra.Dependencies {
	return wahc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) LifecycleManagement() *terra.Lifecycle {
	return wahc.Lifecycle
}

// Attributes returns the attributes for [WebAppHybridConnection].
func (wahc *WebAppHybridConnection) Attributes() webAppHybridConnectionAttributes {
	return webAppHybridConnectionAttributes{ref: terra.ReferenceResource(wahc)}
}

// ImportState imports the given attribute values into [WebAppHybridConnection]'s state.
func (wahc *WebAppHybridConnection) ImportState(av io.Reader) error {
	wahc.state = &webAppHybridConnectionState{}
	if err := json.NewDecoder(av).Decode(wahc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wahc.Type(), wahc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebAppHybridConnection] has state.
func (wahc *WebAppHybridConnection) State() (*webAppHybridConnectionState, bool) {
	return wahc.state, wahc.state != nil
}

// StateMust returns the state for [WebAppHybridConnection]. Panics if the state is nil.
func (wahc *WebAppHybridConnection) StateMust() *webAppHybridConnectionState {
	if wahc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wahc.Type(), wahc.LocalName()))
	}
	return wahc.state
}

// WebAppHybridConnectionArgs contains the configurations for azurerm_web_app_hybrid_connection.
type WebAppHybridConnectionArgs struct {
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
	Timeouts *webapphybridconnection.Timeouts `hcl:"timeouts,block"`
}
type webAppHybridConnectionAttributes struct {
	ref terra.Reference
}

// Hostname returns a reference to field hostname of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("namespace_name"))
}

// Port returns a reference to field port of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(wahc.ref.Append("port"))
}

// RelayId returns a reference to field relay_id of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) RelayId() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("relay_id"))
}

// RelayName returns a reference to field relay_name of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) RelayName() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("relay_name"))
}

// SendKeyName returns a reference to field send_key_name of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) SendKeyName() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("send_key_name"))
}

// SendKeyValue returns a reference to field send_key_value of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) SendKeyValue() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("send_key_value"))
}

// ServiceBusNamespace returns a reference to field service_bus_namespace of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) ServiceBusNamespace() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("service_bus_namespace"))
}

// ServiceBusSuffix returns a reference to field service_bus_suffix of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) ServiceBusSuffix() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("service_bus_suffix"))
}

// WebAppId returns a reference to field web_app_id of azurerm_web_app_hybrid_connection.
func (wahc webAppHybridConnectionAttributes) WebAppId() terra.StringValue {
	return terra.ReferenceAsString(wahc.ref.Append("web_app_id"))
}

func (wahc webAppHybridConnectionAttributes) Timeouts() webapphybridconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webapphybridconnection.TimeoutsAttributes](wahc.ref.Append("timeouts"))
}

type webAppHybridConnectionState struct {
	Hostname            string                                `json:"hostname"`
	Id                  string                                `json:"id"`
	NamespaceName       string                                `json:"namespace_name"`
	Port                float64                               `json:"port"`
	RelayId             string                                `json:"relay_id"`
	RelayName           string                                `json:"relay_name"`
	SendKeyName         string                                `json:"send_key_name"`
	SendKeyValue        string                                `json:"send_key_value"`
	ServiceBusNamespace string                                `json:"service_bus_namespace"`
	ServiceBusSuffix    string                                `json:"service_bus_suffix"`
	WebAppId            string                                `json:"web_app_id"`
	Timeouts            *webapphybridconnection.TimeoutsState `json:"timeouts"`
}
