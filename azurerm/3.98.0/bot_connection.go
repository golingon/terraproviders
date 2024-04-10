// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	botconnection "github.com/golingon/terraproviders/azurerm/3.98.0/botconnection"
	"io"
)

// NewBotConnection creates a new instance of [BotConnection].
func NewBotConnection(name string, args BotConnectionArgs) *BotConnection {
	return &BotConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotConnection)(nil)

// BotConnection represents the Terraform resource azurerm_bot_connection.
type BotConnection struct {
	Name      string
	Args      BotConnectionArgs
	state     *botConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotConnection].
func (bc *BotConnection) Type() string {
	return "azurerm_bot_connection"
}

// LocalName returns the local name for [BotConnection].
func (bc *BotConnection) LocalName() string {
	return bc.Name
}

// Configuration returns the configuration (args) for [BotConnection].
func (bc *BotConnection) Configuration() interface{} {
	return bc.Args
}

// DependOn is used for other resources to depend on [BotConnection].
func (bc *BotConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(bc)
}

// Dependencies returns the list of resources [BotConnection] depends_on.
func (bc *BotConnection) Dependencies() terra.Dependencies {
	return bc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotConnection].
func (bc *BotConnection) LifecycleManagement() *terra.Lifecycle {
	return bc.Lifecycle
}

// Attributes returns the attributes for [BotConnection].
func (bc *BotConnection) Attributes() botConnectionAttributes {
	return botConnectionAttributes{ref: terra.ReferenceResource(bc)}
}

// ImportState imports the given attribute values into [BotConnection]'s state.
func (bc *BotConnection) ImportState(av io.Reader) error {
	bc.state = &botConnectionState{}
	if err := json.NewDecoder(av).Decode(bc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bc.Type(), bc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotConnection] has state.
func (bc *BotConnection) State() (*botConnectionState, bool) {
	return bc.state, bc.state != nil
}

// StateMust returns the state for [BotConnection]. Panics if the state is nil.
func (bc *BotConnection) StateMust() *botConnectionState {
	if bc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bc.Type(), bc.LocalName()))
	}
	return bc.state
}

// BotConnectionArgs contains the configurations for azurerm_bot_connection.
type BotConnectionArgs struct {
	// BotName: string, required
	BotName terra.StringValue `hcl:"bot_name,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scopes: string, optional
	Scopes terra.StringValue `hcl:"scopes,attr"`
	// ServiceProviderName: string, required
	ServiceProviderName terra.StringValue `hcl:"service_provider_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *botconnection.Timeouts `hcl:"timeouts,block"`
}
type botConnectionAttributes struct {
	ref terra.Reference
}

// BotName returns a reference to field bot_name of azurerm_bot_connection.
func (bc botConnectionAttributes) BotName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("bot_name"))
}

// ClientId returns a reference to field client_id of azurerm_bot_connection.
func (bc botConnectionAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_bot_connection.
func (bc botConnectionAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_bot_connection.
func (bc botConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_connection.
func (bc botConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_bot_connection.
func (bc botConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_bot_connection.
func (bc botConnectionAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bc.ref.Append("parameters"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_connection.
func (bc botConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_bot_connection.
func (bc botConnectionAttributes) Scopes() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("scopes"))
}

// ServiceProviderName returns a reference to field service_provider_name of azurerm_bot_connection.
func (bc botConnectionAttributes) ServiceProviderName() terra.StringValue {
	return terra.ReferenceAsString(bc.ref.Append("service_provider_name"))
}

// Tags returns a reference to field tags of azurerm_bot_connection.
func (bc botConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bc.ref.Append("tags"))
}

func (bc botConnectionAttributes) Timeouts() botconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botconnection.TimeoutsAttributes](bc.ref.Append("timeouts"))
}

type botConnectionState struct {
	BotName             string                       `json:"bot_name"`
	ClientId            string                       `json:"client_id"`
	ClientSecret        string                       `json:"client_secret"`
	Id                  string                       `json:"id"`
	Location            string                       `json:"location"`
	Name                string                       `json:"name"`
	Parameters          map[string]string            `json:"parameters"`
	ResourceGroupName   string                       `json:"resource_group_name"`
	Scopes              string                       `json:"scopes"`
	ServiceProviderName string                       `json:"service_provider_name"`
	Tags                map[string]string            `json:"tags"`
	Timeouts            *botconnection.TimeoutsState `json:"timeouts"`
}
