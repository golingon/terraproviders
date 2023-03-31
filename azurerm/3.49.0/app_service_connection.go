// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceconnection "github.com/golingon/terraproviders/azurerm/3.49.0/appserviceconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAppServiceConnection(name string, args AppServiceConnectionArgs) *AppServiceConnection {
	return &AppServiceConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceConnection)(nil)

type AppServiceConnection struct {
	Name  string
	Args  AppServiceConnectionArgs
	state *appServiceConnectionState
}

func (asc *AppServiceConnection) Type() string {
	return "azurerm_app_service_connection"
}

func (asc *AppServiceConnection) LocalName() string {
	return asc.Name
}

func (asc *AppServiceConnection) Configuration() interface{} {
	return asc.Args
}

func (asc *AppServiceConnection) Attributes() appServiceConnectionAttributes {
	return appServiceConnectionAttributes{ref: terra.ReferenceResource(asc)}
}

func (asc *AppServiceConnection) ImportState(av io.Reader) error {
	asc.state = &appServiceConnectionState{}
	if err := json.NewDecoder(av).Decode(asc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asc.Type(), asc.LocalName(), err)
	}
	return nil
}

func (asc *AppServiceConnection) State() (*appServiceConnectionState, bool) {
	return asc.state, asc.state != nil
}

func (asc *AppServiceConnection) StateMust() *appServiceConnectionState {
	if asc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asc.Type(), asc.LocalName()))
	}
	return asc.state
}

func (asc *AppServiceConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(asc)
}

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
	// DependsOn contains resources that AppServiceConnection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type appServiceConnectionAttributes struct {
	ref terra.Reference
}

func (asc appServiceConnectionAttributes) AppServiceId() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("app_service_id"))
}

func (asc appServiceConnectionAttributes) ClientType() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("client_type"))
}

func (asc appServiceConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("id"))
}

func (asc appServiceConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("name"))
}

func (asc appServiceConnectionAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("target_resource_id"))
}

func (asc appServiceConnectionAttributes) VnetSolution() terra.StringValue {
	return terra.ReferenceString(asc.ref.Append("vnet_solution"))
}

func (asc appServiceConnectionAttributes) Authentication() terra.ListValue[appserviceconnection.AuthenticationAttributes] {
	return terra.ReferenceList[appserviceconnection.AuthenticationAttributes](asc.ref.Append("authentication"))
}

func (asc appServiceConnectionAttributes) SecretStore() terra.ListValue[appserviceconnection.SecretStoreAttributes] {
	return terra.ReferenceList[appserviceconnection.SecretStoreAttributes](asc.ref.Append("secret_store"))
}

func (asc appServiceConnectionAttributes) Timeouts() appserviceconnection.TimeoutsAttributes {
	return terra.ReferenceSingle[appserviceconnection.TimeoutsAttributes](asc.ref.Append("timeouts"))
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
