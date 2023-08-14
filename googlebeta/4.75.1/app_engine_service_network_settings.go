// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	appengineservicenetworksettings "github.com/golingon/terraproviders/googlebeta/4.75.1/appengineservicenetworksettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineServiceNetworkSettings creates a new instance of [AppEngineServiceNetworkSettings].
func NewAppEngineServiceNetworkSettings(name string, args AppEngineServiceNetworkSettingsArgs) *AppEngineServiceNetworkSettings {
	return &AppEngineServiceNetworkSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineServiceNetworkSettings)(nil)

// AppEngineServiceNetworkSettings represents the Terraform resource google_app_engine_service_network_settings.
type AppEngineServiceNetworkSettings struct {
	Name      string
	Args      AppEngineServiceNetworkSettingsArgs
	state     *appEngineServiceNetworkSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) Type() string {
	return "google_app_engine_service_network_settings"
}

// LocalName returns the local name for [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) LocalName() string {
	return aesns.Name
}

// Configuration returns the configuration (args) for [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) Configuration() interface{} {
	return aesns.Args
}

// DependOn is used for other resources to depend on [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(aesns)
}

// Dependencies returns the list of resources [AppEngineServiceNetworkSettings] depends_on.
func (aesns *AppEngineServiceNetworkSettings) Dependencies() terra.Dependencies {
	return aesns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) LifecycleManagement() *terra.Lifecycle {
	return aesns.Lifecycle
}

// Attributes returns the attributes for [AppEngineServiceNetworkSettings].
func (aesns *AppEngineServiceNetworkSettings) Attributes() appEngineServiceNetworkSettingsAttributes {
	return appEngineServiceNetworkSettingsAttributes{ref: terra.ReferenceResource(aesns)}
}

// ImportState imports the given attribute values into [AppEngineServiceNetworkSettings]'s state.
func (aesns *AppEngineServiceNetworkSettings) ImportState(av io.Reader) error {
	aesns.state = &appEngineServiceNetworkSettingsState{}
	if err := json.NewDecoder(av).Decode(aesns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aesns.Type(), aesns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineServiceNetworkSettings] has state.
func (aesns *AppEngineServiceNetworkSettings) State() (*appEngineServiceNetworkSettingsState, bool) {
	return aesns.state, aesns.state != nil
}

// StateMust returns the state for [AppEngineServiceNetworkSettings]. Panics if the state is nil.
func (aesns *AppEngineServiceNetworkSettings) StateMust() *appEngineServiceNetworkSettingsState {
	if aesns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aesns.Type(), aesns.LocalName()))
	}
	return aesns.state
}

// AppEngineServiceNetworkSettingsArgs contains the configurations for google_app_engine_service_network_settings.
type AppEngineServiceNetworkSettingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// NetworkSettings: required
	NetworkSettings *appengineservicenetworksettings.NetworkSettings `hcl:"network_settings,block" validate:"required"`
	// Timeouts: optional
	Timeouts *appengineservicenetworksettings.Timeouts `hcl:"timeouts,block"`
}
type appEngineServiceNetworkSettingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_app_engine_service_network_settings.
func (aesns appEngineServiceNetworkSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aesns.ref.Append("id"))
}

// Project returns a reference to field project of google_app_engine_service_network_settings.
func (aesns appEngineServiceNetworkSettingsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aesns.ref.Append("project"))
}

// Service returns a reference to field service of google_app_engine_service_network_settings.
func (aesns appEngineServiceNetworkSettingsAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(aesns.ref.Append("service"))
}

func (aesns appEngineServiceNetworkSettingsAttributes) NetworkSettings() terra.ListValue[appengineservicenetworksettings.NetworkSettingsAttributes] {
	return terra.ReferenceAsList[appengineservicenetworksettings.NetworkSettingsAttributes](aesns.ref.Append("network_settings"))
}

func (aesns appEngineServiceNetworkSettingsAttributes) Timeouts() appengineservicenetworksettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appengineservicenetworksettings.TimeoutsAttributes](aesns.ref.Append("timeouts"))
}

type appEngineServiceNetworkSettingsState struct {
	Id              string                                                 `json:"id"`
	Project         string                                                 `json:"project"`
	Service         string                                                 `json:"service"`
	NetworkSettings []appengineservicenetworksettings.NetworkSettingsState `json:"network_settings"`
	Timeouts        *appengineservicenetworksettings.TimeoutsState         `json:"timeouts"`
}