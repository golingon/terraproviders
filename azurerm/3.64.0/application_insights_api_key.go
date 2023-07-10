// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsightsapikey "github.com/golingon/terraproviders/azurerm/3.64.0/applicationinsightsapikey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsightsApiKey creates a new instance of [ApplicationInsightsApiKey].
func NewApplicationInsightsApiKey(name string, args ApplicationInsightsApiKeyArgs) *ApplicationInsightsApiKey {
	return &ApplicationInsightsApiKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsightsApiKey)(nil)

// ApplicationInsightsApiKey represents the Terraform resource azurerm_application_insights_api_key.
type ApplicationInsightsApiKey struct {
	Name      string
	Args      ApplicationInsightsApiKeyArgs
	state     *applicationInsightsApiKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) Type() string {
	return "azurerm_application_insights_api_key"
}

// LocalName returns the local name for [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) LocalName() string {
	return aiak.Name
}

// Configuration returns the configuration (args) for [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) Configuration() interface{} {
	return aiak.Args
}

// DependOn is used for other resources to depend on [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) DependOn() terra.Reference {
	return terra.ReferenceResource(aiak)
}

// Dependencies returns the list of resources [ApplicationInsightsApiKey] depends_on.
func (aiak *ApplicationInsightsApiKey) Dependencies() terra.Dependencies {
	return aiak.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) LifecycleManagement() *terra.Lifecycle {
	return aiak.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsightsApiKey].
func (aiak *ApplicationInsightsApiKey) Attributes() applicationInsightsApiKeyAttributes {
	return applicationInsightsApiKeyAttributes{ref: terra.ReferenceResource(aiak)}
}

// ImportState imports the given attribute values into [ApplicationInsightsApiKey]'s state.
func (aiak *ApplicationInsightsApiKey) ImportState(av io.Reader) error {
	aiak.state = &applicationInsightsApiKeyState{}
	if err := json.NewDecoder(av).Decode(aiak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aiak.Type(), aiak.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsightsApiKey] has state.
func (aiak *ApplicationInsightsApiKey) State() (*applicationInsightsApiKeyState, bool) {
	return aiak.state, aiak.state != nil
}

// StateMust returns the state for [ApplicationInsightsApiKey]. Panics if the state is nil.
func (aiak *ApplicationInsightsApiKey) StateMust() *applicationInsightsApiKeyState {
	if aiak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aiak.Type(), aiak.LocalName()))
	}
	return aiak.state
}

// ApplicationInsightsApiKeyArgs contains the configurations for azurerm_application_insights_api_key.
type ApplicationInsightsApiKeyArgs struct {
	// ApplicationInsightsId: string, required
	ApplicationInsightsId terra.StringValue `hcl:"application_insights_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReadPermissions: set of string, optional
	ReadPermissions terra.SetValue[terra.StringValue] `hcl:"read_permissions,attr"`
	// WritePermissions: set of string, optional
	WritePermissions terra.SetValue[terra.StringValue] `hcl:"write_permissions,attr"`
	// Timeouts: optional
	Timeouts *applicationinsightsapikey.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsApiKeyAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(aiak.ref.Append("api_key"))
}

// ApplicationInsightsId returns a reference to field application_insights_id of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceAsString(aiak.ref.Append("application_insights_id"))
}

// Id returns a reference to field id of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiak.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aiak.ref.Append("name"))
}

// ReadPermissions returns a reference to field read_permissions of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) ReadPermissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiak.ref.Append("read_permissions"))
}

// WritePermissions returns a reference to field write_permissions of azurerm_application_insights_api_key.
func (aiak applicationInsightsApiKeyAttributes) WritePermissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiak.ref.Append("write_permissions"))
}

func (aiak applicationInsightsApiKeyAttributes) Timeouts() applicationinsightsapikey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsightsapikey.TimeoutsAttributes](aiak.ref.Append("timeouts"))
}

type applicationInsightsApiKeyState struct {
	ApiKey                string                                   `json:"api_key"`
	ApplicationInsightsId string                                   `json:"application_insights_id"`
	Id                    string                                   `json:"id"`
	Name                  string                                   `json:"name"`
	ReadPermissions       []string                                 `json:"read_permissions"`
	WritePermissions      []string                                 `json:"write_permissions"`
	Timeouts              *applicationinsightsapikey.TimeoutsState `json:"timeouts"`
}
