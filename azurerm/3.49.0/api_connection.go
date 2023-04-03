// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apiconnection "github.com/golingon/terraproviders/azurerm/3.49.0/apiconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiConnection creates a new instance of [ApiConnection].
func NewApiConnection(name string, args ApiConnectionArgs) *ApiConnection {
	return &ApiConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiConnection)(nil)

// ApiConnection represents the Terraform resource azurerm_api_connection.
type ApiConnection struct {
	Name      string
	Args      ApiConnectionArgs
	state     *apiConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiConnection].
func (ac *ApiConnection) Type() string {
	return "azurerm_api_connection"
}

// LocalName returns the local name for [ApiConnection].
func (ac *ApiConnection) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [ApiConnection].
func (ac *ApiConnection) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [ApiConnection].
func (ac *ApiConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [ApiConnection] depends_on.
func (ac *ApiConnection) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiConnection].
func (ac *ApiConnection) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [ApiConnection].
func (ac *ApiConnection) Attributes() apiConnectionAttributes {
	return apiConnectionAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [ApiConnection]'s state.
func (ac *ApiConnection) ImportState(av io.Reader) error {
	ac.state = &apiConnectionState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiConnection] has state.
func (ac *ApiConnection) State() (*apiConnectionState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [ApiConnection]. Panics if the state is nil.
func (ac *ApiConnection) StateMust() *apiConnectionState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// ApiConnectionArgs contains the configurations for azurerm_api_connection.
type ApiConnectionArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedApiId: string, required
	ManagedApiId terra.StringValue `hcl:"managed_api_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParameterValues: map of string, optional
	ParameterValues terra.MapValue[terra.StringValue] `hcl:"parameter_values,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *apiconnection.Timeouts `hcl:"timeouts,block"`
}
type apiConnectionAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_api_connection.
func (ac apiConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_connection.
func (ac apiConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// ManagedApiId returns a reference to field managed_api_id of azurerm_api_connection.
func (ac apiConnectionAttributes) ManagedApiId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("managed_api_id"))
}

// Name returns a reference to field name of azurerm_api_connection.
func (ac apiConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// ParameterValues returns a reference to field parameter_values of azurerm_api_connection.
func (ac apiConnectionAttributes) ParameterValues() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("parameter_values"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_connection.
func (ac apiConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_api_connection.
func (ac apiConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

func (ac apiConnectionAttributes) Timeouts() apiconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apiconnection.TimeoutsAttributes](ac.ref.Append("timeouts"))
}

type apiConnectionState struct {
	DisplayName       string                       `json:"display_name"`
	Id                string                       `json:"id"`
	ManagedApiId      string                       `json:"managed_api_id"`
	Name              string                       `json:"name"`
	ParameterValues   map[string]string            `json:"parameter_values"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Tags              map[string]string            `json:"tags"`
	Timeouts          *apiconnection.TimeoutsState `json:"timeouts"`
}
