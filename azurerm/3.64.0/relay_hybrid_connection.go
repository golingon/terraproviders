// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	relayhybridconnection "github.com/golingon/terraproviders/azurerm/3.64.0/relayhybridconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRelayHybridConnection creates a new instance of [RelayHybridConnection].
func NewRelayHybridConnection(name string, args RelayHybridConnectionArgs) *RelayHybridConnection {
	return &RelayHybridConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RelayHybridConnection)(nil)

// RelayHybridConnection represents the Terraform resource azurerm_relay_hybrid_connection.
type RelayHybridConnection struct {
	Name      string
	Args      RelayHybridConnectionArgs
	state     *relayHybridConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RelayHybridConnection].
func (rhc *RelayHybridConnection) Type() string {
	return "azurerm_relay_hybrid_connection"
}

// LocalName returns the local name for [RelayHybridConnection].
func (rhc *RelayHybridConnection) LocalName() string {
	return rhc.Name
}

// Configuration returns the configuration (args) for [RelayHybridConnection].
func (rhc *RelayHybridConnection) Configuration() interface{} {
	return rhc.Args
}

// DependOn is used for other resources to depend on [RelayHybridConnection].
func (rhc *RelayHybridConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(rhc)
}

// Dependencies returns the list of resources [RelayHybridConnection] depends_on.
func (rhc *RelayHybridConnection) Dependencies() terra.Dependencies {
	return rhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RelayHybridConnection].
func (rhc *RelayHybridConnection) LifecycleManagement() *terra.Lifecycle {
	return rhc.Lifecycle
}

// Attributes returns the attributes for [RelayHybridConnection].
func (rhc *RelayHybridConnection) Attributes() relayHybridConnectionAttributes {
	return relayHybridConnectionAttributes{ref: terra.ReferenceResource(rhc)}
}

// ImportState imports the given attribute values into [RelayHybridConnection]'s state.
func (rhc *RelayHybridConnection) ImportState(av io.Reader) error {
	rhc.state = &relayHybridConnectionState{}
	if err := json.NewDecoder(av).Decode(rhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rhc.Type(), rhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RelayHybridConnection] has state.
func (rhc *RelayHybridConnection) State() (*relayHybridConnectionState, bool) {
	return rhc.state, rhc.state != nil
}

// StateMust returns the state for [RelayHybridConnection]. Panics if the state is nil.
func (rhc *RelayHybridConnection) StateMust() *relayHybridConnectionState {
	if rhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rhc.Type(), rhc.LocalName()))
	}
	return rhc.state
}

// RelayHybridConnectionArgs contains the configurations for azurerm_relay_hybrid_connection.
type RelayHybridConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RelayNamespaceName: string, required
	RelayNamespaceName terra.StringValue `hcl:"relay_namespace_name,attr" validate:"required"`
	// RequiresClientAuthorization: bool, optional
	RequiresClientAuthorization terra.BoolValue `hcl:"requires_client_authorization,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UserMetadata: string, optional
	UserMetadata terra.StringValue `hcl:"user_metadata,attr"`
	// Timeouts: optional
	Timeouts *relayhybridconnection.Timeouts `hcl:"timeouts,block"`
}
type relayHybridConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("name"))
}

// RelayNamespaceName returns a reference to field relay_namespace_name of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) RelayNamespaceName() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("relay_namespace_name"))
}

// RequiresClientAuthorization returns a reference to field requires_client_authorization of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) RequiresClientAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(rhc.ref.Append("requires_client_authorization"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("resource_group_name"))
}

// UserMetadata returns a reference to field user_metadata of azurerm_relay_hybrid_connection.
func (rhc relayHybridConnectionAttributes) UserMetadata() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("user_metadata"))
}

func (rhc relayHybridConnectionAttributes) Timeouts() relayhybridconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[relayhybridconnection.TimeoutsAttributes](rhc.ref.Append("timeouts"))
}

type relayHybridConnectionState struct {
	Id                          string                               `json:"id"`
	Name                        string                               `json:"name"`
	RelayNamespaceName          string                               `json:"relay_namespace_name"`
	RequiresClientAuthorization bool                                 `json:"requires_client_authorization"`
	ResourceGroupName           string                               `json:"resource_group_name"`
	UserMetadata                string                               `json:"user_metadata"`
	Timeouts                    *relayhybridconnection.TimeoutsState `json:"timeouts"`
}
