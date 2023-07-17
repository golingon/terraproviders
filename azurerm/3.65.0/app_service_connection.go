// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceconnection "github.com/golingon/terraproviders/azurerm/3.65.0/appserviceconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceConnection creates a new instance of [AppServiceConnection].
func NewAppServiceConnection(name string, args AppServiceConnectionArgs) *AppServiceConnection {
	return &AppServiceConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceConnection)(nil)

// AppServiceConnection represents the Terraform resource azurerm_app_service_connection.
type AppServiceConnection struct {
	Name      string
	Args      AppServiceConnectionArgs
	state     *appServiceConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceConnection].
func (asc *AppServiceConnection) Type() string {
	return "azurerm_app_service_connection"
}

// LocalName returns the local name for [AppServiceConnection].
func (asc *AppServiceConnection) LocalName() string {
	return asc.Name
}

// Configuration returns the configuration (args) for [AppServiceConnection].
func (asc *AppServiceConnection) Configuration() interface{} {
	return asc.Args
}

// DependOn is used for other resources to depend on [AppServiceConnection].
func (asc *AppServiceConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(asc)
}

// Dependencies returns the list of resources [AppServiceConnection] depends_on.
func (asc *AppServiceConnection) Dependencies() terra.Dependencies {
	return asc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceConnection].
func (asc *AppServiceConnection) LifecycleManagement() *terra.Lifecycle {
	return asc.Lifecycle
}

// Attributes returns the attributes for [AppServiceConnection].
func (asc *AppServiceConnection) Attributes() appServiceConnectionAttributes {
	return appServiceConnectionAttributes{ref: terra.ReferenceResource(asc)}
}

// ImportState imports the given attribute values into [AppServiceConnection]'s state.
func (asc *AppServiceConnection) ImportState(av io.Reader) error {
	asc.state = &appServiceConnectionState{}
	if err := json.NewDecoder(av).Decode(asc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asc.Type(), asc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceConnection] has state.
func (asc *AppServiceConnection) State() (*appServiceConnectionState, bool) {
	return asc.state, asc.state != nil
}

// StateMust returns the state for [AppServiceConnection]. Panics if the state is nil.
func (asc *AppServiceConnection) StateMust() *appServiceConnectionState {
	if asc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asc.Type(), asc.LocalName()))
	}
	return asc.state
}

// AppServiceConnectionArgs contains the configurations for azurerm_app_service_connection.
type AppServiceConnectionArgs struct {
	// AppServiceId: string, required
	AppServiceId terra.StringValue `hcl:"app_service_id,attr" validate:"required"`
	// ClientType: string, optional
	ClientType terra.StringValue `hcl:"client_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// VnetSolution: string, optional
	VnetSolution terra.StringValue `hcl:"vnet_solution,attr"`
	// Authentication: required
	Authentication *appserviceconnection.Authentication `hcl:"authentication,block" validate:"required"`
	// SecretStore: optional
	SecretStore *appserviceconnection.SecretStore `hcl:"secret_store,block"`
	// Timeouts: optional
	Timeouts *appserviceconnection.Timeouts `hcl:"timeouts,block"`
}
type appServiceConnectionAttributes struct {
	ref terra.Reference
}

// AppServiceId returns a reference to field app_service_id of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("app_service_id"))
}

// ClientType returns a reference to field client_type of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) ClientType() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("client_type"))
}

// Id returns a reference to field id of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("name"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("target_resource_id"))
}

// VnetSolution returns a reference to field vnet_solution of azurerm_app_service_connection.
func (asc appServiceConnectionAttributes) VnetSolution() terra.StringValue {
	return terra.ReferenceAsString(asc.ref.Append("vnet_solution"))
}

func (asc appServiceConnectionAttributes) Authentication() terra.ListValue[appserviceconnection.AuthenticationAttributes] {
	return terra.ReferenceAsList[appserviceconnection.AuthenticationAttributes](asc.ref.Append("authentication"))
}

func (asc appServiceConnectionAttributes) SecretStore() terra.ListValue[appserviceconnection.SecretStoreAttributes] {
	return terra.ReferenceAsList[appserviceconnection.SecretStoreAttributes](asc.ref.Append("secret_store"))
}

func (asc appServiceConnectionAttributes) Timeouts() appserviceconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceconnection.TimeoutsAttributes](asc.ref.Append("timeouts"))
}

type appServiceConnectionState struct {
	AppServiceId     string                                     `json:"app_service_id"`
	ClientType       string                                     `json:"client_type"`
	Id               string                                     `json:"id"`
	Name             string                                     `json:"name"`
	TargetResourceId string                                     `json:"target_resource_id"`
	VnetSolution     string                                     `json:"vnet_solution"`
	Authentication   []appserviceconnection.AuthenticationState `json:"authentication"`
	SecretStore      []appserviceconnection.SecretStoreState    `json:"secret_store"`
	Timeouts         *appserviceconnection.TimeoutsState        `json:"timeouts"`
}
